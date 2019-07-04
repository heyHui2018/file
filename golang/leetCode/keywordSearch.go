package main

import (
	"bufio"
	"fmt"
	"github.com/ngaut/log"
	"io"
	"os"
	"strings"
)

type Question struct {
	Id       string
	Target   string
	KeyWord  string
	Solution string
}

// key-questionId,value-Question
var questionMap map[string]*Question

func logInint() {
	log.SetLevelByString("info")
}

func dataInit() {
	questionMap = make(map[string]*Question)
	question := new(Question)

	file, err := os.Open("text.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	for {
		data, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		list := strings.Split(string(data), " ")
		if len(list) == 4 {
			question.Id = list[0]
			question.Target = list[1]
			question.KeyWord = list[2]
			question.Solution = list[3]
			log.Infof("dataInit,question = %+v", question)
		}
	}
}

// todo 接口1，根据questionId查询
// todo 接口2，根据keyword匹配查询
func main() {
	logInint()
	dataInit()
	fmt.Println("hello world")
}
