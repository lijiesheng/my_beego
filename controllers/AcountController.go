package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"time"
	"zs403_mbook_copy/common"
	"zs403_mbook_copy/models"
	"zs403_mbook_copy/utils"
)

type AcountController struct {
	BaseController
}

func init()  {
	// 使用beego缓存系统存储验证码数据
}


// 注册
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
	c.Data["RandomStr"] = time.Now().Unix()

	//存储标识，以标记是哪个用户，在完善用户信息的时候跟传递过来的auth和id进行校验
	c.SetSession("auth", fmt.Sprintf("%v-%v", "email", id))  // todo 这个不明白
	c.TplName = "account/bind.html"
}

// 登录
func (c *AcountController) Login() {
	var remember CookieRemember
	// 验证 cookie
	if cookie, ok := c.GetSecureCookie(common.AppKey(), "login"); ok {
		if err := utils.Decode(cookie, &remember); err != nil {
			if err = c.login(remember.MemberId); err == nil {
				c.Redirect(beego.URLFor("HomeController.Index"), 302)
				return
			}
		}
	}
	c.TplName = "account/login.html"

	if c.Ctx.Input.IsPost() {  // 如果是 post 请求
		account := c.GetString("account")
		password := c.GetString("password")
		member, err := (&models.Member{}).Login(account, password)
	}
}

// 退出登录
func (c *AcountController) Logout() {
	c.SetMember(models.Member{})
	c.SetSecureCookie(common.AppKey(), "login", "",-3600)
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

