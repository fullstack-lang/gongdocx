// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongdocx/go/db"
	"github.com/fullstack-lang/gongdocx/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_TableStyle_sql sql.NullBool
var dummy_TableStyle_time time.Duration
var dummy_TableStyle_sort sort.Float64Slice

// TableStyleAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model tablestyleAPI
type TableStyleAPI struct {
	gorm.Model

	models.TableStyle_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	TableStylePointersEncoding TableStylePointersEncoding
}

// TableStylePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type TableStylePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Node is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	NodeID sql.NullInt64
}

// TableStyleDB describes a tablestyle in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model tablestyleDB
type TableStyleDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field tablestyleDB.Name
	Name_Data sql.NullString

	// Declation for basic field tablestyleDB.Content
	Content_Data sql.NullString

	// Declation for basic field tablestyleDB.Val
	Val_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	TableStylePointersEncoding
}

// TableStyleDBs arrays tablestyleDBs
// swagger:response tablestyleDBsResponse
type TableStyleDBs []TableStyleDB

// TableStyleDBResponse provides response
// swagger:response tablestyleDBResponse
type TableStyleDBResponse struct {
	TableStyleDB
}

// TableStyleWOP is a TableStyle without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type TableStyleWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Content string `xlsx:"2"`

	Val string `xlsx:"3"`
	// insertion for WOP pointer fields
}

var TableStyle_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Content",
	"Val",
}

type BackRepoTableStyleStruct struct {
	// stores TableStyleDB according to their gorm ID
	Map_TableStyleDBID_TableStyleDB map[uint]*TableStyleDB

	// stores TableStyleDB ID according to TableStyle address
	Map_TableStylePtr_TableStyleDBID map[*models.TableStyle]uint

	// stores TableStyle according to their gorm ID
	Map_TableStyleDBID_TableStylePtr map[uint]*models.TableStyle

	db db.DBInterface

	stage *models.StageStruct
}

func (backRepoTableStyle *BackRepoTableStyleStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoTableStyle.stage
	return
}

func (backRepoTableStyle *BackRepoTableStyleStruct) GetDB() db.DBInterface {
	return backRepoTableStyle.db
}

// GetTableStyleDBFromTableStylePtr is a handy function to access the back repo instance from the stage instance
func (backRepoTableStyle *BackRepoTableStyleStruct) GetTableStyleDBFromTableStylePtr(tablestyle *models.TableStyle) (tablestyleDB *TableStyleDB) {
	id := backRepoTableStyle.Map_TableStylePtr_TableStyleDBID[tablestyle]
	tablestyleDB = backRepoTableStyle.Map_TableStyleDBID_TableStyleDB[id]
	return
}

