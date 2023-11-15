package main

import (
	"bufio"
	"fmt"
	"product-parser/comparator"
	"product-parser/flags"
	"product-parser/parser"
	"product-parser/picker"
)

type Product struct {
	Name   string `json:"product" csv:"Product"`
	Price  int    `json:"price" csv:"Price"`
	Rating int    `json:"rating" csv:"Rating"`
}

// Greater price compares p instance's price with given instance's price
// Returns the instance which price is greater
func (p Product) GreaterPrice(with Product) Product {
	if p.Price >= with.Price {
		return p
	}

	return with
}

// Greater price compares p instance's price with given instance's rating
// Returns the instance which rating is greater
func (p Product) GreaterRating(with Product) Product {
	if p.Rating >= with.Rating {
		return p
	}

	return with
}

func (p Product) String() string {
	return fmt.Sprintf("Name: %s, Price: %d, Rating: %d", p.Name, p.Price, p.Rating)
}

type ParserComparable[T comparator.Comparable[T]] struct {
	Parser parser.Parser[T]
	file   *picker.File
}

func (p *ParserComparable[T]) GetProduct() {

	count := 0

	scanner := bufio.NewScanner(p.file.FD)
	var currentMaxPrice T
	var currentMaxRating T
	for scanner.Scan() {
		text := scanner.Text()
		product, err := p.Parser.Parse(text)
		if err != nil {
			continue
		}

		currentMaxPrice = currentMaxPrice.GreaterPrice(product)
		currentMaxRating = currentMaxRating.GreaterRating(product)
		count++
	}
	p.file.FD.Close()

	if count == 0 {
		fmt.Println("There is no products :(")
		return
	}

	fmt.Println("Product with highest price: ", currentMaxPrice)
	fmt.Println("Product with highest rating: ", currentMaxRating)

}

func New[T comparator.Comparable[T]](filepath string) ParserComparable[T] {
	file := picker.Open(filepath)
	p := parser.New[T](file.Ext)
	return ParserComparable[T]{
		Parser: p,
		file:   &file,
	}

}

func main() {
	flags.ReadConfig()
	pc := New[Product](flags.Filepath)
	fmt.Println("Started parsing file...")
	pc.GetProduct()
}
