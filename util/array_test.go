package util

import (
	"reflect"
	"testing"
)

func TestCleanStringArray(t *testing.T) {

	input := []string{"", "hej1", "hej2\r\n", "hej3\n", "hej4\r", "\r\n", " "}
	expected := []string{"hej1", "hej2", "hej3", "hej4"}

	result := CleanStringArray(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CleanStringArray expected length: %d. Got %d", len(expected), len(result))
	}
}

func TestStringArrayToIntArray(t *testing.T) {

	input := []string{"1", "2", "3"}
	expected := []int{1, 2, 3}

	result := StringArrayToIntArray(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CleanStringArray expected length: %d. Got %d", len(expected), len(result))
	}
}
