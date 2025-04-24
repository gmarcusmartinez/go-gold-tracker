package main

import "testing"

func TestApp_getPriceText(t *testing.T) {
	open, current, change := testApp.getPriceText()

	if open.Text != "Open: $3285.3575 USD" {
		t.Error("wrong open price returned: ", open)
	}

	if current.Text != "Current: $3331.8950 USD" {
		t.Error("wrong current price returned: ", current)
	}

	if change.Text != "Change: $46.5375 USD" {
		t.Error("wrong change price returned: ", change)
	}
}
