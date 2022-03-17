package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readForbidden(file *os.File) []string {
	var bad_words = make([]string, 0)
	fscan := bufio.NewScanner(file)
	for i := 0; fscan.Scan(); i++ {
		bad_words = append(bad_words, fscan.Text())
	}
	return bad_words
}

func checkAndPrint(bad_words []string) {
	var ok bool
	var line string
	var words []string
	inscan := bufio.NewScanner(os.Stdin)
	inscan.Split(bufio.ScanLines)

	fmt.Print("Enter your message here(if you want to exit enter \"exit\"):\n>")
	for inscan.Scan(){
		line = inscan.Text()
		words = strings.Split(line, " ")
		for i := 0; i < len(words); i++ {
			for j := 0; j < len(bad_words); j++ {
				if strings.ToLower(words[i]) == strings.ToLower(bad_words[j]) {
					ok = true
					break
				}
			}
			if ok {
				for k := 0; k < len(words[i]); k++ {
					fmt.Print("*")
				}
			} else if strings.ToLower(words[0]) == "exit" {
				fmt.Println("Bye!")
				return
			} else {
				fmt.Print(words[i])
			}
			ok = false
			fmt.Print(" ")
		}
		fmt.Println()
		fmt.Print(">")
	}
}

func main() {
	var filename string
	var bad_words []string

	fmt.Println("Enter the name of file with forbidden words:")
	fmt.Scan(&filename)
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bad_words = readForbidden(file)
	checkAndPrint(bad_words)
}
