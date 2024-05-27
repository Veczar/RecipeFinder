package models

type Item struct {
	Name string `json:"name"`
}

type Nutrients struct {
	Carbs    string `json:"carbs"`
	Proteins string `json:"protein"`
	Calories string `json:"calories"`
}

type Recipe struct {
	ID                int    `json:"id"`
	Title             string `json:"title"`
	UsedIngredients   []Item `json:"usedIngredients"`
	MissedIngredients []Item `json:"missedIngredients"`
	Carbs             string `json:"carbs,omitempty"`
	Proteins          string `json:"proteins,omitempty"`
	Calories          string `json:"calories,omitempty"`
}
