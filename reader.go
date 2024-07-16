package filestorage

import (
	"fmt"
	"os"
	"path"
)

// readFiles reads all files in the given directory and its subdirectories
// and writes them in a filePaths map with subdir/filename as keys and filepath as values
// and returns all keys in filePaths map
func readFiles(filePaths map[string]string, basePath string) ([]string, error) {
	const operation = "read files from base dir"

	items, err := os.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", operation, err)
	}

	files := make([]string, 0, len(items))

	for _, item := range items {
		if item.IsDir() {
			dirname := item.Name()
			dirfiles, err := readSubdir(filePaths, path.Join(basePath, dirname), dirname)
			if err != nil {
				return nil, err
			}
			files = append(files, dirfiles...)
		} else {
			filename := item.Name()

			files = append(files, filename)
			filePaths[filename] = path.Join(basePath, filename)
		}
	}

	return files, nil
}

func readSubdir(filePaths map[string]string, basePath, subDir string) ([]string, error) {
	operation := fmt.Sprintf("read files from subdir %s", subDir)

	items, err := os.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", operation, err)
	}

	files := make([]string, 0, len(items))

	for _, item := range items {
		if item.IsDir() {
			dirname := item.Name()
			dirfiles, err := readSubdir(
				filePaths,
				path.Join(basePath, dirname),
				fmt.Sprintf("%s/%s", subDir, dirname),
			)
			if err != nil {
				return nil, err
			}
			files = append(files, dirfiles...)
		} else {
			filename := item.Name()
			key := fmt.Sprintf("%s/%s", subDir, filename)

			files = append(files, key)
			filePaths[key] = path.Join(basePath, filename)
		}
	}

	return files, nil
}
