package goPushSdk

import "github.com/ddliu/go-httpclient"

const (
	pushThroughMessageByPushId      = PUSH_API_SERVER + "/garcia/api/server/push/unvarnished/pushByPushId"
	pushNotificationMessageByPushId = PUSH_API_SERVER + "/garcia/api/server/push/varnished/pushByPushId"
	pushThroughMessageByAlias       = PUSH_API_SERVER + "/garcia/api/server/push/unvarnished/pushByAlias"
	pushNotificationMessageByAlias  = PUSH_API_SERVER + "/garcia/api/server/push/varnished/pushByAlias"
)

/**
 * 通过PushId推送透传消息
 */
func PushThroughMessageByPushId(appId string, pushIds string, messageJson string, appKey string) PushResponse {
	pushThroughMessageMap := map[string]string{
		"appId":       appId,
		"pushIds":     pushIds,
		"messageJson": messageJson,
	}

	sign := GenerateSign(pushThroughMessageMap, appKey)
	pushThroughMessageMap["sign"] = sign

	res, err := httpclient.Post(pushThroughMessageByPushId, pushThroughMessageMap)

	return ResolvePushResponse(res, err)
}

//pushId推送接口（通知栏消息）
func PushNotificationMessageByPushId(appId string, pushIds string, messageJson string, appKey string) PushResponse {
	pushNotificationMessageMap := map[string]string{
		"appId":       appId,
		"pushIds":     pushIds,
		"messageJson": messageJson,
	}

	sign := GenerateSign(pushNotificationMessageMap, appKey)
	pushNotificationMessageMap["sign"] = sign

	res, err := httpclient.Post(pushNotificationMessageByPushId, pushNotificationMessageMap)

	return ResolvePushResponse(res, err)
}

//别名推送接口（透传消息
func PushThroughMessageByAlias(appId string, alias string, messageJson string, appKey string) PushResponse {
	pushThroughMessageMap := map[string]string{
		"appId":       appId,
		"alias":       alias,
		"messageJson": messageJson,
	}

	sign := GenerateSign(pushThroughMessageMap, appKey)
	pushThroughMessageMap["sign"] = sign

	res, err := httpclient.Post(pushThroughMessageByAlias, pushThroughMessageMap)

	return ResolvePushResponse(res, err)
}

//别名推送接口（通知栏消息）
func PushNotificationMessageByAlias(appId string, alias string, messageJson string, appKey string) PushResponse {
	pushNotificationMessageMap := map[string]string{
		"appId":       appId,
		"alias":       alias,
		"messageJson": messageJson,
	}

	sign := GenerateSign(pushNotificationMessageMap, appKey)
	pushNotificationMessageMap["sign"] = sign

	res, err := httpclient.Post(pushNotificationMessageByAlias, pushNotificationMessageMap)

	return ResolvePushResponse(res, err)
}
