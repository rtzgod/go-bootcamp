package internal

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/rtzgod/go-bootcamp/Day01/models"
)

type XMLReader struct {
	recipe *models.Recipe
}

func NewXMLReader() *XMLReader {
	return &XMLReader{}
}

func (r *XMLReader) Parse(data []byte) (*models.Recipe, error) {
	r.recipe = new(models.Recipe)

	err := xml.Unmarshal(data, r.recipe)
	if err != nil {
		return nil, err
	}

	return r.recipe, nil
}

func (r *XMLReader) Print() {
	text, _ := json.MarshalIndent(r.recipe, "", "    ")

	fmt.Println(string(text))
}
