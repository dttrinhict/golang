package models

import (
	"errors"
	json "github.com/goccy/go-json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
)

type UserModel struct {
	Id int
	FullName string
	Email string
	Phone string
	Age int
	Sex string
}

type User interface {
	GetUsers(users []UserModel, page int, limit int) (*[]UserModel, error)
	CreateUser(users []UserModel, user UserModel) ([]UserModel, error)
	GetUserByID(users []UserModel, id int) ([]UserModel, error)
	UpdateUserByID(users []UserModel, user UserModel) ([]UserModel, error)
	DeleteUserByID(users []UserModel, id int) ([]UserModel, error)
}


func NewUser() User {
	return &UserModel{}
}

/* Get users with paging
http://localhost:8080/users?page=1&limit=4
*/
func (u *UserModel) GetUsers(users []UserModel, page int, limit int) (*[]UserModel, error) {
	us := []UserModel{}
	sort.Slice(users, func(i, j int) bool {
		return users[i].Id < users[j].Id
	})
	if len(users)/limit < page && page > 1 {
		return &us, nil
	}
	if len(users) <=  limit  {
		return &users, nil
	}
	if len(users) <=  page*limit {
		us = append(us, users[(page-1)*limit:len(users)]...)
	}else{
		us = append(us, users[(page-1)*limit:page*limit]...)
	}
	return &us, nil
}

/* Create User
curl --location --request POST 'http://localhost:8080/create-user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 3,
    "fullname": "c",
    "email": "c@a.a",
    "phone": "0912111113",
    "age": 30,
    "sex": "male"
}'
*/
func (u *UserModel) CreateUser(users []UserModel, user UserModel) ([]UserModel, error) {
	if CheckExsisted(users, user) {
		return users, errors.New("User ID đã tồn tài trong hệ thống")
	}
	users = append(users, user)
	return users, nil
}


func (u *UserModel) GetUserByID(users []UserModel,id int) ([]UserModel, error) {
	getUsers := []UserModel{}
	for _, v := range users {
		if v.Id == id {
			getUsers = append(getUsers, v)
		}
	}
	return getUsers, nil
}

/*Update user

*/
func (u *UserModel) UpdateUserByID(users []UserModel, user UserModel) ([]UserModel, error) {
	for index, v := range users {
		if v.Id == user.Id {
			users[index] = user
		}
	}
	return users, nil
}

/*Delete user by ids
curl --location --request DELETE 'http://localhost:8080/delete-user?id=1&id=2'
*/
func (u *UserModel) DeleteUserByID(users []UserModel, id int) ([]UserModel, error) {
	us := []UserModel{}
	for index, user := range users {
		if user.Id == id {
			us = append(us, users[:index]...)
			us = append(us, users[index+1:]...)
			return us, nil
		}
	}
	return users, nil
}


func CheckExsisted(users []UserModel, user UserModel) bool {
	for _, v := range users {
		if v.Id == user.Id {
			return true
		}
	}
	return false
}


var users = []UserModel{
	{
		Id: 1,
		FullName: "A",
		Email: "A@a.a",
		Phone: "0912111111",
		Age: 20,
		Sex: "female",
	},
	{
		Id: 2,
		FullName: "B",
		Email: "B@a.a",
		Phone: "0912111112",
		Age: 20,
		Sex: "male",
	},
}


type Server struct {
	user User
}

func NewServer(user User) Server {
	return Server{
		user: user,
	}
}

/* Handle health check server */
func (s *Server)HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("OK"))
}


/* Get users with paging
http://localhost:8080/users?page=1&limit=4
*/
func (s *Server)GetUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	queryPage, ok := query["page"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing query string"))
		return
	}
	queryLimit, ok := query["limit"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing query string"))
		return
	}
	page, _ := strconv.Atoi(queryPage[0])
	limit, _ := strconv.Atoi(queryLimit[0])
	us, _ := s.user.GetUsers(users, page, limit)
	jsonUs, _ := json.Marshal(us)
	w.Write(jsonUs)
}

/* Create User
curl --location --request POST 'http://localhost:8080/create-user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "fullname": "a",
    "email": "a@a.a",
    "phone": "0912111111",
    "age": 30,
    "sex": "male"
}'
*/
func (s *Server)CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var u UserModel
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		err = json.Unmarshal(body, &u)
		if err != nil {
			w.Write([]byte("error:"+err.Error()))
			return
		}
		users, err = s.user.CreateUser(users, u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}else{
			jsonUs, _ := json.Marshal(users)
			w.WriteHeader(http.StatusCreated)
			w.Write(jsonUs)
		}
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupport method"))
	}
}


/*Get user by ids
http://localhost:8080/get-user?id=1
http://localhost:8080/get-user?id=1&id=2
*/
func (s *Server) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	getUsers := []UserModel{}
	if r.Method == "GET" {
		query := r.URL.Query()
		queryIDs, ok := query["id"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing query string"))
			return
		}
		for _, i := range queryIDs {
			id, err := strconv.Atoi(i)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Wrong data input"))
				return
			}
			us, err := s.user.GetUserByID(users, id)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Error while get user"))
				return
			}
			getUsers = append(getUsers, us...)
		}
		jsonUs, _ := json.Marshal(getUsers)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonUs)
	}else{
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Unsupport method"))
	}
}

/*Update user
curl --location --request PUT 'http://localhost:8080/update-user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "fullname": "a",
    "email": "a@a.a",
    "phone": "0912111111",
    "age": 30,
    "sex": "male"
}'
 */
func (s *Server) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "PUT" {
		var u UserModel
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		err = json.Unmarshal(body, &u)
		if err != nil {
			w.Write([]byte("error:"+err.Error()))
			return
		}
		users, err = s.user.UpdateUserByID(users, u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}else{
			jsonUs, _ := json.Marshal(users)
			w.WriteHeader(http.StatusCreated)
			w.Write(jsonUs)
		}
	}else{
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Unsupport method"))
	}
}

/*Delete user by ids
curl --location --request DELETE 'http://localhost:8080/delete-user?id=1&id=2'
*/
func (s *Server) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "DELETE" {
		query := r.URL.Query()
		queryIDs, ok := query["id"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing query string"))
			return
		}
		for _, i := range queryIDs {
			id, err := strconv.Atoi(i)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Wrong query data input"))
				return
			}
			users, err = s.user.DeleteUserByID(users, id)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Error while get user"))
				return
			}
		}
		jsonUs, _ := json.Marshal(users)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonUs)
	}else{
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Unsupport method"))
	}
}

func NewHttp(server Server) {
	http.HandleFunc("/users", RequestInfo(server.GetUsers)) //middleware
	http.HandleFunc("/create-user", RequestInfo(server.CreateUser))
	http.HandleFunc("/get-user", RequestInfo(server.GetUserByID))
	http.HandleFunc("/update-user", RequestInfo(server.UpdateUserById))
	http.HandleFunc("/delete-user", RequestInfo(server.DeleteUserByID))
	// Run server on port
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/* Middleware return request method and path
*/
func RequestInfo(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request method: %v - %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}