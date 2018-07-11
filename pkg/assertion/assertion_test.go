package assertion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidOperator(t *testing.T) {
	_, err := NewOperator("UnavailableOp")
	assert.NotNil(t, err)
}

func TestEquals(t *testing.T) {
	op, err := NewOperator(EqualOp)
	assert.Nil(t, err)

	r, err := op("val", "val")
	assert.Nil(t, err)
	assert.True(t, r)

	r, err = op("val", "val2")
	assert.NotNil(t, err)
	assert.False(t, r)
}

func TestNotEquals(t *testing.T) {
	op, err := NewOperator(NotEqualOp)
	assert.Nil(t, err)

	r, err := op("hello", "hallo")
	assert.Nil(t, err)
	assert.True(t, r)

	r, err = op("hello", "hello")
	assert.NotNil(t, err)
	assert.False(t, r)
}

func TestAvailableOps(t *testing.T) {
	assert.Len(t, AvailableOps(), len(opMap))
}

func TestContains(t *testing.T) {
	op, err := NewOperator(ContainsOp)
	assert.Nil(t, err)

	r, err := op("hello", "hell")
	assert.Nil(t, err)
	assert.True(t, r)

	r, err = op("hello", "hallo")
	assert.NotNil(t, err)
	assert.False(t, r)
}

func TestNotContains(t *testing.T) {
	op, err := NewOperator(NotContainsOp)
	assert.Nil(t, err)

	r, err := op("hello", "hallo")
	assert.Nil(t, err)
	assert.True(t, r)

	r, err = op("hello", "hell")
	assert.NotNil(t, err)
	assert.False(t, r)
}
