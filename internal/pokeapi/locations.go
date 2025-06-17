package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/senaphim/pokedexcli/internal/pokecache"
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

func ListLocations(url *string, cache *pokecache.Cache) (Locations20, error) {
	apiUrl := baseurl + "/location-area"
	if url != nil {
		apiUrl = *url
	}

	if data, ok := cache.Get(apiUrl); ok {
		var locations Locations20
		if err := json.Unmarshal(data, &locations); err != nil {
			return Locations20{}, nil
		}
		return locations, nil
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

	dat, _ := io.ReadAll(res.Body)
	cache.Add(*url, dat)
	return locations, nil
}
