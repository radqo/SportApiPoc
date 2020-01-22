package model

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