package classpath
import "io/ioutil"
import "path/filepath"
type ZipEntry struct{
	absPath string
}
func newZipEntry(path string) *ZipEntry {
	absPath,err:=filepath.Abs(path)
	if err!=nil{
		panic(err)
	}
	return &ZipEntry{absPath}
}
func (self *ZipEntry) readClass(className string) ([]byte,Entry,error) {
	r,err:=zip.OpenReader(self.absPath)
	if err !=nil{
		return nil,nil,err
	}
	defer r.Close()
	for _,f:=range r.File{
		if f.Name==
	}
}
func (self *ZipEntry) string() string {
	return self.absPath
}