package main

import (
	"reflect"
	"testing"
)

func TestURL(t *testing.T) {
	test_suits := []struct {
		name string
		input string
		expected []string
	}{
		{		
			name: "",
			input: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
	}

	for i, tc := range test_suits {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.input, "")
			if err != nil {
				t.Errorf("Error occured while parsing the html")
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}