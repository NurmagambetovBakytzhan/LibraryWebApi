package members

type Member struct {
	ID  uint   `db:"id"`
	FIO string `db:"FIO"`
}
