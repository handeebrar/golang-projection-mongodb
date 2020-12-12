package product

import (
	"log"
	"projection/events/product"
)

type ProductMapper interface {
	MapForCreate(event *product.Created, product *Product)
}

type productMapper struct {
	logger log.Logger
}

func (self *productMapper) MapForCreate(event *product.Created, product *Product) {
	panic("implement me")
}

func NewProductmapper(logger log.Logger) ProductMapper{
	return &productMapper{logger: logger}
}