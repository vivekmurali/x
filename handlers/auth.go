package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty"
)

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("template/index.html")
	tmpl.Execute(w, nil)
}

func Callback(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("template/callback.html")
	tmpl.Execute(w, nil)
}

func PrintUserDetails(w http.ResponseWriter, r *http.Request) {

	sessionContainer := session.GetSessionFromRequestContext(r.Context())
	userID := sessionContainer.GetUserID()
	userInfo, err := thirdparty.GetUserByID(userID)
	if err != nil {
		log.Fatalf("Error printing user details, %v", err)
	}
	fmt.Println("userInfo", userInfo)

	fmt.Fprintf(w, "UserInfo is %+w", userInfo)
}
