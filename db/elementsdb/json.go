package elementsdb

import (
	"context"
	"encoding/json"
	"io"

	"github.com/chenyanchen/oni/model"
)

type jsonFileKV struct {
	rw io.ReadWriter
}

func NewJsonKV(rw io.ReadWriter) *jsonFileKV {
	return &jsonFileKV{rw: rw}
}

func (s *jsonFileKV) Get(ctx context.Context, k struct{}) ([]*model.Element, error) {
	var result fileFormat

	err := json.NewDecoder(s.rw).Decode(&result)

	return result.Elements, err
}

func (s *jsonFileKV) Set(ctx context.Context, k struct{}, elements []*model.Element) error {
	return json.NewEncoder(s.rw).Encode(fileFormat{Elements: elements})
}

func (s *jsonFileKV) Del(ctx context.Context, k struct{}) error {
	return nil
}
