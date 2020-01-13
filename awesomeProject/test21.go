package main

import (
	"fmt"
	"io"
)

type DataWriter interface {
	WriteData(data interface{})error
	CanWrite()bool
}
type file struct {

}

func (f *file) CanWrite() bool {
	panic("implement me")
}

func (f *file) WriteData(data interface{}) error{
	fmt.Println("WriteData:",data)
	return nil
}

type Socket struct {

}
func (s *Socket) Write(p []byte)(n int,err error){
	return 0,nil
}
func (s *Socket) Close()error{
	return nil
}

type Writer interface {
	Write(p []byte)(n int,err error)
}
type Closer interface {
	Close()error
}

func userWriter(writer io.Writer){
	writer.Write(nil)
}
func userCloser(closer io.Closer){
	closer.Close()
}

func main() {
	f:=new(file)
	var write DataWriter
	write=f
	write.WriteData("data")

	s:=new(Socket)
	userWriter(s)
	userCloser(s)
}
