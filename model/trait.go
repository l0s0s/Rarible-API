package model

type TraitProperty struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TraitRarity struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Rarity string `json:"rarity"`
}

type TraitsRarity struct {
	Traits []TraitRarity `json:"traits"`
}
