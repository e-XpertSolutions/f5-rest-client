package f5

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Paths for file upload.
const (
	PathUploadImage = "/mgmt/cm/autodeploy/software-image-uploads"
	PathUploadFile  = "/mgmt/shared/file-transfer/uploads"
	PathUploadUCS   = "mgmt/shared/file-transfer/ucs-uploads"

	// For backward compatibility
	// DEPRECATED
	UploadRESTPath = PathUploadFile
)

// Paths for file download.
const (
	PathDownloadUCS   = "/mgmt/shared/file-transfer/ucs-downloads"
	PathDownloadImage = "/mgmt/cm/autodeploy/software-image-downloads"
)

// MaxChunkSize is the maximum chunk size allowed by the iControl REST
const MaxChunkSize = 1048576

// FileTransferOptions contains SSH configuration for downloading and uploading
// UCS using SFTP.
type FileTransferOptions struct {
	UseSFTP      bool
	ClientConfig *ssh.ClientConfig
}

// FileTransferOption is a function type to set the transfer options.
type FileTransferOption func(*FileTransferOptions)

// WithSFTP sets the ssh configuration for file transfer.
func WithSFTP(config *ssh.ClientConfig) FileTransferOption {
	return func(o *FileTransferOptions) {
		o.UseSFTP = true
		o.ClientConfig = config
	}
}

