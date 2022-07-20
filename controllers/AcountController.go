package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/captcha"
	"regexp"
	"strings"
	"time"
	"zs403_mbook_copy/common"
	"zs403_mbook_copy/models"
	"zs403_mbook_copy/utils"
)

type AcountController struct {
	BaseController
}

var cpt *captcha.Captcha

func init()  {
	// 使用beego缓存系统存储验证码数据
}


// 注册
// 需要传入的参数 返回给前端
func (c *AcountController) Regist() {
	var (
		nickname string     // 昵称
		avatar string       // 头像的http链接地址
		email string        // 邮箱地址
		username string     // 用户名
		id  interface{}     // 用户 id
		captchaOn bool      // 是否开启了验证码
	)

	// 如果开启了验证码
	if v, ok := c.Option["ENABLED_CAPTCHA"]; ok && strings.EqualFold(v, "true") {
		captchaOn = true
		c.Data["CaptchaOn"] = captchaOn
	}

	c.Data["Nickname"] = nickname
	c.Data["Avatar"] = avatar
	c.Data["Email"] = email
	c.Data["Username"] = username
	c.Data["Id"] = id
	c.Data["RandomStr"] = time.Now().Unix()   // 随机数

	//存储标识，以标记是哪个用户，在完善用户信息的时候跟传递过来的auth和id进行校验
	c.SetSession("auth", fmt.Sprintf("%v-%v", "email", id))  // todo 这个不明白
	c.TplName = "account/bind.html"
}

// 返回一个登录页面
// 获取 cookie 中的 remember 信息 【就是回写进去】
// 如果是 post 请求【登录接口】
// 1、校验用户名和密码
// 2、更新上一次登录时间
// 3、更新CookieRemember 存入到 cookie中
func (c *AcountController) Login() {
	var remember CookieRemember
	// 验证 cookie
	if cookie, ok := c.GetSecureCookie(common.AppKey(), "login"); ok {
		if err := utils.Decode(cookie, &remember); err == nil {
			if err = c.login(remember.MemberId); err == nil {
				c.Redirect(beego.URLFor("HomeController.Index"), 302)
				return
			}
		}
	}
	c.TplName = "account/login.html"

	if c.Ctx.Input.IsPost() {
		account := c.GetString("account")
		password := c.GetString("password")
		member, err := (&models.Member{}).Login(account, password)
		fmt.Println(err)
		if err != nil {
			c.JsonResult(1, "登录失败", nil)
		}
		member.LastLoginTime = time.Now()
		member.Update()
		c.SetMember(*member)


		remember.MemberId = member.MemberId
		remember.Account = member.Account
		remember.Time = time.Now()
		v, err := utils.Encode(remember)
		if err == nil {
			c.SetSecureCookie(common.AppKey(), "login", v, 24*3600*365)
		}
		c.JsonResult(0, "ok")
	}

	c.Data["RandomStr"] = time.Now().Unix()
}


// 注册
// 用户注册后，会产生一条 member 数据
// 将用户的信息 存入到 cookie 中
func (c *AcountController) DoRegist() {
	var err error
	account := c.GetString("account")
	nickname := strings.TrimSpace(c.GetString("nickname"))
	password1 := c.GetString("password1")
	password2 := c.GetString("password2")
	email := c.GetString("email")

	member := &models.Member{}
	if password1 != password2 {
		c.JsonResult(1, "登录密码与确认密码不一致")
	}

	if l := strings.Count(password1, ""); password1 == "" || l > 20 || l < 6 {
		c.JsonResult(1, "密码必须在6-20个字符之间")
	}

	if ok, err := regexp.MatchString(common.RegexpEmail, email); !ok || err != nil || email == "" {
		c.JsonResult(1, "邮箱格式错误")
	}

	if l := strings.Count(nickname, "") - 1; l < 2 || l > 20 {
		c.JsonResult(1, "用户昵称限制在2-20个字符")
	}

	member.Account = account
	member.Nickname = nickname
	member.Password = password1
	if account == "admin" || account == "administrator" {  // 特殊账号
		member.Role = common.MemberSuperRole
	} else {
		member.Role = common.MemberGeneralRole
	}
	member.Avatar = common.DefaultAvatar()    // 头像
	member.CreateAt = 0
	member.Email = email
	member.Status = 0
	if err := member.Add(); err != nil {
		beego.Error(err)
		c.JsonResult(1, err.Error())
	}
	if err = c.login(member.MemberId); err != nil {
		beego.Error(err)
		c.JsonResult(1, err.Error())
	}
	c.JsonResult(0, "注册成功")
}

// 退出登录
func (c *AcountController) Logout() {
	c.SetMember(models.Member{})   // 清空用户信息
	c.SetSecureCookie(common.AppKey(), "login", "",-3600)  // cookie 设置过期
	c.Redirect(beego.URLFor("AccountController.Login"), 302) // 跳转到登录页面
}

/*
* 私有函数
 */
//封装一个内部调用的函数，login
// todo 这里为啥没有校验密码
func (c *AcountController) login(memberId int) (err error) {
	member, err := (&models.Member{}).Find(memberId)
	if err != nil {
		return err
	}
	if member.MemberId == 0 {
		return errors.New("用户不存在")
	}
	// todo 校验密码

	// 更新最后登录时间
	member.LastLoginTime = time.Now()
	member.Update()
	// 将信息  __mbook_session__   member
	//  uid  member.MemberId
	c.SetMember(*member)
	var remember CookieRemember
	remember.MemberId = member.MemberId
	remember.Account = member.Account
	remember.Time = time.Now()

	// 序列化为字节流
	v, err := utils.Encode(remember)
	if err == nil {
		c.SetSecureCookie(common.AppKey(), "login", v, 24 * 3600 * 365)  // 保存了 remember 结构
	}
	return err
}

