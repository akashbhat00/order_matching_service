package db

import (
	"database/sql"
	"fmt"
	"os"

	"bitbucket.org/liamstask/goose/lib/goose"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)


var gormDB *gorm.DB
var err error

// DB ...
type DB struct{}

func Init() {
	dbUserName := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUserName, dbName, dbPassword) //Build connection string

	pqDB, err := sql.Open("postgres", dbURI)

	if err != nil {
		fmt.Println("Failed to connect to DB", dbURI, err.Error())
		os.Exit(1)
	}
	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: pqDB,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("Failed to connect to DB", dbURI, err.Error())
		os.Exit(1)
	}
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Not able to fetch the working directory")
		os.Exit(1)
	}

	workingDir = workingDir + "/db/migrations"
	migrateConf := &goose.DBConf{
		MigrationsDir: workingDir,
		Driver: goose.DBDriver{
			Name:    "postgres",
			OpenStr: dbURI,
			Import:  "github.com/lib/pq",
			Dialect: &goose.PostgresDialect{},
		},
	}
	fmt.Println("Fetching the most recent DB version")
	latest, err := goose.GetMostRecentDBVersion(migrateConf.MigrationsDir)
	if err != nil {
		fmt.Println("Unable to get recent goose db version", err)

	}
	fmt.Println(" Most recent DB version ", latest)
	fmt.Println("Running the migrations on db", workingDir)
	err = goose.RunMigrationsOnDb(migrateConf, migrateConf.MigrationsDir, latest, pqDB)
	if err != nil {
		fmt.Println("Error while running migrations", err)
		os.Exit(1)
	}
}

// GetDB : Get an instance of DB to connect to the database connection pool
func (d DB) GetDB() *gorm.DB {
	return gormDB
}
