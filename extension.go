package filestorage

import (
	"fmt"

	"github.com/dop251/goja"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/filestorage", new(FileStorageExt))
}

type FileStorageExt struct{}

// XFileStorage creates a new FileStorage instance with the given base path.
// Usage in k6 scripts: `const storage = new FileStorage("./testdata");`
func (b *FileStorageExt) XFileStorage(call goja.ConstructorCall, rt *goja.Runtime) *goja.Object {
	var basePath string

	err := rt.ExportTo(call.Argument(0), &basePath)
	if err != nil {
		panic(fmt.Errorf("error reading argument: %w", err))
	}

	return rt.ToValue(NewFileStorage(basePath)).ToObject(rt)
}
