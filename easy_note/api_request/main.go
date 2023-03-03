// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/cloudwego/biz-demo/easy_note/api_request/api_service"
	"github.com/cloudwego/biz-demo/easy_note/hertz_gen/demoapi"
)

var action = flag.String("action", "", "执行的操作")

func main() {
	flag.Parse()
	fmt.Println(*action)
	switch *action {
	case "register":
		CreateUser()
	case "login":
		CheckUser()
	case "createNote":
		CreateNote()
	case "queryNote":
		QueryNote()
	case "updateNote":
		UpdateNote()
	case "deleteNote":
		DeleteNote()
	default:
		fmt.Printf("wrong action: %s\n", *action)
	}
}

func CreateUser() {
	req := demoapi.CreateUserRequest{
		Username: "lorain",
		Password: "123456",
	}
	resp, rawResp, err := api_service.CreateUser(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println(string(rawResp.Body()))
}

func CheckUser() {
	req := demoapi.CheckUserRequest{
		Username: "lorain",
		Password: "123456",
	}
	_, rawResp, err := api_service.CheckUser(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	var res map[string]interface{}
	err = json.Unmarshal(rawResp.Body(), &res)
	if err != nil {
		panic(err)
	}
	token = res["token"].(string)
	defaultHeader.Add("Authorization", "Bearer "+token)
	fmt.Printf("token:%s\n", token)
}

var (
	defaultHeader = http.Header{}
	token         string
)

func CreateNote() {
	CheckUser()
	authorizationClient, _ := api_service.NewApiServiceClient("http://127.0.0.1:8080", api_service.WithHeader(defaultHeader))
	req := demoapi.CreateNoteRequest{
		Title:   "test title",
		Content: "test content",
	}
	resp, rawResp, err := authorizationClient.CreateNote(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println(string(rawResp.Body()))
}

func QueryNote() {
	CheckUser()
	authorizationClient, _ := api_service.NewApiServiceClient("http://127.0.0.1:8080", api_service.WithHeader(defaultHeader))
	key := "test"
	req := demoapi.QueryNoteRequest{
		Offset:    0,
		Limit:     20,
		SearchKey: &key,
	}
	_, rawResp, err := authorizationClient.QueryNote(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rawResp.Body()))
}

func UpdateNote() {
	CheckUser()
	authorizationClient, _ := api_service.NewApiServiceClient("http://127.0.0.1:8080", api_service.WithHeader(defaultHeader))
	title := "test"
	content := "test"
	req := demoapi.UpdateNoteRequest{
		Title:   &title,
		Content: &content,
		NoteID:  1,
	}
	resp, rawResp, err := authorizationClient.UpdateNote(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println(string(rawResp.Body()))
}

func DeleteNote() {
	CheckUser()
	authorizationClient, _ := api_service.NewApiServiceClient("http://127.0.0.1:8080", api_service.WithHeader(defaultHeader))
	req := demoapi.DeleteNoteRequest{
		NoteID: 1,
	}
	_, rawResp, err := authorizationClient.DeleteNote(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rawResp.Body()))
}
