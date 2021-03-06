package mappers

import (
	"database/sql"
	"fmt"
	entities "managerApp/app/models/entities"
)

type CandidateMapper struct {
	db *sql.DB
}

func (m *CandidateMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}

//выбрать всех кандидатов
func (m *CandidateMapper) SelectAllCandidates() ([]*entities.Candidate, error) {
	var (
		c_id           sql.NullInt64
		c_last_name    sql.NullString
		c_first_name   sql.NullString
		c_mid_name     sql.NullString
		c_email        sql.NullString
		c_phone_number sql.NullString
		c_education    sql.NullString
		c_birth_date   sql.NullString
	)
	candidates := make([]*entities.Candidate, 0)
	query := `SELECT c_id, c_last_name, c_first_name, c_mid_name, c_email, c_phone_number, 
		to_char(c_birth_date, 'YYYY-MM-DD'), c_education 
		FROM t_candidate`
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("CandidateMapper::SelectAllCandidates:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&c_id, &c_last_name, &c_first_name, &c_mid_name, &c_email, &c_phone_number, &c_birth_date, &c_education)
		candidate := &entities.Candidate{ID: c_id.Int64,
			Surname:     c_last_name.String,
			Name:        c_first_name.String,
			Patronymic:  c_mid_name.String,
			Email:       c_email.String,
			PhoneNumber: c_phone_number.String,
			Education:   c_education.String,
			BirthDate:   c_birth_date.String,
		}
		candidates = append(candidates, candidate)
	}
	fmt.Println(candidates)
	return candidates, nil
}

//получить всех кандидатов, которые участвуют в выбранном асссесменте
func (m *CandidateMapper) Select(assessmentId int64) ([]*entities.Candidate, error) {
	var (
		c_id           sql.NullInt64
		c_last_name    sql.NullString
		c_first_name   sql.NullString
		c_mid_name     sql.NullString
		c_email        sql.NullString
		c_phone_number sql.NullString
		c_education    sql.NullString
		c_birth_date   sql.NullString
		c_status_name  sql.NullString
	)

	candidates := make([]*entities.Candidate, 0)
	//запрос к БД
	query := `SELECT u.c_id, u.c_last_name, u.c_first_name, u.c_mid_name, u.c_email, u.c_phone_number, 
	to_char(u.c_birth_date, 'DD.MM.YYYY'), u.c_education, v.c_s_name 
	FROM t_candidate u INNER JOIN toc_assessment_candidate d ON u.c_id = d.a_c_candidate_id 
	INNER JOIN t_candidate_status v ON d.a_c_candidate_status = v.c_s_id WHERE d.a_c_assessment_id = $1`
	rows, err := m.db.Query(query, assessmentId)
	if err != nil {
		return nil, fmt.Errorf("CandidateMapper::Select:%v", err)
	}
	defer rows.Close()
	for rows.Next() {
		//считываем данные и записываем в candidate
		rows.Scan(&c_id, &c_last_name, &c_first_name, &c_mid_name, &c_email, &c_phone_number, &c_birth_date, &c_education, &c_status_name)
		candidate := &entities.Candidate{ID: c_id.Int64,
			Surname:     c_last_name.String,
			Name:        c_first_name.String,
			Patronymic:  c_mid_name.String,
			Email:       c_email.String,
			PhoneNumber: c_phone_number.String,
			Education:   c_education.String,
			BirthDate:   c_birth_date.String,
			StatusName:  c_status_name.String,
		}
		//добавляем candidate к созданному срезу
		candidates = append(candidates, candidate)
	}
	fmt.Println("Candidates in assessment:", candidates)
	return candidates, nil
}

