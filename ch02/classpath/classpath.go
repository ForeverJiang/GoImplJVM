package classpath

import "os"
import "path/filepath"

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

/*
Classpath结构体有三个字段，分别存放三种类路径。
Parse（）函数使用-Xjre选项解析启动类路径和扩展类路径，
使用-classpath/-cp选项解析用户类路径
*/
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// ReadClass（）方法依次从启动类路径、扩展类路径和用户类路径中搜索class文件
// 注意，传递给ReadClass（）方法的类名不包含“.class”后缀。
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}

/*
优先使用用户输入的-Xjre选项作为jre目录。如果没有输入该
选项，则在当前目录下寻找jre目录。如果找不到，尝试使用
JAVA_HOME环境变量。
*/
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// 如果用户没有提供-classpath/-cp选项，则使用当前目录作为用户类路径
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

// exists（）函数用于判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
