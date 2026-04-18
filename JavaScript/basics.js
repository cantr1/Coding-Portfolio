// var is old way (function-scoped), let and const are block-scoped
const user_name = "Kelz";
let favorite_language = "JavaScript";

console.log(`Hello ${user_name}! \nYour favorite language is ${favorite_language}`);

// undefinded
let age;
console.log(age); // Output: undefined

// null
let city = null;
console.log(city); // Output: null

// === Strict equality operator checks for both value and type
console.log(5 === '5'); // Output: false
console.log(5 === 5);   // Output: true

// == Loose equality operator checks for value after type coercion
console.log(5 == '5');  // Output: true
console.log(5 == 5);    // Output: true

// Control flow
let score = 85;
if (score >= 90) {
    console.log("Pass")
} else if (score >= 80) {
    console.log("pass")
} else {
    console.log("fail")
}

// Switch
const os = "mac";
let creator;
switch (os) {
  case "linux":
    creator = "Linus Torvalds";
    break;
  case "windows":
    creator = "Bill Gates";
    break;
  case "mac":
    creator = "Steve";
    break;
  default:
    creator = "Unknown";
    break;
}

// ternary
const price = isMember ? "$2.00" : "$10.00";

// nullish coalescing operator
let myName = null;
console.log(myName ?? "Anonymous"); // "Anonymous"

myName = "kelz";
console.log(myName ?? "Anonymous"); // "kelz"