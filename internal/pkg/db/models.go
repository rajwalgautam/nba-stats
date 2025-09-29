package db

type Team struct {
	FullName     string `json:"full_name"`
	Abbreviation string `json:"abbreviation"` // TODO: add abbreviation validation
	CityState    string `json:"city_state"`
}
