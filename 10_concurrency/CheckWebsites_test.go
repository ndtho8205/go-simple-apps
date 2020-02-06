package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"http://google.com",
		"http://ndtho8205.github.io",
		"waat://furhurterwe.geds",
	}

	got := CheckWebsites(mockWebsiteChecker, urls)
	want := map[string]bool{
		"http://google.com":          true,
		"http://ndtho8205.github.io": true,
		"waat://furhurterwe.geds":    false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}
