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

// AptCheckupdatesRequester performs a RPC request to package#apt_checkupdates
type AptCheckupdatesRequester struct {
	r    *requester
	outc chan *AptCheckupdatesOutput
}

// AptCheckupdatesOutput is the output from the apt_checkupdates action
type AptCheckupdatesOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// AptCheckupdatesResult is the result from a apt_checkupdates action
type AptCheckupdatesResult struct {
	ddl        *agent.DDL
	stats      *rpcclient.Stats
	outputs    []*AptCheckupdatesOutput
	rpcreplies []*replyfmt.RPCReply
	mu         sync.Mutex
}

func (d *AptCheckupdatesResult) RenderResults(w io.Writer, format RenderFormat, displayMode DisplayMode, verbose bool, silent bool, colorize bool, log Log) error {
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
func (d *AptCheckupdatesResult) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *AptCheckupdatesOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *AptCheckupdatesOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *AptCheckupdatesOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// ParseAptCheckupdatesOutput parses the result value from the AptCheckupdates action into target
func (d *AptCheckupdatesOutput) ParseAptCheckupdatesOutput(target interface{}) error {
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
func (d *AptCheckupdatesRequester) Do(ctx context.Context) (*AptCheckupdatesResult, error) {
	dres := &AptCheckupdatesResult{ddl: d.r.client.ddl}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		// filtered by expr filter
		if r == nil {
			return
		}

		output := &AptCheckupdatesOutput{
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
func (d *AptCheckupdatesResult) EachOutput(h func(r *AptCheckupdatesOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// Exitcode is the value of the exitcode output
//
// Description: The exitcode from the apt command
func (d *AptCheckupdatesOutput) Exitcode() interface{} {
	val, ok := d.reply["exitcode"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}

// OutdatedPackages is the value of the outdated_packages output
//
// Description: Outdated packages
func (d *AptCheckupdatesOutput) OutdatedPackages() interface{} {
	val, ok := d.reply["outdated_packages"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}

// Output is the value of the output output
//
// Description: Output from APT
func (d *AptCheckupdatesOutput) Output() interface{} {
	val, ok := d.reply["output"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}
