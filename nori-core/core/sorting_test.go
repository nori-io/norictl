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
	item4 := new(MockPlugin)
	item5_verr := new(MockPlugin)
	item6 := new(MockPlugin)
	item7_kerr := new(MockPlugin)

	item0.On("GetMeta").Return(&plugins.PluginMetaStruct{
		PluginName:   "item31",
		Namespace:    "ns31",
		Dependencies: []string{},
		Version:      "2.0.0",
	})
	item1.On("GetMeta").Return(&plugins.PluginMetaStruct{
		PluginName:   "item14",
		Namespace:    "ns14",
		Dependencies: []string{"ns31/item31"},
		Kind:         plugins.HTTP,
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
	item4.On("GetMeta").Return(&plugins.PluginMetaStruct{
		PluginName:   "item4",
		Namespace:    "ns3",
		Dependencies: []string{"ns31/item31:>=2"},
	})
	item5_verr.On("GetMeta").Return(&plugins.PluginMetaStruct{
		PluginName:   "item5",
		Namespace:    "ns5",
		Dependencies: []string{"ns31/item31:>2"},
	})
	item6.On("GetMeta").Return(&plugins.PluginMetaStruct{
		PluginName:   "item6",
		Namespace:    "ns6",
		Dependencies: []string{"kind:http"},
	})
	item7_kerr.On("GetMeta").Return(&plugins.PluginMetaStruct{
		PluginName:   "item7",
		Namespace:    "ns7",
		Dependencies: []string{"kind:cache"},
	})

	list := PluginList{item2, item0, item3_err, item7_kerr, item5_verr, item4, item1, item6}

	sortedList := SortPlugins(list)

	assert.Equal(sortedList[0].GetMeta().GetPluginName(), "item31")
	assert.Nil(sortedList[0].error)

	assert.Equal(sortedList[1].GetMeta().GetPluginName(), "item14")
	assert.Nil(sortedList[1].error)

	assert.Equal(sortedList[2].GetMeta().GetPluginName(), "item4")
	assert.Nil(sortedList[2].error)

	assert.Equal(sortedList[3].GetMeta().GetPluginName(), "item6")
	assert.Nil(sortedList[3].error)

	assert.Equal(sortedList[4].GetMeta().GetPluginName(), "item2")
	assert.Nil(sortedList[4].error)

	assert.Equal(sortedList[5].GetMeta().GetPluginName(), "item3")
	assert.Equal(sortedList[5].error.Error(), "Dependencies err/err for plugin ns3/item3 not found.")

	assert.Equal(sortedList[6].GetMeta().GetPluginName(), "item7")
	assert.Equal(sortedList[6].error.Error(), "Dependencies kind:cache for plugin ns7/item7 not found.")

	assert.Equal(sortedList[7].GetMeta().GetPluginName(), "item5")
	assert.Equal(sortedList[7].error.Error(), "Wrong version for plugin ns5/item5. Have: 2.0.0. Want: >2")
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
