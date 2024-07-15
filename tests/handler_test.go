package functions

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"asciiweb/functions"
)

// Setup function to create temporary template files
func setup() {
	os.MkdirAll("templates", os.ModePerm)

	ioutil.WriteFile("templates/index.html", []byte(`
<!DOCTYPE html>
<html>
<head>
    <title>{{ .pageTitle }}</title>
</head>
<body>
    {{ if .Result }}
    <pre>{{ .Result }}</pre>
    {{ else }}
    <h1>{{ .pageTitle }}</h1>
    {{ end }}
</body>
</html>
`), 0o644)

	ioutil.WriteFile("templates/error.html", []byte(`
<!DOCTYPE html>
<html>
<head>
    <title>Error</title>
</head>
<body>
    <h1>{{ . }}</h1>
</body>
</html>
`), 0o644)
}

// Teardown function to clean up temporary template files
func teardown() {
	os.RemoveAll("templates")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestAscii(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		form       url.Values
		statusCode int
		expected   string
	}{
		{
			name:       "Incorrect URL Path",
			method:     "POST",
			path:       "/wrongpath",
			form:       url.Values{"text": {"Hello"}, "banner": {"standard"}},
			statusCode: http.StatusNotFound,
			expected:   "page not found",
		},
		{
			name:       "Incorrect Method",
			method:     "GET",
			path:       "/ascii-art",
			form:       nil,
			statusCode: http.StatusMethodNotAllowed,
			expected:   "Method Not Allowed",
		},
		{
			name:       "Missing Text Field",
			method:     "POST",
			path:       "/ascii-art",
			form:       url.Values{"banner": {"standard"}},
			statusCode: http.StatusBadRequest,
			expected:   "Bad Request",
		},
		{
			name:       "Missing Banner Field",
			method:     "POST",
			path:       "/ascii-art",
			form:       url.Values{"text": {"Hello"}},
			statusCode: http.StatusBadRequest,
			expected:   "Bad Request",
		},
		{
			name:       "Invalid Banner",
			method:     "POST",
			path:       "/ascii-art",
			form:       url.Values{"text": {"Hello"}, "banner": {"invalid"}},
			statusCode: http.StatusInternalServerError,
			expected:   "Internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			var err error

			if tt.form != nil {
				req, err = http.NewRequest(tt.method, tt.path, strings.NewReader(tt.form.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			} else {
				req, err = http.NewRequest(tt.method, tt.path, nil)
			}
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(functions.Ascii)
			handler.ServeHTTP(rr, req)

			// Check status code
			if status := rr.Code; status != tt.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.statusCode)
			}

			// Check response body contains expected content
			if !strings.Contains(rr.Body.String(), tt.expected) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expected)
			}
		})
	}
}
