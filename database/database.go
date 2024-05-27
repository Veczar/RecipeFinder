package database

import (
	"RecipeFinder/models"
	"RecipeFinder/utils"
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) {
	createRecipesTable := `
    CREATE TABLE IF NOT EXISTS recipes (
        id INTEGER PRIMARY KEY,
        title TEXT,
        used_ingredients TEXT,
        missed_ingredients TEXT,
        carbs TEXT,
        proteins TEXT,
        calories TEXT,
        ingredients_hash TEXT UNIQUE
    );`

	createIngredientsIndex := `
    CREATE UNIQUE INDEX IF NOT EXISTS idx_ingredients_hash ON recipes (ingredients_hash);`

	if _, err := db.Exec(createRecipesTable); err != nil {
		log.Fatalf("Error creating recipes table: %v", err)
	}

	if _, err := db.Exec(createIngredientsIndex); err != nil {
		log.Fatalf("Error creating ingredients index: %v", err)
	}
}

func GetCachedRecipes(db *sql.DB, ingredients string) ([]models.Recipe, bool) {
	ingredientsHash := utils.HashIngredients(ingredients)
	query := `SELECT title, used_ingredients, missed_ingredients, carbs, proteins, calories FROM recipes WHERE ingredients_hash = ?`
	rows, err := db.Query(query, ingredientsHash)
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}
	defer rows.Close()

	var recipes []models.Recipe
	for rows.Next() {
		var title, usedIngredients, missedIngredients, carbs, proteins, calories string
		if err := rows.Scan(&title, &usedIngredients, &missedIngredients, &carbs, &proteins, &calories); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}

		recipes = append(recipes, models.Recipe{
			Title:             title,
			UsedIngredients:   utils.ParseIngredients(usedIngredients),
			MissedIngredients: utils.ParseIngredients(missedIngredients),
			Carbs:             carbs,
			Proteins:          proteins,
			Calories:          calories,
		})
	}

	return recipes, len(recipes) > 0
}

func StoreRecipes(db *sql.DB, ingredients string, recipes []models.Recipe) {
	for _, recipe := range recipes {
		ingredientsHash := utils.HashIngredients(ingredients + recipe.Title)
		usedIngredients := utils.JoinIngredients(recipe.UsedIngredients)
		missedIngredients := utils.JoinIngredients(recipe.MissedIngredients)

		query := `
        INSERT OR IGNORE INTO recipes (title, used_ingredients, missed_ingredients, carbs, proteins, calories, ingredients_hash)
        VALUES (?, ?, ?, ?, ?, ?, ?)`
		_, err := db.Exec(query, recipe.Title, usedIngredients, missedIngredients, recipe.Carbs, recipe.Proteins, recipe.Calories, ingredientsHash)
		if err != nil {
			log.Fatalf("Error inserting recipe: %v", err)
		}
	}
}
