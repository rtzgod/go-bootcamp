package main

import (
	"flag"
	"fmt"
	"github.com/rtzgod/go-bootcamp/Day01/internal"
	"github.com/rtzgod/go-bootcamp/Day01/models"
	"github.com/google/go-cmp/cmp"
)
type DBReader interface {
	Parse(data []byte) (*models.Recipes, error)
	Print()
}

func main() {
	var (
		oldDBPath string
		newDBPath string
	)

	flag.StringVar(&oldDBPath, "old", "", "This flag indicates path to your old db")
	flag.StringVar(&newDBPath, "new", "", "This flag indicates path to your new db")
	flag.Parse()

	data, err := internal.ReadFiles(oldDBPath, newDBPath)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	reader := internal.NewDBReader(oldDBPath)
	oldRecipe, err := reader.Parse(data[0])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	reader = internal.NewDBReader(newDBPath)
	newRecipe, err := reader.Parse(data[1])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

}

func Compare(oldRecipe, newRecipe *models.Recipes) {

}