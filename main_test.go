package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	query := "iphone"
	key := "All"
	records := []Record{
		{
			ID:      63983811,
			Title:   "Be discovered",
			Content: "Apple Search Ads helps people discover your app when they search on the. App Store, matching customers with your app right when they’re looking.",
			Tags:    []string{"Apple"},
		},
		{
			ID:      63983811,
			Title:   "IPhone 12 Pro 256",
			Content: "IPhone 12 Pro 256 iphone 12 pro The device is still new",
			Tags:    []string{"Iphone"},
		},
		{
			ID:      63983811,
			Title:   "(Renewed) Apple iPhone 8, US Version, 64GB, Space Gray - Unlocked",
			Content: "IPhone 8 for sale, very clean, space 256, what maintenance income (the remaining specifications are in the pictures) Price: on the price tag (Accept instead of iPhone XS, iPhone XR, and iPhone 8 Plus) They are 128 and more",
			Tags:    []string{"Iphone"},
		},
	}
	expected := []Record{
		{
			ID:      63983811,
			Title:   "(Renewed) Apple iPhone 8, US Version, 64GB, Space Gray - Unlocked",
			Content: "IPhone 8 for sale, very clean, space 256, what maintenance income (the remaining specifications are in the pictures) Price: on the price tag (Accept instead of iPhone XS, iPhone XR, and iPhone 8 Plus) They are 128 and more",
			Tags:    []string{"Iphone"},
			Matches: 6,
		},
		{
			ID:      63983811,
			Title:   "IPhone 12 Pro 256",
			Content: "IPhone 12 Pro 256 iphone 12 pro The device is still new",
			Tags:    []string{"Iphone"},
			Matches: 4,
		},
	}
	searcher := Searcher{}
	searcher.records = records

	result, _ := searcher.filter(query, key)

	assert.Equal(t, result, expected, "The result should be the same as expected.")
}
