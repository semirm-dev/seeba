package musicmoz

import (
	"encoding/xml"
	"io/ioutil"
)

// search will search for music data previously exported to a destination
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

	var music MatchingReleases
	if err = xml.Unmarshal(data, &music); err != nil {
		return nil, err
	}

	return music, nil
}
