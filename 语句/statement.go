package main;
func main(){
	if(10<20){
		println("10<20");
	}else{
		println("10>=20");
	}

	//switch
	var grade = "B";
	var marks= 90;
	switch marks{
	case 90:
		grade = "A"
	case 80 :
		grade = "B"
	case 50,60,70:
		grade= "C"
	}

	switch{
	case grade == "A": println("优秀");
	case grade == "B": println("良好");
	case grade == "C": println("一般");
	}

	//for 循环
	// 一定不要像JS那样加分号
	for i:=0;i<10;i++ {
		println("value is " ,i )
	}
	//函数的调用
	println(max(21,8))

}
func max(a int,b int) int{
	if(a>b){
		return a;
	}else{
		return b;
	}
	// return a>b ? a : b;
}