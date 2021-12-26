package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python"

type extractedJob struct {
	id       string
	title    string
	company  string
	location string
	salary   string
	summary  string
}

func main() {
	mainC := make(chan []extractedJob)
	var jobs []extractedJob
	pageNums := GetPageNums(baseURL)
	for i := 0; i < pageNums; i++ {
		go GetPage(i, mainC)
	}

	for i := 0; i < pageNums; i++ {
		job := <-mainC
		jobs = append(jobs, job...)
	}
	fmt.Println("Finish indeed crawling")

}

func GetPage(page int, mainC chan []extractedJob) {
	c := make(chan extractedJob)
	var jobs []extractedJob

	url := baseURL + "&start=" + strconv.Itoa(page*10)
	fmt.Println("Requesting... ", url)
	res, err := http.Get(url)
	checkRequest(err)
	checkStatus(res)

	// 채용 정보 데이터 얻기 -> go routine으로 최적화하기
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkRequest(err)
	tapItem := doc.Find(".tapItem")
	tapItem.Each(func(i int, card *goquery.Selection) {
		go ExtractJob(card, c)
	})

	// go routine 에서 main으로 channel 통해 데이터 전송 받기
	for i := 0; i < tapItem.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func ExtractJob(card *goquery.Selection, c chan extractedJob) {
	id, _ := card.Attr("data-jk")
	title := card.Find(".jobTitle").Text()
	company := card.Find(".companyName").Text()
	location := card.Find(".companyLocation").Text()
	salary := card.Find(".salary-snippet").Text()
	summary := card.Find(".job-snippet").Text()

	c <- extractedJob{
		id:       id,
		title:    title,
		company:  company,
		location: location,
		salary:   salary,
		summary:  summary}
}

// calculate the number of pages
func GetPageNums(url string) int {
	var pageNums int
	res, err := http.Get(url)
	checkRequest(err)
	checkStatus(res)

	// goquery 이용해서 HTML response 파싱 가능
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkRequest(err)
	pagination := doc.Find(".pagination-list")
	pagination.Each(func(i int, s *goquery.Selection) {
		pageNums = s.Find("a").Length()
	})
	return pageNums
}

// handling request errors
func checkRequest(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatus(res *http.Response) {
	if res.StatusCode >= 400 {
		log.Fatalln("Request failed with StatusCode: ", res.StatusCode)
	}
}
