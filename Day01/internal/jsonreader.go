package internal

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/rtzgod/go-bootcamp/Day01/models"
)

type JSONReader struct {
	recipe *models.Recipe
}

func NewJSONReader() *JSONReader {
	return &JSONReader{}
}

func (r *JSONReader) Parse(data []byte) (*models.Recipe, error) {
	r.recipe = new(models.Recipe)

	err := json.Unmarshal(data, r.recipe)
	if err != nil {
		return nil, err
	}

	return r.recipe, nil
}

func (r *JSONReader) Print() {
	text, _ := xml.MarshalIndent(r.recipe, "", "    ")

	fmt.Println(string(text))
}