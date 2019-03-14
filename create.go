package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"regexp"
	"os"
	"strings"

	"github.com/google/go-github/github"
)

var (
	todo = regexp.MustCompile("TODO.*:")
	bug  = regexp.MustCompile("BUG.*:")
	ver  = regexp.MustCompile(`(v[0-9]\.[0-9]\.[0-9])`)
	assign = `\s+((?:[^,\s]*(?:,\s)?)*)`
)

func createIssue(ctx context.Context, is *github.IssuesService) {
	var labels []string
	var assignees []string

	scanner := bufio.NewScanner(os.Stdin)
	if ! scanner.Scan() {
		log.Fatal(scanner.Err())
	}
	input := scanner.Text()
	input = cleanComments(input) + "\n"
	switch {
	case todo.MatchString(input):
		labels = append(labels, "todo")
		t := regexp.MustCompile("TODO" + assign + ":")
		h := t.FindStringSubmatch(input)
		if len(h) > 0 {
			c := csv.NewReader(strings.NewReader(h[1]))
			assignees, _= c.Read()
		}
		var size int
		if len(assignees) > 0 {
			size = len(t.FindString(input))
		} else {
			size = len(todo.FindString(input))
		}
		input = input[size + 1:]
	case bug.MatchString(input):
		labels = append(labels, "bug")
		t := regexp.MustCompile(`BUG\(([^,\s]+)\)`)
		h := t.FindStringSubmatch(input)
		if len(h) > 0 {
			c := csv.NewReader(strings.NewReader(h[1]))
			assignees, _ = c.Read()
		}
		var size int
		if len(assignees) > 0 {
			size = len(t.FindString(input))
		} else {
			size = len(bug.FindString(input))
		}
		input = input[size + 1:]
	}
	for scanner.Scan() {
		line := scanner.Text()
		line = cleanComments(line)
		if line != "" {
			input += line + "\n"
		}
	}
	if ver.MatchString(input) {
		h := ver.FindStringSubmatch(input)
		if len(h) > 0 {
			labels = append(labels, h[1])
		}
	}	
	ir := &github.IssueRequest{
		Title: title,
		Body: &input,
		Labels: &labels,
		Assignees: &assignees,
	}
	issue, _, err := is.Create(ctx, *username, os.Args[1], ir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(issue.URL)

}

// Just let the library do all the lifting
func cleanComments(input string) string {
	input = strings.TrimSpace(input)
	input = strings.TrimPrefix(input, "//")
	input = strings.TrimPrefix(input, "/*")
	input = strings.TrimSuffix(input, "*/")
	input = strings.TrimPrefix(input, "#")
	input = strings.TrimSpace(input)
	return input
}