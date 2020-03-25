/*
 *    Copyright 2020 opensourceai
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

/*
 * @Package dto
 * @Author Quan Chen
 * @Date 2020/3/19
 * @Description
 *
 */
package dto

type CommentUpdate struct {
	ID             int    `json:"id"`              // 评论ID
	CommentContent string `json:"comment_content"` // 评论的内容
}
type Comment struct {
	ID             int    `json:"id"`              // 评论ID
	CommentContent string `json:"comment_content"` // 评论内容
	FromUser       User   `json:"from_user"`       // 评论者
	ToUser         User   `json:"to_user"`         // 被评论者
}