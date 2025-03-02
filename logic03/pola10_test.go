package logic03

import (
	"testing"

	"github.com/achmad-aditiansah/latihan/logic"
	"github.com/stretchr/testify/assert"
)

func TestPola10(t *testing.T) {
	n := 9
	result := logic.Logic3Pola10(n)

	assert.Equal(t, len(result), n)
	assert.Equal(t, result[n-5][n-9], 9)
	assert.Equal(t, result[n-5][n-5], 1)
}
