mod arithemetic; // import other file
mod calculation;
use std::io::{self, Write}; // I/O Library

fn take_input(prompt: String) -> String {
    print!("{}", prompt);
    // ensure prompt displays immediately, not waiting for input
    io::stdout().flush().expect("Failed to flush stdout");

    // Create a mutable to contain inputs
    let mut user_input = String::new();

    // Read from stdin
    io::stdin().read_line(&mut user_input).expect("Failed to read line");

    // Clean input
    let cleaned_input = user_input.trim();

    return cleaned_input.to_string();
}

fn int_to_str(str_int: String) -> i32 {
    // Type conversion to i32, implicit return
    str_int.parse().unwrap()
}

fn main() {
    // Take user input
    let x = int_to_str(take_input("Enter x: ".to_string()));
    let y = int_to_str(take_input("Enter y: ".to_string()));
    let action = take_input("Enter operation: * + - / ".to_string());

    // Create a mutable variable to store result
    let mut result: i32 = 0;

    // Determine result by action
    match action.as_str() {
        "+" => result = arithemetic::add_values(x, y),
        "-" => result = arithemetic::subtract_values(x, y),
        "*" => result = arithemetic::multiply_values(x, y),
        "/" => result = arithemetic::divide_values(x, y),
        _ => println!("Unrecognized input") // The underscore (_) acts as the default case
    }
    
    // Parse inputs into a struct to write to file
    let calc = calculation::build_calculation(x, y, action, result);
    println!("Addition: {}", calc.result)
}
