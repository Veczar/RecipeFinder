package main

import (
	"RecipeFinder/api"
	"RecipeFinder/database"
	"RecipeFinder/utils"
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	_ "modernc.org/sqlite"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "RecipeFinder",
	Short: "Recipe Finder CLI",
	Long:  "A CLI application to find recipes based on provided ingredients.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please use a subcommand.")
	},
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch recipes based on ingredients",
	Run: func(cmd *cobra.Command, args []string) {
		ingredients, _ := cmd.Flags().GetString("ingredients")
		numberOfRecipes, _ := cmd.Flags().GetInt("numberOfRecipes")

		if ingredients == "" {
			fmt.Println("Please provide a list of ingredients")
			os.Exit(1)
		}

		// Connect to the database
		db, err := sql.Open("sqlite", "./recipes.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Create tables if they don't exist
		database.CreateTables(db)

		// Check if the result is already in the database
		cachedRecipes, found := database.GetCachedRecipes(db, ingredients)
		if found {
			utils.PrintRecipes(cachedRecipes)
		} else {
			// Fetch recipes from the API
			recipes, err := api.FetchRecipes(ingredients, numberOfRecipes)
			if err != nil {
				fmt.Printf("Error fetching recipes: %v\n", err)
				os.Exit(1)
			}

			// Store recipes in the database
			database.StoreRecipes(db, ingredients, recipes)

			// Print the recipes
			utils.PrintRecipes(recipes)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().String("ingredients", "", "Comma-separated list of ingredients")
	fetchCmd.Flags().Int("numberOfRecipes", 5, "Maximum number of recipes to retrieve")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		println("Error:", err)
	}
}
