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

	"github.com/fullstack-lang/gongsvg/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_SVG_sql sql.NullBool
var dummy_SVG_time time.Duration
var dummy_SVG_sort sort.Float64Slice

// SVGAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model svgAPI
type SVGAPI struct {
	gorm.Model

	models.SVG

	// encoding of pointers
	SVGPointersEnconding
}

// SVGPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SVGPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// field StartRect is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	StartRectID sql.NullInt64

	// field EndRect is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	EndRectID sql.NullInt64
}

// SVGDB describes a svg in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model svgDB
type SVGDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field svgDB.Name
	Name_Data sql.NullString

	// Declation for basic field svgDB.DrawingState
	DrawingState_Data sql.NullString

	// Declation for basic field svgDB.IsEditable
	// provide the sql storage for the boolan
	IsEditable_Data sql.NullBool
	// encoding of pointers
	SVGPointersEnconding
}

// SVGDBs arrays svgDBs
// swagger:response svgDBsResponse
type SVGDBs []SVGDB

// SVGDBResponse provides response
// swagger:response svgDBResponse
type SVGDBResponse struct {
	SVGDB
}

// SVGWOP is a SVG without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SVGWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DrawingState models.DrawingState `xlsx:"2"`

	IsEditable bool `xlsx:"3"`
	// insertion for WOP pointer fields
}

var SVG_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DrawingState",
	"IsEditable",
}

type BackRepoSVGStruct struct {
	// stores SVGDB according to their gorm ID
	Map_SVGDBID_SVGDB map[uint]*SVGDB

	// stores SVGDB ID according to SVG address
	Map_SVGPtr_SVGDBID map[*models.SVG]uint

	// stores SVG according to their gorm ID
	Map_SVGDBID_SVGPtr map[uint]*models.SVG

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoSVG *BackRepoSVGStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoSVG.stage
	return
}

func (backRepoSVG *BackRepoSVGStruct) GetDB() *gorm.DB {
	return backRepoSVG.db
}

// GetSVGDBFromSVGPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSVG *BackRepoSVGStruct) GetSVGDBFromSVGPtr(svg *models.SVG) (svgDB *SVGDB) {
	id := backRepoSVG.Map_SVGPtr_SVGDBID[svg]
	svgDB = backRepoSVG.Map_SVGDBID_SVGDB[id]
	return
}

// BackRepoSVG.CommitPhaseOne commits all staged instances of SVG to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSVG *BackRepoSVGStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for svg := range stage.SVGs {
		backRepoSVG.CommitPhaseOneInstance(svg)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, svg := range backRepoSVG.Map_SVGDBID_SVGPtr {
		if _, ok := stage.SVGs[svg]; !ok {
			backRepoSVG.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSVG.CommitDeleteInstance commits deletion of SVG to the BackRepo
func (backRepoSVG *BackRepoSVGStruct) CommitDeleteInstance(id uint) (Error error) {

	svg := backRepoSVG.Map_SVGDBID_SVGPtr[id]

	// svg is not staged anymore, remove svgDB
	svgDB := backRepoSVG.Map_SVGDBID_SVGDB[id]
	query := backRepoSVG.db.Unscoped().Delete(&svgDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoSVG.Map_SVGPtr_SVGDBID, svg)
	delete(backRepoSVG.Map_SVGDBID_SVGPtr, id)
	delete(backRepoSVG.Map_SVGDBID_SVGDB, id)

	return
}

// BackRepoSVG.CommitPhaseOneInstance commits svg staged instances of SVG to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSVG *BackRepoSVGStruct) CommitPhaseOneInstance(svg *models.SVG) (Error error) {

	// check if the svg is not commited yet
	if _, ok := backRepoSVG.Map_SVGPtr_SVGDBID[svg]; ok {
		return
	}

	// initiate svg
	var svgDB SVGDB
	svgDB.CopyBasicFieldsFromSVG(svg)

	query := backRepoSVG.db.Create(&svgDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoSVG.Map_SVGPtr_SVGDBID[svg] = svgDB.ID
	backRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID] = svg
	backRepoSVG.Map_SVGDBID_SVGDB[svgDB.ID] = &svgDB

	return
}

// BackRepoSVG.CommitPhaseTwo commits all staged instances of SVG to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVG *BackRepoSVGStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, svg := range backRepoSVG.Map_SVGDBID_SVGPtr {
		backRepoSVG.CommitPhaseTwoInstance(backRepo, idx, svg)
	}

	return
}

