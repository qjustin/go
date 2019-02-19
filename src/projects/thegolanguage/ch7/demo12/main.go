package main

import "io"

/*
	7.12. 通过类型断言询问行为

	我们不能对任意io.Writer类型的变量w，假设它也拥有WriteString方法。
	但是我们可以定义一个只有这个方法的新接口并且使用类型断言来检测是否w的动态类型满足这个新接口。
*/

func main() {
}

func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) // avoid a copy
	}
	return w.Write([]byte(s)) // allocate temporary copy
}

func writeHeader(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-Type: "); err != nil {
		return err
	}
	if _, err := writeString(w, contentType); err != nil {
		return err
	}

	return nil
}
