package paths

import (
	"fmt"
	"os"
)

var name = "Paths"

// type Paths struct {
// 	html        FileSystemPath
// 	images      FileSystemPath
// 	backups     FileSystemPath
// 	dumps       FileSystemPath
// 	database    FileSystemPath
// 	data        FileSystemPath
// 	application FileSystemPath
// }

type FileSystemPath struct {
	path string
}

func (f FileSystemPath) String() string {
	return f.path
}

func HTML() FileSystemPath {
	return FileSystemPath{Res().String() + "/html/templates/"}
}

func HTMLTemplates() FileSystemPath {
	return FileSystemPath{Res().String() + "/html/"}
}

func HTMLPage(in string) string {
	return HTML().String() + in + ".html"
}

func HTMLTemplate() string {
	return HTMLTemplates().String() + "templates.html"
}

func Images() FileSystemPath {
	return FileSystemPath{Res().String() + "/img"}
}

func Backups() FileSystemPath {
	return FileSystemPath{Data().String() + "/backups"}
}

func Dumps() FileSystemPath {
	return FileSystemPath{Data().String() + "/dumps"}
}

func Database() FileSystemPath {
	return FileSystemPath{Data().String() + "/database"}
}

func Config() FileSystemPath {
	return FileSystemPath{Data().String() + "/config"}
}

func Logs() FileSystemPath {
	return FileSystemPath{Data().String() + "/logs"}
}

func Data() FileSystemPath {
	return FileSystemPath{"/data"}
}

func Res() FileSystemPath {
	return FileSystemPath{"./res"}
}

func Application() FileSystemPath {
	return FileSystemPath{fullPath()}
}

func Seperator() string {
	return string(os.PathSeparator)
}

func (F *FileSystemPath) Is(in FileSystemPath) bool {
	return F.path == in.path
}

func fullPath() string {
	// Get the full path of the current directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("[%v] Getting current directory", err)
		panic(err)
	}
	return dir
}
