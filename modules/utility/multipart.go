package utility

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"strings"
)

// MultipartEngine provides multipart upload utilities.
type MultipartEngine struct{}

// NewMultipartEngine creates a new MultipartEngine.
func NewMultipartEngine() *MultipartEngine {
	return &MultipartEngine{}
}

// BuildForm builds a multipart form with the given fields.
func (e *MultipartEngine) BuildForm(fields map[string]string) (string, []byte, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	for key, value := range fields {
		err := writer.WriteField(key, value)
		if err != nil {
			return "", nil, fmt.Errorf("failed to write field %s: %w", key, err)
		}
	}

	err := writer.Close()
	if err != nil {
		return "", nil, fmt.Errorf("failed to close writer: %w", err)
	}

	return writer.FormDataContentType(), buf.Bytes(), nil
}

// BuildFormWithFile builds a multipart form with a file field.
func (e *MultipartEngine) BuildFormWithFile(filename string, content []byte) (string, []byte, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return "", nil, fmt.Errorf("failed to create form file: %w", err)
	}

	_, err = part.Write(content)
	if err != nil {
		return "", nil, fmt.Errorf("failed to write file content: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return "", nil, fmt.Errorf("failed to close writer: %w", err)
	}

	return writer.FormDataContentType(), buf.Bytes(), nil
}

// ParseBoundary extracts the boundary from a content-type header.
func (e *MultipartEngine) ParseBoundary(contentType string) string {
	parts := strings.Split(contentType, "boundary=")
	if len(parts) < 2 {
		return ""
	}
	return parts[1]
}

// InjectField injects a field into an existing multipart body.
func (e *MultipartEngine) InjectField(body []byte, boundary string, key string, value string) ([]byte, error) {
	newPart := fmt.Sprintf("\r\n--%s\r\nContent-Disposition: form-data; name=\"%s\"\r\n\r\n%s", boundary, key, value)

	sep := []byte("\r\n--" + boundary + "--\r\n")
	idx := bytes.LastIndex(body, sep)
	if idx == -1 {
		return nil, fmt.Errorf("could not find multipart end boundary")
	}

	result := append(body[:idx], []byte(newPart)...)
	result = append(result, sep...)
	return result, nil
}