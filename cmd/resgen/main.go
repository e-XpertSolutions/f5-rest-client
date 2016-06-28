// Copyright 2014-2016 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// resgen is a small CLI tool that automatically generates Resource from an API.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	gojson "github.com/ChimeraCoder/gojson"

	"e-xpert_solutions/f5-rest-client/f5"
)

var tmpl = `// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package {{ .Pkg }}

import "e-xpert_solutions/f5-rest-client/f5"

// {{ .Name }}ConfigList holds a list of {{ .Name }} configuration.
type {{ .Name }}ConfigList struct {
	Items []{{ .Name }}Config ` + "`json:\"items\"`" + `
	Kind string ` + "`json:\"kind\"`" + `
	SelfLink string ` + "`json:\"selflink\"`" + `
}

// {{ .Name }}Config holds the configuration of a single {{ .Name }}.
type {{ .Name }}Config struct {
	{{ .ItemDef }}
}


// {{ .Name }}Endpoint represents the REST resource for managing {{ .Name }}.
const {{ .Name }}Endpoint = "{{ .Endpoint }}"

// {{ .Name }}Resource provides an API to manage {{ .Name }} configurations.
type {{ .Name }}Resource struct {
	c f5.Client
}

// ListAll  lists all the {{ .Name }} configurations.
func (r *{{ .Name }}Resource) ListAll() (*{{ .Name }}ConfigList, error) {
	var list {{ .Name }}ConfigList
	if err := r.c.ReadQuery(BasePath+{{ .Name }}Endpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single {{ .Name }} configuration identified by id.
func (r *{{ .Name }}Resource) Get(id string) (*{{ .Name }}Config, error) {
	var item {{ .Name }}Config
	if err := r.c.ReadQuery(BasePath+{{ .Name }}Endpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new {{ .Name }} configuration.
func (r *{{ .Name }}Resource) Create(item {{ .Name }}Config) error {
	if err := r.c.ModQuery("POST", BasePath + {{ .Name }}Endpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a {{ .Name }} configuration identified by id.
func (r *{{ .Name }}Resource) Edit(id string, item {{ .Name }}Config) error {
	if err := r.c.ModQuery("PUT", BasePath + {{ .Name }}Endpoint + "/" + id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single {{ .Name }} configuration identified by id.
func (r *{{ .Name }}Resource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath + {{ .Name }}Endpoint + "/"+id, nil); err != nil {
		return err
	}
	return nil
}
`

var t = template.Must(template.New("resource").Parse(tmpl))

type templateData struct {
	Name     string
	ItemDef  string
	Pkg      string
	Endpoint string
}

func renderTemplate(outputDir string, data templateData) error {
	path := filepath.Join(outputDir, strings.ToLower(data.Name)+".go")
	if _, err := os.Stat(path); err == nil {
		return errors.New(path + " already exists")
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create output file '%s' for %s resource: %v", path, data.Name, err)
	}
	defer f.Close()
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to render template for %s resource: %v", data.Name, err)
	}
	bs, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format generated file: %v", err)
	}
	if _, err := f.Write(bs); err != nil {
		return fmt.Errorf("failed to write generated file: %v", err)
	}
	return nil
}

var re = regexp.MustCompile("Items \\[\\]struct \\{(.*)\\} `json:\"items\"`")

func findItemDef(s string) string {
	log.Print("s = ", s)
	matches := re.FindAllStringSubmatch(strings.Replace(s, "\n", "<°BREAK°>", -1), -1)
	if len(matches) > 0 && len(matches[0]) > 1 {
		return strings.Replace(matches[0][1], "<°BREAK°>", "\n", -1)
	}
	return ""
}

type config struct {
	BaseURL  string `json:"base_url"`
	User     string `json:"username"`
	Password string `json:"password"`
	API      map[string]struct {
		RESTPath string `json:"rest_path"`
		Endpoint string `json:"endpoint"`
	} `json:"api"`
}

func parseConfig(path string) (*config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer f.Close()
	var cfg config
	dec := json.NewDecoder(f)
	if err := dec.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}
	return &cfg, nil
}

var (
	cfgPath   = flag.String("c", "resgen.conf", "configuration file path")
	outputDir = flag.String("o", ".", "output directory")
	pkgName   = flag.String("p", "foo", "package name")
)

func main() {
	flag.Parse()

	if info, err := os.Stat(*outputDir); err != nil {
		log.Fatal(err)
	} else if !info.IsDir() {
		log.Fatalf("'%s' is not a directory", *outputDir)
	}

	cfg, err := parseConfig(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	f5Client, err := f5.NewBasicClient(cfg.BaseURL, cfg.User, cfg.Password)
	if err != nil {
		log.Fatal(err)
	}
	f5Client.DisableCertCheck()

	for name, settings := range cfg.API {
		req, err := f5Client.MakeRequest("GET", settings.RESTPath, nil)
		if err != nil {
			log.Print("failed to create http request", settings.RESTPath, ": ", err)
			continue
		}
		resp, err := f5Client.Do(req)
		if err != nil {
			log.Print("failed to query ", settings.RESTPath, ": ", err)
			continue
		}
		defer resp.Body.Close()
		listDef, err := gojson.Generate(resp.Body, name, *pkgName)
		if err != nil {
			log.Print("failed to generate structure for ", settings.RESTPath, ": ", err)
			continue
		}
		data := templateData{
			Name:     name,
			ItemDef:  findItemDef(string(listDef)),
			Pkg:      *pkgName,
			Endpoint: settings.Endpoint,
		}
		if err := renderTemplate(*outputDir, data); err != nil {
			log.Print("failed to render template: ", err)
		}
	}
}
