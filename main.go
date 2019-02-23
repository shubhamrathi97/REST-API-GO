package main

import (
	"encoding/json"
	"fmt"
    "log"
	"net/http"
	// "json"
)

type User struct {
    ID        string   `json:"id,omitempty"`
    Name 	  string   `json:"Name,omitempty"`
    DOB       string   `json:"DOB,omitempty"`
    Friend   []*User     `json:"user,omitempty"`
}

type FriendData struct {
	UserID string
	FriendUserID  string
}

var users []User

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var user User
		_ = json.NewDecoder(r.Body).Decode(&user)
		users = append(users, user)
		json.NewEncoder(w).Encode(user)
	default:
		fmt.Fprintf(w, "Method Not Allow")
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id := r.URL.Query().Get("id")
		if id != "" {
			for _, item := range users {
				if item.ID == id {
					json.NewEncoder(w).Encode(item)
            		return
				}
			}
		}
		json.NewEncoder(w).Encode(&User{})
	default:
		fmt.Fprintf(w, "Method Not Allow")
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		var user User
		_ = json.NewDecoder(r.Body).Decode(&user)
		if user.ID != "" {
			for i, item := range users {
				if item.ID == user.ID {
					users[i] = user
					// users = append(users, item)
					json.NewEncoder(w).Encode(user)
				}
			}
		}
	default:
		fmt.Fprintf(w, "Method Not Allow")
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(users)
	default:
		fmt.Fprintf(w, "Method Not Allow")
	}
}

func getFriends(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Hello Allow")
	default:
		fmt.Fprintf(w, "Method Not Allow")
	}
}

func addFriend(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var friendData FriendData
		var friendUser User
		var user User
		_ = json.NewDecoder(r.Body).Decode(&friendData)
		if friendData.UserID != "" && friendData.FriendUserID != "" {
			for _, item := range users {
				if item.ID == friendData.FriendUserID {
					friendUser = item			
					for i, item := range users {					
						if item.ID == friendData.UserID {
							user = item
							user.Friend = append(user.Friend, &friendUser)
							users[i] = user
							json.NewEncoder(w).Encode(user)
							return
						}
					}		
				}
				fmt.Fprintf(w, "friend not found")
			}
		}
		fmt.Fprintf(w, "Friend ID or User ID is missing")
	default:
		fmt.Fprintf(w, "Method Not Allow")
	}
}
	// our main function
func main() {
	users = append(users, User{ID: "1", Name: "shubham", DOB:"10-01-1997"})
	users = append(users, User{ID: "2", Name: "abc", DOB:"10-01-1997"})


	http.HandleFunc("/", sayhelloName) // set router
	http.HandleFunc("/createuser", createUser) //post create
	http.HandleFunc("/user", getUser) //get & post for update
	http.HandleFunc("/updateuser", updateUser) //get & post for update	
	http.HandleFunc("/users", getUsers) // return all user

	http.HandleFunc("/userfriends", getFriends) //return all friends of user 
	http.HandleFunc("/addfriend", addFriend) //post to add friend
	
	err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
