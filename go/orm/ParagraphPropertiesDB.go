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
var dummy_ParagraphProperties_sql sql.NullBool
var dummy_ParagraphProperties_time time.Duration
var dummy_ParagraphProperties_sort sort.Float64Slice

// ParagraphPropertiesAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model paragraphpropertiesAPI
type ParagraphPropertiesAPI struct {
	gorm.Model

	models.ParagraphProperties_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ParagraphPropertiesPointersEncoding ParagraphPropertiesPointersEncoding
}

// ParagraphPropertiesPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ParagraphPropertiesPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field ParagraphStyle is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ParagraphStyleID sql.NullInt64

	// field Node is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	NodeID sql.NullInt64
}

// ParagraphPropertiesDB describes a paragraphproperties in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model paragraphpropertiesDB
type ParagraphPropertiesDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field paragraphpropertiesDB.Name
	Name_Data sql.NullString

	// Declation for basic field paragraphpropertiesDB.Content
	Content_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ParagraphPropertiesPointersEncoding
}

// ParagraphPropertiesDBs arrays paragraphpropertiesDBs
// swagger:response paragraphpropertiesDBsResponse
type ParagraphPropertiesDBs []ParagraphPropertiesDB

// ParagraphPropertiesDBResponse provides response
// swagger:response paragraphpropertiesDBResponse
type ParagraphPropertiesDBResponse struct {
	ParagraphPropertiesDB
}

// ParagraphPropertiesWOP is a ParagraphProperties without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ParagraphPropertiesWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Content string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var ParagraphProperties_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Content",
}

type BackRepoParagraphPropertiesStruct struct {
	// stores ParagraphPropertiesDB according to their gorm ID
	Map_ParagraphPropertiesDBID_ParagraphPropertiesDB map[uint]*ParagraphPropertiesDB

	// stores ParagraphPropertiesDB ID according to ParagraphProperties address
	Map_ParagraphPropertiesPtr_ParagraphPropertiesDBID map[*models.ParagraphProperties]uint

	// stores ParagraphProperties according to their gorm ID
	Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr map[uint]*models.ParagraphProperties

	db db.DBInterface

	stage *models.StageStruct
}

func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoParagraphProperties.stage
	return
}

func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) GetDB() db.DBInterface {
	return backRepoParagraphProperties.db
}

// GetParagraphPropertiesDBFromParagraphPropertiesPtr is a handy function to access the back repo instance from the stage instance
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) GetParagraphPropertiesDBFromParagraphPropertiesPtr(paragraphproperties *models.ParagraphProperties) (paragraphpropertiesDB *ParagraphPropertiesDB) {
	id := backRepoParagraphProperties.Map_ParagraphPropertiesPtr_ParagraphPropertiesDBID[paragraphproperties]
	paragraphpropertiesDB = backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB[id]
	return
}

// BackRepoParagraphProperties.CommitPhaseOne commits all staged instances of ParagraphProperties to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for paragraphproperties := range stage.ParagraphPropertiess {
		backRepoParagraphProperties.CommitPhaseOneInstance(paragraphproperties)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, paragraphproperties := range backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr {
		if _, ok := stage.ParagraphPropertiess[paragraphproperties]; !ok {
			backRepoParagraphProperties.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoParagraphProperties.CommitDeleteInstance commits deletion of ParagraphProperties to the BackRepo
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) CommitDeleteInstance(id uint) (Error error) {

	paragraphproperties := backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr[id]

	// paragraphproperties is not staged anymore, remove paragraphpropertiesDB
	paragraphpropertiesDB := backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB[id]
	db, _ := backRepoParagraphProperties.db.Unscoped()
	_, err := db.Delete(paragraphpropertiesDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoParagraphProperties.Map_ParagraphPropertiesPtr_ParagraphPropertiesDBID, paragraphproperties)
	delete(backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr, id)
	delete(backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB, id)

	return
}

// BackRepoParagraphProperties.CommitPhaseOneInstance commits paragraphproperties staged instances of ParagraphProperties to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) CommitPhaseOneInstance(paragraphproperties *models.ParagraphProperties) (Error error) {

	// check if the paragraphproperties is not commited yet
	if _, ok := backRepoParagraphProperties.Map_ParagraphPropertiesPtr_ParagraphPropertiesDBID[paragraphproperties]; ok {
		return
	}

	// initiate paragraphproperties
	var paragraphpropertiesDB ParagraphPropertiesDB
	paragraphpropertiesDB.CopyBasicFieldsFromParagraphProperties(paragraphproperties)

	_, err := backRepoParagraphProperties.db.Create(&paragraphpropertiesDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoParagraphProperties.Map_ParagraphPropertiesPtr_ParagraphPropertiesDBID[paragraphproperties] = paragraphpropertiesDB.ID
	backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr[paragraphpropertiesDB.ID] = paragraphproperties
	backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB[paragraphpropertiesDB.ID] = &paragraphpropertiesDB

	return
}

// BackRepoParagraphProperties.CommitPhaseTwo commits all staged instances of ParagraphProperties to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, paragraphproperties := range backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr {
		backRepoParagraphProperties.CommitPhaseTwoInstance(backRepo, idx, paragraphproperties)
	}

	return
}

