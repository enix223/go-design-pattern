package main

import (
	"flag"
	"fmt"

	abstractfactory "github.com/enix223/go-design-pattern/abstract_factory"
	"github.com/enix223/go-design-pattern/adapter"
	"github.com/enix223/go-design-pattern/bridge"
	"github.com/enix223/go-design-pattern/builder"
	"github.com/enix223/go-design-pattern/composite"
	"github.com/enix223/go-design-pattern/decorator"
	"github.com/enix223/go-design-pattern/example"
	"github.com/enix223/go-design-pattern/facade"
	"github.com/enix223/go-design-pattern/factory"
	"github.com/enix223/go-design-pattern/flyweight"
	"github.com/enix223/go-design-pattern/prototype"
	"github.com/enix223/go-design-pattern/proxy"
	"github.com/enix223/go-design-pattern/singleton"
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
	case "bridge":
		e = &bridge.Example{}
	case "composite":
		e = &composite.Example{}
	case "decorator":
		e = &decorator.Example{}
	case "facade":
		e = &facade.Example{}
	case "flyweight":
		e = &flyweight.Example{}
	case "proxy":
		e = &proxy.Example{}
	default:
		fmt.Println("no pattern found")
		return
	}
	fmt.Println("---------------------------------------------")
	fmt.Printf("Running <%s> pattern example...\n", *pattern)
	fmt.Println("---------------------------------------------")
	e.Run()
}
