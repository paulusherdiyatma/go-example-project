# Go Example Project

This is a Go example project that follows a layered architecture.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Running with Docker](#running-with-docker)
- [Architecture](#architecture)
- [Dependencies](#dependencies)
- [Others](#others)


## Contributing
Contributions to this project are welcome. To contribute, follow these steps:
1. Fork the repository.
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes and commit them: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin feature/your-feature-name`
5. Submit a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project is an example implementation of a Go application that follows a layered architecture. It demonstrates how to separate concerns and organize code into different layers, such as the presentation layer, business logic layer, and data access layer.

## Installation
To install and run this project, follow these steps:
1. Clone the repository: `git clone https://github.com/your-username/go-example-project.git`
2. Change into the project directory: `cd go-example-project`
3. Create an environment variable file: `cp .env.example .env`
4. Update the `.env` file with your desired configuration.
5. Install the dependencies: `go mod download`
6. Run the program: `go run cmd/example-app/main.go`

## Running with Docker
To run the program using Docker, follow these steps:
1. Make sure you have Docker installed and running on your system. You can download Docker from the official website: [Docker](https://www.docker.com/).
2. Build the Docker image by running the following command in the project directory:
    ```
    docker build -t go-example-project .
    ```
3. Once the image is built, you can run the program in a Docker container using the following command:
    ```
    docker run -p 8080:8080 go-example-project
    ```
    This command maps port 8080 of the container to port 8080 of your local machine.
4. Open your web browser and navigate to `http://localhost:8080` to access the application.

Running the program via Docker allows for easy deployment and portability of the application, as it encapsulates all the dependencies and configurations within a container.

## Architecture
The architecture of this Go example project follows a layered approach, which helps in separating concerns and organizing code into different layers. The layers in this project include:

1. Presentation Layer: This layer is responsible for handling user interactions and displaying information to the user. It includes components such as handlers, controllers, and views.

2. Business Logic Layer: This layer contains the core logic of the application. It implements the business rules and processes data received from the presentation layer. It includes components such as services, managers, and validators.

3. Data Access Layer: This layer is responsible for interacting with the data storage and retrieval mechanisms. It includes components such as repositories, data access objects (DAOs), and database connectors.

The layered architecture promotes separation of concerns, modularity, and maintainability of the codebase. It allows for easier testing, reusability, and scalability of the application.

## Dependencies
To build and run this project, you will need the following dependencies:
- **Database**: PostgreSQL
- **API Library**: Fiber
- **ORM**: GORM

### Database
This project uses PostgreSQL as the database. Make sure you have PostgreSQL installed and running on your system. You can download PostgreSQL from the official website: [PostgreSQL](https://www.postgresql.org/).

### API Library
This project uses Fiber as the API library. Fiber is a web framework for Go that is designed to be fast and efficient. You can find more information about Fiber and how to install it on the official website: [Fiber](https://gofiber.io/).

### ORM
This project uses GORM as the ORM (Object-Relational Mapping) library. GORM is a powerful and flexible ORM for Go that provides a convenient way to interact with databases. You can find more information about GORM and how to install it on the official website: [GORM](https://gorm.io/).

## Others
### linter
Install and run golangcli-lint as a Go linter aggregator and static code analysis tool. Execute this command to run:

```
golangci-lint run 
```

### Formmater
Install and run Gofumpt as a formatting tool for Go code that enforces a specific style convention, automatically adjusting code to adhere to its rules, which are designed to enhance readability and consistency across projects.
```
gofumpt -l -w .
```