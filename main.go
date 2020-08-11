package main

import (
	"fmt"
	"os"
	"time"

	"github.com/meteocima/namelist-prepare/namelist"
)

func parseArgs() (start, end time.Time, hours int) {
	startdateS := os.Args[1]
	enddateS := os.Args[2]

	startdate, err := time.Parse("2006010215", startdateS)
	if err != nil {
		fmt.Printf("Cannot parse startdate argument: %s. You must use YYYYMMDDHH format.", err.Error())
		os.Exit(1)
	}

	enddate, err := time.Parse("2006010215", enddateS)
	if err != nil {
		fmt.Printf("Cannot parse enddate argument: %s. You must use YYYYMMDDHH format.", err.Error())
		os.Exit(1)
	}

	lenghtHours := int(enddate.Sub(startdate) / time.Hour)

	return startdate, enddate, lenghtHours
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: cat templatefile.tmpl | namelist-prepare startdate enddate > file.out\n")
		fmt.Printf("\twhere startdate and enddate are in format YYYYMMDDHH\n\n")
		os.Exit(1)
	}

	startdate, enddate, hours := parseArgs()
	nlRenderer := namelist.Tmpl{}
	nlRenderer.ReadTemplateFrom(os.Stdin)
	nlRenderer.RenderTo(namelist.Args{
		Start: startdate,
		End:   enddate,
		Hours: hours,
	}, os.Stdout)
}
