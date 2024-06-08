package main

import (
	"bufio"
	"fmt"
	"os"
)

const NOME_FILE string = "arquivo.txt"

func main() {
	f, err := os.Create(NOME_FILE)

	if err != nil {
		panic(err)
	}

	//tamanho, err := f.WriteString("Olá mundo!")

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo!")) // Ideal para escrever dados
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso, tamanho: %d bytes\n", tamanho)
	f.Close()

	arquivo, err := os.ReadFile(NOME_FILE)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conteúdo do arquivo:", string(arquivo))

	// Lendo um arquivo sem carregar todo ele de uma só vez
	// Abre o arquivo, sem carregar todo o conteúdo
	dados, err := os.Open(NOME_FILE)
	if err != nil {
		panic(err)
	}
	dados.Close()

	reader := bufio.NewReader(dados)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// Removendo arquivo
	err = os.Remove(NOME_FILE)
	if err != nil {
		panic(err)
	}
}
