package mappers

import (
	"database/sql"
	"fmt"
	"managerApp/app/models/entities"
)

type InterviewerMapper struct {
	db *sql.DB
}

func (m *InterviewerMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}

//выбрать всех сотрудников
func (m *InterviewerMapper) SelectAll() ([]*entities.Interviewer, error) {
	var (
		c_id           sql.NullInt64
		c_surname      sql.NullString
		c_name         sql.NullString
		c_patronymic   sql.NullString
		c_email        sql.NullString
		c_phone_number sql.NullString
		c_position     sql.NullString
	)

	interviewers := make([]*entities.Interviewer, 0)
	//выбираем всех сотрудников из таблицы t_interviewer
	query := `SELECT i_id, i_last_name, i_first_name, i_mid_name, i_email, i_phone_num, i_position FROM t_interviewer`
	rows, err := m.db.Query(query)
	if err != nil {
		fmt.Print(err)
		return nil, fmt.Errorf("InterviewerMapper::Select:%v", err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&c_id, &c_surname, &c_name, &c_patronymic, &c_email, &c_phone_number, &c_position)
		interviewer := &entities.Interviewer{
			ID:          c_id.Int64,
			Surname:     c_surname.String,
			Name:        c_name.String,
			Patronymic:  c_patronymic.String,
			Email:       c_email.String,
			PhoneNumber: c_phone_number.String,
			Position:    c_position.String,
		}
		interviewers = append(interviewers, interviewer)
	}
	return interviewers, nil
}

//получить сотрудников в выбранном ассессменте
func (m *InterviewerMapper) Select(assessmentId int64) ([]*entities.Interviewer, error) {
	var (
		c_id           sql.NullInt64
		c_surname      sql.NullString
		c_name         sql.NullString
		c_patronymic   sql.NullString
		c_email        sql.NullString
		c_phone_number sql.NullString
		c_position     sql.NullString
	)

	interviewers := make([]*entities.Interviewer, 0)
	rows, err := m.db.Query("SELECT u.i_id, u.i_last_name, u.i_first_name, u.i_mid_name, i_email, i_phone_num, i_position FROM t_interviewer u INNER JOIN toc_assessment_interviewer d ON u.i_id = d.a_i_interviewer_id WHERE d.a_i_assessment_id = $1", assessmentId)
	if err != nil {
		fmt.Print(err)
		return nil, fmt.Errorf("InterviewerMapper::Select:%v", err)
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&c_id, &c_surname, &c_name, &c_patronymic, &c_email, &c_phone_number, &c_position)
		interviewer := &entities.Interviewer{
			ID:          c_id.Int64,
			Surname:     c_surname.String,
			Name:        c_name.String,
			Patronymic:  c_patronymic.String,
			Email:       c_email.String,
			PhoneNumber: c_phone_number.String,
			Position:    c_position.String,
		}
		interviewers = append(interviewers, interviewer)
	}
	return interviewers, nil
}

//получить выбранного сотрудника
func (m *InterviewerMapper) SelectById(interviewerId int64) (*entities.Interviewer, error) {
	var (
		c_id           sql.NullInt64
		c_surname      sql.NullString
		c_name         sql.NullString
		c_patronymic   sql.NullString
		c_email        sql.NullString
		c_phone_number sql.NullString
		c_position     sql.NullString
	)

	row := m.db.QueryRow("SELECT i_id, i_last_name, i_first_name, i_mid_name, i_email, i_phone_num, i_position FROM t_interviewer WHERE i_id = $1", interviewerId)

	err := row.Scan(&c_id, &c_surname, &c_name, &c_patronymic, &c_email, &c_phone_number, &c_position)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("InterviewerMapper::SelectById:%v", err)
	} else if err != nil {
		return nil, fmt.Errorf("InterviewerMapper::SelectById:%v", err)
	}
	interviewer := &entities.Interviewer{
		ID:          c_id.Int64,
		Surname:     c_surname.String,
		Name:        c_name.String,
		Patronymic:  c_patronymic.String,
		Email:       c_email.String,
		PhoneNumber: c_phone_number.String,
		Position:    c_position.String,
	}
	return interviewer, nil
}

