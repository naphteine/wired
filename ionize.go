package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Program details
const (
	Program   = "ionize"
	Version   = "v1.0"
	Copyright = "All rights reserved. (c) 2021"
	Host      = ""
	Port      = 8080
	hashCost  = 12

	dbHost   = "localhost"
	dbPort   = 5432
	dbUser   = "root"
	dbPasswd = "password"
	dbName   = "duga1"

	urlLogin      = "/login"
	urlSignup     = "/sign-up"
	urlLogout     = "/logout"
	urlUser       = "/u"
	urlPostLogin  = "/post/login"
	urlPostSignup = "/post/signup"

	tmplParts  = "templates/parts.html"
	tmplIndex  = "templates/index.html"
	tmplLogin  = "templates/login.html"
	tmplSignup = "templates/signup.html"
	tmplUser   = "templates/user.html"
)

var router = mux.NewRouter()
var tmpl = make(map[string]*template.Template)
var db *sql.DB

func getDate() string {
	current := time.Now().UTC()
	return current.Format("2006-01-02 15:04:05 -0700")
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	// Prepare index data
	type IndexData struct {
		Username    string
		ProfileLink string
	}

	data := &IndexData{
		Username:    getUserName(request),
		ProfileLink: "/u/" + getUserName(request),
	}

	// Run
	err := tmpl[tmplIndex].Execute(response, data)

	if err != nil {
		return
	}
}

func userHandler(response http.ResponseWriter, request *http.Request) {
	if getUserName(request) == "" {
		http.Redirect(response, request, "/", 302)
		return
	}

	vars := mux.Vars(request)
	userName := vars["user"]

	var userID int
	var userEmail sql.NullString
	var userRegDate sql.NullString
	var userLastLogin sql.NullString

	result := db.QueryRow("SELECT user_id, email, register_date, last_login FROM users WHERE username=$1", userName)
	err := result.Scan(&userID, &userEmail, &userRegDate, &userLastLogin)

	if err != nil {
		fmt.Printf("ERROR userHandler: %s\n", err)
		return
	}

	type UserData struct {
		ID           int
		Title        string
		ProfileLink  string
		Username     string
		Name         string
		Email        string
		RegisterDate string
		LastLogin    string
	}

	data := &UserData{
		ID:           userID,
		Title:        "Ionize - " + userName,
		ProfileLink:  "/u/" + getUserName(request),
		Username:     getUserName(request),
		Name:         userName,
		Email:        userEmail.String,
		RegisterDate: userRegDate.String,
		LastLogin:    userLastLogin.String,
	}

	err = tmpl[tmplUser].Execute(response, data)

	if err != nil {
		fmt.Printf("ERROR userHandler: %s\n", err)
		return
	}
}

func main() {
	// Get flags
	portPtr := flag.Int("port", Port, "HTTP server port")
	flag.Parse()

	// Print details at start
	fmt.Printf("%s :: %s\nStarting up...\n", Copyright, Program)
	fmt.Printf("Version:\t%s\n", Version)
	fmt.Printf("Date:\t\t%s\n", getDate())
	fmt.Printf("Port:\t\t%s\n", strconv.Itoa(*portPtr))

	// Init DB connection
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPasswd, dbName)

	if db, err = sql.Open("postgres", psqlconn); err != nil {
		panic(err)
	}

	// Parse the template files
	tmpl[tmplIndex] = template.Must(template.ParseFiles(tmplIndex, tmplParts))
	tmpl[tmplLogin] = template.Must(template.ParseFiles(tmplLogin, tmplParts))
	tmpl[tmplSignup] = template.Must(template.ParseFiles(tmplSignup, tmplParts))
	tmpl[tmplUser] = template.Must(template.ParseFiles(tmplUser, tmplParts))

	// Handle pages
	router.HandleFunc("/", indexHandler)
	router.HandleFunc(urlLogin, loginHandler)
	router.HandleFunc(urlSignup, signupHandler)
	router.HandleFunc(urlLogout, logoutHandler)

	router.HandleFunc(urlPostLogin, postLoginHandler).Methods("POST")
	router.HandleFunc(urlPostSignup, postSignupHandler).Methods("POST")

	router.HandleFunc(urlUser+"/{user}", userHandler)

	// File server
	router.PathPrefix("/res/").Handler(http.StripPrefix("/res/", http.FileServer(http.Dir("static"))))

	// Start server
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(Host+":"+strconv.Itoa(*portPtr), nil))

	// Shutting down
	db.Close()
}
