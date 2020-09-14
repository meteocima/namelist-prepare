package namelist

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"time"
)

type dateArg struct {
	Iso   string
	Year  int
	Month int
	Day   int
	Hour  int
}

type tmplArgs struct {
	MetGridConstants string
	Start            dateArg
	End              dateArg
	OneHBeforeStart  string
	OneHAfterStart   string
	Hours            int
	MetgridLevels    int
}

func createTemplateArgs(start, end time.Time, hours int) tmplArgs {
	var args tmplArgs

	if hours > 24 {
		args.MetGridConstants = "constants_name                = 'TAVGSFC',"
	} else {
		args.MetGridConstants = ""
	}
	args.Start.Day = start.Day()
	args.Start.Month = int(start.Month())
	args.Start.Year = start.Year()
	args.Start.Hour = start.Hour()
	args.Start.Iso = start.Format("2006-01-02_15:00:00")

	args.OneHBeforeStart = start.Add(-1 * time.Hour).Format("2006-01-02_15:00:00")
	args.OneHAfterStart = start.Add(1 * time.Hour).Format("2006-01-02_15:00:00")

	args.End.Day = end.Day()
	args.End.Month = int(end.Month())
	args.End.Year = end.Year()
	args.End.Hour = end.Hour()
	args.End.Iso = end.Format("2006-01-02_15:00:00")

	metgridLevels := 34
	if start.Before(time.Date(2019, time.June, 12, 12, 0, 0, 0, time.UTC)) {
		metgridLevels = 32
	}
	if start.Before(time.Date(2016, time.May, 11, 12, 0, 0, 0, time.UTC)) {
		metgridLevels = 27
	}

	args.MetgridLevels = metgridLevels
	args.Hours = hours
	return args
}

// Args ...
type Args struct {
	Start, End time.Time
	Hours      int
}

// Tmpl ...
type Tmpl struct {
	TemplateContent string
}

// RenderTo ...
func (this *Tmpl) RenderTo(args Args, w io.Writer) {
	tmpl, err := template.New("namelist").Parse(this.TemplateContent)
	if err != nil {
		fmt.Printf("Error while parsing template: %s", err.Error())
		os.Exit(1)
	}

	tmplArgs := createTemplateArgs(args.Start, args.End, args.Hours)

	err = tmpl.Execute(w, tmplArgs)
	if err != nil {
		fmt.Printf("Error while evaluating template: %s", err.Error())
		os.Exit(1)
	}

}

// ReadTemplateFrom ...
func (this *Tmpl) ReadTemplateFrom(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Printf("Error while reading from stdin: %s", scanner.Err().Error())
		os.Exit(1)
	}

	this.TemplateContent = strings.Join(lines, "\n") + "\n"
}
