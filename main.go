package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/drone/drone-plugin-go/plugin"
	try "gopkg.in/matryer/try.v1"
)

type DockerHub struct {
	Token string `json:"token"`
	Repo  string `json:"repo"`
}

var (
	buildCommit string
)

func main() {
	fmt.Printf("Drone DockerHub Plugin built from %s\n", buildCommit)

	vargs := DockerHub{}

	plugin.Param("vargs", &vargs)
	if err := plugin.Parse(); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	endpoint := fmt.Sprintf("https://registry.hub.docker.com/u/%s/trigger/%s/", vargs.Repo, vargs.Token)
	values := url.Values{"build": {"true"}}

	var resp *http.Response
	err := try.Do(func(attempt int) (bool, error) {
		var err error

		resp, err = http.PostForm(endpoint, values)
		return attempt < 5, err
	})

	if err != nil {
		re := regexp.MustCompile("(.*?/trigger/)[^/]+")
		fmt.Println(re.ReplaceAllString(err.Error(), "${1}HIDDEN"))

		os.Exit(1)
	}
	resp.Body.Close()
}
