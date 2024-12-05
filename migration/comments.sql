/*
	указана только структура таблицы. 
	Таблиц будет столько сколько обсуждений будет открыто,
	но не больше чем суммарное количетско жалоб. 
	К имени будем добавлять либо номер либо uuid жалобы.
	По истечении года талица будет удаляться.
*/
CREATE TABLE IF NOT EXISTS comments (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_uuid UUID,
	comment TEXT,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,	
    updated_at DATE NOT NULL DEFAULT CURRENT_DATE	
);