// BackRepoSVG.CommitPhaseTwoInstance commits {{structname }} of models.SVG to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVG *BackRepoSVGStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, svg *models.SVG) (Error error) {

	// fetch matching svgDB
	if svgDB, ok := backRepoSVG.Map_SVGDBID_SVGDB[idx]; ok {

		svgDB.CopyBasicFieldsFromSVG(svg)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers svg.Layers into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, layerAssocEnd := range svg.Layers {

			// get the back repo instance at the association end
			layerAssocEnd_DB :=
				backRepo.BackRepoLayer.GetLayerDBFromLayerPtr(layerAssocEnd)

			// encode reverse pointer in the association end back repo instance
			layerAssocEnd_DB.SVG_LayersDBID.Int64 = int64(svgDB.ID)
			layerAssocEnd_DB.SVG_LayersDBID.Valid = true
			layerAssocEnd_DB.SVG_LayersDBID_Index.Int64 = int64(idx)
			layerAssocEnd_DB.SVG_LayersDBID_Index.Valid = true
			if q := backRepoSVG.db.Save(layerAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		// commit pointer value svg.StartRect translates to updating the svg.StartRectID
		svgDB.StartRectID.Valid = true // allow for a 0 value (nil association)
		if svg.StartRect != nil {
			if StartRectId, ok := backRepo.BackRepoRect.Map_RectPtr_RectDBID[svg.StartRect]; ok {
				svgDB.StartRectID.Int64 = int64(StartRectId)
				svgDB.StartRectID.Valid = true
			}
		}

		// commit pointer value svg.EndRect translates to updating the svg.EndRectID
		svgDB.EndRectID.Valid = true // allow for a 0 value (nil association)
		if svg.EndRect != nil {
			if EndRectId, ok := backRepo.BackRepoRect.Map_RectPtr_RectDBID[svg.EndRect]; ok {
				svgDB.EndRectID.Int64 = int64(EndRectId)
				svgDB.EndRectID.Valid = true
			}
		}

		query := backRepoSVG.db.Save(&svgDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SVG intance %s", svg.Name))
		return err
	}

	return
}

// BackRepoSVG.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSVG *BackRepoSVGStruct) CheckoutPhaseOne() (Error error) {

	svgDBArray := make([]SVGDB, 0)
	query := backRepoSVG.db.Find(&svgDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	svgInstancesToBeRemovedFromTheStage := make(map[*models.SVG]any)
	for key, value := range backRepoSVG.stage.SVGs {
		svgInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, svgDB := range svgDBArray {
		backRepoSVG.CheckoutPhaseOneInstance(&svgDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		svg, ok := backRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID]
		if ok {
			delete(svgInstancesToBeRemovedFromTheStage, svg)
		}
	}

	// remove from stage and back repo's 3 maps all svgs that are not in the checkout
	for svg := range svgInstancesToBeRemovedFromTheStage {
		svg.Unstage(backRepoSVG.GetStage())

		// remove instance from the back repo 3 maps
		svgID := backRepoSVG.Map_SVGPtr_SVGDBID[svg]
		delete(backRepoSVG.Map_SVGPtr_SVGDBID, svg)
		delete(backRepoSVG.Map_SVGDBID_SVGDB, svgID)
		delete(backRepoSVG.Map_SVGDBID_SVGPtr, svgID)
	}

	return
}

// CheckoutPhaseOneInstance takes a svgDB that has been found in the DB, updates the backRepo and stages the
// models version of the svgDB
func (backRepoSVG *BackRepoSVGStruct) CheckoutPhaseOneInstance(svgDB *SVGDB) (Error error) {

	svg, ok := backRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID]
	if !ok {
		svg = new(models.SVG)

		backRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID] = svg
		backRepoSVG.Map_SVGPtr_SVGDBID[svg] = svgDB.ID

		// append model store with the new element
		svg.Name = svgDB.Name_Data.String
		svg.Stage(backRepoSVG.GetStage())
	}
	svgDB.CopyBasicFieldsToSVG(svg)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	svg.Stage(backRepoSVG.GetStage())

	// preserve pointer to svgDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SVGDBID_SVGDB)[svgDB hold variable pointers
	svgDB_Data := *svgDB
	preservedPtrToSVG := &svgDB_Data
	backRepoSVG.Map_SVGDBID_SVGDB[svgDB.ID] = preservedPtrToSVG

	return
}

// BackRepoSVG.CheckoutPhaseTwo Checkouts all staged instances of SVG to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVG *BackRepoSVGStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, svgDB := range backRepoSVG.Map_SVGDBID_SVGDB {
		backRepoSVG.CheckoutPhaseTwoInstance(backRepo, svgDB)
	}
	return
}

