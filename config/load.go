package config

import (
	"flag"
	"github.com/AlecAivazis/survey/v2"
	"runtime"
	"strings"
)

type Config struct {
	AppName    string
	Package    string
	Middleware []string
	Cmd        string
	Mode       int
	ImportPath string
	GoVer      string
}

func NewConfig() *Config {
	return &Config{}
}

const (
	ModeGen = iota + 1
	ModeHelp
)

func (c *Config) Load() error {
	flag.Parse()
	c.Cmd = flag.Arg(0)
	c.ImportPath = flag.Arg(1)
	if c.ImportPath == "" {
		c.ImportPath = "."
	}
	c.GoVer = strings.Join(strings.Split(runtime.Version(), ".")[:2], ".")[2:]
	switch c.Cmd {
	case "new":
		c.Mode = ModeGen
	case "help":
		c.Mode = ModeHelp
	}
	return nil
}

var qs = []*survey.Question{
	{
		Name: "appName",
		Prompt: &survey.Input{
			Message: "Please your app name",
			Help:    "please input your app name",
		},
		Validate: survey.Required,
	},
	{
		Name: "package",
		Prompt: &survey.Input{
			Message: "Please your go module name",
			Help:    "please input like github.com/gari8/worgen",
		},
		Validate: survey.Required,
	},
	{
		Name: "middleware",
		Prompt: &survey.MultiSelect{
			Message: "Please select your middleware",
			Options: []string{"minio"},
			Default: nil,
			Help:    "please input your middleware",
		},
	},
}

type answer struct {
	AppName    string   `survey:"appName"`
	Package    string   `survey:"package"`
	Middleware []string `survey:"middleware"`
}

func (c *Config) Conversation() error {
	var ans answer
	if err := survey.Ask(qs, &ans); err != nil {
		return err
	}
	c.Package = ans.Package
	c.Middleware = ans.Middleware
	c.AppName = ans.AppName
	return nil
}
