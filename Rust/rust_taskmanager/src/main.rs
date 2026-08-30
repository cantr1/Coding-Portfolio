use std::io;
use std::io::Write;
use std::path::Path;
mod task;

fn take_user_input(prompt: &str) -> String {
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

    cleaned_input.to_string()
}

fn create_new_task(task_stats: &mut task::TaskStatistics, task_list: &mut Vec<task::Task>) {
    // Get description of task from user
    let task_description = take_user_input("Enter task description: ");

    // Create the task
    let new_task: task::Task = task::build_task(task_stats.next_task, &task_description, false);

    // Add the task to the list
    task_list.push(new_task);

    // Increment task statistics
    task::increment_next_task(task_stats);
}

fn view_task_list(task_list: &Vec<task::Task>) {
    println!();
    for task in task_list {
        println!(
            "[{}] ID: {} - {}",
            if task.complete { "*" } else { " " },
            task.id,
            task.description
        )
    }
    println!();
}

fn mark_task_complete(task_statistics: &mut task::TaskStatistics, task_list: &mut Vec<task::Task>) {
    let task_id_str = take_user_input("Enter task ID to complete: ");

    // convert to i32
    let task_id = match task_id_str.parse::<i32>() {
        Ok(id) => id,
        Err(_) => {
            println!("please enter a valid numeric task id");
            return;
        }
    };

    // parse task list, find matching id and mark complete
    for task in task_list.iter_mut() {
        if task.id == task_id {
            task::complete_task(task);
            task::increment_task_complete(task_statistics);
            println!("task marked complete");
            return;
        }
    }

    println!("no matching task id found for {task_id_str}");
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
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
            "Choices\n-------\nd - display tasks\nn - create new task\nc - complete task\nq - quit\n~:",
        );
        match user_choice.as_str() {
            "d" => view_task_list(&task_list),
            "n" => create_new_task(&mut task_stats, &mut task_list),
            "c" => mark_task_complete(&mut task_stats, &mut task_list),
            "q" => break 'main_loop,
            _ => println!("unrecognized input"),
        }
    }

    // End of operations, write statistics and task list
    task::write_task_statistics(task_stats_file, &task_stats)?;
    task::write_task_list(task_list_file, &task_list)?;

    Ok(())
}
