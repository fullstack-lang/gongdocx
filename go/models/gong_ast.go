// generated code - do not edit
package models

import (
	"bufio"
	"errors"
	"go/ast"
	"go/doc/comment"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError
var dummy_time_import time.Time

// swagger:ignore
type GONG__ExpressionType string

const (
	GONG__STRUCT_INSTANCE      GONG__ExpressionType = "STRUCT_INSTANCE"
	GONG__FIELD_OR_CONST_VALUE GONG__ExpressionType = "FIELD_OR_CONST_VALUE"
	GONG__FIELD_VALUE          GONG__ExpressionType = "FIELD_VALUE"
	GONG__ENUM_CAST_INT        GONG__ExpressionType = "ENUM_CAST_INT"
	GONG__ENUM_CAST_STRING     GONG__ExpressionType = "ENUM_CAST_STRING"
	GONG__IDENTIFIER_CONST     GONG__ExpressionType = "IDENTIFIER_CONST"
)

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(stage *StageStruct, pathToFile string) error {

	ReplaceOldDeclarationsInFile(pathToFile)

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	// startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	// log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileFromAst(stage *StageStruct, inFile *ast.File, fset *token.FileSet) error {
	// if there is a meta package import, it is the third import
	if len(inFile.Imports) > 3 {
		log.Fatalln("Too many imports in file", inFile.Name)
	}
	if len(inFile.Imports) == 3 {
		stage.MetaPackageImportAlias = inFile.Imports[2].Name.Name
		stage.MetaPackageImportPath = inFile.Imports[2].Path.Value
	}

	// astCoordinate := "File "
	// log.Println(// astCoordinate)
	for _, decl := range inFile.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl
			// astCoordinate := // astCoordinate + "\tFunction " + funcDecl.Name.Name
			if name := funcDecl.Name; name != nil {
				isOfInterest := strings.Contains(funcDecl.Name.Name, "_")
				if !isOfInterest {
					continue
				}
				// log.Println(// astCoordinate)
			}
			if doc := funcDecl.Doc; doc != nil {
				// astCoordinate := // astCoordinate + "\tDoc"
				for _, comment := range doc.List {
					_ = comment
					// astCoordinate := // astCoordinate + "\tComment: " + comment.Text
					// log.Println(// astCoordinate)
				}
			}
			if body := funcDecl.Body; body != nil {
				// astCoordinate := // astCoordinate + "\tBody: "
				for _, stmt := range body.List {
					switch stmt := stmt.(type) {
					case *ast.ExprStmt:
						exprStmt := stmt
						// astCoordinate := // astCoordinate + "\tExprStmt: "
						switch expr := exprStmt.X.(type) {
						case *ast.CallExpr:
							// astCoordinate := // astCoordinate + "\tCallExpr: "
							callExpr := expr
							switch fun := callExpr.Fun.(type) {
							case *ast.Ident:
								ident := fun
								_ = ident
								// astCoordinate := // astCoordinate + "\tIdent: " + ident.Name
								// log.Println(// astCoordinate)
							}
						}
					case *ast.AssignStmt:
						// Create an ast.CommentMap from the ast.File's comments.
						// This helps keeping the association between comments
						// and AST nodes.
						cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName :=
							UnmarshallGongstructStaging(
								stage, &cmap, assignStmt, astCoordinate)
						_ = instance
						_ = id
						_ = gongstruct
						_ = fieldName
					}
				}
			}
		case *ast.GenDecl:
			genDecl := decl
			// log.Println("\tAST GenDecl: ")
			if doc := genDecl.Doc; doc != nil {
				for _, comment := range doc.List {
					_ = comment
					// log.Println("\tAST Comment: ", comment.Text)
				}
			}
			for _, spec := range genDecl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec
					if path := importSpec.Path; path != nil {
						// log.Println("\t\tAST Path: ", path.Value)
					}
				case *ast.ValueSpec:
					ident := spec.Names[0]
					_ = ident
					if !strings.HasPrefix(ident.Name, "_") {
						continue
					}
					// declaration of a variable without initial value
					if len(spec.Values) == 0 {
						continue
					}
					switch compLit := spec.Values[0].(type) {
					case *ast.CompositeLit:
						var key string
						_ = key
						var value string
						_ = value
						for _, elt := range compLit.Elts {

							// each elt is an expression for struct or for field such as
							// for struct
							//
							//         "dummy.Dummy": &(dummy.Dummy{})
							//
							// or, for field
							//
							//          "dummy.Dummy.Name": (dummy.Dummy{}).Name,
							//
							// first node in the AST is a key value expression
							var ok bool
							var kve *ast.KeyValueExpr
							if kve, ok = elt.(*ast.KeyValueExpr); !ok {
								log.Fatal("Expression should be key value expression" +
									fset.Position(kve.Pos()).String())
							}

							switch bl := kve.Key.(type) {
							case *ast.BasicLit:
								key = bl.Value // "\"dumm.Dummy\"" is the value

								// one remove the ambracing double quotes
								key = strings.TrimPrefix(key, "\"")
								key = strings.TrimSuffix(key, "\"")
							}
							var expressionType GONG__ExpressionType = GONG__STRUCT_INSTANCE
							var docLink GONG__Identifier

							var fieldName string
							var ue *ast.UnaryExpr
							if ue, ok = kve.Value.(*ast.UnaryExpr); !ok {
								expressionType = GONG__FIELD_OR_CONST_VALUE
							}

							var callExpr *ast.CallExpr
							if callExpr, ok = kve.Value.(*ast.CallExpr); ok {

								var se *ast.SelectorExpr
								if se, ok = callExpr.Fun.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector expression" +
										fset.Position(callExpr.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}

								// check the arg type to select wether this is a int or a string enum
								var bl *ast.BasicLit
								if bl, ok = callExpr.Args[0].(*ast.BasicLit); ok {
									switch bl.Kind {
									case token.STRING:
										expressionType = GONG__ENUM_CAST_STRING
									case token.INT:
										expressionType = GONG__ENUM_CAST_INT
									}
								} else {
									log.Fatal("Expression should be a basic lit" +
										fset.Position(se.Pos()).String())
								}

								docLink.Ident = id.Name + "." + se.Sel.Name
								_ = callExpr
							}

							var se2 *ast.SelectorExpr
							switch expressionType {
							case GONG__FIELD_OR_CONST_VALUE:
								if se2, ok = kve.Value.(*ast.SelectorExpr); ok {

									var ident *ast.Ident
									if _, ok = se2.X.(*ast.ParenExpr); ok {
										expressionType = GONG__FIELD_VALUE
										fieldName = se2.Sel.Name
									} else if ident, ok = se2.X.(*ast.Ident); ok {
										expressionType = GONG__IDENTIFIER_CONST
										docLink.Ident = ident.Name + "." + se2.Sel.Name
									} else {
										log.Fatal("Expression should be a selector expression or an ident" +
											fset.Position(kve.Pos()).String())
									}
								} else {

								}
							}

							var pe *ast.ParenExpr
							switch expressionType {
							case GONG__STRUCT_INSTANCE:
								if pe, ok = ue.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							case GONG__FIELD_VALUE:
								if pe, ok = se2.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							}
							switch expressionType {
							case GONG__FIELD_VALUE, GONG__STRUCT_INSTANCE:
								// expect a Composite Litteral with no Element <type>{}
								var cl *ast.CompositeLit
								if cl, ok = pe.X.(*ast.CompositeLit); !ok {
									log.Fatal("Expression should be a composite lit" +
										fset.Position(pe.Pos()).String())
								}

								var se *ast.SelectorExpr
								if se, ok = cl.Type.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector" +
										fset.Position(cl.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}
								docLink.Ident = id.Name + "." + se.Sel.Name
							}

							switch expressionType {
							case GONG__FIELD_VALUE:
								docLink.Ident += "." + fieldName
							}

							// if map_DocLink_Identifier has the same ident, this means
							// that no renaming has occured since the last processing of the
							// file. But it is neccessary to keep it in memory for the
							// marshalling
							if docLink.Ident == key {
								// continue
							}

							// otherwise, one stores the new ident (after renaming) in the
							// renaming map
							docLink.Type = expressionType
							stage.Map_DocLink_Renaming[key] = docLink
						}
					}
				}
			}
		}

	}
	return nil
}

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_Body = make(map[string]*Body)
var __gong__map_Document = make(map[string]*Document)
var __gong__map_Docx = make(map[string]*Docx)
var __gong__map_File = make(map[string]*File)
var __gong__map_Node = make(map[string]*Node)
var __gong__map_Paragraph = make(map[string]*Paragraph)
var __gong__map_ParagraphProperties = make(map[string]*ParagraphProperties)
var __gong__map_ParagraphStyle = make(map[string]*ParagraphStyle)
var __gong__map_Rune = make(map[string]*Rune)
var __gong__map_RuneProperties = make(map[string]*RuneProperties)
var __gong__map_Table = make(map[string]*Table)
var __gong__map_TableColumn = make(map[string]*TableColumn)
var __gong__map_TableProperties = make(map[string]*TableProperties)
var __gong__map_TableRow = make(map[string]*TableRow)
var __gong__map_TableStyle = make(map[string]*TableStyle)
var __gong__map_Text = make(map[string]*Text)

// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
// to be recognized as a proper identifier.
// While this was introduced in go 1.19, it is not yet implemented in
// gopls (see [issue](https://github.com/golang/go/issues/57559)
func lookupPackage(name string) (importPath string, ok bool) {
	return name, true
}
func lookupSym(recv, name string) (ok bool) {
	if recv == "" {
		return true
	}
	return false
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(stage *StageStruct, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {

	// used for debug purposes
	astCoordinate := "\tAssignStmt: "

	//
	// First parse all comment groups in the assignement
	// if a comment "//gong:ident [DocLink]" is met and is followed by a string assignement.
	// modify the following AST assignement to assigns the DocLink text to the string value
	//

	// get the comment group of the assigStmt
	commentGroups := (*cmap)[assignStmt]
	// get the the prefix
	var hasGongIdentDirective bool
	var commentText string
	var docLinkText string
	for _, commentGroup := range commentGroups {
		for _, comment := range commentGroup.List {
			if strings.HasPrefix(comment.Text, "//gong:ident") {
				hasGongIdentDirective = true
				commentText = comment.Text
			}
		}
	}
	if hasGongIdentDirective {
		// parser configured to find doclinks
		var docLinkFinder comment.Parser
		docLinkFinder.LookupPackage = lookupPackage
		docLinkFinder.LookupSym = lookupSym
		doc := docLinkFinder.Parse(commentText)

		for _, block := range doc.Content {
			switch paragraph := block.(type) {
			case *comment.Paragraph:
				_ = paragraph
				for _, text := range paragraph.Text {
					switch docLink := text.(type) {
					case *comment.DocLink:
						if docLink.Recv == "" {
							docLinkText = docLink.ImportPath + "." + docLink.Name
						} else {
							docLinkText = docLink.ImportPath + "." + docLink.Recv + "." + docLink.Name
						}

						// we check wether the doc link has been renamed
						// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
						if renamed, ok := (stage.Map_DocLink_Renaming)[docLinkText]; ok {
							docLinkText = renamed.Ident
						}
					}
				}
			}
		}
	}

	for rank, expr := range assignStmt.Lhs {
		if rank > 0 {
			continue
		}
		switch expr := expr.(type) {
		case *ast.Ident:
			// we are on a variable declaration
			ident := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			// log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			// log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		// astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			// astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				// astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					// astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						// astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							// astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								// astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									// astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										_ = ident
										// astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										// log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										// astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										// log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							_ = astCoordinate2
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								// astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									_ = ident
									// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									// log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									// log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers
									case "Body":
										instanceBody := new(Body)
										instanceBody.Name = instanceName
										instanceBody.Stage(stage)
										instance = any(instanceBody)
										__gong__map_Body[identifier] = instanceBody
									case "Document":
										instanceDocument := new(Document)
										instanceDocument.Name = instanceName
										instanceDocument.Stage(stage)
										instance = any(instanceDocument)
										__gong__map_Document[identifier] = instanceDocument
									case "Docx":
										instanceDocx := new(Docx)
										instanceDocx.Name = instanceName
										instanceDocx.Stage(stage)
										instance = any(instanceDocx)
										__gong__map_Docx[identifier] = instanceDocx
									case "File":
										instanceFile := new(File)
										instanceFile.Name = instanceName
										instanceFile.Stage(stage)
										instance = any(instanceFile)
										__gong__map_File[identifier] = instanceFile
									case "Node":
										instanceNode := new(Node)
										instanceNode.Name = instanceName
										instanceNode.Stage(stage)
										instance = any(instanceNode)
										__gong__map_Node[identifier] = instanceNode
									case "Paragraph":
										instanceParagraph := new(Paragraph)
										instanceParagraph.Name = instanceName
										instanceParagraph.Stage(stage)
										instance = any(instanceParagraph)
										__gong__map_Paragraph[identifier] = instanceParagraph
									case "ParagraphProperties":
										instanceParagraphProperties := new(ParagraphProperties)
										instanceParagraphProperties.Name = instanceName
										instanceParagraphProperties.Stage(stage)
										instance = any(instanceParagraphProperties)
										__gong__map_ParagraphProperties[identifier] = instanceParagraphProperties
									case "ParagraphStyle":
										instanceParagraphStyle := new(ParagraphStyle)
										instanceParagraphStyle.Name = instanceName
										instanceParagraphStyle.Stage(stage)
										instance = any(instanceParagraphStyle)
										__gong__map_ParagraphStyle[identifier] = instanceParagraphStyle
									case "Rune":
										instanceRune := new(Rune)
										instanceRune.Name = instanceName
										instanceRune.Stage(stage)
										instance = any(instanceRune)
										__gong__map_Rune[identifier] = instanceRune
									case "RuneProperties":
										instanceRuneProperties := new(RuneProperties)
										instanceRuneProperties.Name = instanceName
										instanceRuneProperties.Stage(stage)
										instance = any(instanceRuneProperties)
										__gong__map_RuneProperties[identifier] = instanceRuneProperties
									case "Table":
										instanceTable := new(Table)
										instanceTable.Name = instanceName
										instanceTable.Stage(stage)
										instance = any(instanceTable)
										__gong__map_Table[identifier] = instanceTable
									case "TableColumn":
										instanceTableColumn := new(TableColumn)
										instanceTableColumn.Name = instanceName
										instanceTableColumn.Stage(stage)
										instance = any(instanceTableColumn)
										__gong__map_TableColumn[identifier] = instanceTableColumn
									case "TableProperties":
										instanceTableProperties := new(TableProperties)
										instanceTableProperties.Name = instanceName
										instanceTableProperties.Stage(stage)
										instance = any(instanceTableProperties)
										__gong__map_TableProperties[identifier] = instanceTableProperties
									case "TableRow":
										instanceTableRow := new(TableRow)
										instanceTableRow.Name = instanceName
										instanceTableRow.Stage(stage)
										instance = any(instanceTableRow)
										__gong__map_TableRow[identifier] = instanceTableRow
									case "TableStyle":
										instanceTableStyle := new(TableStyle)
										instanceTableStyle.Name = instanceName
										instanceTableStyle.Stage(stage)
										instance = any(instanceTableStyle)
										__gong__map_TableStyle[identifier] = instanceTableStyle
									case "Text":
										instanceText := new(Text)
										instanceText.Name = instanceName
										instanceText.Stage(stage)
										instance = any(instanceText)
										__gong__map_Text[identifier] = instanceText
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					// astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					// log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					// astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						// log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Fatalln("gongstructName not found for identifier", identifier)
						}
						switch gongstructName {
						// insertion point for basic lit assignments
						case "Body":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Document":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Docx":
							switch fieldName {
							// insertion point for date assign code
							}
						case "File":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Node":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Paragraph":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ParagraphProperties":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ParagraphStyle":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Rune":
							switch fieldName {
							// insertion point for date assign code
							}
						case "RuneProperties":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Table":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TableColumn":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TableProperties":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TableRow":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TableStyle":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Text":
							switch fieldName {
							// insertion point for date assign code
							}
						}
					}
				}
			case *ast.Ident:
				// append function
				ident := fun
				_ = ident
				// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			for _, arg := range callExpr.Args {
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident:
					ident := arg
					_ = ident
					// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
					// log.Println(astCoordinate)
					var ok bool
					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Fatalln("gongstructName not found for identifier", identifier)
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "Body":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Paragraphs":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Paragraph[targetIdentifier]
							__gong__map_Body[identifier].Paragraphs =
								append(__gong__map_Body[identifier].Paragraphs, target)
						case "Tables":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Table[targetIdentifier]
							__gong__map_Body[identifier].Tables =
								append(__gong__map_Body[identifier].Tables, target)
						}
					case "Document":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Docx":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Files":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_File[targetIdentifier]
							__gong__map_Docx[identifier].Files =
								append(__gong__map_Docx[identifier].Files, target)
						}
					case "File":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Node":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Nodes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Node[targetIdentifier]
							__gong__map_Node[identifier].Nodes =
								append(__gong__map_Node[identifier].Nodes, target)
						}
					case "Paragraph":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Runes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Rune[targetIdentifier]
							__gong__map_Paragraph[identifier].Runes =
								append(__gong__map_Paragraph[identifier].Runes, target)
						}
					case "ParagraphProperties":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ParagraphStyle":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Rune":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "RuneProperties":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Table":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "TableRows":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_TableRow[targetIdentifier]
							__gong__map_Table[identifier].TableRows =
								append(__gong__map_Table[identifier].TableRows, target)
						}
					case "TableColumn":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Paragraphs":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Paragraph[targetIdentifier]
							__gong__map_TableColumn[identifier].Paragraphs =
								append(__gong__map_TableColumn[identifier].Paragraphs, target)
						}
					case "TableProperties":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "TableRow":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "TableColumns":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_TableColumn[targetIdentifier]
							__gong__map_TableRow[identifier].TableColumns =
								append(__gong__map_TableRow[identifier].TableColumns, target)
						}
					case "TableStyle":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Text":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					}
				case *ast.SelectorExpr:
					slcExpr := arg
					// astCoordinate := astCoordinate + "\tSelectorExpr"
					switch X := slcExpr.X.(type) {
					case *ast.Ident:
						ident := X
						_ = ident
						// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
						// log.Println(astCoordinate)

					}
					if Sel := slcExpr.Sel; Sel != nil {
						// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
						// log.Println(astCoordinate)
					}
				}
			}
		case *ast.BasicLit, *ast.UnaryExpr:

			var basicLit *ast.BasicLit
			var exprSign = 1.0
			_ = exprSign // in case this is not used

			if bl, ok := expr.(*ast.BasicLit); ok {
				// expression is  for instance ... = 18.000
				basicLit = bl
			} else if ue, ok := expr.(*ast.UnaryExpr); ok {
				// expression is  for instance ... = -18.000
				// we want to extract a *ast.BasicLit from the *ast.UnaryExpr
				basicLit = ue.X.(*ast.BasicLit)
				exprSign = -1
			}

			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			// substitute the RHS part of the assignment if a //gong:ident directive is met
			if hasGongIdentDirective {
				basicLit.Value = "[" + docLinkText + "]"
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "Body":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Body[identifier].Name = fielValue
				}
			case "Document":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Document[identifier].Name = fielValue
				}
			case "Docx":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Docx[identifier].Name = fielValue
				}
			case "File":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_File[identifier].Name = fielValue
				}
			case "Node":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Node[identifier].Name = fielValue
				}
			case "Paragraph":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Paragraph[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Paragraph[identifier].Content = fielValue
				case "Text":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Paragraph[identifier].Text = fielValue
				}
			case "ParagraphProperties":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ParagraphProperties[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ParagraphProperties[identifier].Content = fielValue
				}
			case "ParagraphStyle":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ParagraphStyle[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ParagraphStyle[identifier].Content = fielValue
				case "ValAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ParagraphStyle[identifier].ValAttr = fielValue
				}
			case "Rune":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rune[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rune[identifier].Content = fielValue
				}
			case "RuneProperties":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RuneProperties[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RuneProperties[identifier].Content = fielValue
				}
			case "Table":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Table[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Table[identifier].Content = fielValue
				}
			case "TableColumn":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TableColumn[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TableColumn[identifier].Content = fielValue
				}
			case "TableProperties":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TableProperties[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TableProperties[identifier].Content = fielValue
				}
			case "TableRow":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TableRow[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TableRow[identifier].Content = fielValue
				}
			case "TableStyle":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TableStyle[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TableStyle[identifier].Content = fielValue
				case "Val":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TableStyle[identifier].Val = fielValue
				}
			case "Text":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].Content = fielValue
				}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			_ = ident
			// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments
			case "Body":
				switch fieldName {
				// insertion point for field dependant code
				case "LastParagraph":
					targetIdentifier := ident.Name
					__gong__map_Body[identifier].LastParagraph = __gong__map_Paragraph[targetIdentifier]
				}
			case "Document":
				switch fieldName {
				// insertion point for field dependant code
				case "File":
					targetIdentifier := ident.Name
					__gong__map_Document[identifier].File = __gong__map_File[targetIdentifier]
				case "Root":
					targetIdentifier := ident.Name
					__gong__map_Document[identifier].Root = __gong__map_Node[targetIdentifier]
				case "Body":
					targetIdentifier := ident.Name
					__gong__map_Document[identifier].Body = __gong__map_Body[targetIdentifier]
				}
			case "Docx":
				switch fieldName {
				// insertion point for field dependant code
				case "Document":
					targetIdentifier := ident.Name
					__gong__map_Docx[identifier].Document = __gong__map_Document[targetIdentifier]
				}
			case "File":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Node":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Paragraph":
				switch fieldName {
				// insertion point for field dependant code
				case "Node":
					targetIdentifier := ident.Name
					__gong__map_Paragraph[identifier].Node = __gong__map_Node[targetIdentifier]
				case "ParagraphProperties":
					targetIdentifier := ident.Name
					__gong__map_Paragraph[identifier].ParagraphProperties = __gong__map_ParagraphProperties[targetIdentifier]
				case "Next":
					targetIdentifier := ident.Name
					__gong__map_Paragraph[identifier].Next = __gong__map_Paragraph[targetIdentifier]
				case "Previous":
					targetIdentifier := ident.Name
					__gong__map_Paragraph[identifier].Previous = __gong__map_Paragraph[targetIdentifier]
				case "EnclosingBody":
					targetIdentifier := ident.Name
					__gong__map_Paragraph[identifier].EnclosingBody = __gong__map_Body[targetIdentifier]
				case "EnclosingTableColumn":
					targetIdentifier := ident.Name
					__gong__map_Paragraph[identifier].EnclosingTableColumn = __gong__map_TableColumn[targetIdentifier]
				}
			case "ParagraphProperties":
				switch fieldName {
				// insertion point for field dependant code
				case "ParagraphStyle":
					targetIdentifier := ident.Name
					__gong__map_ParagraphProperties[identifier].ParagraphStyle = __gong__map_ParagraphStyle[targetIdentifier]
				case "Node":
					targetIdentifier := ident.Name
					__gong__map_ParagraphProperties[identifier].Node = __gong__map_Node[targetIdentifier]
				}
			case "ParagraphStyle":
				switch fieldName {
				// insertion point for field dependant code
				case "Node":
					targetIdentifier := ident.Name
					__gong__map_ParagraphStyle[identifier].Node = __gong__map_Node[targetIdentifier]
				}
			case "Rune":
				switch fieldName {
				// insertion point for field dependant code
				case "Node":
					targetIdentifier := ident.Name
					__gong__map_Rune[identifier].Node = __gong__map_Node[targetIdentifier]
				case "Text":
					targetIdentifier := ident.Name
					__gong__map_Rune[identifier].Text = __gong__map_Text[targetIdentifier]
				case "RuneProperties":
					targetIdentifier := ident.Name
					__gong__map_Rune[identifier].RuneProperties = __gong__map_RuneProperties[targetIdentifier]
				case "EnclosingParagraph":
					targetIdentifier := ident.Name
					__gong__map_Rune[identifier].EnclosingParagraph = __gong__map_Paragraph[targetIdentifier]
				}
			case "RuneProperties":
				switch fieldName {
				// insertion point for field dependant code
				case "Node":
					targetIdentifier := ident.Name
					__gong__map_RuneProperties[identifier].Node = __gong__map_Node[targetIdentifier]
				case "IsBold":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RuneProperties[identifier].IsBold = fielValue
				case "IsStrike":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RuneProperties[identifier].IsStrike = fielValue
				case "IsItalic":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RuneProperties[identifier].IsItalic = fielValue
				}
			case "Table":
				switch fieldName {
				// insertion point for field dependant code
				case "Node":
					targetIdentifier := ident.Name
					__gong__map_Table[identifier].Node = __gong__map_Node[targetIdentifier]
				case "TableProperties":
					targetIdentifier := ident.Name
					__gong__map_Table[identifier].TableProperties = __gong__map_TableProperties[targetIdentifier]
				}
			case "TableColumn":
				switch fieldName {
				// insertion point for field dependant code
				case "Node":
					targetIdentifier := ident.Name
					__gong__map_TableColumn[identifier].Node = __gong__map_Node[targetIdentifier]
				}
			case "TableProperties":
				switch fieldName {
				// insertion point for field dependant code
				case "Node":
					targetIdentifier := ident.Name
					__gong__map_TableProperties[identifier].Node = __gong__map_Node[targetIdentifier]
				case "TableStyle":
					targetIdentifier := ident.Name
					__gong__map_TableProperties[identifier].TableStyle = __gong__map_TableStyle[targetIdentifier]
				}
			case "TableRow":
				switch fieldName {
				// insertion point for field dependant code
				case "Node":
					targetIdentifier := ident.Name
					__gong__map_TableRow[identifier].Node = __gong__map_Node[targetIdentifier]
				}
			case "TableStyle":
				switch fieldName {
				// insertion point for field dependant code
				case "Node":
					targetIdentifier := ident.Name
					__gong__map_TableStyle[identifier].Node = __gong__map_Node[targetIdentifier]
				}
			case "Text":
				switch fieldName {
				// insertion point for field dependant code
				case "Node":
					targetIdentifier := ident.Name
					__gong__map_Text[identifier].Node = __gong__map_Node[targetIdentifier]
				case "PreserveWhiteSpace":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Text[identifier].PreserveWhiteSpace = fielValue
				case "EnclosingRune":
					targetIdentifier := ident.Name
					__gong__map_Text[identifier].EnclosingRune = __gong__map_Rune[targetIdentifier]
				}
			}
		case *ast.SelectorExpr:
			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for enums assignments
				case "Body":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Document":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Docx":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "File":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Node":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Paragraph":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ParagraphProperties":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ParagraphStyle":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Rune":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "RuneProperties":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Table":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "TableColumn":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "TableProperties":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "TableRow":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "TableStyle":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Text":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}

// ReplaceOldDeclarationsInFile replaces specific text in a file at the given path.
func ReplaceOldDeclarationsInFile(pathToFile string) error {
	// Open the file for reading.
	file, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// replacing function with Injection
	pattern := regexp.MustCompile(`\b\w*Injection\b`)
	pattern2 := regexp.MustCompile(`\bmap_DocLink_Identifier_\w*\b`)

	// Temporary slice to hold lines from the file.
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Replace the target text with the desired text.
		line := strings.Replace(scanner.Text(), "var ___dummy__Time_stage time.Time", "var _ time.Time", -1)
		line = pattern.ReplaceAllString(line, "_")
		line = pattern2.ReplaceAllString(line, "_")

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Re-open the file for writing.
	file, err = os.Create(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the modified lines back to the file.
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
