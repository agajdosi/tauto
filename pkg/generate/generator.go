package generate

import (
	"math/rand"

	"github.com/lucasjones/reggen"
)

func StupidQuestion() string {
	templates := []string{
		"To nemyslíte (fakt |doopravdy |skutečně |)vážně, (panebože |ježiši |kristepane ||||)že ne\\?(|||!)",
		"Myslíte to (opravdu |skutečně |)vážně\\?(|||!)",
		"(To|Tohlencto|Todle|Tohle|Toto) (jako |)myslíte (fakt |doopravdy |)vážně\\?(|||!)",
		"Co (přesně|konkrétně) tím (vlastně |)(myslíte|naznačujete|chcete říct)\\?(|||!)",
	}

	x := rand.Intn(len(templates))
	result, _ := reggen.Generate(templates[x], 5)

	return result
}
