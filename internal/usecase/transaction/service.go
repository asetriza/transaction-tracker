package transaction

type Service interface {
	CancelLatestOddRecords(limit int) error
}
