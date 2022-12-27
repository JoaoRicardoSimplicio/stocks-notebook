package notebook

import (
	"strconv"
	"strings"
	"time"
)

type StockOrder struct {
	Date   time.Time
	Code   string
	Action string
	Amount int
	Price  float64
}

const DATEFORMAT string = "2006-01-02"

func transformRecordsInStockOrders(records *[]string) []*StockOrder {
	var stockOrders []*StockOrder

	for _, record := range *records {
		recordItems := strings.Split(record, ";")
		date, _ := time.Parse(DATEFORMAT, recordItems[0])
		amount, _ := strconv.Atoi(recordItems[3])
		price, _ := strconv.ParseFloat(recordItems[4], 64)
		stockOrder := StockOrder{
			Date:   date,
			Code:   recordItems[1],
			Action: recordItems[2],
			Amount: amount,
			Price:  price,
		}
		stockOrders = append(stockOrders, &stockOrder)
	}

	return stockOrders
}

func groupStockOrdersByStockCode(stockOrders []*StockOrder) map[string][]*StockOrder {
	groupedStockOrders := make(map[string][]*StockOrder)
	for _, stockOrder := range stockOrders {
		if _, ok := groupedStockOrders[stockOrder.Code]; !ok {
			groupedStockOrders[stockOrder.Code] = make([]*StockOrder, 0)
		}
		groupedStockOrders[stockOrder.Code] = append(
			groupedStockOrders[stockOrder.Code],
			stockOrder,
		)
	}
	return groupedStockOrders
}
