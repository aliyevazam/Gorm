// package main

// import (
// 	"fmt"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type CreditCard struct {
// 	gorm.Model
// 	Number string
// 	UserID uint
// }

// type User struct {
// 	gorm.Model
// 	Name       string
// 	CreditCard CreditCard
// }

// func main() {
// 	dsn := "host=localhost user=azam password=Azam_2000 dbname=test1 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("Error 1")
// 	}

// 	db.AutoMigrate(&User{})
// 	db.AutoMigrate(&CreditCard{})
// 	db.Create(&User{
// 		Name:       "jinzhu",
// 		CreditCard: CreditCard{Number: "411111111111"},
// 	})
// 	// user := User{Name: "Golang", Age: 10, Birthday: time.Now()}
// 	// result := db.Create(&user)

// }

package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name   string
	Models string
	Stores []Store `gorm:"many2many:product_storages;"`
}
type Store struct {
	gorm.Model
	Name     string
	Adresses []Address `gorm:"many2many:storage_addresses;"`
}

type Address struct {
	gorm.Model
	Street   string
	District string
}

type Info struct {
	Name         string
	Categoryname string
	Typename     string
	Models       string
	Price        float64
	Amount       int
	Storename    string
	District     string
	Street       string
}

func main() {
	dsn := "host=localhost user=azam password=Azam_2000 dbname=test1 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error while connecting db", err)
	}
	tx := db.Begin()
	tx.AutoMigrate(&Address{}, &Store{}, &Product{}, &Info{})
	if err := tx.Create(&Product{
		Name:   "Iphone 14",
		Models: "Pro max",
		Stores: []Store{
			{
				Name: "Malika TC",
				Adressess: []Address{
					{
						District: "Yunusobod",
						Street:   "7 kvartal",
					},
				},
			},
		},
	}).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
	}
	fmt.Println(tx.Commit())
}
