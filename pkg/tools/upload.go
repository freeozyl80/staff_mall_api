package tools

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func Upload(file multipart.File, filename string) error {

	out, err := os.Create(filename)
	if err != nil {
		fmt.Println("写入文件失败")
	}
	defer out.Close()
	io.Copy(out, file)

	fmt.Printf("%T \n", filename)

	return err
}
