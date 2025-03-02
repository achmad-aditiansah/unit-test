package logic02

import (
	"testing"

	"github.com/achmad-aditiansah/latihan/logic"
	"github.com/stretchr/testify/assert"
)

func TestPola12(t *testing.T) {
	n := 9
	result := logic.Logic2Pola12(n)

	assert.Equal(t, len(result), n)
	assert.Equal(t, result[n-9][n-9], 1)
	assert.Equal(t, result[n-5][n-5], 9)
}
