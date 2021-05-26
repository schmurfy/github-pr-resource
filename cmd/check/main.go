package main

import (
	"encoding/json"
	"log"
	"os"

	resource "github.com/telia-oss/github-pr-resource"
)

func main() {
	var request resource.CheckRequest

	decoder := json.NewDecoder(os.Stdin)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&request); err != nil {
		log.Fatalf("failed to unmarshal request: %s", err)
	}

	f, err := os.Create("/tmp/check_params")
	if err != nil {
		log.Fatalf("failed to open /tmp/check_params: %s", err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(&request)
	if err != nil {
		log.Fatalf("failed to save inputs: %s", err)
	}

	if err := request.Source.Validate(); err != nil {
		log.Fatalf("invalid source configuration: %s", err)
	}
	github, err := resource.NewGithubClient(&request.Source)
	if err != nil {
		log.Fatalf("failed to create github manager: %s", err)
	}
	response, err := resource.Check(request, github)
	if err != nil {
		log.Fatalf("check failed: %s", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		log.Fatalf("failed to marshal response: %s", err)
	}
}
