package main

import (
	"fmt"
	"strconv"
)

func init() {
	product := product{
		name:     "Lápis",
		price:    1.79,
		discount: 0.05,
		items: []item{
			{id: "001", value: 2.79493},
			{id: "002", value: 2.793043},
		},
	}

	fmt.Println(product)
	fmt.Println("Price with discount", product.priceWithDiscount())
	fmt.Println("total item values", product.getTotalItemValues())
}

type item struct {
	id    string
	value float64
}

type product struct {
	name     string
	price    float64
	discount float64
	items    []item
}

func (p product) priceWithDiscount() float64 {
	newPrice := p.price * (1 - p.discount)
	priceFormatted := fmt.Sprintf("%.*f", 2, newPrice)

	parsedNumber, _ := strconv.ParseFloat(priceFormatted, 64)
	return parsedNumber

}

func (p product) getTotalItemValues() float64 {
	total := 0.0
	for _, item := range p.items {
		total += float64(item.value)
	}

	convertedTotal, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", total), 64)

	return convertedTotal
}
