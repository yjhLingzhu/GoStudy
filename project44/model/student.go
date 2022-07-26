package model

type student struct {
	Name string
	Age  int
}

// GetStudent 因为student结果体首字母是小写，因此只能在model使用
// 我们通过工厂模式来解决
func GetStudent(Name string, Age int) *student {
	return &student{
		Name: Name,
		Age:  Age,
	}
}
