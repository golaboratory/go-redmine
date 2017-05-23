package redmine

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type userResult struct {
	User User `json:"user"`
}

type usersResult struct {
	Users []User `json:"users"`
}

type User struct {
	Id          int          `json:"id"`
	Login       string       `json:"login"`
	Firstname   string       `json:"firstname"`
	Lastname    string       `json:"lastname"`
	Mail        string       `json:"mail"`
	CreatedOn   string       `json:"created_on"`
	LatLoginOn  string       `json:"last_login_on"`
	APIKey      string       `json:"apt_key"`
	Memberships []Membership `json:"memberships"`
}

type authResult struct {
	User User `json:"user"`
}

func (c *Client) Auth(username string, password string) (*User, error) {

	var authstring = url.QueryEscape(username) + ":" + url.QueryEscape(password) + "@"
	var endpoint = c.endpoint[0:strings.Index(c.endpoint, "//")+2] + authstring + c.endpoint[strings.Index(c.endpoint, "//")+2:]
	res, err := c.Get(endpoint + "/users/current.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r authResult
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.User, nil
}

func (c *Client) Users() ([]User, error) {
	res, err := c.Get(c.endpoint + "/users.json?key=" + c.apikey + c.getPaginationClause())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r usersResult
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return r.Users, nil
}

func (c *Client) User(id int) (*User, error) {
	res, err := c.Get(c.endpoint + "/users/" + strconv.Itoa(id) + ".json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r userResult
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.User, nil
}

func (c *Client) CurrentUser() (*User, error) {
	res, err := c.Get(c.endpoint + "/users/current.json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r userResult
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.User, nil
}
