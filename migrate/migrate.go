package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Ghezak-git/synapsis/initializers"
	"github.com/Ghezak-git/synapsis/models"
	"github.com/Ghezak-git/synapsis/utils"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.ShopCart{},
		&models.TransactionHeader{},
		&models.TransactioDetail{},
		&models.PaymentMethod{},
	)
	now := time.Now()
	hashedPassword, _ := utils.HashPassword("admin")
	initializers.DB.Create(
		&models.User{
			Email:     "admin@gmail.com",
			Password:  hashedPassword,
			FullName:  "Admin",
			Role:      "admin",
			CreatedAt: now,
			UpdatedAt: now,
		},
	)
	initializers.DB.Create(
		[]models.Category{
			{
				Name:        "Clothes",
				Description: "Example Description Clothes",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				Name:        "Technology",
				Description: "Example Description Technology",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		},
	)
	initializers.DB.Create(
		[]models.Product{
			{
				IDCategory:  1,
				Name:        "Clothes ASD",
				Description: "Example Description Clothes ASD",
				Price:       56000,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				IDCategory:  1,
				Name:        "Clothes QWE",
				Description: "Example Description Clothes QWE",
				Price:       56000,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				IDCategory:  2,
				Name:        "Laptop ASD",
				Description: "Example Description Laptop ASD",
				Price:       56000,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				IDCategory:  2,
				Name:        "Laptop QWE",
				Description: "Example Description Laptop QWE",
				Price:       56000,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		},
	)
	initializers.DB.Create(
		[]models.PaymentMethod{
			{
				Nama:        "BRI",
				Description: "18236791263 a/n BRI",
			},
			{
				Nama:        "BCA",
				Description: "17236712232 a/n BCA",
			},
			{
				Nama:        "BNI",
				Description: "98327498732 a/n BNI",
			},
		},
	)
	fmt.Println("👍 Migration complete")
}
