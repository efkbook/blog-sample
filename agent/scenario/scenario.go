package scenario

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type StatusError struct {
	StatusCode int
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("server returns status: %v", e.StatusCode)
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode != 200 {
		return &StatusError{resp.StatusCode}
	}
	return nil
}

const (
	indexURL  = "http://localhost/"
	loginURL  = "http://localhost/login"
	signupURL = "http://localhost/signup"
	saveURL   = "http://localhost/save"
)

func login(email, password string) error {
	form := url.Values{}
	form.Set("email", email)
	form.Set("password", password)

	resp, err := http.PostForm(loginURL, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return errors.Wrap(checkResponse(resp), "failed to login")
}

func signup(email, name, password string) error {
	form := url.Values{}
	form.Set("email", email)
	form.Set("name", name)
	form.Set("password", password)

	resp, err := http.PostForm(signupURL, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return errors.Wrap(checkResponse(resp), "failed to signup")
}

func loginOrSignup(email, name, password string) (err error) {
	err = login(email, password)
	if err != nil {
		if _, ok := errors.Cause(err).(*StatusError); ok {
			return signup(email, name, password)
		}
	}
	return err
}

func create(title, body string) error {
	form := url.Values{}
	form.Set("title", title)
	form.Set("body", body)

	resp, err := http.PostForm(saveURL, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return errors.Wrap(checkResponse(resp), "failed to create a entry")
}
