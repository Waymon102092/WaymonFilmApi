package services

import (
	"Waymon_api/cache"
	"Waymon_api/dao"
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/log"
	"Waymon_api/pkg/res"
	"Waymon_api/pkg/waymon"
	"Waymon_api/pkg/wechat/mini"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

type AuthorService struct {
	Code string `json:"code" form:"code"`
}

type LoginService struct {
	MediaId       int64  `json:"media_id" form:"media_id"`
	PromoteId     int64  `json:"promote_id" form:"promote_id"`
	MemberId      int64  `json:"member_id" form:"member_id"`
	Province      string `json:"province" form:"province"`
	City          string `json:"city" form:"city"`
	District      string `json:"district" form:"district"`
	UnionId       string `json:"unionid" form:"unionid"`
	OpenId        string `json:"openid" form:"openid"`
	SessionKey    string `json:"session_key" form:"session_key"`
	EncryptedData string `json:"encryptedData" form:"encryptedData"`
	Iv            string `json:"iv" form:"iv"`
	Code          string `json:"code" form:"code"`
	NickName      string `json:"nick_name" form:"nick_name"`
	AvatarUrl     string `json:"avatar_url" form:"avatar_url"`
}

func (author *AuthorService) Author(ctx context.Context) res.Response {
	code := e.Success
	if author.Code == "" {
		code = e.ParamError
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	appId := viper.GetString("mini.appId")
	appSecret := viper.GetString("mini.appSecret")
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appId, appSecret, author.Code)
	json, err := waymon.GetJson(url)
	if err != nil {
		code = e.GetJsonError
		log.WaymonLogger.Error("GetJson" + err.Error())
		zap.S().Error("GetJson" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   json,
	}
}

func (author *AuthorService) Session(ctx context.Context) res.Response {
	code := e.Success
	if author.Code == "" {
		code = e.ParamError
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	appId := viper.GetString("tt.appId")
	appSecret := viper.GetString("tt.appSecret")
	url := "https://developer.toutiao.com/api/apps/v2/jscode2session"
	params := make(map[string]interface{})
	params["appid"] = appId
	params["secret"] = appSecret
	params["code"] = author.Code
	json, err := waymon.PostJson(url, params)
	if err != nil {
		code = e.GetJsonError
		log.WaymonLogger.Error("GetJson" + err.Error())
		zap.S().Error("GetJson" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   json,
	}
}

func (login *LoginService) Login(ctx context.Context) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	condition["union_id"] = login.UnionId
	memberDao := dao.NewMemberDao()
	member, _ := memberDao.MemberInfo(condition)
	//
	c := make(map[string]interface{})
	c["id"] = login.Code
	parent, _ := memberDao.MemberInfo(c)
	var status int
	var staffId int64
	var parentId int64
	if parent.Status == 3 { //员工
		parentId = 0
		staffId = int64(parent.ID)
		status = 1
	}
	if parent.Status == 1 { //代理商
		parentId = int64(parent.ID)
		staffId = 0
		status = 0
	}
	if member.ID == 0 {
		//创建用户
		member.Tag = 1
		member.MediaId = login.MediaId
		member.PromoteId = login.PromoteId
		member.ParentId = parentId
		member.StaffId = staffId
		member.UnionId = login.UnionId
		member.OpenId = login.OpenId
		member.NickName = login.NickName
		member.AvatarUrl = login.AvatarUrl
		member.Province = login.Province
		member.City = login.City
		member.District = login.District
		member.Time = time.Now().Unix()
		member.Status = status
		err := memberDao.MemberAdd(&member)
		if err != nil {
			code = e.MemberAddError
			zap.S().Error("MemberAddError" + err.Error())
			log.WaymonLogger.Error("MemberAddError" + err.Error())
			return res.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   nil,
			}
		}
		//统计数据
		go func() {
			if login.MediaId > 0 {
				cache.IncrByMediaMember(login.MediaId)
			}
			if login.PromoteId > 0 {
				cache.IncrByPromoteMember(login.PromoteId)
			}
		}()
	} else { //当前用户已存在
		if member.Status == 0 {
			member.ParentId = parentId
			member.StaffId = staffId
			member.Status = status
		}
		if member.OpenId == "" {
			member.OpenId = login.OpenId
		}
		if member.Province != "" {
			member.Province = login.Province
		}
		if member.City != "" {
			member.City = login.City
		}
		if member.District != "" {
			member.District = login.District
		}
		err := memberDao.MemberEdit(int64(member.ID), &member)
		if err != nil {
			code = e.MemberEditError
			zap.S().Error("MemberEditError" + err.Error())
			log.WaymonLogger.Error("MemberEditError" + err.Error())
			return res.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   nil,
			}
		}
	}
	//token
	token, err := waymon.GenerteToken(member.ID, member.OpenId)
	if err != nil {
		code = e.GenerateTokenError
		zap.S().Error("GenerateTokenError" + err.Error())
		log.WaymonLogger.Error("GenerateTokenError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	result := make(map[string]interface{})
	result["token"] = token
	result["member"] = map[string]interface{}{
		"id":         member.ID,
		"nick_name":  member.NickName,
		"avatar_url": member.AvatarUrl,
		"tel":        member.Tel,
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   result,
	}
}

func (login *LoginService) TTLogin(ctx context.Context) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	condition["tt_union_id"] = login.UnionId
	memberDao := dao.NewMemberDao()
	member, _ := memberDao.MemberInfo(condition)
	//
	c := make(map[string]interface{})
	c["id"] = login.Code
	parent, _ := memberDao.MemberInfo(c)
	var status int
	var staffId int64
	var parentId int64
	if parent.Status == 3 { //员工
		parentId = 0
		staffId = int64(parent.ID)
		status = 1
	}
	if parent.Status == 1 { //代理商
		parentId = int64(parent.ID)
		staffId = 0
		status = 0
	}
	if member.ID == 0 {
		//创建用户
		member.Tag = 3
		member.MediaId = login.MediaId
		member.PromoteId = login.PromoteId
		member.ParentId = parentId
		member.StaffId = staffId
		member.TTUnionId = login.UnionId
		member.TTOpenId = login.OpenId
		member.NickName = login.NickName
		member.AvatarUrl = login.AvatarUrl
		member.Province = login.Province
		member.City = login.City
		member.District = login.District
		member.Time = time.Now().Unix()
		member.Status = status
		err := memberDao.MemberAdd(&member)
		if err != nil {
			code = e.MemberAddError
			zap.S().Error("MemberAddError" + err.Error())
			log.WaymonLogger.Error("MemberAddError" + err.Error())
			return res.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   nil,
			}
		}
		//统计数据
		go func() {
			if login.MediaId > 0 {
				cache.IncrByMediaMember(login.MediaId)
			}
			if login.PromoteId > 0 {
				cache.IncrByPromoteMember(login.PromoteId)
			}
		}()
	} else { //当前用户已存在
		if member.Status == 0 && parentId > 0 {
			member.ParentId = parentId
		}
		if member.OpenId == "" {
			member.TTOpenId = login.OpenId
		}
		if member.UnionId == "" {
			member.TTUnionId = login.UnionId
		}
		if member.Province != "" {
			member.Province = login.Province
		}
		if member.City != "" {
			member.City = login.City
		}
		if member.District != "" {
			member.District = login.District
		}
		err := memberDao.MemberEdit(int64(member.ID), &member)
		if err != nil {
			code = e.MemberEditError
			zap.S().Error("MemberEditError" + err.Error())
			log.WaymonLogger.Error("MemberEditError" + err.Error())
			return res.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   nil,
			}
		}
	}
	//token
	token, err := waymon.GenerteToken(member.ID, member.TTOpenId)
	if err != nil {
		code = e.GenerateTokenError
		zap.S().Error("GenerateTokenError" + err.Error())
		log.WaymonLogger.Error("GenerateTokenError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	result := make(map[string]interface{})
	result["token"] = token
	result["member"] = map[string]interface{}{
		"id":         member.ID,
		"nick_name":  member.NickName,
		"avatar_url": member.AvatarUrl,
		"tel":        member.Tel,
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   result,
	}
}

// 小程序手机号码登录

func (login *LoginService) Register(ctx context.Context) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	condition["union_id"] = login.UnionId
	memberDao := dao.NewMemberDao()
	member, _ := memberDao.MemberInfo(condition)
	//
	c := make(map[string]interface{})
	c["id"] = login.Code
	parent, _ := memberDao.MemberInfo(c)
	var status int
	var staffId int64
	var parentId int64
	if parent.Status == 3 { //员工
		parentId = 0
		staffId = int64(parent.ID)
		status = 1
	}
	if parent.Status == 1 { //代理商
		parentId = int64(parent.ID)
		staffId = 0
		status = 0
	}
	//
	if member.ID == 0 {
		//解密
		tel, err := mini.Decrypt(login.EncryptedData, login.SessionKey, login.Iv)
		if err != nil {
			code = e.DecryptError
			log.WaymonLogger.Error("DecryptError" + err.Error())
			zap.S().Error("DecryptError" + err.Error())
			return res.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   nil,
			}
		}
		//创建用户
		member.Tag = 1
		member.MediaId = login.MediaId
		member.PromoteId = login.PromoteId
		member.ParentId = parentId
		member.StaffId = staffId
		member.Tel = tel
		member.UnionId = login.UnionId
		member.OpenId = login.OpenId
		member.NickName = waymon.WaymonTel(tel)
		member.AvatarUrl = "https://cinema.img.sainiao.net/avatar.png"
		member.Province = login.Province
		member.City = login.City
		member.District = login.District
		member.Time = time.Now().Unix()
		member.Status = status
		err = memberDao.MemberAdd(&member)
		if err != nil {
			code = e.MemberAddError
			zap.S().Error("MemberAddError" + err.Error())
			log.WaymonLogger.Error("MemberAddError" + err.Error())
			return res.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   nil,
			}
		}
		//统计数据
		go func() {
			if login.MediaId > 0 {
				cache.IncrByMediaMember(login.MediaId)
			}
			if login.PromoteId > 0 {
				cache.IncrByPromoteMember(login.PromoteId)
			}
		}()
	} else { //当前用户已存在
		if member.Status == 0 {
			member.ParentId = parentId
			member.StaffId = staffId
			member.Status = status
		}
		if member.OpenId == "" {
			member.OpenId = login.OpenId
		}
		if member.Province != "" {
			member.Province = login.Province
		}
		if member.City != "" {
			member.City = login.City
		}
		if member.District != "" {
			member.District = login.District
		}
		err := memberDao.MemberEdit(int64(member.ID), &member)
		if err != nil {
			code = e.MemberEditError
			zap.S().Error("MemberEditError" + err.Error())
			log.WaymonLogger.Error("MemberEditError" + err.Error())
			return res.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   nil,
			}
		}
	}
	//token
	token, err := waymon.GenerteToken(member.ID, member.OpenId)
	if err != nil {
		code = e.GenerateTokenError
		zap.S().Error("GenerateTokenError" + err.Error())
		log.WaymonLogger.Error("GenerateTokenError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	result := make(map[string]interface{})
	result["token"] = token
	result["member"] = map[string]interface{}{
		"id":         member.ID,
		"nick_name":  member.NickName,
		"avatar_url": member.AvatarUrl,
		"tel":        member.Tel,
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   result,
	}
}

// 抖音手机号码登录

func (login *LoginService) TTRegister(ctx context.Context) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	condition["tt_union_id"] = login.UnionId
	memberDao := dao.NewMemberDao()
	member, _ := memberDao.MemberInfo(condition)
	//
	c := make(map[string]interface{})
	c["id"] = login.Code
	parent, _ := memberDao.MemberInfo(c)
	var status int
	var staffId int64
	var parentId int64
	if parent.Status == 3 { //员工
		parentId = 0
		staffId = int64(parent.ID)
		status = 1
	}
	if parent.Status == 1 { //代理商
		parentId = int64(parent.ID)
		staffId = 0
		status = 0
	}
	//
	if member.ID == 0 {
		//解密
		tel, err := mini.Decrypt(login.EncryptedData, login.SessionKey, login.Iv)
		if err != nil {
			code = e.DecryptError
			log.WaymonLogger.Error("DecryptError" + err.Error())
			zap.S().Error("DecryptError" + err.Error())
			return res.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   nil,
			}
		}
		//创建用户
		member.Tag = 3
		member.MediaId = login.MediaId
		member.PromoteId = login.PromoteId
		member.ParentId = parentId
		member.StaffId = staffId
		member.Tel = tel
		member.TTUnionId = login.UnionId
		member.TTOpenId = login.OpenId
		member.NickName = waymon.WaymonTel(tel)
		member.AvatarUrl = "https://cinema.img.sainiao.net/avatar.png"
		member.Province = login.Province
		member.City = login.City
		member.District = login.District
		member.Time = time.Now().Unix()
		member.Status = status
		err = memberDao.MemberAdd(&member)
		if err != nil {
			code = e.MemberAddError
			zap.S().Error("MemberAddError" + err.Error())
			log.WaymonLogger.Error("MemberAddError" + err.Error())
			return res.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   nil,
			}
		}
	} else { //当前用户已存在
		if member.Status == 0 && parentId > 0 {
			member.ParentId = parentId
		}
		if member.OpenId == "" {
			member.OpenId = login.OpenId
		}
		if member.Province != "" {
			member.Province = login.Province
		}
		if member.City != "" {
			member.City = login.City
		}
		if member.District != "" {
			member.District = login.District
		}
		err := memberDao.MemberEdit(int64(member.ID), &member)
		if err != nil {
			code = e.MemberEditError
			zap.S().Error("MemberEditError" + err.Error())
			log.WaymonLogger.Error("MemberEditError" + err.Error())
			return res.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   nil,
			}
		}
	}
	//token
	token, err := waymon.GenerteToken(member.ID, member.OpenId)
	if err != nil {
		code = e.GenerateTokenError
		zap.S().Error("GenerateTokenError" + err.Error())
		log.WaymonLogger.Error("GenerateTokenError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	result := make(map[string]interface{})
	result["token"] = token
	result["member"] = map[string]interface{}{
		"id":         member.ID,
		"nick_name":  member.NickName,
		"avatar_url": member.AvatarUrl,
		"tel":        member.Tel,
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   result,
	}
}
