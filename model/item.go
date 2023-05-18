package model

type InputBook struct {
	Title    string
	Price    int
	SubTitle string `json:"sub_title"` //ini dipake buat nangkep key json dari postman
}

type ItemStruct struct {
	Name     string `json:"name", binding:"required"`
	Price    int    `json:"price", binding:"required,number"`
	Quantity int    `json:"qty", binding:"required,number"` //ini dipake buat nangkep key json dari postman
}

// type Tabler interface {
// 	TableName() string
// }

// func (Mst_user) TableName() string {
// 	return "mst_user"
// }
