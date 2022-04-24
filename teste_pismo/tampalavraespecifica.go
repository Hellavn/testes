package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var qtdrepeticao int

func countWords(st string, p string) map[string]int {

	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {

		if word == p {
			qtdrepeticao++
		}
	}
	return wc
}

// Main function
func main() {
	fmt.Println("Digite a frase desejada!")
	texto := bufio.NewScanner(os.Stdin)
	texto.Scan()

	fmt.Println("Digite qual palavra deseja procurar!")
	palavra := bufio.NewScanner(os.Stdin)
	palavra.Scan()

	input := texto.Text()
	p := palavra.Text()
	countWords(input, p)
	fmt.Printf("A palavra %s foi  repetida %d vezes", p, qtdrepeticao)
}
