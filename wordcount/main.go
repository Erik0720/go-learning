package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Bitte geben Sie eine Datei an!")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rowCnt int
	var wordCnt int
	var charCnt int

	for scanner.Scan() {
		rowCnt++

		row := scanner.Text()
		charCnt += len(row)

		words := strings.Fields(row)
		wordCnt += len(words)
	}

	fmt.Printf("Die Datei %s hat %d Zeilen, %d Wörtern und %d Zeichen.", os.Args[1], rowCnt, wordCnt, charCnt)
}
