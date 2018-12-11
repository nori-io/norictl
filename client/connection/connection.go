package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/secure2work/norictl/client/consts"
)

type Connection struct {
	Name     string `json:"-"`
	Host     string `json:"host"`
	Port     uint64 `json:"port"`
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
		return fmt.Errorf("File %s already exist. Use %s for overwrite", path, "force")
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

		conn, err := open(name)
		if err != nil {
			return nil, err
		}

		conn.Name = strings.TrimSuffix(f.Name(), ext)
		conn.path = path
		connections = append(connections, *conn)
	}
	return connections, nil
}

func open(name string) (*Connection, error) {
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
	return conn, nil
}

func Get(path, name string) (*Connection, error) {
	name = filepath.Join(path, name+ext)
	return open(name)
}

func (c Connection) Remove() error {
	name := filepath.Join(c.path, c.Name+ext)
	return os.Remove(name)
}

func (c Connection) Use(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(c.Name)
	if err != nil {
		return err
	}

	return nil
}

func (c Connection) Render(format string) (string, error) {
	return render(format, c)
}

func (cs Connections) Render(format string) (string, error) {
	return render(format, cs...)
}

func render(format string, cs ...Connection) (string, error) {
	var b bytes.Buffer

	switch format {
	case "table":
		if len(cs) > 0 {
			table := tablewriter.NewWriter(&b)
			table.SetHeader([]string{"Name", "Host", "Port", "CertPath"})
			for _, c := range cs {
				table.Append([]string{c.Name, c.Host, strconv.Itoa(int(c.Port)), c.CertPath})
			}
			table.Render()
		} else {
			return "", fmt.Errorf("Connections not exists.")
		}
	case "json":
		enc := json.NewEncoder(&b)
		err := enc.Encode(&cs)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("Format %s not supported", format)
	}
	return b.String(), nil
}

func (cs Connections) FilterByName(name string) (*Connection, error) {
	for _, c := range cs {
		if c.Name == name {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("Connection with name %q not exist", name)
}

func (cs Connections) FilterBySecure(secure bool) Connections {
	newConns := make(Connections, 0)
	for _, c := range cs {
		if (len(c.CertPath) > 0) == secure {
			newConns = append(newConns, c)
		}
	}
	return newConns
}

func (cs Connections) FilterByHost(host string) Connections {
	newConns := make(Connections, 0)
	for _, c := range cs {
		if c.Host == host {
			newConns = append(newConns, c)
		}
	}
	return newConns
}

func CurrentConnection() (*Connection, error) {
	bs, err := ioutil.ReadFile(consts.UseFilePath)
	if err != nil {
		return nil, err
	}

	name := string(bytes.TrimSpace(bs))

	conn, err := Get(consts.ConnectionsDir, name)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (c Connection) HostPort() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(int(c.Port)))
}
