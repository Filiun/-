package modes

import (
	"crm/db"
	"github.com/go-xorm/xorm"
)

func Db(index int) *xorm.Engine {
	return db.UseHand(index)
}