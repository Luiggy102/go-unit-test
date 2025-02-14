package sometests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSuccessTestify(t *testing.T) {
	c := require.New(t) // easy asertions
	result := Add(20, 2)
	expect := 22

	c.Equal(expect, result)
}
