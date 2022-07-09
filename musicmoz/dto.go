package musicmoz

import (
	"encoding/xml"
)

type RawRecords struct {
	XMLName xml.Name  `xml:"records"`
	Records []*Record `xml:"record"`
}

type Record struct {
	XMLName      xml.Name      `xml:"record"`
	Name         string        `xml:"name"`
	ReleaseDate  string        `xml:"releasedate"`
	TrackListing *TrackListing `xml:"tracklisting"`
}

type TrackListing struct {
	XMLName xml.Name `xml:"tracklisting"`
	Tracks  []string `xml:"track"`
}

type MatchingReleases struct {
	XMLName  xml.Name   `xml:"matchingReleases"`
	Releases []*Release `xml:"release"`
}

type Release struct {
	XMLName    xml.Name `xml:"release"`
	Name       string   `xml:"name"`
	TrackCount int      `xml:"trackCount"`
}
