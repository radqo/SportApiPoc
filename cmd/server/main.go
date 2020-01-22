package main

import (
	"github.com/radqo/SportApiPoc/adapter/sportclient"
	"github.com/radqo/SportApiPoc/service/api"
	"log"
	"net/http"
	"time"
)

func main() {

	log.Println("Server start")

	client := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}

	c := sportclient.CreateDemoConfiguration()

	s := &api.Server{PlayerInfo: sportclient.NewService(c, client)}

	s.Run("50300")

}
