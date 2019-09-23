package template

const (
	DefaultRunAddr   = ":31000"
	DefaultName      = "/template"
	DefaultApiPrefix = "/path"
)

func setCfgDefaultOption() *Cfg {
	return &Cfg{
		RunAddr:   DefaultRunAddr,
		APIBase:   DefaultName,
		ApiPrefix: DefaultApiPrefix,
	}
}

type Cfg struct {
	RunAddr   string
	APIBase   string
	ApiPrefix string
}

type CfgFunc func(*Cfg)

var (
	defaultCfgOption = setCfgDefaultOption()
)

func NewCfg(opts ...CfgFunc) (opt *Cfg) {
	opt = defaultCfgOption
	for _, o := range opts {
		o(opt)
	}
	defaultCfgOption = setCfgDefaultOption()
	return
}

func WithRunAddr(runAddr string) CfgFunc {
	return func(o *Cfg) {
		o.RunAddr = runAddr
	}
}

func WithAPIBase(apiBase string) CfgFunc {
	return func(o *Cfg) {
		o.APIBase = apiBase
	}
}

func WithApiPrefix(apiPrefix string) CfgFunc {
	return func(o *Cfg) {
		o.ApiPrefix = apiPrefix
	}
}
