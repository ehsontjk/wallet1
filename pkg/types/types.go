
package types
type Money int64

type PaymentCategory string
type PaymentStatus string

const (
	PaymentStatusOk PaymentStatus = "OK"
	PaymentStatusFail PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

    

type Payment struct {
	ID string // 'card'
	AccountID int64
	Amount Money // номер вида '5058 xxxx xxxx 8888'
	Category PaymentCategory // баланс в дирамах
    Status PaymentStatus
}
type Phone string
type Account struct {
	ID  int64
	Phone Phone
	Balance Money
}