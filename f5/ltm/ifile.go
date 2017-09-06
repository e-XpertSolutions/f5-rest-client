package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type IFileConfigList struct {
	Items    []IFileConfig `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

type IFileConfig struct {
	AppService  string `json:"appService,omitempty"`
	Description string `json:"description,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	TMPartition string `json:"tmPartition,omitempty"`
}

const IFileEndpoint = "/ifile"

type IFileResource struct {
	c *f5.Client
}

func (r *IFileResource) ListAll() (*IFileConfigList, error) {
	var list IFileConfigList
	if err := r.c.ReadQuery(BasePath+IFileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *IFileResource) Get(id string) (*IFileConfig, error) {
	var item IFileConfig
	if err := r.c.ReadQuery(BasePath+IFileEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *IFileResource) Create(id, fileObject string) error {
	data := map[string]string{
		"name":      id,
		"file-name": fileObject,
	}
	if err := r.c.ModQuery("POST", BasePath+IFileEndpoint, data); err != nil {
		return err
	}
	return nil
}

func (r *IFileResource) Edit(id, fileObject string) error {
	data := map[string]string{
		"name":      id,
		"file-name": fileObject,
	}
	if err := r.c.ModQuery("PUT", BasePath+IFileEndpoint+"/"+id, data); err != nil {
		return err
	}
	return nil
}

func (r *IFileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+IFileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
