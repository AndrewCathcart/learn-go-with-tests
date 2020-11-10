package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://andrewcathcart.github.io" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://youtube.com",
		"http://andrewcathcart.github.io",
	}

	expected := map[string]bool{
		"http://google.com":               true,
		"http://youtube.com":              true,
		"http://andrewcathcart.github.io": false,
	}

	actual := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}
