package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	var status string
	if p.Stock < 3 {
		status = "Stock is almost running out"
	} else if p.Stock < 10 {
		status = "Stock is limited"
	}
	return status
}
