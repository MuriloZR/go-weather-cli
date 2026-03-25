package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"encoding/json"
)

type Response struct {
	CurrentCondition []CurrentCondition `json:"current_condition"`
	NearestArea 	 []NearestArea `json:"nearest_area"`
	Weather 		 []Weather `json:"weather"`
}

type CurrentCondition struct {
	TempC   string `json:"temp_C"`
	FeelC   string `json:"FeelsLikeC"`
	Humity  string `json:"humidity"`
	ObsTime string `json:"localObsDateTime"`
	WeatherDesc []ValueObject `json:"weatherDesc"`
}

type NearestArea struct {
	AreaName []ValueObject `json:"areaName"`
	Country  []ValueObject `json:"country"`
	Region   []ValueObject `json:"region"`
}

type Weather struct {
	Date string `json:"date"`
	MaxTempC string `json:"maxtempC"`
	MinTempC string `json:"mintempC"`
	Hourly []Hourly `json:"hourly"`
}

type Hourly struct {
	ChanceRain string `json:"chanceofrain"`
}

type ValueObject struct {
	Value string `json:"value"`
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
	current := response.CurrentCondition[0]
	area := response.NearestArea[0]
	weather := response.Weather[0]

	city := area.AreaName[0].Value
	// region := area.Region[0].Value
	country := area.Country[0].Value
	desc := current.WeatherDesc[0].Value

	fmt.Println()
	fmt.Println("╔═════════════════════════════════════════╗")
	fmt.Printf("║  %s, %s\n", city, country)
	fmt.Println("╠═════════════════════════════════════════╣")
	fmt.Printf("║  %s\n", desc)
	fmt.Printf("║  Temperatura:\n")
	fmt.Printf("║    Atual: %s°C Máxima: %s°C Mínima: %s°C\n", current.TempC, weather.MaxTempC, weather.MinTempC)
	fmt.Printf("║  Sensação:      %s°C\n", current.FeelC)
	fmt.Printf("║  Umidade:       %s%%\n", current.Humity)
	// fmt.Println("╠═════════════════════════════════════════╣")
	fmt.Println("╚═════════════════════════════════════════╝")
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Você deve especificar qual cidade quer buscar o clima\n")
		os.Exit(1)
	}

	response := fetchWeather(os.Args[1])
	printWeather(response)
}