# Go Snapshot Server

This repository hosts the "Go Snapshot Server", a project utilizing the Domain-Driven Design (DDD) pattern for its development. It's structured to effectively manage domain complexity and facilitate maintainability.

## Prerequisites

Before starting, ensure you have Go installed on your machine. You can verify this by running `go version` in your terminal.

## Installation and Execution

Follow these steps to set up and run the project:

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/arturdolzan/go-snapshot-server.git
    ```

2. Navigate to the project directory:

    ```bash
    cd go-snapshot-server
    ```

3. Synchronize the project's dependencies:

    ```bash
    go mod tidy
    ```

4. Execute the application:

    ```bash
    go run .
    ```

## Layered Architecture

The project is structured into the following layers, adhering to the DDD principles:

- **Presentation**: Handles all the UI and API presentation logic.
- **Application**: Contains the application's use cases and orchestration logic.
- **Domain**: Includes the core business logic and entities.
- **Infrastructure**: Comprises the implementations of external interfaces like databases, file systems, etc.

## REST API Implementation with Gin Gonic

This project is implemented as a REST API using the Gin Gonic framework, which provides a robust set of tools for building efficient and high-performance web applications and services.

## Dependency Injection with Wire

This project leverages Wire for managing dependency injection, ensuring a cleaner and more maintainable configuration. For more information on how to use Wire, refer to the [Wire documentation](https://github.com/google/wire).