//добавить сотрудника в ассессмент
func (m *InterviewerMapper) Insert(newInterviewer *entities.Interviewer, assessmentId int64) (int64, error) {
	var insertedId int64
	//добавление сотрудника к списку сотрудников
	insertQuery := `INSERT INTO t_interviewer 
		(i_id, i_last_name, i_first_name, i_mid_name, i_email, i_phone_num, i_position) 
		SELECT nextval('interviewer_id'), $1, $2, $3, $4, $5, $6 
		WHERE NOT EXISTS(SELECT i_id, i_last_name, i_first_name, i_mid_name, i_email, i_phone_num, i_position FROM t_interviewer WHERE i_last_name = $7 AND i_first_name = $8 AND i_mid_name = $9 AND i_position = $10)`
	_, err := m.db.Exec(insertQuery, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Email, newInterviewer.PhoneNumber, newInterviewer.Position, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Position)
	if err != nil {
		return 0, fmt.Errorf("Ошибка вставки сотрудника: %v", err)
	}
	//возвращаем его ID
	row := m.db.QueryRow(`select i_id FROM t_interviewer WHERE i_last_name = $1 AND i_first_name = $2 AND i_mid_name = $3 AND i_position = $4`, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Position)

	err = row.Scan(&insertedId)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	} else if err != nil {
		return 0, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	}
	//добавление сотрдуника в таблицу связи toc_assessment_interviewer
	insertQueryToAssess := `INSERT INTO toc_assessment_interviewer 
		(a_i_id, a_i_assessment_id, a_i_interviewer_id) 
		SELECT nextval('assessment_interviewer_id'), $1, $2 
		WHERE NOT EXISTS(SELECT a_i_id, a_i_assessment_id, a_i_interviewer_id FROM toc_assessment_interviewer WHERE a_i_interviewer_id = $3 AND a_i_assessment_id = $4)`
	_, err = m.db.Exec(insertQueryToAssess, assessmentId, insertedId, insertedId, assessmentId)
	if err != nil {
		return 0, fmt.Errorf("Ошибка вставки сотрудника в ассессмент: %v", err)
	}
	fmt.Println("New interviewer:", insertedId)
	return insertedId, nil
}

//изменить сотрудника
func (m *InterviewerMapper) Update(newInterviewer *entities.Interviewer, interviewerId int64) (int64, error) {
	insertQuery := `UPDATE t_interviewer 
		SET i_last_name = $1, i_first_name = $2, i_mid_name = $3, i_email = $4, i_phone_num = $5, i_position = $6
		WHERE i_id = $7`
	_, err := m.db.Exec(insertQuery, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Email, newInterviewer.PhoneNumber, newInterviewer.Position, interviewerId)
	if err != nil {
		return 0, fmt.Errorf("Ошибка обновления сотрудника: %v", err)
	}
	return interviewerId, nil
}

//добавить сотрудника к списку сотрдуников
func (m *InterviewerMapper) InsertInterviewer(newInterviewer *entities.Interviewer) (int64, error) {
	var insertedId int64
	//добавить к списку
	insertQuery := `INSERT INTO t_interviewer 
		(i_id, i_last_name, i_first_name, i_mid_name, i_email, i_phone_num, i_position) 
		SELECT nextval('interviewer_id'), $1, $2, $3, $4, $5, $6 
		WHERE NOT EXISTS(SELECT i_id, i_last_name, i_first_name, i_mid_name, i_email, i_phone_num, i_position FROM t_interviewer WHERE i_last_name = $7 AND i_first_name = $8 AND i_mid_name = $9 AND i_position = $10)`
	_, err := m.db.Exec(insertQuery, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Email, newInterviewer.PhoneNumber, newInterviewer.Position, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Position)
	if err != nil {
		return 0, fmt.Errorf("Ошибка вставки сотрудника: %v", err)
	}
	//вернуть его ID
	row := m.db.QueryRow(`select i_id FROM t_interviewer WHERE i_last_name = $1 AND i_first_name = $2 AND i_mid_name = $3 AND i_position = $4`, newInterviewer.Surname, newInterviewer.Name, newInterviewer.Patronymic, newInterviewer.Position)
	err = row.Scan(&insertedId)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	} else if err != nil {
		return 0, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	}
	return insertedId, nil
}

//удаление сотрудника из ассессмента
func (m *InterviewerMapper) Delete(id int64, idAssessment int64) error {
	_, err := m.db.Exec("DELETE FROM toc_assessment_interviewer WHERE a_i_interviewer_id = $1 AND a_i_assessment_id = $2", id, idAssessment)
	if err == sql.ErrNoRows {
		return fmt.Errorf("InterviewerMapper::Delete:%v", err)
	} else if err != nil {
		return fmt.Errorf("InterviewerMapper::Delete:%v", err)
	}
	return nil
}

//удалить из словаря
func (m *InterviewerMapper) DeleteFromD(id int64) error {

	deleteQuery := `DELETE FROM toc_assessment_interviewer WHERE a_i_interviewer_id = $1;`

	_, err := m.db.Exec(deleteQuery, id)
	if err == sql.ErrNoRows {
		return fmt.Errorf("InterviewerMapper::DeleteFromD:%v", err)
	} else if err != nil {
		return fmt.Errorf("InterviewerMapper::DeleteFromD:%v", err)
	}

	deleteQuery = `DELETE FROM t_interviewer WHERE i_id = $1;`

	_, err = m.db.Exec(deleteQuery, id)
	if err == sql.ErrNoRows {
		return fmt.Errorf("InterviewerMapper::DeleteFromD:%v", err)
	} else if err != nil {
		return fmt.Errorf("InterviewerMapper::DeleteFromD:%v", err)
	}
	return nil
}
