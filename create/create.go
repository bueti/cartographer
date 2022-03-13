/*
Copyright © 2022 Benjamin Buetikofer <bbu@ik.me>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package create

import (
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

const templateDir = "templates/"

type ApplicationCrd struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name       string
		Namespace  string
		Finalizers []string `yaml:""`
	}
	Spec struct {
		Destination struct {
			Namespace string
			Server    string
		}
		Project string
		Source  struct {
			Helm struct {
				ValueFiles []string `yaml:"valueFiles"`
			}
			RepoURL        string `yaml:"repoURL"`
			Path           string
			TargetRevision string `yaml:"targetRevision"`
		}
		SyncPolicy struct {
			Automated struct {
				Prune      bool
				AllowEmpty bool `yaml:"allowEmpty"`
				SelfHeal   bool `yaml:"selfHeal"`
			}
			SyncOptions []string `yaml:"syncOptions"`
		} `yaml:"syncPolicy"`
	}
}

type Chart struct {
	ApiVersion   string
	Name         string
	Description  string
	Type         string
	Version      string
	AppVersion   string
	Dependencies []string `yaml:",flow"`
}

func CreateApplicationCrd(chartName, chartRepository, chartVersion, project, namespace, valueFiles string, secrets bool) {
	reader := reader{fileName: "application-crd.yaml"}
	configVal, _ := getConfig(&reader)

	configVal.Metadata.Name = chartName
	configVal.Spec.Destination.Namespace = namespace
	configVal.Spec.Project = project
	configVal.Spec.Source.RepoURL = chartRepository
	configVal.Spec.Source.Path = "configuration/" + chartName
	configVal.Spec.Source.TargetRevision = chartVersion
	files := strings.Split(valueFiles, ",")
	if secrets {
		files = append(files, "secrets://secrets.yaml")
	}
	configVal.Spec.Source.Helm.ValueFiles = files

	d, err := yaml.Marshal(&configVal)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	ioutil.WriteFile(chartName+".yaml", d, 0755)
}

func getConfig(reader YamlReader) (*ApplicationCrd, error) {
	file, _ := reader.readConfig()

	conf := ApplicationCrd{}
	err := yaml.Unmarshal(file, &conf)
	if err != nil {
		return &ApplicationCrd{}, err
	}

	return &conf, nil
}

type YamlReader interface {
	readConfig() ([]byte, error)
}
type reader struct {
	fileName string
}

func (r *reader) readConfig() ([]byte, error) {

	yfile, err := ioutil.ReadFile(templateDir + r.fileName)

	if err != nil {
		log.Fatal(err)
	}

	return yfile, err
}
