package providers

import (
	"database/sql"
	"fmt"
	"managerApp/app/models/entities"
	"managerApp/app/models/mappers"
	"managerApp/app/helpers"
	_ "github.com/lib/pq"
)

type AuthProvider struct {
	db   *sql.DB
	auth *mappers.AuthMapper
}

func (p *AuthProvider) Init() error {
	var err error
	p.db, err = helpers.DBInit()
	if err != nil {
		return err
	}
	p.auth = new(mappers.AuthMapper)
	p.auth.Init(p.db)
	return nil
}

func (p *AuthProvider) Login(userName string, password string) (*entities.User, error) {
	user, err := p.auth.Login(userName, password)
	if err != nil {
		fmt.Printf("AuthProvider::Login:%v", err)
		return nil, err
	}
	return user, nil
}

func (p *AuthProvider) Logout() error {
	defer p.db.Close()

	return nil
}
