package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"task2/internal/model"
)


type dao struct {
	db *sql.DB
}

//返回一个dao的对象吧
func NewDao() (*dao,error){
	b, err :=NewDB()
	if err != nil {
		err = errors.Wrap(err, "NewDB Failed!")
		return nil, err
	}
	d := &dao {
		db: b,
	}
	return d, nil
}

func (d *dao) QueryStudent(id int64) (model.Students, error) {
	var student model.Students
	err := d.db.QueryRow("select * from student where id = ? ", id).Scan(&student.ID,&student.Name,&student.Phone)
	if err != nil {
		err = errors.Wrap(err,"Student is not Exit!")
		return model.Students{}, err
	}
	return student, nil
}