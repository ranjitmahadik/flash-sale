package prepare

import (
	"log"

	"github.com/ranjitmahadik/flash-sale/storage"
)

func CreateTable() error {
	tx, err := storage.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Printf("[prepare]: Rollback failed: %v", rollbackErr)
			}
		}
	}()

	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS sales (
			sale_id SERIAL PRIMARY KEY,
			product_id integer not null,
			user_id integer,
			sold_at datetime,
			locked_at bigint
		)
	`)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	log.Println("[prepare]:", "table created successfully.")
	return nil
}

func DropTable() error {
	tx, err := storage.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		DROP TABLE IF EXISTS sales
	`)

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Printf("[prepare]: Rollback failed: %v", rollbackErr)
			}
		}
	}()

	err = tx.Commit()

	if err != nil {
		return err
	}
	log.Println("[prepare]:", "table dropped successfully.")
	return nil
}
