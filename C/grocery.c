#include <stdlib.h>
#include <stdio.h>

typedef struct GroceryList grocery_t;
typedef struct Recipe recipe_t;
typedef struct Ingredient ingredient_t;

typedef struct GroceryList {
    recipe_t **recipes;
    float total_cost;
    size_t recipe_count;
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

void calculate_total_cost(grocery_t *g) {
    if (g == NULL) {
        return;
    }

    g->total_cost = 0.0;

    for (size_t i = 0; i < g->recipe_count; i++) {
        //calculate_recipe_cost(g->recipes[i]);
        g->total_cost = g->total_cost + g->recipes[i]->recipe_cost;
    }

    return;
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

int main(void) {
    ingredient_t *cucumber_ptr = create_ingredient("cucumber", 1.25);
    ingredient_t *rice_ptr = create_ingredient("rice", 0.75);

    ingredient_t *arr[] = {cucumber_ptr, rice_ptr};

    recipe_t *sample_r = create_recipe("sample", arr, 2);

    free(cucumber_ptr);
    free(rice_ptr);
    free(sample_r);

}