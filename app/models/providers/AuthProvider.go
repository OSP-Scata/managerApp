package providers

import (
	"database/sql"
	"fmt"
	"managerApp/app/models/entities"
	"managerApp/app/models/mappers"

	_ "github.com/lib/pq"
)

type AuthProvider struct {
	db   *sql.DB
	auth *mappers.AuthMapper
}

func (p *AuthProvider) Init() {
	//defer p.db.Close()
	//подключение к БД
	var err error
	connStr := "user=postgres password=password port=5433 dbname=AssessmentManager sslmode=disable"
	p.db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	p.auth = new(mappers.AuthMapper)
	p.auth.Init(p.db)
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
