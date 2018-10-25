# Discord Bot

Simple and structure boilerplate to develop bots with Golang for Discord.

# Guides

## Create new command

To develop a new bot command start by creating a file under `commands` folder. In this example I will create a file called, `ping.go`. Then we insert the default structure of a command.

```go
package commands

// NewCommandName is a default function
func NewCommandName(dcs DiscordSession, args []string) {
	_, _ = dcs.Session.ChannelMessageSend(dcs.Msg.ChannelID, "Message")
}
```

This takes two attributes. The first called `dcs` which is a DiscordSession type attribute and the second the `args` which is a slice of strings. For this example we'll create a simple Ping Pong example. On que `ping.go` file insert the following code:

```go
package commands

// Ping command responds with a "pong" text
func Ping(dcs DiscordSession, args []string) {
	_, _ = dcs.Session.ChannelMessageSend(dcs.Msg.ChannelID, "pong!")
}
```

Now what we need to do, is to declare this new command on the message handler function. Enter the `bot` folder and inside on the `bot.go` file we'll find the function `messageHandler`. Bellow the following line:

```go
handler.Handle("help", dcs, commands.Help)
```

add this:

```go
handler.Handle("ping", dcs, commands.Ping)
```

The first argument is a string to the command name, for example this will be called like this: `!prefix ping`. `dcs` is the DiscordSession struct we mentioned before and `commands.Ping` is the function we created on the commands package inside `ping.go`.

## DiscordSession

DiscordSession is a new type struct to make it easier to pass `bwmarrin/discordgo` library variables. The structure is pretty simple and looks like this:

```go
// DiscordSession type eases the construction of a command function
type DiscordSession struct {
	Msg     *discordgo.MessageCreate
	Session *discordgo.Session
}
```

## discordgo

To get more information on the core library, check this link [http://bwmarrin.github.io/discordgo/](http://bwmarrin.github.io/discordgo/).
