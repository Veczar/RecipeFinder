package utils

import (
	"RecipeFinder/models"
	"fmt"
	"strings"
)

func HashIngredients(ingredients string) string {
	return strings.ToLower(strings.ReplaceAll(ingredients, ",", "_"))
}

func JoinIngredients(items []models.Item) string {
	var names []string
	for _, item := range items {
		names = append(names, item.Name)
	}
	return strings.Join(names, ",")
}

func ParseIngredients(ingredients string) []models.Item {
	var items []models.Item
	for _, name := range strings.Split(ingredients, ",") {
		items = append(items, models.Item{Name: name})
	}
	return items
}

func PrintRecipes(recipes []models.Recipe) {
	for _, recipe := range recipes {
		fmt.Printf("Recipe: %s\n", recipe.Title)
		fmt.Println("Ingredients present:")
		for _, item := range recipe.UsedIngredients {
			fmt.Printf("- %s\n", item.Name)
		}
		fmt.Println("Missing ingredients:")
		for _, item := range recipe.MissedIngredients {
			fmt.Printf("- %s\n", item.Name)
		}
		fmt.Printf("Nutritional Info: Carbs: %s, Proteins: %s, Calories: %s\n\n",
			recipe.Carbs, recipe.Proteins, recipe.Calories)
	}
}
