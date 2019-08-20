package main

import (
	"fmt"
	"log"
	"net/http"

	"text/template"

	"HotelWebsite/database"

	"github.com/gorilla/mux"
)

func main() {

	r := NewRouter()
	r.HandleFunc("/Hotel-Website", hotelhandler)
	r.HandleFunc("/Hotel-Website/About", abouthandler)
	r.HandleFunc("/Hotel-Website/Contact", contacthandler).Methods("POST", "GET")
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)

}

func hotelhandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("C:/Users/yashi/go/src/static-site/index.html")
	if err != nil {
		log.Fatal("Could not parse template files\n")
	}

	er := t.Execute(w, "")
	if er != nil {
		log.Fatal("could not execute the files\n")
	}

}
func abouthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yahaan aagya")
	t, err := template.ParseFiles("C:/Users/yashi/go/src/static-site/about.html")
	if err != nil {
		log.Fatal("Could not parse template files\n")
	}

	er := t.Execute(w, "")
	if er != nil {
		log.Fatal("could not execute the files\n")
	}

}

func contacthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yahaan aagya")
	clctn, client := database.Createdb()
	switch r.Method {

	case "GET":
		{
			fmt.Println("yeh chlra hai")
			t, err := template.ParseFiles("C:/Users/yashi/go/src/static-site/contact.html")
			if err != nil {
				log.Fatal("Could not parse template files\n")
			}
			er := t.Execute(w, "")
			if er != nil {
				log.Fatal("could not execute the files\n")
			}
		}
		log.Print("working")
	case "POST":
		{
			fmt.Println(" lets see if it works ")
			a := r.FormValue("First name")

			b := r.FormValue("Last name")
			c := r.FormValue("Email")
			d := r.FormValue("Message")
			db1 := &database.Data{
				Fname: a,
				Lname: b,
				Email: c,
				Msg:   d,
			}
			fmt.Println(db1)
			//email := db1.Email
			database.Insertintodb(clctn, db1)

			http.Redirect(w, r, "./Contact", 302)

		}

		database.Disconnectdb(client)
	}
}

//NewRouter .....
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return r
}
