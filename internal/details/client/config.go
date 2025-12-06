package client

type Config struct {
	proxyURL          string
	disableKeepAlives bool
	afterDialHook     func()
}

func (c Config) AfterDialHook() func() {
	return c.afterDialHook
}

func (c Config) DisableKeepAlives() bool {
	return c.disableKeepAlives
}

func (c Config) ProxyURL() string {
	return c.proxyURL
}

type Option func(*Config)

func WithProxyURL(proxyURL string) Option {
	return func(c *Config) {
		c.proxyURL = proxyURL
	}
}

func WithDisableKeepAlives() Option {
	return func(c *Config) {
		c.disableKeepAlives = true
	}
}

func WithAfterDialHook(afterDialHook func()) Option {
	return func(c *Config) {
		c.afterDialHook = afterDialHook
	}
}
