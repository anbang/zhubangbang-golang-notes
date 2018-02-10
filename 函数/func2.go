package main;
// import ( "fmt")
func main (){

	test1 := getSequence();

	println(test1());//1
	println(test1());//2
	println(test1());//3

	test1 = getSequence();
	println(test1());//1
	println(test1());//2
	println(test1());//3
}

//返回参数有2个，func int
func getSequence() func() int {
	i:=0;
	return func () int {
		i++;
		return i;
	}
}