package database

import (
	"Order-Matching-System/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	return sql.Open("mysql", dsn)
}

func CreateTables(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS orders (
            id INT AUTO_INCREMENT PRIMARY KEY,
            symbol VARCHAR(10) NOT NULL,
            side ENUM('buy', 'sell') NOT NULL,
            type ENUM('limit', 'market') NOT NULL,
            price DECIMAL(10,2),
            quantity INT NOT NULL,
            status ENUM('open', 'partially_filled', 'filled', 'cancelled') DEFAULT 'open',
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS trades (
            id INT AUTO_INCREMENT PRIMARY KEY,
            buy_order_id INT,
            sell_order_id INT,
            symbol VARCHAR(10) NOT NULL,
            price DECIMAL(10,2) NOT NULL,
            quantity INT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (buy_order_id) REFERENCES orders(id),
            FOREIGN KEY (sell_order_id) REFERENCES orders(id)
        )
    `)
	return err
}
