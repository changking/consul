package command

import (
	"strings"
	"testing"

	"github.com/hashicorp/consul/command/base"
	"github.com/mitchellh/cli"
)

func TestReloadCommand_implements(t *testing.T) {
	var _ cli.Command = &ReloadCommand{}
}

func TestReloadCommandRun(t *testing.T) {
	a1 := testAgentWithConfig(t, nil)
	defer a1.Shutdown()

	// Setup a dummy response to errCh to simulate a successful reload
	go func() {
		errCh := <-a1.agent.ReloadCh()
		errCh <- nil
	}()

	ui := new(cli.MockUi)
	c := &ReloadCommand{
		Command: base.Command{
			UI:    ui,
			Flags: base.FlagSetClientHTTP,
		},
	}
	args := []string{"-http-addr=" + a1.httpAddr}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), "reload triggered") {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}
