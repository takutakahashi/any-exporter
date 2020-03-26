package metrics

import (
	"testing"

	"github.com/takutakahashi/any-exporter/pkg/config"
)

func TestMetricsStore(t *testing.T) {
	ms := NewMetricsStore()
	mc := config.MetricsConfig{WorkDir: "../../example/test_metrics", ScriptPath: "./metrics", Args: []string{"takutakahashi"}}
	m := NewMetrics(mc)
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
	m := NewMetrics(mc)
	t.Log(mc)
	t.Log(m.c)
	err := m.Execute()
	if err != nil {
		t.Fatal(err)
	}
}
