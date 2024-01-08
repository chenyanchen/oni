package elementsdb

import (
	"context"
	"io"

	"gopkg.in/yaml.v3"

	"github.com/chenyanchen/oni/model"
)

type yamlFileKV struct {
	rw io.ReadWriter
}

func NewYamlKV(rw io.ReadWriter) *yamlFileKV {
	return &yamlFileKV{rw: rw}
}

func (s *yamlFileKV) Get(ctx context.Context, k struct{}) ([]*model.Element, error) {
	var result fileFormat

	err := yaml.NewDecoder(s.rw).Decode(&result)

	return result.Elements, err
}

func (s *yamlFileKV) Set(ctx context.Context, k struct{}, elements []*model.Element) error {
	return yaml.NewEncoder(s.rw).Encode(fileFormat{Elements: elements})
}

func (s *yamlFileKV) Del(ctx context.Context, k struct{}) error {
	return nil
}

type fileFormat struct {
	Elements []*model.Element `yaml:"elements"`
}
