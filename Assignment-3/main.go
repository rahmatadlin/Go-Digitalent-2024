package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	status := generateStatus()
	saveStatusToFile(status)

	data := getStatusFromFile()

	waterStatus := getStatusLabel(data.Water, 8, 6)
	windStatus := getStatusLabel(data.Wind, 15, 7)

	fmt.Println("Water Level:", data.Water, ", Status:", waterStatus)
	fmt.Println("Wind Speed:", data.Wind, ", Status:", windStatus)
}

func generateStatus() Status {
	rand.Seed(time.Now().UnixNano())
	return Status{
		Water: rand.Intn(100),
		Wind:  rand.Intn(100),
	}
}

func saveStatusToFile(status Status) {
	file, err := os.Create("status.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonData, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(string(jsonData))
	if err != nil {
		panic(err)
	}
}

func getStatusFromFile() Status {
	dataFile, err := os.ReadFile("status.json")
	if err != nil {
		panic(err)
	}

	var status Status
	err = json.Unmarshal(dataFile, &status)
	if err != nil {
		panic(err)
	}

	return status
}

func getStatusLabel(value, highThreshold, mediumThreshold int) string {
	if value > highThreshold {
		return "Danger"
	} else if value >= mediumThreshold {
		return "Alert"
	}
	return "Safe"
}
