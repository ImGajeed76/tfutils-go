package Systembeschreibung

import (
	"fmt"
	"github.com/ImGajeed76/charmer/pkg/charmer/console"
	"github.com/ImGajeed76/charmer/pkg/charmer/path"
	pathmodels "github.com/ImGajeed76/charmer/pkg/charmer/path/models"
	"log"
	"sort"
	"tfutils-go/internal/config"
	"tfutils-go/internal/helpers"
)

/*
CreateNewSystemDescription godoc
@Charm
@Title Neue Systembeschreibung erstellen
@Description

# Neue Systembeschreibung erstellen

Mit diesem Charm können Sie eine neue Systembeschreibung erstellen.
*/
func CreateNewSystemDescription() {
	officePath := path.New("/t_lernende/E/LIVE/02_Vorlagen/01_Office", config.GetSFTPConfig())
	files, err := officePath.Glob("Systembeschreibung_Vorlage_v*.dotx")
	if err != nil {
		log.Fatal(err)
		return
	}

	if len(files) == 0 {
		log.Println("No system description template found")
		return
	}

	// sort files by name, higher version numbers first
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() > files[j].Name()
	})

	latestFile := files[0]
	currentDirFile := path.Cwd().Join("Systembeschreibung.dotx")

	progressBar := console.NewProgressBar()
	defer progressBar.Close()

	err = latestFile.CopyTo(currentDirFile, pathmodels.CopyOptions{
		PathOption:   pathmodels.DefaultPathOption(),
		ProgressFunc: progressBar.Update,
	})
	if err != nil {
		fmt.Printf("Error while copying file: %v\n", err)
		return
	}

	progressBar.Finish()

	fmt.Printf("File copied to %s\n", currentDirFile)

	// Convert file to docx
	newFile := currentDirFile.Parent().Join("Systembeschreibung.docx")
	err = helpers.ConvertDOTXtoDOCX(currentDirFile, newFile)
	if err != nil {
		fmt.Printf("Error while converting file: %v\n", err)
		return
	}

	err = currentDirFile.Remove(true, false)
	if err != nil {
		fmt.Printf("Error while removing file: %v\n", err)
		return
	}

	fmt.Printf("File converted to %s\n", newFile)
}
