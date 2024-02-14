package main

import (
	"log"
	"sync"

	"github.com/ranjitmahadik/flash-sale/prepare"
	"github.com/ranjitmahadik/flash-sale/sale"
	"github.com/ranjitmahadik/flash-sale/storage"
)

func main() {
	_, err := storage.Init()
	if err != nil {
		log.Fatal("[main]", "Failed to init db")
		panic("db init failed")
	}
	err = prepare.CreateTable()
	if err != nil {
		log.Fatalf("[main]:", "Failed to create table with err %+v", err)
		panic(err)
	}

	// defer func() {
	// 	_ = prepare.DropTable()
	// }()

	// prepare.PopulateFakeData()	// run only once to populate fake data.

	wg := sync.WaitGroup{}
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func(userId int) {
			defer wg.Done()
			productBookingRequest := sale.ProductBookingRequest{ProductId: 1, UserId: userId}
			sale.BookProduct(productBookingRequest)
		}(i + 1)
	}

	wg.Wait()
}
