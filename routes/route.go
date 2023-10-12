package routes

import (
	api "Waymon_api/controller/v1"
	"Waymon_api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	//api
	v1 := r.Group("api/v1")
	{
		//微信小程序
		v1.POST("author", api.Author)
		v1.POST("login", api.Login)
		v1.POST("register", api.Register)
		//抖音小程序
		v1.POST("session", api.Session)
		v1.POST("tt_login", api.TTLogin)
		v1.POST("tt_register", api.TTRegister)
		//支付宝小程序

		//city
		v1.GET("city/add", api.CityAdd)
		v1.GET("city/current", api.CityCurrent)
		//movie
		v1.GET("movie/hot", api.MovieHot)
		v1.GET("movie/coming", api.MovieComing)
		//district
		v1.GET("district/filter", api.DistrictFilter)
		//cinema
		v1.GET("cinema/list", api.CinemaList)

		//退款回调
		//v1.POST("order/wx_refund_notify", api.OrderWxRefundNotify)
		//v1.POST("order/tt_refund_notify", api.OrderTTRefundNotify)
		//v1.POST("order/ali_refund_notify", api.OrderAliRefundNotify)

		token := v1.Group("")
		token.Use(middleware.Jwt())
		{
			// 工具类
			//banner
			v1.GET("banner/list", api.BannerList)
			//brand
			v1.GET("brand/list", api.BrandList)
			//report
			token.POST("report/add", api.ReportAdd)
			//report_category
			v1.GET("report_category/list", api.ReportCategoryList)
			//help
			v1.GET("help/info", api.HelpInfo)
			v1.GET("help/list", api.HelpList)
			//poster
			token.GET("poster/list", api.PosterList)
			token.GET("poster/code", api.PosterCode)
			//custom
			v1.GET("custom/info", api.CustomInfo)
			//upload
			v1.POST("upload", api.Upload)

			//账户类
			//account
			token.GET("account/info", api.AccountInfo)
			token.POST("account/add", api.AccountAdd)
			//amount
			token.GET("amount/list", api.AmountList)
			token.GET("amount/money", api.AmountMoney)
			token.GET("amount/settle", api.AmountSettle)
			token.GET("amount/accumulate", api.AmountAccumulate)
			//withdraw
			token.GET("withdraw/info", api.WithdrawInfo)
			token.GET("withdraw/list", api.WithdrawList)
			token.POST("withdraw/add", api.WithdrawAdd)
			token.POST("withdraw/edit", api.WithdrawEdit)
			token.POST("withdraw/status", api.WithdrawStatus)
			token.GET("withdraw/money", api.WithdrawMoney)

			//用户类
			//fans
			token.GET("fans/list", api.FansList)
			token.GET("fans/count", api.FansCount)
			token.GET("fans/order_count", api.FansOrderCount)
			token.GET("fans/order", api.FansOrder)

			//订单类

			//refund
			//v1.GET("order/wx_refund", api.OrderWxRefund)
			//v1.POST("order/tt_refund", api.OrderTtPayRefund)
			//v1.POST("order/ali_refund", api.OrderAliPayRefund)
		}
	}
	return r
}
