package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	//"path/filepath"
	
	"gopkg.in/yaml.v2"
	log "github.com/sirupsen/logrus"
)

type MetaData struct {
	Name string `yaml: name`
	Labels map[string]string `yaml: labels`
}

type ConfigMap struct {
	ApiVersion string `yaml: apiVersion`
	Kind string `yaml: kind`
	Metadata MetaData `yaml: metadata`
	Data map[string]string `yaml: data`
}


func EmptyConfigMap(name string) (*ConfigMap) {
	cm := &ConfigMap{
		ApiVersion: "v1",
		Kind: "ConfigMap",
		Metadata: MetaData{
			Name: name,
		},
		Data: map[string]string{},
	}
	return cm
}

// Adds a file
func (c *ConfigMap) AddFile(f *ConfigMapFile) (error) {
	c.Data[f.Name] = string(f.Contents)
	return nil
}

func (c *ConfigMap) DumpYaml() {
	yml, err := yaml.Marshal(*c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(yml))
}

/**
 * ConfigMapFile
 *
 * Will be come a key in the ConfigMap
 */
type ConfigMapFile struct {
	// Actual FS path
	Path string
	// Name/key for configmap (basename(Path))
	Name string `yaml: name`
	// Contents (as bytes)
	Contents []byte
}

func NewConfigMapFile(fpath string) (*ConfigMapFile, error) {
	contents, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	
	cm := &ConfigMapFile{
		Path: fpath,
		Name: path.Base(fpath),
		Contents: contents,
	}
	return cm, nil
}


func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	name := flag.String("name", "my-config", "The ConfigMap Metadata.Name")
	dir := flag.String("dir", cwd, "The input directory")

	flag.Parse()
	
	//var files []string
	files, err := ioutil.ReadDir(*dir)
	
	if err != nil {
		panic(err)
	}

	cm := EmptyConfigMap(*name)
	
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		
		fullpath := path.Join(*dir, file.Name())
		cmf, err := NewConfigMapFile(fullpath)
		if err != nil {
			log.Warnf("Problem with file: %s", err)
			continue
		}
		err = cm.AddFile(cmf)
		if err != nil {
			log.Warnf("Couldn't add file %s (%s)", fullpath, err)
		}
	}

	cm.DumpYaml()
}
