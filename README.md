# Hermes-API

Welcome to **Hermes-API**, the backend powering a restaurant platform. This API enables restaurants, delivery personnel, and customers to interact seamlessly through a robust and scalable system.

## Features

- 🍽️ **Restaurant Management** - Manage menus, orders, and operational details.
- 🚀 **Order Processing** - Handle real-time order tracking and fulfillment.
- 📦 **Delivery System** - Assign and track deliveries efficiently.
- 🔑 **Authentication & Authorization** - Secure access using JWT and OAuth.
- ⚡ **Scalable Architecture** - Built with Golang, PostgreSQL, Redis, and Docker.

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/doc/install) (latest stable version recommended)
- [Docker](https://www.docker.com/get-started) (for containerized deployment)
- [PostgreSQL](https://www.postgresql.org/download/) (if running the database locally)
- [Redis](https://redis.io/download) (for caching and queue management)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/hermes-api.git
   ```
2. Navigate to the project directory:
   ```sh
   cd hermes-api
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```

### Configuration

- Copy the `.env.example` file and rename it to `.env`:
  ```sh
  cp .env.example .env
  ```
- Update the `.env` file with your database credentials and API keys.

### Running the Application

#### Using Docker

```sh
docker-compose up --build
```

#### Running Without Docker

```sh
go run main.go
```

### Running Tests

Run tests to ensure everything is working correctly:
```sh
go test ./...
```

## Contributing

We welcome contributions! Please check our [Contributing Guide](./CONTRIBUTING.md) for details on how to contribute.

## License

Hermes-API is licensed under the MIT License.

---

🚀 **Built with love by the Hermes-API Team**
