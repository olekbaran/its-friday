<h1>
  <img src="https://raw.githubusercontent.com/olek-arsee/its-friday/main/images/its-friday.png" width="123px" />
  <p>its-friday</p>
</h1>
<h1>Description</h1>

🐬 Discord BOT built with Go.

Technologies used:

- [discordgo](https://github.com/bwmarrin/discordgo)
- [cron](https://github.com/robfig/cron)

<h1>Features</h1>
<h3>Commands</h3>

- #### `help`

  Sends the embed with listed all available commands.

- #### `author`

  Embed about the developer.

- #### `ping`

  A Simple method to check if we have communication with the BOT.

- #### `pong`

  Like ping but vice versa.

- #### `when-friday`

  Sends how many days are left to the next Friday.

- #### `add-friday`

  Adds channel to the Friday message sending list. You have to put a channel ID after the space!

- #### `delete-friday`

  Deletes channel from the Friday message sending list. You have to put a channel ID after the space!

<h3>Other</h3>

- #### Friday embed
  itsFriday automatically sends embed when Friday starts. The BOT does this to make users smile and remind them that this is the last day of the working week. When Monday comes, the message about Friday is deleted.
