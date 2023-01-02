package main

import (
	"fmt"
	"stocks-notebook/calculator"
	"stocks-notebook/notebook"
)

func main() {
	stocksOrdersGroupedByStockCode := notebook.GetStockOrdersGroupedByStockCode()
	for code, stockOrders := range stocksOrdersGroupedByStockCode {
		stockStats := calculator.CalculateStocksStats(code, stockOrders)
		fmt.Printf(
			"Stock Code: %s\nAvarage Price: %.3f\nTotal Cost Price: %.3f\nTotal Amount: %d\n\n",
			stockStats.Code,
			stockStats.AvaragePrice,
			stockStats.TotalCostPrice,
			stockStats.TotalAmount,
		)
	}
}
