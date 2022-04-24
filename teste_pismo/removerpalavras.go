package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var res string

func RemoveWords(st string, minLenght int) map[string]int {

	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {

		length := len(word)
		if length >= minLenght {
			st = strings.ReplaceAll(st, word, "")
		}

	}
	res = st
	return wc
}

// Main function
func main() {
	fmt.Println("Digite a frase desejada!")
	texto := bufio.NewScanner(os.Stdin)
	texto.Scan()

	fmt.Println("Digite qual o tamanho palavra deseja remover")
	tamanho := bufio.NewScanner(os.Stdin)
	tamanho.Scan()

	input := texto.Text()
	tam := tamanho.Text()
	min, erro := strconv.Atoi(tam)
	if erro != nil {
		log.Fatal(erro)
	}
	RemoveWords(input, min)
	fmt.Println(res)

}
