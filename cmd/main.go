package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"task2/internal/dao"
)

func main() {
	d,err := dao.NewDao()
	if err != nil {
		//打印错误源头
		fmt.Printf("Source error: %T %v\n", errors.Cause(err), errors.Cause(err))
		//打印堆栈
		fmt.Printf("stack strace: \n%+v\n", err)
		panic(err)
	}

	student, err := d.QueryStudent(123)
	if err != nil {
		//打印错误源头
		fmt.Printf("Source error: %T %v\n", errors.Cause(err), errors.Cause(err))
		return
	}
	fmt.Println(student)

}