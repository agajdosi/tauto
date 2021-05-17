package generate

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/lucasjones/reggen"
	"github.com/spf13/viper"
)

func FromTemplateByName(templateName string) string {
	templates := viper.GetStringSlice(templateName)
	x := rand.Intn(len(templates))

	g, err := reggen.NewGenerator(templates[x])
	if err != nil {
		fmt.Println("Error while generating tweet from template:")
		log.Fatal(err)
	}
	g.SetSeed(time.Now().UTC().UnixNano())

	return g.Generate(5)
}

func FromTemplate(template string) string {
	g, err := reggen.NewGenerator(template)
	if err != nil {
		fmt.Println("Error setting generator.")
		log.Fatal(err)
	}
	g.SetSeed(time.Now().UTC().UnixNano())

	return g.Generate(5)
}
