package cal

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	assert.Equal(t, 55, res)
}
