package models

type Flavortext struct {
	Flavortext string `json:"flavortext"`
	Language   string `json:"language"`
}
type Name struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}
type Pokemon struct {
	Name           string       `json:"name"`
	Names          []Name       `json:"names"`
	Type1          string       `json:"type1"`
	Type2          string       `json:"type2"`
	Generation     int          `json:"generation"`
	EvolutionState int          `json:"evolutionState"`
	IsFullyEvolved string       `json:"isFullyEvolved"`
	Habitat        string       `json:"habitat"`
	Color          string       `json:"color"`
	FlavorTexts    []Flavortext `json:"flavortexts"`
}
