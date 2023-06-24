package db

import (
	// "context"
	// "database/sql"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	// "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	// "google.golang.org/appengine/log"
	log "github.com/sirupsen/logrus"
)

type DBConfig struct {
	DbUserName string
	DbName     string
	DbPassword string
	DbPort     string
	DbHost     string
}

type DbHandler struct {
	DBMigrate *migrate.Migrate
	SqlDB     *sql.DB
}

func DbInit(conn DBConfig) (DbHandler, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)
	logger.SetOutput(os.Stdout)

	dbUrl := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", conn.DbUserName, conn.DbPassword, conn.DbPort, conn.DbName)

	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		logger.Errorf("not open postgres %+v", err)
		return DbHandler{}, err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Errorf(" postgres instance issue %+v", err)
		return DbHandler{}, err

	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./pkg/migrations", dbUrl, driver)

	if err != nil {
		logger.Errorf(" New With Database Instance instance issue %+v", err)
		return DbHandler{}, err
	}

	//m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

	// Apply the database migrations
	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		logger.Errorf(" migration issue issue %+v", err)
		return DbHandler{}, err
	}

	logger.Infoln("Database migrations applied successfully!")

	defer m.Close()

	return DbHandler{
		DBMigrate: m,
		SqlDB:     db,
	}, nil
}

func NewDatabaseConfig(config DBConfig) (string, error) {
	var dbParams []string
	dbParams = append(dbParams, fmt.Sprintf("user=%s", config.DbUserName))
	if host := config.DbHost; host == "" {
		dbParams = append(dbParams, fmt.Sprintf("host=%s", "localhost"))
	} else {
		dbParams = append(dbParams, fmt.Sprintf("host=%s", config.DbHost))
	}
	dbParams = append(dbParams, fmt.Sprintf("port=%s", config.DbPort))
	dbParams = append(dbParams, fmt.Sprintf("dbname=%s", config.DbName))
	if password := config.DbPassword; password != "" {
		dbParams = append(dbParams, fmt.Sprintf("password=%s", password))
	}
	// dbParams = append(dbParams, fmt.Sprintf("sslmode=%s",
	// 	config.GetString("database.sslMode")))
	// dbParams = append(dbParams, fmt.Sprintf("connect_timeout=%d",
	// 	config.GetInt("database.connectionTimeout")))
	// dbParams = append(dbParams, fmt.Sprintf("statement_timeout=%d",
	// 	config.GetInt("database.statementTimeout")))
	// dbParams = append(dbParams, fmt.Sprintf("idle_in_transaction_session_timeout=%d",
	// 	config.GetInt("database.idleInTransactionSessionTimeout")))

	return strings.Join(dbParams, " "), nil
}
