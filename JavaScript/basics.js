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

// functions
function getSum(a, b) {
    return a + b;
}

// loop
for (let i = 0; i < 5; i++) {
  console.log("Iteration number: " + i);
}

// IIFE
const result = (function (a, b) {
  return a + b;
})(1, 2);

console.log(result);
// 3

/* Optional chaining
function getRegion(campaign) {
  return campaign.location?.region;
}
*/

// Object methods
const quest = {
  title: "Save the Kingdom",
  completed: false,
  completeQuest() {
    this.completed = true;
    return `You have completed the quest: ${this.title}`;
  },
};

// Spread syntax
const all_employees = { ...engineering_dept, ...video_dept };

// destructuring
const apple = {
  radius: 2,
  color: "red",
};

const { radius, color } = apple;

// bind method
const user = {
  firstName: "Lane",
  lastName: "Wagner",
  getFullName() {
    return `${this.firstName} ${this.lastName}`;
  },
};

function getGreeting(introduction, nameCallback) {
  return `${introduction}, ${nameCallback()}`;
}

const boundGetFullName = user.getFullName.bind(user);
console.log(getGreeting("Hello", boundGetFullName));