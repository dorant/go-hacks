package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"text/template"

	"github.com/joncalhoun/pipe"
)

type data struct {
	Type   string
	Name   string
	Output string
}

func main() {
	var d data
	flag.StringVar(&d.Type, "type", "", "The subtype used for the queue being generated")
	flag.StringVar(&d.Name, "name", "", "The name used for the queue being generated. This should start with a capital letter so that it is exported.")
	flag.StringVar(&d.Output, "output", "", "Save result to output file")
	flag.Parse()

	t := template.Must(template.New("queue").Parse(queueTemplate))

	rc, wc, _ := pipe.Commands(
		exec.Command("gofmt"),
		exec.Command("goimports"),
	)
	t.Execute(wc, d)
	wc.Close()

	if d.Output == "" {
		log.Fatal("Outputfile missing, add -output=xx")
	}

	outFile, err := os.Create(d.Output)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	io.Copy(outFile, rc)
}

var queueTemplate = `
package queue

import (
  "container/list"
)

func New{{.Name}}() *{{.Name}} {
  return &{{.Name}}{list.New()}
}

type {{.Name}} struct {
  list *list.List
}

func (q *{{.Name}}) Len() int {
  return q.list.Len()
}

func (q *{{.Name}}) Enqueue(i {{.Type}}) {
  q.list.PushBack(i)
}

func (q *{{.Name}}) Dequeue() {{.Type}} {
  if q.list.Len() == 0 {
    panic(ErrEmptyQueue)
  }
  raw := q.list.Remove(q.list.Front())
  if typed, ok := raw.({{.Type}}); ok {
    return typed
  }
  panic(ErrInvalidType)
}
`
