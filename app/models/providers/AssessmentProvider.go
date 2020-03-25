package providers

import (
	"database/sql"
	"fmt"
	"managerApp/app/models/entities"
	"managerApp/app/models/mappers"

	_ "github.com/lib/pq"
	_ "github.com/revel/revel"
)

type AssessmentProvider struct {
	db          *sql.DB
	assessments *mappers.AssessmentMapper
}

//инициализация и подключение к БД
func (p *AssessmentProvider) Init() {
	var err error
	connStr := "user=postgres password=password port=5433 dbname=AssessmentManager sslmode=disable"
	p.db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = p.db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PostgreSQL.")
	p.assessments = new(mappers.AssessmentMapper)
	p.assessments.Init(p.db)
}

//изменение ассессмента
func (p *AssessmentProvider) PostAssessment(newAssessment *entities.Assessment, assessmentId int64) (*entities.Assessment, error) {
	defer p.db.Close()
	id, err := p.assessments.Update(newAssessment, assessmentId)
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("AssessmentProvider::PostAssessment:%v", err)
		}
	}
	updatedAssessment, err := p.assessments.SelectById(id)
	if err != nil {
		return nil, fmt.Errorf("AssessmentProvider::PostAssessment:%v", err)
	}
	return updatedAssessment, nil
}

//удаление ассессмента
func (p *AssessmentProvider) DeleteAssessment(assessmentId int64) error {
	defer p.db.Close()
	err := p.assessments.Delete(assessmentId)
	if err != nil {
		return fmt.Errorf("AssessmentProvider::DeleteAssessment:%v", err)
	}
	return nil
}

//создание ассессмента
func (p *AssessmentProvider) PutAssessment(newAssessment entities.Assessment) (*entities.Assessment, error) {
	defer p.db.Close()
	//задаём ID начального статуса 1 ("назначен")
	newAssessment.Status = 1
	id, err := p.assessments.Insert(newAssessment)
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("AssessmentProvider::PutAssessment:%v", err)
		}
	}
	createdAssessment, err := p.assessments.SelectById(id)
	if err != nil {
		return nil, fmt.Errorf("AssessmentProvider::PutAssessment:%v", err)
	}
	return createdAssessment, nil
}

//получить все ассессменты из базы
func (p *AssessmentProvider) GetAssessments() ([]entities.Assessment, error) {
	assessments, err := p.assessments.Select()
	if err != nil {
		return nil, err
	}
	return assessments, nil
}

//получить выбранный ассессмент
func (p *AssessmentProvider) GetAssessmentById(id int64) (*entities.Assessment, error) {
	defer p.db.Close()
	assessment, err := p.assessments.SelectById(id)
	return assessment, err
}

// получить возможные статусы ассессмента
func (p *AssessmentProvider) GetAssessmentStatus(id int64) ([]*entities.AssessmentStatus, error) {
	defer p.db.Close()
	assessment, err := p.assessments.SelectStatus(id)
	return assessment, err
}

//задать статус ассессмента
func (p *AssessmentProvider) SetAssessmentStatus(newStatus *entities.AssessmentStatus, statusId int64, assessmentId int64) (int64, error) {
	defer p.db.Close()
	status, err := p.assessments.SetStatus(newStatus, statusId, assessmentId)
	return status, err
}
