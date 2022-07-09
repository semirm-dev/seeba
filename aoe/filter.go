package aoe

import (
	"context"
	"encoding/xml"
	"github.com/sirupsen/logrus"
)

type filter struct {
}

func NewFilter() *filter {
	return &filter{}
}

func (ftr *filter) Apply(ctx context.Context, musicData []byte) <-chan []byte {
	filtered := make(chan []byte)

	go func() {
		defer close(filtered)

		select {
		case <-ctx.Done():
			return
		default:
			var rawRecords *RawRecords
			if err := xml.Unmarshal(musicData, &rawRecords); err != nil {
				logrus.Error(err)
				return
			}

			queryResult, err := applyFilter(rawRecords.Records)
			if err != nil {
				logrus.Error(err)
				return
			}

			var releases []*Release
			for _, rec := range queryResult {
				releases = append(releases, &Release{
					Name:       rec.Name,
					TrackCount: len(rec.TrackListing.Tracks),
				})
			}

			bytes, err := xml.Marshal(&MatchingReleases{
				Releases: releases,
			})
			if err != nil {
				logrus.Error(err)
				return
			}

			filtered <- bytes
		}
	}()

	return filtered
}

func applyFilter(records []*Record) ([]*Record, error) {
	var buf []*Record

	for _, r := range records {
		if len(r.TrackListing.Tracks) > 10 {
			buf = append(buf, r)
		}
	}

	return buf, nil
}
