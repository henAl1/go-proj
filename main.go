package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//User struct ....
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

//APIResponse ...
type APIResponse struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Total   int    `json:"total_pages"`
	Data    []User `json:"data"`
}

func getUsers(body []byte) (*APIResponse, error) {
	var s = new(APIResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func main() {
	res, err := http.Get("https://reqres.in/api/users")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	s, err := getUsers([]byte(body))
	fmt.Println(len(s.Data))
	for _, v := range s.Data {

		strr := fmt.Sprintf("%s %s with avatar  %s", v.FirstName, v.LastName, v.Avatar)
		fmt.Println(strr)
	}

}

