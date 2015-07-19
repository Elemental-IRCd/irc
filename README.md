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

```go
var ErrDisconnected = errors.New("Disconnect Called")
```

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


#### func  New

```go
func New(nick, user string) *Connection
```
Create a connection with the (publicly visible) nickname and username. The
nickname is later used to address the user. Returns nil if nick or user are
empty.

#### func (*Connection) Action

```go
func (irc *Connection) Action(target, message string)
```
Send (action) message to a target (channel or nickname). No clear RFC on this
one...

#### func (*Connection) Actionf

```go
func (irc *Connection) Actionf(target, format string, a ...interface{})
```
Send formatted (action) message to a target (channel or nickname).

#### func (*Connection) AddCallback

```go
func (irc *Connection) AddCallback(eventcode string, callback func(*Event)) string
```
Register a callback to a connection and event code. A callback is a function
which takes only an Event pointer as parameter. Valid event codes are all
IRC/CTCP commands and error/response codes. This function returns the ID of the
registered callback for later management.

#### func (*Connection) ClearCallback

```go
func (irc *Connection) ClearCallback(eventcode string) bool
```
Remove all callbacks from a given event code. It returns true if given event
code is found and cleared.

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
Returns true if the connection is connected to an IRC server.

#### func (*Connection) Disconnect

```go
func (irc *Connection) Disconnect()
```
A disconnect sends all buffered messages (if possible), stops all goroutines and
then closes the socket.

#### func (*Connection) ErrorChan

```go
func (irc *Connection) ErrorChan() chan error
```

#### func (*Connection) GetNick

```go
func (irc *Connection) GetNick() string
```
Determine nick currently used with the connection.

#### func (*Connection) Join

```go
func (irc *Connection) Join(channel string)
```
Use the connection to join a given channel. RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.2.1

#### func (*Connection) Loop

```go
func (irc *Connection) Loop()
```
Main loop to control the connection.

#### func (*Connection) Mode

```go
func (irc *Connection) Mode(target string, modestring ...string)
```
Set different modes for a target (channel or nickname). RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.2.3

#### func (*Connection) Nick

```go
func (irc *Connection) Nick(n string)
```
Set (new) nickname. RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.1.2

#### func (*Connection) Notice

```go
func (irc *Connection) Notice(target, message string)
```
Send a notification to a nickname. This is similar to Privmsg but must not
receive replies. RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.4.2

#### func (*Connection) Noticef

```go
func (irc *Connection) Noticef(target, format string, a ...interface{})
```
Send a formated notification to a nickname. RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.4.2

#### func (*Connection) Part

```go
func (irc *Connection) Part(channel string)
```
Leave a given channel. RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.2.2

#### func (*Connection) Privmsg

```go
func (irc *Connection) Privmsg(target, message string)
```
Send (private) message to a target (channel or nickname). RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.4.1

#### func (*Connection) Privmsgf

```go
func (irc *Connection) Privmsgf(target, format string, a ...interface{})
```
Send formated string to specified target (channel or nickname).

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
Remove callback i (ID) from the given event code. This functions returns true
upon success, false if any error occurs.

#### func (*Connection) ReplaceCallback

```go
func (irc *Connection) ReplaceCallback(eventcode string, i string, callback func(*Event))
```
Replace callback i (ID) associated with a given event code with a new callback
function.

#### func (*Connection) RunCallbacks

```go
func (irc *Connection) RunCallbacks(event *Event)
```
Execute all callbacks associated with a given event.

#### func (*Connection) SendRaw

```go
func (irc *Connection) SendRaw(message string)
```
Send raw string.

#### func (*Connection) SendRawf

```go
func (irc *Connection) SendRawf(format string, a ...interface{})
```
Send raw formated string.

#### func (*Connection) Who

```go
func (irc *Connection) Who(nick string)
```
Query information about a given nickname in the server. RFC 1459 details:
https://tools.ietf.org/html/rfc1459#section-4.5.1

#### func (*Connection) Whois

```go
func (irc *Connection) Whois(nick string)
```
Query information about a particular nickname. RFC 1459:
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

A struct to represent an event.

#### func (*Event) Message

```go
func (e *Event) Message() string
```
Retrieve the last message from Event arguments. This function leaves the
arguments untouched and returns an empty string if there are none.
