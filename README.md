# GitHub Activity CLI

A simple command line interface (CLI) application written in Go that fetches and displays the recent activity of a specified GitHub user. This project helps you practice working with APIs, handling JSON data, and building a CLI application.

## Features

- Fetches recent activity of a GitHub user using the GitHub API.
- Displays various types of activities such as pushes, issues opened, and stars.
- Handles errors gracefully, including invalid usernames and API failures.

## Requirements

- Go 1.16 or later

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Ssajaia/github-activity.git
   cd github-activity
Build the application (optional):
bash

Copy Code
go build -o main.go
Usage
Run the application from the command line, providing a GitHub username as an argument:

bash

Copy Code
go run main.go <username>
Replace <username> with the actual GitHub username you want to check.

Example
bash

Copy Code
go run main.go kamranahmedse
Output
The output will display the recent activities of the specified user, for example:

Code

Copy Code
Pushed 3 commits to kamranahmedse/developer-roadmap
Opened a new issue in kamranahmedse/developer-roadmap
Starred kamranahmedse/developer-roadmap
Error Handling
The application will handle errors such as:

Invalid GitHub usernames
API request failures
Network issues
Contributing
Contributions are welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Acknowledgments
GitHub API Documentation
https://roadmap.sh/projects/github-user-activity
