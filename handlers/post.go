package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

type Post struct {
	Logger *log.Logger
}

func NewPost(logger *log.Logger) *Post {
	return &Post{logger}
}

type YAML struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
}

func (post *Post) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write([]byte("Received request for 'POST'\n"))
	post.Logger.Println("Received request for 'POST'")

	file, _, err := r.FormFile("file")
	if err != nil {
		post.Logger.Println("Error retrieving the file: ", err)
		return
	}
	defer file.Close()

	// read the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		post.Logger.Println("Error reading file: ", err)
		return
	}

	// unmarshal the yaml
	var job YAML
	err = yaml.Unmarshal(data, &job)
	if err != nil {
		post.Logger.Println("Error unmarshalling yaml: ", err)
		return
	}
	fmt.Println(job)
}
