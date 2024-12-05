package storage

// заготовка ReportStore
import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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
	// создание таблиц
	// types
	if err := loadSQLFile(db, "migration/types.sql"); err != nil {
		fmt.Printf("Create types error %v\n", err)
	} else {
		fmt.Println("Create type ok")
	}
	// reports
	if err := loadSQLFile(db, "migration/reports.sql"); err != nil {
		log.Fatalf("Create reports error %v\n", err)
		return nil, err
	}
	fmt.Println("Reports ok")
	// groups
	if err := loadSQLFile(db, "migration/groupps.sql"); err != nil {
		log.Fatalf("Create groupps error %v\n", err)
		return nil, err
	}
	fmt.Println("Groupps ok")
	// users
	if err := loadSQLFile(db, "migration/users.sql"); err != nil {
		log.Fatalf("Create users error %v\n", err)
		return nil, err
	}
	fmt.Println("Users ok")
	// Проверка соединения
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		// ошибка подключения к базе
		return nil, err
	}
	fmt.Println("База подключена (Ping Ok)")
	return db, err
}

func loadSQLFile(db *pg.DB, sqlFile string) error {
	file, err := os.ReadFile(sqlFile)
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		tx.Rollback()
	}()
	for _, q := range strings.Split(string(file), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		if _, err := tx.Exec(q); err != nil {
			return err
		}
	}
	return tx.Commit()
}
