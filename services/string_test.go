package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
		{"CC23FRANASPERMASTER", false, "", "", 0},
		{"TX3F", false, "", "", 0},
	}
	for _, testData := range cases {
		var c Chain
		r, err := ValidChain(testData.Input)
		assert.Equal(t, err == nil, testData.Success)
		if !r {
			y, _ := c.ModChain(testData.Input, r)

			assert.Equal(t, testData.Type, y.Type)
			assert.Equal(t, testData.Value, y.Value)
			assert.Equal(t, testData.Length, y.Length)

		}
	}
}
