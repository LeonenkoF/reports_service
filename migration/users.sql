	/*
	// таблица пользователей
	   	id 			идентификатор
	   	user_uuid 	UUID,
	   	mail 		почта
	   	phone 		номер телефона (*** не обязательно)
	   	nick 		ник (*** необязательно)
	   	role INT	роль (user admin *** superuser *** superadmin)
	       created_at 	дата регистрации
	       updated_at 	дата инзменения (*** не обязательно)
	*/

CREATE TABLE IF NOT EXISTS users (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_uuid UUID,
	mail VARCHAR(64),
	phone VARCHAR(16),
	nick VARCHAR(16),
	role ROLE_TYPE,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,	
    updated_at DATE NOT NULL DEFAULT CURRENT_DATE	
);
