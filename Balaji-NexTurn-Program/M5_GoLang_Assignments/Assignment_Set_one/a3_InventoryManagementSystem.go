// Case Study:
// A store needs to manage its inventory of products. Build an application that includes
// the following:
// 1. Product Struct: Create a struct to represent a product with fields for ID, name,
// price (float64), and stock (int).
// 2. Add Product: Write a function to add new products to the inventory. Use type
// casting to ensure price inputs are converted to float64.
// 3. Update Stock: Implement a function to update the stock of a product. Use
// conditions to validate the input (e.g., stock cannot be negative).
// 4. Search Product: Allow users to search for products by name or ID. If a product is
// not found, return a custom error message.
// 5. Display Inventory: Use loops to display all available products in a formatted
// table.

package main

import (
	"errors"
	"fmt"
	"strings"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory []Product

func addProduct(id int, name string, price interface{}, stock int) error {
	// Validate price type and convert to float64
	priceFloat, ok := price.(float64)
	if !ok {
		return errors.New("price must be a float64")
	}

	// Check for unique ID
	for _, product := range inventory {
		if product.ID == id {
			return errors.New("product ID must be unique")
		}
	}

	inventory = append(inventory, Product{
		ID:    id,
		Name:  name,
		Price: priceFloat,
		Stock: stock,
	})
	return nil
}

func updateStock(id int, stock int) error {
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}
	for i := range inventory {
		if inventory[i].ID == id {
			inventory[i].Stock = stock
			return nil
		}
	}
	return errors.New("product not found")
}

func searchProduct(query string) (Product, error) {
	for _, product := range inventory {
		if fmt.Sprintf("%d", product.ID) == query || product.Name == query {
			return product, nil
		}
	}
	return Product{}, errors.New("product not found")
}

func displayInventory() {
	fmt.Printf("%-10s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 50))
	for _, product := range inventory {
		fmt.Printf("%-10d %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func main() {
	// Add some products
	if err := addProduct(1, "Laptop", 999.99, 10); err != nil {
		fmt.Println("Error:", err)
	}
	if err := addProduct(2, "Mouse", 19.99, 50); err != nil {
		fmt.Println("Error:", err)
	}
	if err := addProduct(3, "Keyboard", 49.99, 30); err != nil {
		fmt.Println("Error:", err)
	}

	// Update stock
	if err := updateStock(1, 8); err != nil {
		fmt.Println("Error:", err)
	}

	// Search for a product
	product, err := searchProduct("2")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found Product: %+v\n", product)
	}

	// Display inventory
	fmt.Println("\nInventory:")
	displayInventory()
}
