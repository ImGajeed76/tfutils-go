package Checkliste

import (
	"fmt"
	"github.com/ImGajeed76/charmer/pkg/charmer/console"
	"github.com/ImGajeed76/charmer/pkg/charmer/path"
	pathmodels "github.com/ImGajeed76/charmer/pkg/charmer/path/models"
	"sort"
	"tfutils-go/internal/config"
)

/*
CreateSchemaChecklist godoc
@Charm
@Title Schema-Checkliste erstellen
@Description

# Schema-Checkliste erstellen

Die Schema-Checkliste ist ein Dokument, das die Struktur und die Anforderungen an ein Schema beschreibt. Sie dient
dazu, die Anforderungen an ein Schema zu dokumentieren und sicherzustellen, dass alle Anforderungen erfüllt sind.

Mit diesem Charm können Sie eine Schema-Checkliste in dem oben angegebenen Pfad erstellen.
*/
func CreateSchemaChecklist() {
	copyChecklist("Checkliste_SCH_v*.docx")
}

/*
CreatePCBChecklist godoc
@Charm
@Title PCB-Checkliste erstellen
@Description

# PCB-Checkliste erstellen

Die PCB-Checkliste ist ein Dokument, das die Struktur und die Anforderungen an ein PCB beschreibt. Sie dient
dazu, die Anforderungen an ein PCB zu dokumentieren und sicherzustellen, dass alle Anforderungen erfüllt sind.

Mit diesem Charm können Sie eine PCB-Checkliste in dem oben angegebenen Pfad erstellen.
*/
func CreatePCBChecklist() {
	copyChecklist("Checkliste_PCB_v*.docx")
}

func copyChecklist(pattern string) {
	checklistTemplatesPath := path.New("/t_lernende/E/LIVE/05_HW_Entwicklung/08_Checklisten", config.GetSFTPConfig())

	checklistFiles, globErr := checklistTemplatesPath.Glob(pattern)
	if globErr != nil {
		fmt.Printf("Error while globbing checklist files: %v\n", globErr)
		return
	}

	if len(checklistFiles) == 0 {
		fmt.Println("No checklist files found")
		return
	}

	// sort checklist files by name, higher version numbers first
	sort.Slice(checklistFiles, func(i, j int) bool {
		return checklistFiles[i].Name() > checklistFiles[j].Name()
	})

	checklistFile := checklistFiles[0]
	currentDirFile := path.Cwd().Join(checklistFile.Name())

	progressBar := console.NewProgressBar()
	defer progressBar.Close()

	err := checklistFile.CopyTo(currentDirFile, pathmodels.CopyOptions{
		PathOption:   pathmodels.DefaultPathOption(),
		ProgressFunc: progressBar.Update,
	})
	if err != nil {
		fmt.Printf("Error while copying checklist file: %v\n", err)
		return
	}

	progressBar.Finish()

	fmt.Printf("Checklist file copied to %s\n", currentDirFile)
}
