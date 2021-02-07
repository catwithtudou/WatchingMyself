package _1_flyweight

import "fmt"

/**
 *@Author tudou
 *@Date 2021/1/14
 **/

//对象工厂
type FileFlyweightFactory struct{
	maps map[string]*FileFlyweight
}

var fileFactory *FileFlyweightFactory

func GetFileFlyweightFactory()*FileFlyweightFactory{
	if fileFactory==nil{
		fileFactory=&FileFlyweightFactory{maps: make(map[string]*FileFlyweight)}
	}
	return fileFactory
}

func (f *FileFlyweightFactory)Get(fileName string)*FileFlyweight{
	file,ok:=f.maps[fileName]
	if !ok{
		file = NewFileFlyweight(fileName)
		f.maps[fileName]=file
	}
	return file
}


type FileFlyweight struct{
	data string
}

func NewFileFlyweight(fileName string)*FileFlyweight{
	data:=fmt.Sprintf("file name:%s",fileName)
	return &FileFlyweight{data: data}
}

func (f *FileFlyweight)Data()string{
	return f.data
}

type FileViewer struct{
	*FileFlyweight
}

func NewFileViewer(fileName string)*FileViewer{
	file:=GetFileFlyweightFactory().Get(fileName)
	return &FileViewer{
		FileFlyweight:file,
	}
}

func (f *FileViewer)PrintData(){
	fmt.Println(f.Data())
}


