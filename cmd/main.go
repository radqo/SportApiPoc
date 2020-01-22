package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
	"../adapter/sportclient"

)

func main() {

	fmt.Println("Test of demo sport client")

	if len(os.Args) == 1 {
		fmt.Println("Provide player surname as argument")
		return
	}

	surname := os.Args[1]

	fmt.Println("Find player : "+ surname)

	c := sportclient.CreateDemoConfiguration()

	client := &http.Client{
	 	Timeout: time.Duration(30) * time.Second,
	 }

	service := sportclient.NewService(c, client)

	x, err := service.FindPlayer(surname)

	if err!=nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(x)
	}

}