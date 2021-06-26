package structer

type Test struct {
	//首字母大写，可被外部包使用
	Ok int
	//首字母小写，不可被外部包使用
	noWay int
}

type Test2 struct {
	Ok    int
	NoWay int
}
