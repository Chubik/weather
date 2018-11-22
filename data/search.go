package data

import (
	"strings"
)

//Search method for searches in string
//TODO: should change search method for performance
func Search(what string) []*CityItem {
	var resp []*CityItem

	for _, city := range Cities {
		if strings.Contains(strings.ToLower(city.Name), what) {
			resp = append(resp, city)
		}
	}
	return resp
}
