package main;

import (
	"fmt"
)
var a,b= 11,21;
var F,T=false,true;
func main(){
	println(a==b)//false
	println(a!=b)//true
	println(a>b)//false
	println(a>=b)//false
	println(a<b)//true
	println(a<=b)//true
	println("\n")

	println(T && F);//false
	println(!T);//flase
	println(!F);//true
	println(T || F);//true

	//赋值
	var STAR=10;
	var END=20;
	var TEMP=STAR+END;
	println(TEMP);//30
	STAR+=END;
	println(STAR);//30

	STAR-=END;
	println(STAR);//10

	STAR*=END
	println(STAR)//200

	STAR/=END
	println(STAR);//10


	// & 和 * 运算符实例
	var x int =4;
	var y int32;
	var z float32;
	var ptr *int;
	println(x)	//4
	println(y)	//0
	println(z)	//+0.000000e+000
	println(ptr)//0x0

	fmt.Printf("fmt.Printf 需要字符串做参数",x)	//%!(EXTRA int=4)
	fmt.Printf("\n",y)	//%!(EXTRA int32=0)
	fmt.Printf("\n",z)	//%!(EXTRA float32=0)
	fmt.Printf("\n",ptr)//%!(EXTRA *int=<nil>)

	fmt.Printf("\n %T",x);//int
	fmt.Printf("\n %T",y);//int32
	fmt.Printf("\n %T",z);//faloat32
	fmt.Printf("\n %T",ptr);//*int

	var test1=111;
	println(&test1);//*int0xc42003bec8
	// var test2=func(){};
	// println("22",*test2)

	//

}