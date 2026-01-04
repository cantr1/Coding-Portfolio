#!/usr/bin/python3

def convert_to_int(char: str) -> tuple[bool, int]:
    try:
        value = int(char)
        return True, value
    except ValueError:
        return False, 0
    
def determine_negative(char: str) -> bool:
    if char == "-":
        return True
    else:
        return False

def process_terms(polynomial: str) -> tuple[list[int], bool]:
    processed_terms = []

    split_polynomial = polynomial.replace('(', '').replace(')', '').split()

    first_term = split_polynomial[0]
    second_term = split_polynomial[2]
    negative = determine_negative(split_polynomial[1])

    preprocessed_terms = [first_term, second_term]

    for i, term in enumerate(preprocessed_terms):
        converted, num = convert_to_int(term)
        if converted:
            if i == 1:
                if negative:
                    processed_terms.append(num * -1)
                else:
                    processed_terms.append(num)
        else:
            processed_terms.append(term)

    return processed_terms, negative

def print_triangle(exponent: int) -> list[int]:
    row = []
    for i in range(exponent + 1):
        row = [1] + [row[i] + row[i+1] for i in range(len(row)-1)] + [1] if row else [1]
        if (i) == exponent:
            print("*" + str(row) + "*")
            return row
        
def factor_polynomial(processed_terms: list[str], coeffs: list[int]):
    a, b_str = processed_terms[0], processed_terms[1]
    try:
        b = int(b_str)
    except ValueError:
        return "Only numeric second term supported for now"

    terms = []
    n = len(coeffs) - 1
    for i, c in enumerate(coeffs):
        val = c * (b ** i)
        if val == 0: continue
        sign = "" if i == 0 else " + " if val > 0 else " - "
        coef = str(abs(val)) if abs(val) != 1 or n - i == 0 else ""
        x_part = "" if i == n else f"{a}" if i == n-1 else f"{a}^{n-i}"
        term = f"{sign}{coef}{x_part}"
        terms.append(term)
    return "".join(terms).replace(" + -", " - ")
    
def main() -> None:
    # Get info from user
    polynomial = input("Please enter your polynomial in this format (x + y): ")
    expo = int(input("Enter the power the polynomial is raised to: "))

    processed_terms, negative = process_terms(polynomial)

    row = print_triangle(expo)

    factored = factor_polynomial(processed_terms, row)
    print(f"{polynomial}^{expo} = {factored}")

if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\nBye!")