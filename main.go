package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"encoding/json"
)

type CurrentCondition struct {
	TempC string `json:"temp_C"`
}

type Response struct {
	CurrentCondition []CurrentCondition `json:"current_condition"`
}

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

	defer resp.Body.Close()
	info, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro na leitura da resposta da requisição\n")
		os.Exit(1)
	}

	var response Response
	json.Unmarshal(info, &response)

	fmt.Printf("Temperatura: %s °C\n", response.CurrentCondition[0].TempC)
}