package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type FriendController struct {
	beego.Controller
}

type FrontFriends struct {
	UserName string `form:"userName"`
	FriendName string `form:"friendName"`
	Message string `form:"message"`
}


// post请求，表单格式 userName|friendName|message
func (this *FriendController) MakeFriends() {
	beego.Informational("好友申请!")
	frontFriends := FrontFriends{}
	err := this.ParseForm(&frontFriends)
	beego.Informational("前端参数： UserName=" + frontFriends.UserName + ", FriendName = "+ frontFriends.FriendName + ", Message = " + frontFriends.Message)
	cookieToken := this.Ctx.GetCookie("token")
	goSessionId := this.Ctx.GetCookie("gosessionid")
	beego.Informational(goSessionId)
	sess, errSess := globalSessions.GetSessionStore(goSessionId)
	if errSess != nil {
		beego.Informational("session获取失败，检查session设置！")

		this.Ctx.WriteString("session设置有问题！！")
		return
	}

	sessionToken := sess.Get(frontFriends.UserName)
	if cookieToken != sessionToken {
		beego.Informational(sess)
		beego.Informational("token不一致，需要重新登录！ cookieToken = " + cookieToken + "; sessionToken = ")
		beego.Informational(sessionToken)
		this.Ctx.WriteString("未登录，请登录！")
		return
	}

	if err != nil {
		beego.Informational("前端传惨错误，请检查！")
		this.Ctx.WriteString("前端传惨错误，请检查！")
		return
	} else {
		user := models.GetUserByName(frontFriends.UserName)
		friend := models.GetUserByName(frontFriends.FriendName)
		applyFriend := models.Friend{UserId:user.Id, FriendId:friend.Id, MarkName:friend.UserName,
			Message:frontFriends.Message, Agree:false}
		err := models.ApplyFriend(&applyFriend)
		if err == nil {
			beego.Informational("好友申请已发送，等待通过！")
			this.Ctx.WriteString("好友申请已发送，等待通过")
			return
		} else {
			beego.Informational(err)
		}
	}
	this.Ctx.WriteString("好友申请失败!")
}

//post请求
func (this *FriendController) AgreeApply() {
	userName := "shun"
	friends := models.GetApplyByName(userName)
	for _, friend := range friends {
		beego.Informational(friend)
	}
	this.Ctx.WriteString("申请列表！")
}