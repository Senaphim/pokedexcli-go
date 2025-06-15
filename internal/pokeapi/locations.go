package pokeapi

import (
	"encoding/json"
	"net/http"
)

type Location struct {
	Name string
	url  string
}

type Locations20 struct {
	count    int
	Next     *string
	Previous *string
	Results  []Location
}

func ListLocations(url *string) (Locations20, error) {
	apiUrl := baseurl + "/location-area"
	if url != nil {
		apiUrl = *url
	}

	res, err := http.Get(apiUrl)
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
