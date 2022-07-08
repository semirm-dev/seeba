package aol

import (
	"encoding/xml"
	"github.com/semirm-dev/seeba/etl"
)

type MatchingReleases struct {
	XMLName  xml.Name   `xml:"matchingReleases"`
	Releases []*Release `xml:"releases"`
}

type Release struct {
	XMLName    xml.Name `xml:"release"`
	Name       string   `xml:"name"`
	TrackCount int      `xml:"trackCount"`
}

func musicDataToDtos(data []*etl.Music) []*Release {
	var releases []*Release

	for _, d := range data {
		releases = append(releases, musicDataToDto(d))
	}

	return releases
}

func musicDataToDto(data *etl.Music) *Release {
	return &Release{
		Name:       data.Name,
		TrackCount: data.TrackCount,
	}
}
