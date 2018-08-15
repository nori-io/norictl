package core

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/go-version"

	"github.com/secure2work/nori/core/plugins"
)

type PluginList []plugins.Plugin

type internalPluginList []*internalPlugin

type nameList map[string]*internalPlugin

type kindList map[plugins.PluginKind]*internalPlugin

type RespPluginList []RespPlugin

type internalPlugin struct {
	weight           int
	calculatedWeight bool
	plugins.Plugin
	error
}

type RespPlugin struct {
	plugins.Plugin
	error
}

const (
	kindPrefix          = "kind:"
	constraintSeparator = ":"
)

func SortPlugins(pList PluginList) RespPluginList {
	// just the map for fast access to a slice item
	nList := make(nameList)
	kList := make(kindList)

	// transforming to internal format
	list := toInternalList(pList)

	// prepare to sort
	for _, plugin := range list {
		// TODO check for duplicates
		nList[plugin.name()] = plugin
		kList[plugin.Plugin.GetMeta().GetKind()] = plugin
	}

	// check dependencies
	for _, plugin := range list {
		err := plugin.checkDeps(nList, kList)
		if err != nil {
			plugin.error = err
		}
	}

	// calculate weight
	for _, plugin := range list {
		plugin.calcWeight(nList, kList)
	}

	// sorting C.O.
	sort.Sort(list)

	// we need return the errors? rewrite if not
	resp := make(RespPluginList, len(list))
	for i, p := range list {
		resp[i] = RespPlugin{
			Plugin: p.Plugin,
			error:  p.error,
		}
	}

	return resp
}

func toInternalList(list PluginList) internalPluginList {
	newList := make(internalPluginList, len(list))
	for i, plugin := range list {
		newList[i] = &internalPlugin{
			Plugin: plugin,
		}
	}
	return newList
}

func (i *internalPlugin) calcWeight(nList nameList, kList kindList) {
	// already calculated
	if i.calculatedWeight {
		return
	}

	deps := i.GetMeta().GetDependencies()

	// without dependencies the weight equals zero
	if len(deps) == 0 {
		i.calculatedWeight = true
		return
	} else {
		// each the dependency add to weight 1
		i.weight = len(deps)

		// also added to the weight a weight of each dependency
		for _, name := range deps {
			if isKind(name) {
				kind := str2kind(name)
				plug, ok := kList[kind]
				if ok {
					i.weight += plug.weight
				} else {
					plug.calcWeight(nList, kList)
					i.weight += plug.weight
				}
			} else {
				plug, ok := nList[name]
				if ok {
					if plug.calculatedWeight {
						// already calculated
						i.weight += plug.weight
					} else {
						// "To understand recursion, you must understand recursion."
						plug.calcWeight(nList, kList)
						i.weight += plug.weight
					}
				}
			}
		}
	}
}

func (i *internalPlugin) checkDeps(nList nameList, kList kindList) error {
	// TODO check for self-dependence
	for _, name := range i.GetMeta().GetDependencies() {
		if isKind(name) {
			kind := str2kind(name)
			_, ok := kList[kind]
			if !ok {
				i.calculatedWeight = true
				return fmt.Errorf("Dependencies %s for plugin %s not found.", name, i.name())
			}
		} else {
			var constraint string
			name, constraint = splitConstraint(name)

			depPlug, ok := nList[name]
			if !ok {
				i.calculatedWeight = true
				return fmt.Errorf("Dependencies %s for plugin %s not found.", name, i.name())
			}

			ver := depPlug.Plugin.GetMeta().GetVersion()
			check, err := versionCheck(ver, constraint)
			if err != nil {
				return err
			}
			if !check {
				return fmt.Errorf("Wrong version for plugin %s. Have: %s. Want: %s", i.name(), ver, constraint)
			}
		}
	}
	return nil
}

func isKind(name string) bool {
	return strings.HasPrefix(name, kindPrefix)
}

func versionCheck(ver, constraint string) (bool, error) {
	if len(constraint) == 0 {
		return true, nil
	}

	v, err := version.NewVersion(ver)
	if err != nil {
		return false, err
	}

	c, err := version.NewConstraint(constraint)
	if err != nil {
		return false, err
	}

	check := c.Check(v)
	return check, nil
}

func splitConstraint(name string) (string, string) {
	ss := strings.Split(name, constraintSeparator)
	if len(ss) == 1 {
		return ss[0], ""
	}
	return ss[0], ss[1]
}

func str2kind(kind string) plugins.PluginKind {
	kind = strings.TrimPrefix(kind, kindPrefix)
	kind = strings.ToLower(kind)
	var str string
	for i := 0; str != "unknown"; i++ {
		str = plugins.PluginKind(i).String()
		str = strings.ToLower(str)
		if str == kind {
			return plugins.PluginKind(i)
		}
	}
	return plugins.PluginKind(-1)
}

// name = namespace + "/" + name
func (i internalPlugin) name() string {
	return i.GetMeta().GetNamespace() + "/" + i.GetMeta().GetPluginName()
}

func (i internalPluginList) Len() int      { return len(i) }
func (i internalPluginList) Swap(x, y int) { i[x], i[y] = i[y], i[x] }
func (i internalPluginList) Less(x, y int) bool {
	// items with error have less priority
	if i[x].error == nil && i[y].error != nil {
		return true
	}
	if i[x].error != nil && i[y].error == nil {
		return false
	}

	// more weight = less priority
	if i[x].weight < i[y].weight {
		return true
	}

	// sorting by name
	if i[x].weight == i[y].weight && i[x].name() < i[y].name() {
		return true
	}

	return false
}
