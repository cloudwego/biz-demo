// Copyright 2023 CloudWeGo Authors
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

// Code generated by hertz generator. DO NOT EDIT.

package douyinapi

import (
	douyinapi "github.com/cloudwego/biz-demo/sample_douyin/cmd/api/hertz_handler/douyinapi"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_comment := _douyin.Group("/comment", _commentMw()...)
			{
				_action := _comment.Group("/action", _actionMw()...)
				_action.POST("/", append(_comment_ctionMw(), douyinapi.CommentAction)...)
			}
			{
				_list := _comment.Group("/list", _listMw()...)
				_list.GET("/", append(_commentlistMw(), douyinapi.CommentList)...)
			}
		}
		{
			_favorite := _douyin.Group("/favorite", _favoriteMw()...)
			{
				_action0 := _favorite.Group("/action", _action0Mw()...)
				_action0.POST("/", append(_favorite_ctionMw(), douyinapi.FavoriteAction)...)
			}
			{
				_list0 := _favorite.Group("/list", _list0Mw()...)
				_list0.GET("/", append(_getfavoritelistMw(), douyinapi.GetFavoriteList)...)
			}
		}
		{
			_feed := _douyin.Group("/feed", _feedMw()...)
			_feed.GET("/", append(_getfeedMw(), douyinapi.GetFeed)...)
		}
		{
			_message := _douyin.Group("/message", _messageMw()...)
			{
				_action1 := _message.Group("/action", _action1Mw()...)
				_action1.POST("/", append(_message_ctionMw(), douyinapi.MessageAction)...)
			}
			{
				_chat := _message.Group("/chat", _chatMw()...)
				_chat.GET("/", append(_messagechatMw(), douyinapi.MessageChat)...)
			}
		}
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_action2 := _publish.Group("/action", _action2Mw()...)
				_action2.POST("/", append(_publishvideoMw(), douyinapi.PublishVideo)...)
			}
			{
				_list1 := _publish.Group("/list", _list1Mw()...)
				_list1.GET("/", append(_getpublishlistMw(), douyinapi.GetPublishList)...)
			}
		}
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			{
				_action3 := _relation.Group("/action", _action3Mw()...)
				_action3.POST("/", append(_relation_ctionMw(), douyinapi.RelationAction)...)
			}
			{
				_follow := _relation.Group("/follow", _followMw()...)
				{
					_list2 := _follow.Group("/list", _list2Mw()...)
					_list2.GET("/", append(_followlistMw(), douyinapi.FollowList)...)
				}
			}
			{
				_follower := _relation.Group("/follower", _followerMw()...)
				{
					_list3 := _follower.Group("/list", _list3Mw()...)
					_list3.GET("/", append(_followerlistMw(), douyinapi.FollowerList)...)
				}
			}
			{
				_friend := _relation.Group("/friend", _friendMw()...)
				{
					_list4 := _friend.Group("/list", _list4Mw()...)
					_list4.GET("/", append(_friendlistMw(), douyinapi.FriendList)...)
				}
			}
		}
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.GET("/", append(_getuserMw(), douyinapi.GetUser)...)
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_checkuserMw(), douyinapi.CheckUser)...)
			}
			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_registuserMw(), douyinapi.RegistUser)...)
			}
		}
	}
}