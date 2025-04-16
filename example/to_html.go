package main

import (
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/oslyak/htmldump"
)

func ToHTML() {
	// Generate a slice of structs
	people := make([]Person, 0)
	for i := 0; i < 5; i++ {
		people = append(people, Person{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
			Age:   gofakeit.Number(18, 60),
		})
	}

	pack := make([]*Wolf, 0)

	for i := 0; i < 5; i++ {
		pack = append(pack, &Wolf{
			Animal: Animal{
				Name:    gofakeit.Name(),
				Species: gofakeit.Animal(),
				Age:     gofakeit.Number(1, 10),
			},
			Fur:    gofakeit.Color(),
			Hungry: gofakeit.Bool(),
		})
	}

	person := Person{
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
		Age:   gofakeit.Number(18, 60),
	}

	message := "Inspect variables via map table"

	magickNumber := 42
	isItCool := true
	sql := "SELECT Name, Species, Age, Sex, Fur, Hungry FROM animals WHERE Species = 'wolf'"

	variables := map[string]interface{}{
		"message string ": message,
		"magicNumber int": magickNumber,
		"isItCool bool":   isItCool,
	}

	colours := map[string]string{
		"red":    "#FF0000",
		"green":  "#00FF00",
		"blue":   "#0000FF",
		"pink":   "#FFC0CB",
		"yellow": "#FFFF00",
		"purple": "#800080",
	}

	htmlFile, err := os.Create("example-2.html")
	if err != nil {
		panic(err)
	}
	defer htmlFile.Close()

	err = htmldump.ToHTML(htmlFile, variables, sql, pack, person, people, colours)
	if err != nil {
		panic(err)
	}

	htmlFile.Close()
}
