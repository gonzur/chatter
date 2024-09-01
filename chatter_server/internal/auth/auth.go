package auth

import (
	"chatter-server/internal/db"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json:"id" sql:"id"`
	Username string `json:"username" sql:"username"`
	// all byte data must be base64 encoded for json to work
	Password []byte `json:"password" sql:"password"`
	Hash     []byte `json:"-" sql:"hash"`
}

func saltedHash(content []byte, salt []byte) ([]byte, error) {
	hasher := sha256.New()
	if _, err := hasher.Write(append(content, salt...)); err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func randomSalt(nBytes int) ([]byte, error) {
	salt := make([]byte, nBytes)
	if _, err := rand.Reader.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func CreateUser(c *gin.Context) {
	if length, err := strconv.Atoi(c.GetHeader("Content-Length")); err != nil || length > 300 {
		c.JSON(500, ErrorMessage{Code: 2, Message: "Request too long"})
		return
	}
	// newerr, _ := io.ReadAll(c.Request.Body)
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.String(400, "%s", err.Error())
		return
	}

	if len(string(user.Password)) > 32 || len(user.Username) > 256 {
		c.JSON(500, ErrorMessage{Code: 3, Message: "Username or Password too long."})
		return
	}

	salt, err := randomSalt(8)
	if err != nil {
		c.Status(500)
		return
	}

	hashedPW, err := saltedHash(user.Password, salt)
	if err != nil {
		log.Println(err.Error())
		c.Status(500)
		return
	}

	cont, cancel := db.DBContext()
	defer cancel()

	if ct, err := db.DB.Exec(cont,
		"INSERT INTO USERS (username, password, hash) VALUES ($1, $2, $3)",
		user.Username, hashedPW, salt); err != nil {
		log.Println(err.Error())
		if ct.RowsAffected() != 1 {
			log.Println("User", user.Username, "Not Inserted")
		}
	}

	c.Status(200)
}

type ErrorMessage struct {
	Message string `json:"message"`
	Code    uint32 `json:"code"`
}

func Login(c *gin.Context) {
	if length, err := strconv.Atoi(c.GetHeader("Content-Length")); err != nil || length < 300 {
		c.JSON(500, ErrorMessage{Code: 2, Message: "Request too long"})
		return
	}

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err.Error())
		c.Status(500)
		return
	}

	if len(string(user.Password)) > 16 || len(user.Username) > 256 {
		c.JSON(500, ErrorMessage{Code: 3, Message: "Password Must be less than 16 Characters. Username must be less than 256 characters."})
		return
	}

	storedUser, err := db.QueryRow[User](
		"SELECT username, password, hash FROM USERS WHERE username == $1",
		user.Username)
	if err != nil || storedUser == nil {
		c.Header("WWW-Authenticate", "Basic realm=Application")
		c.JSON(401, ErrorMessage{Message: "Wrong Username/Password", Code: 1})
		return
	}

	hashedPW, err := saltedHash(user.Password, storedUser.Hash)
	if err != nil {
		c.Status(500)
		return
	}

	if !reflect.DeepEqual(hashedPW, user.Password) {
		c.Header("WWW-Authenticate", "Basic realm=Application")
		c.JSON(401, ErrorMessage{Message: "Wrong Username/Password", Code: 1})
		return
	}

	sKey := AddSession(user.Username)
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("session", string(sKey), 0, "/", "localhost", false, true)
	c.Status(200)
}
