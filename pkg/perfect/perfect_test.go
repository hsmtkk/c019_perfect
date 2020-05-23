package perfect_test

import (
	"github.com/hsmtkk/c019_perfect/pkg/perfect"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJudge(t *testing.T) {
	assert.Equal(t, perfect.Perfect, perfect.Judge(28), "should be equal")
	assert.Equal(t, perfect.Nearly, perfect.Judge(16), "should be equal")
	assert.Equal(t, perfect.Neither, perfect.Judge(777), "should be equal")
}

func TestDivisor(t *testing.T) {
	want := []int{1, 2, 4, 8}
	got := perfect.Divisor(16)
	assert.Equal(t, want, got, "should be equal")

	want = []int{1, 2, 4, 7, 14}
	got = perfect.Divisor(28)
	assert.Equal(t, want, got, "should be equal")
}
