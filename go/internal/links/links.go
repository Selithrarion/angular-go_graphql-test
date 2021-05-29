package links

import (
	"github.com/selithrarion/angular-go_graphql-test/internal/pkg/database"
	"github.com/selithrarion/angular-go_graphql-test/internal/users"
	"log"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func GetAll() []Link {
	stmt, err := database.DB.Prepare("SELECT id, title, address FROM Links")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}

func (link Link) Save() int64 {
	stmt, err := database.DB.Prepare("INSERT INTO Links(Title,Address, UserID) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	return id
}
