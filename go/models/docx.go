package models

import (
	"log"
	"path/filepath"
)

type Docx struct {
	Name string
}

func NewDocx(gongdocx_stage *StageStruct, path string, embed bool) (docx *Docx) {

	docx = (&Docx{Name: filepath.Base(path)}).Stage(gongdocx_stage)

	if err := Docx2md(path, embed); err != nil {
		log.Fatal(err)
	}

	return
}
