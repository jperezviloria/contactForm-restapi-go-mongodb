package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db: %s\n")
	}
	db.AutoMigrate(&Person{})

	app.Post("/create", func(ctx *fiber.Ctx) error {
		db.Create(&Person{Code: "D43", Price: 123})
		return err
	})

	app.Get("/get", func(ctx *fiber.Ctx) error {

		var person Person
		db.First(&person, 1)
		db.First(&person, "code = ?", "D43")
		return err
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello")
	})

	app.Listen(":3000")

}
