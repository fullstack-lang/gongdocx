// insertion point sub template for components imports 
  import { DocxsTableComponent } from './docxs-table/docxs-table.component'
  import { DocxSortingComponent } from './docx-sorting/docx-sorting.component'
  import { FilesTableComponent } from './files-table/files-table.component'
  import { FileSortingComponent } from './file-sorting/file-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfDocxsComponents: Map<string, any> = new Map([["DocxsTableComponent", DocxsTableComponent],])
  export const MapOfDocxSortingComponents: Map<string, any> = new Map([["DocxSortingComponent", DocxSortingComponent],])
  export const MapOfFilesComponents: Map<string, any> = new Map([["FilesTableComponent", FilesTableComponent],])
  export const MapOfFileSortingComponents: Map<string, any> = new Map([["FileSortingComponent", FileSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Docx", MapOfDocxsComponents],
      ["File", MapOfFilesComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Docx", MapOfDocxSortingComponents],
      ["File", MapOfFileSortingComponents],
    ]
  )
