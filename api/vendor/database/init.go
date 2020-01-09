package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var DB *gorm.DB
var err error

type Post struct {
	gorm.Model
	Identity    int64
	CreditScore int64
}

func addDatabase(dbname string) error {
	// db ismini aarak eğer db tanımlı değil ise tanımla
	DB.Exec("CREATE DATABASE " + dbname)

	// tanımlanmış db ye bağlan ve parametreleri tanımla
	connectionParams := "dbname=" + dbname + " user=docker password=docker sslmode=disable host=db"
	DB, err = gorm.Open("postgres", connectionParams)
	if err != nil {
		return err
	}

	return nil
}

func Init() (*gorm.DB, error) {
	//  database bağlantısını kur
	connectionParams := "user=docker password=docker sslmode=disable host=db"
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open("postgres", connectionParams) // gorm ayakta olup olmadığını anlamak için ping atıyor
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return DB, err
	}

	// datatable tanımlı değil ise yaratıyor
	if !DB.HasTable(&Post{}) {
		DB.CreateTable(&Post{})
	}

	testPost1 := Post{Identity: 12345678910, CreditScore: 1280}
	testPost2 := Post{Identity: 14345678910, CreditScore: 100}
	testPost3 := Post{Identity: 13345678910, CreditScore: 430}

	DB.Create(&testPost1)
	DB.Create(&testPost2)
	DB.Create(&testPost3)

	return DB, err
}
