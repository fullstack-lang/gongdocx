// insertion point sub template for components imports 
  import { DocumentsTableComponent } from './documents-table/documents-table.component'
  import { DocumentSortingComponent } from './document-sorting/document-sorting.component'
  import { DocxsTableComponent } from './docxs-table/docxs-table.component'
  import { DocxSortingComponent } from './docx-sorting/docx-sorting.component'
  import { FilesTableComponent } from './files-table/files-table.component'
  import { FileSortingComponent } from './file-sorting/file-sorting.component'
  import { NodesTableComponent } from './nodes-table/nodes-table.component'
  import { NodeSortingComponent } from './node-sorting/node-sorting.component'
  import { ParagraphsTableComponent } from './paragraphs-table/paragraphs-table.component'
  import { ParagraphSortingComponent } from './paragraph-sorting/paragraph-sorting.component'
  import { ParagraphPropertiessTableComponent } from './paragraphpropertiess-table/paragraphpropertiess-table.component'
  import { ParagraphPropertiesSortingComponent } from './paragraphproperties-sorting/paragraphproperties-sorting.component'
  import { ParagraphStylesTableComponent } from './paragraphstyles-table/paragraphstyles-table.component'
  import { ParagraphStyleSortingComponent } from './paragraphstyle-sorting/paragraphstyle-sorting.component'
  import { RunesTableComponent } from './runes-table/runes-table.component'
  import { RuneSortingComponent } from './rune-sorting/rune-sorting.component'
  import { RunePropertiessTableComponent } from './runepropertiess-table/runepropertiess-table.component'
  import { RunePropertiesSortingComponent } from './runeproperties-sorting/runeproperties-sorting.component'
  import { TablesTableComponent } from './tables-table/tables-table.component'
  import { TableSortingComponent } from './table-sorting/table-sorting.component'
  import { TablePropertiessTableComponent } from './tablepropertiess-table/tablepropertiess-table.component'
  import { TablePropertiesSortingComponent } from './tableproperties-sorting/tableproperties-sorting.component'
  import { TableStylesTableComponent } from './tablestyles-table/tablestyles-table.component'
  import { TableStyleSortingComponent } from './tablestyle-sorting/tablestyle-sorting.component'
  import { TextsTableComponent } from './texts-table/texts-table.component'
  import { TextSortingComponent } from './text-sorting/text-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfDocumentsComponents: Map<string, any> = new Map([["DocumentsTableComponent", DocumentsTableComponent],])
  export const MapOfDocumentSortingComponents: Map<string, any> = new Map([["DocumentSortingComponent", DocumentSortingComponent],])
  export const MapOfDocxsComponents: Map<string, any> = new Map([["DocxsTableComponent", DocxsTableComponent],])
  export const MapOfDocxSortingComponents: Map<string, any> = new Map([["DocxSortingComponent", DocxSortingComponent],])
  export const MapOfFilesComponents: Map<string, any> = new Map([["FilesTableComponent", FilesTableComponent],])
  export const MapOfFileSortingComponents: Map<string, any> = new Map([["FileSortingComponent", FileSortingComponent],])
  export const MapOfNodesComponents: Map<string, any> = new Map([["NodesTableComponent", NodesTableComponent],])
  export const MapOfNodeSortingComponents: Map<string, any> = new Map([["NodeSortingComponent", NodeSortingComponent],])
  export const MapOfParagraphsComponents: Map<string, any> = new Map([["ParagraphsTableComponent", ParagraphsTableComponent],])
  export const MapOfParagraphSortingComponents: Map<string, any> = new Map([["ParagraphSortingComponent", ParagraphSortingComponent],])
  export const MapOfParagraphPropertiessComponents: Map<string, any> = new Map([["ParagraphPropertiessTableComponent", ParagraphPropertiessTableComponent],])
  export const MapOfParagraphPropertiesSortingComponents: Map<string, any> = new Map([["ParagraphPropertiesSortingComponent", ParagraphPropertiesSortingComponent],])
  export const MapOfParagraphStylesComponents: Map<string, any> = new Map([["ParagraphStylesTableComponent", ParagraphStylesTableComponent],])
  export const MapOfParagraphStyleSortingComponents: Map<string, any> = new Map([["ParagraphStyleSortingComponent", ParagraphStyleSortingComponent],])
  export const MapOfRunesComponents: Map<string, any> = new Map([["RunesTableComponent", RunesTableComponent],])
  export const MapOfRuneSortingComponents: Map<string, any> = new Map([["RuneSortingComponent", RuneSortingComponent],])
  export const MapOfRunePropertiessComponents: Map<string, any> = new Map([["RunePropertiessTableComponent", RunePropertiessTableComponent],])
  export const MapOfRunePropertiesSortingComponents: Map<string, any> = new Map([["RunePropertiesSortingComponent", RunePropertiesSortingComponent],])
  export const MapOfTablesComponents: Map<string, any> = new Map([["TablesTableComponent", TablesTableComponent],])
  export const MapOfTableSortingComponents: Map<string, any> = new Map([["TableSortingComponent", TableSortingComponent],])
  export const MapOfTablePropertiessComponents: Map<string, any> = new Map([["TablePropertiessTableComponent", TablePropertiessTableComponent],])
  export const MapOfTablePropertiesSortingComponents: Map<string, any> = new Map([["TablePropertiesSortingComponent", TablePropertiesSortingComponent],])
  export const MapOfTableStylesComponents: Map<string, any> = new Map([["TableStylesTableComponent", TableStylesTableComponent],])
  export const MapOfTableStyleSortingComponents: Map<string, any> = new Map([["TableStyleSortingComponent", TableStyleSortingComponent],])
  export const MapOfTextsComponents: Map<string, any> = new Map([["TextsTableComponent", TextsTableComponent],])
  export const MapOfTextSortingComponents: Map<string, any> = new Map([["TextSortingComponent", TextSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Document", MapOfDocumentsComponents],
      ["Docx", MapOfDocxsComponents],
      ["File", MapOfFilesComponents],
      ["Node", MapOfNodesComponents],
      ["Paragraph", MapOfParagraphsComponents],
      ["ParagraphProperties", MapOfParagraphPropertiessComponents],
      ["ParagraphStyle", MapOfParagraphStylesComponents],
      ["Rune", MapOfRunesComponents],
      ["RuneProperties", MapOfRunePropertiessComponents],
      ["Table", MapOfTablesComponents],
      ["TableProperties", MapOfTablePropertiessComponents],
      ["TableStyle", MapOfTableStylesComponents],
      ["Text", MapOfTextsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Document", MapOfDocumentSortingComponents],
      ["Docx", MapOfDocxSortingComponents],
      ["File", MapOfFileSortingComponents],
      ["Node", MapOfNodeSortingComponents],
      ["Paragraph", MapOfParagraphSortingComponents],
      ["ParagraphProperties", MapOfParagraphPropertiesSortingComponents],
      ["ParagraphStyle", MapOfParagraphStyleSortingComponents],
      ["Rune", MapOfRuneSortingComponents],
      ["RuneProperties", MapOfRunePropertiesSortingComponents],
      ["Table", MapOfTableSortingComponents],
      ["TableProperties", MapOfTablePropertiesSortingComponents],
      ["TableStyle", MapOfTableStyleSortingComponents],
      ["Text", MapOfTextSortingComponents],
    ]
  )
