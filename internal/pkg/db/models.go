package db

type Team struct {
	FullName     string `json:"full_name"`
	Abbreviation string `json:"abbreviation"`
	CityState    string `json:"city_state"`
}
