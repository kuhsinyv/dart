package dart_test

import (
	_ "embed"
	"github.com/kuhsinyv/dart"
	"strings"
	"testing"
)

//go:embed test_Chinese.txt
var chinesePatterns string

//go:embed test_English.txt
var englishPatterns string

//go:embed test_mix.txt
var mixPatterns string

var patternsSlice = []string{
	chinesePatterns,
	englishPatterns,
	mixPatterns,
}

func TestBuild(t *testing.T) {
	for _, patterns := range patternsSlice {
		d := new(dart.Dart)

		_, _, err := d.Build(strings.Split(patterns, "\n"))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestExactMatchSearch(t *testing.T) {
	for _, patterns := range patternsSlice {
		d := new(dart.Dart)

		dat, _, err := d.Build(strings.Split(patterns, "\n"))
		if err != nil {
			t.Fatal(err)
		}

		for _, pattern := range strings.Split(patterns, "\n") {
			if !dat.ExactMatchSearch([]rune(pattern), 0) {
				t.Fatalf("not match error: %v\n", pattern)
			}
		}
	}
}
