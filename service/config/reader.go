package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

// ReadConfiguration - read configuration file
func ReadConfiguration(filePath string, destination interface{}) (err error) {

	defer func() {
		if r := recover(); r != nil {
			log.Println(err)
			err = errors.New("Error reading file")
		}
	}()

	log.Println("Reading " + filePath)

	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = json.Unmarshal(content, &destination)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
