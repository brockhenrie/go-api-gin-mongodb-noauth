package models

type Address struct{
	Address1 		string `json:"address1" bson:"address1"`
	Address2 		string	`json:"address2" bson:"address2"`
	City 			string`json:"city" bson:"city"`
	State			string`json:"state" bson:"state"`
	Zip 			int`json:"zip" bson:"zip"`
}

type User struct {
	Name		string `json:"name" bson:"user_name"`
	Age 		int `json:"age" bson:"user_age"`
	Address		Address `json:"address" bson:"user_address"`
}

