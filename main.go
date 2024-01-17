package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

func main() {
	var q string
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	res, err := http.Get("https://api.weatherapi.com/v1/current.json?key=0f756d64873242bcb0d125233241701&q=" + q + "")
	if err != nil {
		log.Fatalln("Server Not Found 404")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("Enter The Location Args")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("Not Getting the body")
	}
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current := weather.Location, weather.Current
	fmt.Printf("%s\n", location.Country)
	fmt.Printf("%s\n", location.Region)
	fmt.Printf("%.0fc\n", current.TempC)
}
