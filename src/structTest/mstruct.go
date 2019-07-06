package structTest

type Shape struct {
	Length int // 导出性也是通过大小写来区分的
	Name string
}

func (s *Shape) OutputLength() int{
	return s.Length;
}
