/* 
 * BombBomb
 *
 * We make it easy to build relationships using simple videos.
 *
 * OpenAPI spec version: 2.0.21432
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bombbomb

import (
	"time"
)

type JerichoConfiguration struct {

	Id string `json:"id,omitempty"`

	ClientGroupId string `json:"clientGroupId,omitempty"`

	// When the email should be sent.
	SendDate time.Time `json:"sendDate,omitempty"`

	// Determines whether this is a static or prompted send.
	IsPrompt bool `json:"isPrompt,omitempty"`

	// Static send: The Email to send on behalf of the group members.
	EmailId string `json:"emailId,omitempty"`

	// Video Prompt: The subject line prompting the user to record a video.
	PromptSubject string `json:"promptSubject,omitempty"`

	// Video Prompt: The HTML body of the email prompting the user to record a video.
	PromptBody string `json:"promptBody,omitempty"`

	// Video Prompt: The subject line of the final sent email
	EmailSubject string `json:"emailSubject,omitempty"`

	// Video Prompt: The HTML body of the final sent email.
	EmailBody string `json:"emailBody,omitempty"`

	// Video Prompt: Whether to send the final email if no video was recorded.
	SendWithoutVideo bool `json:"sendWithoutVideo,omitempty"`

	// The state of the send.
	Status string `json:"status,omitempty"`
}
