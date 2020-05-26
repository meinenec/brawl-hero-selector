package heroes

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

// Hero is a struct definition of a hero
type Hero struct {
	Name string
	Role string
}

// Heroes is a slice of Hero
type Heroes []Hero

var brawlBlacklist = []string{"Cho", "Gall", "Sgt. Hammer", "Abathur"}

// Assign returns a slice of n unique heroes to play
func Assign(n int, hPool string) Heroes {
	heroes := getHeroes()
	rand.Seed(time.Now().UnixNano())
	h := Heroes{}

	for len(h) < n {
		hero := heroes[rand.Intn(len(heroes))]
		if h.accepts(hero, hPool) {
			h = append(h, hero)
		}
	}

	return h
}

func (h *Heroes) accepts(hero Hero, hPool string) bool {
	for _, a := range *h {
		if a == hero {
			return false
		}
	}

	if hPool == "brawl" {
		for _, a := range brawlBlacklist {
			if a == hero.Name {
				return false
			}
		}
	}

	return true
}

func getHeroes() Heroes {
	var h Heroes
	data, err := ioutil.ReadFile("heroes.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &h)
	if err != nil {
		panic(err)
	}
	return h
}
