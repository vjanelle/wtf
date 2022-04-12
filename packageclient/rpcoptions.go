// generated code; DO NOT EDIT

package packageclient

import (
	"time"

	coreclient "github.com/choria-io/go-choria/client/client"
	rpcclient "github.com/choria-io/go-choria/providers/agent/mcorpc/client"
)

// OptionReset resets the client options to use across requests to an empty list
func (p *PackageClient) OptionReset() *PackageClient {
	p.Lock()
	defer p.Unlock()

	p.clientRPCOpts = []rpcclient.RequestOption{}
	p.ns = p.clientOpts.ns
	p.targets = []string{}
	p.filters = []FilterFunc{
		FilterFunc(coreclient.AgentFilter("package")),
	}

	return p
}

// OptionIdentityFilter adds an identity filter
func (p *PackageClient) OptionIdentityFilter(f ...string) *PackageClient {
	p.Lock()
	defer p.Unlock()

	for _, i := range f {
		if i == "" {
			continue
		}

		p.filters = append(p.filters, FilterFunc(coreclient.IdentityFilter(i)))
	}

	p.ns.Reset()

	return p
}

// OptionClassFilter adds a class filter
func (p *PackageClient) OptionClassFilter(f ...string) *PackageClient {
	p.Lock()
	defer p.Unlock()

	for _, i := range f {
		if i == "" {
			continue
		}

		p.filters = append(p.filters, FilterFunc(coreclient.ClassFilter(i)))
	}

	p.ns.Reset()

	return p
}

// OptionFactFilter adds a fact filter
func (p *PackageClient) OptionFactFilter(f ...string) *PackageClient {
	p.Lock()
	defer p.Unlock()

	for _, i := range f {
		if i == "" {
			continue
		}

		p.filters = append(p.filters, FilterFunc(coreclient.FactFilter(i)))
	}

	p.ns.Reset()

	return p
}

// OptionAgentFilter adds an agent filter
func (p *PackageClient) OptionAgentFilter(a ...string) *PackageClient {
	p.Lock()
	defer p.Unlock()

	for _, f := range a {
		if f == "" {
			continue
		}

		p.filters = append(p.filters, FilterFunc(coreclient.AgentFilter(f)))
	}

	p.ns.Reset()

	return p
}

// OptionCombinedFilter adds a combined filter
func (p *PackageClient) OptionCombinedFilter(f ...string) *PackageClient {
	p.Lock()
	defer p.Unlock()

	for _, i := range f {
		if i == "" {
			continue
		}

		p.filters = append(p.filters, FilterFunc(coreclient.CombinedFilter(i)))
	}

	p.ns.Reset()

	return p
}

// OptionCompoundFilter adds a compound filter
func (p *PackageClient) OptionCompoundFilter(f ...string) *PackageClient {
	p.Lock()
	defer p.Unlock()

	for _, i := range f {
		if i == "" {
			continue
		}

		p.filters = append(p.filters, FilterFunc(coreclient.CompoundFilter(i)))
	}

	p.ns.Reset()

	return p
}

// OptionCollective sets the collective to target
func (p *PackageClient) OptionCollective(c string) *PackageClient {
	p.Lock()
	defer p.Unlock()

	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.Collective(c))
	return p
}

// OptionInBatches performs requests in batches
func (p *PackageClient) OptionInBatches(size int, sleep int) *PackageClient {
	p.Lock()
	defer p.Unlock()

	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.InBatches(size, sleep))
	return p
}

// OptionDiscoveryTimeout configures the request discovery timeout, defaults to configured discovery timeout
func (p *PackageClient) OptionDiscoveryTimeout(t time.Duration) *PackageClient {
	p.Lock()
	defer p.Unlock()

	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.DiscoveryTimeout(t))
	return p
}

// OptionLimitMethod configures the method to use when limiting targets - "random" or "first"
func (p *PackageClient) OptionLimitMethod(m string) *PackageClient {
	p.Lock()
	defer p.Unlock()

	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.LimitMethod(m))
	return p
}

// OptionLimitSize sets limits on the targets, either a number of a percentage like "10%"
func (p *PackageClient) OptionLimitSize(s string) *PackageClient {
	p.Lock()
	defer p.Unlock()

	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.LimitSize(s))
	return p
}

// OptionLimitSeed sets the random seed used to select targets when limiting and limit method is "random"
func (p *PackageClient) OptionLimitSeed(s int64) *PackageClient {
	p.Lock()
	defer p.Unlock()

	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.LimitSeed(s))
	return p
}

// OptionTargets sets specific node targets which would avoid discovery for all action calls until reset
func (p *PackageClient) OptionTargets(t []string) *PackageClient {
	p.Lock()
	defer p.Unlock()

	p.targets = t
	return p
}

// OptionWorkers sets how many worker connections should be started to the broker
func (p *PackageClient) OptionWorkers(w int) *PackageClient {
	p.Lock()
	defer p.Unlock()

	p.workers = w
	return p
}

// OptionExprFilter sets a filter expression that will remove results from the result set
func (p *PackageClient) OptionExprFilter(f string) *PackageClient {
	p.Lock()
	defer p.Unlock()

	p.exprFilter = f
	return p
}

// OptionReplyTo sets a custom reply target
func (p *PackageClient) OptionReplyTo(t string) *PackageClient {
	p.Lock()
	defer p.Unlock()

	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.ReplyTo(t))
	p.noReplies = true
	p.clientOpts.progress = false

	return p
}
