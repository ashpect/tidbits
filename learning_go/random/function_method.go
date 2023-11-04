package main

type Contents struct {
	paths               []string
	excludedPaths       []string
	preservePermissions bool
}

func NewContents(paths []string, excludedPaths []string, preservePermissions bool) Contents {
	return Contents{paths: paths, excludedPaths: excludedPaths, preservePermissions: preservePermissions}
}

func (b Contents) Push(someargument string) (string, error) {
	return "", nil
}

func main() {
	//Usage example:
	//imageURL, err := bundle.NewContents(po.FileFlags.Files, po.FileFlags.ExcludedFilePaths, po.FileFlags.PreservePermissions).Push(uploadRef, po.LabelFlags.Labels, registry, logger)

	imageURL, err := NewContents([]string{"file1.txt", "file2.txt", "file3.txt"}, []string{"file2.txt"}, true).Push("someargument")

	//OR AS

	hi := NewContents([]string{"file1.txt", "file2.txt", "file3.txt"}, []string{"file2.txt"}, true)
	imageURL, err = hi.Push("someargument")

	if err != nil {
		panic(err)
	}
	println(imageURL)

}
