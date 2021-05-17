package modes

import "fmt"

type Article struct {
	Id      int64  `json:"id"  xorm:"id"`
	Name    string `json:"name"  xorm:"name"`
	Content string `json:"content" xorm:"content"`
	Author  string `json:"author" xorm:"author"`
	Uid     int64 `json:"uid"  xorm:"uid"`
}

func (this *Article) TableName() string {
	return "article"
}

func (this *Article)Add()(int64, error)  {
	return Db(0).Insert(this)
}

func (this *Article)List(uid  int)([]Article,error) {

	list:=make([]Article,0)
	where:=fmt.Sprintf("uid=%d",uid)
	err:= Db(0).Where(where).Find(&list)
	return list,err
}

func (this *Article)Info(id int64) error{

	where:=fmt.Sprintf("id=%d",id)
	_,err:=Db(0).Where(where).Get(this)
	return err
}



