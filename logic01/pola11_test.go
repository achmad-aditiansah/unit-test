package logic01

import (
	"testing"

	"github.com/achmad-aditiansah/latihan/logic"
	"github.com/stretchr/testify/assert"
)

func TestPola11(t *testing.T) {
	result := logic.Logic1Pola11(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[0], "buzz")
}
