package sale

import "time"

type Sale struct {
	SaleId    *int
	ProductId *int
	UserId    *int
	SoldAt    *time.Time
	LockedAt  *int64
}

func New(SaleId, ProductId, UserId int, SoldAt time.Time, LockedAt int64) *Sale {
	return &Sale{
		SaleId:    &SaleId,
		ProductId: &ProductId,
		UserId:    &UserId,
		SoldAt:    &SoldAt,
		LockedAt:  &LockedAt,
	}
}
