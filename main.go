package main

import (
	"flag"
	"fmt"

	abstractfactory "github.com/enix223/go-design-pattern/abstract_factory"
	"github.com/enix223/go-design-pattern/builder"
	"github.com/enix223/go-design-pattern/example"
	"github.com/enix223/go-design-pattern/factory"
	"github.com/enix223/go-design-pattern/prototype"
	"github.com/enix223/go-design-pattern/singleton"
	"github.com/enix223/go-design-pattern/adapter"
)

var pattern = flag.String("pattern", "", `design pattern to run, support folowing patterns:
	factory                  Factory method pattern
	abstract_factory         Abstract factory pattern
	builder                  Builder pattern
	singleton                Singleton pattern
	prototype                Prototype pattern
	adapter                	 Adapter pattern
`)

func main() {
	flag.Parse()

	var e example.PatternExample
	switch *pattern {
	case "factory":
		e = &factory.Example{}
	case "abstract_factory":
		e = &abstractfactory.Example{}
	case "builder":
		e = &builder.Example{}
	case "singleton":
		e = &singleton.Example{}
	case "prototype":
		e = &prototype.Example{}
	case "adapter":
		e = &adapter.Example{}
	default:
		fmt.Println("no pattern found")
		return
	}
	fmt.Println("---------------------------------------------")
	fmt.Printf("Running <%s> pattern example...\n", *pattern)
	fmt.Println("---------------------------------------------")
	e.Run()
}
