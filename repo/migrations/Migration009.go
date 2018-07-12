package migrations

import (
	"database/sql"

	_ "github.com/mutecomm/go-sqlcipher"
)

type Migration009 struct{}

func (Migration009) Up(repoPath string, dbPassword string, testnet bool) (err error) {
	db, err := newDB(repoPath, dbPassword, testnet)
	if err != nil {
		return err
	}

	err = withTransaction(db, func(tx *sql.Tx) error {
		for _, stmt := range []string{
			"ALTER TABLE cases ADD COLUMN coinType text;",
			"ALTER TABLE sales ADD COLUMN coinType text;",
			"ALTER TABLE purchases ADD COLUMN coinType text;",
			"ALTER TABLE cases ADD COLUMN paymentCoin text;",
			"ALTER TABLE sales ADD COLUMN paymentCoin text;",
			"ALTER TABLE purchases ADD COLUMN paymentCoin text;",
		} {
			_, err := tx.Exec(stmt)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = writeRepoVer(repoPath, 10)
	if err != nil {
		return err
	}
	return nil
}

func (Migration009) Down(repoPath string, dbPassword string, testnet bool) error {
	db, err := newDB(repoPath, dbPassword, testnet)
	if err != nil {
		return err
	}

	err = withTransaction(db, func(tx *sql.Tx) error {
		for _, stmt := range []string{
			"ALTER TABLE cases DROP COLUMN coinType text;",
			"ALTER TABLE sales DROP COLUMN coinType text;",
			"ALTER TABLE purchases DROP COLUMN coinType text;",
			"ALTER TABLE cases DROP COLUMN paymentCoin text;",
			"ALTER TABLE sales DROP COLUMN paymentCoin text;",
			"ALTER TABLE purchases DROP COLUMN paymentCoin text;",
		} {
			_, err := tx.Exec(stmt)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = writeRepoVer(repoPath, 9)
	if err != nil {
		return err
	}
	return nil
}
