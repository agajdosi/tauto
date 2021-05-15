package generate

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/lucasjones/reggen"
	"github.com/spf13/viper"
)

func FromTemplateByName(templateName string) string {
	templates := viper.GetStringSlice(templateName)
	x := rand.Intn(len(templates))

	result, err := reggen.Generate(templates[x], 5)
	if err != nil {
		fmt.Println("Error while generating tweet from template:")
		log.Fatal(err)
	}

	return result
}

func FromTemplate(template string) string {
	result, err := reggen.Generate(template, 5)
	if err != nil {
		fmt.Println("Error while generating tweet from template:")
		log.Fatal(err)
	}

	return result
}
