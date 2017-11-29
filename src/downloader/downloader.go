package downloader

import (
	"os"
	"net/http"
	"io"
	"path/filepath"
)

// Download file from given URL
// and save it under specified path/filename
func Start(targetFilepath string, url string) error {

	// Create target directory
	err := createTargetDirectory(targetFilepath)
	if err != nil  {
		return err
	}

	// Get file from URL
	err = readAndSaveFileFromUrl(targetFilepath, url)
	if err != nil  {
		return err
	}

	return nil
}

// Creates directory based on provided filepath
func createTargetDirectory(targetFilepath string) error {

	// extract directory name from full file path
	dir, _ := filepath.Split(targetFilepath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Make directory bc it not exists
		err := os.MkdirAll(dir, 0755)
		if err != nil  {
			return err
		}
	}

	return nil
}

// Reads content of remote file and stores it locally
// under given path and file name
func readAndSaveFileFromUrl(targetFilepath string, url string) error {

	// Create the file
	output, err := os.Create(targetFilepath)
	if err != nil  {
		return err
	}
	defer output.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(output, resp.Body)
	if err != nil  {
		return err
	}

	return nil
}