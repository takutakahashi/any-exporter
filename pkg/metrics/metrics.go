package metrics

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/takutakahashi/any-exporter/pkg/config"
	"golang.org/x/sync/errgroup"
)

type Metrics struct {
	c      config.MetricsConfig
	sc     config.Config
	Result MetricsResult
}

type MetricsResult struct {
	Name   string        `json:"name"`
	Values []MetricsBody `json:"values"`
}

type MetricsBody struct {
	Labels map[string]string `json:"labels"`
	Value  float64           `json:"value"`
}

type MetricsStore struct {
	store map[string]*Metrics
}

func (m MetricsStore) Add(v Metrics) {
	m.store[v.c.WorkDir] = &v
}

func (m MetricsStore) Delete(v Metrics) {
	delete(m.store, v.c.WorkDir)
}

func (ms MetricsStore) Length() int {
	return len(ms.store)
}

func (ms *MetricsStore) ExecuteAll() error {
	g := new(errgroup.Group)
	for _, v := range ms.store {
		g.Go(v.Execute)
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (m *Metrics) Sleep() {
	duration, err := time.ParseDuration(m.c.Resolution)
	if err != nil {
		duration, err = time.ParseDuration(m.sc.Resolution)
		if err != nil {
			duration = 30 * time.Second
		}
	}
	time.Sleep(duration)
}

func (m *Metrics) Execute() error {
	ctx := context.TODO()
	prev, err := filepath.Abs(".")
	if err != nil {
		return err
	}
	defer os.Chdir(prev)
	os.Chdir(m.c.WorkDir)
	cmd := exec.CommandContext(ctx, m.c.ScriptPath, m.c.Args...)
	cmd.Env = os.Environ()
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}
	output := strings.Split(string(stdout), "\n")
	_ = []byte(output[len(output)-1])
	err = json.Unmarshal([]byte(output[len(output)-1]), &m.Result)
	return err
}

func (m Metrics) String() string {
	result := ""
	for _, value := range m.Result.Values {
		label := []string{}
		for key, val := range value.Labels {
			if key != "value" {
				label = append(label, fmt.Sprintf("%s=\"%s\"", key, val))
			}
		}
		result += fmt.Sprintf("%s{%s} %f\n", m.Result.Name, strings.Join(label, ","), value.Value)
	}
	return result
}

func (m MetricsStore) String() string {
	result := ""
	for _, v := range m.store {
		vr := v.String()
		result += vr
	}
	return result
}

func NewMetrics(c config.MetricsConfig, sc config.Config) Metrics {
	return Metrics{c: c, sc: sc, Result: MetricsResult{}}
}

func NewMetricsStore() MetricsStore {
	return MetricsStore{store: map[string]*Metrics{}}
}
