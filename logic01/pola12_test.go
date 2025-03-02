package logic01

import (
	"testing"

	"github.com/achmad-aditiansah/latihan/logic"
	"github.com/stretchr/testify/assert"
)

func TestPola12(t *testing.T) {
	result := logic.Logic1Pola12(12)

	assert.Equal(t, len(result), 12)
	assert.Equal(t, result[11], 7)
}
