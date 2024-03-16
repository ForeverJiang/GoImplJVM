package main

import "fmt"
import "strings"
import "jvmgo/ch02/classpath"

/*
本章讨论了Java虚拟机从哪里寻找class文件，对类路径和-classpath命令行选项有了较为深入的了解，并且把抽象的类路径概
念转变成了具体的代码。下一章将研究class文件格式，实现class文
件解析。

go install D:\go\workspace\src\jvmgo\ch02
D:\go\workspace\bin\ch02.exe -Xjre "C:\Program Files\Java\jdk1.8.0_102\jre" java.lang.Object
*/
func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

/*
startJVM（）先打印出命令行参数，然后读取主类数据，并打印
到控制台。虽然还是无法真正启动Java虚拟机，不过相比第1章，已
经有了很大的进步。打开命令行窗口，执行下面的命令编译本章代
码。
*/
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
