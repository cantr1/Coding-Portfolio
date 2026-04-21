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

// Class
class User {
  constructor(name, age) {
    this.name = name;
    this.age = age;
  }
}

const user_object = new User("Kelz", 28);

//private fields
class Movie {
  #title;
  constructor(title, rating) {
    this.#title = title;
    this.rating = rating;
  }
}

// Static methods
// A static method or property is bound to the class itself, not the instance of the class (an object). 
class User {
  static numUsers = 0;

  constructor(name, age) {
    this.name = name;
    this.age = age;
    User.numUsers++;
  }

  static getNumUsers() {
    return User.numUsers;
  }
}

const lane = new User("Lane", 30);
console.log(User.getNumUsers()); // 1
const allan = new User("Allan", 30);
console.log(User.getNumUsers()); // 2

// This doesn't work because its not a method on the object
console.log(lane.getNumUsers());
// TypeError: lane.getNumUsers is not a function
//    at main.js:20:18

// Geter and Setter
class User {
  constructor(name, age) {
    this.name = name;
    this._age = age;
  }

  get age() {
    return this._age;
  }

  set age(value) {
    if (value < 0) {
      throw new Error("Age can't be negative.");
    }
    this._age = value;
  }
}

// str slicing
const text = "JavaScript";
console.log(text.slice(0, 4)); // Output: "Java"

// inheritance
class Titan {
  constructor(name) {
    this.name = name;
  }

  speak() {
    // this gets overridden in the BeastTitan class
    console.log("*titan noises*");
  }
}

class BeastTitan extends Titan {
  speak() {
    console.log(`${this.name} says, "I'm the Beast Titan"`);
  }
}

const pureTitan = new Titan("Eren's mom");
pureTitan.speak();
// *titan noises*

const beast = new BeastTitan("Zeke");
beast.speak();
// Zeke says, "I'm the Beast Titan"


// Super
class Titan {
  constructor(name) {
    this.name = name;
  }
  toString() {
    return `Titan - Name: ${this.name}`;
  }
}

class BeastTitan extends Titan {
  constructor(name, power) {
    // call the parent's constructor
    super(name);
    this.power = power;
  }

  toString() {
    // call the parent's `toString` method
    return `${super.toString()}, Power: ${this.power}`;
  }
}

/*
In JavaScript, for...in iterates over the indices (0, 1, 2...) rather than the values. 
For iterating over values, use for...of instead
*/

// Errors
const err = new Error("We've run out of baked salmon");
console.log(err.message);
// We've run out of baked salmon

// Try block
try {
  const titan = {};
  console.log(titan.neck.thickness);
  console.log("what's a titan?");
} catch (err) {
  console.log(err.message);
} finally {
  console.log("This will always run regardless of any errors.");
}

// Throw errors
const sendMessage = (msg) => {
  if (msg.length > 70) {
    throw new Error("Message is too long")
  } else {
    return msg
  }
};

// Sets
const set = new Set([1, 2, 3, 4, 5, 5, 5, 5]);
console.log(set);
// Set { 1, 2, 3, 4, 5 }
set.add(6);
console.log(set);
// Set { 1, 2, 3, 4, 5, 6 }
set.delete(3);
console.log(set);
// Set { 1, 2, 4, 5, 6 }
console.log(set.has(4));
// true
console.log(set.size);
// 5

// Maps
const map = new Map();
map.set("bertholdt", "shifter");
map.set("reiner", "warrior");
map.set("annie", "shifter");
map.set("bertholdt", "colossal titan");
console.log(map);
// Map { 'bertholdt' => 'colossal titan', 'reiner' => 'warrior', 'annie' => 'shifter' }

map.delete("annie");
console.log(map);
// Map { 'bertholdt' => 'colossal titan', 'reiner' => 'warrior' }

// Asynchronous programming
console.log("I print first");
setTimeout(
  () => console.log("I print third because I'm waiting 100 milliseconds"),
  100,
);
console.log("I print second");

const promise = new Promise((resolve, reject) => {
  setTimeout(() => {
    if (getRandomBool()) {
      resolve("resolved!");
    } else {
      reject("rejected!");
    }
  }, 1000);
});

function getRandomBool() {
  return Math.random() < 0.5;
}

function getPromiseForUserData() {
  return new Promise((resolve) => {
    fetchDataFromServer().then(function (user) {
      resolve(user);
    });
  });
}
 // async/await syntax
const promise = getPromiseForUserData();
async function getPromiseForUserData() {
  const user = await fetchDataFromServer();
  return user;
}

const promise = getPromiseForUserData();

function sleep(ms) {
  const promise = new Promise((resolve) => {
    setTimeout(() => {
        resolve();
    }, ms);
  });
  return promise;
};


// microtasks and macrotasks
function main() {
  console.log("main start");

  setTimeout(() => {
    console.log("macrotask 1 finished");
  }, 0);

  Promise.resolve()
    .then(() => {
      console.log("microtask 1 finished");
    })
    .then(() => {
      console.log("microtask 2 finished");
    });

  console.log("main end");
}

main();
// Output:
/*
main start
main end
microtask 1 finished
microtask 2 finished
macrotask 1 finished
*/

// module exports
function moo(name) {
    return `moo ${name}!`;
}

module.exports = {
    moo,
};

// Import
const moo = require("./moo.js");

console.log(moo)

