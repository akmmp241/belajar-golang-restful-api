package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfErr(errorRollback)
		panic(err)
	}
	errorCommit := tx.Commit()
	PanicIfErr(errorCommit)
}
