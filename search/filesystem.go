package search

import "github.com/semirm-dev/seeba/etl"

type filesystem struct {
	src string
}

func NewFileSystem(src string) *filesystem {
	return &filesystem{
		src: src,
	}
}

func (srch *filesystem) All() ([]*etl.Music, error) {
	return nil, nil
}
