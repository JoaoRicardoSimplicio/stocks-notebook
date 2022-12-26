package main

import (
	"fmt"
	"stocks-notebook/calculator"
	"stocks-notebook/notebook"
)

func main() {
	stocksOrdersGroupedByStockCode := notebook.GetStockOrdersGroupedByStockCode()
	for code, stockOrders := range stocksOrdersGroupedByStockCode {
		result := calculator.CalculateStocksStats(code, stockOrders)
		fmt.Println(result)
	}
}
