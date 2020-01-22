package sportapi

// PlayerInfo - contains player id and basic information
type PlayerInfo struct {
	ID          int
	FirstName   string
	LastName    string
	Nationality string
	Age         int
	Weight      string
	Height      string
}

// PlayerInfoFinder - groups methods for player information search
type PlayerInfoFinder interface {
	FindPlayer(surname string) (palyers []PlayerInfo, err error)
	//GeAPIErrorics(playerID int)
}

// APIError Error returned by api call or internal server error 500
type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return e.Message
}
