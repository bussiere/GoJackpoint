package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
)

type Jack struct {
	Id                    rune
	Created               int64
	Updated               int64
	Nom                   string
	Email                 string
	Skill_Jack_Id         []rune
	Carac_Jack_Id         []rune
	Item_Jack_Id          []rune
	Skill_Jack_Private_Id []rune
	Item_Jack_Private_Id  []rune
	Password              string
	Key_public            string
	Key_private           string
	Message_Id            []rune
	Statut                string
	Avatar                string
}

type Configuration struct {
	ServeurSmtp []rune
	ServeurPop  []rune
}

type Configuration_Admin struct {
	Email      string
	Key_public string
}

type Tag struct {
	Id  rune
	Nom string
}

type ServeurSmtp struct {
	Id      rune
	Email   rune
	Adresse rune
	Login   rune
	MDP     rune
	Created int64
	Updated int64
}

type ServeurPop struct {
	Id      rune
	Email   rune
	Adresse rune
	Login   rune
	MDP     rune
	Created int64
	Updated int64
}

type Admin struct {
	Id          rune
	Key_public  string
	Key_private string
	Jack_Id     rune
	Email       rune
	Created     int64
	Updated     int64
}

type Hand struct {
	Id            rune
	Created       int64
	Updated       int64
	Skill_Jack_Id []rune
	Carac_Jack_Id []rune
	Item_Jack_Id  []rune
	Message       string
	Tag_Id        []rune
}

type Skill struct {
	Id          rune
	Created     int64
	Updated     int64
	Nom         string
	Description string
	Tag_Id      []rune
}

type Filiation_Skill struct {
	Id              rune
	Created         int64
	Updated         int64
	Parent_Skill_Id rune
	Enfant_Skill_Id rune
}

type Carac struct {
	Id          rune
	Created     int64
	Updated     int64
	Nom         string
	Description string
	Tag_Id      []rune
}

type Item struct {
	Id          rune
	Created     int64
	Updated     int64
	Nom         string
	Description string
	Tag_Id      []rune
}

type Item_Carac struct {
	Id       rune
	Created  int64
	Updated  int64
	Carac_Id rune
	Item_Id  rune
}

type Item_Skill struct {
	Id       rune
	Created  int64
	Updated  int64
	Skill_Id rune
	Item_Id  rune
}

type Skill_Jack struct {
	Id       rune
	Created  int64
	Updated  int64
	Skill_Id rune
	Jack_Id  rune
}
type Item_Jack struct {
	Id      rune
	Created int64
	Updated int64
	Item_Id rune
	Jack_Id rune
}
type Carac_Jack struct {
	Id      rune
	Created int64
	Updated int64
	Item_Id rune
	Jack_Id rune
}

type Skill_Jack_Private struct {
	Id       rune
	Created  int64
	Updated  int64
	Skill_Id rune
	Jack_Id  rune
}
type Item_Jack_Private struct {
	Id      rune
	Created int64
	Updated int64
	Item_Id rune
	Jack_Id rune
}

type Admin_Private struct {
	Id                    rune
	Admin_Id              rune
	Id_Item_Jack_Private  Item_Jack_Private
	Id_Skill_Jack_Private Skill_Jack_Private
	Created               int64
	Updated               int64
}

func IndexPage(c http.ResponseWriter, req *http.Request) {

	result := "<html><body><form action='/login/' method='post'><table><tr><td><label for='login'><strong>Nom de compte</strong></label></td><td><input type='text' name='login' id='login'/></td></tr><tr><td><label for='pass'><strong>Mot de passe</strong></label></td><td><input type='password' name='pass' id='pass'/></td></tr></table><input type='submit' name='connexion' value='Se connecter'/></form></body></html>"
	c.Header().Set("Content-Type", "text/html")
	c.Header().Set("Content-Length", strconv.Itoa(len(result)))
	io.WriteString(c, result)
}

func ConfigurationPage(c http.ResponseWriter, req *http.Request) {

	result := "<html><body><form action='/login/' method='post'><table><tr><td><label for='login'><strong>Nom de compte</strong></label></td><td><input type='text' name='login' id='login'/></td></tr><tr><td><label for='pass'><strong>Mot de passe</strong></label></td><td><input type='password' name='pass' id='pass'/></td></tr></table><input type='submit' name='connexion' value='Se connecter'/></form></body></html>"
	c.Header().Set("Content-Type", "text/html")
	c.Header().Set("Content-Length", strconv.Itoa(len(result)))
	io.WriteString(c, result)
}

func loginPage(c http.ResponseWriter, req *http.Request) {
	login := req.FormValue("login")
	fmt.Printf(login)
	result := ""
	c.Header().Set("Content-Type", "text/html")
	c.Header().Set("Content-Length", strconv.Itoa(len(result)))
	io.WriteString(c, result)

}

func login(email string, password string) {

}

func initdb() {
	db, err := sql.Open("sqlite3", "foo.db")
	if err != nil {
		fmt.Printf(err.Error())
	}

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	t1 := dbmap.AddTableWithName(Jack{}, "Jack_test").SetKeys(true, "Id")
	fmt.Printf(t1.TableName)
	dbmap.CreateTables()
}


func test()
{
		j := new(Jack)
	j.Id = 2
	b, err := json.Marshal(j)
	if err != nil {
		fmt.Println(err)
		return
	}
	var s Jack
	json.Unmarshal(b, &s)
	fmt.Println(string(b))
	fmt.Printf("%d\n", j.Id)
	fmt.Printf("%d\n", s.Id)

	j.Key_private = "toto"
}


func decrypt(message)
{
	
}

func mail() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"user@example.com",
		"password",
		"mail.example.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"mail.example.com:25",
		auth,
		"sender@example.org",
		[]string{"recipient@example.net"},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initdb()


	http.Handle("/login/", http.HandleFunc(login))
	http.Handle("/", http.HandlerFunc(IndexPage))
	http.Handle("/configuration/", http.HandlerFunc(ConfigurationPage))
	http.ListenAndServe(":8050", nil)
}
