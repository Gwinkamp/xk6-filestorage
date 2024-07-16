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
		basePath  string
		files     []string
		filePaths map[string]string
		extCache  map[string][]string
	}

	File struct {
		Name    string
		Path    string
		Content []byte
	}
)

func NewFileStorage(basePath string) *FileStorage {
	baseAbsPath, err := filepath.Abs(basePath)
	if err != nil {
		panic(err)
	}

	filePaths := make(map[string]string)
	files, err := readFiles(filePaths, baseAbsPath)
	if err != nil {
		panic(err)
	}

	if len(files) == 0 {
		panic("no files found in " + basePath)
	}

	return &FileStorage{
		basePath:  baseAbsPath,
		files:     files,
		filePaths: filePaths,
		extCache:  make(map[string][]string),
	}
}

func (fs *FileStorage) ListFiles() []string {
	return fs.files
}

func (fs *FileStorage) HasFile(key string) bool {
	if _, ok := fs.filePaths[key]; ok {
		return true
	}
	return false
}

func (fs *FileStorage) ReadFile(key string) (File, error) {
	if filepath, ok := fs.filePaths[key]; ok {
		data, err := os.ReadFile(filepath)
		if err != nil {
			return File{}, err
		}
		return File{
			Name:    parseFilenameFromKey(key),
			Path:    filepath,
			Content: data,
		}, nil
	}
	return File{}, fmt.Errorf("file '%s' not found", key)
}

// ReadRandFile reads a random file from the storage
func (fs *FileStorage) ReadRandFile() (File, error) {
	return fs.readRandFile(fs.files)
}

// ReadRandFileWithExt reads a random file with the given extension from the storage
func (fs *FileStorage) ReadRandFileWithExt(ext string) (File, error) {
	ext = strings.ToLower(ext)

	var files []string

	if files, ok := fs.extCache[ext]; ok {
		return fs.readRandFile(files)
	}

	files = make([]string, 0, len(fs.files))
	for _, filename := range fs.files {
		if strings.ToLower(path.Ext(filename)) == ext {
			files = append(files, filename)
		}
	}

	if len(files) == 0 {
		return File{}, fmt.Errorf("no files found with extension '%s'", ext)
	}

	fs.extCache[ext] = files

	return fs.readRandFile(files)
}

// readRandFile reads a random file from the given slice of filenames
func (fs *FileStorage) readRandFile(files []string) (File, error) {
	randIndex := rand.IntN(len(files))
	key := files[randIndex]

	filename := parseFilenameFromKey(key)
	filepath := fs.filePaths[key]

	data, err := os.ReadFile(filepath)
	if err != nil {
		return File{}, err
	}
	return File{
		Name:    filename,
		Path:    filepath,
		Content: data,
	}, nil
}

// parseFilenameFromKey extracts filename from given key (subdir/filename)
// and returns it as a string
func parseFilenameFromKey(key string) string {
	lastIndex := strings.LastIndex(key, "/")
	return key[lastIndex+1:]
}
