package main;
// import "fmt";
import "unsafe";

var name = "zhuzhu"
var age int = 28;

const WOMANVAL,MANVAL,UNKNOW=0,1,2;

func main(){
	say := "hello"
	var name1, name2 ,name3 = "zhu","bang","bang";
	println(name1,name2,name3)

	name4,name5 :="zhuzhu4",234;
	println(name4,name5);
	println(WOMANVAL,MANVAL,UNKNOW)
	println(len(name1),unsafe.Sizeof(UNKNOW));
	println(name,age,say)
}