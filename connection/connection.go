package connection

import (
	"auto_sql/settings"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectWithDatabase() (*sql.DB, error) {
	databaseCredentials := settings.GetCredentials()

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		databaseCredentials.Host, databaseCredentials.Port, databaseCredentials.User,
		databaseCredentials.Password, databaseCredentials.Database)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalln("Error opening database")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("Fail to connect into database ", err)
		return nil, err
	}

	log.Printf("Connected with %s database!", databaseCredentials.Database)
	return db, nil
}
