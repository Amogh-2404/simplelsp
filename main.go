package main

import (
	"bufio"
	"educationalLSP/rpc"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hi")
	logger := getLogger("/Users/r.amogh/projects/simplelsp/log.txt")
	logger.Println("Hi! I started!")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}
func handleMessage(_ any) {

}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didn't give me a good file")
	}

	return log.New(logfile, "[educationalLSP]", log.Ldate|log.Ltime|log.Lshortfile)
}
