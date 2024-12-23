package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

type MyNote struct {
	gorm.Model
	Name string `gorm:"name"`
}

func main() {
	//TIP Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined or highlighted text
	// to see how GoLand suggests fixing it.
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	dsn := "client:pass123@tcp(127.0.0.1:3309)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	log.SetPrefix("Here is main PKG :: ")
	//log.SetFlags(40)

	if err != nil {
		//fmt.Printf("Connection to database is failed :: %s!\n", err)
		log.Fatalln("Connection to database is failed", err)
	}

	db = db.Debug()

	log.Println("Connection to database is success!", db)

	//n := MyNote{
	//	Name: "haha",
	//}
	//
	//errCreate := db.Create(&n).Error
	//
	//if errCreate != nil {
	//	log.Fatalln(errCreate)
	//}
	//
	//log.Fatalln(n)

	//var myNote MyNote
	//myNote.Name = "haha1"
	//
	//if err := db.Where(myNote).
	//	First(&myNote).
	//	Error; err != nil {
	//	log.Println("Note NOT FOUND!")
	//}
	//
	//log.Println(myNote)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:7000") // listen and serve on 0.0.0.0:8080
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
