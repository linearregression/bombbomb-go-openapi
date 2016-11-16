/* 
 * BombBomb
 *
 * We make it easy to build relationships using simple videos.
 *
 * OpenAPI spec version: 2.0.21454
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

// Reports on the aggregate performance of a Jericho send
type JerichoPerformance struct {

	// The number of emails sent
	Sent int32 `json:"sent,omitempty"`

	// The unique number of people that viewed the email
	UniqueViews int32 `json:"uniqueViews,omitempty"`

	// The unique number of people that visited the landing page
	UniqueLandingPageViews int32 `json:"uniqueLandingPageViews,omitempty"`

	// The total number of people that visisted the landing page
	LandingPageViews int32 `json:"landingPageViews,omitempty"`

	// The number of emails delivered, likely less than sent due to bounces, and other common delivery issues
	Delivered int32 `json:"delivered,omitempty"`

	// The number of emails that bounced as undeliverable
	Bounce int32 `json:"bounce,omitempty"`

	// The total number of times the emails were viewed
	Open int32 `json:"open,omitempty"`

	// The total number of times links in the emails were clicked
	Click int32 `json:"click,omitempty"`

	// The total number of times videos in the emails were played
	VideoPlay int32 `json:"videoPlay,omitempty"`

	// The number of recipients that marked the message as abusive
	AbuseComplaints int32 `json:"abuseComplaints,omitempty"`

	// The total number of contacts submitted to be sent, may be more than was sent to
	Contacts int32 `json:"contacts,omitempty"`
}
