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

	"github.com/fullstack-lang/gongdocx/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Rune_sql sql.NullBool
var dummy_Rune_time time.Duration
var dummy_Rune_sort sort.Float64Slice

// RuneAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model runeAPI
type RuneAPI struct {
	gorm.Model

	models.Rune

	// encoding of pointers
	RunePointersEnconding
}

// RunePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type RunePointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// field Node is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	NodeID sql.NullInt64
}

// RuneDB describes a rune in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model runeDB
type RuneDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field runeDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	RunePointersEnconding
}

// RuneDBs arrays runeDBs
// swagger:response runeDBsResponse
type RuneDBs []RuneDB

// RuneDBResponse provides response
// swagger:response runeDBResponse
type RuneDBResponse struct {
	RuneDB
}

// RuneWOP is a Rune without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type RuneWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Rune_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoRuneStruct struct {
	// stores RuneDB according to their gorm ID
	Map_RuneDBID_RuneDB map[uint]*RuneDB

	// stores RuneDB ID according to Rune address
	Map_RunePtr_RuneDBID map[*models.Rune]uint

	// stores Rune according to their gorm ID
	Map_RuneDBID_RunePtr map[uint]*models.Rune

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoRune *BackRepoRuneStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoRune.stage
	return
}

func (backRepoRune *BackRepoRuneStruct) GetDB() *gorm.DB {
	return backRepoRune.db
}

// GetRuneDBFromRunePtr is a handy function to access the back repo instance from the stage instance
func (backRepoRune *BackRepoRuneStruct) GetRuneDBFromRunePtr(rune *models.Rune) (runeDB *RuneDB) {
	id := backRepoRune.Map_RunePtr_RuneDBID[rune]
	runeDB = backRepoRune.Map_RuneDBID_RuneDB[id]
	return
}

// BackRepoRune.CommitPhaseOne commits all staged instances of Rune to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoRune *BackRepoRuneStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for rune := range stage.Runes {
		backRepoRune.CommitPhaseOneInstance(rune)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, rune := range backRepoRune.Map_RuneDBID_RunePtr {
		if _, ok := stage.Runes[rune]; !ok {
			backRepoRune.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoRune.CommitDeleteInstance commits deletion of Rune to the BackRepo
func (backRepoRune *BackRepoRuneStruct) CommitDeleteInstance(id uint) (Error error) {

	rune := backRepoRune.Map_RuneDBID_RunePtr[id]

	// rune is not staged anymore, remove runeDB
	runeDB := backRepoRune.Map_RuneDBID_RuneDB[id]
	query := backRepoRune.db.Unscoped().Delete(&runeDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoRune.Map_RunePtr_RuneDBID, rune)
	delete(backRepoRune.Map_RuneDBID_RunePtr, id)
	delete(backRepoRune.Map_RuneDBID_RuneDB, id)

	return
}

// BackRepoRune.CommitPhaseOneInstance commits rune staged instances of Rune to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoRune *BackRepoRuneStruct) CommitPhaseOneInstance(rune *models.Rune) (Error error) {

	// check if the rune is not commited yet
	if _, ok := backRepoRune.Map_RunePtr_RuneDBID[rune]; ok {
		return
	}

	// initiate rune
	var runeDB RuneDB
	runeDB.CopyBasicFieldsFromRune(rune)

	query := backRepoRune.db.Create(&runeDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoRune.Map_RunePtr_RuneDBID[rune] = runeDB.ID
	backRepoRune.Map_RuneDBID_RunePtr[runeDB.ID] = rune
	backRepoRune.Map_RuneDBID_RuneDB[runeDB.ID] = &runeDB

	return
}

// BackRepoRune.CommitPhaseTwo commits all staged instances of Rune to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRune *BackRepoRuneStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, rune := range backRepoRune.Map_RuneDBID_RunePtr {
		backRepoRune.CommitPhaseTwoInstance(backRepo, idx, rune)
	}

	return
}

// BackRepoRune.CommitPhaseTwoInstance commits {{structname }} of models.Rune to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRune *BackRepoRuneStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, rune *models.Rune) (Error error) {

	// fetch matching runeDB
	if runeDB, ok := backRepoRune.Map_RuneDBID_RuneDB[idx]; ok {

		runeDB.CopyBasicFieldsFromRune(rune)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value rune.Node translates to updating the rune.NodeID
		runeDB.NodeID.Valid = true // allow for a 0 value (nil association)
		if rune.Node != nil {
			if NodeId, ok := backRepo.BackRepoNode.Map_NodePtr_NodeDBID[rune.Node]; ok {
				runeDB.NodeID.Int64 = int64(NodeId)
				runeDB.NodeID.Valid = true
			}
		}

		query := backRepoRune.db.Save(&runeDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Rune intance %s", rune.Name))
		return err
	}

	return
}

