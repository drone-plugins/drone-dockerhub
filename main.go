package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/drone/drone-plugin-go/plugin"
	try "gopkg.in/matryer/try.v1"
	"net/http"
	"os"
	"regexp"
)

type DockerHub struct {
	Token string `json:"token"`
	Repo  string `json:"repo"`
}

type DockerHubValues struct {
	SourceType string `json:"source_type"`
	SourceName string `json:"source_name"`
}

var (
	buildDate string
)

var re = regexp.MustCompile("(.*?/trigger/)[^/]+")

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
	values := DockerHubValues{SourceType: "Branch", SourceName: build.Branch}
	values_json, err := json.Marshal(values)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(values_json))

	if err != nil {
		fmt.Println(re.ReplaceAllString(err.Error(), "${1}HIDDEN"))

		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	var resp *http.Response
	err = try.Do(func(attempt int) (bool, error) {
		var err error

		resp, err = client.Do(req)
		return attempt < 5, err
	})

	if err != nil {
		fmt.Println(re.ReplaceAllString(err.Error(), "${1}HIDDEN"))

		os.Exit(1)
	}
	resp.Body.Close()
}
