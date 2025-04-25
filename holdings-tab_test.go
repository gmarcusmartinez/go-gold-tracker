package main

import "testing"

func TestConfig_currentHoldings(t *testing.T) {
	all, err := testApp.currentHoldings()
	if err != nil {
		t.Error("failed to get current holdings from database")
	}

	if len(all) != 2 {
		t.Error("wrong number of holdins returned")
	}
}
