package f5

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Paths for file upload.
const (
	PathUploadImage = "/mgmt/cm/autodeploy/software-image-uploads"
	PathUploadFile  = "/mgmt/shared/file-transfer/uploads"

	// For backward compatibility
	// DEPRECATED
	UploadRESTPath = PathUploadFile
)

// An UploadResponse holds the responses send by the BigIP API while uploading
// files.
type UploadResponse struct {
	RemainingByteCount int64          `json:"remainingByteCount"`
	UsedChunks         map[string]int `json:"usedChunks"`
	TotalByteCount     int64          `json:"totalByteCount"`
	LocalFilePath      string         `json:"localFilePath"`
	TemporaryFilePath  string         `json:"temporaryFilePath"`
	Generation         int64          `json:"generation"`
	LastUpdateMicros   int64          `json:"lastUpdateMicros"`
}

// UploadFile reads the content of a file from r and uploads it to the BigIP.
// The uploaded file will be named according to the provided filename.
//
// filesize must be the exact file of the file.
//
// The file is split into small chunk, therefore this method may send multiple
// request.
//
// This method returns the latest upload response received.
func (c *Client) UploadFile(r io.Reader, filename string, filesize int64) (*UploadResponse, error) {
	var uploadResp UploadResponse
	for bytesSent := int64(0); bytesSent < filesize; {
		var chunk int64
		if remainingBytes := filesize - bytesSent; remainingBytes >= 512*1024 {
			chunk = 512 * 1024
		} else {
			chunk = remainingBytes
		}

		req, err := c.MakeUploadRequest(PathUploadFile+"/"+filename, io.LimitReader(r, chunk), bytesSent, chunk, filesize)
		if err != nil {
			return nil, err
		}
		resp, err := c.Do(req)
		if err != nil {
			return nil, err
		}
		if err := c.ReadError(resp); err != nil {
			resp.Body.Close()
			return nil, err
		}

		if filesize-bytesSent <= 512*1024 {
			dec := json.NewDecoder(resp.Body)
			if err := dec.Decode(&uploadResp); err != nil {
				resp.Body.Close()
				return nil, err
			}
		}

		bytesSent += chunk
	}
	return &uploadResp, nil
}

// MakeUploadRequest constructs a single upload request.
//
// restPath can be either PathUploadImage or PathUploadFile.
//
// The file to be uploaded is read from r and must not exceed 524288 bytes.
//
// off represents the number of bytes already sent while chunk is the size of
// chunk to be send in this request.
//
// filesize denotes the size of the entire file.
func (c *Client) MakeUploadRequest(restPath string, r io.Reader, off, chunk, filesize int64) (*http.Request, error) {
	if chunk > 512*1024 {
		return nil, fmt.Errorf("chunk size greater than %d is not supported", 512*1024)
	}
	req, err := http.NewRequest("POST", c.makeURL(restPath), r)
	if err != nil {
		return nil, fmt.Errorf("failed to create F5 authenticated request: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Range", fmt.Sprintf("%d-%d/%d", off, off+chunk-1, filesize))
	req.Header.Set("Content-Type", "application/octet-stream")
	if err := c.makeAuth(req); err != nil {
		return nil, err
	}
	return req, nil
}
