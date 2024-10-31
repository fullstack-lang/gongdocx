// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongdocx/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	bodyDBs map[uint]*BodyDB

	nextIDBodyDB uint

	documentDBs map[uint]*DocumentDB

	nextIDDocumentDB uint

	docxDBs map[uint]*DocxDB

	nextIDDocxDB uint

	fileDBs map[uint]*FileDB

	nextIDFileDB uint

	nodeDBs map[uint]*NodeDB

	nextIDNodeDB uint

	paragraphDBs map[uint]*ParagraphDB

	nextIDParagraphDB uint

	paragraphpropertiesDBs map[uint]*ParagraphPropertiesDB

	nextIDParagraphPropertiesDB uint

	paragraphstyleDBs map[uint]*ParagraphStyleDB

	nextIDParagraphStyleDB uint

	runeDBs map[uint]*RuneDB

	nextIDRuneDB uint

	runepropertiesDBs map[uint]*RunePropertiesDB

	nextIDRunePropertiesDB uint

	tableDBs map[uint]*TableDB

	nextIDTableDB uint

	tablecolumnDBs map[uint]*TableColumnDB

	nextIDTableColumnDB uint

	tablepropertiesDBs map[uint]*TablePropertiesDB

	nextIDTablePropertiesDB uint

	tablerowDBs map[uint]*TableRowDB

	nextIDTableRowDB uint

	tablestyleDBs map[uint]*TableStyleDB

	nextIDTableStyleDB uint

	textDBs map[uint]*TextDB

	nextIDTextDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		bodyDBs: make(map[uint]*BodyDB),

		documentDBs: make(map[uint]*DocumentDB),

		docxDBs: make(map[uint]*DocxDB),

		fileDBs: make(map[uint]*FileDB),

		nodeDBs: make(map[uint]*NodeDB),

		paragraphDBs: make(map[uint]*ParagraphDB),

		paragraphpropertiesDBs: make(map[uint]*ParagraphPropertiesDB),

		paragraphstyleDBs: make(map[uint]*ParagraphStyleDB),

		runeDBs: make(map[uint]*RuneDB),

		runepropertiesDBs: make(map[uint]*RunePropertiesDB),

		tableDBs: make(map[uint]*TableDB),

		tablecolumnDBs: make(map[uint]*TableColumnDB),

		tablepropertiesDBs: make(map[uint]*TablePropertiesDB),

		tablerowDBs: make(map[uint]*TableRowDB),

		tablestyleDBs: make(map[uint]*TableStyleDB),

		textDBs: make(map[uint]*TextDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *BodyDB:
		db.nextIDBodyDB++
		v.ID = db.nextIDBodyDB
		db.bodyDBs[v.ID] = v
	case *DocumentDB:
		db.nextIDDocumentDB++
		v.ID = db.nextIDDocumentDB
		db.documentDBs[v.ID] = v
	case *DocxDB:
		db.nextIDDocxDB++
		v.ID = db.nextIDDocxDB
		db.docxDBs[v.ID] = v
	case *FileDB:
		db.nextIDFileDB++
		v.ID = db.nextIDFileDB
		db.fileDBs[v.ID] = v
	case *NodeDB:
		db.nextIDNodeDB++
		v.ID = db.nextIDNodeDB
		db.nodeDBs[v.ID] = v
	case *ParagraphDB:
		db.nextIDParagraphDB++
		v.ID = db.nextIDParagraphDB
		db.paragraphDBs[v.ID] = v
	case *ParagraphPropertiesDB:
		db.nextIDParagraphPropertiesDB++
		v.ID = db.nextIDParagraphPropertiesDB
		db.paragraphpropertiesDBs[v.ID] = v
	case *ParagraphStyleDB:
		db.nextIDParagraphStyleDB++
		v.ID = db.nextIDParagraphStyleDB
		db.paragraphstyleDBs[v.ID] = v
	case *RuneDB:
		db.nextIDRuneDB++
		v.ID = db.nextIDRuneDB
		db.runeDBs[v.ID] = v
	case *RunePropertiesDB:
		db.nextIDRunePropertiesDB++
		v.ID = db.nextIDRunePropertiesDB
		db.runepropertiesDBs[v.ID] = v
	case *TableDB:
		db.nextIDTableDB++
		v.ID = db.nextIDTableDB
		db.tableDBs[v.ID] = v
	case *TableColumnDB:
		db.nextIDTableColumnDB++
		v.ID = db.nextIDTableColumnDB
		db.tablecolumnDBs[v.ID] = v
	case *TablePropertiesDB:
		db.nextIDTablePropertiesDB++
		v.ID = db.nextIDTablePropertiesDB
		db.tablepropertiesDBs[v.ID] = v
	case *TableRowDB:
		db.nextIDTableRowDB++
		v.ID = db.nextIDTableRowDB
		db.tablerowDBs[v.ID] = v
	case *TableStyleDB:
		db.nextIDTableStyleDB++
		v.ID = db.nextIDTableStyleDB
		db.tablestyleDBs[v.ID] = v
	case *TextDB:
		db.nextIDTextDB++
		v.ID = db.nextIDTextDB
		db.textDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BodyDB:
		delete(db.bodyDBs, v.ID)
	case *DocumentDB:
		delete(db.documentDBs, v.ID)
	case *DocxDB:
		delete(db.docxDBs, v.ID)
	case *FileDB:
		delete(db.fileDBs, v.ID)
	case *NodeDB:
		delete(db.nodeDBs, v.ID)
	case *ParagraphDB:
		delete(db.paragraphDBs, v.ID)
	case *ParagraphPropertiesDB:
		delete(db.paragraphpropertiesDBs, v.ID)
	case *ParagraphStyleDB:
		delete(db.paragraphstyleDBs, v.ID)
	case *RuneDB:
		delete(db.runeDBs, v.ID)
	case *RunePropertiesDB:
		delete(db.runepropertiesDBs, v.ID)
	case *TableDB:
		delete(db.tableDBs, v.ID)
	case *TableColumnDB:
		delete(db.tablecolumnDBs, v.ID)
	case *TablePropertiesDB:
		delete(db.tablepropertiesDBs, v.ID)
	case *TableRowDB:
		delete(db.tablerowDBs, v.ID)
	case *TableStyleDB:
		delete(db.tablestyleDBs, v.ID)
	case *TextDB:
		delete(db.textDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BodyDB:
		db.bodyDBs[v.ID] = v
		return db, nil
	case *DocumentDB:
		db.documentDBs[v.ID] = v
		return db, nil
	case *DocxDB:
		db.docxDBs[v.ID] = v
		return db, nil
	case *FileDB:
		db.fileDBs[v.ID] = v
		return db, nil
	case *NodeDB:
		db.nodeDBs[v.ID] = v
		return db, nil
	case *ParagraphDB:
		db.paragraphDBs[v.ID] = v
		return db, nil
	case *ParagraphPropertiesDB:
		db.paragraphpropertiesDBs[v.ID] = v
		return db, nil
	case *ParagraphStyleDB:
		db.paragraphstyleDBs[v.ID] = v
		return db, nil
	case *RuneDB:
		db.runeDBs[v.ID] = v
		return db, nil
	case *RunePropertiesDB:
		db.runepropertiesDBs[v.ID] = v
		return db, nil
	case *TableDB:
		db.tableDBs[v.ID] = v
		return db, nil
	case *TableColumnDB:
		db.tablecolumnDBs[v.ID] = v
		return db, nil
	case *TablePropertiesDB:
		db.tablepropertiesDBs[v.ID] = v
		return db, nil
	case *TableRowDB:
		db.tablerowDBs[v.ID] = v
		return db, nil
	case *TableStyleDB:
		db.tablestyleDBs[v.ID] = v
		return db, nil
	case *TextDB:
		db.textDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BodyDB:
		if existing, ok := db.bodyDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Body github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *DocumentDB:
		if existing, ok := db.documentDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Document github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *DocxDB:
		if existing, ok := db.docxDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Docx github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *FileDB:
		if existing, ok := db.fileDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db File github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *NodeDB:
		if existing, ok := db.nodeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Node github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *ParagraphDB:
		if existing, ok := db.paragraphDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Paragraph github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *ParagraphPropertiesDB:
		if existing, ok := db.paragraphpropertiesDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ParagraphProperties github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *ParagraphStyleDB:
		if existing, ok := db.paragraphstyleDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ParagraphStyle github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *RuneDB:
		if existing, ok := db.runeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Rune github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *RunePropertiesDB:
		if existing, ok := db.runepropertiesDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db RuneProperties github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *TableDB:
		if existing, ok := db.tableDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Table github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *TableColumnDB:
		if existing, ok := db.tablecolumnDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db TableColumn github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *TablePropertiesDB:
		if existing, ok := db.tablepropertiesDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db TableProperties github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *TableRowDB:
		if existing, ok := db.tablerowDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db TableRow github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *TableStyleDB:
		if existing, ok := db.tablestyleDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db TableStyle github.com/fullstack-lang/gongdocx/go, record not found")
		}
	case *TextDB:
		if existing, ok := db.textDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Text github.com/fullstack-lang/gongdocx/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]BodyDB:
		*ptr = make([]BodyDB, 0, len(db.bodyDBs))
		for _, v := range db.bodyDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DocumentDB:
		*ptr = make([]DocumentDB, 0, len(db.documentDBs))
		for _, v := range db.documentDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DocxDB:
		*ptr = make([]DocxDB, 0, len(db.docxDBs))
		for _, v := range db.docxDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FileDB:
		*ptr = make([]FileDB, 0, len(db.fileDBs))
		for _, v := range db.fileDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]NodeDB:
		*ptr = make([]NodeDB, 0, len(db.nodeDBs))
		for _, v := range db.nodeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ParagraphDB:
		*ptr = make([]ParagraphDB, 0, len(db.paragraphDBs))
		for _, v := range db.paragraphDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ParagraphPropertiesDB:
		*ptr = make([]ParagraphPropertiesDB, 0, len(db.paragraphpropertiesDBs))
		for _, v := range db.paragraphpropertiesDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ParagraphStyleDB:
		*ptr = make([]ParagraphStyleDB, 0, len(db.paragraphstyleDBs))
		for _, v := range db.paragraphstyleDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RuneDB:
		*ptr = make([]RuneDB, 0, len(db.runeDBs))
		for _, v := range db.runeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RunePropertiesDB:
		*ptr = make([]RunePropertiesDB, 0, len(db.runepropertiesDBs))
		for _, v := range db.runepropertiesDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TableDB:
		*ptr = make([]TableDB, 0, len(db.tableDBs))
		for _, v := range db.tableDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TableColumnDB:
		*ptr = make([]TableColumnDB, 0, len(db.tablecolumnDBs))
		for _, v := range db.tablecolumnDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TablePropertiesDB:
		*ptr = make([]TablePropertiesDB, 0, len(db.tablepropertiesDBs))
		for _, v := range db.tablepropertiesDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TableRowDB:
		*ptr = make([]TableRowDB, 0, len(db.tablerowDBs))
		for _, v := range db.tablerowDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TableStyleDB:
		*ptr = make([]TableStyleDB, 0, len(db.tablestyleDBs))
		for _, v := range db.tablestyleDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TextDB:
		*ptr = make([]TextDB, 0, len(db.textDBs))
		for _, v := range db.textDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, Do not process when conds is not a single parameter")
	}

	str, ok := conds[0].(string)

	if !ok {
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, conds[0] is not a string")
	}

	i, err := strconv.ParseUint(str, 10, 32) // Base 10, 32-bit unsigned int
	if err != nil {
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, conds[0] is not a string number")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *BodyDB:
		tmp, ok := db.bodyDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Body Unkown entry %d", i))
		}

		bodyDB, _ := instanceDB.(*BodyDB)
		*bodyDB = *tmp
		
	case *DocumentDB:
		tmp, ok := db.documentDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Document Unkown entry %d", i))
		}

		documentDB, _ := instanceDB.(*DocumentDB)
		*documentDB = *tmp
		
	case *DocxDB:
		tmp, ok := db.docxDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Docx Unkown entry %d", i))
		}

		docxDB, _ := instanceDB.(*DocxDB)
		*docxDB = *tmp
		
	case *FileDB:
		tmp, ok := db.fileDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First File Unkown entry %d", i))
		}

		fileDB, _ := instanceDB.(*FileDB)
		*fileDB = *tmp
		
	case *NodeDB:
		tmp, ok := db.nodeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Node Unkown entry %d", i))
		}

		nodeDB, _ := instanceDB.(*NodeDB)
		*nodeDB = *tmp
		
	case *ParagraphDB:
		tmp, ok := db.paragraphDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Paragraph Unkown entry %d", i))
		}

		paragraphDB, _ := instanceDB.(*ParagraphDB)
		*paragraphDB = *tmp
		
	case *ParagraphPropertiesDB:
		tmp, ok := db.paragraphpropertiesDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ParagraphProperties Unkown entry %d", i))
		}

		paragraphpropertiesDB, _ := instanceDB.(*ParagraphPropertiesDB)
		*paragraphpropertiesDB = *tmp
		
	case *ParagraphStyleDB:
		tmp, ok := db.paragraphstyleDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ParagraphStyle Unkown entry %d", i))
		}

		paragraphstyleDB, _ := instanceDB.(*ParagraphStyleDB)
		*paragraphstyleDB = *tmp
		
	case *RuneDB:
		tmp, ok := db.runeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Rune Unkown entry %d", i))
		}

		runeDB, _ := instanceDB.(*RuneDB)
		*runeDB = *tmp
		
	case *RunePropertiesDB:
		tmp, ok := db.runepropertiesDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First RuneProperties Unkown entry %d", i))
		}

		runepropertiesDB, _ := instanceDB.(*RunePropertiesDB)
		*runepropertiesDB = *tmp
		
	case *TableDB:
		tmp, ok := db.tableDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Table Unkown entry %d", i))
		}

		tableDB, _ := instanceDB.(*TableDB)
		*tableDB = *tmp
		
	case *TableColumnDB:
		tmp, ok := db.tablecolumnDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First TableColumn Unkown entry %d", i))
		}

		tablecolumnDB, _ := instanceDB.(*TableColumnDB)
		*tablecolumnDB = *tmp
		
	case *TablePropertiesDB:
		tmp, ok := db.tablepropertiesDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First TableProperties Unkown entry %d", i))
		}

		tablepropertiesDB, _ := instanceDB.(*TablePropertiesDB)
		*tablepropertiesDB = *tmp
		
	case *TableRowDB:
		tmp, ok := db.tablerowDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First TableRow Unkown entry %d", i))
		}

		tablerowDB, _ := instanceDB.(*TableRowDB)
		*tablerowDB = *tmp
		
	case *TableStyleDB:
		tmp, ok := db.tablestyleDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First TableStyle Unkown entry %d", i))
		}

		tablestyleDB, _ := instanceDB.(*TableStyleDB)
		*tablestyleDB = *tmp
		
	case *TextDB:
		tmp, ok := db.textDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Text Unkown entry %d", i))
		}

		textDB, _ := instanceDB.(*TextDB)
		*textDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gongdocx/go, Unkown type")
	}
	
	return db, nil
}

