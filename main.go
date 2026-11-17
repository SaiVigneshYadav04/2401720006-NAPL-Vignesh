package main

import (
	"errors"
	"fmt"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

var items []Item

func AddItem(item Item) error {
	if item.ID <= 0 || item.Name == "" || item.Quantity < 0 || item.Price < 0 {
		return errors.New("invalid item")
	}

	for _, v := range items {
		if v.ID == item.ID {
			return errors.New("item already exists")
		}
	}

	items = append(items, item)
	return nil
}

func UpdateItem(id, quantity int, price float64) error {
	if quantity < 0 || price < 0 {
		return errors.New("invalid details")
	}

	for i := range items {
		if items[i].ID == id {
			items[i].Quantity = quantity
			items[i].Price = price
			return nil
		}
	}

	return errors.New("item not found")
}

func RemoveItem(id int) error {
	for i, v := range items {
		if v.ID == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}

	return errors.New("item not found")
}

type Stats struct {
	TotalValue float64
	ItemCount  int
}

func CalculateStats() Stats {
	ch := make(chan Stats)

	go func() {
		var s Stats

		for _, v := range items {
			s.TotalValue += float64(v.Quantity) * v.Price
			s.ItemCount++
		}

		ch <- s
	}()

	return <-ch
}

func main() {
	items = []Item{}

	fmt.Println("INVENTORY MANAGEMENT SYSTEM")

	fmt.Println("\nAdding Items:")
	if err := AddItem(Item{1, "Laptop", 5, 50000}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Laptop added successfully")
	}

	if err := AddItem(Item{2, "Mouse", 10, 500}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Mouse added successfully")
	}

	fmt.Println("\nUpdating Item:")
	if err := UpdateItem(1, 8, 55000); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Laptop updated successfully")
		fmt.Println("New Quantity: 8")
		fmt.Println("New Price: 55000")
	}

	fmt.Println("\nRemoving Item:")
	if err := RemoveItem(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Mouse removed successfully")
	}

	fmt.Println("\nFinal Inventory:")
	for _, v := range items {
		fmt.Printf("ID: %d, Name: %s, Quantity: %d, Price: %.2f\n",
			v.ID, v.Name, v.Quantity, v.Price)
	}

	s := CalculateStats()

	fmt.Println("\nInventory Summary:")
	fmt.Println("Total Inventory Value:", s.TotalValue)
	fmt.Println("Total Items:", s.ItemCount)
}
