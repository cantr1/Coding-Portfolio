// Basic Types
const bootupMessage: string = "starting server...";
const port: number = 3000;
const isOnline: boolean = true;
const noValue: null = null;
const notDefined: undefined = undefined;

// Unions
function safeSquare(val: string | number): number {
  if (typeof val === "string") {
    val = parseInt(val, 10);
  }
  // now val is only a number
  return val * val;
}

let result = safeSquare("5");
console.log(result);
// 25

result = safeSquare(5);
console.log(result);
// 25

// Optional Parameters
function greet(name: string, title?: string): string {
  if (title) {
    return `Hello, ${title} ${name}!`;
  }
  return `Hello, ${name}!`;
}

greet("Gandalf");           // "Hello, Gandalf!"
greet("Gandalf", "Wizard"); // "Hello, Wizard Gandalf!"

// Literal Types
//function move(direction: "north") {
  // Implementation...
//}

type Direction = "north" | "south" | "east" | "west";

function move(direction: Direction) {
  // Implementation...
}

/* for simple return statements
if (level === "low") return 1;
*/

type Class = "wizard" | "warrior" | "rogue";
type Race = "elf" | "human" | "dwarf";
type Hero = `Hero: ${Race} ${Class}`;
// Hero: elf wizard | Hero: elf warrior and so on...

// easy iteration with .forEach()

// discriminated unions
type MultipleChoiceLesson = {
  kind: "multiple-choice"; // Discriminant property
  question: string;
  studentAnswer: string;
  correctAnswer: string;
};

type CodingLesson = {
  kind: "coding"; // Discriminant property
  studentCode: string;
  solutionCode: string;
};

type Lesson = MultipleChoiceLesson | CodingLesson;

function isCorrect(lesson: Lesson): boolean {
  switch (lesson.kind) {
    case "multiple-choice":
      return lesson.studentAnswer === lesson.correctAnswer;
    case "coding":
      return lesson.studentCode === lesson.solutionCode;
  }
}

// Sets
// A Set that contains only strings
const justiceLeague = new Set<string>();

justiceLeague.add("Green Arrow");
justiceLeague.add("Flash");

// Error: Argument of type '2' is not assignable to parameter of type 'string'
justiceLeague.add(2);

const justiceLeague = new Set<string>(["Atom", "Black Canary", "Blue Beetle"]);

console.log(justiceLeague.size); // 3

justiceLeague.delete("Blue Beetle");
console.log(justiceLeague.has("Blue Beetle")); // false

justiceLeague.forEach((member) => console.log(member));
// Atom
// Black Canary

// Basic Maps
// A Map with string keys and number values
const podracerSpeeds = new Map<string, number>();

podracerSpeeds.set("Anakin Skywalker", 947);
podracerSpeeds.set("Sebulba", 941);

podracerSpeeds.set("R2-D2", true);
// Error: Argument of type 'true' is not assignable to parameter of type 'number'

podracerSpeeds.set(420, 69);
// Error: Argument of type 'number' is not assignable to parameter of type 'string'

// iterate over map entries
for (const [racer, speed] of podracerSpeeds) {
  console.log(`${racer} raced at ${speed} speed`);
}
// Anakin raced at 947 speed
// Sebulba raced at 941 speed

// Dynamic keys
type UserMetrics = {
  [key: string]: number;
};

// says can have any number of properties, as long as the keys are strings and the values are numbers

type FormData = {
  [field: string]: string | number | boolean;
  email: string;
  password: string;
  age: number;
};

// this says email, password, and age are required properties, but you can also have any number of additional properties with string keys and values that are either string, number, or boolean.

// Readonly Properties
type Point = {
  readonly x: number;
  y: number;
};

// Satisfies
type ColorMap = {
  red: string | number;
  green: string | number;
  blue: string | number;
  yellow: string | number;
};

const colorsSatisfies = {
  red: "#ff0000",
  green: "#00ff00",
  blue: 255,
  yellow: "#ffff00",
  // Error: "yelow" is not in type ColorMap
  // yelow: "#ffff00"
} satisfies ColorMap;

// We keep the narrowed types!
type RedHexSatisfies = typeof colorsSatisfies.red;
const redUpper = colorsSatisfies.red.toUpperCase(); // "#FF0000"

