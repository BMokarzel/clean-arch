package database

import (
	"database/sql"

	"github.com/BMokarzel/clean-arch.git/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Migrate() error {
	query := ("CREATE TABLE IF NOT EXISTS orders (id VARCHAR(36) PRIMARY KEY, price DECIMAL(10,2) NOT NULL, tax DECIMAL(10,2) NOT NULL, final_price DECIMAL(10,2) NOT NULL);")

	_, err := r.Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) List() ([]entity.Order, error) {
	rows, err := r.Db.Query("SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []entity.Order

	for rows.Next() {
		var o entity.Order
		err := rows.Scan(&o.ID, &o.Price, &o.Tax, &o.FinalPrice)
		if err != nil {
			return nil, err
		}
		items = append(items, o)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
