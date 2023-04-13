package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/novalagung/gubrak/v2"
)

func main() {

	for {
		randomWater := gubrak.RandomInt(0, 100)
		randomWind := gubrak.RandomInt(0, 100)

		GetDataJSON(randomWater, randomWind)

		water := statusWater(randomWater)
		wind := statusWind(randomWater)

		fmt.Println("status water: ", water)
		fmt.Println("status wind: ", wind)

		fmt.Println(strings.Repeat("-", 20))
		time.Sleep(15 * time.Second)
	}

}

func statusWater(n int) string {
	print := ""
	if n <= 5 {
		print = "aman"
	} else if n >= 6 && n <= 8 {
		print = "siaga"
	} else if n > 8 {
		print = "bahaya"
	}

	return print
}

func statusWind(n int) string {
	print := ""
	if n <= 6 {
		print = "aman"
	} else if n >= 7 && n <= 15 {
		print = "siaga"
	} else if n > 15 {
		print = "bahaya"
	}

	return print
}

func GetDataJSON(countWater, countWind int) {
	data := map[string]interface{}{
		"water": countWater,
		"wind":  countWind,
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

}
