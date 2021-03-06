---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/slack.svg"
brand_color: "#7C2852"
display_name: "Slack"
name: "slack"
description: "Steampipe plugin for querying Slack Conversations, Groups, Users and other resources."
---

# Slack

The Slack plugin is used to query conversations, users and other data.


## Installation

Download and install the latest Slack plugin:

```bash
$ steampipe plugin install slack
Installing plugin slack...
$
```


## Configure API Token

API tokens in Slack are associated with Apps. To use Steampipe, you need to
create an App in Slack with the appropriate permissions.

1. Sign in to the Slack website, and view Your Apps at https://api.slack.com/apps
2. Create New App for your workspace, e.g. `Steampipe CLI`.
3. In "Add features & functionality", choose "Permissions".
4. Grant permissions in a User token scope. This means the Slack App is acting on your behalf. See below for required scopes by table.
5. (Re-)Install your app.
6. Get the user oauth token for your team. It looks like `xoxp-2556146250-EXAMPLE-1646968370949-df954218b5da5b8614c85cc454136b27`

Set Slack API token as an environment variable (Mac, Linux):

```bash
export SLACK_TOKEN="xoxp-2556146250-EXAMPLE-1646968370949-df954218b5da5b8614c85cc454136b27"
```

Run a query:

```bash
$ steampipe query
Welcome to Steampipe v0.0.11
Type ".inspect" for more information.
> select * from slack_connection;
+-------------------------+------+--------+-----------+-----------+---------------+--------+
|             url         | team |  user  |  team_id  |  user_id  | enterprise_id | bot_id |
+-------------------------+------+--------+-----------+-----------+---------------+--------+
| https://acme.slack.com/ | ACME | dwight | T02GD4B7C | U03GC6A7E |               |        |
+-------------------------+------+--------+-----------+-----------+---------------+--------+
```


## Permissions and Scopes

Scopes are used to determine the permissions and access granted to your App in Slack.
Steampipe requires different permissions for each table. We recommend granting all
of the scopes in the table below, but you can restrict them to specific tables if you
prefer.

| Table | Scope Required |
|--|--|
| `slack_access_log` | `admin` ([paid plan required](https://slack.com/help/articles/360002084807-View-Access-Logs-for-your-workspace)) |
| `slack_connection` | _None_ |
| `slack_conversation` | `channels:read`, `groups:read`, `im:read`, `mpim:read` |
| `slack_emoji` | `emoji:read` |
| `slack_group` | `usergroups:read` |
| `slack_search` | `search:read` |
| `slack_user` | `users:read`, `users:read.email` |
