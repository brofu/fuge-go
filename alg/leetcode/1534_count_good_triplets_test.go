package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountGoodTriplets(t *testing.T) {

	arr := []int{7, 3, 7, 3, 12, 1, 12, 2, 3}
	a := 5
	b := 8
	c := 1

	result := countGoodTripletsV2(arr, a, b, c)

	assert.Equal(t, 12, result)
}
