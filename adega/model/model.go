package model

type Wine struct {
	ID 			string 		`json: "id,omitempty"`
	Name 		string 		`json:"name,omitempty"`
	Country		string 		`json:"country,omitempty"`
	Description	string		`json:"description,omitempty"`
	Amount		int			`json:"amount,omitempty"`
	Year		int			`json:"year,omitempty"`
}

var Wines []Wine
