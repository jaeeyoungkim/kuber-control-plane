package lib

import (
	"context"
	"fmt"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"strings"
)

var rawMetrics map[string]*io_prometheus_client.MetricFamily

func getMetrics(familyName string) []*io_prometheus_client.Metric {
	var apiserverRequestTotalMetricFamily *io_prometheus_client.MetricFamily
	apiserverRequestTotalMetricFamily = rawMetrics[familyName]
	metric := apiserverRequestTotalMetricFamily.GetMetric()
	return metric
}

func GetRawMetrics() {
	client, err, done := GetKubernetesClient()
	if !done {
		fmt.Printf("error getting client: %v\n", err)
	}

	bytes, err := client.RESTClient().Get().AbsPath("/metrics").DoRaw(context.Background())
	if err != nil {
		fmt.Printf("error getting metrics: %v\n", err)
	}

	var parser expfmt.TextParser
	metricFamilies, err := parser.TextToMetricFamilies(strings.NewReader(string(bytes)))
	if err != nil {
		fmt.Printf("error parsing metrics: %v\n", err)
	}

	fmt.Println("raw metrics 갱신")
	rawMetrics = metricFamilies
}

func GetApiserverRequestTotal() []ApiServerRequestTotal {
	metrics := getMetrics("apiserver_request_total")
	var dataArr []ApiServerRequestTotal
	for _, m := range metrics {
		label := m.GetLabel()
		data := ApiServerRequestTotal{}

		for _, l := range label {
			name := l.GetName()
			value := l.GetValue()
			logging("apiserver_request_total", name, value)
			switch name {
			case "code":
				data.Code = value
			case "component":
				data.Component = value
			case "dryRun":
				data.DryRun = value
			case "group":
				data.Group = value
			case "resource":
				data.Resource = value
			case "scope":
				data.Scope = value
			case "subresource":
				data.Subresource = value
			case "verb":
				data.Verb = value
			case "version":
				data.Version = value
			}
		}
		data.Counter = m.GetCounter().GetValue()
		dataArr = append(dataArr, data)
	}
	return dataArr
}

func GetApiserverRequestDurationSecondsBucket() []ApiserverRequestDurationSecondsBucket {
	metrics := getMetrics("apiserver_request_duration_seconds")
	var dataArr []ApiserverRequestDurationSecondsBucket
	for _, m := range metrics {
		label := m.GetLabel()
		labelMap := make(map[string]string)
		for _, l := range label {
			name := l.GetName()
			value := l.GetValue()
			logging("apiserver_request_duration_seconds", name, value)
			switch name {
			case "component":
				labelMap["component"] = value
			case "dryRun":
				labelMap["dryRun"] = value
			case "group":
				labelMap["group"] = value
			case "resource":
				labelMap["resource"] = value
			case "scope":
				labelMap["scope"] = value
			case "subresource":
				labelMap["subresource"] = value
			case "verb":
				labelMap["verb"] = value
			case "version":
				labelMap["version"] = value
			}
		}

		bucket := m.GetHistogram().GetBucket()
		for _, histogram := range bucket {
			data := ApiserverRequestDurationSecondsBucket{}
			data.Component = labelMap["component"]
			data.DryRun = labelMap["dryRun"]
			data.Group = labelMap["group"]
			data.Resource = labelMap["resource"]
			data.Scope = labelMap["scope"]
			data.Subresource = labelMap["subresource"]
			data.Verb = labelMap["verb"]
			data.Version = labelMap["version"]
			data.Le = histogram.GetUpperBound()
			data.CumulativeCount = histogram.GetCumulativeCount()
			dataArr = append(dataArr, data)
		}
	}
	return dataArr
}

func GetApiserverCurrentInflightRequests() []ApiserverCurrentInflightRequests {
	metrics := getMetrics("apiserver_current_inflight_requests")
	var dataArr []ApiserverCurrentInflightRequests

	for _, m := range metrics {
		label := m.GetLabel()
		data := ApiserverCurrentInflightRequests{}
		for _, l := range label {
			name := l.GetName()
			value := l.GetValue()
			logging("apiserver_request_total", name, value)
			switch name {
			case "request_kind":
				data.RequestKind = value
			}
		}
		data.Gauge = m.GetGauge().GetValue()
		dataArr = append(dataArr, data)
	}
	return dataArr
}

func GetApiserverAuditLevelTotal() []ApiserverAuditLevelTotal {
	metrics := getMetrics("apiserver_audit_level_total")
	var dataArr []ApiserverAuditLevelTotal
	for _, m := range metrics {
		label := m.GetLabel()
		data := ApiserverAuditLevelTotal{}

		for _, l := range label {
			name := l.GetName()
			value := l.GetValue()
			logging("apiserver_audit_level_total", name, value)
			switch name {
			case "level":
				data.Level = value
			}
		}
		data.Counter = m.GetCounter().GetValue()
		dataArr = append(dataArr, data)
	}
	return dataArr
}

func GetGoGoroutines() []GoGoroutines {
	metrics := getMetrics("go_goroutines")
	var dataArr []GoGoroutines
	for _, m := range metrics {
		label := m.GetLabel()
		data := GoGoroutines{}

		for _, l := range label {
			name := l.GetName()
			value := l.GetValue()
			logging("go_goroutines", name, value)
		}
		data.Gauge = m.GetGauge().GetValue()
		dataArr = append(dataArr, data)
	}
	return dataArr
}

func GetGoThreads() []GoThreads {
	metrics := getMetrics("go_threads")
	var dataArr []GoThreads
	for _, m := range metrics {
		label := m.GetLabel()
		data := GoThreads{}

		for _, l := range label {
			name := l.GetName()
			value := l.GetValue()
			logging("go_goroutines", name, value)
		}
		data.Gauge = m.GetGauge().GetValue()
		dataArr = append(dataArr, data)
	}
	return dataArr
}

func GetEtcdRequestDurationSecondsBucket() []EtcdRequestDurationSecondsBucket {
	metrics := getMetrics("etcd_request_duration_seconds")
	var dataArr []EtcdRequestDurationSecondsBucket
	for _, m := range metrics {
		label := m.GetLabel()
		labelMap := make(map[string]string)
		for _, l := range label {
			name := l.GetName()
			value := l.GetValue()
			logging("etcd_request_duration_seconds", name, value)
			switch name {
			case "operation":
				labelMap["operation"] = value
			case "type":
				labelMap["type"] = value
			}
		}

		bucket := m.GetHistogram().GetBucket()
		for _, histogram := range bucket {
			data := EtcdRequestDurationSecondsBucket{}
			data.Operation = labelMap["operation"]
			data.Type = labelMap["type"]
			data.Le = histogram.GetUpperBound()
			data.CumulativeCount = histogram.GetCumulativeCount()
			dataArr = append(dataArr, data)
		}
	}
	return dataArr

}

func logging(family string, label string, value string) {
	fmt.Printf(family, label, value)
}
