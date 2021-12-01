// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 115.

// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"log"
	"os"

	"gopl.io/ch4/github"
)

//!+template
import "html/template"

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

//!-template

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

//!-
/*
//$ go build gopl.io--/ch4/issueshtml
//$ ./issueshtml repo:golang/go is:open json decoder
<h1>57 issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>

<tr>
  <td><a href='https://github.com/golang/go/pull/43716'>43716</a></td>
  <td>open</td>
  <td><a href='https://github.com/ggaaooppeenngg'>ggaaooppeenngg</a></td>
  <td><a href='https://github.com/golang/go/pull/43716'>encoding/json: fix byte counter increments when using decoder.Token()</a></td>
</tr>

<tr>
  <td><a href='https://github.com/golang/go/pull/33416'>33416</a></td>
  <td>open</td>
  <td><a href='https://github.com/bserdar'>bserdar</a></td>
  <td><a href='https://github.com/golang/go/pull/33416'>encoding/json: This CL adds Decoder.InternKeys</a></td>
</tr>

*/
