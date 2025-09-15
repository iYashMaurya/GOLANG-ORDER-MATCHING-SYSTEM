# Order Matching Engine

A simple Go service that provides an API to place and manage stock orders. The project uses Go, Gorilla Mux, and a local MySQL database (via GORM).

## Prerequisites

- Go 1.20+ installed → [Download Go](https://golang.org/dl/)
- Homebrew (for macOS) → [Install Homebrew](https://brew.sh/)
- MySQL:
  - **Linux (Ubuntu/Debian)**
    ```bash
    sudo apt update
    sudo apt install mysql-server -y
    sudo systemctl start mysql
    ```
  - **macOS (Homebrew)**
    ```bash
    brew install mysql
    brew services start mysql
    ```
  - **Windows**
    - Download installer: [MySQL Installer](https://dev.mysql.com/downloads/installer/)
    - Start MySQL from Services or MySQL Workbench.
- Postman or curl (for testing)

## Setup Instructions

1. **Clone the Repository**
   ```bash
   git clone https://github.com/<your-username>/<your-repo>.git
   cd <your-repo>
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Install MySQL**
   - **Linux (Ubuntu/Debian)**
     ```bash
     sudo apt update
     sudo apt install mysql-server -y
     sudo systemctl start mysql
     ```
   - **macOS (Homebrew)**
     ```bash
     brew install mysql
     brew services start mysql
     ```
   - **Windows**
     - Download installer: [MySQL Installer](https://dev.mysql.com/downloads/installer/)
     - Start MySQL from Services or MySQL Workbench.

4. **Database Setup**
   Login to MySQL:
   ```bash
   mysql -u root -p
   ```
   Create the database:
   ```sql
   CREATE DATABASE order_matching_system;
   ```
   (Optional) Verify:
   ```sql
   SHOW DATABASES;
   ```

5. **Environment Variables**
   Copy `.env.example` into `.env`:
   ```bash
   cp .env.example .env
   ```
   **.env.example**
   ```
   DB_USER=root
   DB_PASS=password
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=order_matching_system
   PORT=8080
   ```
   Edit `.env` to match your local MySQL credentials.

6. **Run the Server**
   ```bash
   go run main.go
   ```

   The server will start on port 8080.

## API Endpoints

1. **Create an Order**

   **POST** `/orders`

   **Request body:**
   ```json
   {
     "symbol": "AAPL",
     "side": "buy",
     "type": "limit",
     "price": 150.5,
     "quantity": 10
   }
   ```

   **Test with curl:**
   ```bash
   curl -X POST http://localhost:8080/orders \
     -H "Content-Type: application/json" \
     -d '{"symbol":"AAPL","side":"buy","type":"limit","price":150.5,"quantity":10}'
   ```

2. **Get Order by ID**

   **GET** `/orders/{id}`

   **Test with curl:**
   ```bash
   curl http://localhost:8080/orders/1
   ```

3. **Get Order Book**

   **GET** `/orderbook`

   **Test with curl:**
   ```bash
   curl http://localhost:8080/orderbook
   ```

4. **Get Trades**

   **GET** `/trades`

   **Test with curl:**
   ```bash
   curl http://localhost:8080/trades
   ```

## Project Structure

```
Order-Matching-System/
├── config
│   └── config.go
├── database
│   └── connection.go
├── go.mod
├── go.sum
├── handler
│   └── api.go
├── main.go
├── models
│   ├── order.go
│   └── trade.go
└── service
    ├── matching_engine.go
    ├── order-book.go
    └── order_heap.go
```
