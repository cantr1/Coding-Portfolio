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
const headers = {method: "GET", mode: "cors", headers: { "Accept": "application/json" }};
const getResponse = await fetch("https://homestarrunner.com/toons", headers);
if (getResponse.ok) {
  const data = await getResponse.json();
  console.log(data);
} else {
  console.error("HTTP error:", getResponse.status);
}