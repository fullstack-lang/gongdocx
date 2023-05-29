package models

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"
)

// Relationship is
type Relationship struct {
	Text       string `xml:",chardata"`
	ID         string `xml:"Id,attr"`
	Type       string `xml:"Type,attr"`
	Target     string `xml:"Target,attr"`
	TargetMode string `xml:"TargetMode,attr"`
}

// Relationships is
type Relationships struct {
	XMLName      xml.Name       `xml:"Relationships"`
	Text         string         `xml:",chardata"`
	Xmlns        string         `xml:"xmlns,attr"`
	Relationship []Relationship `xml:"Relationship"`
}

// TextVal is
type TextVal struct {
	Text string `xml:",chardata"`
	Val  string `xml:"val,attr"`
}

// NumberingLvl is
type NumberingLvl struct {
	Text      string  `xml:",chardata"`
	Ilvl      string  `xml:"ilvl,attr"`
	Tplc      string  `xml:"tplc,attr"`
	Tentative string  `xml:"tentative,attr"`
	Start     TextVal `xml:"start"`
	NumFmt    TextVal `xml:"numFmt"`
	LvlText   TextVal `xml:"lvlText"`
	LvlJc     TextVal `xml:"lvlJc"`
	PPr       struct {
		Text string `xml:",chardata"`
		Ind  struct {
			Text    string `xml:",chardata"`
			Left    string `xml:"left,attr"`
			Hanging string `xml:"hanging,attr"`
		} `xml:"ind"`
	} `xml:"pPr"`
	RPr struct {
		Text string `xml:",chardata"`
		U    struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"u"`
		RFonts struct {
			Text string `xml:",chardata"`
			Hint string `xml:"hint,attr"`
		} `xml:"rFonts"`
	} `xml:"rPr"`
}

// Numbering is
type Numbering struct {
	XMLName     xml.Name `xml:"numbering"`
	Text        string   `xml:",chardata"`
	Wpc         string   `xml:"wpc,attr"`
	Cx          string   `xml:"cx,attr"`
	Cx1         string   `xml:"cx1,attr"`
	Mc          string   `xml:"mc,attr"`
	O           string   `xml:"o,attr"`
	R           string   `xml:"r,attr"`
	M           string   `xml:"m,attr"`
	V           string   `xml:"v,attr"`
	Wp14        string   `xml:"wp14,attr"`
	Wp          string   `xml:"wp,attr"`
	W10         string   `xml:"w10,attr"`
	W           string   `xml:"w,attr"`
	W14         string   `xml:"w14,attr"`
	W15         string   `xml:"w15,attr"`
	W16se       string   `xml:"w16se,attr"`
	Wpg         string   `xml:"wpg,attr"`
	Wpi         string   `xml:"wpi,attr"`
	Wne         string   `xml:"wne,attr"`
	Wps         string   `xml:"wps,attr"`
	Ignorable   string   `xml:"Ignorable,attr"`
	AbstractNum []struct {
		Text                       string         `xml:",chardata"`
		AbstractNumID              string         `xml:"abstractNumId,attr"`
		RestartNumberingAfterBreak string         `xml:"restartNumberingAfterBreak,attr"`
		Nsid                       TextVal        `xml:"nsid"`
		MultiLevelType             TextVal        `xml:"multiLevelType"`
		Tmpl                       TextVal        `xml:"tmpl"`
		Lvl                        []NumberingLvl `xml:"lvl"`
	} `xml:"abstractNum"`
	Num []struct {
		Text          string  `xml:",chardata"`
		NumID         string  `xml:"numId,attr"`
		AbstractNumID TextVal `xml:"abstractNumId"`
	} `xml:"num"`
}

type file struct {
	rels  Relationships
	num   Numbering
	r     *zip.ReadCloser
	embed bool
	list  map[string]int
}

// Node_ is
type Node_ struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:"-"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Node_    `xml:",any"`
}

// UnmarshalXML is
func (n *Node_) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Attrs = start.Attr
	type node Node_

	return d.DecodeElement((*node)(n), &start)
}

func escape(s, set string) string {
	replacer := []string{}
	for _, r := range []rune(set) {
		rs := string(r)
		replacer = append(replacer, rs, `\`+rs)
	}
	return strings.NewReplacer(replacer...).Replace(s)
}

func (zf *file) extract(rel *Relationship, w io.Writer) error {
	err := os.MkdirAll(filepath.Dir(rel.Target), 0755)
	if err != nil {
		return err
	}
	for _, f := range zf.r.File {
		if f.Name != "word/"+rel.Target {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		b := make([]byte, f.UncompressedSize64)
		n, err := rc.Read(b)
		if err != nil && err != io.EOF {
			return err
		}
		if zf.embed {
			fmt.Fprintf(w, "![](data:image/png;base64,%s)",
				base64.StdEncoding.EncodeToString(b[:n]))
		} else {
			err = ioutil.WriteFile(rel.Target, b, 0644)
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "![](%s)", escape(rel.Target, "()"))
		}
		break
	}
	return nil
}

func attr(attrs []xml.Attr, name string) (string, bool) {
	for _, attr := range attrs {
		if attr.Name.Local == name {
			return attr.Value, true
		}
	}
	return "", false
}

func (zf *file) walk(
	node *Node,
	gongdocxStage *StageStruct,
	node_ *Node_,
	w io.Writer) error {

	var nodeCounter = 0

	switch node_.XMLName.Local {
	case "hyperlink":
		fmt.Fprint(w, "[")
		var cbuf bytes.Buffer
		for _, n := range node_.Nodes {
			if err := zf.walk(node, gongdocxStage, &n, &cbuf); err != nil {
				return err
			}
		}
		fmt.Fprint(w, escape(cbuf.String(), "[]"))
		fmt.Fprint(w, "]")

		fmt.Fprint(w, "(")
		if id, ok := attr(node_.Attrs, "id"); ok {
			for _, rel := range zf.rels.Relationship {
				if id == rel.ID {
					fmt.Fprint(w, escape(rel.Target, "()"))
					break
				}
			}
		}
		fmt.Fprint(w, ")")
	case "t":
		node__ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter)}).Stage(gongdocxStage)
		node.Nodes = append(node.Nodes, node__)
		nodeCounter = nodeCounter + 1

		text := (&Text{Name: node__.Name}).Stage(gongdocxStage)
		text.Node = node__
		text.Content = string(node_.Content)

		fmt.Fprint(w, string(node_.Content))
	case "pPr":
		code := false

		node__ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter)}).Stage(gongdocxStage)
		node.Nodes = append(node.Nodes, node__)
		nodeCounter = nodeCounter + 1

		paragraphproperties_ := (&ParagraphProperties{Name: node__.Name}).Stage(gongdocxStage)
		paragraphproperties_.Node = node__
		paragraphproperties_.Content = string(node_.Content)

		for _, n := range node_.Nodes {
			switch n.XMLName.Local {
			case "ind":
				if left, ok := attr(n.Attrs, "left"); ok {
					if i, err := strconv.Atoi(left); err == nil && i > 0 {
						fmt.Fprint(w, strings.Repeat("  ", i/360))
					}
				}
			case "pStyle":
				if val, ok := attr(n.Attrs, "val"); ok {
					if strings.HasPrefix(val, "Heading") {
						if i, err := strconv.Atoi(val[7:]); err == nil && i > 0 {
							fmt.Fprint(w, strings.Repeat("#", i)+" ")
						}
					} else if val == "Code" {
						code = true
					} else {
						if i, err := strconv.Atoi(val); err == nil && i > 0 {
							fmt.Fprint(w, strings.Repeat("#", i)+" ")
						}
					}
				}
			case "numPr":
				numID := ""
				ilvl := ""
				numFmt := ""
				start := 1
				ind := 0
				for _, nn := range n.Nodes {
					if nn.XMLName.Local == "numId" {
						if val, ok := attr(nn.Attrs, "val"); ok {
							numID = val
						}
					}
					if nn.XMLName.Local == "ilvl" {
						if val, ok := attr(nn.Attrs, "val"); ok {
							ilvl = val
						}
					}
				}
				for _, num := range zf.num.Num {
					if numID != num.NumID {
						continue
					}
					for _, abnum := range zf.num.AbstractNum {
						if abnum.AbstractNumID != num.AbstractNumID.Val {
							continue
						}
						for _, ablvl := range abnum.Lvl {
							if ablvl.Ilvl != ilvl {
								continue
							}
							if i, err := strconv.Atoi(ablvl.Start.Val); err == nil {
								start = i
							}
							if i, err := strconv.Atoi(ablvl.PPr.Ind.Left); err == nil {
								ind = i / 360
							}
							numFmt = ablvl.NumFmt.Val
							break
						}
						break
					}
					break
				}

				fmt.Fprint(w, strings.Repeat("  ", ind))
				switch numFmt {
				case "decimal", "aiueoFullWidth":
					key := fmt.Sprintf("%s:%d", numID, ind)
					cur, ok := zf.list[key]
					if !ok {
						zf.list[key] = start
					} else {
						zf.list[key] = cur + 1
					}
					fmt.Fprintf(w, "%d. ", zf.list[key])
				case "bullet":
					fmt.Fprint(w, "* ")
				}
			}
		}
		if code {
			fmt.Fprint(w, "`")
		}
		for _, n := range node_.Nodes {
			if err := zf.walk(node__, gongdocxStage, &n, w); err != nil {
				return err
			}
		}
		if code {
			fmt.Fprint(w, "`")
		}
	case "tbl":
		var rows [][]string
		for _, tr := range node_.Nodes {
			if tr.XMLName.Local != "tr" {
				continue
			}
			var cols []string
			for _, tc := range tr.Nodes {
				if tc.XMLName.Local != "tc" {
					continue
				}
				var cbuf bytes.Buffer
				if err := zf.walk(node, gongdocxStage, &tc, &cbuf); err != nil {
					return err
				}
				cols = append(cols, strings.Replace(cbuf.String(), "\n", "", -1))
			}
			rows = append(rows, cols)
		}
		maxcol := 0
		for _, cols := range rows {
			if len(cols) > maxcol {
				maxcol = len(cols)
			}
		}
		widths := make([]int, maxcol)
		for _, row := range rows {
			for i := 0; i < maxcol; i++ {
				if i < len(row) {
					width := runewidth.StringWidth(row[i])
					if widths[i] < width {
						widths[i] = width
					}
				}
			}
		}
		for i, row := range rows {
			if i == 0 {
				for j := 0; j < maxcol; j++ {
					fmt.Fprint(w, "|")
					fmt.Fprint(w, strings.Repeat(" ", widths[j]))
				}
				fmt.Fprint(w, "|\n")
				for j := 0; j < maxcol; j++ {
					fmt.Fprint(w, "|")
					fmt.Fprint(w, strings.Repeat("-", widths[j]))
				}
				fmt.Fprint(w, "|\n")
			}
			for j := 0; j < maxcol; j++ {
				fmt.Fprint(w, "|")
				if j < len(row) {
					width := runewidth.StringWidth(row[j])
					fmt.Fprint(w, escape(row[j], "|"))
					fmt.Fprint(w, strings.Repeat(" ", widths[j]-width))
				} else {
					fmt.Fprint(w, strings.Repeat(" ", widths[j]))
				}
			}
			fmt.Fprint(w, "|\n")
		}
		fmt.Fprint(w, "\n")
	case "r":

		node__ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter)}).Stage(gongdocxStage)
		node.Nodes = append(node.Nodes, node__)
		nodeCounter = nodeCounter + 1

		rune_ := (&Rune{Name: node__.Name}).Stage(gongdocxStage)
		rune_.Node = node__
		rune_.Content = string(node_.Content)

		bold := false
		italic := false
		strike := false
		for _, n := range node_.Nodes {
			if n.XMLName.Local != "rPr" {
				continue
			}

			nodeCounter_ := 0
			node___ := (&Node{Name: fmt.Sprintf(node__.Name+".%d", nodeCounter_)}).Stage(gongdocxStage)
			node.Nodes = append(node__.Nodes, node___)
			nodeCounter_ = nodeCounter_ + 1

			runeProperties := (&RuneProperties{Name: node__.Name}).Stage(gongdocxStage)
			runeProperties.Node = node__
			runeProperties.Content = string(n.Content)

			for _, nn := range n.Nodes {

				switch nn.XMLName.Local {
				case "b":
					bold = true
					runeProperties.IsBold = true
				case "i":
					italic = true
					runeProperties.IsItalic = true
				case "strike":
					strike = true
					runeProperties.IsStrike = true
				}
			}
		}
		if strike {
			fmt.Fprint(w, "~~")
		}
		if bold {
			fmt.Fprint(w, "**")
		}
		if italic {
			fmt.Fprint(w, "*")
		}
		var cbuf bytes.Buffer
		for _, n := range node_.Nodes {
			if err := zf.walk(node__, gongdocxStage, &n, &cbuf); err != nil {
				return err
			}
		}
		fmt.Fprint(w, escape(cbuf.String(), `*~\`))
		if italic {
			fmt.Fprint(w, "*")
		}
		if bold {
			fmt.Fprint(w, "**")
		}
		if strike {
			fmt.Fprint(w, "~~")
		}
	case "p":
		paragraph := (&Paragraph{Name: node.Name}).Stage(gongdocxStage)
		paragraph.Node = node
		paragraph.Content = string(node_.Content)

		for _, n := range node_.Nodes {

			node__ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter)}).Stage(gongdocxStage)
			node.Nodes = append(node.Nodes, node__)
			nodeCounter = nodeCounter + 1

			if err := zf.walk(node__, gongdocxStage, &n, w); err != nil {
				return err
			}
		}
		fmt.Fprintln(w)
	case "blip":
		if id, ok := attr(node_.Attrs, "embed"); ok {
			for _, rel := range zf.rels.Relationship {
				if id != rel.ID {
					continue
				}
				if err := zf.extract(&rel, w); err != nil {
					return err
				}
			}
		}
	case "Fallback":
	case "txbxContent":
		var cbuf bytes.Buffer
		for _, n := range node_.Nodes {
			if err := zf.walk(node, gongdocxStage, &n, &cbuf); err != nil {
				return err
			}
		}
		fmt.Fprintln(w, "\n```\n"+cbuf.String()+"```")
	default:
		for _, n := range node_.Nodes {
			node__ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter)}).Stage(gongdocxStage)
			node.Nodes = append(node.Nodes, node__)
			nodeCounter = nodeCounter + 1

			if err := zf.walk(node__, gongdocxStage, &n, w); err != nil {
				return err
			}
		}
	}

	return nil
}

func readFile(f *zip.File) (*Node_, error) {
	rc, err := f.Open()
	defer rc.Close()

	b, _ := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	var node Node_
	err = xml.Unmarshal(b, &node)
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func findFile(files []*zip.File, target string) *zip.File {
	for _, f := range files {
		if ok, _ := path.Match(target, f.Name); ok {
			return f
		}
	}
	return nil
}

func docx2md(docx *Docx, gongdocx_stage *StageStruct, arg string, embed bool) error {
	r, err := zip.OpenReader(arg)
	if err != nil {
		return err
	}
	defer r.Close()

	var rels Relationships
	var num Numbering

	for _, f := range r.File {

		file := (&File{Name: f.Name}).Stage(gongdocx_stage)
		docx.Files = append(docx.Files, file)

		switch f.Name {
		case "word/_rels/document.xml.rels":
			rc, err := f.Open()
			defer rc.Close()

			b, _ := ioutil.ReadAll(rc)
			if err != nil {
				return err
			}

			err = xml.Unmarshal(b, &rels)
			if err != nil {
				return err
			}
		case "word/numbering.xml":
			rc, err := f.Open()
			defer rc.Close()

			b, _ := ioutil.ReadAll(rc)
			if err != nil {
				return err
			}

			err = xml.Unmarshal(b, &num)
			if err != nil {
				return err
			}
		}
	}

	f := findFile(r.File, "word/document*.xml")
	if f == nil {
		return errors.New("incorrect document")
	}

	// match to the File
	document := (&Document{Name: f.Name}).Stage(gongdocx_stage)

	file_ := (*GetGongstructInstancesMap[File](gongdocx_stage))["word/document.xml"]
	document.File = file_

	node_, err := readFile(f)
	if err != nil {
		return err
	}
	node := (&Node{Name: "0"}).Stage(gongdocx_stage)
	document.Root = node

	var buf bytes.Buffer
	zf := &file{
		r:     r,
		rels:  rels,
		num:   num,
		embed: embed,
		list:  make(map[string]int),
	}
	err = zf.walk(node, gongdocx_stage, node_, &buf)
	if err != nil {
		return err
	}
	fmt.Print(buf.String())

	return nil
}
