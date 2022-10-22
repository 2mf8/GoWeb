package data

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"gopkg.in/guregu/null.v3"
	_ "gopkg.in/guregu/null.v3/zero"
)

type CuberPrice struct {
	Id          int         `json:"id"`
	GroupId     int         `json:"group_id"`
	Brand       null.String `json:"brand"`
	Item        string      `json:"item"`
	Price       null.String `json:"price"`
	Shipping    null.String `json:"shipping"`
	Updater     int         `json:"updater"`
	GmtModified null.Time   `json:"gmt_modified"`
}

func GetItem(item string) (cp CuberPrice, err error) {
	cp = CuberPrice{}
	err = Db.QueryRow("select * from [kequ5060].[dbo].[zbot_price] where item = $1", item).Scan(&cp.Id, &cp.GroupId, &cp.Brand, &cp.Item, &cp.Price, &cp.Shipping, &cp.Updater, &cp.GmtModified)
	return
}

func GetItems(key string) (cps []CuberPrice, err error) {
	statment := fmt.Sprintf("select * from [kequ5060].[dbo].[zbot_price] where item like '%%%s%%'", key)
	rows, err := Db.Query(statment)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		cp := CuberPrice{}
		err = rows.Scan(&cp.Id, &cp.GroupId, &cp.Brand, &cp.Item, &cp.Price, &cp.Shipping, &cp.Updater, &cp.GmtModified)
		cps = append(cps, cp)
	}
	return
}

func (cp *CuberPrice) ItemCreate() (err error) {
	statement := "insert into [kequ5060].[dbo].[zbot_price] values ($1, $2, $3, $4, $5, $6, $7) select @@identity"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(cp.GroupId, cp.Brand, cp.Item, cp.Price, cp.Shipping, cp.Updater, cp.GmtModified).Scan(&cp.Id)
	return
}

func (cp *CuberPrice) ItemUpdate() (err error) {
	_, err = Db.Exec("update [kequ5060].[dbo].[zbot_price] set group_id = $2, brand = $3, item = $4, price = $5, shipping = $6, updater = $7, gmt_modified = $8 where ID = $1", cp.Id, cp.GroupId, cp.Brand, cp.Item, cp.Price, cp.Shipping, cp.Updater, cp.GmtModified)
	return
}

func (cp *CuberPrice) ItemDelete() (err error) {
	_, err = Db.Exec("delete from [kequ5060].[dbo].[zbot_price] where item = $1", cp.Item)
	return
}

func (cp *CuberPrice) ItemDeleteById() (err error) {
	_, err = Db.Exec("delete from [kequ5060].[dbo].[zbot_price] where ID = $1", cp.Id)
	return
}

/*func ItemSave(groupId int64, brand string, item string, price string, shipping string, updater int64, gmtModified time.Time) (err error) {
	cp := CuberPrice{
		GroupId:     groupId,
		Brand:       brand,
		Item:        item,
		Price:       price,
		Shipping:    shipping,
		Updater:     updater,
		GmtModified: gmtModified,
	}
	cp_get, err := GetItem(item)
	if err != nil {
		err = cp.ItemCreate()
		return
	}
	err = cp_get.ItemUpdate()
	return
}*/
