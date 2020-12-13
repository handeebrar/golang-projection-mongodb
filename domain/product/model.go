package product

import (
	"gopkg.in/mgo.v2/bson"
	mongoClient "projection/mongo"
)

type Product struct {
	Id          bson.ObjectId   `bson:"_id" json:"id" `
	Name		string			`json:"name"`
	UnitPrice 	string      	`json:"unitPrice"`
	Quantity 	int 			`json:"quantity"`
}

func (product Product) InsertProduct() error {
	product.Id = bson.NewObjectId()
	if err := mongoClient.Db.C("Product").Insert(&product); err != nil {
		return err
	} else {
		return nil
	}
}