// BackRepoTableStyle.CommitPhaseOne commits all staged instances of TableStyle to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoTableStyle *BackRepoTableStyleStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for tablestyle := range stage.TableStyles {
		backRepoTableStyle.CommitPhaseOneInstance(tablestyle)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, tablestyle := range backRepoTableStyle.Map_TableStyleDBID_TableStylePtr {
		if _, ok := stage.TableStyles[tablestyle]; !ok {
			backRepoTableStyle.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoTableStyle.CommitDeleteInstance commits deletion of TableStyle to the BackRepo
func (backRepoTableStyle *BackRepoTableStyleStruct) CommitDeleteInstance(id uint) (Error error) {

	tablestyle := backRepoTableStyle.Map_TableStyleDBID_TableStylePtr[id]

	// tablestyle is not staged anymore, remove tablestyleDB
	tablestyleDB := backRepoTableStyle.Map_TableStyleDBID_TableStyleDB[id]
	db, _ := backRepoTableStyle.db.Unscoped()
	_, err := db.Delete(tablestyleDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoTableStyle.Map_TableStylePtr_TableStyleDBID, tablestyle)
	delete(backRepoTableStyle.Map_TableStyleDBID_TableStylePtr, id)
	delete(backRepoTableStyle.Map_TableStyleDBID_TableStyleDB, id)

	return
}

// BackRepoTableStyle.CommitPhaseOneInstance commits tablestyle staged instances of TableStyle to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoTableStyle *BackRepoTableStyleStruct) CommitPhaseOneInstance(tablestyle *models.TableStyle) (Error error) {

	// check if the tablestyle is not commited yet
	if _, ok := backRepoTableStyle.Map_TableStylePtr_TableStyleDBID[tablestyle]; ok {
		return
	}

	// initiate tablestyle
	var tablestyleDB TableStyleDB
	tablestyleDB.CopyBasicFieldsFromTableStyle(tablestyle)

	_, err := backRepoTableStyle.db.Create(&tablestyleDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoTableStyle.Map_TableStylePtr_TableStyleDBID[tablestyle] = tablestyleDB.ID
	backRepoTableStyle.Map_TableStyleDBID_TableStylePtr[tablestyleDB.ID] = tablestyle
	backRepoTableStyle.Map_TableStyleDBID_TableStyleDB[tablestyleDB.ID] = &tablestyleDB

	return
}

// BackRepoTableStyle.CommitPhaseTwo commits all staged instances of TableStyle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTableStyle *BackRepoTableStyleStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, tablestyle := range backRepoTableStyle.Map_TableStyleDBID_TableStylePtr {
		backRepoTableStyle.CommitPhaseTwoInstance(backRepo, idx, tablestyle)
	}

	return
}

// BackRepoTableStyle.CommitPhaseTwoInstance commits {{structname }} of models.TableStyle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTableStyle *BackRepoTableStyleStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, tablestyle *models.TableStyle) (Error error) {

	// fetch matching tablestyleDB
	if tablestyleDB, ok := backRepoTableStyle.Map_TableStyleDBID_TableStyleDB[idx]; ok {

		tablestyleDB.CopyBasicFieldsFromTableStyle(tablestyle)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value tablestyle.Node translates to updating the tablestyle.NodeID
		tablestyleDB.NodeID.Valid = true // allow for a 0 value (nil association)
		if tablestyle.Node != nil {
			if NodeId, ok := backRepo.BackRepoNode.Map_NodePtr_NodeDBID[tablestyle.Node]; ok {
				tablestyleDB.NodeID.Int64 = int64(NodeId)
				tablestyleDB.NodeID.Valid = true
			}
		} else {
			tablestyleDB.NodeID.Int64 = 0
			tablestyleDB.NodeID.Valid = true
		}

		_, err := backRepoTableStyle.db.Save(tablestyleDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown TableStyle intance %s", tablestyle.Name))
		return err
	}

	return
}