//получить выбранного кандидата
func (m *CandidateMapper) SelectById(id int64, assessmentId int64) (*entities.Candidate, error) {
	var (
		c_id           sql.NullInt64
		c_last_name    sql.NullString
		c_first_name   sql.NullString
		c_mid_name     sql.NullString
		c_email        sql.NullString
		c_phone_number sql.NullString
		c_education    sql.NullString
		c_birth_date   sql.NullString
		c_status_name  sql.NullString
	)
	//обращаемся к БД
	query := `SELECT u.c_id, u.c_last_name, u.c_first_name, u.c_mid_name, u.c_email, u.c_phone_number, 
		u.c_birth_date, u.c_education, v.c_s_name 
		FROM t_candidate u INNER JOIN toc_assessment_candidate d ON u.c_id = d.a_c_candidate_id 
		INNER JOIN t_candidate_status v ON d.a_c_candidate_status = v.c_s_id WHERE d.a_c_assessment_id = $1 AND u.c_id = $2`
	//выполняем
	row := m.db.QueryRow(query, assessmentId, id)
	//считываем
	err := row.Scan(&c_id, &c_last_name, &c_first_name, &c_mid_name, &c_email, &c_phone_number, &c_birth_date, &c_education, &c_status_name)
	//выдаем ошибку если по результту ничего не найдено или произошла иная ошибка
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	} else if err != nil {
		return nil, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	}

	//записываем данные если нет ошибок
	candidate := &entities.Candidate{ID: c_id.Int64,
		Surname:     c_last_name.String,
		Name:        c_first_name.String,
		Patronymic:  c_mid_name.String,
		Email:       c_email.String,
		PhoneNumber: c_phone_number.String,
		Education:   c_education.String,
		BirthDate:   c_birth_date.String,
		StatusName:  c_status_name.String,
	}

	return candidate, nil
}

