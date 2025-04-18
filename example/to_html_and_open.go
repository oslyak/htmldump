package main

import (
	"github.com/oslyak/htmldump"

	"github.com/brianvoe/gofakeit/v7"
)

type Person struct {
	Name  string
	Email string
	Age   int
}

type Sex string

const (
	Male   Sex = "Male"
	Female Sex = "Female"
)

type Animal struct {
	Name    string
	Species string
	Age     int
	Sex     Sex
}

type Wolf struct {
	Animal
	Fur    string
	Hungry bool
}

func ToHTMLAndOpen() {
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

	htmldump.ToHTMLAndOpen("example-1.html", variables, sql, pack, person, people, colours)
}
