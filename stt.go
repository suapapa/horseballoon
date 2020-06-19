package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type sttResponse struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// https://developers.kakao.com/docs/latest/ko/voice/rest-api
func kakaoSTT(ctx context.Context, r io.Reader) {
	req, err := http.NewRequestWithContext(ctx, "POST",
		"https://kakaoi-newtone-openapi.kakao.com/v1/recognize",
		r,
	)
	chk(err)
	req.Header.Set("Transfer-Encoding", "chunked")
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Authorization", "KakaoAK "+os.Getenv("DEVKAKAO_APIKEY"))

	resp, err := http.DefaultClient.Do(req)
	chk(err)
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
	_, params, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	mr := multipart.NewReader(resp.Body, params["boundary"])
	for part, err := mr.NextPart(); err == nil; part, err = mr.NextPart() {
		select {
		case <-ctx.Done():
			return
		default:
		}
		var resp sttResponse
		err := json.NewDecoder(part).Decode(&resp)
		chk(err)
		if resp.Type == "finalResult" {
			// log.Println(resp.Value)
			// fmt.Println(translate(ctx, resp.Value))
			en, err := translate(ctx, resp.Value)
			if err != nil {
				log.Printf("translate failed: %v", err)
				continue
			}

			game.Lock()
			game.Str = en
			game.Start = time.Now()
			game.Unlock()
		}
	}
}
