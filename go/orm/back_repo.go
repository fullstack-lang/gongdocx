package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gongdocx/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoDocument BackRepoDocumentStruct

	BackRepoDocx BackRepoDocxStruct

	BackRepoFile BackRepoFileStruct

	BackRepoNode BackRepoNodeStruct

	BackRepoText BackRepoTextStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	if err != nil {
		panic("Failed to connect to database!")
	}

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		&DocumentDB{},
		&DocxDB{},
		&FileDB{},
		&NodeDB{},
		&TextDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoDocument = BackRepoDocumentStruct{
		Map_DocumentDBID_DocumentPtr: make(map[uint]*models.Document, 0),
		Map_DocumentDBID_DocumentDB:  make(map[uint]*DocumentDB, 0),
		Map_DocumentPtr_DocumentDBID: make(map[*models.Document]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDocx = BackRepoDocxStruct{
		Map_DocxDBID_DocxPtr: make(map[uint]*models.Docx, 0),
		Map_DocxDBID_DocxDB:  make(map[uint]*DocxDB, 0),
		Map_DocxPtr_DocxDBID: make(map[*models.Docx]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFile = BackRepoFileStruct{
		Map_FileDBID_FilePtr: make(map[uint]*models.File, 0),
		Map_FileDBID_FileDB:  make(map[uint]*FileDB, 0),
		Map_FilePtr_FileDBID: make(map[*models.File]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNode = BackRepoNodeStruct{
		Map_NodeDBID_NodePtr: make(map[uint]*models.Node, 0),
		Map_NodeDBID_NodeDB:  make(map[uint]*NodeDB, 0),
		Map_NodePtr_NodeDBID: make(map[*models.Node]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoText = BackRepoTextStruct{
		Map_TextDBID_TextPtr: make(map[uint]*models.Text, 0),
		Map_TextDBID_TextDB:  make(map[uint]*TextDB, 0),
		Map_TextPtr_TextDBID: make(map[*models.Text]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoDocument.CommitPhaseOne(stage)
	backRepo.BackRepoDocx.CommitPhaseOne(stage)
	backRepo.BackRepoFile.CommitPhaseOne(stage)
	backRepo.BackRepoNode.CommitPhaseOne(stage)
	backRepo.BackRepoText.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoDocument.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDocx.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFile.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNode.CommitPhaseTwo(backRepo)
	backRepo.BackRepoText.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoDocument.CheckoutPhaseOne()
	backRepo.BackRepoDocx.CheckoutPhaseOne()
	backRepo.BackRepoFile.CheckoutPhaseOne()
	backRepo.BackRepoNode.CheckoutPhaseOne()
	backRepo.BackRepoText.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoDocument.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDocx.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFile.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNode.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoText.CheckoutPhaseTwo(backRepo)
}

var _backRepo *BackRepoStruct

var once sync.Once

func GetDefaultBackRepo() *BackRepoStruct {
	once.Do(func() {
		_backRepo = NewBackRepo(models.GetDefaultStage(), "")
	})
	return _backRepo
}

func GetLastCommitFromBackNb() uint {
	return GetDefaultBackRepo().GetLastCommitFromBackNb()
}

func GetLastPushFromFrontNb() uint {
	return GetDefaultBackRepo().GetLastPushFromFrontNb()
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoDocument.Backup(dirPath)
	backRepo.BackRepoDocx.Backup(dirPath)
	backRepo.BackRepoFile.Backup(dirPath)
	backRepo.BackRepoNode.Backup(dirPath)
	backRepo.BackRepoText.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoDocument.BackupXL(file)
	backRepo.BackRepoDocx.BackupXL(file)
	backRepo.BackRepoFile.BackupXL(file)
	backRepo.BackRepoNode.BackupXL(file)
	backRepo.BackRepoText.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoDocument.RestorePhaseOne(dirPath)
	backRepo.BackRepoDocx.RestorePhaseOne(dirPath)
	backRepo.BackRepoFile.RestorePhaseOne(dirPath)
	backRepo.BackRepoNode.RestorePhaseOne(dirPath)
	backRepo.BackRepoText.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoDocument.RestorePhaseTwo()
	backRepo.BackRepoDocx.RestorePhaseTwo()
	backRepo.BackRepoFile.RestorePhaseTwo()
	backRepo.BackRepoNode.RestorePhaseTwo()
	backRepo.BackRepoText.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoDocument.RestoreXLPhaseOne(file)
	backRepo.BackRepoDocx.RestoreXLPhaseOne(file)
	backRepo.BackRepoFile.RestoreXLPhaseOne(file)
	backRepo.BackRepoNode.RestoreXLPhaseOne(file)
	backRepo.BackRepoText.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}
