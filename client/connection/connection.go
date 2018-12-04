package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"

	. "github.com/secure2work/norictl/client/consts"
)

type Connection struct {
	Name     string `json:"-"`
	Hostname string `json:"hostname"`
	Secure   bool   `json:"secure"`
	CertPath string `json:"cert_path"`
	path     string `json:"-"`
}

const ext = ".json"

type Connections []Connection

func (c Connection) Save(path string, force bool) error {
	if strings.HasSuffix(c.Name, ext) {
		path = filepath.Join(path, c.Name)
	} else {
		path = filepath.Join(path, c.Name+ext)
	}

	var f *os.File
	if _, err := os.Stat(path); os.IsNotExist(err) || force {
		f, err = os.Create(path)
		defer f.Close()
		if err != nil {
			return fmt.Errorf("File %s could not be created: %s", path, err)
		}
	} else {
		return fmt.Errorf("File %s already exist. Use %s for overwrite", path, CONN_CREATE_FORCE)
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	err := enc.Encode(&c)
	if err != nil {
		return fmt.Errorf("Error %s with encoding to file %s", err, path)
	}
	return nil
}

func List(path string) (Connections, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	connections := make(Connections, 0)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		if filepath.Ext(f.Name()) != ext {
			continue
		}

		name := filepath.Join(path, f.Name())

		file, err := os.Open(name)
		if err != nil {
			return nil, fmt.Errorf("Open error \"%s\" with file %s", err, name)
		}
		conn := new(Connection)
		err = json.NewDecoder(file).Decode(conn)
		file.Close()
		if err != nil {
			return nil, fmt.Errorf("Decode error \"%s\" with file %s", err, name)
		}
		conn.Name = strings.TrimSuffix(f.Name(), ext)
		conn.path = path
		connections = append(connections, *conn)
	}
	return connections, nil
}

func (c Connection) Remove() error {
	name := filepath.Join(c.path, c.Name+ext)
	return os.Remove(name)
}

func (c Connection) Render(format string) (string, error) {
	var b bytes.Buffer
	switch format {
	case "table":
		table := tablewriter.NewWriter(&b)
		table.SetHeader([]string{"Name", "Hostname", "Secure", "CertPath"})
		table.Append([]string{c.Name, c.Hostname, fmt.Sprintf("%t", c.Secure), c.CertPath})
		table.Render()
		return b.String(), nil
	case "json":
		enc := json.NewEncoder(&b)
		err := enc.Encode(&c)
		if err != nil {
			return "", err
		}
		return b.String(), nil
	default:
		return "", fmt.Errorf("Format %s not supported", format)
	}
}

func (c Connection) Test(verbose bool) error {
	return nil
}

func (cs Connections) FilterByName(name string) (*Connection, error) {
	for _, c := range cs {
		if c.Name == name {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("Connection with name %q not exist", name)
}
