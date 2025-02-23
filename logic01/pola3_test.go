package logic01

import (
	"testing"

	"github.com/achmad-aditiansah/latihan/logic"
	"github.com/stretchr/testify/assert"
)

func TestPola3(t *testing.T) {
	result := logic.Logic1Pola3(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[9], 30)
}