package controllers

import (
	"beefile/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller//匿名字段
}

func (c *MainController) Get() {
//	name1:=c.GetString("name")
//	age1,err:=c.GetInt("age")
	name:=c.Ctx.Input.Query("name")
	age:=c.Ctx.Input.Query("age")
	sex:=c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
//以admin，18为正确数据进行验证
	if name != "admin" || age != "18" {
		c.Ctx.ResponseWriter.Write([]byte("数据验证不成功"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据验证成功"))

	//c.Data["Website"] = "www.baidu.com"
	//c.Data["Email"] = "2692596075@qq.com"
	//c.TplName = "index.tpl"
}
//该post方法是处理post类型的请求的时候要调用的方法
func (c *MainController) Post(){
	//fmt.Println("post类型的请求")
	//user:=c.Ctx.Request.FormValue("user")
	//fmt.Println("用户名为：",user)
	//psd:=c.Ctx.Request.FormValue("psd")
	//
	//if user!= "yiliu"||psd!="123456" {
    //  c.Ctx.ResponseWriter.Write([]byte("对不起，数据不正确"))
	//	return
	//}
	//c.Ctx.ResponseWriter.Write([]byte("恭喜你，数据准确"))
	//fmt.Println("密码是：",psd)
	//c.Data["Website"] = "www.baidu.com"
	//c.Data["Email"] = "2692596075@qq.com"
	//c.TplName = "index.tpl"
	/*body:=c.Ctx.Request.Body*/
	dataByes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收失败，请重试")
		return
	}
	var person models.Person
	json.Unmarshal(dataByes,&person)
	if err != nil {
		c.Ctx.WriteString("数据接收失败，请重试")
		return
	}
	fmt.Println("用户名：",person.User,",年龄：",person.Birthday,"地址：",person.Address,"nick:",person.Nick)
	c.Ctx.WriteString("用户名是："+person.User,)
}
