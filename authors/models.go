package authors

type Author struct {
	ID             uint   `db:"id"`
	FIO            string `db:"fio"`
	Pseudonym      string `db:"pseudonym"`
	Specialization string `db:"specialization"`
}
