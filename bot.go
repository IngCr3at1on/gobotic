package gobotic

import (
	"context"

	"github.com/ds0nt/gobotic/transports/types"
)

type Bot struct {
	transport Transport
	router    *CommandRouter
}

func NewBot(t Transport, r *CommandRouter) *Bot {
	return &Bot{
		transport: t,
		router:    r,
	}
}

func (c *Bot) Run(ctx context.Context) error {
	err := c.transport.Connect(ctx)
	if err != nil {
		return err
	}
	c.transport.OnMessage(c.OnMessage)
	c.transport.OnError(c.OnError)

	return nil
}
func (c *Bot) OnMessage(msg types.MessageEvent) error {
	if msg.IsCommand {
		err := c.router.Run(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Bot) OnError(err error) {
	c.transport.SendError(err)
}
