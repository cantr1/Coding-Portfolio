const numberButtons = document.querySelectorAll("[data-value]");
const operationButtons = document.querySelectorAll("[data-operation]");
const equalsButton = document.getElementById("equals-button");
const clearButton = document.getElementById("clear-button");
const expressionDisplay = document.getElementById("expression");
const resultDisplay = document.getElementById("result");

let firstValue = "";
let secondValue = "";
let selectedOperation = null;

numberButtons.forEach(function (button) {
    button.addEventListener("click", function () {
        addNumber(button.dataset.value);
        updateDisplay();
    });
});

operationButtons.forEach(function (button) {
    button.addEventListener("click", function () {
        chooseOperation(button.dataset.operation);
        updateDisplay();
    });
});

equalsButton.addEventListener("click", function () {
    calculateResult();
});

clearButton.addEventListener("click", function () {
    clearCalculator();
    updateDisplay();
});

function addNumber(number) {
    if (selectedOperation === null) {
        firstValue += number;
        return;
    }

    secondValue += number;
}

function chooseOperation(operation) {
    if (firstValue === "") {
        return;
    }

    selectedOperation = operation;
}

function calculateResult() {
    if (firstValue === "" || secondValue === "" || selectedOperation === null) {
        return;
    }

    const x = Number(firstValue);
    const y = Number(secondValue);
    const result = process_inputs(x, y, selectedOperation);

    resultDisplay.textContent = `= ${result}`;
}

function clearCalculator() {
    firstValue = "";
    secondValue = "";
    selectedOperation = null;
    resultDisplay.textContent = "Awaiting calculation...";
}

function updateDisplay() {
    const displayFirstValue = firstValue || "0";
    const displayOperation = selectedOperation || "";
    const displaySecondValue = secondValue || "";

    expressionDisplay.textContent = `> ${displayFirstValue} ${displayOperation} ${displaySecondValue}`;
}

function process_inputs(x, y, operation) {
    switch (operation) {
        case '+':
            return x + y;
        case '-':
            return x - y;
        case '/':
            if (y === 0) {
                return "Division by zero";
            }

            return x / y;
        case '*':
            return x * y;
    }
}