// Function overloads
// note: function overloads need to be declared above the implementation
type Employee = {
    name: string;
    department: string;
};

function formatEmployeeMessage(employee: Employee): string;
function formatEmployeeMessage(
  employee: Employee,
  isNew: true,
  onBoardedDate: Date,
): string;

// this says if you call formatEmployeeMessage with an Employee, it will return a string. If you call it with an Employee, true, and a Date, it will also return a string. 
// But if you call it with an Employee and false, it will not match any overload and will give an error.

// simple tuple
// [string, number]
const nameAndAge: [string, number] = ["John Jones", 104];

// readonly tuple, the way they should be used since under the hood simply an array, and arrays are mutable by default. So we can use readonly to prevent accidental mutations.
const nameAndAge: readonly [string, number] = ["Martha Jones", 24];
// Error: Property 'push' does not exist on type 'readonly [string, number]'.
nameAndAge.push("Donna Noble");

// optional tuple elements
export type Ticket = readonly [id: number, comment: string, label?: string];

export function formatTicket(ticket: Ticket): string {
  const [id, comment, label] = ticket;
    return `#${id} ${comment}${label ? ` [${label}]` : ""}`;
}

// Type Intersection
type IndividualContributor = {
  id: number;
  name: string;
  tasks: string[];
};

type Manager = {
  directReports: number[];
};

type GoodManager = IndividualContributor & Manager;

const hunter: GoodManager = {
  id: 1,
  name: "Hunter Backmann",
  tasks: ["Fixing Lane's B*llsh*t code", "Vibe Coding"],
  directReports: [2, 3, 4],
};

type Point2D = {
  x: number;
  y: number;
};

type Point3D = Point2D & {
  z: number;
};

// Equivalent to:
// type Point3D = {
//   x: number;
//   y: number;
//   z: number;
// };

// Define a function in a type
export type TextBot = SupportBot & {
  messageLog: string[];
  sendMessage(message: string): string;
}

// Superset union
export type EmploymentStatus =
  | "employed"
  | "unemployed"
  | "student" 
  | (string & {});

// don't touch below this line

export function updateEmploymentStatus(status: EmploymentStatus) {
  return `Employment status updated: ${status}`;
}

// Interfaces
interface Character {
  name: string;
  level: number;
}

interface Wizard extends Character {
  spellbook: string[];
  mana: number;
}

// multiple interface extension
type Character = {
  name: string;
  level: number;
};

interface Magical {
  mana: number;
  castSpell(spell: string): void;
}

interface Physical {
  strength: number;
  attack(): void;
}

interface BattleMage extends Character, Magical, Physical {
  combineAttacks(): void;
}

// you can override types but they must match the original type
interface Character {
  rank: string | number;
  name: string;
  level: number;
}

interface Wizard extends Character {
  // Wizards only have a number rank
  // This is allowed because
  // `number` is assignable to `string | number`
  rank: number;
  mana: number;
}

// Enums
enum Direction {
  North, // 0
  East, // 1
  South, // 2
  West, // 3
}

let myDirection: Direction = Direction.North;
console.log(myDirection); // Outputs: 0

// Example
export type SupportRequest = {
  id: string;
  severity: RequestSeverity;
  description: string;
};

// don't touch above this line

export enum RequestSeverity {
  Low,
  Medium,
  High,
  Critical,
}

export function isCritical(request: SupportRequest) {
  return (request.severity === RequestSeverity.Critical);
}



// In operator
type TextMessage = {
  content: string;
  sentAt: Date;
};

type ImageMessage = {
  caption: string;
  sentAt: Date;
};

type VideoMessage = {
  duration: number;
  sentAt: Date;
};

type Message = TextMessage | ImageMessage | VideoMessage;

function displayMessage(message: Message) {
  if ("content" in message) {
    // TypeScript knows this is a TextMessage
    // because it's the only one with a 'content' property
    console.log(`Text content is: ${message.content}`);
  } else if ("caption" in message) {
    // TypeScript knows this is an ImageMessage
    // because it's the only one with an 'caption' property
    console.log(`Image caption is ${message.caption}`);
  } else {
    // TypeScript knows this is a VideoMessage because
    // it's the only other option
    console.log(`Video length is ${message.duration}`);
  }
}