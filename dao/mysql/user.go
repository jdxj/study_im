package mysql

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int64
	Nickname string
	Password string
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Create() error {
	query := fmt.Sprintf(`insert into %s (nickname,password) values(?,?)`,
		u.TableName())
	result, err := db.Exec(query, u.Nickname, u.Password)
	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()
	return err
}

func (u *User) Get() (bool, error) {
	query := fmt.Sprintf(`select nickname from %s
where id=?`, u.TableName())
	row := db.QueryRowx(query, u.ID)
	err := row.Scan(&u.Nickname)
	if err == nil {
		return true, nil
	}

	if err == sql.ErrNoRows {
		return false, nil
	}
	return false, err
}
