/*
	Типы перечислений проекта
	*** ROLE_TYPE намеренно расширен дополнительными ролями. Вопрос на обсуждении.
*/
	CREATE TYPE PRIORITY_TYPE AS ENUM ('high','medium','low');
	CREATE TYPE STAGE_TYPE AS ENUM ('new','inprogress','done', 'canceled');
	CREATE TYPE ROLE_TYPE AS ENUM ('admin','user','superadmin','superuser','guest');
    