// BackRepoTableStyle.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoTableStyle *BackRepoTableStyleStruct) CheckoutPhaseOne() (Error error) {

	tablestyleDBArray := make([]TableStyleDB, 0)
	_, err := backRepoTableStyle.db.Find(&tablestyleDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	tablestyleInstancesToBeRemovedFromTheStage := make(map[*models.TableStyle]any)
	for key, value := range backRepoTableStyle.stage.TableStyles {
		tablestyleInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, tablestyleDB := range tablestyleDBArray {
		backRepoTableStyle.CheckoutPhaseOneInstance(&tablestyleDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		tablestyle, ok := backRepoTableStyle.Map_TableStyleDBID_TableStylePtr[tablestyleDB.ID]
		if ok {
			delete(tablestyleInstancesToBeRemovedFromTheStage, tablestyle)
		}
	}

	// remove from stage and back repo's 3 maps all tablestyles that are not in the checkout
	for tablestyle := range tablestyleInstancesToBeRemovedFromTheStage {
		tablestyle.Unstage(backRepoTableStyle.GetStage())

		// remove instance from the back repo 3 maps
		tablestyleID := backRepoTableStyle.Map_TableStylePtr_TableStyleDBID[tablestyle]
		delete(backRepoTableStyle.Map_TableStylePtr_TableStyleDBID, tablestyle)
		delete(backRepoTableStyle.Map_TableStyleDBID_TableStyleDB, tablestyleID)
		delete(backRepoTableStyle.Map_TableStyleDBID_TableStylePtr, tablestyleID)
	}

	return
}

// CheckoutPhaseOneInstance takes a tablestyleDB that has been found in the DB, updates the backRepo and stages the
// models version of the tablestyleDB
func (backRepoTableStyle *BackRepoTableStyleStruct) CheckoutPhaseOneInstance(tablestyleDB *TableStyleDB) (Error error) {

	tablestyle, ok := backRepoTableStyle.Map_TableStyleDBID_TableStylePtr[tablestyleDB.ID]
	if !ok {
		tablestyle = new(models.TableStyle)

		backRepoTableStyle.Map_TableStyleDBID_TableStylePtr[tablestyleDB.ID] = tablestyle
		backRepoTableStyle.Map_TableStylePtr_TableStyleDBID[tablestyle] = tablestyleDB.ID

		// append model store with the new element
		tablestyle.Name = tablestyleDB.Name_Data.String
		tablestyle.Stage(backRepoTableStyle.GetStage())
	}
	tablestyleDB.CopyBasicFieldsToTableStyle(tablestyle)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	tablestyle.Stage(backRepoTableStyle.GetStage())

	// preserve pointer to tablestyleDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_TableStyleDBID_TableStyleDB)[tablestyleDB hold variable pointers
	tablestyleDB_Data := *tablestyleDB
	preservedPtrToTableStyle := &tablestyleDB_Data
	backRepoTableStyle.Map_TableStyleDBID_TableStyleDB[tablestyleDB.ID] = preservedPtrToTableStyle

	return
}

// BackRepoTableStyle.CheckoutPhaseTwo Checkouts all staged instances of TableStyle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTableStyle *BackRepoTableStyleStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, tablestyleDB := range backRepoTableStyle.Map_TableStyleDBID_TableStyleDB {
		backRepoTableStyle.CheckoutPhaseTwoInstance(backRepo, tablestyleDB)
	}
	return
}

// BackRepoTableStyle.CheckoutPhaseTwoInstance Checkouts staged instances of TableStyle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTableStyle *BackRepoTableStyleStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, tablestyleDB *TableStyleDB) (Error error) {

	tablestyle := backRepoTableStyle.Map_TableStyleDBID_TableStylePtr[tablestyleDB.ID]

	tablestyleDB.DecodePointers(backRepo, tablestyle)

	return
}

func (tablestyleDB *TableStyleDB) DecodePointers(backRepo *BackRepoStruct, tablestyle *models.TableStyle) {

	// insertion point for checkout of pointer encoding
	// Node field	
	{
		id := tablestyleDB.NodeID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: tablestyle.Node, unknown pointer id", id)
				tablestyle.Node = nil
			} else {
				// updates only if field has changed
				if tablestyle.Node == nil || tablestyle.Node != tmp {
					tablestyle.Node = tmp
				}
			}
		} else {
			tablestyle.Node = nil
		}
	}
	
	return
}

