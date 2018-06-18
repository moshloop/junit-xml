package main

import (
	"flag"
	"os"
	"time"

	"github.com/beevik/etree"
	"io/ioutil"
)

func usage() {
	println(`Usage: junit-xml [OPTIONS] pass|fail|skip class name FILE

Utility to append testing results to a Junit formatted XML file
Options:`)
	flag.PrintDefaults()

}

var (
	millis  time.Duration
	message string
)

func main() {

	flag.DurationVar(&millis, "time", 0*time.Nanosecond, "Duration of test in milliseconds")
	flag.StringVar(&message, "message", "", "Failure message")
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 && flag.NFlag() == 0 {
		usage()
		os.Exit(1)
	}

	action := flag.Arg(0)
	class := flag.Arg(1)
	name := flag.Arg(2)

	file := flag.Arg(3)

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(file); err != nil {
		panic(err)
	}

	testcase := doc.FindElement("./testsuites/testsuite").CreateElement("testcase")
	testcase.CreateAttr("classname", class)
	testcase.CreateAttr("name", name)
	if action == "fail" {
		testcase.CreateElement("failure").CreateAttr("message", message)
	}

	if millis.Nanoseconds() > 0 {
		testcase.CreateAttr("time", millis.String())
	}
	doc.Indent(2)
	output, err := doc.WriteToBytes()
	if err != nil {
		panic(err)
	}

	info, _ := os.Stat(file)

	ioutil.WriteFile(file, output, info.Mode())
}
