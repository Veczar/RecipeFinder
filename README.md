# Recipe Finder

Recipe Finder is a command-line application that helps you find recipes based on the ingredients you have available. The application fetches recipes from the Spoonacular API, provides nutritional information, and stores results locally using SQLite for quicker retrieval.

## Features

- Fetch recipes based on a list of ingredients.
- Display nutritional information including carbs, proteins, and calories.
- Cache recipes in a local SQLite database for faster access on subsequent queries.

## Requirements

- Go 1.16+
- Spoonacular API Key (Sign up at [Spoonacular](https://spoonacular.com/food-api) to get an API key)
- SQLite3

## Usage
- `--ingredients`: Comma-separated list of ingredients you have (e.g., `--ingredients=tomatoes,eggs,pasta`).
- `--numberOfRecipes`: Maximum number of recipes to fetch (e.g., `--numberOfRecipes=5`).

## Examples
Fetch recipes with tomatoes, eggs, and pasta as ingredients, and limit the result to 1 recipe:

`````./recipeFinder fetch --ingredients=tomatoes,eggs,pasta --numberOfRecipes=1 `````
or use this command in main directory
````` go run main.go fetch --ingredients=tomatoes,eggs,pasta --numberOfRecipes=1 `````

## Output
The application will output a list of recipes, including:

- Recipe name
- Ingredients present
- Missing ingredients
- Nutritional information (carbs, proteins, calories)

## Sample Output
```
Recipe: Totally Fresh Tomato Lasagna
Ingredients present:
- egg
- no-boil Lasagna noodles
- tomatoes
Missing ingredients:
- butter
- mozzarella
- milk
- oregano
- parmesan
Nutritional Info: Carbs: 56g, Proteins: 25g, Calories: 540
```

## Database
The application uses SQLite to store cached recipes. The database file is named recipes.db and is located in the root directory.

## Database Schema
The recipes table stores the following information:

- id (INTEGER PRIMARY KEY)
- title (TEXT)
- used_ingredients (TEXT)
- missed_ingredients (TEXT)
- carbs (TEXT)
- proteins (TEXT)
- calories (TEXT)
- ingredients_hash (TEXT UNIQUE)

## Functions
**Create Tables**: The tables are created automatically if they do not exist.

**Cache Recipes**: Recipes fetched from the API are stored in the database.

**Retrieve Cached Recipes**: If the same ingredients are queried again, the cached recipes are used.

## Acknowledgements
- [Spoonacular](https://spoonacular.com/food-api) API for providing the recipe data.
- [Cobra](https://github.com/spf13/cobra) for the CLI framework.
- [Go-SQLite3](https://github.com/mattn/go-sqlite3) for the SQLite driver.
