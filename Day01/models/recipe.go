package models

type Recipe struct {
	XMLName string `xml:"recipes" json:"-"`
	Cake    []struct {
		Name        string `xml:"name" json:"name"`
		Time  string `xml:"stovetime" json:"time"`
		Ingredients []struct {
			Name  string `xml:"itemname" json:"ingredient_name"`
			Count string `xml:"itemcount" json:"ingredient_count"`
			Unit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
		} `xml:"ingredients>item" json:"ingredients"`
	} `xml:"cake" json:"cake"`
}