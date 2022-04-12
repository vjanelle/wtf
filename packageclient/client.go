// generated code; DO NOT EDIT

package packageclient

import (
	"fmt"
	"sync"
	"time"

	"context"

	coreclient "github.com/choria-io/go-choria/client/client"
	"github.com/choria-io/go-choria/config"
	"github.com/choria-io/go-choria/inter"
	"github.com/choria-io/go-choria/protocol"
	rpcclient "github.com/choria-io/go-choria/providers/agent/mcorpc/client"
	"github.com/choria-io/go-choria/providers/agent/mcorpc/ddl/agent"
)

// Stats are the statistics for a request
type Stats interface {
	Agent() string
	Action() string
	All() bool
	NoResponseFrom() []string
	UnexpectedResponseFrom() []string
	DiscoveredCount() int
	DiscoveredNodes() *[]string
	FailCount() int
	OKCount() int
	ResponsesCount() int
	PublishDuration() (time.Duration, error)
	RequestDuration() (time.Duration, error)
	DiscoveryDuration() (time.Duration, error)
	OverrideDiscoveryTime(start time.Time, end time.Time)
	UniqueRequestID() string
}

// NodeSource discovers nodes
type NodeSource interface {
	Reset()
	Discover(ctx context.Context, fw inter.Framework, filters []FilterFunc) ([]string, error)
}

// FilterFunc can generate a Choria filter
type FilterFunc func(f *protocol.Filter) error

// RenderFormat is the format used by the RenderResults helper
type RenderFormat int

const (
	// JSONFormat renders the results as a JSON document
	JSONFormat RenderFormat = iota

	// TextFormat renders the results as a Choria typical result set in line with choria req output
	TextFormat

	// TableFormat renders all successful responses in a table
	TableFormat

	// TXTFooter renders only the request summary statistics
	TXTFooter
)

// DisplayMode overrides the DDL display hints
type DisplayMode uint8

const (
	// DisplayDDL shows results based on the configuration in the DDL file
	DisplayDDL = DisplayMode(iota)
	// DisplayOK shows only passing results
	DisplayOK
	// DisplayFailed shows only failed results
	DisplayFailed
	// DisplayAll shows all results
	DisplayAll
	// DisplayNone shows no results
	DisplayNone
)

type Log interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

// PackageClient to the package agent
type PackageClient struct {
	fw            inter.Framework
	cfg           *config.Config
	ddl           *agent.DDL
	ns            NodeSource
	clientOpts    *initOptions
	clientRPCOpts []rpcclient.RequestOption
	filters       []FilterFunc
	targets       []string
	workers       int
	exprFilter    string
	noReplies     bool

	sync.Mutex
}

