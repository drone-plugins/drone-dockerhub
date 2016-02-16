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
	buildDate string
)

func main() {
	fmt.Printf("Drone DockerHub Plugin built at %s\n", buildDate)

	vargs := DockerHub{}
	build := plugin.Build{}

	plugin.Param("vargs", &vargs)
	plugin.Param("build", &build)

	if err := plugin.Parse(); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	endpoint := fmt.Sprintf("https://registry.hub.docker.com/u/%s/trigger/%s/", vargs.Repo, vargs.Token)
	values := url.Values{"source_type": {"Branch"}, "source_name": {build.Branch}}

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
