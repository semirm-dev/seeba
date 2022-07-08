package aoe

import (
	"encoding/xml"
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

type Music struct {
	Name       string
	TrackCount int
}

func musicDataToDtos(data []*Music) []*Release {
	var releases []*Release

	for _, d := range data {
		releases = append(releases, musicDataToDto(d))
	}

	return releases
}

func musicDataToDto(data *Music) *Release {
	return &Release{
		Name:       data.Name,
		TrackCount: data.TrackCount,
	}
}
