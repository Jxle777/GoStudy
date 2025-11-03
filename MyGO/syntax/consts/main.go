package consts

const external = "包外"
const internal = "包外"
const (
	a = 123
)

const (
	StatusA = iota //自动帮你+1
	StatusB
	StatusC
	StatusD
	StatusE = 100
	StatusF //可被打断
)

const (
	DayA = iota*12 + 13 //相当于0*12+13
)

func main() {
	const a = 123
	//a = 456 常量无法更改值（恒量）
}
