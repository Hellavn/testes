package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var qtdrepeticao int

func countWords(st string) map[string]int {

	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

// Main function
func main() {
	fmt.Println("Digite a frase desejada!")
	texto := bufio.NewScanner(os.Stdin)
	texto.Scan()

	input := texto.Text()

	for index, element := range countWords(input) {
		fmt.Println(index, "=", element)
	}
}
