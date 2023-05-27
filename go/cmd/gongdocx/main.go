package main

import (
	"flag"
	"log"

	gongdocx_go "github.com/fullstack-lang/gongdocx/go"
	gongdocx_fullstack "github.com/fullstack-lang/gongdocx/go/fullstack"
	gongdocx_models "github.com/fullstack-lang/gongdocx/go/models"
	gongdocx_static "github.com/fullstack-lang/gongdocx/go/static"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func main() {

	log.SetPrefix("gongdocx: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongdocx_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	var stage *gongdocx_models.StageStruct

	stage = gongdocx_fullstack.NewStackInstance(r, "github.com/fullstack-lang/gongdocx/go/models")

	gongdoc_load.Load(
		"gongdocx",
		"github.com/fullstack-lang/gongdocx/go/models",
		gongdocx_go.GoModelsDir,
		gongdocx_go.GoDiagramsDir,
		r,
		*embeddedDiagrams,
		&stage.Map_GongStructName_InstancesNb)

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}
