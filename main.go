package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"encoding/json"
)

type WeatherDesc struct {
	Value string `json:"value"`
}

type CurrentCondition struct {
	TempC string `json:"temp_C"`
	FeelC string `json:"FeelsLikeC"`
	Humity string `json:"humidity"`
	ObsTime string `json:"localObsDateTime"`
	WeatherDesc []WeatherDesc `json:"weatherDesc"`
}

type AreaName struct {
	Value string `json:"value"`
}

type NearestArea struct {
	AreaName []AreaName `json:"areaName"`
}

type Response struct {
	CurrentCondition []CurrentCondition `json:"current_condition"`
	NearestArea []NearestArea `json:"nearest_area"`
}

func fetchWeather(s string) Response{
	url := fmt.Sprintf("https://wttr.in/%s?format=j1", s)

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
	err = json.Unmarshal(info, &response)
	if err != nil {
		fmt.Printf("Erro na decodificação do JSON\n")
		os.Exit(1)
	}

	return response
}

func printWeather(response Response) {
	fmt.Printf("Tempo em %s:\n", response.NearestArea[0].AreaName[0].Value)
	fmt.Printf("Última Atualização em: %s\n", response.CurrentCondition[0].ObsTime)
	fmt.Printf("%s\n", response.CurrentCondition[0].WeatherDesc[0].Value)
	fmt.Printf("Temperatura: %s °C\n", response.CurrentCondition[0].TempC)
	fmt.Printf("Sensação Térmica: %s °C\n", response.CurrentCondition[0].FeelC)
	fmt.Printf("Umidade: %s%%\n", response.CurrentCondition[0].Humity)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Você deve especificar qual cidade quer buscar o clima\n")
		os.Exit(1)
	}

	response := fetchWeather(os.Args[1])
	printWeather(response)
}