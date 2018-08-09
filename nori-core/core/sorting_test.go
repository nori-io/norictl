package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/secure2work/nori/core/plugins"
)

func TestSorting(t *testing.T) {
	assert := assert.New(t)

	item0 := new(MockPlugin)
	item1 := new(MockPlugin)
	item2 := new(MockPlugin)
	item3_err := new(MockPlugin)

	item0.On("GetMeta").Return(&plugins.PluginMetaStruct{
		Description:  "0",
		PluginName:   "item0",
		Namespace:    "ns0",
		Dependencies: []string{},
	})
	item1.On("GetMeta").Return(&plugins.PluginMetaStruct{
		Description:  "1",
		PluginName:   "item1",
		Namespace:    "ns1",
		Dependencies: []string{"ns0/item0"},
	})
	item2.On("GetMeta").Return(&plugins.PluginMetaStruct{
		Description:  "2",
		PluginName:   "item2",
		Namespace:    "ns2",
		Dependencies: []string{"ns0/item0", "ns1/item1"},
	})
	item3_err.On("GetMeta").Return(&plugins.PluginMetaStruct{
		Description:  "3",
		PluginName:   "item3",
		Namespace:    "ns3",
		Dependencies: []string{"err/err"},
	})

	list := PluginList{item2, item0, item3_err, item1}

	sortedList := SortPlugins(list)

	assert.Equal(sortedList[0].GetMeta().GetDescription(), "0")
	assert.Nil(sortedList[0].error)

	assert.Equal(sortedList[1].GetMeta().GetDescription(), "1")
	assert.Nil(sortedList[1].error)

	assert.Equal(sortedList[2].GetMeta().GetDescription(), "2")
	assert.Nil(sortedList[2].error)

	assert.Equal(sortedList[3].GetMeta().GetDescription(), "3")
	assert.Equal(sortedList[3].Error(), "Dependencies err/err for plugin ns3/item3 not found.")
}

type MockPlugin struct {
	mock.Mock
}

func (_ MockPlugin) GetInstance() interface{} {
	return nil
}

func (m MockPlugin) GetMeta() plugins.PluginMeta {
	args := m.Mock.Called()
	return args.Get(0).(plugins.PluginMeta)
}

func (_ MockPlugin) Install(_ context.Context, _ plugins.PluginRegistry) error {
	return nil
}

func (_ MockPlugin) UnInstall(_ context.Context, _ plugins.PluginRegistry) error {
	return nil
}

func (_ MockPlugin) Start(_ context.Context, _ plugins.PluginRegistry) error {
	return nil
}

func (_ MockPlugin) Stop(_ context.Context) error {
	return nil
}
