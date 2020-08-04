package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type JobConfigFile struct {
	Presubmits map[string][]JobData `json:"presubmits"`
	Periodics  []JobData            `json:"periodics"`
}

type JobData struct {
	Name        string            `json:"name"`
	Annotations map[string]string `json:"annotations"`
}

const anotationName = "testgrid-alert-email"

func main() {

	if len(os.Args) < 2 {
		logrus.Fatal("Specify the path to the k/test-infra repo")
	}
	testInfraPath := os.Args[1]

	var fileList []string
	jobsPath := testInfraPath + string(os.PathSeparator) + "config/jobs/kubernetes/sig-release/"
	err := filepath.Walk(jobsPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml" {
			fileList = append(fileList, path)
		}
		return nil
	})

	jobsFound := make([]JobData, 0)

	if err != nil {
		logrus.Fatal(err)
	}

	output := ""
	for _, filePath := range fileList {
		output += fmt.Sprintf("File: %s\n", strings.TrimPrefix(filePath, testInfraPath))

		yamlReader, err := os.Open(filePath)
		if err != nil {
			logrus.Fatal(errors.Wrap(err, "opening job files"))
		}

		confData := JobConfigFile{}
		decoder := yaml.NewDecoder(yamlReader)
		for {
			if err := decoder.Decode(&confData); err == io.EOF {
				break
			} else if err != nil {
				logrus.Fatal(errors.Wrap(err, "decoding file"))
			}

			for _, job := range confData.Periodics {
				jobsFound = append(jobsFound, job)
			}

			for repo, collection := range confData.Presubmits {
				logrus.Infof("Checo jobs en %s", repo)
				for _, job := range collection {
					jobsFound = append(jobsFound, job)
				}
			}
		}

		for _, job := range jobsFound {
			jobMailAddr := ""
			for label, value := range job.Annotations {
				if label == anotationName {
					jobMailAddr = value
				}
			}
			if jobMailAddr == "" {
				logrus.WithField("Job", job.Name).Warnf("Mail is missing")
				output += fmt.Sprintf("%s %s missing\n", job.Name, anotationName)
			}
		}
		output += "\n"
	}

	fmt.Print(output)
}
