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

float calculate_total_cost(grocery_t *g) {
    if (g == NULL) {
        return 0.0;
    }

    float total_cost = 0.0;

    for (size_t i = 0; i < g->recipe_count; i++) {
        total_cost = total_cost + g->recipes[i]->recipe_cost;
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

grocery_t *create_grocery_list(recipe_t **reciepes, size_t recipe_count) {
    grocery_t *obj = malloc(sizeof(grocery_t));
    if (obj == NULL) {
        return NULL;
    }

    obj->recipes = reciepes;
    obj->recipe_count = recipe_count;
    obj->total_cost = calculate_total_cost(obj);

    return obj;
}

int main(void) {
    ingredient_t *chicken_ptr = create_ingredient("chicken", 7.25);
    ingredient_t *rice_ptr = create_ingredient("rice", 0.75);
    ingredient_t *broc_ptr = create_ingredient("broccoli", 4.50);

    ingredient_t *pasta_ptr = create_ingredient("pasta", 1.25);
    ingredient_t *parm_ptr = create_ingredient("parmesan", 2.50);

    ingredient_t *cnr[3] = {chicken_ptr, rice_ptr, broc_ptr};
    ingredient_t *cnp[3] = {chicken_ptr, pasta_ptr, parm_ptr};

    recipe_t *chicken_n_rice = create_recipe("chicken_n_rice", cnr, 3);
    recipe_t *chicken_pasta = create_recipe("chicken_pasta", cnp, 3);

    recipe_t *rcp[] = {chicken_n_rice, chicken_pasta};

    grocery_t *g = create_grocery_list(rcp, 2);

    free(chicken_ptr);
    free(rice_ptr);
    free(broc_ptr);
    free(pasta_ptr);
    free(parm_ptr);
    free(chicken_n_rice);
    free(chicken_pasta);
    free(g);

}