package gildedrose_test

import (
	"kata/gildedrose"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExpectedItem struct {
	SellIn  int
	Quality int
}

func TestSulfuras(t *testing.T) {
	assert := assert.New(t)

	items := []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 1, 80},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
	}
	expected := []*ExpectedItem{
		{1, 80},
		{0, 80},
		{-1, 80},
	}

	gildedrose.UpdateQuality(items)

	for i := 0; i < len(items); i++ {
		assert.Equal(items[i].SellIn, expected[i].SellIn)
		assert.Equal(items[i].Quality, expected[i].Quality)
	}
}

func TestAgedBrie(t *testing.T) {
	assert := assert.New(t)

	items := []*gildedrose.Item{
		{"Aged Brie", 1, 0},
		{"Aged Brie", 0, 0},
		{"Aged Brie", -1, 0},
		{"Aged Brie", 1, 50},
		{"Aged Brie", 0, 50},
		{"Aged Brie", -1, 50},
		{"Aged Brie", 0, 49},
		{"Aged Brie", -1, 49},
	}
	expected := []*ExpectedItem{
		{0, 1},
		{-1, 2},
		{-2, 2},
		{0, 50},
		{-1, 50},
		{-2, 50},
		{-1, 50},
		{-2, 50},
	}

	gildedrose.UpdateQuality(items)

	for i := 0; i < len(items); i++ {
		assert.Equal(items[i].SellIn, expected[i].SellIn)
		assert.Equal(items[i].Quality, expected[i].Quality)
	}
}

func TestBackstagePasses(t *testing.T) {
	assert := assert.New(t)

	items := []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 20},
	}
	expected := []*ExpectedItem{
		{14, 21},
		{9, 22},
		{9, 50},
		{4, 23},
		{4, 50},
		{-1, 0},
	}

	gildedrose.UpdateQuality(items)

	for i := 0; i < len(items); i++ {
		assert.Equal(items[i].SellIn, expected[i].SellIn)
		assert.Equal(items[i].Quality, expected[i].Quality)
	}
}

func TestConjured(t *testing.T) {
	assert := assert.New(t)

	items := []*gildedrose.Item{
		{"Conjured", 3, 6},
		{"Conjured", 0, 6},
		{"Conjured", -1, 6},
		{"Conjured", 3, 0},
	}
	expected := []*ExpectedItem{
		{2, 4},
		{-1, 2},
		{-2, 2},
		{2, 0},
	}

	gildedrose.UpdateQuality(items)

	for i := 0; i < len(items); i++ {
		assert.Equal(items[i].SellIn, expected[i].SellIn)
		assert.Equal(items[i].Quality, expected[i].Quality)
	}
}

func TestNormal(t *testing.T) {
	assert := assert.New(t)

	items := []*gildedrose.Item{
		{"Normal", 10, 20},
		{"Normal", 0, 20},
		{"Normal", -1, 20},
		{"Normal", 5, 0},
	}
	expected := []*ExpectedItem{
		{9, 19},
		{-1, 18},
		{-2, 18},
		{4, 0},
	}

	gildedrose.UpdateQuality(items)

	for i := 0; i < len(items); i++ {
		assert.Equal(items[i].SellIn, expected[i].SellIn)
		assert.Equal(items[i].Quality, expected[i].Quality)
	}
}