// BackRepoParagraphProperties.CommitPhaseTwoInstance commits {{structname }} of models.ParagraphProperties to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, paragraphproperties *models.ParagraphProperties) (Error error) {

	// fetch matching paragraphpropertiesDB
	if paragraphpropertiesDB, ok := backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB[idx]; ok {

		paragraphpropertiesDB.CopyBasicFieldsFromParagraphProperties(paragraphproperties)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value paragraphproperties.ParagraphStyle translates to updating the paragraphproperties.ParagraphStyleID
		paragraphpropertiesDB.ParagraphStyleID.Valid = true // allow for a 0 value (nil association)
		if paragraphproperties.ParagraphStyle != nil {
			if ParagraphStyleId, ok := backRepo.BackRepoParagraphStyle.Map_ParagraphStylePtr_ParagraphStyleDBID[paragraphproperties.ParagraphStyle]; ok {
				paragraphpropertiesDB.ParagraphStyleID.Int64 = int64(ParagraphStyleId)
				paragraphpropertiesDB.ParagraphStyleID.Valid = true
			}
		} else {
			paragraphpropertiesDB.ParagraphStyleID.Int64 = 0
			paragraphpropertiesDB.ParagraphStyleID.Valid = true
		}

		// commit pointer value paragraphproperties.Node translates to updating the paragraphproperties.NodeID
		paragraphpropertiesDB.NodeID.Valid = true // allow for a 0 value (nil association)
		if paragraphproperties.Node != nil {
			if NodeId, ok := backRepo.BackRepoNode.Map_NodePtr_NodeDBID[paragraphproperties.Node]; ok {
				paragraphpropertiesDB.NodeID.Int64 = int64(NodeId)
				paragraphpropertiesDB.NodeID.Valid = true
			}
		} else {
			paragraphpropertiesDB.NodeID.Int64 = 0
			paragraphpropertiesDB.NodeID.Valid = true
		}

		_, err := backRepoParagraphProperties.db.Save(paragraphpropertiesDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown ParagraphProperties intance %s", paragraphproperties.Name))
		return err
	}

	return
}

