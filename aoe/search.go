package aoe

type search struct {
	src string
}

func NewSearchApi(src string) *search {
	return &search{
		src: src,
	}
}

func (srch *search) All() (interface{}, error) {
	music := []*Music{
		{
			Name:       "music 1",
			TrackCount: 1,
		},
	}

	return &MatchingReleases{
		Releases: musicDataToDtos(music),
	}, nil
}
