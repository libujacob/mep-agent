/*
 *  Copyright 2020 Huawei Technologies Co., Ltd.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

// get token service
package service

import (
	"encoding/json"
	"mep-agent/src/config"
	"mep-agent/src/model"
	"mep-agent/src/util"
	"unsafe"
	log "github.com/sirupsen/logrus"
)

// Returns token
func GetMepToken(auth model.Auth) (*model.TokenModel, error) {

	log.Info("begin to get token from mep_auth")
	server, errGetServer := config.GetServerUrl()
	if errGetServer != nil {
		// clear sk
		sk := auth.SecretKey
		util.ClearByteArray(*sk)
		return nil, errGetServer
	}

	url := server.MepAuthUrl
	resp, errPostRequest := PostTokenRequest("", url, auth)
	if errPostRequest != nil {
		return nil, errPostRequest
	}

	var token model.TokenModel
	errJson := json.Unmarshal([]byte(resp), &token)
	respMsg := *(*[]byte)(unsafe.Pointer(&resp))
	util.ClearByteArray(respMsg)
	if errJson != nil {
		return nil, errJson
	}

	log.Info("get token success.")
	return &token, nil
}