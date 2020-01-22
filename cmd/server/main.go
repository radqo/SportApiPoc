package main

import (
	"github.com/radqo/SportApiPoc/adapter/sportclient"
	"github.com/radqo/SportApiPoc/model"
	"github.com/radqo/SportApiPoc/service/api"
	"github.com/radqo/SportApiPoc/service/config"
	"log"
	"net/http"
	"strconv"
	"time"
)

var appConf *model.AppConfiguration
var clientConf *sportclient.Configuration

func init() {

	err := config.ReadConfiguration("../../config/appconfig.json", &appConf)

	if err != nil {
		log.Fatal("configuration error")
	}

	err = config.ReadConfiguration("../../config/sportclientconfig.json", &clientConf)

	if err != nil {
		log.Fatal("configuration error")
	}
}

func main() {

	log.Println("Server start")

	client := &http.Client{
		Timeout: time.Duration(appConf.ClientTimeoutSec) * time.Second,
	}

	s := &api.Server{PlayerInfo: sportclient.NewService(*clientConf, client)}

	s.Run(strconv.Itoa(appConf.Port))

}
