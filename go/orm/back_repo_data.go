// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	BodyAPIs []*BodyAPI

	DocumentAPIs []*DocumentAPI

	DocxAPIs []*DocxAPI

	FileAPIs []*FileAPI

	NodeAPIs []*NodeAPI

	ParagraphAPIs []*ParagraphAPI

	ParagraphPropertiesAPIs []*ParagraphPropertiesAPI

	ParagraphStyleAPIs []*ParagraphStyleAPI

	RuneAPIs []*RuneAPI

	RunePropertiesAPIs []*RunePropertiesAPI

	TableAPIs []*TableAPI

	TableColumnAPIs []*TableColumnAPI

	TablePropertiesAPIs []*TablePropertiesAPI

	TableRowAPIs []*TableRowAPI

	TableStyleAPIs []*TableStyleAPI

	TextAPIs []*TextAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, bodyDB := range backRepo.BackRepoBody.Map_BodyDBID_BodyDB {

		var bodyAPI BodyAPI
		bodyAPI.ID = bodyDB.ID
		bodyAPI.BodyPointersEncoding = bodyDB.BodyPointersEncoding
		bodyDB.CopyBasicFieldsToBody_WOP(&bodyAPI.Body_WOP)

		backRepoData.BodyAPIs = append(backRepoData.BodyAPIs, &bodyAPI)
	}

	for _, documentDB := range backRepo.BackRepoDocument.Map_DocumentDBID_DocumentDB {

		var documentAPI DocumentAPI
		documentAPI.ID = documentDB.ID
		documentAPI.DocumentPointersEncoding = documentDB.DocumentPointersEncoding
		documentDB.CopyBasicFieldsToDocument_WOP(&documentAPI.Document_WOP)

		backRepoData.DocumentAPIs = append(backRepoData.DocumentAPIs, &documentAPI)
	}

	for _, docxDB := range backRepo.BackRepoDocx.Map_DocxDBID_DocxDB {

		var docxAPI DocxAPI
		docxAPI.ID = docxDB.ID
		docxAPI.DocxPointersEncoding = docxDB.DocxPointersEncoding
		docxDB.CopyBasicFieldsToDocx_WOP(&docxAPI.Docx_WOP)

		backRepoData.DocxAPIs = append(backRepoData.DocxAPIs, &docxAPI)
	}

	for _, fileDB := range backRepo.BackRepoFile.Map_FileDBID_FileDB {

		var fileAPI FileAPI
		fileAPI.ID = fileDB.ID
		fileAPI.FilePointersEncoding = fileDB.FilePointersEncoding
		fileDB.CopyBasicFieldsToFile_WOP(&fileAPI.File_WOP)

		backRepoData.FileAPIs = append(backRepoData.FileAPIs, &fileAPI)
	}

	for _, nodeDB := range backRepo.BackRepoNode.Map_NodeDBID_NodeDB {

		var nodeAPI NodeAPI
		nodeAPI.ID = nodeDB.ID
		nodeAPI.NodePointersEncoding = nodeDB.NodePointersEncoding
		nodeDB.CopyBasicFieldsToNode_WOP(&nodeAPI.Node_WOP)

		backRepoData.NodeAPIs = append(backRepoData.NodeAPIs, &nodeAPI)
	}

	for _, paragraphDB := range backRepo.BackRepoParagraph.Map_ParagraphDBID_ParagraphDB {

		var paragraphAPI ParagraphAPI
		paragraphAPI.ID = paragraphDB.ID
		paragraphAPI.ParagraphPointersEncoding = paragraphDB.ParagraphPointersEncoding
		paragraphDB.CopyBasicFieldsToParagraph_WOP(&paragraphAPI.Paragraph_WOP)

		backRepoData.ParagraphAPIs = append(backRepoData.ParagraphAPIs, &paragraphAPI)
	}

	for _, paragraphpropertiesDB := range backRepo.BackRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB {

		var paragraphpropertiesAPI ParagraphPropertiesAPI
		paragraphpropertiesAPI.ID = paragraphpropertiesDB.ID
		paragraphpropertiesAPI.ParagraphPropertiesPointersEncoding = paragraphpropertiesDB.ParagraphPropertiesPointersEncoding
		paragraphpropertiesDB.CopyBasicFieldsToParagraphProperties_WOP(&paragraphpropertiesAPI.ParagraphProperties_WOP)

		backRepoData.ParagraphPropertiesAPIs = append(backRepoData.ParagraphPropertiesAPIs, &paragraphpropertiesAPI)
	}

	for _, paragraphstyleDB := range backRepo.BackRepoParagraphStyle.Map_ParagraphStyleDBID_ParagraphStyleDB {

		var paragraphstyleAPI ParagraphStyleAPI
		paragraphstyleAPI.ID = paragraphstyleDB.ID
		paragraphstyleAPI.ParagraphStylePointersEncoding = paragraphstyleDB.ParagraphStylePointersEncoding
		paragraphstyleDB.CopyBasicFieldsToParagraphStyle_WOP(&paragraphstyleAPI.ParagraphStyle_WOP)

		backRepoData.ParagraphStyleAPIs = append(backRepoData.ParagraphStyleAPIs, &paragraphstyleAPI)
	}

	for _, runeDB := range backRepo.BackRepoRune.Map_RuneDBID_RuneDB {

		var runeAPI RuneAPI
		runeAPI.ID = runeDB.ID
		runeAPI.RunePointersEncoding = runeDB.RunePointersEncoding
		runeDB.CopyBasicFieldsToRune_WOP(&runeAPI.Rune_WOP)

		backRepoData.RuneAPIs = append(backRepoData.RuneAPIs, &runeAPI)
	}

	for _, runepropertiesDB := range backRepo.BackRepoRuneProperties.Map_RunePropertiesDBID_RunePropertiesDB {

		var runepropertiesAPI RunePropertiesAPI
		runepropertiesAPI.ID = runepropertiesDB.ID
		runepropertiesAPI.RunePropertiesPointersEncoding = runepropertiesDB.RunePropertiesPointersEncoding
		runepropertiesDB.CopyBasicFieldsToRuneProperties_WOP(&runepropertiesAPI.RuneProperties_WOP)

		backRepoData.RunePropertiesAPIs = append(backRepoData.RunePropertiesAPIs, &runepropertiesAPI)
	}

	for _, tableDB := range backRepo.BackRepoTable.Map_TableDBID_TableDB {

		var tableAPI TableAPI
		tableAPI.ID = tableDB.ID
		tableAPI.TablePointersEncoding = tableDB.TablePointersEncoding
		tableDB.CopyBasicFieldsToTable_WOP(&tableAPI.Table_WOP)

		backRepoData.TableAPIs = append(backRepoData.TableAPIs, &tableAPI)
	}

	for _, tablecolumnDB := range backRepo.BackRepoTableColumn.Map_TableColumnDBID_TableColumnDB {

		var tablecolumnAPI TableColumnAPI
		tablecolumnAPI.ID = tablecolumnDB.ID
		tablecolumnAPI.TableColumnPointersEncoding = tablecolumnDB.TableColumnPointersEncoding
		tablecolumnDB.CopyBasicFieldsToTableColumn_WOP(&tablecolumnAPI.TableColumn_WOP)

		backRepoData.TableColumnAPIs = append(backRepoData.TableColumnAPIs, &tablecolumnAPI)
	}

	for _, tablepropertiesDB := range backRepo.BackRepoTableProperties.Map_TablePropertiesDBID_TablePropertiesDB {

		var tablepropertiesAPI TablePropertiesAPI
		tablepropertiesAPI.ID = tablepropertiesDB.ID
		tablepropertiesAPI.TablePropertiesPointersEncoding = tablepropertiesDB.TablePropertiesPointersEncoding
		tablepropertiesDB.CopyBasicFieldsToTableProperties_WOP(&tablepropertiesAPI.TableProperties_WOP)

		backRepoData.TablePropertiesAPIs = append(backRepoData.TablePropertiesAPIs, &tablepropertiesAPI)
	}

	for _, tablerowDB := range backRepo.BackRepoTableRow.Map_TableRowDBID_TableRowDB {

		var tablerowAPI TableRowAPI
		tablerowAPI.ID = tablerowDB.ID
		tablerowAPI.TableRowPointersEncoding = tablerowDB.TableRowPointersEncoding
		tablerowDB.CopyBasicFieldsToTableRow_WOP(&tablerowAPI.TableRow_WOP)

		backRepoData.TableRowAPIs = append(backRepoData.TableRowAPIs, &tablerowAPI)
	}

	for _, tablestyleDB := range backRepo.BackRepoTableStyle.Map_TableStyleDBID_TableStyleDB {

		var tablestyleAPI TableStyleAPI
		tablestyleAPI.ID = tablestyleDB.ID
		tablestyleAPI.TableStylePointersEncoding = tablestyleDB.TableStylePointersEncoding
		tablestyleDB.CopyBasicFieldsToTableStyle_WOP(&tablestyleAPI.TableStyle_WOP)

		backRepoData.TableStyleAPIs = append(backRepoData.TableStyleAPIs, &tablestyleAPI)
	}

	for _, textDB := range backRepo.BackRepoText.Map_TextDBID_TextDB {

		var textAPI TextAPI
		textAPI.ID = textDB.ID
		textAPI.TextPointersEncoding = textDB.TextPointersEncoding
		textDB.CopyBasicFieldsToText_WOP(&textAPI.Text_WOP)

		backRepoData.TextAPIs = append(backRepoData.TextAPIs, &textAPI)
	}

}
