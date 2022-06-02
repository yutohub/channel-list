package repository

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/yutohub/channel-list/model"
)

func FindUserByToken(token string) (*model.User, error) {
	var user model.User
	row := db.QueryRow(
		`SELECT id,name FROM user JOIN user_session
			ON user.id = user_session.user_id
				WHERE user_session.token = ? && user_session.expires_at > ?
				LIMIT 1`, token, time.Now(),
	)
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return &user, fmt.Errorf("FindUserByToken %s: no such token", token)
		}
		return &user, fmt.Errorf("FindUserByToken %s: %v", token, err)
	}
	return &user, nil
}

func CreateNewToken(userID uint64, expiresAt time.Time) (string, error) {
	token := generateToken()
	now := time.Now()
	// CreateNewToken(userID, token, expiresAt)
	_, err := db.Exec(
		`INSERT INTO user_session
			(user_id, token, expires_at, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?)`,
		userID, token, expiresAt, now, now,
	)
	if err != nil {
		return "", err
	}
	return token, err
}

func generateToken() string {
	table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_@"
	l := len(table)
	ret := make([]byte, 128)
	src := make([]byte, 128)
	rand.Read(src)
	for i := 0; i < 128; i++ {
		ret[i] = table[int(src[i])%l]
	}
	return string(ret)
}

func CreateNewUser(name string, password string) error {
	// Generate an unused ID
	id, err := generateID()
	if err != nil {
		return err
	}
	// password -> passwordHash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	now := time.Now()
	_, err = db.Exec(
		`INSERT INTO user
			(id, name, password_hash, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?)`,
		id, name, passwordHash, now, now,
	)
	return err
}

func findPasswordHashByName(name string) (string, error) {
	var hash string
	row := db.QueryRow(
		`SELECT password_hash FROM user
			WHERE name = ? LIMIT 1`, name,
	)
	if err := row.Scan(&hash); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("FindPasswordHashByName %s: no such name", name)
		}
		return "", fmt.Errorf("FindPasswordHashByName %s: %v", name, err)
	}
	return hash, nil
}

func generateID() (uint64, error) {
	var id uint64
	row := db.QueryRow("SELECT UUID_SHORT()")
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		}
		return 0, err
	}
	return id, nil
}

func FindUserByName(name string) (*model.User, error) {
	var user model.User
	row := db.QueryRow("SELECT id,name FROM user WHERE name = ? LIMIT 1", name)
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("FindUserByName %s: no such user", name)
		}
		return nil, fmt.Errorf("FindUserByName %s: %v", name, err)
	}
	return &user, nil
}

func LoginUser(name string, password string) (bool, error) {
	passwordHash, err := findPasswordHashByName(name)
	if err != nil {
		return false, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
