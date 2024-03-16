package classpath

/*
ZipEntry表示ZIP或JAR文件形式的类路径。
*/

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

// absPath字段存放ZIP或JAR文件的绝对路径。构造函数和String（）与DirEntry大同小异
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

/*
首先打开ZIP文件，如果这一步出错的话，直接返回。
然后遍历ZIP压缩包里的文件，看能否找到class文件。如果能找到，则打开
class文件，把内容读取出来，并返回。如果找不到，或者出现其他错
误，则返回错误信息。有两处使用了defer语句来确保打开的文件得
以关闭。readClass（）方法每次都要打开和关闭ZIP文件，因此效率不
是很高。笔者进行了优化，但鉴于篇幅有限，就不展示具体代码
了。感兴趣的读者可以阅读ch02\classpath\entry_zip2.go文件。
*/
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
