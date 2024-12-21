package main

import (
	"flag"
	"fmt"
	"github.com/rtzgod/go-bootcamp/Day01/internal"
)

func main() {
	var path string

	// Обработка флага -f для пути к файлу
	flag.StringVar(&path, "f", "path", "This flag indicates path to your file")
	flag.Parse()

	// Чтение данных из файла
	data, err := internal.ReadFiles(path)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// хз как назвать
	reader := internal.NewDBReader(path)

	recipe, err := reader.Parse(data[0])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	reader.Print(recipe)
}
