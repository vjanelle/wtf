// generated code; DO NOT EDIT"
//
// Client for Choria RPC Agent 'package' Version 5.1.0 generated using Choria version 0.25.1

package packageclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/choria-io/go-choria/protocol"
	rpcclient "github.com/choria-io/go-choria/providers/agent/mcorpc/client"
	"github.com/choria-io/go-choria/providers/agent/mcorpc/ddl/agent"
	"github.com/choria-io/go-choria/providers/agent/mcorpc/replyfmt"
)

// UninstallRequester performs a RPC request to package#uninstall
type UninstallRequester struct {
	r    *requester
	outc chan *UninstallOutput
}

// UninstallOutput is the output from the uninstall action
type UninstallOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// UninstallResult is the result from a uninstall action
type UninstallResult struct {
	ddl        *agent.DDL
	stats      *rpcclient.Stats
	outputs    []*UninstallOutput
	rpcreplies []*replyfmt.RPCReply
	mu         sync.Mutex
}

func (d *UninstallResult) RenderResults(w io.Writer, format RenderFormat, displayMode DisplayMode, verbose bool, silent bool, colorize bool, log Log) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.stats == nil {
		return fmt.Errorf("result stats is not set, result was not completed")
	}

	results := &replyfmt.RPCResults{
		Agent:   d.stats.Agent(),
		Action:  d.stats.Action(),
		Replies: d.rpcreplies,
		Stats:   d.stats,
	}

	addl, err := d.ddl.ActionInterface(d.stats.Action())
	if err != nil {
		return err
	}

	switch format {
	case JSONFormat:
		return results.RenderJSON(w, addl)
	case TableFormat:
		return results.RenderTable(w, addl)
	case TXTFooter:
		results.RenderTXTFooter(w, verbose)
		return nil
	default:
		return results.RenderTXT(w, addl, verbose, silent, replyfmt.DisplayMode(displayMode), colorize, log)
	}
}

// Stats is the rpc request stats
func (d *UninstallResult) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *UninstallOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *UninstallOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *UninstallOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// ParseUninstallOutput parses the result value from the Uninstall action into target
func (d *UninstallOutput) ParseUninstallOutput(target interface{}) error {
	j, err := d.JSON()
	if err != nil {
		return fmt.Errorf("could not access payload: %s", err)
	}

	err = json.Unmarshal(j, target)
	if err != nil {
		return fmt.Errorf("could not unmarshal JSON payload: %s", err)
	}

	return nil
}

// Do performs the request
func (d *UninstallRequester) Do(ctx context.Context) (*UninstallResult, error) {
	dres := &UninstallResult{ddl: d.r.client.ddl}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		// filtered by expr filter
		if r == nil {
			return
		}

		output := &UninstallOutput{
			reply: make(map[string]interface{}),
			details: &ResultDetails{
				sender:  pr.SenderID(),
				code:    int(r.Statuscode),
				message: r.Statusmsg,
				ts:      pr.Time(),
			},
		}

		err := json.Unmarshal(r.Data, &output.reply)
		if err != nil {
			d.r.client.errorf("Could not decode reply from %s: %s", pr.SenderID(), err)
		}

		// caller wants a channel so we dont return a resulset too (lots of memory etc)
		// this is unused now, no support for setting a channel
		if d.outc != nil {
			d.outc <- output
			return
		}

		// else prepare our result set
		dres.mu.Lock()
		dres.outputs = append(dres.outputs, output)
		dres.rpcreplies = append(dres.rpcreplies, &replyfmt.RPCReply{
			Sender:   pr.SenderID(),
			RPCReply: r,
		})
		dres.mu.Unlock()
	}

	res, err := d.r.do(ctx, handler)
	if err != nil {
		return nil, err
	}

	dres.stats = res

	return dres, nil
}

// EachOutput iterates over all results received
func (d *UninstallResult) EachOutput(h func(r *UninstallOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// Arch is the value of the arch output
//
// Description: Package architecture
func (d *UninstallOutput) Arch() interface{} {
	val, ok := d.reply["arch"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}

// Ensure is the value of the ensure output
//
// Description: Full package version
func (d *UninstallOutput) Ensure() interface{} {
	val, ok := d.reply["ensure"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}

// Epoch is the value of the epoch output
//
// Description: Package epoch number
func (d *UninstallOutput) Epoch() interface{} {
	val, ok := d.reply["epoch"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}

// Name is the value of the name output
//
// Description: Package name
func (d *UninstallOutput) Name() interface{} {
	val, ok := d.reply["name"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}

// Output is the value of the output output
//
// Description: Output from the package manager
func (d *UninstallOutput) Output() interface{} {
	val, ok := d.reply["output"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}

// Provider is the value of the provider output
//
// Description: Provider used to retrieve information
func (d *UninstallOutput) Provider() interface{} {
	val, ok := d.reply["provider"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}

// Release is the value of the release output
//
// Description: Package release number
func (d *UninstallOutput) Release() interface{} {
	val, ok := d.reply["release"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}

// Version is the value of the version output
//
// Description: Version number
func (d *UninstallOutput) Version() interface{} {
	val, ok := d.reply["version"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}