package dumper

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

func readDb(db *sql.DB, dbChan chan<- *User) {
	defer db.Close()
	rows, err := db.Query("select user_id, name from users limit ?, ?", 0, 10)
	if err != nil {
		log.Fatal("Error running query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.UserId, &user.Name)
		if err != nil {
			log.Fatal("Error scanning rows:", err)
		}
		//log.Printf("User ID: %d, User Name: %s \n", user.UserId, user.Name)
		dbChan <- user
	}

	close(dbChan)
}

func channelMerge(cs []chan *User) <-chan *User {
	var wg sync.WaitGroup
	out := make(chan *User, 5)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan *User) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func dbConnect(config Configuration, msgChan chan string) {
	var dbChans []chan *User
	dbChans = make([]chan *User, len(config.Dbstrings), len(config.Dbstrings))

	for i, dbString := range config.Dbstrings {
		dbChans[i] = make(chan *User, 5)
		db, err := sql.Open("mysql", dbString)
		if err != nil {
			log.Fatal("Config loading error: ", err)
		} else {
			log.Println("Config loaded")
		}

		err = db.Ping()
		if err != nil {
			log.Fatal("DB connection error: ", err)
		} else {
			log.Println("DB connection success")
			msgChan <- "Connection success"
			go readDb(db, dbChans[i])
		}
	}
	dbChan := channelMerge(dbChans)
	for user := range dbChan {
		fmt.Printf("User ID: %d, User Name: %s \n", user.UserId, user.Name)
	}

	close(msgChan)
}

func sq(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for n := range in {
			out <- n
		}
		close(out)
	}()
	return out
}

func RunDump(config Configuration) {
	// Set up the pipeline.
	c := make(chan string)
	go dbConnect(config, c)
	out := sq(c)

	// Consume the output.
	for output := range out {
		fmt.Println(output)
	}
}
