package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type NullString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}

func main() {
	//use sql.Open to connect database by return *sql.db
	db, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/myDB")
	if err != nil {
		log.Fatal(err)
	}
	/*var (
		id int
		//data     string
		//username string
		//password string
		name      string
		Average   float32
		rate      int
		minplayer int
		maxplayer int
		year      int
		weight    float32
		playtime  int
	)*/
	type board struct {
		id        int
		name      string
		Average   float32
		rate      int
		minplayer int
		maxplayer int
		year      int
		weight    float32
		playtime  int
	}
	/*type user struct {
		Username string
		Password string
	}
	type article struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}*/
	//var boardlist []board
	//var datas []string
	/*
			user{
				id int
		 		data string
			}*/
	//drop table
	/*query := "Drop TABLE board"
	_, err = db.Exec(query)
	if err != nil {
		// do something here
		log.Fatal(err)
	}*/
	//create table
	/*query := "CREATE TABLE board (ID integer, Name varchar(255),Average DECIMAL(10,5),rate integer,weight DECIMAL(10,5),minplayer integer,maxplayer integer,playtime integer,year integer)"
	_, err = db.Exec(query)
	if err != nil {
		// do something here
		log.Fatal(err)
	}

	query = "LOAD DATA  INFILE 'D:/mysql/data/Uploads/out.csv' INTO TABLE board FIELDS TERMINATED BY ',' LINES TERMINATED BY '" + "\n" + "'IGNORE 1 ROWS;"
	_, err = db.Exec(query)
	if err != nil {
		// do something here
		log.Fatal(err)
	}
	log.Println("load table ")*/
	//value: 'string' is necessary
	/*id := 1
	title := "first"
	content := "first article"
	i := 0*/
	//
	//create article
	//
	/*_, err = db.Exec("INSERT INTO article(Id,Title,Content) VALUES(" + strconv.Itoa(id) + ", '" + title + "','" + content + "')")
	if err != nil {
		log.Fatal(err)
	}*/
	//var s string
	//
	//select all article
	//
	/*rows, err := db.Query("SELECT * FROM board")
	if err != nil {
		// do something here
		log.Fatal(err)
	}
	i := 0
	board_t := board{}
	if !rows.Next() {
		log.Println("return null list")
	} else {
		err = rows.Scan(&id, &name, &Average, &rate, &weight, &minplayer, &maxplayer, &playtime, &year)
		i++
		if err != nil {
			// do something here
			log.Fatal(err)
		}
		board_t.id = id
		board_t.name = name
		board_t.Average = Average
		board_t.rate = rate
		board_t.weight = weight
		board_t.minplayer = minplayer
		board_t.maxplayer = maxplayer
		board_t.playtime = playtime
		board_t.year = year
		boardlist = append(boardlist, board_t)
	}
	for rows.Next() {
		err = rows.Scan(&id, &name, &Average, &rate, &weight, &minplayer, &maxplayer, &playtime, &year)
		i++
		if err != nil {
			// do something here
			log.Fatal(err)
		}
		board_t.id = id
		board_t.name = name
		board_t.Average = Average
		board_t.rate = rate
		board_t.weight = weight
		board_t.minplayer = minplayer
		board_t.maxplayer = maxplayer
		board_t.playtime = playtime
		board_t.year = year
		boardlist = append(boardlist, board_t)
		if i > 10 {

			break
		}
	}*/

	/*s := "Denny"
	ss := "1111"
	rows, err := db.Query("SELECT * FROM Users where Username = '" + s + "' AND Password = '" + ss + "'")
	if err != nil {
		// do something here
		log.Fatal(err)
	}
	user_t := user{}
	//calll next will go to next rows
	if !rows.Next() {
		log.Println("no match data")
	} else {
		err := rows.Scan(&username, &password)
		if err != nil {
			log.Fatal(err)
		}
		user_t = user{
			Username: username,
			Password: password,
		}
		log.Println(user_t)
	}
	user_a := user{
		Username: "Denny",
		Password: "1111",
	}
	for rows.Next() {
		err := rows.Scan(&username, &password)
		if err != nil {
			log.Fatal(err)
		}
		user_t := user{
			Username: username,
			Password: password,
		}
		log.Println(user_t)

	}
	if user_a == user_t {
		log.Println("same")
	} else {
		log.Println(user_t)
	}
	// insert and select ** 'string' is neccesary
	/*for i := 1; i < 10; i++ {
		stmt, err := db.Prepare("INSERT INTO user(id,data) VALUES(" + strconv.Itoa(i) + ",'dolly')")
		if err != nil {
			log.Fatal(err)
		}
		_, err = stmt.Exec()
		if err != nil {
			log.Fatal(err)
		}
	}*/
	/*rows, err := db.Query("select * from user")
	//rows, err := db.Query("select id, data from user where id = ?", 3)
	if err != nil {
		// do something here
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &data)
		if err != nil {
			log.Fatal(err)
		}
		ids = append(ids, id)
		datas = append(datas, data)
	}
	defer rows.Close()
	//ids datas store all data
	log.Println(ids)
	log.Println(datas)
	/*query := "INSERT INTO user VALUES (id,data) VALUE(2,denny);"
	_, err = db.Exec(query)
	if err != nil {
		// do something here
		log.Fatal(err)
	}
	/*query1 := "INSERT INTO example  VALUES ("
	for i := 2; i < 10; i++ {
		_, err = db.Exec(query1 + strconv.Itoa(i) + ", denny" + strconv.Itoa(i) + ");")
		if err != nil {
			// do something here
			log.Fatal(err)
		}
	}
	query = "Select * From example;"
	rows, err := db.Query(query)
	if err != nil {
		// do something here
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &data)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, data)
	}
	defer rows.Close()
	err = db.Ping()
	if err != nil {
		// do something here
		log.Fatal(err)
	}*/
	print("success")
	defer db.Close()
}
