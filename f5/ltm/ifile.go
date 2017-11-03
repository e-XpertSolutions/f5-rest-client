package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type IFileList struct {
	Items    []IFile `json:"items"`
	Kind     string  `json:"kind"`
	SelfLink string  `json:"selflink"`
}

type IFile struct {
	AppService  string `json:"appService,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	TMPartition string `json:"tmPartition,omitempty"`
	Partition   string `json:"partition,omitempty"`
}

const IFileEndpoint = "/ifile"

type IFileResource struct {
	c *f5.Client
}

func (r *IFileResource) ListAll() (*IFileList, error) {
	var list IFileList
	if err := r.c.ReadQuery(BasePath+IFileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *IFileResource) Get(id string) (*IFile, error) {
	var item IFile
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
