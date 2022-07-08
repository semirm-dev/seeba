package aol

import "github.com/semirm-dev/seeba/etl"

type search struct {
	src string
}

func NewSearchApi(src string) *search {
	return &search{
		src: src,
	}
}

func (srch *search) All() (interface{}, error) {
	music := []*etl.Music{
		{
			Name:       "music 1",
			TrackCount: 1,
		},
	}

	return &MatchingReleases{
		Releases: musicDataToDtos(music),
	}, nil
}
