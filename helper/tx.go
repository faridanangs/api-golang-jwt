package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRolback := tx.Rollback()
		PanicError(errRolback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicError(errCommit)
	}
}
