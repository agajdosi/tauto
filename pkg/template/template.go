package template

import (
	"github.com/lucasjones/reggen"
)

//Generate generates a string from provided template.
func Generate(template string) (string, error) {

	generator, err := reggen.NewGenerator(template)
	if err != nil {
		return "", err
	}

	text := generator.Generate(10)

	return text, nil
}
