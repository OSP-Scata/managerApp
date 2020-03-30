--База для приложения - менеджера собеседований, который мне надо разработать как тестовый проект

CREATE TABLE IF NOT EXISTS t_user ( 
	u_id                 serial PRIMARY KEY NOT NULL,
	u_login              varchar(32)  NOT NULL,
	u_password           varchar(32)  NOT NULL
 );
 
 CREATE TABLE IF NOT EXISTS t_assessment_status ( 
	a_s_id               serial PRIMARY KEY NOT NULL,
	a_s_name             varchar(20),
	a_s_fk INTEGER REFERENCES t_assessment_status (a_s_id) 
 );

CREATE TABLE IF NOT EXISTS t_candidate_status ( 
	c_s_id               serial PRIMARY KEY NOT NULL,
	c_s_name             varchar(20),
	c_s_fk INTEGER REFERENCES t_candidate_status (c_s_id)
 );

CREATE TABLE IF NOT EXISTS t_assessment ( 
	a_id                 serial PRIMARY KEY NOT NULL,
	a_date               timestamp  NOT NULL,
	a_status             integer,
	FOREIGN KEY (a_status) REFERENCES t_assessment_status (a_s_fk)
 );

CREATE SEQUENCE IF NOT EXISTS assessment_id
	START 1
    INCREMENT BY 1
	OWNED BY t_assessment.a_id;
	
CREATE TABLE IF NOT EXISTS t_interviewer ( 
	i_id                 serial PRIMARY KEY NOT NULL,
	i_last_name          varchar(100)  NOT NULL,
	i_first_name         varchar(50),
	i_mid_name           varchar(100),
	i_email              varchar(100),
	i_phone_num          varchar(20),
	i_position           varchar(100)
 );

 CREATE SEQUENCE IF NOT EXISTS interviewer_id
	START 1
    INCREMENT BY 1
	OWNED BY t_interviewer.i_id;

CREATE TABLE IF NOT EXISTS t_candidate ( 
	c_id                 serial PRIMARY KEY NOT NULL,
	c_last_name          varchar(100)  NOT NULL,
	c_first_name         varchar(50),
	c_mid_name           varchar(100),
	c_birth_date         timestamp,
	c_email              varchar(100),
	c_phone_num          varchar(20),
	c_education          varchar(200),
	c_status             integer,
	FOREIGN KEY (c_status) REFERENCES t_candidate_status (c_s_id)
 );

CREATE SEQUENCE IF NOT EXISTS candidate_id
	START 1
    INCREMENT BY 1
	OWNED BY t_candidate.c_id;

CREATE TABLE IF NOT EXISTS toc_assessment_candidate ( 
	a_c_id               integer PRIMARY KEY NOT NULL,
	a_c_assessment_id    integer,
	a_c_candidate_id     integer,
	a_c_candidate_status varchar(20),
	FOREIGN KEY (a_c_assessment_id) REFERENCES t_assessment (a_id),
	FOREIGN KEY (a_c_candidate_id) REFERENCES t_candidate (c_id)
 );

CREATE TABLE IF NOT EXISTS toc_assessment_interviewer ( 
	a_i_id               integer PRIMARY KEY NOT NULL,
	a_i_assessment_id    integer,
	a_i_interviewer_id   integer,
	FOREIGN KEY (a_i_assessment_id) REFERENCES t_assessment (a_id),
	FOREIGN KEY (a_i_interviewer_id) REFERENCES t_interviewer (i_id)
 );