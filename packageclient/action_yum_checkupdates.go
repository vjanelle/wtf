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

// YumCheckupdatesRequester performs a RPC request to package#yum_checkupdates
type YumCheckupdatesRequester struct {
	r    *requester
	outc chan *YumCheckupdatesOutput
}

// YumCheckupdatesOutput is the output from the yum_checkupdates action
type YumCheckupdatesOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// YumCheckupdatesResult is the result from a yum_checkupdates action
type YumCheckupdatesResult struct {
	ddl        *agent.DDL
	stats      *rpcclient.Stats
	outputs    []*YumCheckupdatesOutput
	rpcreplies []*replyfmt.RPCReply
	mu         sync.Mutex
}

func (d *YumCheckupdatesResult) RenderResults(w io.Writer, format RenderFormat, displayMode DisplayMode, verbose bool, silent bool, colorize bool, log Log) error {
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
func (d *YumCheckupdatesResult) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *YumCheckupdatesOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *YumCheckupdatesOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *YumCheckupdatesOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// ParseYumCheckupdatesOutput parses the result value from the YumCheckupdates action into target
func (d *YumCheckupdatesOutput) ParseYumCheckupdatesOutput(target interface{}) error {
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
func (d *YumCheckupdatesRequester) Do(ctx context.Context) (*YumCheckupdatesResult, error) {
	dres := &YumCheckupdatesResult{ddl: d.r.client.ddl}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		// filtered by expr filter
		if r == nil {
			return
		}

		output := &YumCheckupdatesOutput{
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
func (d *YumCheckupdatesResult) EachOutput(h func(r *YumCheckupdatesOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// Exitcode is the value of the exitcode output
//
// Description: The exitcode from the yum command
func (d *YumCheckupdatesOutput) Exitcode() interface{} {
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
func (d *YumCheckupdatesOutput) OutdatedPackages() interface{} {
	val, ok := d.reply["outdated_packages"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}

// Output is the value of the output output
//
// Description: Output from YUM
func (d *YumCheckupdatesOutput) Output() interface{} {
	val, ok := d.reply["output"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}
