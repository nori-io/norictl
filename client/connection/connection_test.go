package connection

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testData1 = Connection{
	Name:     "test1",
	Hostname: "testhost1",
	Secure:   true,
	CertPath: "testpath1",
}

var testData2 = Connection{
	Name:     "test2",
	Hostname: "testhost2",
	Secure:   true,
	CertPath: "testpath2",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func tmpDir() string {
	tmp := filepath.Join(os.TempDir(), fmt.Sprintf("%d", rand.Int63()))
	os.Mkdir(tmp, os.ModePerm)
	return tmp
}

func TestSave(t *testing.T) {
	assert := assert.New(t)

	tmp := tmpDir()
	defer os.RemoveAll(tmp)

	err := testData1.Save(tmp, false)
	assert.Nil(err)

	err = testData1.Save(tmp, false)
	assert.NotNil(err)

	err = testData1.Save(tmp, true)
	assert.Nil(err)
}

func TestList(t *testing.T) {
	assert := assert.New(t)

	tmp := tmpDir()
	defer os.RemoveAll(tmp)

	err := testData1.Save(tmp, false)
	assert.Nil(err)

	err = testData2.Save(tmp, false)
	assert.Nil(err)

	connLst, err := List(tmp)
	assert.Nil(err)
	assert.Len(connLst, 2)

	assert.Equal(connLst[0].Name, testData1.Name)
	assert.Equal(connLst[1].Name, testData2.Name)
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)

	tmp := tmpDir()
	defer os.RemoveAll(tmp)

	err := testData1.Save(tmp, false)
	assert.Nil(err)

	_, err = os.Stat(filepath.Join(tmp, testData1.Name+ext))
	assert.False(os.IsNotExist(err))

	connLst, err := List(tmp)
	assert.Nil(err)
	conn, err := connLst.FilterByName(testData1.Name)
	assert.Nil(err)
	err = conn.Remove()
	assert.Nil(err)

	_, err = os.Stat(filepath.Join(tmp, testData1.Name+ext))
	assert.True(os.IsNotExist(err))
}

func TestRender(t *testing.T) {
	assert := assert.New(t)

	tmp := tmpDir()
	defer os.RemoveAll(tmp)

	str, err := testData1.Render("json")
	assert.Nil(err)
	assert.NotEmpty(str)

	str, err = testData1.Render("table")
	assert.Nil(err)
	assert.NotEmpty(str)
}

func TestFilterByName(t *testing.T) {
	assert := assert.New(t)

	tmp := tmpDir()
	defer os.RemoveAll(tmp)

	err := testData1.Save(tmp, false)
	assert.Nil(err)

	connLst, err := List(tmp)
	assert.Nil(err)
	_, err = connLst.FilterByName(testData1.Name)
	assert.Nil(err)

	_, err = connLst.FilterByName(testData2.Name)
	assert.NotNil(err)
}
