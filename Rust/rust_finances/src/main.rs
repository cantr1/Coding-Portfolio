mod money;
mod transaction;

fn main() {
    println!("Testing transaction creation");
    let test_t = transaction::Transaction::new();

    test_t.print_balance();
}
