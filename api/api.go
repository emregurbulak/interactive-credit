package main

import (
	"database"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	defer database.DB.Close()

	// add router and routes
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/creditApprovement", customerApproveHandler)
	router.OPTIONS("/*any", corsHandler)

	// add database
	_, err := database.Init()
	if err != nil {
		log.Println("connection to DB failed, aborting...")
		log.Fatal(err)
	}

	log.Println("connected to DB")

	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	http.ListenAndServe(":8080", router)
}

func showPostHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	var post database.Post
	database.DB.Where("ID = ?", ps.ByName("postId")).First(&post)
	res, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(res)
}

func indexPostHandler(w http.ResponseWriter, identity string, _ httprouter.Params) (CustomerScore, error) {
	customerScoreTableData := CustomerScore{}
	setCors(w)
	var posts []database.Post
	database.DB.Where("Identity = ?", identity).First(&customerScoreTableData)

	res, err := json.Marshal(customerScoreTableData)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func customerApproveHandler() {

}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
	fmt.Fprintf(w, "This is the RESTful api") //api kontrol ediliyor
}

// tanımsız url lerde
func corsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
}

func getCustomerScore(w http.ResponseWriter, identity string, _ httprouter.Params) (CustomerScore, error) {
	customerScoreTableData := CustomerScore{}
	setCors(w)
	var posts []database.Post
	database.DB.Where("Identity = ?", identity).First(&customerScoreTableData)

	res, err := json.Marshal(customerScoreTableData)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// util
func getFrontendUrl() string {
	if os.Getenv("APP_ENV") == "production" {
		return "http://localhost:3000" // prod ortamda url ekleyebilirsin
	} else {
		return "http://localhost:3000"
	}
}

func setCors(w http.ResponseWriter) {
	frontendUrl := getFrontendUrl()
	w.Header().Set("Access-Control-Allow-Origin", frontendUrl)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

type CustomerScore struct {
	Identity    int64 `json:"identity"`
	CreditScore int64 `json:"creditScore"`
}
