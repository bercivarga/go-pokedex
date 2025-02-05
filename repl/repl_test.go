package repl

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	input := "BulBasaur Charmander PIKACHU"
	expected := []string{"bulbasaur", "charmander", "pikachu"}
	actual := CleanInput(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	} else {
		t.Logf("Expected %v, got %v", expected, actual)
	}
}
