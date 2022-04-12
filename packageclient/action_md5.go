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

// Md5Requester performs a RPC request to package#md5
type Md5Requester struct {
	r    *requester
	outc chan *Md5Output
}

// Md5Output is the output from the md5 action
type Md5Output struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// Md5Result is the result from a md5 action
type Md5Result struct {
	ddl        *agent.DDL
	stats      *rpcclient.Stats
	outputs    []*Md5Output
	rpcreplies []*replyfmt.RPCReply
	mu         sync.Mutex
}

func (d *Md5Result) RenderResults(w io.Writer, format RenderFormat, displayMode DisplayMode, verbose bool, silent bool, colorize bool, log Log) error {
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
func (d *Md5Result) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *Md5Output) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *Md5Output) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *Md5Output) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// ParseMd5Output parses the result value from the Md5 action into target
func (d *Md5Output) ParseMd5Output(target interface{}) error {
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
func (d *Md5Requester) Do(ctx context.Context) (*Md5Result, error) {
	dres := &Md5Result{ddl: d.r.client.ddl}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		// filtered by expr filter
		if r == nil {
			return
		}

		output := &Md5Output{
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
func (d *Md5Result) EachOutput(h func(r *Md5Output)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// Exitcode is the value of the exitcode output
//
// Description: The exitcode from the rpm/dpkg command
func (d *Md5Output) Exitcode() interface{} {
	val, ok := d.reply["exitcode"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}

// Output is the value of the output output
//
// Description: md5 of list of packages installed
func (d *Md5Output) Output() interface{} {
	val, ok := d.reply["output"]
	if !ok || val == nil {
		// we have to avoid returning nil.(interface{})
		return nil
	}

	return val

}
