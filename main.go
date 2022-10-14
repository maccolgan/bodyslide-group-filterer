package main

import (
	"encoding/xml"
	"flag"
	"os"
	"regexp"
)

type Member struct {
	Name string `xml:"name,attr"`
}

type SliderGroups struct {
	XMLName xml.Name `xml:"SliderGroups"`
	Group   struct {
		Name   string   `xml:"name,attr"`
		Member []Member `xml:"Member"`
	} `xml:"Group"`
}

func main() {
	var input, output, regex string

	flag.StringVar(&input, "input", "input.xml", "File to read from")
	flag.StringVar(&output, "output", "output.xml", "Output file")
	flag.StringVar(&regex, "regex", ".*", "Regex to filter with")
	flag.Parse()

	realRegex := regexp.MustCompile(regex)

	inputBytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	var sliderGroups SliderGroups
	err = xml.Unmarshal(inputBytes, &sliderGroups)
	if err != nil {
		panic(err)
	}

	var newMembers []Member
	for _, mem := range sliderGroups.Group.Member {
		if realRegex.MatchString(mem.Name) {
			newMembers = append(newMembers, mem)
		}
	}
	sliderGroups.Group.Member = newMembers

	bytes, err := xml.MarshalIndent(sliderGroups, " ", "\t")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(output, bytes, 0666)
	if err != nil {
		panic(err)
	}
}
