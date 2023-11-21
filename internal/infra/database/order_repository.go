package database

import (
	"database/sql"

	"github.com/yagoinacio/golang-intro-fullcycle-1/internal/entities"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Create() error {
	_, err := r.Db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY(id))")

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) Save(order *entities.Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (id, price, tax, final_price) values (?, ?, ?, ?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetTotalTransactions() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
