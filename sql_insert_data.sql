INSERT INTO public.t_user(
	u_id, u_login, u_password)
	VALUES (1, 'user', 'pass');

INSERT INTO public.t_assessment_status(
	a_s_id, a_s_name, a_s_fk)
	VALUES
		(1, 'Назначен', null),
        (2, 'Проведён', 1),
        (3, 'Отменён', 1);

INSERT INTO public.t_candidate_status(
    c_s_id, c_s_name, c_s_fk)
    VALUES
        (1, 'Приглашён', null),
        (2, 'Не явился', 1),
        (3, 'Завершил', 1),
        (4, 'Не завершил', 1),
        (5, 'Принят на обучение', 3),
        (6, 'Принят на работу', 3),
        (7, 'Не принят', 3);