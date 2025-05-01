package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestApp_getToolbar(t *testing.T) {
	tb := testApp.getToolbar()

	if len(tb.Items) != 4 {
		t.Error("wrong number of items in toolbar: ", len(tb.Items))
	}
}

func TestApp_addHoldingsDialog(t *testing.T) {
	testApp.addHoldingsDialog()

	test.Type(testApp.PurchaseAmountEntry, "1")
	test.Type(testApp.PurchasePriceEntry, "1000")
	test.Type(testApp.PurchaseDateEntry, "2022-01-01")

	if testApp.PurchaseAmountEntry.Text != "1" {
		t.Error("amount not correct")
	}

	if testApp.PurchasePriceEntry.Text != "1000" {
		t.Error("price not correct")
	}

	if testApp.PurchaseDateEntry.Text != "2022-01-01" {
		t.Error("date not correct")
	}
}
