package meta

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cnosdatabase/cnosdb/pkg/utils"
)

// ResponseLogger is wrapper of http.ResponseWriter that keeps track of its HTTP status
// code and body size
type ResponseLogger struct {
	w      http.ResponseWriter
	status int
	size   int
}

func NewResponseLogger(w http.ResponseWriter) *ResponseLogger {
	return &ResponseLogger{
		w: w,
	}
}

func (l *ResponseLogger) CloseNotify() <-chan bool {
	if notifier, ok := l.w.(http.CloseNotifier); ok {
		return notifier.CloseNotify()
	}
	// needed for response recorder for testing
	return make(<-chan bool)
}

func (l *ResponseLogger) Header() http.Header {
	return l.w.Header()
}

func (l *ResponseLogger) Flush() {
	l.w.(http.Flusher).Flush()
}

func (l *ResponseLogger) Write(b []byte) (int, error) {
	if l.status == 0 {
		// Set status if WriteHeader has not been called
		l.status = http.StatusOK
	}

	size, err := l.w.Write(b)
	l.size += size
	return size, err
}

func (l *ResponseLogger) WriteHeader(s int) {
	l.w.WriteHeader(s)
	l.status = s
}

func (l *ResponseLogger) Status() int {
	if l.status == 0 {
		// This can happen if we never actually write data, but only set response headers.
		l.status = http.StatusOK
	}
	return l.status
}

func (l *ResponseLogger) Size() int {
	return l.size
}

// redact any occurrence of a password parameter, 'p'
func redactPassword(r *http.Request) {
	q := r.URL.Query()
	if p := q.Get("p"); p != "" {
		q.Set("p", "[REDACTED]")
		r.URL.RawQuery = q.Encode()
	}
}

// Common Log Format: http://en.wikipedia.org/wiki/Common_Log_Format

// buildLogLine creates a common log format
// in addition to the common fields, we also append referrer, user agent and request ID
func buildLogLine(l *ResponseLogger, r *http.Request, start time.Time) string {

	redactPassword(r)

	username := parseUsername(r)

	host, _, err := net.SplitHostPort(r.RemoteAddr)

	if err != nil {
		host = r.RemoteAddr
	}

	uri := r.URL.RequestURI()

	referer := r.Referer()

	userAgent := r.UserAgent()

	fields := []string{
		host,
		"-",
		detect(username, "-"),
		fmt.Sprintf("[%s]", start.Format("02/Jan/2006:15:04:05 -0700")),
		r.Method,
		uri,
		r.Proto,
		detect(strconv.Itoa(l.Status()), "-"),
		strconv.Itoa(l.Size()),
		detect(referer, "-"),
		detect(userAgent, "-"),
		r.Header.Get("Request-Id"),
		fmt.Sprintf("%s", time.Since(start)),
	}

	return strings.Join(fields, " ")
}

// detect detects the first presense of a non blank string and returns it
func detect(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

// parses the username either from the url or auth header
func parseUsername(r *http.Request) string {
	var (
		username = ""
		url      = r.URL
	)

	// get username from the url if passed there
	if url.User != nil {
		if name := url.User.Username(); name != "" {
			username = name
		}
	}

	// Try to get the username from the query param 'u'
	q := url.Query()
	if u := q.Get("u"); u != "" {
		username = u
	}

	// Try to get it from the authorization header if set there
	if username == "" {
		if u, _, ok := r.BasicAuth(); ok {
			username = u
		}
	}
	return username
}

// sanitize redacts passwords from query string for logging.
func sanitize(r *http.Request) {
	values := r.URL.Query()
	for i, q := range values["q"] {
		values["q"][i] = utils.Sanitize(q)
	}
	r.URL.RawQuery = values.Encode()
}