// BackRepoParagraphProperties.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) CheckoutPhaseOne() (Error error) {

	paragraphpropertiesDBArray := make([]ParagraphPropertiesDB, 0)
	_, err := backRepoParagraphProperties.db.Find(&paragraphpropertiesDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	paragraphpropertiesInstancesToBeRemovedFromTheStage := make(map[*models.ParagraphProperties]any)
	for key, value := range backRepoParagraphProperties.stage.ParagraphPropertiess {
		paragraphpropertiesInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, paragraphpropertiesDB := range paragraphpropertiesDBArray {
		backRepoParagraphProperties.CheckoutPhaseOneInstance(&paragraphpropertiesDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		paragraphproperties, ok := backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr[paragraphpropertiesDB.ID]
		if ok {
			delete(paragraphpropertiesInstancesToBeRemovedFromTheStage, paragraphproperties)
		}
	}

	// remove from stage and back repo's 3 maps all paragraphpropertiess that are not in the checkout
	for paragraphproperties := range paragraphpropertiesInstancesToBeRemovedFromTheStage {
		paragraphproperties.Unstage(backRepoParagraphProperties.GetStage())

		// remove instance from the back repo 3 maps
		paragraphpropertiesID := backRepoParagraphProperties.Map_ParagraphPropertiesPtr_ParagraphPropertiesDBID[paragraphproperties]
		delete(backRepoParagraphProperties.Map_ParagraphPropertiesPtr_ParagraphPropertiesDBID, paragraphproperties)
		delete(backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB, paragraphpropertiesID)
		delete(backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr, paragraphpropertiesID)
	}

	return
}

// CheckoutPhaseOneInstance takes a paragraphpropertiesDB that has been found in the DB, updates the backRepo and stages the
// models version of the paragraphpropertiesDB
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) CheckoutPhaseOneInstance(paragraphpropertiesDB *ParagraphPropertiesDB) (Error error) {

	paragraphproperties, ok := backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr[paragraphpropertiesDB.ID]
	if !ok {
		paragraphproperties = new(models.ParagraphProperties)

		backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr[paragraphpropertiesDB.ID] = paragraphproperties
		backRepoParagraphProperties.Map_ParagraphPropertiesPtr_ParagraphPropertiesDBID[paragraphproperties] = paragraphpropertiesDB.ID

		// append model store with the new element
		paragraphproperties.Name = paragraphpropertiesDB.Name_Data.String
		paragraphproperties.Stage(backRepoParagraphProperties.GetStage())
	}
	paragraphpropertiesDB.CopyBasicFieldsToParagraphProperties(paragraphproperties)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	paragraphproperties.Stage(backRepoParagraphProperties.GetStage())

	// preserve pointer to paragraphpropertiesDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ParagraphPropertiesDBID_ParagraphPropertiesDB)[paragraphpropertiesDB hold variable pointers
	paragraphpropertiesDB_Data := *paragraphpropertiesDB
	preservedPtrToParagraphProperties := &paragraphpropertiesDB_Data
	backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB[paragraphpropertiesDB.ID] = preservedPtrToParagraphProperties

	return
}

// BackRepoParagraphProperties.CheckoutPhaseTwo Checkouts all staged instances of ParagraphProperties to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, paragraphpropertiesDB := range backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB {
		backRepoParagraphProperties.CheckoutPhaseTwoInstance(backRepo, paragraphpropertiesDB)
	}
	return
}

// BackRepoParagraphProperties.CheckoutPhaseTwoInstance Checkouts staged instances of ParagraphProperties to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, paragraphpropertiesDB *ParagraphPropertiesDB) (Error error) {

	paragraphproperties := backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr[paragraphpropertiesDB.ID]

	paragraphpropertiesDB.DecodePointers(backRepo, paragraphproperties)

	return
}

func (paragraphpropertiesDB *ParagraphPropertiesDB) DecodePointers(backRepo *BackRepoStruct, paragraphproperties *models.ParagraphProperties) {

	// insertion point for checkout of pointer encoding
	// ParagraphStyle field	
	{
		id := paragraphpropertiesDB.ParagraphStyleID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoParagraphStyle.Map_ParagraphStyleDBID_ParagraphStylePtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: paragraphproperties.ParagraphStyle, unknown pointer id", id)
				paragraphproperties.ParagraphStyle = nil
			} else {
				// updates only if field has changed
				if paragraphproperties.ParagraphStyle == nil || paragraphproperties.ParagraphStyle != tmp {
					paragraphproperties.ParagraphStyle = tmp
				}
			}
		} else {
			paragraphproperties.ParagraphStyle = nil
		}
	}
	
	// Node field	
	{
		id := paragraphpropertiesDB.NodeID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: paragraphproperties.Node, unknown pointer id", id)
				paragraphproperties.Node = nil
			} else {
				// updates only if field has changed
				if paragraphproperties.Node == nil || paragraphproperties.Node != tmp {
					paragraphproperties.Node = tmp
				}
			}
		} else {
			paragraphproperties.Node = nil
		}
	}
	
	return
}

