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
		PluginName:   "item31",
		Namespace:    "ns31",
		Dependencies: []string{},
	})
	item1.On("GetMeta").Return(&plugins.PluginMetaStruct{
		PluginName:   "item14",
		Namespace:    "ns14",
		Dependencies: []string{"ns31/item31"},
	})
	item2.On("GetMeta").Return(&plugins.PluginMetaStruct{
		PluginName:   "item2",
		Namespace:    "ns2",
		Dependencies: []string{"ns31/item31", "ns14/item14"},
	})
	item3_err.On("GetMeta").Return(&plugins.PluginMetaStruct{
		PluginName:   "item3",
		Namespace:    "ns3",
		Dependencies: []string{"err/err"},
	})

	list := PluginList{item2, item0, item3_err, item1}

	sortedList := SortPlugins(list)

	assert.Equal(sortedList[0].GetMeta().GetPluginName(), "item31")
	assert.Nil(sortedList[0].error)

	assert.Equal(sortedList[1].GetMeta().GetPluginName(), "item14")
	assert.Nil(sortedList[1].error)

	assert.Equal(sortedList[2].GetMeta().GetPluginName(), "item2")
	assert.Nil(sortedList[2].error)

	assert.Equal(sortedList[3].GetMeta().GetPluginName(), "item3")
	assert.Equal(sortedList[3].error.Error(), "Dependencies err/err for plugin ns3/item3 not found.")
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
