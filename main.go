package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type PokemonProduct struct {
	url, image, name, price string
}

func main() {
	fmt.Println("hello there bitches, bros and non binary hoes")

	var pokemonProducts []PokemonProduct

	c := colly.NewCollector()

	c.Visit("https://scrapeme.live/shop/")

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting the target DADDY! ", r.URL)
	})

	// Scraping logic

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}

		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})
	// making a the CSVEEEEE file

	file, err := os.Create("productsinfo.csv")
	if err != nil {
		log.Fatal("Failed to create the file ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	header := []string{"url", "image", "name", "price"}
	writer.Write(header)

	for _, pokemonProduct := range pokemonProducts {
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}
		writer.Write(record)
	}
	defer writer.Flush()
}
