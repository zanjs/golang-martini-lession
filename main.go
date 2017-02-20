package main

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Product is ...
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create

	db.Create(&Product{Code: "l1", Price: 1000})

	// Read

	var product Product
	pro := db.First(&product, 1)         // find product with id 1
	db.First(&product, "code = ?", "l1") //find product with code l1

	db.Model(&product).Update("Proce", 2000)

	log.Println(pro)

	// Delect - delect product
	db.Delete(&product)

	m := martini.Classic()

	// log 记录请求完成前后  (*译者注: 很巧妙，掌声鼓励.)
	m.Use(func(c martini.Context, log *log.Logger) {
		log.Println("before a request")

		c.Next()

		log.Println("after a request")
	})

	// 验证api密匙
	m.Use(func(res http.ResponseWriter, req *http.Request) {
		if req.Header.Get("X-API-KEY") != "secret123" {

			log.Println("auth log...")

			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte("auth err"))
		}
	})

	m.Get("/", func() string {
		return "hello dage"
	})
	m.Run()
}
