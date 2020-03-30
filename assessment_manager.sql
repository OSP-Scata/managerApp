CREATE SCHEMA assessment_manager;

CREATE TABLE IF NOT EXISTS assessment_manager.t_assessment_status ( 
	a_s_id               serial PRIMARY KEY NOT NULL ,
	a_s_name             varchar(20)  NOT NULL ,
	CONSTRAINT pk_t_assessment_status_a_s_id PRIMARY KEY ( a_s_id ),
	CONSTRAINT unq_t_assessment_status_a_s_name UNIQUE ( a_s_name ) 
 );

CREATE TABLE IF NOT EXISTS assessment_manager.t_candidate_status ( 
	c_s_id               serial  NOT NULL ,
	c_s_name             varchar(20)  NOT NULL ,
	CONSTRAINT pk_t_candidate_status_c_s_id PRIMARY KEY ( c_s_id ),
	CONSTRAINT unq_t_candidate_status_c_s_name UNIQUE ( c_s_name ) 
 );

CREATE TABLE IF NOT EXISTS assessment_manager.t_interviewer ( 
	i_id                 serial PRIMARY KEY NOT NULL ,
	i_last_name          varchar(100)  NOT NULL ,
	i_first_name         varchar(50)   ,
	i_mid_name           varchar(100)   ,
	i_email              varchar(100)   ,
	i_phone_num          varchar(20)   ,
	i_position           varchar(100)   ,
	CONSTRAINT pk_t_interviewer_i_id PRIMARY KEY ( i_id )
 );

 CREATE SEQUENCE IF NOT EXISTS interviewer_id
	START 1
    INCREMENT BY 1
	OWNED BY t_interviewer.i_id;

CREATE TABLE IF NOT EXISTS assessment_manager.t_user ( 
	u_id                 serial PRIMARY KEY NOT NULL ,
	u_login              varchar(32)  NOT NULL ,
	u_password           varchar(32)  NOT NULL ,
	CONSTRAINT pk_t_user_u_id PRIMARY KEY ( u_id )
 );

CREATE TABLE IF NOT EXISTS assessment_manager.t_assessment ( 
	a_id                 serial PRIMARY KEY NOT NULL ,
	a_date               timestamp  NOT NULL ,
	a_status             varchar(20)  NOT NULL ,
	CONSTRAINT pk_t_assessments_id PRIMARY KEY ( a_id ),
	CONSTRAINT unq_t_assessment_a_status UNIQUE ( a_status ) 
 );

 CREATE SEQUENCE IF NOT EXISTS assessment_id
	START 1
    INCREMENT BY 1
	OWNED BY t_assessment.a_id;

COMMENT ON COLUMN assessment_manager.t_assessment.a_id IS 'Assessment ID';

COMMENT ON COLUMN assessment_manager.t_assessment.a_date IS 'Assessment date and time';

CREATE TABLE IF NOT EXISTS assessment_manager.t_candidate ( 
	c_id                 serial PRIMARY KEY NOT NULL ,
	c_last_name          varchar(100)  NOT NULL ,
	c_first_name         varchar(50)   ,
	c_mid_name           varchar(100)   ,
	c_birth_date         timestamp   ,
	c_email              varchar(100)   ,
	c_phone_num          varchar(20)   ,
	c_education          varchar(200)   ,
	c_status             varchar(20)  NOT NULL ,
	CONSTRAINT pk_t_candidate_c_id PRIMARY KEY ( c_id ),
	CONSTRAINT unq_t_candidate_c_status UNIQUE ( c_status ) 
 );

CREATE SEQUENCE IF NOT EXISTS candidate_id
	START 1
    INCREMENT BY 1
	OWNED BY t_candidate.c_id;

CREATE TABLE IF NOT EXISTS assessment_manager.toc_assessment_candidate ( 
	a_c_id               integer PRIMARY KEY NOT NULL   ,
	a_c_assessment_id    integer   ,
	a_c_candidate_id     integer   ,
	a_c_candidate_status varchar(20)   
 );

CREATE TABLE IF NOT EXISTS assessment_manager.toc_assessment_interviewer ( 
	a_i_id               integer PRIMARY KEY NOT NULL   ,
	a_i_assessment_id    integer   ,
	a_i_interviewer_id   integer   
 );

ALTER TABLE assessment_manager.t_assessment ADD CONSTRAINT fk_t_assessment_t_assessment_status FOREIGN KEY ( a_status ) REFERENCES assessment_manager.t_assessment_status( a_s_name );

ALTER TABLE assessment_manager.t_candidate ADD CONSTRAINT fk_t_candidate_t_candidate_status FOREIGN KEY ( c_status ) REFERENCES assessment_manager.t_candidate_status( c_s_name );

ALTER TABLE assessment_manager.toc_assessment_candidate ADD CONSTRAINT fk_toc_assessment_candidate_t_assessment FOREIGN KEY ( a_c_assessment_id ) REFERENCES assessment_manager.t_assessment( a_id );

ALTER TABLE assessment_manager.toc_assessment_candidate ADD CONSTRAINT fk_toc_assessment_candidate_t_candidate FOREIGN KEY ( a_c_candidate_id ) REFERENCES assessment_manager.t_candidate( c_id );

ALTER TABLE assessment_manager.toc_assessment_candidate ADD CONSTRAINT fk_toc_assessment_candidate_t_candidate_status FOREIGN KEY ( a_c_candidate_status ) REFERENCES assessment_manager.t_candidate_status( c_s_name );

ALTER TABLE assessment_manager.toc_assessment_interviewer ADD CONSTRAINT fk_toc_assessment_interviewer_t_assessment FOREIGN KEY ( a_i_assessment_id ) REFERENCES assessment_manager.t_assessment( a_id );

ALTER TABLE assessment_manager.toc_assessment_interviewer ADD CONSTRAINT fk_toc_assessment_interviewer_t_interviewer FOREIGN KEY ( a_i_interviewer_id ) REFERENCES assessment_manager.t_interviewer( i_id );
