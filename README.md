# Order Matching Engine

A simple Go service that provides an API to place and manage stock orders. The project uses Go, Gorilla Mux, and a local SQLite database (via GORM).

## Prerequisites

- Go 1.20+ installed → [Download Go](https://golang.org/dl/)
- Homebrew (for macOS) → [Install Homebrew](https://brew.sh/)
- SQLite (installed via brew):
  ```bash
  brew install sqlite
  ```
- Postman or curl (for testing)

## Setup Instructions

1. **Clone the Repository**
   ```bash
   https://github.com/iYashMaurya/GOLANG-ORDER-MATCHING-SYSTEM.git
   cd GOLANG-ORDER-MATCHING-SYSTEM
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Run Database Migrations**

   The app uses SQLite by default. The DB file will be auto-created. You don’t need to run migrations manually — just start the app.

4. **Run the Server**
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

## Running Tests (Optional)

```bash
go test ./...
```

## Project Structure

```
.
├── main.go
├── db/
│   └── db.go
├── handlers/
│   └── order_handlers.go
├── models/
│   └── order.go
├── routes/
│   └── routes.go
├── go.mod
├── go.sum
└── README.md
```
