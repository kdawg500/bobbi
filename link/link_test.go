package link

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	linkString := "facebook.com"
	expected := "https://facebook.com"
	actual := Validate(linkString)
	assert.Equal(t, expected, actual)
}
