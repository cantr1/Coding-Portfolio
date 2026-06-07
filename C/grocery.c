#include <stdlib.h>

typedef struct GroceryList grocery_t;
typedef struct Recipe recipe_t;
typedef struct Ingredient ingredient_t;

typedef struct GroceryList {
    recipe_t **recipes;
    float total_cost;
    int recipe_count;
} grocery_t;

typedef struct Recipe {
    char *name;
    ingredient_t **ingredient;
    float recipe_cost;
    int num_ingredients;
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

    for (int i = 0; i < g->recipe_count; i++) {
        calculate_recipe_cost(g->recipes[i]);
        g->total_cost = g->total_cost + g->recipes[i]->recipe_cost;
    }

    return;
}

void calculate_recipe_cost(recipe_t *r) {
    if (r == NULL) {
        return;
    }

    r->recipe_cost = 0.0;

    for (int i = 0; i < r->num_ingredients; i++) {
        r->recipe_cost = r->recipe_cost + r->ingredient[i]->cost;
    }

    return;
}

ingredient_t *create_ingredient(char *name, float cost) {
    ingredient_t *obj = malloc(sizeof(ingredient_t));
    if (obj == NULL) {
        return NULL;
    }

    obj->name = name;
    obj->cost = cost;
}

int main(void) {

}