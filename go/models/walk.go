package models

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"
)

func walk(
	zf *file,
	parentNode any,
	node *Node,
	gongdocxStage *StageStruct,
	node_ *Node_,
	w io.Writer) error {

	var nodeCounter = 0
	var dummyNode *Node

	switch node_.XMLName.Local {
	case "hyperlink":
		fmt.Fprint(w, "[")
		var cbuf bytes.Buffer
		for _, n := range node_.Nodes {
			if err := walk(zf, dummyNode, node, gongdocxStage, &n, &cbuf); err != nil {
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

		// check if the parent node is a rune
		switch rune_ := parentNode.(type) {
		case *Rune:
			rune_.Text = text
		}

		fmt.Fprint(w, string(node_.Content))
	case "pPr":
		code := false

		paragraphProperties := (&ParagraphProperties{Name: node.Name}).Stage(gongdocxStage)
		paragraphProperties.Node = node
		paragraphProperties.Content = string(node_.Content)

		// check if the parent node is a paragraph
		switch paragraph := parentNode.(type) {
		case *Paragraph:
			paragraph.ParagraphProperties = paragraphProperties
		}

		for _, n := range node_.Nodes {

			switch n.XMLName.Local {
			case "ind":
				if left, ok := attr(n.Attrs, "left"); ok {
					if i, err := strconv.Atoi(left); err == nil && i > 0 {
						fmt.Fprint(w, strings.Repeat("  ", i/360))
					}
				}
			case "pStyle":

				nodeCounter_ := 0
				node___ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter_)}).Stage(gongdocxStage)
				node.Nodes = append(node.Nodes, node___)
				nodeCounter_ = nodeCounter_ + 1

				paragraphStyle := (&ParagraphStyle{Name: node___.Name}).Stage(gongdocxStage)
				paragraphStyle.Node = node
				paragraphStyle.Content = string(n.Content)
				paragraphProperties.ParagraphStyle = paragraphStyle

				if val, ok := attr(n.Attrs, "val"); ok {

					paragraphStyle.ValAttr = val
					paragraphStyle.Name = paragraphStyle.Name + " " + paragraphStyle.ValAttr

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
			if err := walk(zf, dummyNode, node, gongdocxStage, &n, w); err != nil {
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
				if err := walk(zf, dummyNode, node, gongdocxStage, &tc, &cbuf); err != nil {
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

		rune_ := (&Rune{Name: node.Name}).Stage(gongdocxStage)
		rune_.Node = node
		rune_.Content = string(node_.Content)

		// check if the parent node is a paragraph
		switch paragraph := parentNode.(type) {
		case *Paragraph:
			paragraph.Runes = append(paragraph.Runes, rune_)
		}

		bold := false
		italic := false
		strike := false
		for _, n := range node_.Nodes {
			if n.XMLName.Local != "rPr" {
				continue
			}

			nodeCounter_ := 0
			node___ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter_)}).Stage(gongdocxStage)
			node.Nodes = append(node.Nodes, node___)
			nodeCounter_ = nodeCounter_ + 1

			runeProperties := (&RuneProperties{Name: node___.Name}).Stage(gongdocxStage)
			runeProperties.Node = node___
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
			if err := walk(zf, rune_, node, gongdocxStage, &n, &cbuf); err != nil {
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

			if err := walk(zf, paragraph, node__, gongdocxStage, &n, w); err != nil {
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
			if err := walk(zf, dummyNode, node, gongdocxStage, &n, &cbuf); err != nil {
				return err
			}
		}
		fmt.Fprintln(w, "\n```\n"+cbuf.String()+"```")
	default:
		for _, n := range node_.Nodes {
			node__ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter)}).Stage(gongdocxStage)
			node.Nodes = append(node.Nodes, node__)
			nodeCounter = nodeCounter + 1

			if err := walk(zf, dummyNode, node__, gongdocxStage, &n, w); err != nil {
				return err
			}
		}
	}

	return nil
}
