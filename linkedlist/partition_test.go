package linkedlist

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	head := Create("h", "g", "f", "e", "d", "c", "b", "a")
	head = Partition(head, "d")

	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") < 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") < 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") < 0, head.Value)
	head = head.Next
	assert.Equal(t, "d", head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.Nil(t, head)
}

func TestPartitionExclude(t *testing.T) {
	head := Create("h", "g", "f", "e", "c", "b", "a")
	head = Partition(head, "d")

	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") < 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") < 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") < 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.Nil(t, head)
}

func TestPartitionLower(t *testing.T) {
	head := Create("d", "c", "b", "a")
	head = Partition(head, "d")

	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") < 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") < 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") < 0, head.Value)
	head = head.Next
	assert.Equal(t, "d", head.Value)
	head = head.Next
	assert.Nil(t, head)
}

func TestPartitionHigher(t *testing.T) {
	head := Create("h", "g", "f", "e", "d")
	head = Partition(head, "d")

	assert.Equal(t, "d", head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.Nil(t, head)
}

func TestPartitionLowerExclude(t *testing.T) {
	head := Create("c", "b", "a")
	head = Partition(head, "d")

	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") < 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") < 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") < 0, head.Value)
	head = head.Next
	assert.Nil(t, head)
}

func TestPartitionHigherExclude(t *testing.T) {
	head := Create("h", "g", "f", "e")
	head = Partition(head, "d")

	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.True(t, strings.Compare(fmt.Sprintf("%v", head.Value), "d") > 0, head.Value)
	head = head.Next
	assert.Nil(t, head)
}
