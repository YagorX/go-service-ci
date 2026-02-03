package model

import (
	"bytes"
	"encoding/json"
)

type Satellite struct {
	Name string
}

func (s *Satellite) Encode() ([]byte, error) {
	res, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Satellite) Decode(in []byte) error {
	var res Satellite
	err := json.NewDecoder(bytes.NewReader(in)).Decode(&res)
	if err != nil {
		return err
	}

	*s = res

	return nil
}
