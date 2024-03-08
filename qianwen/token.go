package qianwen

import (
	"fmt"
	apiClient "github.com/alibabacloud-go/bailian-20230601/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"tongyiqwen/config"
)

func CreateToken() {
	conf := config.GetConfig()
	accessKeyId := conf.AccessKey
	accessKeySecret := conf.AccessSecretKey
	agentKey := conf.AgentKey
	//appId := conf.ID
	endpoint := "bailian.cn-beijing.aliyuncs.com"

	option := &openapi.Config{AccessKeyId: &accessKeyId,
		AccessKeySecret: &accessKeySecret,
		Endpoint:        &endpoint,
	}

	tokenClient, err := apiClient.NewClient(option)
	if err != nil {
		fmt.Printf("failed to new client, err: %v\n", err)
		return
	}

	request := &apiClient.CreateTokenRequest{AgentKey: &agentKey}
	result, err := tokenClient.CreateToken(request)
	if err != nil {
		fmt.Printf("failed to create token, err: %v\n", err)
		return
	}

	resultBody := result.Body
	if !(*resultBody.Success) {
		requestId := resultBody.RequestId
		if requestId == nil {
			requestId = result.Headers["x-acs-request-id"]
		}

		errMessage := fmt.Sprintf("Failed to create token, reason: %s RequestId: %s", *resultBody.Message, *requestId)
		fmt.Printf("%v\n", errMessage)
	}
	token := *resultBody.Data.Token
	fmt.Printf("token: %s, expiredTime : %d\n", token, *resultBody.Data.ExpiredTime)
	config.RefreshToken(token)
}
