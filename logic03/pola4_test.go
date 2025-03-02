package logic03

import (
	"testing"

	"github.com/achmad-aditiansah/latihan/logic"
	"github.com/stretchr/testify/assert"
)

func TestPola4(t *testing.T) {
	n := 9
	result := logic.Logic3Pola4(n)

	assert.Equal(t, len(result), n)
	assert.Equal(t, result[n-9][n-1], 1)
	assert.Equal(t, result[n-1][n-1], 73)
}
