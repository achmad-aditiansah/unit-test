package logic01

import (
	"testing"

	"github.com/achmad-aditiansah/latihan/logic"
	"github.com/stretchr/testify/assert"
)

func TestPola4(t *testing.T) {
	result := logic.Logic1Pola4(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[9], 1)
}