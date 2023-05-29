package main

import (
	"flag"
	"log"
	"os"

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

	var embed bool
	var showVersion bool
	flag.BoolVar(&embed, "embed", false, "embed resources")
	flag.BoolVar(&showVersion, "v", false, "Print the version")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// setup the static file server and get the controller
	r := gongdocx_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	gongdocxStage := gongdocx_fullstack.NewStackInstance(r, "github.com/fullstack-lang/gongdocx/go/models")

	for _, arg := range flag.Args() {
		gongdocx_models.NewDocx(gongdocxStage, arg, embed)
	}
	gongdocxStage.Commit()

	//
	gongdocx_models.ExtractStyleText("Titre1", gongdocxStage)

	gongdoc_load.Load(
		"gongdocx",
		"github.com/fullstack-lang/gongdocx/go/models",
		gongdocx_go.GoModelsDir,
		gongdocx_go.GoDiagramsDir,
		r,
		*embeddedDiagrams,
		&gongdocxStage.Map_GongStructName_InstancesNb)

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}
