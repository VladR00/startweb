package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DBpath = "./sql.db"
)

type User struct {
	Login    string
	Password string
	Time     int64
}

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DBpath)
	if err != nil {
		return nil, fmt.Errorf("Can't open DB: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Can't ping DB: %w", err)
	}
	//fmt.Println("DB successfully connecting")
	return db, nil
}

func InitiateTables() error {
	if err := CreateTable("users"); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
func CreateTable(table string) error {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()
	var q string
	switch table {
	case "users":
		q = `CREATE TABLE IF NOT EXISTS users (
			login TEXT PRIMARY KEY,
			password TEXT,
			time INTEGER)`
	}

	query, err := db.Prepare(q)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Can't preparing query for creating table %s: %w", table, err)
	}
	defer query.Close()

	if _, err = query.Exec(); err != nil {
		fmt.Println(err)
		return fmt.Errorf("Can't execute create table %s: %w", table, err)
	}

	return nil
}

func IsTableExists(db *sql.DB, tableName string) bool {
	query := `SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?;`
	var count int
	err := db.QueryRow(query, tableName).Scan(&count)

	if err != nil {
		fmt.Println(err)
	}

	return count > 0
}

func (s *User) InsertNew() error {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query, err := db.Prepare("INSERT INTO users (login, password, time) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Can't preparing query for insert new users into users: %w", err)
	}
	defer query.Close()

	if _, err = query.Exec(s.Login, s.Password, s.Time); err != nil {
		fmt.Println(err)
		return fmt.Errorf("Can't execute inserting new users into users: %w", err)
	}
	return nil
}

func OutputAllUsers() ([]*User, error) {
	db, err := OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := ("SELECT * FROM users")

	var userlist []*User

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.Login, &user.Password, &user.Time)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Staff not found while reads staff:", err)
			}
			fmt.Println("Undefined error while reads staff:", err)
		}
		userlist = append(userlist, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return userlist, nil //staff.Time = time.Unix(registrationTime, 0).Format("2006-01-02 15:04")
}

func SaveToSQL(login string, pass string) error {
	user := User{
		Login:    login,
		Password: pass,
		Time:     time.Now().Unix(),
	}
	if err := user.InsertNew(); err != nil {
		return err
	}
	fmt.Println(user)
	return nil
}

func ReadUser(login string, pass string) (*User, error) {
	db, err := OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := (`SELECT * FROM users 
				WHERE login = ? AND password = ?`)

	user := &User{}

	row := db.QueryRow(query, login, pass)
	err = row.Scan(&user.Login, &user.Password, &user.Time)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User not found while reads user: %w", err)
		}
		return nil, fmt.Errorf("Undefined error while reads user: %w", err)
	}
	return user, nil
}

func Login(login string, pass string) (*User, error) {
	user, err := ReadUser(login, pass)
	if err != nil {
		return nil, err
	}
	return user, nil
}
