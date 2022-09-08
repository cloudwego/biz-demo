package configparser

import (
	"errors"
	"fmt"
)

type fileProvider struct{}

func NewFile() Provider {
	return &fileProvider{}
}

func (fl *fileProvider) Get() (*Parser, error) {
	fileName := getConfigFlag()
	if fileName == "" {
		return nil, errors.New("config file not specified")
	}

	cp, err := NewParserFromFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error loading config file %q: %v", fileName, err)
	}

	return cp, nil
}
