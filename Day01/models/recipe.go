package models

type Recipes struct {
	XMLName string `xml:"recipes" json:"-"`
	Recipe []Recipe `xml:"cake" json:"cake"`
}

type Recipe struct {
	Name        string `xml:"name" json:"name"`
	Time  string `xml:"stovetime" json:"time"`
	Ingredient []Ingredient `xml:"ingredients>item" json:"ingredients"`
}

type Ingredient struct {
	Name  string `xml:"itemname" json:"ingredient_name"`
	Count string `xml:"itemcount" json:"ingredient_count"`
	Unit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
}