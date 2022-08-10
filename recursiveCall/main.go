package main

import "reflect"

type Menu struct {
	Id       int64   //id
	SiteId   int64   //站点id
	MenuName string  //目录名称
	PmenuId  int64   //父级id
	MenuUrl  string  //链接地址
	IconId   string  //目录图标
	Sorts    int64   //排序
	Child    []*Menu //多个子级目录
}

func GetMenu() []*Menu {

	//定义指针切片用来存储所有菜单
	var menus []*Menu

	//定义指针切片返回控制器
	var res []*Menu

	//查询所有pmenu_id为0的数据,也就是一级菜单,site_id是站点,可以忽略
	Db.Table("hb_menu").Where("pmenu_id=?", 0).Where("site_id=?", 1).Order("id desc").Find(&menus)

	//判断是否存在数据,存在进行树状图重构
	if reflect.ValueOf(menus).IsValid() {
		//将一级菜单传递给回调函数
		res = tree(menus)
	}
	return res
}

//生成树结构
func tree(menus []*Menu) []*Menu {
	//定义子节点目录
	var nodes []*Menu
	if reflect.ValueOf(menus).IsValid() {
		//循环所有一级菜单
		for k, v := range menus {
			//查询所有该菜单下的所有子菜单
			Db.Table("hb_menu").Where("pmenu_id= ?", v.Id).Find(&nodes)

			//将子菜单的数据循环赋值给父菜单
			for kk, _ := range nodes {
				menus[k].Child = append(menus[k].Child, nodes[kk])
			}
			//将刚刚查询出来的子菜单进行递归,查询出三级菜单和四级菜单
			tree(nodes)
		}
	}
	return menus
}

func main() {

}
