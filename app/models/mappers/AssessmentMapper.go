package mappers

import (
	"database/sql"
	"fmt"
	entities "managerApp/app/models/entities"
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
	)
	//делаем запрос к бд, находим ассессмент по ID
	//и считываем
	row := m.db.QueryRow("SELECT a_id, a_date FROM t_assessment WHERE a_id = $1", assessmentId)
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
	updateQuery := `UPDATE t_assessment SET a_date = $1 WHERE a_id = $2`
	err := m.db.QueryRow(updateQuery, newAssessment.Date, assessmentId)
	if err != nil {
		return 0, fmt.Errorf("Ошибка вставки ассессмента: ", err)
	}
	return assessmentId, nil
}

//удаление ассессмента
func (m *AssessmentMapper) Delete(assessmentId int64) error {
	//удаляем из таблицы смежности toc_assessment_candidate
	_, err := m.db.Exec("DELETE FROM toc_assessment_candidate WHERE a_c_assessment_id = $1", assessmentId)
	if err == sql.ErrNoRows {
		return fmt.Errorf("AssessmentMapper::Delete:%v", err)
	} else if err != nil {
		return fmt.Errorf("AssessmentMapper::Delete:%v", err)
	}
	//удаляем из таблицы смежности toc_assessment_interviewer
	_, err = m.db.Exec("DELETE FROM toc_assessment_interviewer WHERE a_i_assessment_id = $1", assessmentId)
	if err == sql.ErrNoRows {
		return fmt.Errorf("AssessmentMapper::Delete:%v", err)
	} else if err != nil {
		return fmt.Errorf("AssessmentMapper::Delete:%v", err)
	}
	//удаляем из таблицы t_assessment
	_, err = m.db.Exec("DELETE FROM t_assessment WHERE a_id = $1", assessmentId)
	if err == sql.ErrNoRows {
		return fmt.Errorf("AssessmentMapper::Delete:%v", err)
	} else if err != nil {
		return fmt.Errorf("AssessmentMapper::Delete:%v", err)
	}
	return nil
}

//выбор всех ассессментов из базы
func (m *AssessmentMapper) Select() ([]entities.Assessment, error) {
	assessments := []entities.Assessment{}
	fmt.Println("Created array:", assessments)
	queryStr := `SELECT up.a_id, to_char(up.a_date, 'DD.MM.YYYY HH24:MI'), down.a_s_name FROM t_assessment up INNER JOIN t_assessment_status down ON down.a_s_id = up.a_status` //
	rows, err := m.db.Query(queryStr)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var (
			id     int64
			date   string
			status string
		)
		err = rows.Scan(&id, &date, &status)
		fmt.Println("Vars:", id, date, status)
		if err == nil {
			assessment := entities.Assessment{ID: id, Date: date, StatusName: status}
			fmt.Println("Assessment:", assessment)
			assessments = append(assessments, assessment)
		}
		fmt.Println("Error:", err)
	}
	defer rows.Close()
	fmt.Println("Mapper GET:", assessments)
	return assessments, nil
}

//вставка ассессмента
func (m *AssessmentMapper) Insert(newAssessment entities.Assessment) (int64, error) {
	var insertedId int64
	insertQuery := `INSERT INTO t_assessment 
		(a_date, a_status) 
		SELECT to_timestamp($1,'YYYY-MM-DD HH24:MI:SS'), $2 
		WHERE NOT EXISTS(SELECT a_date, a_status FROM t_assessment WHERE a_date = to_timestamp($3,'YYYY-MM-DD HH24:MI:SS'))
		returning a_id;`
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
	query := `SELECT a_s_id, a_s_name FROM t_assessment_status u INNER JOIN t_assessment d
			ON a_id = $1`
	rows, err := m.db.Query(query, assessmentId)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выбора всех статусов:%v", err)
	}
	defer rows.Close()
	//считываем данные
	for rows.Next() {
		rows.Scan(&c_id, &c_status)
		status := &entities.AssessmentStatus{
			ID:   c_id,
			Name: c_status,
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}

//задать статус ассессменту
func (m *AssessmentMapper) SetStatus(newStatus *entities.AssessmentStatus, statusId int64, assessmentId int64) (int64, error) {
	insertQuery := `UPDATE t_assessment SET a_status = $1 WHERE a_id = $2`
	_, err := m.db.Exec(insertQuery, statusId, assessmentId)
	if err != nil {
		return 0, fmt.Errorf("Ошибка изменения статуса ассессмента: %v", err)
	}
	return assessmentId, nil
}
