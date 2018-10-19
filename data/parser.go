package data

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

//CityItem type for one city
type CityItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Cts []*CityItem

var (
	Cities Cts
)

//Init initialization data to slice of cities
func Init(file string) error {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&Cities)
	sort.Sort(Cts(Cities))
	fmt.Println()
	return nil
}

func (a Cts) Len() int           { return len(a) }
func (a Cts) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a Cts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
