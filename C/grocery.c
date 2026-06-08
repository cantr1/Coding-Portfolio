#include <stdlib.h>
#include <stdio.h>

typedef struct GroceryList grocery_t;
typedef struct Recipe recipe_t;
typedef struct Ingredient ingredient_t;
typedef struct SingleItem item_t;

typedef struct GroceryList {
    recipe_t **recipes;
    item_t **items;
    float total_cost;
    size_t recipe_count;
    size_t item_count;
} grocery_t;

typedef struct Recipe {
    char *name;
    ingredient_t **ingredients;
    float recipe_cost;
    size_t num_ingredients;
} recipe_t;

typedef struct Ingredient {
    char *name;
    float cost;
} ingredient_t;

typedef struct SingleItem {
    char *name;
    float cost;
} item_t;


float calculate_total_cost(grocery_t *g) {
    if (g == NULL) {
        return 0.0;
    }

    float total_cost = 0.0;

    for (size_t i = 0; i < g->recipe_count; i++) {
        total_cost = total_cost + g->recipes[i]->recipe_cost;
    }

    for (size_t j = 0; j < g->item_count; j++) {
        total_cost = total_cost + g->items[j]->cost;
    }

    printf("--------------------------------------\n| Total Cost for grocery list: %.2f |\n--------------------------------------\n", total_cost);

    return total_cost;
}

float calculate_recipe_cost(recipe_t *r) {
    if (r == NULL) {
        return 0.0;
    }

    printf("Recipe: %s\n", r->name);

    float recipe_cost = 0.0;

    for (size_t i = 0; i < r->num_ingredients; i++) {
        printf("Ingredient %s costs: %.2f\n", r->ingredients[i]->name, r->ingredients[i]->cost);
        recipe_cost += r->ingredients[i]->cost;
    }

    printf("Total Cost: %.2f\n", recipe_cost);
    return recipe_cost;
}


item_t *create_single_item(char *name, float cost) {
    item_t *obj = malloc(sizeof(item_t));
    if (obj == NULL) {
        return NULL;
    }
    obj->name = name;
    obj->cost = cost;

    return obj;
}

ingredient_t *create_ingredient(char *name, float cost) {
    ingredient_t *obj = malloc(sizeof(ingredient_t));
    if (obj == NULL) {
        return NULL;
    }

    obj->name = name;
    obj->cost = cost;

    return obj;
}

recipe_t *create_recipe(char *name, ingredient_t **ingredients, size_t num_ingredients) {
    recipe_t *obj = malloc(sizeof(recipe_t));
    if (obj == NULL) {
        return NULL;
    }
    obj->name = name;
    obj->ingredients = ingredients;
    obj->num_ingredients = num_ingredients;
    obj->recipe_cost = calculate_recipe_cost(obj);

    return obj;
}

grocery_t *create_grocery_list(recipe_t **reciepes, item_t **items, size_t recipe_count, size_t item_count) {
    grocery_t *obj = malloc(sizeof(grocery_t));
    if (obj == NULL) {
        return NULL;
    }

    obj->recipes = reciepes;
    obj->items = items;
    obj->recipe_count = recipe_count;
    obj->item_count = item_count;
    obj->total_cost = calculate_total_cost(obj);

    return obj;
}

void free_grocery_memory(grocery_t *g) {
    for (size_t i = 0; i < g->recipe_count; i++) {
        for (size_t j = 0; j < g->recipes[i]->num_ingredients; j++) {
            free(g->recipes[i]->ingredients[j]);
        }
        free(g->recipes[i]);
    }

    for (size_t l = 0; l < g->item_count; l++) {
        free(g->items[l]);
    }

    free(g);
}

int main(void) {
    // TODO: Make a TUI to select recipes, allow the user to continue adding
    // items until satisfied, then calculate

    // Create Individual item
    item_t *made_good_ptr = create_single_item("made good bars", 4.50);

    // Create Array of Individual Items
    item_t *items_arr[] = {made_good_ptr};

    // Create individual ingredients
    ingredient_t *chicken_ptr1 = create_ingredient("chicken", 7.25);
    ingredient_t *chicken_ptr2 = create_ingredient("chicken", 7.25);
    ingredient_t *rice_ptr = create_ingredient("rice", 0.75);
    ingredient_t *broc_ptr = create_ingredient("broccoli", 4.50);

    ingredient_t *pasta_ptr = create_ingredient("pasta", 1.25);
    ingredient_t *parm_ptr = create_ingredient("parmesan", 2.50);

    // Create array of ingredients for recipe
    ingredient_t *cnr[3] = {chicken_ptr1, rice_ptr, broc_ptr};
    ingredient_t *cnp[3] = {chicken_ptr2, pasta_ptr, parm_ptr};

    // Create recipes
    recipe_t *chicken_n_rice = create_recipe("chicken_n_rice", cnr, 3);
    recipe_t *chicken_pasta = create_recipe("chicken_pasta", cnp, 3);

    // Create array of recipes
    recipe_t *rcp[] = {chicken_n_rice, chicken_pasta};

    // Create grocery list
    grocery_t *g = create_grocery_list(rcp, items_arr, 2, 1);

    free_grocery_memory(g);
}