package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("decimal sample")

	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875

	fmt.Printf("total.String()      : %s\n", total.String())
	fmt.Printf("total.StringFixed(2): %s\n", total.StringFixed(2))
	fmt.Printf("total.StringFixed(3): %s\n", total.StringFixed(3))

	sum := decimal.NewFromInt(0)
	unit, err := decimal.NewFromString("0.1")
	if err != nil {
		panic(err)
	}

	var sumF float64 = 0
	var unitF float64 = 0.1

	for i := 0; i < 100; i++ {
		sum = sum.Add(unit)
		sumF += unitF
	}

	fmt.Printf("bigdecimal sum: %s\n", sum.String())
	fmt.Println("float64 sum:    ", sumF)
}
