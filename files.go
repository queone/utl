package utl

import (
	"os"
	"time"
)

// Read and recode given filePath as text byte slice.
// Returns the byte slice and error if any.
func LoadFileText(filePath string) (rawBytes []byte, err error) {
	rawBytes, err = os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return rawBytes, nil
}

// Saves given byte slice as text file.
// Returns error is any.
func SaveFileText(filePath string, rawBytes []byte) error {
	err := os.WriteFile(filePath, rawBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Removes given filepath
func RemoveFile(filePath string) {
	if FileExist(filePath) {
		if err := os.Remove(filePath); err != nil {
			panic(err.Error())
		}
	}
}

// Returns true if filepath exists and has some content. False otherwise.
func FileUsable(filePath string) (e bool) {
	if FileExist(filePath) && FileSize(filePath) > 0 {
		return true
	}
	return false
}

// Returns true if given filePath exists. False otherwise.
func FileExist(filePath string) (e bool) {
	if _, err := os.Stat(filePath); err == nil || os.IsExist(err) {
		return true
	}
	return false
}

// Returns true is given filePath does not exist. False otherwise.
func FileNotExist(filePath string) (e bool) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return true
	}
	return false
}

// Returns size of given filePath as int64
func FileSize(filePath string) int64 {
	f, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return f.Size()
}

// Returns given filePath modified time in Unix epoch int
func FileModTime(filePath string) int {
	f, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return int(f.ModTime().Unix())
}

// Returns given filePath age in seconds int64
func FileAge(filePath string) int64 {
	if FileUsable(filePath) {
		fileEpoc := int64(FileModTime(filePath))
		return int64(time.Now().Unix()) - fileEpoc
	}
	return int64(0)
}