// Metadata is the agent metadata
type Metadata struct {
	License     string `json:"license"`
	Author      string `json:"author"`
	Timeout     int    `json:"timeout"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

// Must create a new client and panics on error
func Must(fw inter.Framework, opts ...InitializationOption) (client *PackageClient) {
	c, err := New(fw, opts...)
	if err != nil {
		panic(err)
	}

	return c
}

// New creates a new client to the package agent
func New(fw inter.Framework, opts ...InitializationOption) (client *PackageClient, err error) {
	c := &PackageClient{
		fw:            fw,
		ddl:           &agent.DDL{},
		clientRPCOpts: []rpcclient.RequestOption{},
		filters: []FilterFunc{
			FilterFunc(coreclient.AgentFilter("package")),
		},
		clientOpts: &initOptions{
			cfgFile: coreclient.UserConfig(),
		},
		targets: []string{},
	}

	for _, opt := range opts {
		opt(c.clientOpts)
	}

	c.cfg = c.fw.Configuration()

	if c.clientOpts.dt > 0 {
		c.cfg.DiscoveryTimeout = int(c.clientOpts.dt.Seconds())
	}

	if c.clientOpts.ns == nil {
		switch c.cfg.DefaultDiscoveryMethod {
		case "choria":
			c.clientOpts.ns = &PuppetDBNS{}
		default:
			c.clientOpts.ns = &BroadcastNS{}
		}
	}
	c.ns = c.clientOpts.ns

	if c.clientOpts.logger == nil {
		c.clientOpts.logger = c.fw.Logger("package")
	} else {
		c.fw.SetLogger(c.clientOpts.logger.Logger)
	}

	c.ddl, err = DDL()
	if err != nil {
		return nil, fmt.Errorf("could not parse embedded DDL: %s", err)
	}

	return c, nil
}

// AgentMetadata is the agent metadata this client supports
func (p *PackageClient) AgentMetadata() *Metadata {
	return &Metadata{
		License:     p.ddl.Metadata.License,
		Author:      p.ddl.Metadata.Author,
		Timeout:     p.ddl.Metadata.Timeout,
		Name:        p.ddl.Metadata.Name,
		Version:     p.ddl.Metadata.Version,
		URL:         p.ddl.Metadata.URL,
		Description: p.ddl.Metadata.Description,
	}
}

// DiscoverNodes performs a discovery using the configured filter and node source
func (p *PackageClient) DiscoverNodes(ctx context.Context) (nodes []string, err error) {
	p.Lock()
	defer p.Unlock()

	return p.ns.Discover(ctx, p.fw, p.filters)
}

// AptCheckupdates performs the apt_checkupdates action
//
// Description: Check for APT updates
func (p *PackageClient) AptCheckupdates() *AptCheckupdatesRequester {
	d := &AptCheckupdatesRequester{
		outc: nil,
		r: &requester{
			args:   map[string]interface{}{},
			action: "apt_checkupdates",
			client: p,
		},
	}

	action, _ := p.ddl.ActionInterface(d.r.action)
	action.SetDefaults(d.r.args)

	return d
}

// AptUpdate performs the apt_update action
//
// Description: Update the apt cache
func (p *PackageClient) AptUpdate() *AptUpdateRequester {
	d := &AptUpdateRequester{
		outc: nil,
		r: &requester{
			args:   map[string]interface{}{},
			action: "apt_update",
			client: p,
		},
	}

	action, _ := p.ddl.ActionInterface(d.r.action)
	action.SetDefaults(d.r.args)

	return d
}

// Checkupdates performs the checkupdates action
//
// Description: Check for updates
func (p *PackageClient) Checkupdates() *CheckupdatesRequester {
	d := &CheckupdatesRequester{
		outc: nil,
		r: &requester{
			args:   map[string]interface{}{},
			action: "checkupdates",
			client: p,
		},
	}

	action, _ := p.ddl.ActionInterface(d.r.action)
	action.SetDefaults(d.r.args)

	return d
}

// Count performs the count action
//
// Description: Get number of packages installed
func (p *PackageClient) Count() *CountRequester {
	d := &CountRequester{
		outc: nil,
		r: &requester{
			args:   map[string]interface{}{},
			action: "count",
			client: p,
		},
	}

	action, _ := p.ddl.ActionInterface(d.r.action)
	action.SetDefaults(d.r.args)

	return d
}

// Install performs the install action
//
// Description: Install a package
//
// Required Inputs:
//    - package (string) - Package to install
//
// Optional Inputs:
//    - version (string) - Version of package to install
func (p *PackageClient) Install(inputPackage string) *InstallRequester {
	d := &InstallRequester{
		outc: nil,
		r: &requester{
			args: map[string]interface{}{
				"package": inputPackage,
			},
			action: "install",
			client: p,
		},
	}

	action, _ := p.ddl.ActionInterface(d.r.action)
	action.SetDefaults(d.r.args)

	return d
}

// Md5 performs the md5 action
//
// Description: Get md5 digest of list of packages installed
func (p *PackageClient) Md5() *Md5Requester {
	d := &Md5Requester{
		outc: nil,
		r: &requester{
			args:   map[string]interface{}{},
			action: "md5",
			client: p,
		},
	}

	action, _ := p.ddl.ActionInterface(d.r.action)
	action.SetDefaults(d.r.args)

	return d
}

// Purge performs the purge action
//
// Description: Purge a package
//
// Required Inputs:
//    - package (string) - Package to purge
func (p *PackageClient) Purge(inputPackage string) *PurgeRequester {
	d := &PurgeRequester{
		outc: nil,
		r: &requester{
			args: map[string]interface{}{
				"package": inputPackage,
			},
			action: "purge",
			client: p,
		},
	}

	action, _ := p.ddl.ActionInterface(d.r.action)
	action.SetDefaults(d.r.args)

	return d
}

// Status performs the status action
//
// Description: Get the status of a package
//
// Required Inputs:
//    - package (string) - Package to retrieve the status of
func (p *PackageClient) Status(inputPackage string) *StatusRequester {
	d := &StatusRequester{
		outc: nil,
		r: &requester{
			args: map[string]interface{}{
				"package": inputPackage,
			},
			action: "status",
			client: p,
		},
	}

	action, _ := p.ddl.ActionInterface(d.r.action)
	action.SetDefaults(d.r.args)

	return d
}

// Uninstall performs the uninstall action
//
// Description: Uninstall a package
//
// Required Inputs:
//    - package (string) - Package to uninstall
func (p *PackageClient) Uninstall(inputPackage string) *UninstallRequester {
	d := &UninstallRequester{
		outc: nil,
		r: &requester{
			args: map[string]interface{}{
				"package": inputPackage,
			},
			action: "uninstall",
			client: p,
		},
	}

	action, _ := p.ddl.ActionInterface(d.r.action)
	action.SetDefaults(d.r.args)

	return d
}

// Update performs the update action
//
// Description: Update a package
//
// Required Inputs:
//    - package (string) - Package to update
func (p *PackageClient) Update(inputPackage string) *UpdateRequester {
	d := &UpdateRequester{
		outc: nil,
		r: &requester{
			args: map[string]interface{}{
				"package": inputPackage,
			},
			action: "update",
			client: p,
		},
	}

	action, _ := p.ddl.ActionInterface(d.r.action)
	action.SetDefaults(d.r.args)

	return d
}

// YumCheckupdates performs the yum_checkupdates action
//
// Description: Check for YUM updates
func (p *PackageClient) YumCheckupdates() *YumCheckupdatesRequester {
	d := &YumCheckupdatesRequester{
		outc: nil,
		r: &requester{
			args:   map[string]interface{}{},
			action: "yum_checkupdates",
			client: p,
		},
	}

	action, _ := p.ddl.ActionInterface(d.r.action)
	action.SetDefaults(d.r.args)

	return d
}

// YumClean performs the yum_clean action
//
// Description: Clean the YUM cache
//
// Optional Inputs:
//    - mode (string) - One of the various supported clean modes
func (p *PackageClient) YumClean() *YumCleanRequester {
	d := &YumCleanRequester{
		outc: nil,
		r: &requester{
			args:   map[string]interface{}{},
			action: "yum_clean",
			client: p,
		},
	}

	action, _ := p.ddl.ActionInterface(d.r.action)
	action.SetDefaults(d.r.args)

	return d
}
