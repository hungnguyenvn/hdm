package hdm

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"fmt"
)

func fetchSource_1() []HotelV1 {
	response, err := http.Get("https://api.myjson.com/bins/gdmqa")
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var hotels []HotelV1
	err = json.Unmarshal(responseData, &hotels)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	return hotels
}

func fetchSource_2() []HotelV2 {
	response, err := http.Get("https://api.myjson.com/bins/1fva3m")
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var hotels []HotelV2
	err = json.Unmarshal(responseData, &hotels)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}

	return hotels
}

func fetchSource_3() []HotelV3 {
	response, err := http.Get("https://api.myjson.com/bins/j6kzm")
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var hotels []HotelV3
	err = json.Unmarshal(responseData, &hotels)

	if err != nil {
		fmt.Print(err.Error())
		return nil
	}

	return hotels
}
