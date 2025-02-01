# Contributing to Hermes-API

Welcome to Hermes-API! We appreciate your interest in contributing to our API, which powers a restaurant platform. By contributing, you help improve the platform for restaurants, delivery personnel, and customers.

## How to Contribute

### 1. Fork and Clone the Repository

- Fork the repository on GitHub.
- Clone your fork to your local machine:

  ```sh
  git clone https://github.com/your-username/hermes-api.git
  ```

- Navigate into the project directory:

  ```sh
  cd hermes-api
  ```

### 2. Set Up the Development Environment

#### Prerequisites

- Go (latest stable version recommended)
- Docker (for containerized development)
- PostgreSQL (if running the database locally)
- Redis (for caching and queue management)

#### Install Dependencies

Run the following command to install dependencies:

```sh
go mod tidy
```

#### Set Up Environment Variables

- Copy the `.env.example` file and rename it to `.env`:

  ```sh
  cp .env.example .env
  ```

- Update the `.env` file with your local database credentials and API keys.

#### Run the Project

- Start the development server with:

  ```sh
  docker-compose up --build
  ```

- If you’re running it without Docker, use:

  ```sh
  go run main.go
  ```

### 3. Branching and Code Contribution

- Create a new branch from `main`:

  ```sh
  git checkout -b feature/your-feature-name
  ```

- Make your changes and commit them following the conventional commit format:

  ```sh
  git commit -m "feat: add new endpoint for orders"
  ```

- Push your changes to your fork:

  ```sh
  git push origin feature/your-feature-name
  ```

- Open a pull request (PR) against the `main` branch.

### 4. Code Guidelines

- Follow the Golang Style Guide.
- Use Gin best practices for API routing.
- Ensure GORM models are properly updated before pushing changes.
- Write unit and integration tests using Testify.

### 5. Running Tests

Before submitting your PR, run tests to ensure nothing is broken:

```sh
go test ./...
```

### 6. Code Review Process

- Once you open a PR, the maintainers will review your code.
- If requested, make necessary changes.
- After approval, your PR will be merged.

## Community and Support

- Join our Discord Server (coming soon) for discussions.
- Open an Issue if you encounter bugs or need improvements.

Thank you for contributing to Hermes-API! 🚀
