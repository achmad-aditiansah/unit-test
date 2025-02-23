package logic01

import (
	"testing"

	"github.com/achmad-aditiansah/latihan/logic"
	"github.com/stretchr/testify/assert"
)

func TestPola8(t *testing.T) {
	result := logic.Logic1Pola8(11)

	assert.Equal(t, len(result), 11)
	assert.Equal(t, result[9], 4)
}