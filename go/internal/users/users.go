package users

import (
	"database/sql"
	"log"

	"github.com/selithrarion/angular-go_graphql-test/internal/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func GetUserIdByUsername(username string) (int, error) {
	statement, err := database.DB.Prepare("select ID from Users WHERE Username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(username)

	var ID int
	err = row.Scan(&ID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return ID, nil
}

func (user *User) Create() {
	statement, err := database.DB.Prepare("INSERT INTO Users(Username, Password) VALUES(?,?)")
	print(statement)
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword, err := HashPassword(user.Password)
	_, err = statement.Exec(user.Username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func (user *User) Authenticate() bool {
	statement, err := database.DB.Prepare("select Password from Users WHERE Username = ?")
	if err != nil {
		log.Fatal(err)
	}

	row := statement.QueryRow(user.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) book {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}