// CommitTableStyle allows commit of a single tablestyle (if already staged)
func (backRepo *BackRepoStruct) CommitTableStyle(tablestyle *models.TableStyle) {
	backRepo.BackRepoTableStyle.CommitPhaseOneInstance(tablestyle)
	if id, ok := backRepo.BackRepoTableStyle.Map_TableStylePtr_TableStyleDBID[tablestyle]; ok {
		backRepo.BackRepoTableStyle.CommitPhaseTwoInstance(backRepo, id, tablestyle)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitTableStyle allows checkout of a single tablestyle (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutTableStyle(tablestyle *models.TableStyle) {
	// check if the tablestyle is staged
	if _, ok := backRepo.BackRepoTableStyle.Map_TableStylePtr_TableStyleDBID[tablestyle]; ok {

		if id, ok := backRepo.BackRepoTableStyle.Map_TableStylePtr_TableStyleDBID[tablestyle]; ok {
			var tablestyleDB TableStyleDB
			tablestyleDB.ID = id

			if _, err := backRepo.BackRepoTableStyle.db.First(&tablestyleDB, id); err != nil {
				log.Fatalln("CheckoutTableStyle : Problem with getting object with id:", id)
			}
			backRepo.BackRepoTableStyle.CheckoutPhaseOneInstance(&tablestyleDB)
			backRepo.BackRepoTableStyle.CheckoutPhaseTwoInstance(backRepo, &tablestyleDB)
		}
	}
}

// CopyBasicFieldsFromTableStyle
func (tablestyleDB *TableStyleDB) CopyBasicFieldsFromTableStyle(tablestyle *models.TableStyle) {
	// insertion point for fields commit

	tablestyleDB.Name_Data.String = tablestyle.Name
	tablestyleDB.Name_Data.Valid = true

	tablestyleDB.Content_Data.String = tablestyle.Content
	tablestyleDB.Content_Data.Valid = true

	tablestyleDB.Val_Data.String = tablestyle.Val
	tablestyleDB.Val_Data.Valid = true
}

// CopyBasicFieldsFromTableStyle_WOP
func (tablestyleDB *TableStyleDB) CopyBasicFieldsFromTableStyle_WOP(tablestyle *models.TableStyle_WOP) {
	// insertion point for fields commit

	tablestyleDB.Name_Data.String = tablestyle.Name
	tablestyleDB.Name_Data.Valid = true

	tablestyleDB.Content_Data.String = tablestyle.Content
	tablestyleDB.Content_Data.Valid = true

	tablestyleDB.Val_Data.String = tablestyle.Val
	tablestyleDB.Val_Data.Valid = true
}

// CopyBasicFieldsFromTableStyleWOP
func (tablestyleDB *TableStyleDB) CopyBasicFieldsFromTableStyleWOP(tablestyle *TableStyleWOP) {
	// insertion point for fields commit

	tablestyleDB.Name_Data.String = tablestyle.Name
	tablestyleDB.Name_Data.Valid = true

	tablestyleDB.Content_Data.String = tablestyle.Content
	tablestyleDB.Content_Data.Valid = true

	tablestyleDB.Val_Data.String = tablestyle.Val
	tablestyleDB.Val_Data.Valid = true
}

// CopyBasicFieldsToTableStyle
func (tablestyleDB *TableStyleDB) CopyBasicFieldsToTableStyle(tablestyle *models.TableStyle) {
	// insertion point for checkout of basic fields (back repo to stage)
	tablestyle.Name = tablestyleDB.Name_Data.String
	tablestyle.Content = tablestyleDB.Content_Data.String
	tablestyle.Val = tablestyleDB.Val_Data.String
}

// CopyBasicFieldsToTableStyle_WOP
func (tablestyleDB *TableStyleDB) CopyBasicFieldsToTableStyle_WOP(tablestyle *models.TableStyle_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	tablestyle.Name = tablestyleDB.Name_Data.String
	tablestyle.Content = tablestyleDB.Content_Data.String
	tablestyle.Val = tablestyleDB.Val_Data.String
}

// CopyBasicFieldsToTableStyleWOP
func (tablestyleDB *TableStyleDB) CopyBasicFieldsToTableStyleWOP(tablestyle *TableStyleWOP) {
	tablestyle.ID = int(tablestyleDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	tablestyle.Name = tablestyleDB.Name_Data.String
	tablestyle.Content = tablestyleDB.Content_Data.String
	tablestyle.Val = tablestyleDB.Val_Data.String
}

// Backup generates a json file from a slice of all TableStyleDB instances in the backrepo
func (backRepoTableStyle *BackRepoTableStyleStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "TableStyleDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*TableStyleDB, 0)
	for _, tablestyleDB := range backRepoTableStyle.Map_TableStyleDBID_TableStyleDB {
		forBackup = append(forBackup, tablestyleDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json TableStyle ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json TableStyle file", err.Error())
	}
}

// Backup generates a json file from a slice of all TableStyleDB instances in the backrepo
func (backRepoTableStyle *BackRepoTableStyleStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*TableStyleDB, 0)
	for _, tablestyleDB := range backRepoTableStyle.Map_TableStyleDBID_TableStyleDB {
		forBackup = append(forBackup, tablestyleDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("TableStyle")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&TableStyle_Fields, -1)
	for _, tablestyleDB := range forBackup {

		var tablestyleWOP TableStyleWOP
		tablestyleDB.CopyBasicFieldsToTableStyleWOP(&tablestyleWOP)

		row := sh.AddRow()
		row.WriteStruct(&tablestyleWOP, -1)
	}
}

// RestoreXL from the "TableStyle" sheet all TableStyleDB instances
func (backRepoTableStyle *BackRepoTableStyleStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoTableStyleid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["TableStyle"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoTableStyle.rowVisitorTableStyle)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoTableStyle *BackRepoTableStyleStruct) rowVisitorTableStyle(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var tablestyleWOP TableStyleWOP
		row.ReadStruct(&tablestyleWOP)

		// add the unmarshalled struct to the stage
		tablestyleDB := new(TableStyleDB)
		tablestyleDB.CopyBasicFieldsFromTableStyleWOP(&tablestyleWOP)

		tablestyleDB_ID_atBackupTime := tablestyleDB.ID
		tablestyleDB.ID = 0
		_, err := backRepoTableStyle.db.Create(tablestyleDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoTableStyle.Map_TableStyleDBID_TableStyleDB[tablestyleDB.ID] = tablestyleDB
		BackRepoTableStyleid_atBckpTime_newID[tablestyleDB_ID_atBackupTime] = tablestyleDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "TableStyleDB.json" in dirPath that stores an array
// of TableStyleDB and stores it in the database
// the map BackRepoTableStyleid_atBckpTime_newID is updated accordingly
func (backRepoTableStyle *BackRepoTableStyleStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoTableStyleid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "TableStyleDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json TableStyle file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*TableStyleDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_TableStyleDBID_TableStyleDB
	for _, tablestyleDB := range forRestore {

		tablestyleDB_ID_atBackupTime := tablestyleDB.ID
		tablestyleDB.ID = 0
		_, err := backRepoTableStyle.db.Create(tablestyleDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoTableStyle.Map_TableStyleDBID_TableStyleDB[tablestyleDB.ID] = tablestyleDB
		BackRepoTableStyleid_atBckpTime_newID[tablestyleDB_ID_atBackupTime] = tablestyleDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json TableStyle file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<TableStyle>id_atBckpTime_newID
// to compute new index
func (backRepoTableStyle *BackRepoTableStyleStruct) RestorePhaseTwo() {

	for _, tablestyleDB := range backRepoTableStyle.Map_TableStyleDBID_TableStyleDB {

		// next line of code is to avert unused variable compilation error
		_ = tablestyleDB

		// insertion point for reindexing pointers encoding
		// reindexing Node field
		if tablestyleDB.NodeID.Int64 != 0 {
			tablestyleDB.NodeID.Int64 = int64(BackRepoNodeid_atBckpTime_newID[uint(tablestyleDB.NodeID.Int64)])
			tablestyleDB.NodeID.Valid = true
		}

		// update databse with new index encoding
		db, _ := backRepoTableStyle.db.Model(tablestyleDB)
		_, err := db.Updates(*tablestyleDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoTableStyle.ResetReversePointers commits all staged instances of TableStyle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTableStyle *BackRepoTableStyleStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, tablestyle := range backRepoTableStyle.Map_TableStyleDBID_TableStylePtr {
		backRepoTableStyle.ResetReversePointersInstance(backRepo, idx, tablestyle)
	}

	return
}

func (backRepoTableStyle *BackRepoTableStyleStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, tablestyle *models.TableStyle) (Error error) {

	// fetch matching tablestyleDB
	if tablestyleDB, ok := backRepoTableStyle.Map_TableStyleDBID_TableStyleDB[idx]; ok {
		_ = tablestyleDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoTableStyleid_atBckpTime_newID map[uint]uint
