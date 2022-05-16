package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Rodrigueslcs/challenge-backend/api/handler"
	"github.com/Rodrigueslcs/challenge-backend/config"
	"github.com/Rodrigueslcs/challenge-backend/infrastructure/repository"
	"github.com/Rodrigueslcs/challenge-backend/usecase/video"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	videoRepo := repository.NewVideoMySQL(db)
	videoService := video.NewService(videoRepo)

	n := negroni.Classic()
	r := mux.NewRouter()

	handler.MakeBookHandlers(r, videoService)

	n.UseHandler(r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ": 8080",
		ErrorLog:     logger,
		Handler:      n,
	}
	fmt.Println("API RODANDO NA PORTA:", server.Addr)
	log.Fatal(server.ListenAndServe())

}

// var maxDigit int
// 	fmt.Scanf("%d", &maxDigit)

// 	var maxString string
// 	var maxInt int
// 	for i := 0; i < 4; i++ {
// 		maxString += strconv.Itoa(maxDigit)
// 	}
// 	maxInt, _ = strconv.Atoi(maxString)

// 	for i := 1000; i < maxInt; i++ {
// 		sun := 0
// 		for _, v := range strconv.Itoa(i) {
// 			digit, _ := strconv.Atoi(string(v))
// 			if digit <= maxDigit {
// 				sun += digit
// 			}
// 		}
// 		if sun == 21 {
// 			cont = true
// 			fmt.Println(i)
// 		}
// 		sun = 0
// 	}

// 	if !cont {
// 		fmt.Println("nulo")
// 	}
