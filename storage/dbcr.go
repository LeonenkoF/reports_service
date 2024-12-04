package storage

// запросы создания таблиц сервица
var (
	schTypes = `
	CREATE TYPE PRIORITY_TYPE AS ENUM ('high','medium','low');
	CREATE TYPE STAGE_TYPE AS ENUM ('new','inprogress','done', 'canceled');
	CREATE TYPE ROLE_TYPE AS ENUM ('admin','user','superadmin','superuser','guest');
`
	//
	schReport string = `
CREATE TABLE IF NOT EXISTS reports (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	uuid UUID,
	user_uuid UUID,
	description TEXT,
	priority PRIORITY_TYPE,
	stage STAGE_TYPE,
	category INT,
	groupp UUID,
	groupp_name VARCHAR(32),
	adm_uuid UUID,
	closing_comment TEXT,
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);
CREATE INDEX IF NOT EXISTS idx_created_at ON reports (created_at);
`
	// таблица группы (под каждую группу своя таблица с уникальным именем)
	// !!! %s имя уникально НЕ groupp !!!
	schGroupp string = `
CREATE TABLE IF NOT EXISTS groupp (
	rep_id INT
)
`
	// или таблица групп (под все группы одна таблица)
	// *** недостаток - дублирование данных. преимущество - предположительно более легкий поиск
	schGroupps string = `
CREATE TABLE IF NOT EXISTS groupps (
	groupp UUID,
	group_name VARCHAR(32),
	rep_id INT
);
CREATE INDEX IF NOT EXISTS idx_id ON groupps (group_name);
`

	// таблица диалога по жалобе. каждому диалогу своя таблица
	/*
			id 			идентификатор
			user_uuid 	автор
			comment 	собтвенно текст
		    created_at 	дата создания
		    updated_at 	дата редакции *** возможно лишнее поле
	*/
	// !!! %s имя уникально НЕ comments !!!
	schComments string = `
CREATE TABLE IF NOT EXISTS comments (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_uuid UUID,
	comment TEXT,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,	
    updated_at DATE NOT NULL DEFAULT CURRENT_DATE	
);
`
	schUsers string = `
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
`
	// список тем *** возможно нужно просто enum, но возможное появление новых тем будет вызывать переписывание кода
	schCategory string = `
CREATE TABLE IF NOT EXISTS categoryes (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	category VARCHAR(32)
);	
`
)
