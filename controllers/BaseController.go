package controllers

import (
	"compress/gzip"
	"encoding/json"
	"github.com/astaxie/beego"
	"io"
	"strings"
	"time"
	"zs403_mbook_copy/common"
	"zs403_mbook_copy/models"
	"zs403_mbook_copy/utils"
)

type BaseController struct {
	beego.Controller
	Member   *models.Member   // todo 为啥把用户放到这里 这里登录的时候回初始化用户的信息【比如是否是管理员】
	EnableAnonymous bool      // 开启匿名访问
	Option          map[string]string //全局设置
}

type CookieRemember struct {
	MemberId int
	Account string
	Time time.Time
}


// 每个子类 Controller 共用方法调用前，都执行一下 Prepare 方法
// 这是一个接口，调用
func (c *BaseController) Prepare() {
	c.Member = &models.Member{}
	c.EnableAnonymous = false
	// 从session 中获取用户的信息
	if member, ok := c.GetSession(common.SessionName).(models.Member); ok && member.MemberId > 0 {
		c.Member = &member
	} else {
		// 如果 Cookie 中存在登录信息，从cookie中获取用户信息
		if cookie , ok := c.GetSecureCookie(common.AppKey(), "login"); ok {
			var cookieRemember  CookieRemember
			err := utils.Decode(cookie, &cookieRemember)
			if err == nil {
				member,err := (&models.Member{}).Find(cookieRemember.MemberId)
				if err == nil {
					c.SetMember(*member)
					c.Member = member
				}
			}
		}
	}
	if "" == c.Member.RoleName {
		c.Member.RoleName = common.Role(c.Member.MemberId)   //todo 这里对不
	}
	c.Data["Member"] = c.Member   // 前端去使用 {{if gt .Member.MemberId 0}}
	c.Data["SITE_NAME"] = "MBOOK"
	//设置全局配置
	c.Option = make(map[string]string)
	c.Option["ENABLED_CAPTCHA"] = "false"
	c.Data["BaseUrl"] = c.BaseUrl()  // todo 这个不懂
}

// 设置用户登录信息
// 获取是退出登录
func (c *BaseController) SetMember(member models.Member) {
	// 退出登录
	if member.MemberId <= 0 {
		c.DelSession(common.SessionName)
		c.DelSession("uid")
		c.DestroySession()
	} else {  // 登录
		c.SetSession(common.SessionName, member)  // 将 member 记录到 session
		c.SetSession("uid", member.MemberId)
	}
}

//todo 不知道啥意思
func (c *BaseController) BaseUrl() string {
	host := beego.AppConfig.String("sitemap_host")
	if len(host) > 0 {
		if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") {
			return host
		}
		return c.Ctx.Input.Scheme() + "://" + host
	}
	return c.Ctx.Input.Scheme() + "://" + c.Ctx.Request.Host
}


// Ajax 接口返回 Json
// todo 完全不懂了
func (c *BaseController) JsonResult (errCode int, errMsg string, data ...interface{}) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["errcode"] = errCode
	jsonData["message"] = errMsg

	if (len(data) > 0 && data[0] != nil) {
		jsonData["data"] = data[0]
	}

	// map 转换为 json 的字节流
	returnJSON, err := json.Marshal(jsonData)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 启用 gzip 压缩
	if strings.Contains(strings.ToLower(c.Ctx.Request.Header.Get("Accept-Encoding")), "gzip") {
		c.Ctx.ResponseWriter.Header().Set("Content-Encoding", "gzip")
		w := gzip.NewWriter(c.Ctx.ResponseWriter)
		defer w.Close()
		w.Write(returnJSON)
		w.Flush()
	} else {
		io.WriteString(c.Ctx.ResponseWriter, string(returnJSON))
	}
	c.StopRun()
}

// 关注或者取消关注
func (c *BaseController) SetFollow() {
	if c.Member.MemberId == 0 {
		c.JsonResult(1, "请先登录")
	}
	uid , _:= c.GetInt(":uid")
	if uid == c.Member.MemberId {
		c.JsonResult(1, "不能关注自己")
	}
	cancel, _ := (&models.Fans{}).FollowOrCancel(uid, c.Member.MemberId)
	if cancel {
		c.JsonResult(0, "已成功取消关注")
	}
	c.JsonResult(0, "已成功关注")
}