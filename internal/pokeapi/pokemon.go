package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/senaphim/pokedexcli/internal/pokecache"
)

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	}
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"forms"`
	GameIndicies []struct {
		GameIndex int
		Version   struct {
			Name int    `json:"name"`
			Url  string `json:"url"`
		} `json:"version"`
	} `json:"game_indicies"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	Id                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"move_learn_method"`
			Order        *string `json:"order"`
			VersionGroup []struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string `json:"name"`
	Order         int    `json:"order"`
	PastAbilities []struct {
		Abilities []struct {
			Ability  *string `json:"ability"`
			IsHidden bool    `json:"is_hidden"`
			Slot     int     `json:"slot"`
		} `json:"abilities"`
		Generation struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"generation"`
	}
	PastTypes []struct {
		Generation struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"generation"`
		Types []struct {
			Slot string `json:"slot"`
			Type struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
	Species struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      *string `json:"back_default"`
		BackFemale       *string `json:"back_female"`
		BackShiny        *string `json:"back_shiny"`
		BackShinyFemale  *string `json:"back_shiny_female"`
		FrontDefault     *string `json:"front_default"`
		FrontFemale      *string `json:"front_female"`
		FrontShiny       *string `jsont:"front_shiny"`
		FrontShinyFemale *string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault *string `json:"front_default"`
				FrontFemale  *string `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     *string `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       *string `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault *string `json:"front_default"`
				FrontShiny   *string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      *string `json:"back_default"`
				BackFemale       *string `json:"back_female"`
				BackShiny        *string `json:"back_shiny"`
				BackShinyFemale  *string `json:"back_shiny_female"`
				FrontDefault     *string `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       *string `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      *string `json:"back_default"`
					BackGray         *string `json:"back_gray"`
					BackTransparent  *string `json:"back_transparent"`
					FrontDefault     *string `json:"front_default"`
					FrontGray        *string `json:"front_gray"`
					FrontTransparent *string `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      *string `json:"back_default"`
					BackGray         *string `json:"back_gray"`
					BackTransparent  *string `json:"back_transparent"`
					FrontDefault     *string `json:"front_default"`
					FrontGray        *string `json:"front_gray"`
					FrontTransparent *string `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationII struct {
				Crystal struct {
					BackDefault           *string `json:"back_default"`
					BackShiny             *string `json:"back_shiny"`
					BackShinyTransparent  *string `json:"back_shiny_transparent"`
					BackTransparent       *string `json:"back_transparent"`
					FrontDefault          *string `json:"front_default"`
					FrontShiny            *string `json:"front_shiny"`
					FrontShinyTransparent *string `json:"front_shiny_transparent"`
					FrontTransparent      *string `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault           *string `json:"back_default"`
					BackShiny             *string `json:"back_shiny"`
					BackShinyTransparent  *string `json:"back_shiny_transparent"`
					BackTransparent       *string `json:"back_transparent"`
					FrontDefault          *string `json:"front_default"`
					FrontShiny            *string `json:"front_shiny"`
					FrontShinyTransparent *string `json:"front_shiny_transparent"`
					FrontTransparent      *string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault           *string `json:"back_default"`
					BackShiny             *string `json:"back_shiny"`
					BackShinyTransparent  *string `json:"back_shiny_transparent"`
					BackTransparent       *string `json:"back_transparent"`
					FrontDefault          *string `json:"front_default"`
					FrontShiny            *string `json:"front_shiny"`
					FrontShinyTransparent *string `json:"front_shiny_transparent"`
					FrontTransparent      *string `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIII struct {
				Emerald struct {
					FrontDefault *string `json:"front_default"`
					FrontShiny   *string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  *string `json:"back_default"`
					BackShiny    *string `json:"back_shiny"`
					FrontDefault *string `json:"front_default"`
					FrontShiny   *string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  *string `json:"back_default"`
					BackShiny    *string `json:"back_shiny"`
					FrontDefault *string `json:"front_default"`
					FrontShiny   *string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIV struct {
				DiamondPearl struct {
					BackDefault      *string `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        *string `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      *string `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        *string `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      *string `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        *string `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      *string `json:"back_default"`
						BackFemale       *string `json:"back_female"`
						BackShiny        *string `json:"back_shiny"`
						BackShinyFemale  *string `json:"back_shiny_female"`
						FrontDefault     *string `json:"front_default"`
						FrontFemale      *string `json:"front_female"`
						FrontShiny       *string `json:"front_shiny"`
						FrontShinyFemale *string `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      *string `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        *string `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVI struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVII struct {
				Icons struct {
					FrontDefault *string `json:"front_default"`
					FrontFemale  *string `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationVIII struct {
				Icons struct {
					FrontDefault *string `json:"front_default"`
					FrontFemale  *string `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func GetPokemon(mon string, cache *pokecache.Cache) (Pokemon, error) {
	apiUrl := baseurl + "/pokemon/" + mon
	var monData Pokemon

	if data, ok := cache.Get(apiUrl); ok {
		if err := json.Unmarshal(data, &monData); err != nil {
			fmtErr := fmt.Errorf("Error encountered when umarshalling from cache: %v", err)
			return Pokemon{}, fmtErr
		}
		return monData, nil
	}

	res, err := http.Get(apiUrl)
	if err != nil {
		fmtErr := fmt.Errorf("Error when getting data from api: %v", err)
		return Pokemon{}, fmtErr
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		fmtErr := fmt.Errorf("Error when streaming body to io: %v", err)
		return Pokemon{}, fmtErr
	}

	if err := json.Unmarshal(dat, &monData); err != nil {
		fmtErr := fmt.Errorf("Error when umarshalling json: %v", err)
		return Pokemon{}, fmtErr
	}

	cache.Add(apiUrl, dat)
	return monData, nil
}
