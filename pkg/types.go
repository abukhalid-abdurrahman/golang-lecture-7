package types

type Money int64
type PAN string
type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type Card struct {
	ID int
	Balance Money
	PAN PAN
	Currency Currency
}