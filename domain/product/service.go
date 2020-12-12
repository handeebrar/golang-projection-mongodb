package product

import (
	"log"
	"projection/events/product"
)

type ProductService interface {
	Create(event *product.Created) error
}

type productService struct {
	logger log.Logger
	//repo
}

func (self *productService) Create(event *product.Created) error{
	return nil
}