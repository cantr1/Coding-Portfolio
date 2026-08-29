use std::io;
use std::io::Write;
use std::path::Path;
mod task;

fn prompt_user_task_description() -> String {
    print!("Enter task description: ");
    io::stdout().flush().expect("failed to flush stdout");

    let mut user_input = String::new();

    io::stdin()
        .read_line(&mut user_input)
        .expect("failed to read line");

    let cleaned_input = user_input.trim();

    cleaned_input.to_string()
}

fn create_new_task(task_stats: &mut task::TaskStatistics, task_list: &mut Vec<task::Task>) {
    // Get description of task from user
    let task_description = prompt_user_task_description();

    // Create the task
    let new_task: task::Task = task::build_task(task_stats.next_task, &task_description, false);

    // Add the task to the list
    task_list.push(new_task);

    // Increment task statistics
    task::increment_next_task(task_stats);
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

    // Create a new task, add to list and increment stats
    create_new_task(&mut task_stats, &mut task_list);

    // End of operations, write statistics and task list
    _ = task::write_task_statistics(task_stats_file, &task_stats);
    _ = task::write_task_list(task_list_file, &task_list);
}
