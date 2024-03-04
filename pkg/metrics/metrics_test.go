package metrics

import (
	"testing"

	"github.com/takutakahashi/any-exporter/pkg/config"
)

func TestMetricsStore(t *testing.T) {
	ms := NewMetricsStore()
	mc := config.MetricsConfig{WorkDir: "../../example/test_metrics", ScriptPath: "./metrics", Args: []string{"takutakahashi"}}
	c := config.Config{Metrics: []config.MetricsConfig{mc}, Resolution: "1s"}
	m := NewMetrics(mc, c)
	ms.Add(m)
	if ms.String() != "" {
		t.Fatal(ms.String())
	}
	err := ms.ExecuteAll()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMetrics(t *testing.T) {
	mc := config.MetricsConfig{WorkDir: "../../example/test_metrics", ScriptPath: "./metrics", Args: []string{"takutakahashi"}}
	c := config.Config{Metrics: []config.MetricsConfig{mc}, Resolution: "1s"}
	m := NewMetrics(mc, c)
	t.Log(mc)
	t.Log(m.c)
	err := m.Execute()
	if err != nil {
		t.Fatal(err)
	}
}
