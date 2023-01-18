package models

type Pokemon struct {
	Name           string `json:"name"`
	Type1          string `json:"type1"`
	Type2          string `json:"type2"`
	Generation     int    `json:"generation"`
	EvolutionState int    `json:"evolutionState"`
	IsFullyEvolved string `json:"isFullyEvolved"`
}
