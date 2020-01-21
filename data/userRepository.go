package data

import (
	"database/sql"
	"fmt"
	"microAuth/model"

	_ "github.com/go-sql-driver/mysql"
)

type SqlUserRepository struct {
	Database *sql.DB
}

func (repo *SqlUserRepository) GetUserById(id int32) (user model.User, err error) {
	row := repo.Database.QueryRow("select userId, userName, passwordHash from users where userId = ?", id)
	err = row.Scan(&user.UserID, &user.UserName, &user.PasswordHash)
	return
}

func (repo *SqlUserRepository) GetUserByName(userName string) (user model.User, err error) {
	row := repo.Database.QueryRow("select userId, userName, passwordHash from users where userName = ?", userName)
	err = row.Scan(&user.UserID, &user.UserName, &user.PasswordHash)
	return
}

func (repo *SqlUserRepository) Create(user *model.User) (err error) {
	tx, err := repo.Database.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("insert into users (userName, passwordHash) values(?, ?)")
	if err != nil {
		return
	}
	res, err := stmt.Exec(user.UserName, user.PasswordHash)
	if err != nil {
		return
	}
	id, err := checkRowCount(res)
	if err != nil {
		return
	}

	user.UserID = int32(id)

	if err = tx.Commit(); err != nil {
		return
	}
	stmt.Close()
	return
}

func (repo *SqlUserRepository) Update(user *model.User) (err error) {
	res, err := repo.Database.Exec("update users SET userName = ?, passwordHash = ? where userId = ?",
		user.UserName, user.PasswordHash, user.UserID)
	if err != nil {
		return
	}
	id, err := checkRowCount(res)
	if id != int64(user.UserID) {
		fmt.Errorf("Expected to update user %d, updated %d.\n", user.UserID, id)
	}
	return
}

func (repo *SqlUserRepository) DeleteById(id int32) (err error) {
	res, err := repo.Database.Exec("delete from users where userId = ?", id)
	if err == nil {
		_, err = checkRowCount(res)
	}
	return
}

func (repo *SqlUserRepository) DeleteByName(userName string) (err error) {
	res, err := repo.Database.Exec("delete from users where userName = ?", userName)
	if err == nil {
		_, err = checkRowCount(res)
	}
	return
}

func checkRowCount(result sql.Result) (lastIndex int64, err error) {
	if lastIndex, err = result.LastInsertId(); err != nil {
		return
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected != 1 {
		err = fmt.Errorf("Expected to affect 1 row, affected %d.\n", rowsAffected)
	}
	return
}