// BackRepoRune.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoRune *BackRepoRuneStruct) CheckoutPhaseOne() (Error error) {

	runeDBArray := make([]RuneDB, 0)
	query := backRepoRune.db.Find(&runeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	runeInstancesToBeRemovedFromTheStage := make(map[*models.Rune]any)
	for key, value := range backRepoRune.stage.Runes {
		runeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, runeDB := range runeDBArray {
		backRepoRune.CheckoutPhaseOneInstance(&runeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		rune, ok := backRepoRune.Map_RuneDBID_RunePtr[runeDB.ID]
		if ok {
			delete(runeInstancesToBeRemovedFromTheStage, rune)
		}
	}

	// remove from stage and back repo's 3 maps all runes that are not in the checkout
	for rune := range runeInstancesToBeRemovedFromTheStage {
		rune.Unstage(backRepoRune.GetStage())

		// remove instance from the back repo 3 maps
		runeID := backRepoRune.Map_RunePtr_RuneDBID[rune]
		delete(backRepoRune.Map_RunePtr_RuneDBID, rune)
		delete(backRepoRune.Map_RuneDBID_RuneDB, runeID)
		delete(backRepoRune.Map_RuneDBID_RunePtr, runeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a runeDB that has been found in the DB, updates the backRepo and stages the
// models version of the runeDB
func (backRepoRune *BackRepoRuneStruct) CheckoutPhaseOneInstance(runeDB *RuneDB) (Error error) {

	rune, ok := backRepoRune.Map_RuneDBID_RunePtr[runeDB.ID]
	if !ok {
		rune = new(models.Rune)

		backRepoRune.Map_RuneDBID_RunePtr[runeDB.ID] = rune
		backRepoRune.Map_RunePtr_RuneDBID[rune] = runeDB.ID

		// append model store with the new element
		rune.Name = runeDB.Name_Data.String
		rune.Stage(backRepoRune.GetStage())
	}
	runeDB.CopyBasicFieldsToRune(rune)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	rune.Stage(backRepoRune.GetStage())

	// preserve pointer to runeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_RuneDBID_RuneDB)[runeDB hold variable pointers
	runeDB_Data := *runeDB
	preservedPtrToRune := &runeDB_Data
	backRepoRune.Map_RuneDBID_RuneDB[runeDB.ID] = preservedPtrToRune

	return
}

// BackRepoRune.CheckoutPhaseTwo Checkouts all staged instances of Rune to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRune *BackRepoRuneStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, runeDB := range backRepoRune.Map_RuneDBID_RuneDB {
		backRepoRune.CheckoutPhaseTwoInstance(backRepo, runeDB)
	}
	return
}

// BackRepoRune.CheckoutPhaseTwoInstance Checkouts staged instances of Rune to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRune *BackRepoRuneStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, runeDB *RuneDB) (Error error) {

	rune := backRepoRune.Map_RuneDBID_RunePtr[runeDB.ID]
	_ = rune // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// Node field
	if runeDB.NodeID.Int64 != 0 {
		rune.Node = backRepo.BackRepoNode.Map_NodeDBID_NodePtr[uint(runeDB.NodeID.Int64)]
	}
	return
}

// CommitRune allows commit of a single rune (if already staged)
func (backRepo *BackRepoStruct) CommitRune(rune *models.Rune) {
	backRepo.BackRepoRune.CommitPhaseOneInstance(rune)
	if id, ok := backRepo.BackRepoRune.Map_RunePtr_RuneDBID[rune]; ok {
		backRepo.BackRepoRune.CommitPhaseTwoInstance(backRepo, id, rune)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitRune allows checkout of a single rune (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutRune(rune *models.Rune) {
	// check if the rune is staged
	if _, ok := backRepo.BackRepoRune.Map_RunePtr_RuneDBID[rune]; ok {

		if id, ok := backRepo.BackRepoRune.Map_RunePtr_RuneDBID[rune]; ok {
			var runeDB RuneDB
			runeDB.ID = id

			if err := backRepo.BackRepoRune.db.First(&runeDB, id).Error; err != nil {
				log.Panicln("CheckoutRune : Problem with getting object with id:", id)
			}
			backRepo.BackRepoRune.CheckoutPhaseOneInstance(&runeDB)
			backRepo.BackRepoRune.CheckoutPhaseTwoInstance(backRepo, &runeDB)
		}
	}
}

// CopyBasicFieldsFromRune
func (runeDB *RuneDB) CopyBasicFieldsFromRune(rune *models.Rune) {
	// insertion point for fields commit

	runeDB.Name_Data.String = rune.Name
	runeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromRuneWOP
func (runeDB *RuneDB) CopyBasicFieldsFromRuneWOP(rune *RuneWOP) {
	// insertion point for fields commit

	runeDB.Name_Data.String = rune.Name
	runeDB.Name_Data.Valid = true
}

// CopyBasicFieldsToRune
func (runeDB *RuneDB) CopyBasicFieldsToRune(rune *models.Rune) {
	// insertion point for checkout of basic fields (back repo to stage)
	rune.Name = runeDB.Name_Data.String
}

// CopyBasicFieldsToRuneWOP
func (runeDB *RuneDB) CopyBasicFieldsToRuneWOP(rune *RuneWOP) {
	rune.ID = int(runeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	rune.Name = runeDB.Name_Data.String
}

// Backup generates a json file from a slice of all RuneDB instances in the backrepo
func (backRepoRune *BackRepoRuneStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "RuneDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*RuneDB, 0)
	for _, runeDB := range backRepoRune.Map_RuneDBID_RuneDB {
		forBackup = append(forBackup, runeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Rune ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Rune file", err.Error())
	}
}

// Backup generates a json file from a slice of all RuneDB instances in the backrepo
func (backRepoRune *BackRepoRuneStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*RuneDB, 0)
	for _, runeDB := range backRepoRune.Map_RuneDBID_RuneDB {
		forBackup = append(forBackup, runeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Rune")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Rune_Fields, -1)
	for _, runeDB := range forBackup {

		var runeWOP RuneWOP
		runeDB.CopyBasicFieldsToRuneWOP(&runeWOP)

		row := sh.AddRow()
		row.WriteStruct(&runeWOP, -1)
	}
}

// RestoreXL from the "Rune" sheet all RuneDB instances
func (backRepoRune *BackRepoRuneStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoRuneid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Rune"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoRune.rowVisitorRune)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoRune *BackRepoRuneStruct) rowVisitorRune(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var runeWOP RuneWOP
		row.ReadStruct(&runeWOP)

		// add the unmarshalled struct to the stage
		runeDB := new(RuneDB)
		runeDB.CopyBasicFieldsFromRuneWOP(&runeWOP)

		runeDB_ID_atBackupTime := runeDB.ID
		runeDB.ID = 0
		query := backRepoRune.db.Create(runeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoRune.Map_RuneDBID_RuneDB[runeDB.ID] = runeDB
		BackRepoRuneid_atBckpTime_newID[runeDB_ID_atBackupTime] = runeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "RuneDB.json" in dirPath that stores an array
// of RuneDB and stores it in the database
// the map BackRepoRuneid_atBckpTime_newID is updated accordingly
func (backRepoRune *BackRepoRuneStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoRuneid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "RuneDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Rune file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*RuneDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_RuneDBID_RuneDB
	for _, runeDB := range forRestore {

		runeDB_ID_atBackupTime := runeDB.ID
		runeDB.ID = 0
		query := backRepoRune.db.Create(runeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoRune.Map_RuneDBID_RuneDB[runeDB.ID] = runeDB
		BackRepoRuneid_atBckpTime_newID[runeDB_ID_atBackupTime] = runeDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Rune file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Rune>id_atBckpTime_newID
// to compute new index
func (backRepoRune *BackRepoRuneStruct) RestorePhaseTwo() {

	for _, runeDB := range backRepoRune.Map_RuneDBID_RuneDB {

		// next line of code is to avert unused variable compilation error
		_ = runeDB

		// insertion point for reindexing pointers encoding
		// reindexing Node field
		if runeDB.NodeID.Int64 != 0 {
			runeDB.NodeID.Int64 = int64(BackRepoNodeid_atBckpTime_newID[uint(runeDB.NodeID.Int64)])
			runeDB.NodeID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoRune.db.Model(runeDB).Updates(*runeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoRuneid_atBckpTime_newID map[uint]uint