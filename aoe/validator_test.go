package aoe_test

import (
	"github.com/semirm-dev/seeba/aoe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultValidator_Valid(t *testing.T) {
	testTable := map[string]struct {
		given    *aoe.Record
		expected bool
	}{
		"given valid record should return true": {
			given: &aoe.Record{
				Name:        "record 1",
				ReleaseDate: "1999.01.15",
				TrackListing: &aoe.TrackListing{
					Tracks: []string{
						"track 1", "track 2", "track 3", "track 4", "track 5",
						"track 6", "track 7", "track 8", "track 9", "track 10",
					},
				},
			},
			expected: true,
		},
		"given record with less than 10 tracks and newer date should return false": {
			given: &aoe.Record{
				Name:        "record 1",
				ReleaseDate: "2010.01.01",
				TrackListing: &aoe.TrackListing{
					Tracks: []string{
						"track 1", "track 2", "track 3", "track 4", "track 5",
					},
				},
			},
			expected: false,
		},
		"given less than 10 tracks should return false": {
			given: &aoe.Record{
				Name:        "record 1",
				ReleaseDate: "1999.01.01",
				TrackListing: &aoe.TrackListing{
					Tracks: []string{
						"track 1", "track 2", "track 3", "track 4", "track 5",
					},
				},
			},
			expected: false,
		},
		"given record newer than 2001.01.01 should return false": {
			given: &aoe.Record{
				Name:        "record 1",
				ReleaseDate: "2001.01.02",
				TrackListing: &aoe.TrackListing{
					Tracks: []string{
						"track 1", "track 2", "track 3", "track 4", "track 5",
						"track 6", "track 7", "track 8", "track 9", "track 10",
					},
				},
			},
			expected: false,
		},
	}

	validator := aoe.NewDefaultValidator()
	for name, suite := range testTable {
		t.Run(name, func(t *testing.T) {
			valid := validator.Valid(suite.given)
			assert.Equal(t, suite.expected, valid)
		})
	}
}
