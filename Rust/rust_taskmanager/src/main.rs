use std::path::Path;
use std::io::Write;
use std::io;
mod task;

fn prompt_user_new_task(task_id: i32) -> task::Task {
    print!("Enter task description: ");
    io::stdout().flush().expect("failed to flush stdout");

    let mut user_input = String::new();

    io::stdin()
    .read_line(&mut user_input)
    .expect("failed to read line");

    let cleaned_input = user_input.trim();

    task::build_task(task_id, cleaned_input, false)
}

fn main() {
    // Initialize task statistics and task list
    let task_stats_file = Path::new("/home/kelz/Work/Coding-Portfolio/Rust/rust_taskmanager/artifacts/task_statistics.json");
    let mut task_stats = task::build_task_statistics(task_stats_file);

    //let task_list: Vec<task::Task> = Vec::new();
    let task_list_file = Path::new("/home/kelz/Work/Coding-Portfolio/Rust/rust_taskmanager/artifacts/task_list.json");
    let mut task_list = task::build_task_list(task_list_file);

    // Setup a new task
    let new_task = prompt_user_new_task(task_stats.next_task);

    // Add new task to list
    task::add_new_task(&mut task_stats, &mut task_list, new_task);

    // End of operations, write statistics and task list
    _ = task::write_task_statistics(task_stats_file, &task_stats);
    _ = task::write_task_list(task_list_file, &task_list);
}
