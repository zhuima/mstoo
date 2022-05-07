package internels

import "os"

func ReadFile(filename string) (*os.File, error) {
	if !FileExists(filename) {
		return nil, os.ErrNotExist
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// defer file.Close()
	return file, nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
