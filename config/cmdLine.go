package config

import (
	"fmt"
	"strings"

	"github.com/ah-its-andy/downton/collections"
	"github.com/ah-its-andy/downton/core"
)

// builder apis
func AddCommandLine(builder core.ConfigurationBuilder, args []string, switchMapping map[string]string) {
	builder.AddSource(NewCmdLineConfigurationSource(switchMapping, args))
}

type CmdLineConfigurationSource struct {
	switchMapping map[string]string
	args          []string
}

func NewCmdLineConfigurationSource(switchMapping map[string]string, args []string) *CmdLineConfigurationSource {
	if switchMapping == nil {
		switchMapping = make(map[string]string)
	}
	if args == nil {
		args = make([]string, 0)
	}
	return &CmdLineConfigurationSource{
		switchMapping: switchMapping,
		args:          args,
	}
}

func (source *CmdLineConfigurationSource) Build(builder core.ConfigurationBuilder) core.ConfigurationProvider {
	return NewCmdLineConfigurationProvider(source.switchMapping, source.args)
}

type CmdLineConfigurationProvider struct {
	ConfigurationProvider

	switchMapping map[string]string
	args          []string
}

func NewCmdLineConfigurationProvider(switchMapping map[string]string, args []string) *CmdLineConfigurationProvider {
	if switchMapping == nil {
		switchMapping = make(map[string]string)
	}
	if args == nil {
		args = make([]string, 0)
	}
	return &CmdLineConfigurationProvider{
		switchMapping:         switchMapping,
		args:                  args,
		ConfigurationProvider: *NewConfigurationProvider(),
	}
}

func (provider *CmdLineConfigurationProvider) Arguments() []string {
	ret := make([]string, len(provider.args))
	copy(ret, provider.args)
	return ret
}

func (provider *CmdLineConfigurationProvider) Load() {
	data := make(map[string]string)
	k, v := "", ""
	argList := collections.NewArrayList[string](0)
	argList.AddRange(provider.args...)
	iterator := argList.GetIterator()
	for iterator.MoveNext() {
		arg := iterator.Current()
		keyStartIndex := 0
		if strings.HasPrefix(arg, "--") {
			keyStartIndex = 2
		} else if strings.HasPrefix(arg, "-") {
			keyStartIndex = 1
		}
		separatorIndex := strings.Index(arg, "=")

		if separatorIndex < 0 {
			if keyStartIndex == 0 {
				// Ignore invalid formats
				continue
			}
			if mappedK, ok := provider.switchMapping[arg]; ok {
				k = mappedK
			} else if keyStartIndex == 1 {
				//If the switch starts with a single "-" and it isn't in given mappings , it is an invalid usage so ignore it
				continue
			} else {
				// Otherwise, use the switch name directly as a key
				k = arg[keyStartIndex:]
			}
			if !iterator.MoveNext() {
				continue
			}
			v = iterator.Current()
		} else {
			keySegment := arg[keyStartIndex : separatorIndex-keyStartIndex]
			if mappedK, ok := provider.switchMapping[keySegment]; ok {
				k = mappedK
			} else if keyStartIndex == 1 {
				panic(fmt.Sprintf("CmdLineProvider.Load : Invalid usage of switch '%s'", arg))
			} else {
				k = keySegment
			}
			v = arg[separatorIndex+1:]
		}
		data[k] = v
	}
	provider.data = data
}
