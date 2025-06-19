package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/senaphim/pokedexcli/internal/pokecache"
)

type LocationDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func ExploreLocation(location string, cache *pokecache.Cache) (LocationDetails, error) {
	apiUrl := baseurl + "/location-area/" + location
	var details LocationDetails

	if data, ok := cache.Get(apiUrl); ok {
		if err := json.Unmarshal(data, &details); err != nil {
			fmtErr := fmt.Errorf("Encountered error when unmashalling cache: %v",
				err)
			return LocationDetails{}, fmtErr
		}
		return details, nil
	}

	res, err := http.Get(apiUrl)
	if err != nil {
		fmtErr := fmt.Errorf("Encountered error with get from api: %v", err)
		return LocationDetails{}, fmtErr
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		fmtErr := fmt.Errorf("Encountered error with io steaming response body: %v", err)
		return LocationDetails{}, fmtErr
	}

	if err := json.Unmarshal(dat, &details); err != nil {
		fmtErr := fmt.Errorf("Encountered error unmarshalling json: %v", err)
		return LocationDetails{}, fmtErr
	}

	cache.Add(apiUrl, dat)
	return details, nil
}
