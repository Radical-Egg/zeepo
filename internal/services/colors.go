package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)
type ColorScheme struct {
	// Return type is a nested array
	Result [][]int `json:"result"`
}

func ConvRGBToHex(colors ...int) string {
	result := ""
	for _, color := range colors {
		result += strconv.FormatInt(int64(color), 16)
	}
	return result
}

func GetColorPallete() ColorScheme {
	url := "http://colormind.io/api/"

	var jsonStr = []byte(`{"model":"default"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(
			`Unable to pull from the colormind API. 
		You probably aren't connected to the internet, or the API is down :c`)
		log.Fatalln(err)
	}
	defer res.Body.Close()

	var colors ColorScheme
	derr := json.NewDecoder(res.Body).Decode(&colors)

	if derr != nil {
		log.Fatalln(derr)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalln(res.Status)
	}


	return colors
}