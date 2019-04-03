package main

import (
	"math"
	"fmt"
	"errors"
)

// func avg(x, y float64) float64{
// 	return (x + y) / 2git 
// }

// func main() {
// 	x := 5.75
// 	y := 6.25

// 	result := avg(x,y)

// 	fmt.Printf("Average of %.1f and of %.1f = %.1f\n", x, y, result)
// }

func getStockPriceChangeWithError(prevPrice, currentPrice float64) (float64, float64, error) {
	if prevPrice == 0 {
		err := errors.New("Previous price cant be zero")
		return 0, 0, err
	}
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange, nil
}

func main() {
	prevStockPrice := 45000.0
	currentStockPrice := 100000.0

	change, percentChange, err := getStockPriceChangeWithError(prevStockPrice, currentStockPrice)

	if err != nil {
		fmt.Println("Sorry there was an error: ", err)
	}else {
	if change < 0 {
		fmt.Printf("The Stock price decreased by $%.2f which is %.2f%% of the prev price\n",
					math.Abs(change), math.Abs(percentChange))
	}else {
		fmt.Printf("The Stock price increased by $%.2f which is %.2f%% of the prev price\n",
					change, percentChange)
	}
}
} 
