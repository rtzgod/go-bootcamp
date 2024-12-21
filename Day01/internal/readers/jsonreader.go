package readers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/rtzgod/go-bootcamp/Day01/models"
)

type JSONReader struct {
}

func NewJSONReader() *JSONReader {
	return &JSONReader{}
}

func (r *JSONReader) Parse(data []byte) (*models.Recipes, error) {
	recipe := new(models.Recipes)

	err := json.Unmarshal(data, recipe)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (r *JSONReader) Print(recipe *models.Recipes) {
	text, _ := xml.MarshalIndent(recipe, "", "    ")

	fmt.Println(string(text))
}