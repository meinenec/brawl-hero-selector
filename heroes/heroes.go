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

// Assign returns a slice of n unique heroes to play
func Assign(n int) Heroes {
	heroes := getHeroes()
	rand.Seed(time.Now().UnixNano())
	h := Heroes{}

	for len(h) < n {
		hero := heroes[rand.Intn(len(heroes))]
		if !h.contains(hero) {
			h = append(h, hero)
		}
	}

	return h
}

func (h *Heroes) contains(hero Hero) bool {
	for _, a := range *h {
		if a == hero {
			return true
		}
	}
	return false
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
