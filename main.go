package main

import (
	"log"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/kyoh86/notifail/internal"
)

// nolint
var (
	version = "snapshot"
	commit  = "snapshot"
	date    = "snapshot"
)

func main() {
	app := kingpin.New("notifail", "Notify failure over OS notification drivers").Version(version).Author("kyoh86")
	var flag struct {
		Level          internal.Level
		Title          string
		Message        string
		Timeout        int
		DetailFilename string
	}
	app.Flag("level", "Notification Level").SetValue(&flag.Level)
	app.Flag("timeout", "Notification Timeout").Default("15").IntVar(&flag.Timeout)
	app.Flag("title", "Notification Title").StringVar(&flag.Title)
	app.Flag("detail-file", "Notification Detail Filename").StringVar(&flag.DetailFilename)
	app.Arg("message", "Notification Message").Required().StringVar(&flag.Message)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if err := internal.Notify(flag.Level, flag.Timeout, flag.Title, flag.Message, flag.DetailFilename); err != nil {
		log.Fatal(err)
	}
}
