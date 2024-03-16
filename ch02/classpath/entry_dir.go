package classpath

import "io/ioutil"
import "path/filepath"

// DirEntry只有一个字段，用于存放目录的绝对路径。
type DirEntry struct {
	absDir string
}

/*
Go结构体不需要显示实现接口，只要方法匹配即可。Go没有
专门的构造函数，本书统一使用new开头的函数来创建结构体实
例，并把这类函数称为构造函数。newDirEntry（）函数。
newDirEntry（）先把参数转换成绝对路径，如果转换过程出现错误，
则调用panic（）函数终止程序执行，否则创建DirEntry实例并返回。
*/
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

/*
readClass（）先把目录和class文件名拼成一个完整的路径，然后
调用ioutil包提供的ReadFile（）函数读取class文件内容，最后返回。
*/
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
