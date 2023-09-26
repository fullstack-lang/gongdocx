package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	gongdocx_go "github.com/fullstack-lang/gongdocx/go"
	gongdocx_fullstack "github.com/fullstack-lang/gongdocx/go/fullstack"
	gongdocx_models "github.com/fullstack-lang/gongdocx/go/models"
	gongdocx_orm "github.com/fullstack-lang/gongdocx/go/orm"
	gongdocx_probe "github.com/fullstack-lang/gongdocx/go/probe"
	gongdocx_static "github.com/fullstack-lang/gongdocx/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")

	style = flag.String("style", "", "style for extracting instances")

	xls = flag.String("xls", "", "output file")
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
	var stage *gongdocx_models.StageStruct
	var backRepo *gongdocx_orm.BackRepoStruct

	stage, backRepo = gongdocx_fullstack.NewStackInstance(r, "gongdocx")

	for _, arg := range flag.Args() {
		gongdocx_models.NewDocx(stage, arg, embed)
	}

	if *xls != "" {
		fileName := fmt.Sprintf("%s_%s.xlsx", *xls, time.Now().Local().Format("2006-01-02-15-04"))
		gongdocx_models.SerializeStage(stage, fileName)
	}

	stage.Commit()

	if *style != "" {
		instances := gongdocx_models.ExtractStyleText(*style, stage)
		log.Println("Instances of ", *style)

		for _, instance := range instances {
			log.Println(instance)
		}

	}

	gongdocx_probe.NewProbe(r, gongdocx_go.GoModelsDir, gongdocx_go.GoDiagramsDir,
		*embeddedDiagrams, "gongdocx", stage, backRepo)

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
