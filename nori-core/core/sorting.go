package core

import (
	"fmt"
	"sort"

	"github.com/secure2work/nori/core/plugins"
)

type PluginList []plugins.Plugin

type internalPluginList []*internalPlugin

type nameList map[string]*internalPlugin

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

func SortPlugins(pList PluginList) RespPluginList {
	// just the map for fast access to a slice item
	nList := make(nameList)

	// transforming to internal format
	list := toInternalList(pList)

	// prepare to sort
	for _, plugin := range list {
		// TODO check for duplicates
		nList[plugin.name()] = plugin
	}

	// check dependencies
	for _, plugin := range list {
		err := plugin.checkDeps(nList)
		if err != nil {
			plugin.error = err
		}
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

func (i *internalPlugin) calcWeight(nList nameList) {
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
			plug := nList[name]
			if plug.calculatedWeight {
				// already calculated
				i.weight += plug.weight
			} else {
				// "To understand recursion, you must understand recursion."
				plug.calcWeight(nList)
				i.weight += plug.weight
			}
		}
	}
}

func (i *internalPlugin) checkDeps(nList nameList) error {
	// TODO check for self-dependence
	for _, name := range i.GetMeta().GetDependencies() {
		if _, ok := nList[name]; !ok {
			return fmt.Errorf("Dependencies %s for plugin %s not found.", name, i.name())
		}
	}
	return nil
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
