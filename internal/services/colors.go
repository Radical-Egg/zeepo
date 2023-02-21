package services

import (
	"bytes"
	"encoding/json"
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

func GetColorPallete() ColorScheme{
	url := "http://colormind.io/api/"

	var jsonStr = []byte(`{"model":"default"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var colors ColorScheme
	derr := json.NewDecoder(res.Body).Decode(&colors)

	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}


	return colors
}