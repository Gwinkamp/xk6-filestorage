package filestorage

import (
	"fmt"
	"math/rand/v2"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type (
	FileStorage struct {
		basePath string
		files    []string
		extCache map[string][]string
	}

	File struct {
		Name    string
		Path    string
		Content []byte
	}
)

func NewFileStorage(basePath string) *FileStorage {
	items, err := os.ReadDir(basePath)
	if err != nil {
		panic(err)
	}

	files := make([]string, 0, len(items))

	for _, item := range items {
		if item.IsDir() {
			continue
		}
		files = append(files, item.Name())
	}

	if len(files) == 0 {
		panic("no files found in " + basePath)
	}

	baseAbsPath, err := filepath.Abs(basePath)
	if err != nil {
		panic(err)
	}

	return &FileStorage{
		basePath: baseAbsPath,
		files:    files,
		extCache: make(map[string][]string),
	}
}

func (fs *FileStorage) ListFiles() []string {
	return fs.files
}

func (fs *FileStorage) HasFile(filename string) bool {
	for _, file := range fs.files {
		if file == filename {
			return true
		}
	}
	return false
}

func (fs *FileStorage) ReadFile(filename string) (File, error) {
	filePath := path.Join(fs.basePath, filename)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return File{}, err
	}
	return File{
		Name:    filename,
		Path:    filePath,
		Content: data,
	}, nil
}

// ReadRandFile reads a random file from the storage.
func (fs *FileStorage) ReadRandFile() (File, error) {
	randIndex := rand.IntN(len(fs.files))
	return fs.ReadFile(fs.files[randIndex])
}

// ReadRandFileWithExt reads a random file with the given extension from the storage.
func (fs *FileStorage) ReadRandFileWithExt(ext string) (File, error) {
	ext = strings.ToLower(ext)

	var files []string

	if files, ok := fs.extCache[ext]; ok {
		randIndex := rand.IntN(len(files))
		return fs.ReadFile(files[randIndex])
	}

	files = make([]string, 0, len(fs.files))
	for _, file := range fs.files {
		if strings.ToLower(path.Ext(file)) == ext {
			files = append(files, file)
		}
	}

	if len(files) == 0 {
		return File{}, fmt.Errorf("no files found with extension '%s'", ext)
	}

	fs.extCache[ext] = files

	randIndex := rand.IntN(len(files))
	return fs.ReadFile(files[randIndex])
}
