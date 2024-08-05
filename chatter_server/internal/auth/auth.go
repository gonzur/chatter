package auth

import (
	"chatter-server/internal/db"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json:"id" sql:"id"`
	Username string `json:"username" sql:"username"`
	Password []byte `json:"password" sql:"password"`
	Hash     []byte `json:"-" sql:"hash"`
}

func CreateUser(c *gin.Context) {
	var user User
	if c.ShouldBindJSON(&user) != nil {
		c.Status(400)
		return
	}

	salt := make([]byte, 16)
	if _, err := rand.Reader.Read(salt); err != nil {
		log.Println(err.Error())
		c.Status(500)
		return
	}

	hasher := sha256.New()
	if _, err := hasher.Write(append([]byte(user.Password), salt...)); err != nil {
		log.Println(err.Error())
		c.Status(500)
		return
	}
	hashedPW := hasher.Sum(nil)

	cont, cancel := db.DBContext()
	defer cancel()

	if ct, err := db.DB.Exec(cont,
		"INSERT INTO USERS (username, password, hash) VALUES ($1, $2, $3)",
		user.Username, hashedPW, salt); err != nil {
		log.Println(err.Error())
	} else if ct.RowsAffected() != 1 {
		log.Println("User", user.Username, "Not Inserted")
		c.Status(500)
		return
	}

	c.Status(200)
}

type ErrorMessage struct {
	Message string `json:"message"`
	Code    uint16 `json:"code"`
}

func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err.Error())
		c.Status(500)
		return
	}

	storedUser, err := db.QueryRow[User](
		"SELECT username, password, hash FROM users WHERE username == $1",
		user.Username)

	if err != nil || storedUser == nil {
		c.JSON(200, ErrorMessage{Message: "Wrong Username/Password", Code: 1})
		return
	}

	hashedPW, err := saltedHash(user.Password, storedUser.Hash)
	if err != nil {
		c.Status(500)
		return
	}

	if !reflect.DeepEqual(hashedPW, user.Password) {
		c.JSON(200, ErrorMessage{Message: "Wrong Username/Password", Code: 1})
		return
	}

	sKey := AddSession(user.Username)

	c.SetCookie("session", string(sKey), 0, "/", "localhost", false, true)
	c.Status(200)
}

func saltedHash(content []byte, salt []byte) ([]byte, error) {
	hasher := sha256.New()
	if _, err := hasher.Write(append(content, salt...)); err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}
