# My Go Web Application

This is a simple Go web application that serves editable content and allows updating it via HTTP requests. It is designed as a basic example and should not be used in a production environment without proper security measures.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Explanation](#Explanation)
- [Endpoints](#endpoints)
- [Security Considerations](#security-considerations)
- [License](#license)

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/Ebrahim-Ramadan/simple-editText-Golang.git
   ```


## Explanation
The React component (Tipta) is responsible for creating the user interface and allowing users to edit the content. It integrates the Tiptap rich text editor within your web application, providing a user-friendly way to edit content.

The Go main.go application, in this context, primarily serves as the backend server that handles network communication. It provides endpoints (e.g., /content and /update) that the React component can communicate with. When a user interacts with the React component (e.g., updates content), the React component sends HTTP requests to the Go server, and the server processes these requests and updates the content as needed.

This separation of concerns is a common architectural approach in web applications, where I have a frontend ([NextJs](https://nextjs.org/)) responsible for the user interface and user interactions, and a backend ([Go](https://go.dev/)) responsible for handling data and network communication.
but we can run the main.go by 
```
curl -X POST -d "content=<strong>New Content</strong>" http://localhost:8080/update
```
## Usage
This web application provides two endpoints:

```/content (GET) ```
Returns the current editable content as HTML.
```/update (POST)```
Allows updating the content by sending a POST request with the new content as a content parameter.
