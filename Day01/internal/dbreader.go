package internal

import (
	"github.com/rtzgod/go-bootcamp/Day01/internal/readers"
	"github.com/rtzgod/go-bootcamp/Day01/models"
	"path/filepath"
)
type DBReader interface {
	Parse(data []byte) (*models.Recipes, error)
	Print(recipe *models.Recipes)
}


func NewDBReader(path string) DBReader {
	fileExtension := filepath.Ext(path)

	switch fileExtension {
	case ".xml":
		return readers.NewXMLReader()
	case ".json":
		return readers.NewJSONReader()
	default:
		return nil
	}
}