//получить возможные статусы
func (m *CandidateMapper) SelectStatus(candidateId int64, assessmentId int64) ([]*entities.CandidateStatus, error) {
	var (
		c_id     sql.NullInt64
		c_status sql.NullString
	)

	statuses := make([]*entities.CandidateStatus, 0)
	//запросы к БД
	//получаем родительский статус кандидата
	query := `SELECT c_s_id, c_s_name FROM t_candidate_status 
		WHERE c_s_id = (select c_s_fk FROM t_candidate_status where c_s_id = 
		(select a_c_candidate_status FROM toc_assessment_candidate where 
		c_id_candidate = $1));`
	//получаем возможные статусы
	query2 := `SELECT u.c_s_id, u.c_s_name FROM t_candidate_status u INNER JOIN toc_assessment_candidate d ON 
	d.a_c_candidate_status = u.c_s_fk WHERE d.a_c_candidate_id = $1 AND d.a_c_assessment_id = $2`
	rows, err := m.db.Query(query, candidateId)
	if err != nil {
		return nil, fmt.Errorf("CandidateMapper::SelectStatus:%v", err)
	}
	rows2, erro := m.db.Query(query2, candidateId, assessmentId)
	if erro != nil {
		return nil, fmt.Errorf("CandidateMapper::SelectStatus:%v", erro)
	}
	defer rows.Close()
	//получаем данные
	for rows.Next() {
		rows.Scan(&c_id, &c_status)
		status := &entities.CandidateStatus{
			ID:   c_id.Int64,
			Name: c_status.String,
		}
		statuses = append(statuses, status)
	}
	defer rows2.Close()
	for rows2.Next() {
		rows2.Scan(&c_id, &c_status)
		status := &entities.CandidateStatus{
			ID:   c_id.Int64,
			Name: c_status.String,
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}

//задать статус кандидата
func (m *CandidateMapper) SetStatus(newStatus *entities.CandidateStatus, statusId int64, candidateId int64) (int64, error) {
	insertQuery := `UPDATE toc_assessment_candidate SET a_c_candidate_status = $1 WHERE a_c_candidate_id = $2`
	_, err := m.db.Exec(insertQuery, statusId, candidateId)
	if err != nil {
		return 0, fmt.Errorf("Ошибка изменения статуса кандидата: %v", err)
	}
	return statusId, nil
}

//задать статус кандидата
func (m *CandidateMapper) SetStatus2(newStatus *entities.CandidateStatus, statusId int64, candidateId int64) (int64, error) {
	insertQuery := `UPDATE toc_assessment_candidate SET a_c_candidate_status = $1 WHERE a_c_candidate_id = $2`
	_, err := m.db.Exec(insertQuery, statusId, candidateId)
	if err != nil {
		return 0, fmt.Errorf("Ошибка изменения статуса кандидата: %v", err)
	}
	return statusId, nil
}

//создание кандидата
func (m *CandidateMapper) Insert(newCandidate *entities.Candidate, assessmentId int64) (int64, error) {
	var insertedId int64
	//обращения к БД
	//добавляем кандидата к списку кандидатов
	insertQuery := `INSERT INTO t_candidate 
		(c_last_name, c_first_name, c_mid_name, c_email, c_phone_number, c_birth_date, c_education) 
		SELECT $1, $2, $3, $4, $5, to_date($6,'YYYY-MM-DD'), $7
		WHERE NOT EXISTS(SELECT c_last_name, c_first_name, c_mid_name, c_email, c_phone_number, c_birth_date, c_education FROM t_candidate WHERE c_last_name = $8 AND c_first_name = $9 AND c_mid_name = $10 AND c_birth_date = to_date($11,'YYYY-MM-DD'))
		`
	_, err := m.db.Exec(insertQuery, newCandidate.Surname, newCandidate.Name, newCandidate.Patronymic, newCandidate.Email, newCandidate.PhoneNumber, newCandidate.BirthDate, newCandidate.Education, newCandidate.Surname, newCandidate.Name, newCandidate.Patronymic, newCandidate.BirthDate)
	if err != nil {
		return 0, fmt.Errorf("Ошибка вставки кандидата: %v", err)
	}
	//получаем его ID
	row := m.db.QueryRow(`select c_id FROM t_candidate WHERE c_last_name = $1 AND c_first_name = $2 AND c_mid_name = $3 AND c_birth_date = to_date($4,'YYYY-MM-DD')`, newCandidate.Surname, newCandidate.Name, newCandidate.Patronymic, newCandidate.BirthDate)

	err = row.Scan(&insertedId)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	} else if err != nil {
		return 0, fmt.Errorf("CandidateMapper::SelectById:%v", err)
	}
	//добавляем кандидата в таблицу связи toc_assessment_candidate
	insertQueryToAssess := `INSERT INTO toc_assessment_candidate 
		(a_c_assessment_id, a_c_candidate_id, a_c_candidate_status) 
		SELECT $1, $2, 1 
		WHERE NOT EXISTS(SELECT a_c_assessment_id, a_c_candidate_id FROM toc_assessment_candidate WHERE a_c_candidate_id = $3 AND a_c_assessment_id = $4)`
	_, err = m.db.Exec(insertQueryToAssess, assessmentId, insertedId, insertedId, assessmentId)

	if err != nil {
		return 0, fmt.Errorf("Ошибка вставки кандидата в ассессмент: %v", err)
	}
	fmt.Println("New candidate:", insertedId)
	return insertedId, nil
}

//изменение кандидата
func (m *CandidateMapper) Update(newCandidate *entities.Candidate, candidateId int64) (int64, error) {
	//обращение к БД
	insertQuery := `UPDATE t_candidate 
		SET c_last_name = $1, c_first_name = $2, c_mid_name = $3, c_email = $4, c_phone_number = $5, c_birth_date = to_date($6,'YYYY-MM-DD'), c_education = $7 
		WHERE c_id = $8`
	_, err := m.db.Exec(insertQuery, newCandidate.Surname, newCandidate.Name, newCandidate.Patronymic, newCandidate.Email, newCandidate.PhoneNumber, newCandidate.BirthDate, newCandidate.Education, candidateId)
	if err != nil {
		return 0, fmt.Errorf("Ошибка обновления кандидата: %v", err)
	}
	return candidateId, nil
}

//удаление
func (m *CandidateMapper) Delete(id int64, idAssessment int64) error {
	//обращение к БД
	_, err := m.db.Exec("DELETE FROM toc_assessment_candidate WHERE a_c_candidate_id = $1 AND a_c_assessment_id = $2", id, idAssessment)
	fmt.Print("ID Candidate: ", id)
	if err == sql.ErrNoRows {
		return fmt.Errorf("CandidateMapper::Delete:%v", err)
	} else if err != nil {
		return fmt.Errorf("CandidateMapper::Delete:%v", err)
	}
	return nil
}
