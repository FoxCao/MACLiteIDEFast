// fast
package main

import (
	"fmt"
	"net/http"
)

func LiteIDEFast(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type":"text/html;charset=utf-8")
	fmt.Fprintf(w, "<div><p>option + 1 ：打开时间记录</p><p>option + 2 : 打开搜索+结果（大范围）</p><p>option + 4 ：打开编译输出</p><p>option + 5 ：打开调试输出</p><p>option + 6 ：打开Go源码查询</p><p>option + 7 ：打开GoLang文档索引</p><p>option + 8 ：打开GoLangApi索引</p><p>option + command + 1 ：目录</p><p>option + command + 2 ：文档</p><p>option + command + 3 ：打开类视图</p><p>option + command + 4 ：打开大纲</p><p>option + command + 5 ：package浏览</p><p>option + command + 6 ：打开文件系统</p><p>option + command + 7 ：打开书签</p><p>command + N ：新建</p><p>command + O ：打开文件</p><p>command + P ： 打开文件</p><p>option + command + P ：快速打开tab页</p><p>shift + command + N ：打开新窗口</p><p>shift + command + W : 关闭窗口</p><p>command + W : 关闭文件</p><p>shift + command + D ：向下复制整行</p><p>shift + command + K ：删除整行</p><p>shift + command + 上箭头 ：行向上移动</p><p>shift + command + 下箭头 ：行向下移动</p><p>�option + command + 上箭头 ：复制行向上</p><p>option+ command + 下箭头 ：复制行向下</p><p>command + L ：跳转到多少行</p><p>command + [ :跳转到上一段</p><p>command + ] :跳转到下一段</p><p>command + J ：连接行，将下面的行和光标所在行合并成一行。</p><p>shift + command + 回车：行前插入</p><p>command + 回车：行后插入</p><p>command + U：选择段</p><p>command + I ：格式化代码</p><p>option + command + I ：格式化代码（调整导入行）</p><p>shift + command + I ： 查看表达式的信息</p><p>shift + command + J ： 跳转到方法的声明</p><p>shift + command + U ： 查找使用（当前文件）</p><p>option+ command + U ： 查找使用、打开搜索页</p><p>command + F ：查找</p><p>command + G ：查找下一个</p><p>shift + command + G ：查找上一个</p><p>option + command + F ：替换</p><p>shift + command + F ：文件检索</p><p>shift + command + H ： Go 源码查询，option + 6</p><p>command + + ：增大字号</p><p>command + — ：减小字号</p><p>command + O ：重置字号</p><p>command + R ：编译并且运行</p><p>option + command + R : 运行</p><p>shift + command + R ：filerun</p><p>command + B ：编译</p><p>option + command + B ：强制编译</p><p>command + F8 : 安装。</p></div>")
}
func main() {
	// fmt.Println("Hello World!")
	server := http.Server{Addr: "localhost:8899"}
	http.HandleFunc("/", LiteIDEFast)
	// http.ListenAndServe(":8899", nil)
	server.ListenAndServe()
}
