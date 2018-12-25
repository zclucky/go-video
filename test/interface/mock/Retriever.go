package mock

// 定义类
type Retriever struct {
	Comment string
}

// 实现接口
func (r *Retriever) Get(url string ) string  {
	return r.Comment
}

// 实现接口
func (r *Retriever)String() string {
	return "{||" + r.Comment + "||}"
}
