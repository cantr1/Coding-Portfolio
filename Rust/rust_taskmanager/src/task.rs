use serde::{Deserialize, Serialize};
use std::fs;
use std::path::Path;

// --- Struct definitions
#[derive(Serialize, Deserialize, Debug)]
pub struct Task {
    pub id: i32,
    pub description: String,
    pub complete: bool,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct TaskStatistics {
    pub next_task: i32,
    pub completed_tasks: i32,
}

pub fn increment_next_task(task_statistics: &mut TaskStatistics) {
    task_statistics.next_task += 1;
}

pub fn build_task(task_id: i32, task_description: &str, task_complete: bool) -> Task {
    Task {
        id: task_id,
        description: task_description.to_string(),
        complete: task_complete,
    }
}

pub fn build_task_statistics(file_path: &Path) -> TaskStatistics {
    println!(
        "checking filepath '{:#?}' for existing statistics",
        file_path
    );
    if file_path.is_file() {
        println!("found task statistics file, opening to parse json");
        // Open and parse file
        let result = std::fs::read_to_string(file_path).unwrap();

        // parse string into JSON
        let stats: TaskStatistics = serde_json::from_str(&result).unwrap();

        println!("found the following statistics: {:#?}", stats);
        stats
    } else {
        println!("file path not found, creating new statistics");
        TaskStatistics {
            next_task: 1,
            completed_tasks: 0,
        }
    }
}

pub fn build_task_list(file_path: &Path) -> Vec<Task> {
    println!(
        "checking filepath '{:#?}' for existing task list",
        file_path
    );
    if file_path.is_file() {
        println!("found task list, parsing to vector");
        let result = std::fs::read_to_string(file_path).unwrap();
        let task_list: Vec<Task> = serde_json::from_str(&result).unwrap();
        println!("found the following task list:\n{:#?}", task_list);

        task_list
    } else {
        println!("file path not found, creating new task list");
        let task_list: Vec<Task> = Vec::new();
        task_list
    }
}

pub fn write_task_statistics(file_path: &Path, stats: &TaskStatistics) -> std::io::Result<()> {
    // Write stats to file, return true if successful
    let json = serde_json::to_string_pretty(&stats).unwrap();

    // Accepts a path and string/byte data
    fs::write(file_path, json)?;
    Ok(())
}

pub fn write_task_list(file_path: &Path, tasks: &Vec<Task>) -> std::io::Result<()> {
    let json = serde_json::to_string_pretty(&tasks).unwrap();
    fs::write(file_path, json)?;
    Ok(())
}
