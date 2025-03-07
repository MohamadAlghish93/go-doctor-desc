# Go Doctor Desc

## Description
Go Doctor Desc is a web application designed to manage and store medical information including users, receipts, medicines, roles, and user roles. It is built using the Go programming language and uses SQLite as the database.

## Features
- User management
- Receipt management
- Medicine management
- Role-based access control

## Installation
### Prerequisites
- Go 1.16 or later

### Steps
1. Clone the repository:
    ```sh
    git clone https://github.com/MohamadAlghish93/go-doctor-desc.git
    ```
2. Navigate to the project directory:
    ```sh
    cd go-doctor-desc
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage
1. Initialize the database and run the application:
    ```sh
    go run main.go
    ```
2. The application will run on `http://localhost:8080`.

## Project Structure
- `main.go`: The entry point of the application.
- `database/database.go`: Initializes the database and performs auto-migration.
- `routes/`: Contains the route definitions (unable to retrieve detailed information).

## Contributing
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add some feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

## License
This project is licensed under the MIT License.
