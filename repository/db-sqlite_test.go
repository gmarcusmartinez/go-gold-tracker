package repository

import (
	"testing"
	"time"
)

func Test_DBSQLiteRepository(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Errorf("Failed to migrate: %v", err)
	}
}

func Test_InsertHolding(t *testing.T) {
	h := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}

	result, err := testRepo.InsertHolding(h)
	if err != nil {
		t.Errorf("Failed to insert holding: %v", err)
	}

	if result.ID <= 0 {
		t.Errorf("Failed to insert holding: %v", err)
	}
}

func Test_GetAllHoldings(t *testing.T) {
	holdings, err := testRepo.AllHoldings()
	if err != nil {
		t.Errorf("Failed to get all holdings: %v", err)
	}

	if len(holdings) != 1 {
		t.Errorf("Expected 1 holding, got %d", len(holdings))
	}
}

func Test_GetHoldingByID(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Errorf("Failed to get holding by id: %v", err)
	}

	if h.PurchasePrice != 1000 {
		t.Errorf("Expected 1000 purchase price, got %d", h.PurchasePrice)
	}

	h, err = testRepo.GetHoldingByID(2)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

}

func Test_UpdateHolding(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Errorf("Failed to get holding by id: %v", err)
	}

	h.PurchasePrice = 2000
	err = testRepo.UpdateHolding(h.ID, *h)
	if err != nil {
		t.Errorf("Failed to update holding: %v", err)
	}

}

func Test_DeleteHolding(t *testing.T) {
	err := testRepo.DeleteHolding(1)
	if err != nil {
		t.Errorf("Failed to delete holding: %v", err)
		if err != errorDeleteFailed {
			t.Errorf("Expected errorDeleteFailed, got %v", err)
		}
	}
	err = testRepo.DeleteHolding(2)
	if err == nil {
		t.Errorf("No error when attempting to delete non-existent holding")
	}
}
