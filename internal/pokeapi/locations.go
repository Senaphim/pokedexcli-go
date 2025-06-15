package pokeapi

import (
	"encoding/json"
	"net/http"
)

type Location struct {
	name string
	url  string
}

type Locations20 struct {
	count    int
	next     *string
	previous *string
	results  []Location
}

func ListLocations(url string) (Locations20, error) {
	res, err := http.Get(url)
	if err != nil {
		return Locations20{}, err
	}
	defer res.Body.Close()

	var locations Locations20
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return Locations20{}, err
	}

	return locations, nil
}
