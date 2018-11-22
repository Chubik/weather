package data

//CityParse struct for load cities from datafile
type CityParse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

//CityItem type for one city
type CityItem struct {
	ID      int
	Name    string
	Country string
}

type Cts []*CityItem

var (
	Cities Cts
)
