package notebook

const FILEPATH string = "./notebook.txt"

func GetStockOrders() []*StockOrder {
	linesOfFile := ReadLinesFromFile(FILEPATH)
	stockOrders := transformRecordsInStockOrders(&linesOfFile)
	return stockOrders
}

func GetStockOrdersGroupedByStockCode() map[string][]*StockOrder {
	stockOrders := GetStockOrders()
	stockOrdersGroupedByDate := groupStockOrdersByStockCode(stockOrders)
	return stockOrdersGroupedByDate
}
