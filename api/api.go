package main

import (
	"database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/julienschmidt/httprouter"
)
const CREDIT_LOWER_LIMIT = 500;
const CREDIT_SCORE_STANDART_LIMIT= 1000;
const CREDIT_LIMIT_SCORE = 4;
func main() {
	defer database.DB.Close()

	// add router and routes
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.POST("/creditApprovement", customerApproveHandler)
	router.OPTIONS("/*any", corsHandler)

	// database oluştur
	_, err := database.Init()
	if err != nil {
		log.Println("connection to DB failed, aborting...")
		log.Fatal(err)
	}

	log.Println("connected to DB")

	// ekran modu log kaydı kontrolü
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	http.ListenAndServe(":8080", router)
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
	fmt.Fprintf(w, "This is the RESTful api") //api kontrol ediliyor
}


func customerApproveHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
	responseBody := ApprovementResponse{};

	decoder := json.NewDecoder(r.Body)
	var body ApproveServiceRequestBody{}
	if err := decoder.Decode(&body); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	respScore := CustomerScore{}

	identity := strconv.FormatInt(body.Identity, 10)
	respScore = getCustomerScore(identity)
	
	if(respScore.CreditScore < CREDIT_SCORE_LOWER_LIMIT){
		responseBody.ApprovementStatus = false;
		responseBody.AssignedCreditAmount= 0; //bu alan omitempty de geçilebilir
	}

	if(respScore.CreditScore > CREDIT_SCORE_LOWER_LIMIT && 
	   respScore.CreditScore < CREDIT_SCORE_STANDART_LIMIT &&
	   body.Salary < 5000 ){
		responseBody.ApprovementStatus = true;
		responseBody.AssignedCreditAmount= 10000;
	}

	if(respScore.CreditScore >= CREDIT_SCORE_STANDART_LIMIT){
		 responseBody.ApprovementStatus = true;
		 responseBody.AssignedCreditAmount= CREDIT_LIMIT_SCORE * body.Salary;
	 }

	//Tüm işlemler başarılı devam ettiğinde ilgili kaydı at
	database.DB.Create(&body);

	w.Write(responseBody)
}

func getCustomerScore(identity string) (customerScoreTableData CustomerScore){
	customerScoreTableData := CustomerScore{}
	database.DB.Where("Identity = ?", identity).First(&customerScoreTableData)
	
	return customerScoreTableData
}

// tanımsız url lerde
func corsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
}


// util
func getFrontendURL() string {
	if os.Getenv("APP_ENV") == "production" {
		return "http://localhost:3000" // prod ortamda url ekleyebilirsin
	} else {
		return "http://localhost:3000"
	}
}

func setCors(w http.ResponseWriter) {
	frontendURL := getFrontendURL()
	w.Header().Set("Access-Control-Allow-Origin", frontendURL)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

type CustomerScore struct {
	Identity    int64 `json:"identity"`
	CreditScore int64 `json:"creditScore"`
}

type ApproveServiceRequestBody struct {
	Identity    int64 	`json:"identity"`
	FirstName   string 	`json:"firstName"`
	LastName  	string 	`json:"lastName"`
	Salary   	float64 `json:"salary"`
	Number    	int64 	`json:"number"`
}

type ApprovementResponse struct {
	ApprovementStatus bool `json:"approvementStatus"`
	AssignedCreditAmount float64 `json:"assignedCreditAmount"`
}