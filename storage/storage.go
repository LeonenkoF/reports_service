package storage

// заготовка ReportStore
import (
	"context"
	"fmt"
	"time"

	//_ "modernc.org/sqlite"
	// _ "github.com/lib/pq" // Импорт драйвера
	"github.com/go-pg/pg/v10"
)

type Report struct {
	ID          int
	uuid        string
	user_uuid   string
	description string
	created_at  time.Time
	update_at   time.Time
}

const (
	template_date = "20060102"
	limit         = 50
)

var (
	err     error
	reports []Report
)

type ReportStore struct {
	DB *pg.DB
}

func NewReportStore(db *pg.DB) ReportStore {
	return ReportStore{DB: db}
}

func (ts ReportStore) Add(report Report) error {
	_, err = ts.DB.Model(&report).Insert()
	return err
}

func (ts ReportStore) Delete(id int) error {
	_, err = ts.DB.Model(&reports).Where("id = ?", id).Delete()
	return err
}

func (ts ReportStore) Find(search string) ([]Report, error) {
	if len(search) == 0 {
		// возвращаем всё если строка пустая
		err = ts.DB.Model(&reports).Limit(limit).Order("creat_at").Order("description").Select()
		return reports, err
	} else {
		search = "%" + search + "%"
		// _, err = ts.DB.Query(&reports, "select * from reports where upper(description) like upper(?)", search).Order("creat_at").Order("description")
		return reports, err
	}
}

func (ts ReportStore) Get(id int) (Report, error) {
	if err = ts.DB.Model(&reports).Where("id = ?", id).Select(); err != nil {
		return Report{}, err
	}
	return reports[0], err
}

func (ts ReportStore) Update(report Report) error {
	// так работает
	_, err = ts.DB.Model(&report).WherePK().Update()
	return err
}

func InitDBase() (*pg.DB, error) {
	fmt.Println("Init Data Base...")
	// Создание конфигурации для подключения к базе данных
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432", // Адрес PostgreSQL сервера
		User:     "postgres",       // Имя пользователя
		Password: "password",       // Пароль
		Database: "testdb",         // Имя базы данных
	})
	// создание таблицы
	//	if _, err := db.Exec(schTypes + schReport + schGroups + schUsers + schCategory); err != nil {
	if _, err := db.Exec(schTypes + schReport + schGroups + schUsers + schCategory); err != nil {
		return nil, err
	}
	// Проверка соединения
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		// ошибка подключения к базе
		return nil, err
	}
	fmt.Println("База подключена (Ping Ok)")
	return db, err
}
