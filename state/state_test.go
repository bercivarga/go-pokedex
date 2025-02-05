package state

import "testing"

func TestStateReset(t *testing.T) {
	state := NewAppState()
	state.LocationAreas = []string{"one", "two", "three"}
	state.AreaPage = 3
	state.Reset()
	if len(state.LocationAreas) != 0 {
		t.Errorf("Expected LocationAreas to be empty, got %v", state.LocationAreas)
	} else {
		t.Logf("Expected LocationAreas to be empty, got %v", state.LocationAreas)
	}
	if state.AreaPage != 0 {
		t.Errorf("Expected AreaPage to be 0, got %d", state.AreaPage)
	} else {
		t.Logf("Expected AreaPage to be 0, got %d", state.AreaPage)
	}
}
