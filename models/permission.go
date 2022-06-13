package models

type Permission struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

//permission in db postgre assign manual ex col id name: 1, view_products
