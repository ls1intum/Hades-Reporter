package main

import (
	"bytes"
	"encoding/json"
	"github.com/caarlos0/env/v9"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type ResultMetadata struct {
	JobName                  string `json:"jobName" env:"JOB_NAME"`
	UUID                     string `json:"uuid" env:"UUID"`
	AssignmentRepoBranchName string `json:"assignmentRepoBranchName" env:"ASSIGNMENT_REPO_BRANCH_NAME" envDefault:"main"`
	IsBuildSuccessful        bool   `json:"isBuildSuccessful" env:"IS_BUILD_SUCCESSFUL"`
	AssignmentRepoCommitHash string `json:"assignmentRepoCommitHash" env:"ASSIGNMENT_REPO_COMMIT_HASH"`
	TestsRepoCommitHash      string `json:"testsRepoCommitHash" env:"TESTS_REPO_COMMIT_HASH"`
	BuildCompletionTime      string `json:"buildCompletionTime" env:"BUILD_COMPLETION_TIME"`
}

type Config struct {
	Endpoint string `env:"ENDPOINT,required"`
}

var config Config
var metadata ResultMetadata

func main() {
	log.Info("Starting the application...")

	loadEnv()

	timestamp := time.Now().Format(time.RFC3339)
	log.Infof("Timestamp: %s", timestamp)

	payload := map[string]string{
		"uuid":           metadata.UUID,
		"buildStartTime": timestamp,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	sendResponse(jsonData)
}

func loadEnv() {
	log.Info("Loading Environment variables...")
	// Load the environment variables
	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}
	if err := env.Parse(&metadata); err != nil {
		log.Fatal(err)
	}

	// Dump the parsed configuration for easier debugging
	dumpEnvConfig()
	log.Info("Environment variables loaded successfully")
}

// dumpEnvConfig prints all configuration and metadata values that were read from the environment
// Sensitive values (like API tokens) are partially masked so they don't leak into CI logs.
func dumpEnvConfig() {
	log.Info("============== Runtime Configuration ==============")
	log.Infof("Endpoint: %s", config.Endpoint)

	log.Info("================= Metadata Values =================")
	log.Infof("JobName: %s", metadata.JobName)
	log.Infof("UUID: %s", metadata.UUID)
	log.Infof("AssignmentRepoBranchName: %s", metadata.AssignmentRepoBranchName)
	log.Infof("IsBuildSuccessful: %v", metadata.IsBuildSuccessful)
	log.Infof("AssignmentRepoCommitHash: %s", metadata.AssignmentRepoCommitHash)
	log.Infof("TestsRepoCommitHash: %s", metadata.TestsRepoCommitHash)
	log.Infof("BuildCompletionTime: %s", metadata.BuildCompletionTime)
	log.Info("===================================================")
}

func sendResponse(json []byte) {
	// Create a new request using http
	log.Info("Sending the response to the API...")
	req, err := http.NewRequest("POST", config.Endpoint, bytes.NewBuffer(json))
	if err != nil {
		log.Debug("Error creating the request")
		log.Fatal(err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Create a Client
	client := &http.Client{}

	// Send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		log.Debug("Error sending the request")
		log.Fatal(err)
	}

	log.Info("Response Status: ", resp.Status)

	// Close response body
	defer resp.Body.Close()
}
