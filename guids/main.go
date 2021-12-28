package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

const fileName string = "guids.txt"

func main() {
	generateData(fileName)
}

func generateData(fileName string) {
	c := make(chan uuid.UUID)

	go loadData(fileName, c)

	// 그런데 데이터가 10개라는 것을 어떻게 미리 알고 for loop를 돌까..? go routine 함수를 통해서 알아낼 순 없을까..?
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

}

func loadData(fileName string, c chan<- uuid.UUID) {
	file, err := os.Open("./data/" + fileName)
	checkErr(err)

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		guid, err := uuid.Parse(line) // UUID form을 디코딩하는 함수
		checkErr(err)

		c <- guid
	}
}

// handle error
func checkErr(err error) {
	if err != nil {
		log.Fatalln("Failed to open file")
	}
}
