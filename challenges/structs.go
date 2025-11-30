package challenges

import "errors"

// Item represents a product in the inventory.
type Item struct {
	ID    string
	Name  string
	Price float64
	Stock int
}

// Inventory manages a collection of items.
type Inventory struct {
	Items map[string]*Item
}

// NewInventory creates a new empty Inventory.
func NewInventory() *Inventory {
	// TODO: Initialize and return a new Inventory
	
	return &Inventory {
		Items: make(map[string]*Item),
	}
}

// AddItem adds a new item to the inventory.
// If an item with the same ID already exists, return an error.
func (inv *Inventory) AddItem(id, name string, price float64, stock int) error {
	// TODO: Implement this method
	if inv.Items[id] != nil {
		return errors.New("Item with this ID already exists")
	}
	inv.Items[id] = &Item{id, name, price, stock}
	return nil
}

// UpdateStock updates the stock of an existing item.
// If the item does not exist, return an error.
// If the new stock is negative, return an error.
func (inv *Inventory) UpdateStock(id string, newStock int) error {
	// TODO: Implement this method
	item := inv.Items[id]
	if item == nil {
		return errors.New("ID does not exist")
	}
	if newStock < 0 {
		return errors.New("Negative Stock")
	}
	item.Stock = newStock
	return nil
}

// GetTotalValue calculates the total value of the inventory (sum of Price * Stock for all items).
func (inv *Inventory) GetTotalValue() float64 {
	// TODO: Implement this method
	var total float64
	for _, i := range inv.Items {
		total += (i.Price * float64 (i.Stock)) 
	}
	return total
}
