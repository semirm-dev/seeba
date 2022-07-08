package aoe

import (
	"encoding/xml"
	"io/ioutil"
)

type search struct {
	src string
}

func NewSearchApi(src string) *search {
	return &search{
		src: src,
	}
}

func (srch *search) All() (interface{}, error) {
	data, err := ioutil.ReadFile(srch.src)
	if err != nil {
		return nil, err
	}

	var music []*Music
	if err = xml.Unmarshal(data, &music); err != nil {
		return nil, err
	}

	return &MatchingReleases{
		Releases: musicDataToDtos(music),
	}, nil
}
