// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gongdocx/go/db"
	"github.com/fullstack-lang/gongdocx/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gongdocx/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoBody BackRepoBodyStruct

	BackRepoDocument BackRepoDocumentStruct

	BackRepoDocx BackRepoDocxStruct

	BackRepoFile BackRepoFileStruct

	BackRepoNode BackRepoNodeStruct

	BackRepoParagraph BackRepoParagraphStruct

	BackRepoParagraphProperties BackRepoParagraphPropertiesStruct

	BackRepoParagraphStyle BackRepoParagraphStyleStruct

	BackRepoRune BackRepoRuneStruct

	BackRepoRuneProperties BackRepoRunePropertiesStruct

	BackRepoTable BackRepoTableStruct

	BackRepoTableColumn BackRepoTableColumnStruct

	BackRepoTableProperties BackRepoTablePropertiesStruct

	BackRepoTableRow BackRepoTableRowStruct

	BackRepoTableStyle BackRepoTableStyleStruct

	BackRepoText BackRepoTextStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	var db db.DBInterface

	db = NewDBLite()

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gongdocx_go",
		&BodyDB{},
		&DocumentDB{},
		&DocxDB{},
		&FileDB{},
		&NodeDB{},
		&ParagraphDB{},
		&ParagraphPropertiesDB{},
		&ParagraphStyleDB{},
		&RuneDB{},
		&RunePropertiesDB{},
		&TableDB{},
		&TableColumnDB{},
		&TablePropertiesDB{},
		&TableRowDB{},
		&TableStyleDB{},
		&TextDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoBody = BackRepoBodyStruct{
		Map_BodyDBID_BodyPtr: make(map[uint]*models.Body, 0),
		Map_BodyDBID_BodyDB:  make(map[uint]*BodyDB, 0),
		Map_BodyPtr_BodyDBID: make(map[*models.Body]uint, 0),

		db:    db,
		stage: stage,
	}
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
	backRepo.BackRepoParagraph = BackRepoParagraphStruct{
		Map_ParagraphDBID_ParagraphPtr: make(map[uint]*models.Paragraph, 0),
		Map_ParagraphDBID_ParagraphDB:  make(map[uint]*ParagraphDB, 0),
		Map_ParagraphPtr_ParagraphDBID: make(map[*models.Paragraph]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoParagraphProperties = BackRepoParagraphPropertiesStruct{
		Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr: make(map[uint]*models.ParagraphProperties, 0),
		Map_ParagraphPropertiesDBID_ParagraphPropertiesDB:  make(map[uint]*ParagraphPropertiesDB, 0),
		Map_ParagraphPropertiesPtr_ParagraphPropertiesDBID: make(map[*models.ParagraphProperties]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoParagraphStyle = BackRepoParagraphStyleStruct{
		Map_ParagraphStyleDBID_ParagraphStylePtr: make(map[uint]*models.ParagraphStyle, 0),
		Map_ParagraphStyleDBID_ParagraphStyleDB:  make(map[uint]*ParagraphStyleDB, 0),
		Map_ParagraphStylePtr_ParagraphStyleDBID: make(map[*models.ParagraphStyle]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRune = BackRepoRuneStruct{
		Map_RuneDBID_RunePtr: make(map[uint]*models.Rune, 0),
		Map_RuneDBID_RuneDB:  make(map[uint]*RuneDB, 0),
		Map_RunePtr_RuneDBID: make(map[*models.Rune]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRuneProperties = BackRepoRunePropertiesStruct{
		Map_RunePropertiesDBID_RunePropertiesPtr: make(map[uint]*models.RuneProperties, 0),
		Map_RunePropertiesDBID_RunePropertiesDB:  make(map[uint]*RunePropertiesDB, 0),
		Map_RunePropertiesPtr_RunePropertiesDBID: make(map[*models.RuneProperties]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTable = BackRepoTableStruct{
		Map_TableDBID_TablePtr: make(map[uint]*models.Table, 0),
		Map_TableDBID_TableDB:  make(map[uint]*TableDB, 0),
		Map_TablePtr_TableDBID: make(map[*models.Table]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTableColumn = BackRepoTableColumnStruct{
		Map_TableColumnDBID_TableColumnPtr: make(map[uint]*models.TableColumn, 0),
		Map_TableColumnDBID_TableColumnDB:  make(map[uint]*TableColumnDB, 0),
		Map_TableColumnPtr_TableColumnDBID: make(map[*models.TableColumn]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTableProperties = BackRepoTablePropertiesStruct{
		Map_TablePropertiesDBID_TablePropertiesPtr: make(map[uint]*models.TableProperties, 0),
		Map_TablePropertiesDBID_TablePropertiesDB:  make(map[uint]*TablePropertiesDB, 0),
		Map_TablePropertiesPtr_TablePropertiesDBID: make(map[*models.TableProperties]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTableRow = BackRepoTableRowStruct{
		Map_TableRowDBID_TableRowPtr: make(map[uint]*models.TableRow, 0),
		Map_TableRowDBID_TableRowDB:  make(map[uint]*TableRowDB, 0),
		Map_TableRowPtr_TableRowDBID: make(map[*models.TableRow]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTableStyle = BackRepoTableStyleStruct{
		Map_TableStyleDBID_TableStylePtr: make(map[uint]*models.TableStyle, 0),
		Map_TableStyleDBID_TableStyleDB:  make(map[uint]*TableStyleDB, 0),
		Map_TableStylePtr_TableStyleDBID: make(map[*models.TableStyle]uint, 0),

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

	backRepo.broadcastNbCommitToBack()

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
	backRepo.BackRepoBody.CommitPhaseOne(stage)
	backRepo.BackRepoDocument.CommitPhaseOne(stage)
	backRepo.BackRepoDocx.CommitPhaseOne(stage)
	backRepo.BackRepoFile.CommitPhaseOne(stage)
	backRepo.BackRepoNode.CommitPhaseOne(stage)
	backRepo.BackRepoParagraph.CommitPhaseOne(stage)
	backRepo.BackRepoParagraphProperties.CommitPhaseOne(stage)
	backRepo.BackRepoParagraphStyle.CommitPhaseOne(stage)
	backRepo.BackRepoRune.CommitPhaseOne(stage)
	backRepo.BackRepoRuneProperties.CommitPhaseOne(stage)
	backRepo.BackRepoTable.CommitPhaseOne(stage)
	backRepo.BackRepoTableColumn.CommitPhaseOne(stage)
	backRepo.BackRepoTableProperties.CommitPhaseOne(stage)
	backRepo.BackRepoTableRow.CommitPhaseOne(stage)
	backRepo.BackRepoTableStyle.CommitPhaseOne(stage)
	backRepo.BackRepoText.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoBody.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDocument.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDocx.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFile.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNode.CommitPhaseTwo(backRepo)
	backRepo.BackRepoParagraph.CommitPhaseTwo(backRepo)
	backRepo.BackRepoParagraphProperties.CommitPhaseTwo(backRepo)
	backRepo.BackRepoParagraphStyle.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRune.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRuneProperties.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTable.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTableColumn.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTableProperties.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTableRow.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTableStyle.CommitPhaseTwo(backRepo)
	backRepo.BackRepoText.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoBody.CheckoutPhaseOne()
	backRepo.BackRepoDocument.CheckoutPhaseOne()
	backRepo.BackRepoDocx.CheckoutPhaseOne()
	backRepo.BackRepoFile.CheckoutPhaseOne()
	backRepo.BackRepoNode.CheckoutPhaseOne()
	backRepo.BackRepoParagraph.CheckoutPhaseOne()
	backRepo.BackRepoParagraphProperties.CheckoutPhaseOne()
	backRepo.BackRepoParagraphStyle.CheckoutPhaseOne()
	backRepo.BackRepoRune.CheckoutPhaseOne()
	backRepo.BackRepoRuneProperties.CheckoutPhaseOne()
	backRepo.BackRepoTable.CheckoutPhaseOne()
	backRepo.BackRepoTableColumn.CheckoutPhaseOne()
	backRepo.BackRepoTableProperties.CheckoutPhaseOne()
	backRepo.BackRepoTableRow.CheckoutPhaseOne()
	backRepo.BackRepoTableStyle.CheckoutPhaseOne()
	backRepo.BackRepoText.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoBody.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDocument.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDocx.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFile.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNode.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoParagraph.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoParagraphProperties.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoParagraphStyle.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRune.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRuneProperties.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTable.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTableColumn.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTableProperties.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTableRow.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTableStyle.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoText.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoBody.Backup(dirPath)
	backRepo.BackRepoDocument.Backup(dirPath)
	backRepo.BackRepoDocx.Backup(dirPath)
	backRepo.BackRepoFile.Backup(dirPath)
	backRepo.BackRepoNode.Backup(dirPath)
	backRepo.BackRepoParagraph.Backup(dirPath)
	backRepo.BackRepoParagraphProperties.Backup(dirPath)
	backRepo.BackRepoParagraphStyle.Backup(dirPath)
	backRepo.BackRepoRune.Backup(dirPath)
	backRepo.BackRepoRuneProperties.Backup(dirPath)
	backRepo.BackRepoTable.Backup(dirPath)
	backRepo.BackRepoTableColumn.Backup(dirPath)
	backRepo.BackRepoTableProperties.Backup(dirPath)
	backRepo.BackRepoTableRow.Backup(dirPath)
	backRepo.BackRepoTableStyle.Backup(dirPath)
	backRepo.BackRepoText.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoBody.BackupXL(file)
	backRepo.BackRepoDocument.BackupXL(file)
	backRepo.BackRepoDocx.BackupXL(file)
	backRepo.BackRepoFile.BackupXL(file)
	backRepo.BackRepoNode.BackupXL(file)
	backRepo.BackRepoParagraph.BackupXL(file)
	backRepo.BackRepoParagraphProperties.BackupXL(file)
	backRepo.BackRepoParagraphStyle.BackupXL(file)
	backRepo.BackRepoRune.BackupXL(file)
	backRepo.BackRepoRuneProperties.BackupXL(file)
	backRepo.BackRepoTable.BackupXL(file)
	backRepo.BackRepoTableColumn.BackupXL(file)
	backRepo.BackRepoTableProperties.BackupXL(file)
	backRepo.BackRepoTableRow.BackupXL(file)
	backRepo.BackRepoTableStyle.BackupXL(file)
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
	backRepo.BackRepoBody.RestorePhaseOne(dirPath)
	backRepo.BackRepoDocument.RestorePhaseOne(dirPath)
	backRepo.BackRepoDocx.RestorePhaseOne(dirPath)
	backRepo.BackRepoFile.RestorePhaseOne(dirPath)
	backRepo.BackRepoNode.RestorePhaseOne(dirPath)
	backRepo.BackRepoParagraph.RestorePhaseOne(dirPath)
	backRepo.BackRepoParagraphProperties.RestorePhaseOne(dirPath)
	backRepo.BackRepoParagraphStyle.RestorePhaseOne(dirPath)
	backRepo.BackRepoRune.RestorePhaseOne(dirPath)
	backRepo.BackRepoRuneProperties.RestorePhaseOne(dirPath)
	backRepo.BackRepoTable.RestorePhaseOne(dirPath)
	backRepo.BackRepoTableColumn.RestorePhaseOne(dirPath)
	backRepo.BackRepoTableProperties.RestorePhaseOne(dirPath)
	backRepo.BackRepoTableRow.RestorePhaseOne(dirPath)
	backRepo.BackRepoTableStyle.RestorePhaseOne(dirPath)
	backRepo.BackRepoText.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoBody.RestorePhaseTwo()
	backRepo.BackRepoDocument.RestorePhaseTwo()
	backRepo.BackRepoDocx.RestorePhaseTwo()
	backRepo.BackRepoFile.RestorePhaseTwo()
	backRepo.BackRepoNode.RestorePhaseTwo()
	backRepo.BackRepoParagraph.RestorePhaseTwo()
	backRepo.BackRepoParagraphProperties.RestorePhaseTwo()
	backRepo.BackRepoParagraphStyle.RestorePhaseTwo()
	backRepo.BackRepoRune.RestorePhaseTwo()
	backRepo.BackRepoRuneProperties.RestorePhaseTwo()
	backRepo.BackRepoTable.RestorePhaseTwo()
	backRepo.BackRepoTableColumn.RestorePhaseTwo()
	backRepo.BackRepoTableProperties.RestorePhaseTwo()
	backRepo.BackRepoTableRow.RestorePhaseTwo()
	backRepo.BackRepoTableStyle.RestorePhaseTwo()
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
	backRepo.BackRepoBody.RestoreXLPhaseOne(file)
	backRepo.BackRepoDocument.RestoreXLPhaseOne(file)
	backRepo.BackRepoDocx.RestoreXLPhaseOne(file)
	backRepo.BackRepoFile.RestoreXLPhaseOne(file)
	backRepo.BackRepoNode.RestoreXLPhaseOne(file)
	backRepo.BackRepoParagraph.RestoreXLPhaseOne(file)
	backRepo.BackRepoParagraphProperties.RestoreXLPhaseOne(file)
	backRepo.BackRepoParagraphStyle.RestoreXLPhaseOne(file)
	backRepo.BackRepoRune.RestoreXLPhaseOne(file)
	backRepo.BackRepoRuneProperties.RestoreXLPhaseOne(file)
	backRepo.BackRepoTable.RestoreXLPhaseOne(file)
	backRepo.BackRepoTableColumn.RestoreXLPhaseOne(file)
	backRepo.BackRepoTableProperties.RestoreXLPhaseOne(file)
	backRepo.BackRepoTableRow.RestoreXLPhaseOne(file)
	backRepo.BackRepoTableStyle.RestoreXLPhaseOne(file)
	backRepo.BackRepoText.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb(ctx context.Context) <-chan int {
	ch := make(chan int)

	backRepoStruct.rwMutex.Lock()
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	backRepoStruct.rwMutex.Unlock()

	// Goroutine to remove subscriber when context is done
	go func() {
		<-ctx.Done()
		backRepoStruct.unsubscribe(ch)
	}()
	return ch
}

// unsubscribe removes a subscriber's channel from the subscribers slice.
func (backRepoStruct *BackRepoStruct) unsubscribe(ch chan int) {
	backRepoStruct.rwMutex.Lock()
	defer backRepoStruct.rwMutex.Unlock()
	for i, subscriber := range backRepoStruct.subscribers {
		if subscriber == ch {
			backRepoStruct.subscribers =
				append(backRepoStruct.subscribers[:i],
					backRepoStruct.subscribers[i+1:]...)
			close(ch) // Close the channel to signal completion
			break
		}
	}
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.rwMutex.RLock()
	subscribers := make([]chan int, len(backRepoStruct.subscribers))
	copy(subscribers, backRepoStruct.subscribers)
	backRepoStruct.rwMutex.RUnlock()

	for _, ch := range subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			// Successfully sent commit from back
		default:
			// Subscriber is not ready to receive; skip to avoid blocking
		}
	}
}