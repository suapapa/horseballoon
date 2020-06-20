// Copyright 2020 Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type translateResp struct {
	TranslatedText [][]string `json:"translated_text"`
}

// https://developers.kakao.com/docs/latest/ko/translate/dev-guide
func translate(ctx context.Context, src string) (string, error) {
	log.Println("kr:", src)
	qVal := make(url.Values)
	qVal.Add("src_lang", "kr")
	qVal.Add("target_lang", "en")
	qVal.Add("query", src)

	body := strings.NewReader(qVal.Encode())
	req, err := http.NewRequestWithContext(ctx,
		"POST",
		"https://kapi.kakao.com/v1/translation/translate",
		body,
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "KakaoAK "+os.Getenv("DEVKAKAO_APIKEY"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var trResp translateResp
	err = json.NewDecoder(resp.Body).Decode(&trResp)
	if err != nil {
		return "", err
	}
	en := trResp.TranslatedText[0][0]
	log.Println("en:", en)

	return en, nil
}
