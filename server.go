package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:michomicho!@tcp(127.0.0.1:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE test_db")
	db.Exec("USE test_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&models.Order{}, &models.Item{})
}

func mapItemsFromInput(itemsInput []*ItemInput) []models.Item {
	var items []models.Item
	for _, itemInput := range itemsInput {
		items = append(items, models.Item{
			ProductCode: itemInput.ProductCode,
			ProductName: itemInput.ProductName,
			Quantity:    itemInput.Quantity,
		})
	}
	return items
}
func (r *mutationResolver) CreateOrder(ctx context.Context, input OrderInput) (*models.Order, error) {
	order := models.Order{
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        mapItemsFromInput(input.Items),
	}
	err := r.DB.Create(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, orderID int, input OrderInput) (*models.Order, error) {
	updatedOrder := models.Order{
		ID:           orderID,
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        mapItemsFromInput(input.Items),
	}
	r.DB.Save(&updatedOrder)
	return &updatedOrder, nil
}
func (r *mutationResolver) DeleteOrder(ctx context.Context, orderID int) (bool, error) {
	r.DB.Where("order_id = ?", orderID).Delete(&models.Item{})
	r.DB.Where("order_id = ?", orderID).Delete(&models.Order{})
	return true, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(backend_personal_webapp.NewExecutableSchema(backend-personal-webapp.Config{Resolvers: &backend - personal - webapp.Resolver{
		DB: db,
	}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
