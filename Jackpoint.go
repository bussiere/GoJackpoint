package main

import (
	"./structure"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func HelloServer(c http.ResponseWriter, req *http.Request) {

	result := "<html><body><form action='/login/' method='post'><table><tr><td><label for='login'><strong>Nom de compte</strong></label></td><td><input type='text' name='login' id='login'/></td></tr><tr><td><label for='pass'><strong>Mot de passe</strong></label></td><td><input type='password' name='pass' id='pass'/></td></tr></table><input type='submit' name='connexion' value='Se connecter'/></form></body></html>"
	c.Header().Set("Content-Type", "text/html")
	c.Header().Set("Content-Length", strconv.Itoa(len(result)))
	io.WriteString(c, result)
}

func login(c http.ResponseWriter, req *http.Request) {
	body := r.FormValue("login")
	fmt.Printf(body)
	result = ""
	c.Header().Set("Content-Type", "text/html")
	c.Header().Set("Content-Length", strconv.Itoa(len(result)))

}

func init() {

}

func main() {
	j := new(struct_jackpoint.Jack)
	j.Id = 2
	fmt.Printf("%d", j.Id)
	http.HandleFunc("/login/", editHandler)

	http.Handle("/", http.HandlerFunc(HelloServer))
	http.ListenAndServe(":8050", nil)
}
