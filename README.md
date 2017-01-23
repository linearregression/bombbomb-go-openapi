# Go API client for bombbomb

We make it easy to build relationships using simple videos.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0.22196
- Package version: 2.0.22196
- Build date: 2017-01-23T19:40:40.423Z
- Build package: class io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./bombbomb"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.bombbomb.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AutomationsApi* | [**GetDripDropStats**](docs/AutomationsApi.md#getdripdropstats) | **Get** /automation/{dripId}/dripdrop/{dripDropId}/stats | Get Automation Email Stats
*AutomationsApi* | [**GetDripStats**](docs/AutomationsApi.md#getdripstats) | **Get** /automation/{id}/stats | Get Automation Stats
*CurriculumApi* | [**GetCurricula**](docs/CurriculumApi.md#getcurricula) | **Get** /curricula/ | Get Curricula
*CurriculumApi* | [**GetUserCurriculumWithProgress**](docs/CurriculumApi.md#getusercurriculumwithprogress) | **Get** /curriculum/getForUserWithProgress | Get Detailed For User
*EmailsApi* | [**CreatePrintingPressEmail**](docs/EmailsApi.md#createprintingpressemail) | **Post** /emails/print | Create an Email with Printing Press
*EmailsApi* | [**GetEmailTracking**](docs/EmailsApi.md#getemailtracking) | **Get** /emails/{emailId}/tracking | Get Email Tracking
*EmailsApi* | [**GetEmailTrackingInteractions**](docs/EmailsApi.md#getemailtrackinginteractions) | **Get** /emails/{emailId}/tracking/interactions | Get Email Tracking Interactions
*EmailsApi* | [**GetHourlyEmailTracking**](docs/EmailsApi.md#gethourlyemailtracking) | **Get** /emails/{emailId}/tracking/hourly | Get Hourly Email Tracking
*PromptsApi* | [**CreatePromptBot**](docs/PromptsApi.md#createpromptbot) | **Post** /prompts/bots | Create a running Prompt Bot for a list
*PromptsApi* | [**CreateVideoEmailPrompt**](docs/PromptsApi.md#createvideoemailprompt) | **Post** /prompt | Prompts user to send a video
*PromptsApi* | [**GetPendingVideoEmailPrompts**](docs/PromptsApi.md#getpendingvideoemailprompts) | **Get** /prompt/pending | List pending prompts
*PromptsApi* | [**GetPromptBots**](docs/PromptsApi.md#getpromptbots) | **Get** /prompts/bots | List Prompt Bots
*PromptsApi* | [**GetPromptCampaigns**](docs/PromptsApi.md#getpromptcampaigns) | **Get** /prompts/campaigns | List Prompt Campaigns
*PromptsApi* | [**GetVideoEmailPrompt**](docs/PromptsApi.md#getvideoemailprompt) | **Get** /prompt/{id} | Gets a prompt
*PromptsApi* | [**GetVideoEmailPrompts**](docs/PromptsApi.md#getvideoemailprompts) | **Get** /prompt/ | List prompts
*PromptsApi* | [**RespondToVideoEmailPrompt**](docs/PromptsApi.md#respondtovideoemailprompt) | **Post** /prompt/{id}/response | Respond to a prompt
*PromptsApi* | [**UpdatePromptBot**](docs/PromptsApi.md#updatepromptbot) | **Put** /prompts/bots/{id} | Update Prompt Bot
*PromptsApi* | [**UpdatePromptCampaign**](docs/PromptsApi.md#updatepromptcampaign) | **Put** /prompts/campaigns/{id} | Update Prompt Campaign
*TeamsApi* | [**AddTeamMember**](docs/TeamsApi.md#addteammember) | **Post** /team/{teamId}/member | Add Member to Team
*TeamsApi* | [**CancelJerichoSend**](docs/TeamsApi.md#canceljerichosend) | **Delete** /team/{teamId}/jericho/{jerichoId} | Cancel a Jericho Send
*TeamsApi* | [**CreateSubteam**](docs/TeamsApi.md#createsubteam) | **Post** /team/{teamId}/subteam | Add a Subteam
*TeamsApi* | [**DeleteSubteam**](docs/TeamsApi.md#deletesubteam) | **Delete** /team/{teamId}/subteam | Delete Subteam
*TeamsApi* | [**GetClientGroupAssets**](docs/TeamsApi.md#getclientgroupassets) | **Get** /team/assets/ | Lists team assets
*TeamsApi* | [**GetJerichoSends**](docs/TeamsApi.md#getjerichosends) | **Get** /team/{teamId}/jericho | List Jericho Sends
*TeamsApi* | [**GetJerichoStats**](docs/TeamsApi.md#getjerichostats) | **Get** /team/{teamId}/jericho/{jerichoId}/performance | Gets Jericho performance statistics
*TeamsApi* | [**GetSubteams**](docs/TeamsApi.md#getsubteams) | **Get** /team/{teamId}/subteam | List Subteams
*TeamsApi* | [**QueueJerichoSend**](docs/TeamsApi.md#queuejerichosend) | **Post** /team/{teamId}/jericho | Creates a Jericho send.
*TeamsApi* | [**RemoveMemberFromTeam**](docs/TeamsApi.md#removememberfromteam) | **Delete** /team/{teamId}/member/{userId} | Remove Member from Team
*TeamsApi* | [**UpdateTeam**](docs/TeamsApi.md#updateteam) | **Post** /team/{teamId} | Update a team
*UtilitiesApi* | [**CreateOAuthClient**](docs/UtilitiesApi.md#createoauthclient) | **Post** /oauthclient | Create an OAuth Client
*UtilitiesApi* | [**DeleteOAuthClient**](docs/UtilitiesApi.md#deleteoauthclient) | **Delete** /oauthclient/{id} | Delete an OAuth Client
*UtilitiesApi* | [**GetOAuthClients**](docs/UtilitiesApi.md#getoauthclients) | **Get** /oauthclient | Lists OAuth Clients
*UtilitiesApi* | [**GetSpec**](docs/UtilitiesApi.md#getspec) | **Get** /spec | Describes this api
*VideosApi* | [**GetVideoRecorder**](docs/VideosApi.md#getvideorecorder) | **Get** /videos/live/getRecorder | Get Live Video Recorder HTML
*VideosApi* | [**MarkLiveRecordingComplete**](docs/VideosApi.md#markliverecordingcomplete) | **Post** /videos/live/markComplete | Completes a live recording
*VideosApi* | [**SignUpload**](docs/VideosApi.md#signupload) | **Post** /video/signedUpload | Generate Signed Url
*WebhooksApi* | [**AddWebHook**](docs/WebhooksApi.md#addwebhook) | **Post** /webhook | Add Webhook
*WebhooksApi* | [**DeleteWebHook**](docs/WebhooksApi.md#deletewebhook) | **Delete** /webhook/{hookId} | Deletes Webhook
*WebhooksApi* | [**GetWebHooks**](docs/WebhooksApi.md#getwebhooks) | **Get** /webhook/ | Lists Webhooks
*WebhooksApi* | [**ListWebHookEvents**](docs/WebhooksApi.md#listwebhookevents) | **Get** /webhook/events | Describe WebHook Events
*WebhooksApi* | [**SendWebhookExample**](docs/WebhooksApi.md#sendwebhookexample) | **Post** /webhook/test | Sends test Webhook


## Documentation For Models

 - [BbWebHook](docs/BbWebHook.md)
 - [Curriculum](docs/Curriculum.md)
 - [CurriculumUserProgress](docs/CurriculumUserProgress.md)
 - [CurriculumWithProgress](docs/CurriculumWithProgress.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse200Items](docs/InlineResponse200Items.md)
 - [JerichoConfiguration](docs/JerichoConfiguration.md)
 - [JerichoPerformance](docs/JerichoPerformance.md)
 - [OAuthClient](docs/OAuthClient.md)
 - [PromptBotBot](docs/PromptBotBot.md)
 - [SignUploadRequest](docs/SignUploadRequest.md)
 - [SignUploadResponse](docs/SignUploadResponse.md)
 - [String](docs/String.md)
 - [TeamPublicRepresentation](docs/TeamPublicRepresentation.md)
 - [VideoEmailPrompt](docs/VideoEmailPrompt.md)
 - [VideoPublicRepresentation](docs/VideoPublicRepresentation.md)
 - [VideoRecorderMethodResponse](docs/VideoRecorderMethodResponse.md)


## Documentation For Authorization


## BBOAuth2

- **Type**: OAuth
- **Flow**: implicit
- **Authorizatoin URL**: https://app.bombbomb.com/auth/authorize
- **Scopes**: 
 - **all:manage**: Manage All
 - **all:read**: Read All
 - **email:manage**: Manage Email
 - **email:read**: Read Email
 - **video:manage**: Manage Video
 - **video:read**: Read Video
 - **contact:manage**: Manage Contact
 - **contact:read**: Read Contact
 - **curriculum:manage**: Manage Curriculum
 - **curriculum:read**: Read Curriculum
 - **automation:manage**: Manage Automation
 - **automation:read**: Read Automation
 - **form:manage**: Manage Form
 - **form:read**: Read Form
 - **team:manage**: Manage Team
 - **team:read**: Read Team
 - **settings:manage**: Manage Settings


## Author



