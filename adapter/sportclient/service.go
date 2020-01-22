package sportclient

import (
	"encoding/json"
	"github.com/radqo/SportApiPoc/model"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPClient - interface used for http client
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

type service struct {
	client HTTPClient
	conf   Configuration
}

// NewService - creates new instance of service
func NewService(c Configuration, client HTTPClient) model.PlayerInfoFinder {
	s := &service{
		conf:   c,
		client: client,
	}
	return s
}

func (s *service) FindPlayer(surname string) (playerInfo []model.PlayerInfo, apiErr *model.APIError) {

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			apiErr = &model.APIError{Code: 500, Message: "Error in api call"}
			playerInfo = nil
		}
	}()

	url := s.conf.APIURL + "/players/search/" + surname

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err.Error())
		return nil, &model.APIError{Code: 500, Message: "Error in request creation"}
	}

	req.Header.Add("x-rapidapi-key", s.conf.APIKEY)

	res, err := s.client.Get(url)

	if err != nil {
		log.Println(err.Error())
		return nil, &model.APIError{Code: 500, Message: "Error in http call"}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err.Error())
		return nil, &model.APIError{Code: 500, Message: "Error reading body"}
	}

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		res, err := parseResult(body)

		if err != nil {
			return nil, &model.APIError{Code: 500, Message: err.Error()}
		}

		return res, nil
	}

	return nil, &model.APIError{Code: res.StatusCode, Message: "External api error"}
}

type responseModel struct {
	Api struct {
		Results int
		Players []struct {
			Player_id   int
			Firstname   string
			Lastname    string
			Age         int
			Nationality string
			Height      string
			Weight      string
		}
	}
}

func parseResult(body []byte) (playerInfo []model.PlayerInfo, err error) {

	x := responseModel{}

	err = json.Unmarshal(body, &x)

	if err != nil {
		return nil, &model.APIError{Code: 500, Message: "Unknown json format"}
	}

	info := []model.PlayerInfo{}

	for a := 0; a < x.Api.Results; a++ {
		p := x.Api.Players[a]
		info = append(info, model.PlayerInfo{
			ID:          p.Player_id,
			FirstName:   p.Firstname,
			LastName:    p.Lastname,
			Age:         p.Age,
			Nationality: p.Nationality,
			Height:      p.Height,
			Weight:      p.Weight,
		})
	}

	return info, nil
}
