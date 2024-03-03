package baseproject

import (
	"fmt"
	"github.com/uberswe/golang-base-project/config"
	"github.com/uberswe/golang-base-project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func connectToDatabase(c config.Config) (db *gorm.DB, err error) {
	return connectLoop(c, 0)
}

func connectLoop(c config.Config, count int) (db *gorm.DB, err error) {
	db, err = attemptConnection(c)
	if err != nil {
		if count > 300 {
			return db, fmt.Errorf("could not connect to database after 300 seconds")
		}
		time.Sleep(1 * time.Second)
		return connectLoop(c, count+1)
	}
	return db, err
}

func attemptConnection(c config.Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=localhost user=postgres password=Just_arys7 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

func migrateDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Token{}, &models.Session{}, &models.Website{})
	seed(db)
	return err
}

func seed(db *gorm.DB) {
	// We seed some websites for our search results
	websites := []models.Website{
		{
			Title:       "A Tour of Go",
			Description: "A Tour of Go has several interactive examples of how Go which you can learn from. There is a menu available if you would like to skip to different sections.",
			URL:         "https://go.dev/tour/welcome/1",
		},
		{
			Title:       "Go by Example",
			Description: "As described on the website: Go by Example is a hands-on introduction to Go using annotated example programs. I have used this site many times as a reference when I need to look something up.",
			URL:         "https://gobyexample.com/",
		},
		{
			Title:       "Go.dev",
			Description: "Learn how to install Go on your machine and read the documentation on the Go website.",
			URL:         "https://go.dev/learn/",
		},
		{
			Title:       "Uberswe on Github",
			Description: "I am the creator of Golang Base Project. This is my Github profile.",
			URL:         "https://github.com/uberswe",
		},
		{
			Title:       "Tournify",
			Description: "A website to create tournaments or free which uses this project as a base.",
			URL:         "https://tournify.io",
		},
		{
			Title:       "GORM",
			Description: "The fantastic ORM library for Golang.",
			URL:         "https://gorm.io/",
		},
		{
			Title:       "Bootstrap",
			Description: "Quickly design and customize responsive mobile-first sites with Bootstrap, the world’s most popular front-end open source toolkit, featuring Sass variables and mixins, responsive grid system, extensive prebuilt components, and powerful JavaScript plugins.",
			URL:         "https://getbootstrap.com/",
		},
		{
			Title:       "Gin Web Framework",
			Description: "Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.",
			URL:         "https://github.com/gin-gonic/gin",
		},
	}

	for _, w := range websites {
		res := db.Where(&w).First(&w)
		// If no record exists we insert
		if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
			db.Save(&w)
		}
	}
}
