package main

import (
	"customerservice/internal/entity"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"
)

///

func dbConnect(dsn string) (*gorm.DB, error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold ...
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func init() {

	godotenv.Load()

}

//Some modify
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address") //it returns a pointer
	f, err := os.OpenFile("../../tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/loanapp?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	fmt.Println(dsn)
	db, err := dbConnect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &App{
		ErrorLog:         errorLog,
		InfoLog:          infoLog,
		MessageAuthModel: &entity.MessageAuthModel{Db: db},
		LoanApp:          &entity.LoanApplicationModel{Db: db},
		BackendUsers:     &entity.BackendUsersModel{Db: db},
		UserContactModel: &entity.UserContactModel{Db: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.registerroutes(),
	}

	infoLog.Printf("Starting server on %s", *addr) // dereferencing the pointer
	// Call the ListenAndServe() method on our new http.Server struct.
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
