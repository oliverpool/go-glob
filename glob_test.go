package glob_test

import (
	"testing"

	glob "github.com/oliverpool/go-glob"
)

func TestGlob(t *testing.T) {
	tt := []struct {
		pattern  string
		matching []string
		failing  []string
	}{
		{
			pattern:  "",
			matching: []string{""},
			failing:  []string{"*", "word", " "},
		},
		{
			pattern:  "a",
			matching: []string{"a"},
			failing:  []string{"", "b", "aa"},
		},
		{
			pattern:  "*",
			matching: []string{"", "*", "word"},
		},
		{
			pattern:  "*****",
			matching: []string{"", "*", "word"},
		},
		{
			pattern:  "*end",
			matching: []string{"end", ".end", "endend"},
			failing:  []string{"end.", "en", "nd", ""},
		},
		{
			pattern:  "start*",
			matching: []string{"start", "start.", "startstart"},
			failing:  []string{".start", "star", "tart", ""},
		},
		{
			pattern:  "start*end",
			matching: []string{"startend", "start.end", "startstartendend"},
			failing:  []string{"starten", "tartend", "start", "", ".startend", "startend."},
		},
		{
			pattern:  "1*2*3",
			matching: []string{"123", "1.2.3", "11111123123"},
			failing:  []string{"12", "123.", ".123", "13", ""},
		},
		{
			pattern:  "1***3",
			matching: []string{"123", "1.2.3", "11111123123", "13"},
			failing:  []string{"12", "123.", ".123", ""},
		},
		{
			pattern:  "*.*",
			matching: []string{".", "a.b", "a.", "b."},
			failing:  []string{"", "a", "b", "*"},
		},
	}

	for _, tc := range tt {
		matcher := glob.Matcher(tc.pattern)
		t.Run(tc.pattern, func(t *testing.T) {
			for _, m := range tc.matching {
				if !matcher(m) {
					t.Errorf("'%s' should match '%s'", m, tc.pattern)
				}
			}
			for _, m := range tc.failing {
				if matcher(m) {
					t.Errorf("'%s' should not match '%s'", m, tc.pattern)
				}
			}
		})
	}
}
