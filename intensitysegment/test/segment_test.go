package test

import (
	"testing"

	is "github.com/lesiona-z/jerryaiinterview/intensitysegment"
	"github.com/stretchr/testify/assert"
)

func TestSetSegment(t *testing.T) {
	segment := is.IntensitySegment{}
	assert.Equal(t, segment.ToString(), "[]")
	err := segment.Set(2, 1, 1)
	assert.Error(t, err)

	// add non-overlapping segments
	err = segment.Set(5, 10, 1)
	assert.Nil(t, err)
	assert.Equal(t, segment.ToString(), "[[5,1], [10,0]]")
	err = segment.Set(15, 20, 2)
	assert.Nil(t, err)
	assert.Equal(t, segment.ToString(), "[[5,1], [10,0], [15,2], [20,0]]")
	err = segment.Set(1, 2, 1)
	assert.Nil(t, err)
	assert.Equal(t, segment.ToString(), "[[1,1], [2,0], [5,1], [10,0], [15,2], [20,0]]")

	// test error cases
	err = segment.Set(8, 12, 3)
	assert.Error(t, err)
	err = segment.Set(3, 8, 4)
	assert.Error(t, err)
	err = segment.Set(3, 12, 5)
	assert.Error(t, err)
	err = segment.Set(6, 8, 6)
	assert.Error(t, err)
}

func TestAddSegment_NonOverlaps(t *testing.T) {
	segment := is.IntensitySegment{}
	assert.Equal(t, "[]", segment.ToString())
	err := segment.Add(2, 1, 1)
	assert.Error(t, err)
	// add non-overlapping segments
	err = segment.Add(5, 10, 1)
	assert.Nil(t, err)
	assert.Equal(t, "[[5,1], [10,0]]", segment.ToString())
	err = segment.Add(15, 20, 2)
	assert.Nil(t, err)
	assert.Equal(t, "[[5,1], [10,0], [15,2], [20,0]]", segment.ToString())
	err = segment.Add(1, 2, 1)
	assert.Nil(t, err)
	assert.Equal(t, "[[1,1], [2,0], [5,1], [10,0], [15,2], [20,0]]", segment.ToString())
	err = segment.Add(11, 14, 1)
	assert.Nil(t, err)
	assert.Equal(t, "[[1,1], [2,0], [5,1], [10,0], [11,1], [14,0], [15,2], [20,0]]", segment.ToString())
	err = segment.Add(10, 11, 1)
	assert.Nil(t, err)
	assert.Equal(t, "[[1,1], [2,0], [5,1], [14,0], [15,2], [20,0]]", segment.ToString())
}

func TestAddSegment_Overlaps_OldContainsNew(t *testing.T) {
	segment := is.IntensitySegment{}
	err := segment.Add(1, 10, 1)
	assert.Nil(t, err)

	err = segment.Add(3, 5, 2)
	assert.Nil(t, err)
	assert.Equal(t, "[[1,1], [3,3], [5,1], [10,0]]", segment.ToString())

	err = segment.Add(3, 5, 1)
	assert.Nil(t, err)
	assert.Equal(t, "[[1,1], [3,4], [5,1], [10,0]]", segment.ToString())
}

func TestAddSegment_Overlaps_NewContainsOld(t *testing.T) {
	segment := is.IntensitySegment{}
	err := segment.Add(5, 10, 1)
	assert.Nil(t, err)

	err = segment.Add(1, 15, 2)
	assert.Nil(t, err)
	assert.Equal(t, "[[1,2], [5,3], [10,2], [15,0]]", segment.ToString())
}

func TestAddSegment_PartialOverlaps_LeftPart(t *testing.T) {
	segment := is.IntensitySegment{}
	err := segment.Add(5, 10, 1)
	assert.Nil(t, err)

	err = segment.Add(1, 7, 2)
	assert.Nil(t, err)
	assert.Equal(t, "[[1,2], [5,3], [7,1], [10,0]]", segment.ToString())

	err = segment.Add(4, 7, 1)
	assert.Nil(t, err)
	assert.Equal(t, "[[1,2], [4,3], [5,4], [7,1], [10,0]]", segment.ToString())
}

func TestAddSegment_PartialOverlaps_RightPart(t *testing.T) {
	segment := is.IntensitySegment{}
	err := segment.Add(5, 10, 1)
	assert.Nil(t, err)

	err = segment.Add(7, 15, 2)
	assert.Nil(t, err)
	assert.Equal(t, "[[5,1], [7,3], [10,2], [15,0]]", segment.ToString())

	err = segment.Add(7, 12, 1)
	assert.Nil(t, err)
	assert.Equal(t, "[[5,1], [7,4], [10,3], [12,2], [15,0]]", segment.ToString())
}
