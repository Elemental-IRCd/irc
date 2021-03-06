# irc
--
    import "github.com/Elemental-IRCd/irc"

Package irc implements a basic IRC client for Go programs, following RFC 1459.

This does not keep track of channel state.

Usage example:

    package main

    import (
    	"github.com/Elemental-IRCd/irc"
    )

    func main() {
    	bot := irc.New("MyBot", "foosmith")
    	bot.UseTLS = true

    	bot.AddCallback("001", func(*irc.Event) {
    		bot.Join("#irc")
    	})

    	err := bot.Connect("irc.ponychat.net:6697")
    	if err != nil {
    		panic(err)
    	}

    	bot.Loop()
    }

## Usage

```go
const (
	VERSION = "Elemental-IRCd irc package 0.1"
)
```
The CTCP-VERSION reply that clients using this package will return.

```go
var ErrDisconnected = errors.New("Disconnect Called")
```
This is thrown when another goroutine calls Disconnect.

#### type Connection

```go
type Connection struct {
	sync.WaitGroup
	Debug     bool
	Error     chan error
	Password  string
	UseTLS    bool
	TLSConfig *tls.Config
	Version   string
	Timeout   time.Duration
	PingFreq  time.Duration
	KeepAlive time.Duration
	Server    string

	VerboseCallbackHandler bool
	Log                    *log.Logger
}
```

Connection is a single IRC connection to a remote server.

#### func  New

```go
func New(nick, user string) *Connection
```
New creates a connection with the (publicly visible) nickname and username. The
nickname is later used to address the user. Returns nil if nick or user are
empty.

#### func (*Connection) Action

```go
func (irc *Connection) Action(target, message string)
```
Action sends a CTCP-ACTION (/me) message to a target (channel or nickname). No
clear RFC on this one...

#### func (*Connection) Actionf

```go
func (irc *Connection) Actionf(target, format string, a ...interface{})
```
Actionf sends a CTCP-ACTION (/me) to a target (channel or nickname).

#### func (*Connection) AddCallback

```go
func (irc *Connection) AddCallback(eventcode string, callback func(*Event)) string
```
AddCallback registers a callback to a connection and event code. A callback is a
function which takes only an Event pointer as parameter. Valid event codes are
all IRC/CTCP commands and error/response codes. This function returns the ID of
the registered callback for later management.

#### func (*Connection) ClearCallback

```go
func (irc *Connection) ClearCallback(eventcode string) bool
```
ClearCallback removes all callbacks from a given event code. It returns true if
given event code is found and cleared.

#### func (*Connection) Connect

```go
func (irc *Connection) Connect(server string) error
```
Connect to a given server using the current connection configuration. This
function also takes care of identification if a password is provided. RFC 1459
details: https://tools.ietf.org/html/rfc1459#section-4.1

#### func (*Connection) Connected

```go
func (irc *Connection) Connected() bool
```
Connected returns true if the connection is connected to an IRC server.

#### func (*Connection) Disconnect

```go
func (irc *Connection) Disconnect()
```
Disconnect sends all buffered messages (if possible), stops all goroutines and
then closes the socket.

#### func (*Connection) ErrorChan

```go
func (irc *Connection) ErrorChan() chan error
```
ErrorChan returns the connections error channel.

#### func (*Connection) GetNick

```go
func (irc *Connection) GetNick() string
```
GetNick returns the nickname in use by the client.

#### func (*Connection) Join

```go
func (irc *Connection) Join(channel string)
```
Join uses the connection to join a given channel. RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.2.1

#### func (*Connection) Loop

```go
func (irc *Connection) Loop()
```
Loop is the main loop to control the connection.

#### func (*Connection) Mode

```go
func (irc *Connection) Mode(target string, modestring ...string)
```
Mode sets different modes for a target (channel or nickname). RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.2.3

#### func (*Connection) Nick

```go
func (irc *Connection) Nick(n string)
```
Nick changes the client nickname to the given value. This may fail, causing the
server to return ERR_NICKNAMEINUSE or ERR_ERRONEUSNICKNAME. RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.1.2

#### func (*Connection) Notice

```go
func (irc *Connection) Notice(target, message string)
```
Notice send a notification to a nickname or channel. This is similar to Privmsg
but must not receive replies. RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.4.2

#### func (*Connection) Noticef

```go
func (irc *Connection) Noticef(target, format string, a ...interface{})
```
Noticef sends a formated notification to a nickname or channel. RFC 1459
details: https://tools.ietf.org/html/rfc1459#section-4.4.2

#### func (*Connection) Part

```go
func (irc *Connection) Part(channel string)
```
Part leaves a given channel. RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.2.2

#### func (*Connection) Privmsg

```go
func (irc *Connection) Privmsg(target, message string)
```
Privmsg sends a message to a target (channel or nickname). RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.4.1

#### func (*Connection) Privmsgf

```go
func (irc *Connection) Privmsgf(target, format string, a ...interface{})
```
Privmsgf sends a formatted message to a specified target (channel or nickname).

#### func (*Connection) Quit

```go
func (irc *Connection) Quit()
```
Quit the current connection and disconnect from the server RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.1.6

#### func (*Connection) Reconnect

```go
func (irc *Connection) Reconnect() error
```
Reconnect to a server using the current connection.

#### func (*Connection) RemoveCallback

```go
func (irc *Connection) RemoveCallback(eventcode string, i string) bool
```
RemoveCallback removes callback i (ID) from the given event code. This functions
returns true upon success, false if any error occurs.

#### func (*Connection) ReplaceCallback

```go
func (irc *Connection) ReplaceCallback(eventcode string, i string, callback func(*Event))
```
ReplaceCallback replaces callback i (ID) associated with a given event code with
a new callback function.

#### func (*Connection) RunCallbacks

```go
func (irc *Connection) RunCallbacks(event *Event)
```
RunCallbacks executes all callbacks associated with a given event.

#### func (*Connection) SendRaw

```go
func (irc *Connection) SendRaw(message string)
```
SendRaw sends a raw message across the wire.

#### func (*Connection) SendRawf

```go
func (irc *Connection) SendRawf(format string, a ...interface{})
```
SendRawf sends a formatted raw message across the wire.

#### func (*Connection) Who

```go
func (irc *Connection) Who(target string)
```
Who fetches detailed information about a given target (nick or channel). RFC
1459 details: https://tools.ietf.org/html/rfc1459#section-4.5.1

#### func (*Connection) Whois

```go
func (irc *Connection) Whois(nick string)
```
Whois fetches information about a given client. RFC 1459:
https://tools.ietf.org/html/rfc1459#section-4.5.2

#### type Event

```go
type Event struct {
	Code       string
	Raw        string
	Nick       string //<nick>
	Host       string //<nick>!<usr>@<host>
	Source     string //<host>
	User       string //<usr>
	Arguments  []string
	Connection *Connection
}
```

Event is a struct to represent an event.

#### func (*Event) Message

```go
func (e *Event) Message() string
```
Message retrieves the last message from Event arguments. This function leaves
the arguments untouched and returns an empty string if there are none.
