package dao

import (
	"gin/api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

type Dao struct {
	db *gorm.DB
}

//添加数据
func (d *Dao) Insert(order *model.DemoOrder) (id int,err error){

	var student model.Student
	student.Number="20163512"
	//开启事务
	tx := Eloquent.Begin()

	if err = tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Create(&student).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	id = order.Id
	return
}

//数据查询
func (d *Dao) Query() (orders []model.DemoOrder, err error) {
	if err = Eloquent.Find(&orders).Error; err != nil {
		return
	}
	return
}

//模糊查询
func (d *Dao) VagueQuery(username string) (orders []model.DemoOrder, err error) {
	if err = Eloquent.Where("username like?","%"+username+"%").Find(&orders).Error; err != nil {
		return
	}
	return
}

//排序查询
func (d *Dao) SortQuery(sort string) (orders []model.DemoOrder, err error) {
	//根据amount排序，排序顺序为desc
	if err = Eloquent.Order(sort+" desc").Find(&orders).Order("amount",true).Error; err != nil {

	}
	return
}

//单个查询
func (d *Dao) Get(id int) (orders model.DemoOrder, err error) {
	if err = Eloquent.Find(&orders,id).Error; err != nil {
		return
	}
	return
}

//修改数据
func (d *Dao) Update(order *model.DemoOrder) (updateOrder model.DemoOrder, err error) {
	//查询ID是否存在
	if err = Eloquent.Select([]string{"id", "username"}).First(&updateOrder, order.Id).Error; err != nil {
		return
	}
	if err = Eloquent.Model(&updateOrder).Updates(&order).Error; err != nil {
		return
	}
	return
}