package sale

import (
	"database/sql"
	"log"

	"github.com/ranjitmahadik/flash-sale/storage"
)

type ProductBookingRequest struct {
	ProductId int
	UserId    int
}

func BookProduct(bookingRequest ProductBookingRequest) error {
	tx, err := storage.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p) // Re-throw panic after Rollback
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var sale Sale
	err = tx.QueryRow(`
        SELECT sale_id, product_id, user_id, sold_at, locked_at 
        FROM sales 
        WHERE product_id = $1 
        AND sold_at IS NULL 
        AND (locked_at + $2 < extract(epoch from now()) OR locked_at IS NULL) 
        LIMIT 1 
        FOR UPDATE
    `, bookingRequest.ProductId, 10).Scan(&sale.SaleId, &sale.ProductId, &sale.UserId, &sale.SoldAt, &sale.LockedAt)
	if err == sql.ErrNoRows {
		return err
	} else if err != nil {
		return err
	}

	_, err = tx.Exec(`
        UPDATE sales 
        SET user_id = $1, locked_at = extract(epoch from now()) 
        WHERE sale_id = $2
    `, bookingRequest.UserId, *sale.SaleId)
	if err != nil {
		return err
	}

	sale.UserId = &bookingRequest.UserId // Update the sale's UserID
	log.Printf("[Booking]: Booked: ProductID: %d by UserID: %d\n", *sale.ProductId, *sale.UserId)
	return nil
}
