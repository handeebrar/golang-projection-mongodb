package product

import "git.hepsiburada.com/oms/projections/events/endor/common"

type Product struct {
	Name		string			`json:"name"`
	UnitPrice 	common.Money	`json:"unitPrice"`
	Quantity 	int 			`json:"quantity"`
}