package user

import (
	"database/sql"
	"fmt"
	"github.com/devjunaeid/ecom-userSide/types"
)

type Store struct {
	db *sql.DB
}

func CreateStore(database *sql.DB) *Store {
	return &Store{
		db: database,
	}
}

func (s *Store) FindUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email=?", email)
	if err != nil {
		return nil, err
	}

	user := new(types.User)

	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == "" {
		return nil, fmt.Errorf("User not Found!!")
	}
	return user, nil
}

func scanRowIntoUser(r *sql.Rows) (*types.User, error) {
	u := new(types.User)
	err := r.Scan(
		&u.ID,
		&u.Email,
		&u.Username,
		&u.Password,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}
