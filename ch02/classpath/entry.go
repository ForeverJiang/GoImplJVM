package classpath

import (
	"os"
	"strings"
)

/**
常量pathListSeparator是string类型，存放路径分隔符，后面会用
到。Entry接口中有两个方法。readClass（）方法负责寻找和加载class
文件；String（）方法的作用相当于Java中的toString（），用于返回变量
的字符串表示。
readClass（）方法的参数是class文件的相对路径，路径之间用斜
线（/）分隔，文件名有.class后缀。比如要读取java.lang.Object类，传
入的参数应该是java/lang/Object.class。返回值是读取到的字节数
据、最终定位到class文件的Entry，以及错误信息。Go的函数或方法
允许返回多个值，按照惯例，可以使用最后一个返回值作为错误信
息。
newEntry（）函数根据参数创建不同类型的Entry实例，代码如下
*/

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
