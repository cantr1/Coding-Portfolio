use std::path::Path;
mod task;

fn main() {
    let task_stats_file = Path::new("/home/kelz/Workspace/Coding-Portfolio/Rust/rust_taskmanager/artifacts/task_statistics.json");

    let task_stats = task::build_task_statistics(task_stats_file);

    task::write_task_statistics(task_stats_file, &task_stats);

}
