package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type LogUser struct {
	Firstname string `json:"Firstname"`
	Password  string `json:"Password"`
}
type User struct {
	UserId     string `json:"UserId"`
	Firstname  string `json:"Firstname"`
	Lastname   string `json:"Lastname"`
	Email      string `json:"Email"`
	Password   string `json:"Password"`
	repassword string `json:"repassword"`
}
type Task struct {
	UserId string `json:"UserId"`
	TaskId string `json:"TaskId"`
	Desc   string `json:"Desc"`
}
type Lis struct {
	Valu string `json:"Valu"`
}
type Del struct {
	Id string `json:"Id"`
}

type Upd struct {
	Id   string `json:"Id"`
	Desc string `json:"Desc"`
}

func dbConn() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:Qazxsw#2!@tcp(localhost:3306)/customer")
	if err != nil {
		fmt.Println("error validating sql.opem arguments")
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.ping")
		panic(err.Error())
	} else {
		fmt.Println("sucess")
	}
	return db

}

//////Show Users;
func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	fmt.Println("Inside the GET user")
	db := dbConn()
	fmt.Println("inside showuser after dbconn")

	selDB, err := db.Query("SELECT * FROM Persons ")
	if err != nil {
		panic(err.Error())
	}

	user := User{}
	res := []User{}
	for selDB.Next() {
		var id string
		var Firstname string
		var Lastname string
		var Email string
		var Password string
		err = selDB.Scan(&id, &Firstname, &Lastname, &Email, &Password)
		if err != nil {
			panic(err.Error())
		}
		user.UserId = id
		user.Firstname = Firstname
		user.Lastname = Lastname
		user.Email = Email
		user.Password = Password
		res = append(res, user)
	}
	jsonResp, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Println(jsonResp)
	w.Write(jsonResp)
	fmt.Println(res)
	fmt.Println("hi")
	defer selDB.Close()
	defer db.Close()

}

////////////Show task
func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	if r.Body == nil {
		fmt.Println("inside errore")
		http.Error(w, "Please send a request body", 400)
		return
	}
	db := dbConn()
	valu := Lis{}
	fmt.Println("Inside the GET task")
	err := json.NewDecoder(r.Body).Decode(&valu)
	if err != nil {
		fmt.Println("eroorr occuring")
		panic(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println("Inside the GET after decode")
	fmt.Println("the value is " + valu.Valu)

	selDB, err := db.Query("SELECT * FROM Task WHERE userid=? ", valu.Valu)
	if err != nil {
		panic(err.Error())
	}

	work := Task{}
	res := []Task{}
	for selDB.Next() {
		var TaskId string
		var Desc string
		var UserId string
		err = selDB.Scan(&TaskId, &UserId, &Desc)
		if err != nil {
			panic(err.Error())
		}
		work.TaskId = TaskId
		work.UserId = UserId
		work.Desc = Desc
		res = append(res, work)

	}
	fmt.Println(res)
	jsonResp, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	//fmt.Println(jsonResp)
	w.Write(jsonResp)
	return

}

///////////SIgnUp User;
func postUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("gonna  post User")
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Println(user.Firstname)
	fmt.Println("cors accepted")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	db := dbConn()
	fmt.Println("inside postUser after dbconn")
	insert, err := db.Prepare("INSERT INTO `customer`.`Persons` VALUES (NULL,?,?,?,?)")
	if err != nil {
		fmt.Println("inside PostError")
		panic(err.Error())
	}
	insert.Exec(user.Firstname, user.Lastname, user.Email, user.Password)
	fmt.Println(insert)
	jsonResp, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Println(jsonResp)
	w.Write(jsonResp)
	return
	defer insert.Close()
	defer db.Close()

	fmt.Fprintf(w, "Hi")
}

////////Add Task
func postTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside post task")
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	work := Task{}

	fmt.Println("body is")
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&work)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Println("id is" + work.UserId)
	fmt.Println(work.Desc)
	fmt.Println("cors accepted")

	db := dbConn()
	fmt.Println("inside postTask after dbconn")
	insert, err := db.Prepare("INSERT INTO `customer`.`Task` VALUES (NULL,?,?)")
	if err != nil {
		fmt.Println("inside PostError")
		panic(err.Error())
	}
	insert.Exec(work.UserId, work.Desc)

	fmt.Println(insert)
	jsonResp, err := json.Marshal(work)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Println(jsonResp)
	w.Write(jsonResp)
	return
	defer insert.Close()
	defer db.Close()

	fmt.Fprintf(w, "Hi")
}

////////Validae user
func validate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside the validate")
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	logval := LogUser{}
	db := dbConn()
	fmt.Println("after the r.body")
	err := json.NewDecoder(r.Body).Decode(&logval)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println("the username is " + logval.Firstname)
	fmt.Println("the Password is " + logval.Password)
	valDb, err := db.Prepare("SELECT * FROM Persons WHERE firstname=? AND password=?")
	if err != nil {
		panic(err.Error())
	}

	user := User{}
	var id string
	var Firstname string
	var Lastname string
	var Email string
	var Password string
	errr := valDb.QueryRow(logval.Firstname, logval.Password).Scan(&id, &Firstname, &Lastname, &Email, &Password)
	if errr != nil {
		if errr == sql.ErrNoRows {

		}

	}
	user.UserId = id
	user.Firstname = Firstname
	user.Lastname = Lastname
	user.Email = Email
	user.Password = Password
	fmt.Println(user)
	jsonResp, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Println(jsonResp)
	w.Write(jsonResp)

}

///////////Delete User
func delTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside deleteUser")
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	work := Del{}
	err := json.NewDecoder(r.Body).Decode(&work)
	if err != nil {
		fmt.Println("eroorr occuring")
		panic(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println("after decode")
	fmt.Println(work.Id)
	db := dbConn()

	fmt.Println("inside deleteuser after dbconn")

	delete, err := db.Prepare("DELETE FROM  Task WHERE taskid= ?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(work.Id)
	jsonResp, err := json.Marshal(work)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Println(jsonResp)
	w.Write(jsonResp)
	defer delete.Close()
	defer db.Close()

	return

}

/////// Update Task
func updTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside UpdateUser")
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	work := Upd{}
	err := json.NewDecoder(r.Body).Decode(&work)
	if err != nil {
		fmt.Println("eroorr occuring")
		panic(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println("after decode")
	fmt.Println(work.Id)
	fmt.Println(work.Desc)
	db := dbConn()
	update, err := db.Prepare("UPDATE Task SET description=? WHERE taskId=?")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(work.Desc, work.Id)
	jsonResp, err := json.Marshal(work)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Println(jsonResp)
	w.Write(jsonResp)
	defer update.Close()
	defer db.Close()

	return

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getUser", getUser).Methods("GET")
	r.HandleFunc("/postUser", postUser).Methods("POST")
	r.HandleFunc("/validate", validate).Methods("POST")
	r.HandleFunc("/postTask", postTask).Methods("POST")
	r.HandleFunc("/getTask", getTask).Methods("POST")
	r.HandleFunc("/delTask", delTask).Methods("POST")
	r.HandleFunc("/updTask", updTask).Methods("POST")
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	errr := http.ListenAndServe(":9090", handlers.CORS(headers, methods, origins)(r)) // setting listening port
	if errr != nil {
		log.Fatal("ListenAndServe: ", errr)

	}

}
