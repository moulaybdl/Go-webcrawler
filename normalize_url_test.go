package main

import (
	"testing"
)

func TestNomrlizeURL(t *testing.T) {
	tests := []struct {
		name string
		inputURL string
		expected string
	} {
			{	
				name: "base test",
				inputURL: "https://blog.boot.dev/path",
				expected: "blog.boot.dev/path",
			},
			{	
				name: "no schema slash test",
				inputURL: "https://blog.boot.dev/path/",
				expected: "blog.boot.dev/path",
			},
			{	
				name: "schema test",
				inputURL: "http://blog.boot.dev/path",
				expected: "blog.boot.dev/path",
			},
			{	
				name: "schema slash base test",
				inputURL: "http://blog.boot.dev/path/",
				expected: "blog.boot.dev/path",
			},
			
	}

		for i, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				actual, err := normalizeURL(tc.inputURL)
				if err != nil {
					t.Errorf("Test %v - '%s' FAIL: unexpected error: %v ", i, tc.name, err)
					return
				}

				if actual != tc.expected {
					t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
				}
			})
		}

}