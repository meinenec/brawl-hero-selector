package heroes

import (
	"testing"
)

func TestHeroes_Accepts(t *testing.T) {
	testH := Heroes{
		Hero{
			Name: "Brightwing",
			Role: "Healer",
		},
		Hero{
			Name: "Malfurion",
			Role: "Healer",
		},
	}

	testHero := Hero{
		Name: "Dehaka",
		Role: "Bruiser",
	}

	testHeroInvalid := Hero{
		Name: "Brightwing",
		Role: "Healer",
	}

	testHeroInvalidBrawl := Hero{
		Name: "Cho",
		Role: "Tank",
	}

	tests := []struct {
		heroes Heroes
		hero   Hero
		got    bool
		want   bool
	}{
		{
			heroes: testH,
			hero:   testHero,
			got:    testH.accepts(testHero, "all"),
			want:   true,
		},
		{
			heroes: testH,
			hero:   testHero,
			got:    testH.accepts(testHero, "brawl"),
			want:   true,
		},
		{
			heroes: testH,
			hero:   testHeroInvalidBrawl,
			got:    testH.accepts(testHeroInvalidBrawl, "brawl"),
			want:   false,
		},
		{
			heroes: testH,
			hero:   testHeroInvalid,
			got:    testH.accepts(testHeroInvalid, "all"),
			want:   false,
		},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("Expected %v got: %v with heroes %v and hero %v", test.want, test.got, test.heroes, test.hero)
		}
	}
}
