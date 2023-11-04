```go
package image

type TarImage struct {
	files           []string
	excludePaths    []string
	logger          string
	keepPermissions bool
}

func NewTarImage(files []string, excludePaths []string, logger string, keepPermissions bool) *TarImage {
	return &TarImage{files, excludePaths, logger, keepPermissions}
}
```
The function returns a pointer to a TarImage struct that is initialized with the provided parameters. It essentially constructs and configures a TarImage instance based on the provided input.
