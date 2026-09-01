use crate::money::{Currency, Money};

struct AccountID(i64);
struct TransactionID(i64);

enum TransactionKind {
    Income,
    Expense,
}

pub struct Transaction {
    id: TransactionID,
    account_id: AccountID,
    amount: Money,
    kind: TransactionKind,
}

impl Transaction {
    pub fn new() -> Self {
        Self {
            id: TransactionID(1),
            account_id: AccountID(1),
            amount: Money::new(100, Currency::Usd),
            kind: TransactionKind::Income,
        }
    }

    pub fn print_balance(&self) {
        println!("Current balance: {:.2}", self.amount.cents() as f64 / 100.0);
    }
}
