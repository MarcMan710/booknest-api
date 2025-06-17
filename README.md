# BookNest API Documentation

## 📌 Project Title:

**BookNest API**

## 🎯 Objective:

Develop a robust RESTful API in Go (Golang) to manage an online bookstore's inventory, customer orders, and user authentication with role-based access control.

---

## 🚀 Getting Started

### Prerequisites

Before running the BookNest API, ensure you have the following installed:

*   **Go (Golang)**: Version 1.22 or higher. [Download Go](https://golang.org/dl/)
*   **PostgreSQL**: Version 13 or higher. [Download PostgreSQL](https://www.postgresql.org/download/)

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/booknest-api.git
    cd booknest-api
    ```

2.  **Set up environment variables:**
    Create a `.env` file in the root directory and add the following variables:
    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_user
    DB_PASSWORD=your_password
    DB_NAME=booknest
    JWT_SECRET=your_jwt_secret_key
    ```
    *Replace `your_user`, `your_password`, and `your_jwt_secret_key` with your actual credentials.*

3.  **Run database migrations:**
    ```bash
    # Assuming you have a migration tool, e.g., Goose or Migrate
    # Example with Goose:
    # go run github.com/pressly/goose/v3/cmd/goose UP
    ```
    *(Please provide the actual migration command once the migration tool is decided.)*

4.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

5.  **Run the application:**
    ```bash
    go run main.go
    ```
    The API will be accessible at `http://localhost:8080` (or your configured port).

---

## 🛠️ Technical Stack

*   **Language**: Golang
*   **Database**: PostgreSQL
*   **Authentication**: JWT (JSON Web Tokens) with Bcrypt for password hashing

---

## ✨ Features

### 🔐 Authentication

*   **User Registration**: Create new user accounts.
*   **User Login**: Authenticate users and generate JWT tokens.
*   **Role-Based Access Control**: Differentiate between `Admin` and `Customer` roles.
*   **Profile Management**: Users can retrieve their profile information.

### 📚 Book Management (Admin Only)

Administrators have full control over the book inventory:

*   **Add Book**: Add new book entries with details such as `Title`, `Author`, `Price`, `Stock`, `ISBN`, `Category`, and `PublishedDate`.
*   **Update Book**: Modify existing book details.
*   **Delete Book**: Remove books from the inventory.
*   **List Books**: Retrieve a list of all available books.
*   **Search Books**: Search for books by `Title`, `Author`, or `Category`.

### 🛒 Order System (Customer Only)

Customers can manage their orders:

*   **Place Order**: Select books and specify quantities to place a new order.
*   **View Order History**: Access a list of all previous orders.
*   **Cancel Pending Order**: Cancel orders that have not yet been processed.

### 📈 Admin Dashboard (API Endpoints)

Admin-specific endpoints to monitor business metrics:

*   **View Total Orders**: Get the total number of orders.
*   **View Revenue**: See the total revenue generated.
*   **View Low-Stock Books**: Identify books with low stock levels.

---

## 🚀 API Endpoints

*(Detailed API endpoint documentation with request/response examples and authentication requirements will be provided here, possibly generated via Swagger/OpenAPI.)*

---

## 🧪 Testing

*   **Unit Testing**: Comprehensive unit tests using Go's built-in `testing` package.
    ```bash
    go test ./...
    ```
*   **Postman Collection**: A Postman collection will be provided for easy testing of all API endpoints.

---

## 🤝 Contributing

We welcome contributions! Please see our `CONTRIBUTING.md` (if available) for guidelines on how to contribute to this project.

---

## 📄 License

This project is licensed under the MIT License - see the `LICENSE` file for details.

---

## 📞 Contact

For any inquiries, please contact [Your Name/Email/GitHub Profile].