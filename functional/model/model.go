package model

// TODO カプセル化
type Employee struct {
	Name    string
	Age     int
	Version int
}

// TODO バリデーションとエラーハンドリング
func SetAge(e Employee) func(int) Employee {
	return func(age int) Employee {
		e.Age = age
		return e
	}
}
