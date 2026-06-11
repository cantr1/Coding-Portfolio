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

    //printf("Recipe: %s\n", r->name);

    float recipe_cost = 0.0;

    for (size_t i = 0; i < r->num_ingredients; i++) {
        //printf("Ingredient %s costs: %.2f\n", r->ingredients[i]->name, r->ingredients[i]->cost);
        recipe_cost += r->ingredients[i]->cost;
    }

    //printf("Total Cost: %.2f\n", recipe_cost);
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

void display_grocery_list(grocery_t *g) {
    printf("-----------------------------\n");
    printf("|    Recipes of the Week    |\n");
    printf("-----------------------------\n");
    for (size_t i = 0; i < g->recipe_count; i++) {
        printf("-> %s\n", g->recipes[i]->name);
        for (size_t l = 0; l < g->recipes[i]->num_ingredients; l++) {
            printf("\t|-> %s = %.2f\n", g->recipes[i]->ingredients[l]->name, g->recipes[i]->ingredients[l]->cost);
        }
    }

    printf("-----------------------------\n");
    printf("| Single Items for the Week |\n");
    printf("-----------------------------\n");
    for (size_t j = 0; j < g->item_count; j++) {
        printf("-> %s = %.2f\n", g->items[j]->name, g->items[j]->cost);
    }
}

void display_total_savings(grocery_t *g) {
    /* Return a display of total savings, assuming each meal
    is about $20 to eat out, and where each meal cooked at home
    feeds a person twice, cooking for two */
    printf("************************************************\n");
    printf("| Total Savings Compared to Eating Out: %.2f |\n", (g->recipe_count * 80) - g->total_cost);
    printf("************************************************\n");
}

void free_grocery_memory(grocery_t *g) {
    if (g == NULL) {
        return;
    }

    for (size_t i = 0; i < g->recipe_count; i++) {
        for (size_t j = 0; j < g->recipes[i]->num_ingredients; j++) {
            free(g->recipes[i]->ingredients[j]);
        }
        free(g->recipes[i]->ingredients);
        free(g->recipes[i]);
    }

    for (size_t l = 0; l < g->item_count; l++) {
        free(g->items[l]);
    }

    free(g);
}

recipe_t *create_cnr(void) {
    ingredient_t *chicken_ptr = create_ingredient("chicken", 7.25);
    ingredient_t *rice_ptr = create_ingredient("rice", 0.75);
    ingredient_t *broc_ptr = create_ingredient("broccoli", 4.50);

    ingredient_t **arr = malloc(sizeof(*arr) * 3);
    arr[0] = chicken_ptr;
    arr[1] = rice_ptr;
    arr[2] = broc_ptr;

    recipe_t *obj = malloc(sizeof(recipe_t));

    obj->name = "chicken_n_rice";
    obj->ingredients = arr;
    obj->num_ingredients = 3;
    obj->recipe_cost = calculate_recipe_cost(obj);

    return obj;
}

recipe_t *create_cnp(void) {
    ingredient_t *chicken_ptr = create_ingredient("chicken", 7.25);
    ingredient_t *pasta_ptr = create_ingredient("pasta", 1.25);
    ingredient_t *parm_ptr = create_ingredient("parmesan", 2.50);

    ingredient_t **arr = malloc(sizeof(*arr) * 3);
    arr[0] = chicken_ptr;
    arr[1] = pasta_ptr;
    arr[2] = parm_ptr;

    recipe_t *obj = malloc(sizeof(recipe_t));

    obj->name = "chicken_n_pasta";
    obj->ingredients = arr;
    obj->num_ingredients = 3;
    obj->recipe_cost = calculate_recipe_cost(obj);

    return obj;
}

recipe_t *create_bnb(void) {
    ingredient_t *beef_ptr = create_ingredient("beef", 8.50);
    ingredient_t *broc_ptr = create_ingredient("broccoli", 4.50);
    ingredient_t *soy_ptr = create_ingredient("soy sauce", 0.25);

    ingredient_t **arr = malloc(sizeof(*arr) * 3);
    arr[0] = beef_ptr;
    arr[1] = broc_ptr;
    arr[2] = soy_ptr;

    recipe_t *obj = malloc(sizeof(recipe_t));

    obj->name = "beef_n_broccoli";
    obj->ingredients = arr;
    obj->num_ingredients = 3;
    obj->recipe_cost = calculate_recipe_cost(obj);

    return obj;
}

int main(void) {
    printf("Welcome to the Grocery List Calculator!\n");
    char *options = "Recipes:\n1. Chicken n Rice\n2. Chicken n Pasta\n3. Beef n Broccoli\n\nSingle Items:\nA. Made Good Bars\nM. Milk\nE. Eggs\nB. Bread\nY. Yogurt\nF. Frozen Fruit\nC. Coffee\n\nPress 'x' to calculate list\nPress 'q' to quit.\n";
    printf("%s", options);

    // Create placeholder arrays
    item_t *items_arr[10];
    int num_items = 0;
    recipe_t *rcp[10];
    int num_recipes = 0;

    while (1) {
        switch (getchar()) {
            case '1':
                printf("You selected Chicken n Rice.\n");
                rcp[num_recipes++] = create_cnr();
                break;
            case '2':
                printf("You selected Chicken n Pasta.\n");
                rcp[num_recipes++] = create_cnp();
                break;
            case '3':
                printf("You selected Beef n Broccoli.\n");
                rcp[num_recipes++] = create_bnb();
                break;
            case 'A':
                printf("You selected Made Good Bars.\n");
                items_arr[num_items++] = create_single_item("made good bars", 4.50);
                break;
            case 'M':
                printf("You selected Milk.\n");
                items_arr[num_items++] = create_single_item("milk", 5.25);
                break;
            case 'E':
                printf("You selected Eggs.\n");
                items_arr[num_items++] = create_single_item("eggs", 4.50);
                break;
            case 'B':
                printf("You selected Bread.\n");
                items_arr[num_items++] = create_single_item("bread", 4.15);
                break;
            case 'Y':
                printf("You selected Yogurt.\n");
                items_arr[num_items++] = create_single_item("yogurt", 6.15);
                break;
            case 'F':
                printf("You selected Frozen Fruit.\n");
                items_arr[num_items++] = create_single_item("frozen fruit", 6.15);
                break;
            case 'C':
                printf("You selected Coffee.\n");
                items_arr[num_items++] = create_single_item("coffee", 6.35);
                break;
            case 'x':
                printf("Calculating grocery list...\n");
                grocery_t *g = create_grocery_list(rcp, items_arr, num_recipes, num_items);
                display_grocery_list(g);
                display_total_savings(g);
                free_grocery_memory(g);
                return 0;
            case 'q':
                printf("Exiting program...\n");
                return 0;
            default:
                printf("Invalid input, try again.\n");
        }
        //Remove newline character from input buffer
        getchar();
    }
}