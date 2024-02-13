package prepare

import (
	"log"
	"sync"

	"github.com/ranjitmahadik/flash-sale/storage"
)

func PopulateFakeData() error {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			tx, err := storage.DB.Begin()

			defer func() {
				if err != nil {
					err = tx.Rollback()
					if err != nil {
						log.Fatalf("[prepare]:", "couldn't rollback transaction", err)
					}
				}
			}()

			defer wg.Done()
			if err != nil {
				log.Fatalf("[prepare]:", "couldn't start transaction to insert fake data ", err)
				return
			}
			productId := 1
			_, err = tx.Exec(`INSERT INTO sales(product_id) values($1)`, productId)
			if err != nil {
				log.Fatalf("[prepare]:", "failed to insert row into db ", err)
				return
			}
			err = tx.Commit()
			if err != nil {
				log.Fatalf("[prepare]:", "failed to insert row into db ", err)
				return
			}
		}()
	}

	wg.Wait()

	log.Println("[prepare]:", "dummy data inserted into database.")

	return nil
}
