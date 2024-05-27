package api

import (
	"RecipeFinder/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const spoonacularAPIKey = "apiKey" // change it to you're own api Key if to work
const spoonacularAPIURL = "https://api.spoonacular.com/recipes/findByIngredients"
const spoonacularNutritionURL = "https://api.spoonacular.com/recipes/%d/nutritionWidget.json"

func FetchRecipes(ingredients string, numberOfRecipes int) ([]models.Recipe, error) {
	params := url.Values{}
	params.Add("ingredients", ingredients)
	params.Add("number", fmt.Sprintf("%d", numberOfRecipes))
	params.Add("apiKey", spoonacularAPIKey)

	resp, err := http.Get(spoonacularAPIURL + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var recipes []models.Recipe
	if err := json.NewDecoder(resp.Body).Decode(&recipes); err != nil {
		return nil, err
	}

	// Fetch nutritional information for each recipe
	for i, recipe := range recipes {
		nutrients, err := fetchNutrients(recipe.ID)
		if err != nil {
			return nil, err
		}
		recipes[i].Carbs = nutrients.Carbs
		recipes[i].Proteins = nutrients.Proteins
		recipes[i].Calories = nutrients.Calories
	}

	return recipes, nil
}

func fetchNutrients(recipeID int) (*models.Nutrients, error) {
	url := fmt.Sprintf(spoonacularNutritionURL, recipeID) + "?apiKey=" + spoonacularAPIKey
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var nutrients models.Nutrients
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&nutrients); err != nil {
		return nil, err
	}

	return &nutrients, nil
}
