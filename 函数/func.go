package main;
import ("strconv")
var c,d = 0 ,0;
var e *int;
var f *int;
func main(){
	println(maxVal(2,4));

	x,y:= maxVal2(2,5);
	println(x,y);

	a,b:=3,7
	c,d = change(a,b);//返回多值
	println(a,b);//值的，不会改变原有数据
	println(c,d);

	age1,age2:= 17,28;
	println(age1,age2);//值的
	e,f = change2( &age1 , &age2 )
	println(age1,age2);//已经倍改变了

	beautifyVal:=func(x int) string {
		var result= "￥" + strconv.Itoa(x);
		return result;
	}
	println(beautifyVal(234));//￥ 234
}
//返回单个
func maxVal(a int,b int) int {
	// return val:= a>b ? a : b;
	if(a>b){
		return a;
	}else{
		return b;
	}
}

//返回多个.返回多值，调用时候需要多个变量接受
func maxVal2(a int,b int) (int,int) {
	// return val:= a>b ? a : b;
	if(a>b){
		return a, b;
	}else{
		return b , a;
	}
}

//交换 
func change(a int,b int) (int,int) {
	a,b=b,a;
	return a,b;
}

func change2(a *int,b *int) (*int,*int) {
	*a,*b=*b,*a;
	return a,b;
}