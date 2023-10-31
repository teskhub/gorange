package gorange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetDisplayWidth_correct(t *testing.T) {
	r := New(50, 10)
	var err error

	err = r.SetDisplayWidth(5)
	assert.Nil(t, err)

	err = r.SetDisplayWidth(9)
	assert.Nil(t, err)

	err = r.SetDisplayWidth(13)
	assert.Nil(t, err)

	err = r.SetDisplayWidth(17)
	assert.Nil(t, err)

	err = r.SetDisplayWidth(21)
	assert.Nil(t, err)

	err = r.SetDisplayWidth(25)
	assert.Nil(t, err)

	err = r.SetDisplayWidth(29)
	assert.Nil(t, err)
}

func Test_SetDisplayWidth_incorrect(t *testing.T) {
	r := New(50, 10)
	var err error

	err = r.SetDisplayWidth(-3)
	assert.NotNil(t, err)

	err = r.SetDisplayWidth(-1)
	assert.NotNil(t, err)

	err = r.SetDisplayWidth(0)
	assert.NotNil(t, err)

	err = r.SetDisplayWidth(1)
	assert.NotNil(t, err)

	err = r.SetDisplayWidth(2)
	assert.NotNil(t, err)

	err = r.SetDisplayWidth(3)
	assert.NotNil(t, err)

	err = r.SetDisplayWidth(4)
	assert.NotNil(t, err)

	err = r.SetDisplayWidth(6)
	assert.NotNil(t, err)

	err = r.SetDisplayWidth(7)
	assert.NotNil(t, err)

	err = r.SetDisplayWidth(8)
	assert.NotNil(t, err)
}

func Test_GetRangeAsText_standard(t *testing.T) {
	r := New(50, 10)

	var text string

	text = r.GetRangeAsText(28)
	assert.Equal(t, "-----[----|----]-----", text)

	text = r.GetRangeAsText(30)
	assert.Equal(t, "█----[----|----]-----", text)

	text = r.GetRangeAsText(32)
	assert.Equal(t, "-█---[----|----]-----", text)

	text = r.GetRangeAsText(34)
	assert.Equal(t, "--█--[----|----]-----", text)

	text = r.GetRangeAsText(36)
	assert.Equal(t, "---█-[----|----]-----", text)

	text = r.GetRangeAsText(38)
	assert.Equal(t, "----█[----|----]-----", text)

	text = r.GetRangeAsText(40)
	assert.Equal(t, "-----█----|----]-----", text)

	text = r.GetRangeAsText(42)
	assert.Equal(t, "-----[█---|----]-----", text)

	text = r.GetRangeAsText(44)
	assert.Equal(t, "-----[-█--|----]-----", text)

	text = r.GetRangeAsText(46)
	assert.Equal(t, "-----[--█-|----]-----", text)

	text = r.GetRangeAsText(48)
	assert.Equal(t, "-----[---█|----]-----", text)

	text = r.GetRangeAsText(50)
	assert.Equal(t, "-----[----█----]-----", text)

	text = r.GetRangeAsText(52)
	assert.Equal(t, "-----[----|█---]-----", text)

	text = r.GetRangeAsText(54)
	assert.Equal(t, "-----[----|-█--]-----", text)

	text = r.GetRangeAsText(56)
	assert.Equal(t, "-----[----|--█-]-----", text)

	text = r.GetRangeAsText(58)
	assert.Equal(t, "-----[----|---█]-----", text)

	text = r.GetRangeAsText(60)
	assert.Equal(t, "-----[----|----█-----", text)

	text = r.GetRangeAsText(62)
	assert.Equal(t, "-----[----|----]█----", text)

	text = r.GetRangeAsText(64)
	assert.Equal(t, "-----[----|----]-█---", text)

	text = r.GetRangeAsText(66)
	assert.Equal(t, "-----[----|----]--█--", text)

	text = r.GetRangeAsText(68)
	assert.Equal(t, "-----[----|----]---█-", text)

	text = r.GetRangeAsText(70)
	assert.Equal(t, "-----[----|----]----█", text)

	text = r.GetRangeAsText(72)
	assert.Equal(t, "-----[----|----]-----", text)
}

func Test_GetRangeAsText_border_values(t *testing.T) {
	r := New(50, 10)

	var text string

	text = r.GetRangeAsText(28.99)
	assert.Equal(t, "-----[----|----]-----", text)

	text = r.GetRangeAsText(29)
	assert.Equal(t, "█----[----|----]-----", text)

	text = r.GetRangeAsText(30.99)
	assert.Equal(t, "█----[----|----]-----", text)

	text = r.GetRangeAsText(31.00)
	assert.Equal(t, "-█---[----|----]-----", text)

	text = r.GetRangeAsText(69.00)
	assert.Equal(t, "-----[----|----]----█", text)

	text = r.GetRangeAsText(70.99)
	assert.Equal(t, "-----[----|----]----█", text)

	text = r.GetRangeAsText(71)
	assert.Equal(t, "-----[----|----]-----", text)
}

func Test_GetRangeAsText_low_range_and_negative(t *testing.T) {
	r := New(-8, 0.5)
	r.SetDisplayWidth(5)

	var text string

	text = r.GetRangeAsText(-9.3)
	assert.Equal(t, "-[|]-", text)

	text = r.GetRangeAsText(-8.8)
	assert.Equal(t, "█[|]-", text)

	text = r.GetRangeAsText(-8.5)
	assert.Equal(t, "-█|]-", text)

	text = r.GetRangeAsText(-8.2)
	assert.Equal(t, "-[█]-", text)

	text = r.GetRangeAsText(-7.7)
	assert.Equal(t, "-[|█-", text)

	text = r.GetRangeAsText(-7.2)
	assert.Equal(t, "-[|]█", text)

	text = r.GetRangeAsText(-6.7)
	assert.Equal(t, "-[|]-", text)
}
