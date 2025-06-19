package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/senaphim/pokedexcli/internal/pokecache"
)

type NamedObject struct {
	Name string
	Url  string
}

type Rate struct {
	rate    int
	version NamedObject
}

type EncounterMethod struct {
	method         NamedObject
	versionDetails []Rate
}

type Language struct {
	language NamedObject
	name     string
}

type EncounterDetails struct {
	chance          int
	conditionValues []any
	maxLevel        int
	method          NamedObject
	minLevel        int
}

type Encounters struct {
	encounterDetails []EncounterDetails
	maxChance        int
	version          NamedObject
}

type Pokemon struct {
	Pokemon        NamedObject
	versionDetails []Encounters
}

type LocationDetails struct {
	encounterMethodRates []EncounterMethod
	gameIndex            int
	id                   int
	location             NamedObject
	name                 string
	names                []Language
	Encounters           []Pokemon
}

func ExploreLocation(location string, cache *pokecache.Cache) (LocationDetails, error) {
	apiUrl := baseurl + "/location-area/" + location
	fmt.Println(fmt.Sprintf("Api target: %v", apiUrl))
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
	// decoder := json.NewDecoder(res.Body)
	// if err := decoder.Decode(&details); err != nil {
	// 	fmtErr := fmt.Errorf("Ecountered error decoding response body: %v", err)
	// 	return LocationDetails{}, fmtErr
	// }
	// fmt.Println("Decoded successfully")

	// No error handling needed as error wouuld be thrown at json decode
	cache.Add(apiUrl, dat)
	return details, nil
}