// BackRepoSVG.CheckoutPhaseTwoInstance Checkouts staged instances of SVG to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVG *BackRepoSVGStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, svgDB *SVGDB) (Error error) {

	svg := backRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID]
	_ = svg // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem svg.Layers in the stage from the encode in the back repo
	// It parses all LayerDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	svg.Layers = svg.Layers[:0]
	// 2. loop all instances in the type in the association end
	for _, layerDB_AssocEnd := range backRepo.BackRepoLayer.Map_LayerDBID_LayerDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if layerDB_AssocEnd.SVG_LayersDBID.Int64 == int64(svgDB.ID) {
			// 4. fetch the associated instance in the stage
			layer_AssocEnd := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[layerDB_AssocEnd.ID]
			// 5. append it the association slice
			svg.Layers = append(svg.Layers, layer_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(svg.Layers, func(i, j int) bool {
		layerDB_i_ID := backRepo.BackRepoLayer.Map_LayerPtr_LayerDBID[svg.Layers[i]]
		layerDB_j_ID := backRepo.BackRepoLayer.Map_LayerPtr_LayerDBID[svg.Layers[j]]

		layerDB_i := backRepo.BackRepoLayer.Map_LayerDBID_LayerDB[layerDB_i_ID]
		layerDB_j := backRepo.BackRepoLayer.Map_LayerDBID_LayerDB[layerDB_j_ID]

		return layerDB_i.SVG_LayersDBID_Index.Int64 < layerDB_j.SVG_LayersDBID_Index.Int64
	})

	// StartRect field
	if svgDB.StartRectID.Int64 != 0 {
		svg.StartRect = backRepo.BackRepoRect.Map_RectDBID_RectPtr[uint(svgDB.StartRectID.Int64)]
	}
	// EndRect field
	if svgDB.EndRectID.Int64 != 0 {
		svg.EndRect = backRepo.BackRepoRect.Map_RectDBID_RectPtr[uint(svgDB.EndRectID.Int64)]
	}
	return
}

// CommitSVG allows commit of a single svg (if already staged)
func (backRepo *BackRepoStruct) CommitSVG(svg *models.SVG) {
	backRepo.BackRepoSVG.CommitPhaseOneInstance(svg)
	if id, ok := backRepo.BackRepoSVG.Map_SVGPtr_SVGDBID[svg]; ok {
		backRepo.BackRepoSVG.CommitPhaseTwoInstance(backRepo, id, svg)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSVG allows checkout of a single svg (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSVG(svg *models.SVG) {
	// check if the svg is staged
	if _, ok := backRepo.BackRepoSVG.Map_SVGPtr_SVGDBID[svg]; ok {

		if id, ok := backRepo.BackRepoSVG.Map_SVGPtr_SVGDBID[svg]; ok {
			var svgDB SVGDB
			svgDB.ID = id

			if err := backRepo.BackRepoSVG.db.First(&svgDB, id).Error; err != nil {
				log.Panicln("CheckoutSVG : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSVG.CheckoutPhaseOneInstance(&svgDB)
			backRepo.BackRepoSVG.CheckoutPhaseTwoInstance(backRepo, &svgDB)
		}
	}
}

// CopyBasicFieldsFromSVG
func (svgDB *SVGDB) CopyBasicFieldsFromSVG(svg *models.SVG) {
	// insertion point for fields commit

	svgDB.Name_Data.String = svg.Name
	svgDB.Name_Data.Valid = true

	svgDB.DrawingState_Data.String = svg.DrawingState.ToString()
	svgDB.DrawingState_Data.Valid = true

	svgDB.IsEditable_Data.Bool = svg.IsEditable
	svgDB.IsEditable_Data.Valid = true
}

// CopyBasicFieldsFromSVGWOP
func (svgDB *SVGDB) CopyBasicFieldsFromSVGWOP(svg *SVGWOP) {
	// insertion point for fields commit

	svgDB.Name_Data.String = svg.Name
	svgDB.Name_Data.Valid = true

	svgDB.DrawingState_Data.String = svg.DrawingState.ToString()
	svgDB.DrawingState_Data.Valid = true

	svgDB.IsEditable_Data.Bool = svg.IsEditable
	svgDB.IsEditable_Data.Valid = true
}

// CopyBasicFieldsToSVG
func (svgDB *SVGDB) CopyBasicFieldsToSVG(svg *models.SVG) {
	// insertion point for checkout of basic fields (back repo to stage)
	svg.Name = svgDB.Name_Data.String
	svg.DrawingState.FromString(svgDB.DrawingState_Data.String)
	svg.IsEditable = svgDB.IsEditable_Data.Bool
}

// CopyBasicFieldsToSVGWOP
func (svgDB *SVGDB) CopyBasicFieldsToSVGWOP(svg *SVGWOP) {
	svg.ID = int(svgDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	svg.Name = svgDB.Name_Data.String
	svg.DrawingState.FromString(svgDB.DrawingState_Data.String)
	svg.IsEditable = svgDB.IsEditable_Data.Bool
}

// Backup generates a json file from a slice of all SVGDB instances in the backrepo
func (backRepoSVG *BackRepoSVGStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SVGDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SVGDB, 0)
	for _, svgDB := range backRepoSVG.Map_SVGDBID_SVGDB {
		forBackup = append(forBackup, svgDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json SVG ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json SVG file", err.Error())
	}
}

// Backup generates a json file from a slice of all SVGDB instances in the backrepo
func (backRepoSVG *BackRepoSVGStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SVGDB, 0)
	for _, svgDB := range backRepoSVG.Map_SVGDBID_SVGDB {
		forBackup = append(forBackup, svgDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SVG")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SVG_Fields, -1)
	for _, svgDB := range forBackup {

		var svgWOP SVGWOP
		svgDB.CopyBasicFieldsToSVGWOP(&svgWOP)

		row := sh.AddRow()
		row.WriteStruct(&svgWOP, -1)
	}
}

// RestoreXL from the "SVG" sheet all SVGDB instances
func (backRepoSVG *BackRepoSVGStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSVGid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SVG"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSVG.rowVisitorSVG)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoSVG *BackRepoSVGStruct) rowVisitorSVG(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var svgWOP SVGWOP
		row.ReadStruct(&svgWOP)

		// add the unmarshalled struct to the stage
		svgDB := new(SVGDB)
		svgDB.CopyBasicFieldsFromSVGWOP(&svgWOP)

		svgDB_ID_atBackupTime := svgDB.ID
		svgDB.ID = 0
		query := backRepoSVG.db.Create(svgDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoSVG.Map_SVGDBID_SVGDB[svgDB.ID] = svgDB
		BackRepoSVGid_atBckpTime_newID[svgDB_ID_atBackupTime] = svgDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SVGDB.json" in dirPath that stores an array
// of SVGDB and stores it in the database
// the map BackRepoSVGid_atBckpTime_newID is updated accordingly
func (backRepoSVG *BackRepoSVGStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSVGid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SVGDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json SVG file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SVGDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SVGDBID_SVGDB
	for _, svgDB := range forRestore {

		svgDB_ID_atBackupTime := svgDB.ID
		svgDB.ID = 0
		query := backRepoSVG.db.Create(svgDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoSVG.Map_SVGDBID_SVGDB[svgDB.ID] = svgDB
		BackRepoSVGid_atBckpTime_newID[svgDB_ID_atBackupTime] = svgDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json SVG file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SVG>id_atBckpTime_newID
// to compute new index
func (backRepoSVG *BackRepoSVGStruct) RestorePhaseTwo() {

	for _, svgDB := range backRepoSVG.Map_SVGDBID_SVGDB {

		// next line of code is to avert unused variable compilation error
		_ = svgDB

		// insertion point for reindexing pointers encoding
		// reindexing StartRect field
		if svgDB.StartRectID.Int64 != 0 {
			svgDB.StartRectID.Int64 = int64(BackRepoRectid_atBckpTime_newID[uint(svgDB.StartRectID.Int64)])
			svgDB.StartRectID.Valid = true
		}

		// reindexing EndRect field
		if svgDB.EndRectID.Int64 != 0 {
			svgDB.EndRectID.Int64 = int64(BackRepoRectid_atBckpTime_newID[uint(svgDB.EndRectID.Int64)])
			svgDB.EndRectID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoSVG.db.Model(svgDB).Updates(*svgDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSVGid_atBckpTime_newID map[uint]uint
