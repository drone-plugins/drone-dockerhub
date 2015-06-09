package main

import (
	"fmt"
	"os"
	"net/http"
	"net/url"

	"github.com/drone/drone-plugin-go/plugin"
)

type DockerHub struct {
	Token string `json:"token"`
	Repo  string `json:"repo"`
}

func main() {
	vargs := DockerHub{}

	plugin.Param("vargs", &vargs)
	if err := plugin.Parse(); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	endpoint := fmt.Sprintf("https://registry.hub.docker.com/u/%s/trigger/%s/", vargs.Repo, vargs.Token)
	values := url.Values{"build": {"true"}}

	resp, err := http.PostForm(endpoint, values)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	resp.Body.Close()
}