// DownloadUCS downloads an UCS file and writes its content to w.
func (c *Client) DownloadUCS(w io.Writer, filename string, opts ...FileTransferOption) (n int64, err error) {

	options := FileTransferOptions{}
	for _, o := range opts {
		o(&options)
	}

	if options.UseSFTP {
		if options.ClientConfig == nil {
			fn := func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
				for _, q := range questions {
					if strings.Contains(q, "Password") {
						return []string{c.password}, nil
					}
				}
				return []string{}, nil
			}
			options.ClientConfig = &ssh.ClientConfig{
				User: c.username,
				Auth: []ssh.AuthMethod{
					ssh.KeyboardInteractive(fn),
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			}
		}
		return c.downloadUsingSSH(w, filename, options)
	}

	// BigIP 12.x.x only support download requests with a Content-Range header,
	// thus, it is required to know the size of the file to download beforehand.
	//
	// BigIP 13.x.x automatically download the first chunk and provide the
	// Content-Range header with all information in the response, which is far
	// more convenient. Unfortunately, we need to support BigIP 12 and as a
	// result, we need to first retrieve the UCS file size information.
	resp, err := c.SendRequest("GET", "/mgmt/tm/sys/ucs", nil)
	if err != nil {
		return 0, fmt.Errorf("cannot retrieve info for ucs file: %v", err)
	}
	defer resp.Body.Close()
	if err := c.ReadError(resp); err != nil {
		return 0, fmt.Errorf("cannot retrieve info for ucs file: %v", err)
	}

	// As far as I know, there is no direct way to fetch UCS file info for a
	// specific file and therefore we need to list all UCS files and search
	// for the one we want in the list.
	var ucsInfo struct {
		Items []struct {
			APIRawValues struct {
				Filename string `json:"filename"`
				FileSize string `json:"file_size"`
			} `json:"apiRawValues"`
		} `json:"items"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&ucsInfo); err != nil {
		return 0, fmt.Errorf("cannot decode ucs file info: %v", err)
	}

	// File size is a raw string and we need to parse it in order to extract the
	// size as an integer.
	var rawFileSize string
	for _, item := range ucsInfo.Items {
		if strings.HasSuffix(item.APIRawValues.Filename, filename) {
			rawFileSize = strings.TrimSuffix(item.APIRawValues.FileSize, " (in bytes)")
			break
		}
	}
	if rawFileSize == "" {
		return 0, errors.New("ucs file does not exist")
	}
	fileSize, err := strconv.ParseInt(rawFileSize, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("malformed file size in ucs file info: %v", err)
	}

	if n, err = c.download(w, PathDownloadUCS+"/"+filename, fileSize, MaxChunkSize); err != nil {
		return 0, fmt.Errorf("cannot download ucs file: %v", err)
	}
	return
}

func (c *Client) downloadUsingSSH(w io.Writer, filename string, opts FileTransferOptions) (int64, error) {

	path := string(os.PathSeparator) + filepath.Join("var", "local", "ucs", filename)

	parsedURL, err := url.Parse(c.baseURL)
	if err != nil {
		return 0, fmt.Errorf("downloadUsingSSH: cannot parse baseURL: %w", err)
	}

	conn, err := ssh.Dial("tcp", parsedURL.Hostname()+":22", opts.ClientConfig)
	if err != nil {
		return 0, fmt.Errorf("downloadUsingSSH: cannot connect via ssh: %w", err)
	}
	defer conn.Close()

	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		return 0, fmt.Errorf("downloadUsingSSH: cannot create sftp client: %w", err)
	}
	defer sftpClient.Close()

	file, err := sftpClient.Open(path)
	if err != nil {
		return 0, fmt.Errorf("downloadUsingSSH: cannot open file %q: %w", filename, err)
	}
	defer file.Close()

	return io.Copy(w, file)
}

// DownloadImage downloads BIG-IP images from the API and writes it to w.
//
// Download can take some time due to the size of the image files.
func (c *Client) DownloadImage(w io.Writer, filename string) (n int64, err error) {
	// This is necessary to get the filesize first using bash command in order
	// to support BIG-IP 12.x.x.
	out, err := c.Exec("wc -c /shared/images/" + filename)
	if err != nil {
		return 0, fmt.Errorf("cannot get file size: %v", err)
	}
	fileSizeStr := strings.TrimSpace(out.CommandResult)
	pos := strings.Index(fileSizeStr, " ")
	if pos != -1 {
		fileSizeStr = fileSizeStr[:pos]
	}
	fileSize, err := strconv.ParseInt(fileSizeStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot read file size: %v", err)
	}

	if n, err = c.download(w, PathDownloadImage+"/"+filename, fileSize, MaxChunkSize); err != nil {
		return 0, fmt.Errorf("cannot download image file: %v", err)
	}
	return
}

func (c *Client) download(w io.Writer, restPath string, filesize, chunkSize int64) (n int64, err error) {
	if filesize < chunkSize {
		chunkSize = filesize
	}
	return c.downloadByChunks(w, restPath, filesize, 0, chunkSize)
}

func (c *Client) downloadByChunks(w io.Writer, restPath string, filesize, offset, chunkSize int64) (n int64, err error) {
	req, err := c.MakeRequest("GET", restPath, nil)
	if err != nil {
		return 0, err
	}

	// Bound limit to filesize
	limit := offset + chunkSize - 1
	if limit >= filesize {
		limit = filesize - 1
	}

	req.Header.Set("Content-Range", fmt.Sprintf("%d-%d/%d", offset, limit, filesize))

	resp, err := c.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if err := c.ReadError(resp); err != nil {
		return 0, err
	}

	if n, err = io.Copy(w, resp.Body); err != nil {
		return 0, err
	}

	if limit < filesize-1 {
		nn, err := c.downloadByChunks(w, restPath, filesize, offset+chunkSize, chunkSize)
		if err != nil {
			return 0, err
		}
		n += nn
	}

	return
}

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
	return c.upload(r, PathUploadFile, filename, filesize)
}

// UploadImage reads the content of an disk image from r and uploads it to the
// BigIP.
//
// The uploaded image will be named according to the provided filename.
//
// filesize must be the exact file of the file.
//
// The file is split into small chunk, therefore this method may send multiple
// request.
//
// This method returns the latest upload response received.
func (c *Client) UploadImage(r io.Reader, filename string, filesize int64) (*UploadResponse, error) {
	return c.upload(r, PathUploadImage, filename, filesize)
}

// UploadUCS reads the content of an UCS archive from r and uploads it to the
// BigIP.
//
// The uploaded UCS archive will be named according to the provided filename.
//
// filesize must be the exact file of the file.
//
// The file is split into small chunk, therefore this method may send multiple
// request.
//
// This method returns the latest upload response received.
func (c *Client) UploadUCS(r io.Reader, filename string, filesize int64, opts ...FileTransferOption) (*UploadResponse, error) {
	return c.upload(r, PathUploadUCS, filename, filesize, opts...)
}

func (c *Client) upload(r io.Reader, restPath, filename string, filesize int64, opts ...FileTransferOption) (*UploadResponse, error) {
	var uploadResp UploadResponse

	options := FileTransferOptions{}
	for _, o := range opts {
		o(&options)
	}

	if options.UseSFTP {
		if options.ClientConfig == nil {
			fn := func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
				for _, q := range questions {
					if strings.Contains(q, "Password") {
						return []string{c.password}, nil
					}
				}
				return []string{}, nil
			}
			options.ClientConfig = &ssh.ClientConfig{
				User: c.username,
				Auth: []ssh.AuthMethod{
					ssh.KeyboardInteractive(fn),
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			}
		}
		_, err := c.uploadUsingSSH(r, filename, options)
		if err != nil {
			return nil, err
		}
		return &uploadResp, nil
	}

	for bytesSent := int64(0); bytesSent < filesize; {
		var chunk int64
		if remainingBytes := filesize - bytesSent; remainingBytes >= 512*1024 {
			chunk = 512 * 1024
		} else {
			chunk = remainingBytes
		}

		req, err := c.makeUploadRequest(restPath+"/"+filename, io.LimitReader(r, chunk), bytesSent, chunk, filesize)
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

		resp.Body.Close()

		bytesSent += chunk
	}
	return &uploadResp, nil
}

func (c *Client) uploadUsingSSH(r io.Reader, filename string, opts FileTransferOptions) (int64, error) {

	path := string(os.PathSeparator) + filepath.Join("var", "local", "ucs", filename)

	parsedURL, err := url.Parse(c.baseURL)
	if err != nil {
		return 0, fmt.Errorf("uploadUsingSSH: cannot parse baseURL: %w", err)
	}

	conn, err := ssh.Dial("tcp", parsedURL.Hostname()+":22", opts.ClientConfig)
	if err != nil {
		return 0, fmt.Errorf("uploadUsingSSH: cannot connect via ssh: %w", err)
	}
	defer conn.Close()

	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		return 0, fmt.Errorf("uploadUsingSSH: cannot create sftp client: %w", err)
	}
	defer sftpClient.Close()

	dstFile, err := sftpClient.Create(path)
	if err != nil {
		return 0, fmt.Errorf("uploadUsingSSH: cannot create file  %q on remote host: %w", filename, err)
	}
	defer dstFile.Close()

	n, err := dstFile.ReadFrom(r)
	if err != nil {
		return n, fmt.Errorf("uploadUsingSSH: cannot copy reader content to the file %q: %w", filename, err)
	}

	return n, nil
}

// makeUploadRequest constructs a single upload request.
//
// restPath can be any of the Path* constants defined at the top of this file.
//
// The file to be uploaded is read from r and must not exceed 524288 bytes.
//
// off represents the number of bytes already sent while chunk is the size of
// chunk to be send in this request.
//
// filesize denotes the size of the entire file.
func (c *Client) makeUploadRequest(restPath string, r io.Reader, off, chunk, filesize int64) (*http.Request, error) {
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
