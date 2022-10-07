package main

import (
	"fmt"
	"net/http"
	//"github.com/stianeikeland/go-rpio/v4"
)

func main() {

	// Reference: https://github.com/stianeikeland/go-rpio
	//https://levelup.gitconnected.com/consuming-a-rest-api-using-golang-b323602ba9d8
	//fmt.Println("Calling API...")

	client := &http.Client{}
	apiKey := "5aRdCP324H39P1290TR4K4PxnayTE8G19C5jiAg9"
	req, err := http.NewRequest("POST", "https://7nbx4tfpk5.execute-api.ap-southeast-2.amazonaws.com/prod", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

}
