package classpath

import "errors"
import "strings"

/*
如前所述，CompositeEntry由更小的Entry组成，正好可以表示
成[]Entry。在Go语言中，数组属于比较低层的数据结构，很少直接
使用。大部分情况下，使用更便利的slice类型。构造函数把参数（路
径列表）按分隔符分成小路径，然后把每个小路径都转换成具体的
Entry实例
*/
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

/*
相信读者已经想到readClass（）方法的代码了：依次调用每一个
子路径的readClass（）方法，如果成功读取到class数据，返回数据即
可；如果收到错误信息，则继续；如果遍历完所有的子路径还没有找
到class文件，则返回错误。
*/
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

/*
String（）方法也不复杂。调用每一个子路径的String（）方法，然
后把得到的字符串用路径分隔符拼接起来即可
*/
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
