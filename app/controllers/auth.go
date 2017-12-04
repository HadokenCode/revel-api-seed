package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/obitux/revel-api-seed/app/models"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	gormc.TxnController
}

// TODO: read secret from conf file?
var hmacSecret = []byte{97, 48, 97, 50, 97, 98, 105, 49, 99, 102, 83, 53, 57, 98, 52, 54, 97, 102, 99, 12, 12, 13, 56, 34, 23, 16, 78, 67, 54, 34, 32, 21}

// Register create a user and returns token to client.
// params: email, password
// result: token with user.id stores in `sub` field.
func (c Auth) Register() revel.Result {
	type Data struct {
		Email    string
		Name     string
		Password string
	}
	jsonData := Data{}
	c.Params.BindJSON(&jsonData)
	email := jsonData.Email
	name := jsonData.Name
	password := jsonData.Password

	// check json
	if email == "" || password == "" || name == "" {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("params not valid.")
	}

	// check if the email have already exists in DB
	var user = models.User{}
	c.Txn.First(&user, "email = ?", email)
	if user.ID != 0 {
		c.Response.Status = http.StatusConflict
		return c.RenderJSON("user have already exists.")
	}

	// Create user struct
	token := encodeToken(email)
	user = models.User{Name: name, Email: email}
	user.SetNewPassword(password)
	user.Active = true

	// Validate user struct
	user.Validate(c.Validation)
	if c.Validation.HasErrors() {
		log.Println(c.Validation.Errors[0].Message)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("bad email address.") // FIXME
	}

	// Save user info to DB
	c.Txn.Create(&user)

	msg := make(map[string]string)
	msg["email"] = email
	msg["result"] = "user created"
	msg["token"] = token
	return c.RenderJSON(msg)
}

// Login authenticate via email and password, if the user is valid,
// returns the token to client.
func (c Auth) Login() revel.Result {
	type Data struct {
		Email    string
		Password string
	}
	jsonData := Data{}
	c.Params.BindJSON(&jsonData)
	email := jsonData.Email
	password := jsonData.Password

	// check if the email exists in DB
	var user = models.User{Email: email}
	c.Txn.First(&user, "email = ?", email)
	if user.ID == 0 {
		c.Response.Status = http.StatusConflict
		return c.RenderJSON("invalid email or password")
	}

	// check password
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("invalid email or password")
	}

	// get token
	msg := make(map[string]string)
	msg["result"] = "login success"
	msg["token"] = encodeToken(email)
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(msg)
}

func encodeToken(email string) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(hmacSecret)

	return tokenString
}

func decodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		log.Println(err)
		return nil, err
	}
	return claims, nil
}

func checkUser(c *revel.Controller) revel.Result {
	// Read the token from Authorization Header then decode it
	claims, err := decodeToken(c.Request.Header.Get("Authorization"))

	// if provided token is invalid
	if err != nil {
		msg := make(map[string]string)
		msg["error"] = "wrong token"
		c.Response.Status = http.StatusForbidden
		return c.RenderJSON(msg)
	}

	// TODO : check for token validity
	// TODO : check if user is_active
	fmt.Println(claims) // FIXME: remove

	return nil
}
