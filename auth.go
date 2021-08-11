package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

// Credentials holds user credential data: Username and Password
type Credentials struct {
	Username string
	Password []byte
	Email    string
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func getUserID(userName string) (userID int) {
	result := db.QueryRow("SELECT user_id FROM users WHERE username=$1", userName)
	err := result.Scan(&userID)

	if err != nil {
		fmt.Printf("ERROR getUserID(%s): %s\n", userName, err)
		return
	}

	return userID
}

func getUserNameFromID(userID int) (userName string) {
	result := db.QueryRow("SELECT username FROM users WHERE user_id=$1", userID)
	err := result.Scan(&userName)

	if err != nil {
		fmt.Printf("ERROR getUserNameFromID(%d): %s\n", userID, err)
		return
	}

	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func loginHandler(response http.ResponseWriter, request *http.Request) {
	if getUserName(request) != "" {
		http.Redirect(response, request, "/", 302)
	}

	// Read cookie to handle notification data
	cookie, err := request.Cookie("notify")
	notifyData := ""
	if err != nil {
		fmt.Printf("\nERROR loginHandler: reading cookie: %s", err)
	} else {
		notifyData = cookie.Value
	}

	type LoginPageData struct {
		Notification string
	}

	switch notifyData {
	case "usernotexists":
		notifyData = "No such user exists"
	case "notprovided":
		notifyData = "Both e-mail and password are needed!"
	case "notmatch":
		notifyData = "Wrong password!"
	default:
		notifyData = ""
	}
	data := LoginPageData{
		Notification: notifyData,
	}

	// Remove the cookie
	notifyCookie := &http.Cookie{
		Name:   "notify",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, notifyCookie)

	// Execute template
	err = tmpl[tmplLogin].Execute(response, data)

	if err != nil {
		return
	}
}

// LogoutHandler clears session cookies and redirects user to homepage
func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

func signupHandler(response http.ResponseWriter, request *http.Request) {
	if getUserName(request) != "" {
		http.Redirect(response, request, "/", 302)
	}

	// Read cookie to handle notification data
	cookie, err := request.Cookie("notify")
	notifyData := ""
	if err != nil {
		fmt.Printf("\nERROR loginHandler: reading cookie: %s", err)
	} else {
		notifyData = cookie.Value
	}

	type LoginPageData struct {
		Notification string
	}

	switch notifyData {
	case "exists":
		notifyData = "User already exists!"
	case "empty":
		notifyData = "All fields are needed!"
	case "email":
		notifyData = "Please provide real e-mail address!"
	default:
		notifyData = ""
	}
	data := LoginPageData{
		Notification: notifyData,
	}

	// Remove the cookie
	notifyCookie := &http.Cookie{
		Name:   "notify",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, notifyCookie)

	// Execute template
	err = tmpl[tmplSignup].Execute(response, data)

	if err != nil {
		return
	}
}

func postLoginHandler(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	pass := request.FormValue("passwd")
	redirectTarget := urlLogin

	if email != "" && pass != "" {
		var err error

		inputPassword := request.FormValue("passwd")

		u := Credentials{
			Email:    request.FormValue("email"),
			Password: nil,
		}

		result := db.QueryRow("SELECT username, password FROM users WHERE email=$1", u.Email)

		if err != nil {
			// If there is an issue with the database, return a 500 error
			response.WriteHeader(http.StatusInternalServerError)
			return
		}

		storedCreds := &Credentials{}

		err = result.Scan(&u.Username, &storedCreds.Password)
		if err != nil {
			// If an entry with the username does not exist;
			if err == sql.ErrNoRows {
				// ..set notification cookie
				notifyCookie := &http.Cookie{
					Name:  "notify",
					Value: "usernotexists",
					Path:  "/",
				}
				http.SetCookie(response, notifyCookie)

				// ..send an "Unauthorized"(401) status; OR NOT
				//response.WriteHeader(http.StatusUnauthorized)

				// ..redirect
				http.Redirect(response, request, urlLogin, 302)
				return
			}
			// If the error is of any other type, send a 500 status
			response.WriteHeader(http.StatusInternalServerError)
			http.Redirect(response, request, urlLogin, 302)
			return
		}

		// Compare the stored hashed password, with the hashed version of the password that was received
		if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(inputPassword)); err != nil {
			// If the two passwords DO NOT MATCH:
			// ..set notification cookie
			notifyCookie := &http.Cookie{
				Name:  "notify",
				Value: "notmatch",
				Path:  "/",
			}
			http.SetCookie(response, notifyCookie)

			// ..return a 401 status; OR DON'T
			//response.WriteHeader(http.StatusUnauthorized)

			// ..set redirect target to login page
			redirectTarget = urlLogin
		} else {
			// If passwords MATCH; set session cookie and send user to homepage
			setSession(u.Username, response)
			redirectTarget = "/"

			// Update user's last login date
			if _, err = db.Query("UPDATE users SET last_login = $1 WHERE email = $2", getDate(), email); err != nil {
				response.WriteHeader(http.StatusInternalServerError)
				fmt.Printf("ERROR postLoginHandler: %s", err)
				return
			}
		}
	} else {
		// If e-mail or password not provided;
		// ..set notification cookie
		notifyCookie := &http.Cookie{
			Name:  "notify",
			Value: "notprovided",
			Path:  "/",
		}
		http.SetCookie(response, notifyCookie)
	}
	http.Redirect(response, request, redirectTarget, 302)
}

func postSignupHandler(response http.ResponseWriter, request *http.Request) {
	// Check request method first
	if request.Method != "POST" {
		fmt.Printf("ERROR postSignupHandler MethodCheck")
		return
	}

	// Check if any field is empty;
	if request.FormValue("name") == "" || request.FormValue("passwd") == "" || request.FormValue("email") == "" {
		// ..set notification cookie
		notifyCookie := &http.Cookie{
			Name:  "notify",
			Value: "empty",
			Path:  "/",
		}
		http.SetCookie(response, notifyCookie)

		// ..log the error
		fmt.Printf("\nERROR postSignupHandler ValueChecks")

		// ..redirect
		http.Redirect(response, request, urlSignup, 302)
		return
	}

	// Check if the e-mail is valid or not;
	if isEmailValid(request.FormValue("email")) != true {
		// ..set notification cookie
		notifyCookie := &http.Cookie{
			Name:  "notify",
			Value: "email",
			Path:  "/",
		}
		http.SetCookie(response, notifyCookie)

		// ..log the error
		fmt.Printf("\nERROR postSignupHandler EmailCheck(%s)", request.FormValue("email"))

		// ..redirect
		http.Redirect(response, request, urlSignup, 302)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.FormValue("passwd")), hashCost)

	if err != nil {
		fmt.Printf("\nERROR postSignupHandler hashedPassword: %s", err)
		return
	}

	// Create data to hold them all for now
	u := Credentials{
		Username: request.FormValue("name"),
		Password: hashedPassword,
		Email:    request.FormValue("email"),
	}

	// Insert data into database
	if _, err = db.Query("INSERT INTO users (username,password,email,register_date,blocked) VALUES ($1,$2,$3,$4,$5)", u.Username, string(u.Password), u.Email, getDate(), false); err != nil {
		// ..set notification cookie
		notifyCookie := &http.Cookie{
			Name:  "notify",
			Value: "exists",
			Path:  "/",
		}
		http.SetCookie(response, notifyCookie)

		// ..log the error
		fmt.Printf("\nERROR postSignupHandler intoDatabase: %s", err)

		// ..redirect
		http.Redirect(response, request, urlSignup, 302)
		return
	}

	// Redirect user to login page
	http.Redirect(response, request, urlLogin, 302)
}
