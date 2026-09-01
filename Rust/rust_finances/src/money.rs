pub enum Currency {
    Usd,
    Eur,
    Yen,
}

pub struct Money {
    cents: i64,
    currency: Currency,
}

impl Money {
    pub fn new(starting_balance: i64, currency: Currency) -> Self {
        Self {
            cents: starting_balance,
            currency,
        }
    }

    pub fn cents(&self) -> i64 {
        self.cents
    }
}
