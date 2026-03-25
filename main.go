package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Você deve especificar qual cidade quer buscar o clima\n")
		os.Exit(1)
	}
	url := fmt.Sprintf("https://wttr.in/%s?format=j1", os.Args[1])
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro na requisição http\n")
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	info, err := io.ReadAll(resp.Body)

	fmt.Printf(string(info))

	defer resp.Body.Close()
}