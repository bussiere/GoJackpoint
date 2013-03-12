package main

import (
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

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
	Key_public            rune
	Key_private           rune
	Message_Id            []rune
	Statut                string
	Avatar                string
}

type Key_Public struct {
	Key string
}

type Key_Private struct {
	Key string
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

func LoginPage(c http.ResponseWriter, req *http.Request) {
	login := req.FormValue("login")
	fmt.Printf(login)
	result := ""
	c.Header().Set("Content-Type", "text/html")
	c.Header().Set("Content-Length", strconv.Itoa(len(result)))
	io.WriteString(c, result)

}

func concatene(first string, second string) (retour string) {
	var (
		buffer bytes.Buffer
	)
	buffer.WriteString(first)
	buffer.WriteString(second)
	return buffer.String()
}

func encryptAes(message string, clef string) {

	plaintext := []byte(message)

	if len(os.Args) > 1 {

		plaintext = []byte(os.Args[1])

	}

	// Setup a key that will encrypt the other text.

	key_text := clef
	fmt.Println(len(key_text))

	if len(os.Args) > 2 {

		key_text = os.Args[2]

	}

	// We chose our cipher type here in this case

	// we are using AES.

	c, err := aes.NewCipher([]byte(key_text))

	if err != nil {

		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)

		os.Exit(-1)

	}

	// We use the CFBEncrypter in order to encrypt

	// the whole stream of plaintext using the

	// cipher setup with c and a iv.

	cfb := cipher.NewCFBEncrypter(c, commonIV)

	ciphertext := make([]byte, len(plaintext))

	cfb.XORKeyStream(ciphertext, plaintext)

	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	// We must now convert the ciphertext to base64

	// this will allow for the encrypted data to be

	// visible to copy and paste into the decrypter.

	// base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(ciphertext)))

	// base64.StdEncoding.Encode(base64Text, []byte(ciphertext))

	// fmt.Printf("base64: %s\n", base64Text)
}

func decryptAes(message string, clef string) {

	plaintext := []byte(message)

	ciphertext := make([]byte, len(message))

	if len(os.Args) > 1 {

		plaintext = []byte(os.Args[1])

	}

	// Setup a key that will encrypt the other text.

	key_text := clef
	fmt.Println(len(key_text))

	if len(os.Args) > 2 {

		key_text = os.Args[2]

	}

	// We chose our cipher type here in this case

	// we are using AES.

	c, err := aes.NewCipher([]byte(key_text))

	if err != nil {

		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)

		os.Exit(-1)

	}

	cfbdec := cipher.NewCFBDecrypter(c, commonIV)

	plaintextCopy := make([]byte, len(plaintext))

	cfbdec.XORKeyStream(plaintextCopy, ciphertext)

	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)

}

func getprivateKeyUser(user Jack) (privatekey string) {

}

func EncodeB64(message string) (retour string) {

	data := base64.StdEncoding.EncodeToString([]byte(message))

	fmt.Printf("%q\n", data)
	return data

}

func DecodeB64(message string) (retour string) {

	data, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Printf("%q\n", data)
	return string(data)

}

func loginfunc(email string, password string) {

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

func test() {
	var test string
	j := new(Jack)
	j.Id = 2
	j.Key_private = 1
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

	c, err := json.Marshal(j.Key_private)
	fmt.Println(string(c))

	k := new(Key_Public)
	k.Key = "titi"
	c, err = json.Marshal(k)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(c))
	j.Key_private = 2
	fmt.Printf("base64: %s\n", j.Key_private)
	test = EncodeB64(string(j.Key_private))
	fmt.Printf("base64: %s\n", test)
	test = DecodeB64(test)
	fmt.Printf("base64: %s\n", test)
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

	http.Handle("/login/", http.HandlerFunc(LoginPage))
	http.Handle("/", http.HandlerFunc(IndexPage))
	http.Handle("/configuration/", http.HandlerFunc(ConfigurationPage))
	test()
	//http.ListenAndServe(":8050", nil)
}
