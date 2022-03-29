package config

import (
	"math"
	"strconv"
	"strings"

	"github.com/ah-its-andy/downton/collections"
	"github.com/ah-its-andy/downton/core"
)

var ConfigurationKeyComparer = func(a1, a2 any) int {
	x := a1.(string)
	y := a2.(string)

	xParts := strings.Split(x, core.DefaultConfigPathKeyDelimiter)
	yParts := strings.Split(y, core.DefaultConfigPathKeyDelimiter)

	count := math.Min(float64(len(xParts)), float64(len(yParts)))
	for i := 0; i < int(count); i++ {
		innerX := xParts[i]
		innerY := yParts[i]

		xIsInt := false
		yIsInt := false
		value1, err := strconv.ParseInt(innerX, 10, 64)
		if err == nil {
			xIsInt = true
		}

		value2, err := strconv.ParseInt(innerY, 10, 64)
		if err == nil {
			yIsInt = true
		}

		result := 0
		if !xIsInt && !yIsInt {
			result = strings.Compare(innerX, innerY)
		} else if xIsInt && yIsInt {
			result = int(value1) - int(value2)
		} else if xIsInt {
			result = -1
		} else {
			result = 1
		}

		if result != 0 {
			return result
		}
	}

	// If we get here, the common parts are equal.
	// If they are of the same length, then they are totally identical
	return len(xParts) - len(yParts)
}

type ConfigurationProvider struct {
	data map[string]string
}

func NewConfigurationProvider() *ConfigurationProvider {
	return &ConfigurationProvider{
		data: make(map[string]string),
	}
}

func (provider *ConfigurationProvider) Map() map[string]string {
	ret := make(map[string]string)
	for k, v := range provider.data {
		ret[k] = v
	}
	return ret
}

func (provider *ConfigurationProvider) TryGet(k string) (string, core.ConfigurationProvider, bool) {
	if v, ok := provider.data[k]; ok {
		return v, provider, true
	} else {
		return "", provider, false
	}
}
func (provider *ConfigurationProvider) Set(k string, v string) {
	provider.data[k] = v
}
func (provider *ConfigurationProvider) Load() {

}
func (provider *ConfigurationProvider) GetChildKeys(earlierKeys []string, parentPath string) []string {
	results := collections.NewArrayList[string](0)
	if parentPath == "" {
		for k, _ := range provider.data {
			results.Add(segment(k, 0))
		}
	} else {
		for k, _ := range provider.data {
			if len(k) > len(parentPath) &&
				strings.HasPrefix(k, parentPath) &&
				k[len(parentPath)] == ':' {
				results.Add(segment(k, len(parentPath)+1))
			}
		}
	}
	results.AddRange(earlierKeys...)
	results = collections.OrderBy[string](results, ConfigurationKeyComparer).ToList()
	return results.ToArray()
}

func segment(k string, prefixLength int) string {
	//int indexOf = key.IndexOf(ConfigurationPath.KeyDelimiter, prefixLength, StringComparison.OrdinalIgnoreCase);
	//       return indexOf < 0 ? key.Substring(prefixLength) : key.Substring(prefixLength, indexOf - prefixLength);
	indexOfDelimiter := -1
	if len(k) > prefixLength {
		indexOfDelimiter = strings.Index(k[prefixLength:], core.DefaultConfigPathKeyDelimiter)

	}
	if indexOfDelimiter < 0 {
		return k[prefixLength:]
	} else {
		return k[prefixLength:indexOfDelimiter]
	}
}
