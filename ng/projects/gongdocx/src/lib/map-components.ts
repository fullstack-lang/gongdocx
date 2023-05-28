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
  import { RunesTableComponent } from './runes-table/runes-table.component'
  import { RuneSortingComponent } from './rune-sorting/rune-sorting.component'
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
  export const MapOfRunesComponents: Map<string, any> = new Map([["RunesTableComponent", RunesTableComponent],])
  export const MapOfRuneSortingComponents: Map<string, any> = new Map([["RuneSortingComponent", RuneSortingComponent],])
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
      ["Rune", MapOfRunesComponents],
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
      ["Rune", MapOfRuneSortingComponents],
      ["Text", MapOfTextSortingComponents],
    ]
  )
