package mappers

import (
	"database/sql"
	"fmt"
	"managerApp/app/models/entities"
)

type AuthMapper struct {
	db *sql.DB
}

func (m *AuthMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}

func (m *AuthMapper) Login(userName string, password string) (*entities.User, error) {

	var (
		c_id        sql.NullInt64
		c_user_name sql.NullString
		c_password  sql.NullString
	)

	selectQuery := `SELECT u_id, u_login, u_password FROM t_user WHERE u_login = $1 AND
	u_password = $2`
	row := m.db.QueryRow(selectQuery, userName, password)
	err := row.Scan(&c_id, &c_user_name, &c_password)

	if err != nil {
		return nil, fmt.Errorf("Login:%v", err)
	}

	user := &entities.User{
		ID:       c_id.Int64,
		UserName: c_user_name.String,
		Password: c_password.String,
	}
	return user, nil
}

func (m *AuthMapper) Logout() (*entities.User, error) {
	return nil, nil
}
