package mappers

import (
	entities "managerApp/app/models/entities"
	"database/sql"
	"fmt"
)

type AssessmentMapper struct {
	db *sql.DB
}

func (m *AssessmentMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}

//получить выбранный ассессмент
func (m *AssessmentMapper) SelectById(assessmentId int64) (*entities.Assessment, error) {
	var (
		c_id   sql.NullInt64
		c_date sql.NullString
		//db *sql.DB
	)
	//делаем запрос к бд, находим ассессмент по ID
	//и считываем
	row := m.db.QueryRow("SELECT c_id, c_date FROM t_assessment WHERE c_id = $1", assessmentId)
	//записываем данные в переменные
	err := row.Scan(&c_id, &c_date)
	//если по запросу не нашлось ни одного ассессмента или случилась иная ошибка:
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("AssessmentMapper::SelectById:%v", err)
	} else if err != nil {
		return nil, fmt.Errorf("AssessmentMapper::SelectById:%v", err)
	}

	//создаем объект и заполняем его полученными данными
	assessment := &entities.Assessment{ID: c_id.Int64,
		Date: c_date.String,
	}
	fmt.Println(assessment)
	//и возвращаем его
	return assessment, nil
}
//изменение ассессмента
func (m *AssessmentMapper) Update(newAssessment *entities.Assessment, assessmentId int64) (int64, error) {
	updateQuery := `UPDATE t_assessment SET c_date = $1 WHERE c_id = $2`
	err := m.db.QueryRow(updateQuery, newAssessment.Date, assessmentId)
	if err != nil {
		return 0, fmt.Errorf("Ошибка вставки ассессмента: ", err)
	}
	return assessmentId, nil
}
//удаление ассессмента
func (m *AssessmentMapper) Delete(assessmentId int64) error {
	//удаляем из таблицы смежности toc_assessment_candidate
	_, err := m.db.Exec("DELETE FROM toc_assessment_candidate WHERE c_id_assessment = $1", assessmentId)
	if err == sql.ErrNoRows {
		return fmt.Errorf("AssessmentMapper::Delete:%v", err)
	} else if err != nil {
		return fmt.Errorf("AssessmentMapper::Delete:%v", err)
	}

	//удаляем из таблицы смежности toc_assessment_interviewer
	_, err = m.db.Exec("DELETE FROM toc_assessment_interviewer WHERE c_id_assessment = $1", assessmentId)
	if err == sql.ErrNoRows {
		return fmt.Errorf("AssessmentMapper::Delete:%v", err)
	} else if err != nil {
		return fmt.Errorf("AssessmentMapper::Delete:%v", err)
	}

	//удаляем из таблицы t_assessment
	_, err = m.db.Exec("DELETE FROM t_assessment WHERE c_id = $1", assessmentId)
	if err == sql.ErrNoRows {
		return fmt.Errorf("AssessmentMapper::Delete:%v", err)
	} else if err != nil {
		return fmt.Errorf("AssessmentMapper::Delete:%v", err)
	}
	return nil
}
//выбор ассессментов
func (m *AssessmentMapper) Select() ([]entities.Assessment, error) {
	//Retrieve
	//m.Init(m.db)
	assessments := []entities.Assessment{}
	queryStr := `SELECT u.c_id, to_char(u.c_date, 'DD.MM.YYYY HH:MM'), d.c_status FROM t_assessment u INNER JOIN t_status_assessment d ON u.fk_status = d.c_id`
	rows, err := m.db.Query(queryStr)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id     int64
			date   string
			status string
		)
		err = rows.Scan(&id, &date, &status)
		if err == nil {
			assessment := entities.Assessment{ID: id, Date: date, StatusName: status}
			assessments = append(assessments, assessment)
		}
	}
	//fmt.Println(assessments)
	return assessments, nil
}
//вставка ассессмента
func (m *AssessmentMapper) Insert(newAssessment entities.Assessment) (int64, error) {
	var insertedId int64
	insertQuery := `INSERT INTO t_assessment 
		(c_id, c_date, fk_status) 
		SELECT nextval('assessment_id'), to_timestamp($1,'YYYY-MM-DD HH24:MI:SS'), $2 
		WHERE NOT EXISTS(SELECT c_id, c_date, fk_status FROM t_assessment WHERE c_date = to_timestamp($3,'YYYY-MM-DD HH24:MI:SS'))
		returning c_id;`
	err := m.db.QueryRow(insertQuery, newAssessment.Date, newAssessment.Status, newAssessment.Date).Scan(&insertedId)
	if err != nil {
		fmt.Println("Ошибка вставки ассессмента: %v", err)
	}
	fmt.Println(newAssessment)
	return insertedId, nil
}
//получить статусы
func (m *AssessmentMapper) SelectStatus(assessmentId int64) ([]*entities.AssessmentStatus, error) {
	var (
		c_id     int64
		c_status string
	)

	statuses := make([]*entities.AssessmentStatus, 0)
	//запрос к БД
	query := `SELECT u.c_id, u.c_status FROM t_status_assessment u INNER JOIN t_assessment d
			ON d.c_id = $1`

	rows, err := m.db.Query(query, assessmentId)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выбора всех статусов:%v", err)
	}
	defer rows.Close()
	//считываем данные
	for rows.Next() {
		rows.Scan(&c_id, &c_status)
		status := &entities.AssessmentStatus{
			ID:     c_id,
			Name: c_status,
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}

//задать статус ассессменту
func (m *AssessmentMapper) SetStatus(newStatus *entities.AssessmentStatus, statusId int64, assessmentId int64) (int64, error) {

	//запрос к БД
	insertQuery := `UPDATE t_assessment SET fk_status = $1 WHERE c_id = $2`
	_, err := m.db.Exec(insertQuery, statusId, assessmentId)
	if err != nil {
		return 0, fmt.Errorf("Ошибка изменения статуса ассессмента: %v", err)
	}
	//возвращаем статус выбранного ассессмента
	return assessmentId, nil
}