package challenges

import (
	"testing"
)

func TestInventory(t *testing.T) {
	inv := NewInventory()
	if inv == nil {
		t.Fatal("NewInventory() returned nil")
	}

	// Test AddItem
	err := inv.AddItem("1", "Laptop", 1000.0, 5)
	if err != nil {
		t.Errorf("AddItem() unexpected error: %v", err)
	}

	// Test AddItem duplicate
	err = inv.AddItem("1", "Laptop", 1000.0, 5)
	if err == nil {
		t.Error("AddItem() expected error for duplicate ID, got nil")
	}

	// Test UpdateStock
	err = inv.UpdateStock("1", 10)
	if err != nil {
		t.Errorf("UpdateStock() unexpected error: %v", err)
	}

	// Test UpdateStock non-existent
	err = inv.UpdateStock("999", 10)
	if err == nil {
		t.Error("UpdateStock() expected error for non-existent ID, got nil")
	}

	// Test UpdateStock negative
	err = inv.UpdateStock("1", -5)
	if err == nil {
		t.Error("UpdateStock() expected error for negative stock, got nil")
	}

	// Test GetTotalValue
	// Current: 10 Laptops @ 1000.0 = 10000.0
	inv.AddItem("2", "Mouse", 20.0, 50)
	// Added: 50 Mice @ 20.0 = 1000.0
	// Total: 11000.0

	wantVal := 11000.0
	gotVal := inv.GetTotalValue()
	if gotVal != wantVal {
		t.Errorf("GetTotalValue() = %v, want %v", gotVal, wantVal)
	}
}
