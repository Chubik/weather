package data

import (
	"strings"
)

//Search method for searches in string
func Search(what string) []*CityItem {
	var resp []*CityItem

	for _, city := range Cities {
		if strings.Contains(strings.ToLower(city.Name), what) {
			// ret, err := json.Marshal(city)
			// if err != nil {
			// 	log.Println("Error parsing city data", err)
			// }

			resp = append(resp, city)
		}
	}
	return resp
}
