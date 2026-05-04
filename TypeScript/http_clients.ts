export {};

// URL API is built into JS
// Create a URL object
const urlObj = new URL("https://homestarrunner.com/toons");

// Access properties
console.log(urlObj.hostname); // "homestarrunner.com"
console.log(urlObj.pathname); // "/toons"

// Headers API is also built into JS
function getContentType(resp: Response): string {
  const contentType = resp.headers.get("Content-Type");
  return contentType ? contentType : ""
  // Alternatively, using nullish coalescing operator:
    // return resp.headers.get("Content-Type") ?? "";
}

// JSON methods are also built into JS
const resp = await fetch("https://homestarrunner.com/toons");
const javascriptObjectResponse = await resp.json()

// Send a POST request with JSON body
const postResponse = await fetch("https://homestarrunner.com/api/data", {
  method: "POST",
  headers: {
    "Content-Type": "application/json"
  },
  body: JSON.stringify({ key: "value" })
});

// Fetch API is built into JS and can be used to make HTTP requests
// Can pass an options object to configure the request
const getResponse = await fetch("https://homestarrunner.com/toons", 
    {method: "GET", mode: "cors", headers: { "Accept": "application/json" }});
if (getResponse.ok) {
  const data = await getResponse.json();
  console.log(data);
} else {
  console.error("HTTP error:", getResponse.status);
}

// Zod is a popular library for schema validation and parsing in TypeScript
import { z } from "zod";

const UserSchema = z.object({
  id: z.number(),
  name: z.string(),
  email: z.string(),
});

// Schema reinforcement with Zod
const UserSchema = z.object({
  id: z.number().positive(), // must be positive
  name: z.string().min(1), // must be non-empty
  email: z.email(), // must be valid email string
});

// Parse with schemas
import { z } from "zod";

const UserSchema = z.object({
  id: z.number(),
  name: z.string(),
});

try {
  const user = UserSchema.parse(unknownData);
  // user is now typed and validated
  console.log(user.name); // TypeScript knows this is a string
} catch (error) {
  if (error instanceof z.ZodError) {
    console.error("Validation failed:", error.errors);
  }
}


// Instead of multiple definitions with validation logic, we can define a single schema and reuse it
const UserSchema = z.object({
  id: z.number(),
  name: z.string(),
});

type User = z.infer<typeof UserSchema>;