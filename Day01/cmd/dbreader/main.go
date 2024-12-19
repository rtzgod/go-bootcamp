package main

import (
	"flag"
	"fmt"
	"github.com/rtzgod/go-bootcamp/Day01/internal"
	"github.com/rtzgod/go-bootcamp/Day01/models"
	"os"
	"path/filepath"
)

type DBReader interface {
	Parse(data []byte) (*models.Recipe, error)
	Print()
}

func main() {
	var path string

	// Обработка флага -f для пути к файлу
	flag.StringVar(&path, "f", "path", "This flag indicates path to your file")
	flag.Parse()

	// Чтение данных из файла
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// хз как назвать
	reader := NewDBReader(path)


	if _, err := reader.Parse(data); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	reader.Print()
}

func NewDBReader(path string) DBReader {
	fileExtension := filepath.Ext(path)

	switch fileExtension {
	case ".xml":
		return internal.NewXMLReader()
	case ".json":
		return internal.NewJSONReader()
	default:
		return nil
	}
}