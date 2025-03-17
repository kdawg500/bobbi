package link

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate_Success(t *testing.T) {
	testCases := []struct {
		name     string
		link     string
		expected string
	}{
		{
			name:     "with https",
			link:     "https://facebook.com",
			expected: "https://facebook.com",
		},
		{
			name:     "without https",
			link:     "facebook.com",
			expected: "https://facebook.com",
		},
	}

	for _, tc := range testCases {
		actual := Validate(tc.link)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestValidate_InvalidLink(t *testing.T) {
	linkString := "deez"
	assert.Panics(t, func() { Validate(linkString) })
}
