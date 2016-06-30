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

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

var tmpl = `// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package {{ .Pkg }}

import "github.com/e-XpertSolutions/f5-rest-client/f5"

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

func isCapital(b byte) bool {
	return strings.ToUpper(string(b)) == string(b)
}

func uppertCamelToUnderscore(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}
	buf := bytes.NewBufferString(strings.ToLower(string(s[0])))
	for i := 1; i < len(s); i++ {
		if isCapital(s[i]) {
			if !isCapital(s[i-1]) || (i != len(s)-1 && !isCapital(s[i+1])) {
				buf.WriteString("_")
			}
		}
		buf.WriteString(strings.ToLower(string(s[i])))
	}
	return buf.String()
}

func renderTemplate(outputDir string, data templateData) error {
	path := filepath.Join(outputDir, uppertCamelToUnderscore(data.Name)+".go")
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

const tmplClient = `/* Client structure definition code:
{{ range $name, $type:= .Fields }}
{{ $name }} {{ $type }}{{ end }}
*/

/* Constructor initialization code:
{{ range $name, $type:= .Fields }}
{{ $name }}: {{ $type }}{c: c},{{ end }}
*/

/* Accessors definition code:
{{ $receiverName := .ClientReceiver }}{{ $receiverType := .ClientType }}
{{ range $name, $type:= .Fields }}
// {{ ucfirst $name }} returns a configured {{ $type }}.
func ({{ $receiverName }} {{ $receiverType }}) {{ ucfirst $name }}() *{{ $type }} {
	return &{{ $receiverName }}.{{ $name }}
}
{{ end }}
*/
`

func UCFirst(s string) string {
	if len(s) < 2 {
		return strings.ToUpper(s)
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func LCFirst(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

var tc = template.Must(template.New("client").Funcs(template.FuncMap{"ucfirst": UCFirst}).Parse(tmplClient))

type clientTemplateData struct {
	ClientReceiver string
	ClientType     string
	Fields         map[string]string
}

func renderClientTemplate(outputDir string, data clientTemplateData) error {
	path := filepath.Join(outputDir, strings.ToLower(data.ClientType)+".go.part")
	if _, err := os.Stat(path); err == nil {
		return errors.New(path + " already exists")
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create output file '%s' for %s client: %v", path, data.ClientType, err)
	}
	defer f.Close()
	if err := tc.Execute(f, data); err != nil {
		return fmt.Errorf("failed to render template for %s client: %v", data.ClientType, err)
	}
	return nil
}

var re = regexp.MustCompile("Items \\[\\]struct \\{(.*)\\} `json:\"items\"`")

func findItemDef(s string) string {
	matches := re.FindAllStringSubmatch(strings.Replace(s, "\n", "<°BREAK°>", -1), -1)
	if len(matches) > 0 && len(matches[0]) > 1 {
		return strings.Replace(matches[0][1], "<°BREAK°>", "\n", -1)
	}
	return ""
}

type config struct {
	BaseURL        string            `json:"base_url"`
	User           string            `json:"username"`
	Password       string            `json:"password"`
	ClientReceiver string            `json:"client_receiver"`
	ClientType     string            `json:"client_type"`
	Module         string            `json:"module"`
	API            map[string]string `json:"api"`
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

func findEndpoint(restPath, module string) string {
	idx := strings.Index(restPath, module)
	if idx == -1 {
		return restPath
	}
	return restPath[idx+len(module):]
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

	clientData := clientTemplateData{
		ClientReceiver: cfg.ClientReceiver,
		ClientType:     cfg.ClientType,
		Fields:         make(map[string]string),
	}

	for name, restPath := range cfg.API {
		req, err := f5Client.MakeRequest("GET", restPath, nil)
		if err != nil {
			log.Print("failed to create http request", restPath, ": ", err)
			continue
		}
		resp, err := f5Client.Do(req)
		if err != nil {
			log.Print("failed to query ", restPath, ": ", err)
			continue
		}
		defer resp.Body.Close()
		listDef, err := gojson.Generate(resp.Body, name, *pkgName)
		if err != nil {
			log.Print("failed to generate structure for ", restPath, ": ", err)
			continue
		}
		data := templateData{
			Name:     name,
			ItemDef:  findItemDef(string(listDef)),
			Pkg:      *pkgName,
			Endpoint: findEndpoint(restPath, cfg.Module),
		}
		if err := renderTemplate(*outputDir, data); err != nil {
			log.Print("failed to render template: ", err)
		}
		clientData.Fields[LCFirst(name)] = name + "Resource"
	}

	if err := renderClientTemplate(*outputDir, clientData); err != nil {
		log.Print("failed to generate client code: ", err)
	}
}
