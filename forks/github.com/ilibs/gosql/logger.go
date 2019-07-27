package gosql

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ie310mu/ie310go/common/logsagent"
)

const (
	fmtLogQuery     = `Query: %s`
	fmtLogArgs      = `Args:  %#v`
	fmtLogError     = `Error: %v`
	fmtLogTimeTaken = `Time:  %0.5fs`
)

var (
	reInvisibleChars = regexp.MustCompile(`[\s\r\n\t]+`)
)

// QueryStatus represents the status of a query after being executed.
type QueryStatus struct {
	Query string
	Args  interface{}

	Start time.Time
	End   time.Time

	Err error
}

// String returns a formatted log message.
func (q *QueryStatus) String() string {
	lines := make([]string, 0, 8)

	if query := q.Query; query != "" {
		query = reInvisibleChars.ReplaceAllString(query, ` `)
		query = strings.TrimSpace(query)
		lines = append(lines, fmt.Sprintf(fmtLogQuery, query))
	}

	if q.Args != nil {
		lines = append(lines, fmt.Sprintf(fmtLogArgs, q.Args))
	}

	if q.Err != nil {
		lines = append(lines, fmt.Sprintf(fmtLogError, q.Err))
	}

	lines = append(lines, fmt.Sprintf(fmtLogTimeTaken, float64(q.End.UnixNano()-q.Start.UnixNano())/float64(1e9)))

	return strings.Join(lines, "\n")
}

// Logger represents a logging collector. You can pass a logging collector to
// gosql.SetLogger(myCollector) to make it collect QueryStatus messages
// after executing a query.
type Logger interface {
	Printf(format string, v ...interface{})
}

type defaultLogger struct {
	logging bool
	log     Logger
}

func (d *defaultLogger) Log(m *QueryStatus) {
	if d.logging {
		//d.log.Printf("\n\t%s\n\n", strings.Replace(m.String(), "\n", "\n\t", -1))
		logsagent.Debug("\n\t%s\n\n", strings.Replace(m.String(), "\n", "\n\t", -1))
	}
}

func (d *defaultLogger) SetLogging(logging bool) {
	d.logging = logging
}

var logger = &defaultLogger{log: log.New(os.Stderr, "", log.LstdFlags)}

//SetLogger ..
func SetLogger(l Logger) {
	logger.log = l
}
