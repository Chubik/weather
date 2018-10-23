package data

//CityParse struct for load cities from datafile
type CityParse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//CityItem type for one city
type CityItem struct {
	ID   int
	Name string
	Code string
}

type Cts []*CityItem

var (
	Cities Cts
)
