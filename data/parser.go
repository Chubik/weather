package data

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
)

//Init initialization data to slice of cities
func Init(file string) error {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return err
	}
	var parse []*CityParse
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&parse)
	if err != nil {
		return err
	}
	for _, i := range parse {
		ct := strings.Split(i.Name, ",")
		c := &CityItem{
			ID:   i.ID,
			Name: ct[0],
			Code: ct[1],
		}
		Cities = append(Cities, c)
	}
	sort.Sort(Cts(Cities))
	return nil
}

func (a Cts) Len() int           { return len(a) }
func (a Cts) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a Cts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
