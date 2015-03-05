package dumper

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func RunDump(config Configuration) {
	fmt.Println(config.Dbstrings[0])
	db, err := sql.Open("mysql", config.Dbstrings[0])
	if err != nil {
		fmt.Println("Config loading error")
	} else {
		fmt.Println("Config loaded")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("DB connection error: ", err)
	} else {
		fmt.Println("DB connection success")
	}

	defer db.Close()

	rows, err := db.Query("select user_id, name from users where user_id between ? and ?", 401, 410)
	if err != nil {
		fmt.Println("Error running query:", err)
	}
	defer rows.Close()
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.UserId, &user.Name)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
		}
		fmt.Printf("User ID: %d, User Name: %s \n", user.UserId, user.Name)
	}

	defer db.Close()
}
