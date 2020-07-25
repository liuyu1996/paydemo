package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (

)

func WxpayTestController(c *gin.Context)  {

	//appId := ""
	//mchId := ""
	//apiKey := ""
	//client := wechat.NewClient(appId, mchId, apiKey, false)
	//
	//number := gotil.GetRandomString(32)
	//xlog.Info("out_trade_no:", number)
	//bm := make(gopay.BodyMap)
	//bm.Set("nonce_str", gotil.GetRandomString(32))
	//bm.Set("body", "H5支付")
	//bm.Set("out_trade_no", number)
	//bm.Set("total_fee", 1)
	//bm.Set("spbill_create_ip", "127.0.0.1")
	//bm.Set("notify_url", "http://www.gopay.ink")
	//bm.Set("trade_type", wechat.TradeType_H5)
	//bm.Set("device_info", "WEB")
	//bm.Set("sign_type", wechat.SignType_MD5)
	//
	//wxResp, err := client.UnifiedOrder(bm)
	//if err != nil {
	//	xlog.Errorf("client.UnifiedOrder(%+v),error:%+v", bm, err)
	//	return
	//}
	//
	//timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	//
	//// 获取小程序支付需要的paySign
	////pac := "prepay_id=" + wxRsp.PrepayId
	////paySign := GetMiniPaySign(appId, wxRsp.NonceStr, pac, SignType_MD5, timeStamp, apiKey)
	////xlog.Info("paySign:", paySign)
	//// 获取H5支付需要的paySign
	//pac := "prepay_id=" + wxResp.PrepayId
	//paySign := wechat.GetH5PaySign(appId, wxResp.NonceStr, pac, wechat.SignType_MD5, timeStamp, apiKey)
	//xlog.Debug("paySign:", paySign)



	c.JSON(http.StatusOK, gin.H{
		"code:" : 0,
		"msg:" : "test",
	})
}
