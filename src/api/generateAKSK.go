package api

import (
	"encoding/json"
	"github.com/yangwoodstar/NovaCore/src/constString"
	"github.com/yangwoodstar/NovaCore/src/httpClient"
	"github.com/yangwoodstar/NovaCore/src/modelStruct"
)

func GenerateAKSK(url, sn, appCode, accessKey string) (*modelStruct.Credentials, error) {
	params := map[string]string{
		"expireTime": constString.CMSExpireTimeStr,
	}
	result, err := httpClient.ProcessGet(url, sn, appCode, accessKey, params)
	if err != nil {
		return nil, err
	}

	var response modelStruct.CmsHttpResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	var credentials modelStruct.Credentials

	err = json.Unmarshal(response.Data, &credentials)
	if err != nil {
		return nil, err
	}
	return &credentials, nil
}
