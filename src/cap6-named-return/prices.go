package main

import "fmt"

func FinalPrice(costPrice float64) (priceDollar float64, priceReal float64) {

	profitFactor := 1.33
	taxConversion := 2.34

	priceDollar = costPrice * profitFactor
	priceReal = priceDollar * taxConversion

	return priceDollar, priceReal

}

func main() {
	priceDollar, priceReal := FinalPrice(34.99)

	fmt.Printf("Final price in dollar: %.2f\n" + "Final price in reals: %.2f", priceDollar, priceReal)
}
