package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

//funcao remover
func RemoveWords(t []string, minLenght int) []string {
	mapeamento := make(map[string]bool)
	list := []string{}

	for _, pos := range t {
		if len(pos) >= minLenght {
			mapeamento[pos] = true
		}
		if mapeamento[pos] != true {
			list = append(list, pos)
		}
	}
	return list
}

func main() {

	// definindo os valores
	sString := []string{"Vandriele", "batata", "abacaxi", "escada", "campos do Jordão", "Yoda"}

	// Apresentano o array original
	fmt.Println(sString)

	//pegando o tammanho minimo da palavra para remover
	fmt.Println("Digite qual o tamanho palavra deseja remover?")
	tamanho := bufio.NewScanner(os.Stdin)
	tamanho.Scan()

	tam := tamanho.Text()
	min, erro := strconv.Atoi(tam)
	if erro != nil {
		log.Fatal(erro)
	}

	//chammando a funcao
	saida := RemoveWords(sString, min)

	// mostrando a saida sem os itens maiores que os definidos
	fmt.Println("As palavrass que sobram são:", saida)
}
