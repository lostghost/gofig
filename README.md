# gofig

Golang configuration library for building friendly, easy to develop and configure 12-factor apps.

For development it is easiest to have local environment configuration files. This makes it easy to run the application locally durring development, however, upon moving to production, we really want to be configuring the application via env vars.

Gofig is a library designed to abstract the configuration process to allow several methods of providing the application configuration so that we can move our application from convenient local development to secure and flexible production configuration.

Possible methods for providing configuration include:

- configuration file(s)
- env vars
- command line flags

## Usage

```go
import github.com/lostghost/gofig

var config struct {
    env string
    port int
}

gofig.RegisterConfigfile('./config.yaml', false)
gofig.String(&config.env, 'ENV', 'Application Environment', 'development')
gofig.Integer(&config.port, 'SERVER_PORT', 'Server Port', 8080)
gofig.Parse()
```

### Configuration File Formats

TBD: INI, YAML, TOML, JSON?