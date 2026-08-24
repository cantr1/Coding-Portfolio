pub struct Calculation {
    pub x: i32,
    pub y: i32,
    pub action: String,
    pub result: i32
}

pub fn build_calculation(input_x: i32, 
    input_y: i32, 
    input_action: String, 
    input_result: i32
) -> Calculation {
    Calculation {
        x: input_x,
        y: input_y,
        action: input_action,
        result: input_result,
    }
}