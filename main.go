package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
  "errors"
)

const GITHUB_API_URL = "https://api.github.com/"
const GITHUB_API_VERSION_HEADER_KEY = "Accept"
const GITHUB_API_VERSION_HEADER_VALUE = "application/vnd.github.v3+json"

type Parent struct {
	FullName string `json:"full_name"`
}

type Repository struct {
	Parent Parent `json:"parent"`
}

func fetchParentRepo(repoName string) (string, error) {
  url := GITHUB_API_URL + "repos/" + repoName
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add(GITHUB_API_VERSION_HEADER_KEY, GITHUB_API_VERSION_HEADER_VALUE)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

  if res.StatusCode != http.StatusOK {
    return "", errors.New("not ok")
  }

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	repo := Repository{}
	err = json.Unmarshal(body, &repo)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	return repo.Parent.FullName, nil
}

func main() {
	original, err := fetchParentRepo("jyane/grpc")

	if err != nil {
    log.Panic(err)
  }
  log.Print(original)
}
