package generate

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/lucasjones/reggen"
	"github.com/spf13/viper"
)

func FromTemplate(template string) string {
	templates := viper.GetStringSlice(template)
	x := rand.Intn(len(templates))

	result, err := reggen.Generate(templates[x], 5)
	if err != nil {
		fmt.Println("Error while generating tweet from template:")
		log.Fatal(err)
	}

	return result
}
