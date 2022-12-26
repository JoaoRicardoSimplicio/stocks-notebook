package calculator

import (
	"stocks-notebook/notebook"
)

type StockStats struct {
	Code           string
	AvaragePrice   float64
	TotalCostPrice float64
	TotalAmount    int
}

func calculateStockOrderBuy(stockOrder *notebook.StockOrder, stockStats *StockStats) {
	stockStats.TotalAmount += stockOrder.Amount
	stockStats.TotalCostPrice += stockOrder.Price * float64(stockOrder.Amount)
}

func calculateStockOrderSell(stockOrder *notebook.StockOrder, stockStats *StockStats) {
	stockStats.TotalAmount -= stockOrder.Amount
	stockStats.TotalCostPrice -= stockOrder.Price * float64(stockOrder.Amount)
}

func CalculateStocksStats(stockCode string, stockOrders []*notebook.StockOrder) StockStats {
	var stockStats StockStats = StockStats{Code: stockCode}

	for _, stockOrder := range stockOrders {
		switch stockOrder.Action {
		case "BUY":
			calculateStockOrderBuy(stockOrder, &stockStats)
		case "SELL":
			calculateStockOrderSell(stockOrder, &stockStats)
		}
	}
	stockStats.AvaragePrice = stockStats.TotalCostPrice / float64(stockStats.TotalAmount)

	return stockStats
}
