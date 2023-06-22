# Go Ecom

![ci-test](https://github.com/paulodhiambo/ecom/actions/workflows/ci.yml/badge.svg)
![badmath](https://img.shields.io/github/languages/top/paulodhiambo/ecom)


## Description

Go Ecom is a simple e-commerce project implemented in Go. The motivation behind building this project was to create a
basic e-commerce system to learn and practice Go programming concepts.

The project aims to solve the problem of building a functional e-commerce API using Go. It provides a foundation for
managing products, handling orders, and integrating with payment services.

Through this project, I learned how to structure and develop a REST API in Go, handle authentication and authorization,
implement CRUD operations for products and orders, integrate with third-party payment services, and handle error
responses.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Credits](#credits)
- [License](#license)
- [Badges](#badges)
- [Features](#features)
- [How to Contribute](#how-to-contribute)
- [Tests](#tests)

## Installation

To install and run the project locally, follow these steps:

1. Clone the repository: `git clone https://github.com/paulodhiambo/ecom.git`
2. Navigate to the project directory: `cd ecom`
3. Install dependencies: `go mod tidy`
4. Set up the database and configure the connection details in the configuration file.
5. Run the application: `go run main.go`

## Usage

To use the Go Ecom API, follow the API documentation and interact with the available endpoints using tools like cURL,
Postman, or any other HTTP client.

For example, to retrieve a list of products, send a GET request to `/products`:

```
GET /products
```

To add a product to the shopping cart, send a POST request to `/cart` with the product ID and quantity:

```
POST /cart
{
  "product_id": "123",
  "quantity": 2
}
```

For more detailed usage instructions and examples, please refer to the API documentation.

## Credits

- [Odhiambo Paul](https://github.com/paulodhiambo) - Project Developer

## License

This project is licensed under the [MIT License](LICENSE).

## Features

- Product management (CRUD operations)
- Shopping cart functionality
- Order management
- Integration with payment services

## How to Contribute

Contributions to Go Ecom are welcome! If you would like to contribute, please follow the guidelines outlined
in [CONTRIBUTING.md](CONTRIBUTING.md).

## Migrate database

```shell
make migrateup
```

## Tests

To run the tests for this project, use the following command:

```shell
make test
```

> Make sure you have the necessary dependencies installed before running the tests.