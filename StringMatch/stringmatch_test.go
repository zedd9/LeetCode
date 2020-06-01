package StringMatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert := assert.New(t)

	words := []string{"mass", "as", "hero", "superhero"}

	result := StringMatching(words)
	expect := []string{"as", "hero"}
	for _, str := range result {
		assert.Contains(expect, str)
	}
	assert.Equal(len(result), 2)
}

func TestCase2(t *testing.T) {
	assert := assert.New(t)

	words := []string{"leetcode", "et", "code"}

	result := StringMatching(words)
	expect := []string{"et", "code"}
	for _, str := range result {
		assert.Contains(expect, str)
	}
	assert.Equal(len(result), 2)
}
