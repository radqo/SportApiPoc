package sportapi

// Configuration of sport api
type Configuration struct {
	APIURL string // ApiUrl - root url for api calls
	APIKEY string // ApiKey - api key
}


// CreateDemoConfiguration - creates configuration for demo server
func CreateDemoConfiguration() Configuration {
	return Configuration{
		APIURL: "http://www.api-football.com/demo/api/v2",
		APIKEY: "SIGN-UP-FOR-KEY",
	}
}