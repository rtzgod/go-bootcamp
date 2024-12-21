package readers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/rtzgod/go-bootcamp/Day01/models"
)

type XMLReader struct {
}

func NewXMLReader() *XMLReader {
	return &XMLReader{}
}

func (r *XMLReader) Parse(data []byte) (*models.Recipes, error) {
	recipe := new(models.Recipes)

	err := xml.Unmarshal(data, recipe)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (r *XMLReader) Print(recipe *models.Recipes) {
	text, _ := json.MarshalIndent(recipe, "", "    ")

	fmt.Println(string(text))
}
