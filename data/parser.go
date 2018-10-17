package data

type CityItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	Cities []*CityItem
)

func loadCities() {

}