// CommitParagraphProperties allows commit of a single paragraphproperties (if already staged)
func (backRepo *BackRepoStruct) CommitParagraphProperties(paragraphproperties *models.ParagraphProperties) {
	backRepo.BackRepoParagraphProperties.CommitPhaseOneInstance(paragraphproperties)
	if id, ok := backRepo.BackRepoParagraphProperties.Map_ParagraphPropertiesPtr_ParagraphPropertiesDBID[paragraphproperties]; ok {
		backRepo.BackRepoParagraphProperties.CommitPhaseTwoInstance(backRepo, id, paragraphproperties)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitParagraphProperties allows checkout of a single paragraphproperties (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutParagraphProperties(paragraphproperties *models.ParagraphProperties) {
	// check if the paragraphproperties is staged
	if _, ok := backRepo.BackRepoParagraphProperties.Map_ParagraphPropertiesPtr_ParagraphPropertiesDBID[paragraphproperties]; ok {

		if id, ok := backRepo.BackRepoParagraphProperties.Map_ParagraphPropertiesPtr_ParagraphPropertiesDBID[paragraphproperties]; ok {
			var paragraphpropertiesDB ParagraphPropertiesDB
			paragraphpropertiesDB.ID = id

			if _, err := backRepo.BackRepoParagraphProperties.db.First(&paragraphpropertiesDB, id); err != nil {
				log.Fatalln("CheckoutParagraphProperties : Problem with getting object with id:", id)
			}
			backRepo.BackRepoParagraphProperties.CheckoutPhaseOneInstance(&paragraphpropertiesDB)
			backRepo.BackRepoParagraphProperties.CheckoutPhaseTwoInstance(backRepo, &paragraphpropertiesDB)
		}
	}
}

// CopyBasicFieldsFromParagraphProperties
func (paragraphpropertiesDB *ParagraphPropertiesDB) CopyBasicFieldsFromParagraphProperties(paragraphproperties *models.ParagraphProperties) {
	// insertion point for fields commit

	paragraphpropertiesDB.Name_Data.String = paragraphproperties.Name
	paragraphpropertiesDB.Name_Data.Valid = true

	paragraphpropertiesDB.Content_Data.String = paragraphproperties.Content
	paragraphpropertiesDB.Content_Data.Valid = true
}

// CopyBasicFieldsFromParagraphProperties_WOP
func (paragraphpropertiesDB *ParagraphPropertiesDB) CopyBasicFieldsFromParagraphProperties_WOP(paragraphproperties *models.ParagraphProperties_WOP) {
	// insertion point for fields commit

	paragraphpropertiesDB.Name_Data.String = paragraphproperties.Name
	paragraphpropertiesDB.Name_Data.Valid = true

	paragraphpropertiesDB.Content_Data.String = paragraphproperties.Content
	paragraphpropertiesDB.Content_Data.Valid = true
}

// CopyBasicFieldsFromParagraphPropertiesWOP
func (paragraphpropertiesDB *ParagraphPropertiesDB) CopyBasicFieldsFromParagraphPropertiesWOP(paragraphproperties *ParagraphPropertiesWOP) {
	// insertion point for fields commit

	paragraphpropertiesDB.Name_Data.String = paragraphproperties.Name
	paragraphpropertiesDB.Name_Data.Valid = true

	paragraphpropertiesDB.Content_Data.String = paragraphproperties.Content
	paragraphpropertiesDB.Content_Data.Valid = true
}

// CopyBasicFieldsToParagraphProperties
func (paragraphpropertiesDB *ParagraphPropertiesDB) CopyBasicFieldsToParagraphProperties(paragraphproperties *models.ParagraphProperties) {
	// insertion point for checkout of basic fields (back repo to stage)
	paragraphproperties.Name = paragraphpropertiesDB.Name_Data.String
	paragraphproperties.Content = paragraphpropertiesDB.Content_Data.String
}

// CopyBasicFieldsToParagraphProperties_WOP
func (paragraphpropertiesDB *ParagraphPropertiesDB) CopyBasicFieldsToParagraphProperties_WOP(paragraphproperties *models.ParagraphProperties_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	paragraphproperties.Name = paragraphpropertiesDB.Name_Data.String
	paragraphproperties.Content = paragraphpropertiesDB.Content_Data.String
}

// CopyBasicFieldsToParagraphPropertiesWOP
func (paragraphpropertiesDB *ParagraphPropertiesDB) CopyBasicFieldsToParagraphPropertiesWOP(paragraphproperties *ParagraphPropertiesWOP) {
	paragraphproperties.ID = int(paragraphpropertiesDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	paragraphproperties.Name = paragraphpropertiesDB.Name_Data.String
	paragraphproperties.Content = paragraphpropertiesDB.Content_Data.String
}

// Backup generates a json file from a slice of all ParagraphPropertiesDB instances in the backrepo
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ParagraphPropertiesDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ParagraphPropertiesDB, 0)
	for _, paragraphpropertiesDB := range backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB {
		forBackup = append(forBackup, paragraphpropertiesDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json ParagraphProperties ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json ParagraphProperties file", err.Error())
	}
}

// Backup generates a json file from a slice of all ParagraphPropertiesDB instances in the backrepo
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ParagraphPropertiesDB, 0)
	for _, paragraphpropertiesDB := range backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB {
		forBackup = append(forBackup, paragraphpropertiesDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("ParagraphProperties")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&ParagraphProperties_Fields, -1)
	for _, paragraphpropertiesDB := range forBackup {

		var paragraphpropertiesWOP ParagraphPropertiesWOP
		paragraphpropertiesDB.CopyBasicFieldsToParagraphPropertiesWOP(&paragraphpropertiesWOP)

		row := sh.AddRow()
		row.WriteStruct(&paragraphpropertiesWOP, -1)
	}
}

// RestoreXL from the "ParagraphProperties" sheet all ParagraphPropertiesDB instances
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoParagraphPropertiesid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["ParagraphProperties"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoParagraphProperties.rowVisitorParagraphProperties)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) rowVisitorParagraphProperties(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var paragraphpropertiesWOP ParagraphPropertiesWOP
		row.ReadStruct(&paragraphpropertiesWOP)

		// add the unmarshalled struct to the stage
		paragraphpropertiesDB := new(ParagraphPropertiesDB)
		paragraphpropertiesDB.CopyBasicFieldsFromParagraphPropertiesWOP(&paragraphpropertiesWOP)

		paragraphpropertiesDB_ID_atBackupTime := paragraphpropertiesDB.ID
		paragraphpropertiesDB.ID = 0
		_, err := backRepoParagraphProperties.db.Create(paragraphpropertiesDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB[paragraphpropertiesDB.ID] = paragraphpropertiesDB
		BackRepoParagraphPropertiesid_atBckpTime_newID[paragraphpropertiesDB_ID_atBackupTime] = paragraphpropertiesDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ParagraphPropertiesDB.json" in dirPath that stores an array
// of ParagraphPropertiesDB and stores it in the database
// the map BackRepoParagraphPropertiesid_atBckpTime_newID is updated accordingly
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoParagraphPropertiesid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ParagraphPropertiesDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json ParagraphProperties file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ParagraphPropertiesDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ParagraphPropertiesDBID_ParagraphPropertiesDB
	for _, paragraphpropertiesDB := range forRestore {

		paragraphpropertiesDB_ID_atBackupTime := paragraphpropertiesDB.ID
		paragraphpropertiesDB.ID = 0
		_, err := backRepoParagraphProperties.db.Create(paragraphpropertiesDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB[paragraphpropertiesDB.ID] = paragraphpropertiesDB
		BackRepoParagraphPropertiesid_atBckpTime_newID[paragraphpropertiesDB_ID_atBackupTime] = paragraphpropertiesDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json ParagraphProperties file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<ParagraphProperties>id_atBckpTime_newID
// to compute new index
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) RestorePhaseTwo() {

	for _, paragraphpropertiesDB := range backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB {

		// next line of code is to avert unused variable compilation error
		_ = paragraphpropertiesDB

		// insertion point for reindexing pointers encoding
		// reindexing ParagraphStyle field
		if paragraphpropertiesDB.ParagraphStyleID.Int64 != 0 {
			paragraphpropertiesDB.ParagraphStyleID.Int64 = int64(BackRepoParagraphStyleid_atBckpTime_newID[uint(paragraphpropertiesDB.ParagraphStyleID.Int64)])
			paragraphpropertiesDB.ParagraphStyleID.Valid = true
		}

		// reindexing Node field
		if paragraphpropertiesDB.NodeID.Int64 != 0 {
			paragraphpropertiesDB.NodeID.Int64 = int64(BackRepoNodeid_atBckpTime_newID[uint(paragraphpropertiesDB.NodeID.Int64)])
			paragraphpropertiesDB.NodeID.Valid = true
		}

		// update databse with new index encoding
		db, _ := backRepoParagraphProperties.db.Model(paragraphpropertiesDB)
		_, err := db.Updates(*paragraphpropertiesDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoParagraphProperties.ResetReversePointers commits all staged instances of ParagraphProperties to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, paragraphproperties := range backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr {
		backRepoParagraphProperties.ResetReversePointersInstance(backRepo, idx, paragraphproperties)
	}

	return
}

func (backRepoParagraphProperties *BackRepoParagraphPropertiesStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, paragraphproperties *models.ParagraphProperties) (Error error) {

	// fetch matching paragraphpropertiesDB
	if paragraphpropertiesDB, ok := backRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesDB[idx]; ok {
		_ = paragraphpropertiesDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoParagraphPropertiesid_atBckpTime_newID map[uint]uint
