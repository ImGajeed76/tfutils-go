package helpers

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/ImGajeed76/charmer/pkg/charmer/path"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ConvertDOTXtoDOCX(dotxPath, docxPath *path.Path) error {
	// Check if both paths are local
	if dotxPath.IsSftp() || dotxPath.IsUrl() || docxPath.IsSftp() || docxPath.IsUrl() {
		return fmt.Errorf("both paths must be local")
	}

	dotxStr := dotxPath.String()
	docxStr := docxPath.String()

	// Open the DOTX file
	dotxReader, err := zip.OpenReader(dotxStr)
	if err != nil {
		return fmt.Errorf("failed to open DOTX file: %w", err)
	}
	defer dotxReader.Close()

	// Create a buffer for the DOCX file
	docxBuffer := new(bytes.Buffer)
	docxWriter := zip.NewWriter(docxBuffer)

	// Copy all files from DOTX to DOCX with necessary modifications
	for _, file := range dotxReader.File {
		// Open the file from the DOTX archive
		rc, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open file in DOTX: %w", err)
		}

		// Read the file content
		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return fmt.Errorf("failed to read file content: %w", err)
		}

		// Modify content-type if necessary
		if filepath.Base(file.Name) == "[Content_Types].xml" {
			content = bytes.Replace(
				content,
				[]byte("application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml"),
				[]byte("application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"),
				-1,
			)
		}

		// Create a new file in the DOCX archive
		newFilename := file.Name
		// In case of document.xml, change its location from templates to documents
		if strings.Contains(newFilename, "word/document.xml") {
			newFilename = strings.Replace(newFilename, "template", "document", -1)
		}

		w, err := docxWriter.Create(newFilename)
		if err != nil {
			return fmt.Errorf("failed to create file in DOCX: %w", err)
		}

		// Write the content
		_, err = w.Write(content)
		if err != nil {
			return fmt.Errorf("failed to write file content: %w", err)
		}
	}

	// Close the DOCX writer
	err = docxWriter.Close()
	if err != nil {
		return fmt.Errorf("failed to close DOCX writer: %w", err)
	}

	// Write the DOCX buffer to file
	err = os.WriteFile(docxStr, docxBuffer.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write DOCX file: %w", err)
	}

	return nil
}
