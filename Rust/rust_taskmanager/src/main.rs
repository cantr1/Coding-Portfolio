use std::io;
use std::io::Write;
use std::path::Path;
mod task;

fn take_user_input(prompt: String) -> String {
    print!("{}", prompt);
    // ensure prompt displays immediately, not waiting for input
    io::stdout().flush().expect("Failed to flush stdout");

    // Create a mutable to contain inputs
    let mut user_input = String::new();

    // Read from stdin
    io::stdin()
        .read_line(&mut user_input)
        .expect("Failed to read line");

    // Clean input
    let cleaned_input = user_input.trim();

    return cleaned_input.to_string();
}

fn create_new_task(task_stats: &mut task::TaskStatistics, task_list: &mut Vec<task::Task>) {
    // Get description of task from user
    let task_description = take_user_input("Enter task description: ".to_string());

    // Create the task
    let new_task: task::Task = task::build_task(task_stats.next_task, &task_description, false);

    // Add the task to the list
    task_list.push(new_task);

    // Increment task statistics
    task::increment_next_task(task_stats);
}

fn view_task_list(task_list: &Vec<task::Task>) {
    println!("");
    for task in task_list {
        println!(
            "[{}] ID: {} - {}",
            if task.complete { "*" } else { " " },
            task.id,
            task.description
        )
    }
    println!("");
}

fn main() {
    // Initialize task statistics
    let task_stats_file = Path::new(
        "/home/kelz/Work/Coding-Portfolio/Rust/rust_taskmanager/artifacts/task_statistics.json",
    );
    let mut task_stats = task::build_task_statistics(task_stats_file);

    // Initialize task list
    let task_list_file = Path::new(
        "/home/kelz/Work/Coding-Portfolio/Rust/rust_taskmanager/artifacts/task_list.json",
    );
    let mut task_list = task::build_task_list(task_list_file);

    'main_loop: loop {
        let user_choice = take_user_input(
            "Choices\n-------\nd - display tasks\nn - create new task\nq - quit\n~:".to_string(),
        );
        match user_choice.as_str() {
            "d" => view_task_list(&task_list),
            "n" => create_new_task(&mut task_stats, &mut task_list),
            "q" => break 'main_loop,
            _ => println!("unrecognized input"),
        }
    }

    // End of operations, write statistics and task list
    _ = task::write_task_statistics(task_stats_file, &task_stats);
    _ = task::write_task_list(task_list_file, &task_list);
}
