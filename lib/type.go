package lib

type ApiserverData struct {
	ApiserverRequestTotal                 []ApiServerRequestTotal
	ApiserverRequestDurationSecondsBucket []ApiserverRequestDurationSecondsBucket
	ApiserverRequestDurationSecondsCount  []ApiserverRequestDurationSecondsCount
	ApiserverRequestDurationSecondsSum    []ApiserverRequestDurationSecondsSum
	ApiserverCurrentInflightRequests      []ApiserverCurrentInflightRequests
	ApiserverAuditLevelTotal              []ApiserverAuditLevelTotal
	GoGoroutines                          []GoGoroutines
	GoThreads                             []GoThreads
	etcdRequestDurationSecondsBucket      []EtcdRequestDurationSecondsBucket
	etcdRequestDurationSecondsCount       []EtcdRequestDurationSecondsCount
	etcdRequestDurationSecondsSum         []EtcdRequestDurationSecondsSum
}

/*
Counter of apiserver requests broken out for each verb, dry run value, group, version, resource, scope, component, and HTTP response code.
Stability Level:STABLE
Type: Counter
*/
type ApiServerRequestTotal struct {
	Code        string  `json:"code"`
	Component   string  `json:"component"`
	DryRun      string  `json:"dry_run"`
	Group       string  `json:"group"`
	Resource    string  `json:"resource"`
	Scope       string  `json:"scope"`
	Subresource string  `json:"subresource"`
	Verb        string  `json:"verb"`
	Version     string  `json:"version"`
	Counter     float64 `json:"counter"`
}

/*
apiserver_request_duration_seconds
Response latency distribution in seconds for each verb, dry run value, group, version, resource, subresource, scope and component.
Stability Level:STABLE
Type: Histogram
*/
type ApiserverRequestDurationSecondsBucket struct {
	Component       string  `json:"component"`
	DryRun          string  `json:"dry_run"`
	Group           string  `json:"group"`
	Resource        string  `json:"resource"`
	Scope           string  `json:"scope"`
	Subresource     string  `json:"subresource"`
	Verb            string  `json:"verb"`
	Version         string  `json:"version"`
	Le              float64 `json:"le"`
	CumulativeCount uint64  `json:"cumulative_count"`
}

type ApiserverRequestDurationSecondsCount struct {
	Component   string  `json:"component"`
	DryRun      string  `json:"dry_run"`
	Group       string  `json:"group"`
	Resource    string  `json:"resource"`
	Scope       string  `json:"scope"`
	Subresource string  `json:"subresource"`
	Verb        string  `json:"verb"`
	Version     string  `json:"version"`
	SampleCount float64 `json:"sample_count"`
}

type ApiserverRequestDurationSecondsSum struct {
	Component   string  `json:"component"`
	DryRun      string  `json:"dry_run"`
	Group       string  `json:"group"`
	Resource    string  `json:"resource"`
	Scope       string  `json:"scope"`
	Subresource string  `json:"subresource"`
	Verb        string  `json:"verb"`
	Version     string  `json:"version"`
	SampleSum   float64 `json:"sample_sum"`
}

/*
Maximal number of currently used inflight request limit of this apiserver per request kind in last second.
Stability Level:STABLE
Type: Gauge
*/
type ApiserverCurrentInflightRequests struct {
	RequestKind string  `json:"request_kind"`
	Gauge       float64 `json:"gauge"`
}

/*
Counter of policy levels for audit events (1 per request).
Stability Level:ALPHA
Type: Counter
*/
type ApiserverAuditLevelTotal struct {
	Level   string  `json:"level"`
	Counter float64 `json:"counter"`
}

/*
Number of goroutines that currently exist.
Stability Level:?
Type: Gauge
*/
type GoGoroutines struct {
	Gauge float64 `json:"gauge"`
}

/*
Number of OS threads created.
Stability Level:?
Type: Gauge
*/
type GoThreads struct {
	Gauge float64 `json:"gauge"`
}

/*
Etcd request latency in seconds for each operation and object type.
Stability Level:ALPHA
Type: Histogram
*/
type EtcdRequestDurationSecondsBucket struct {
	Operation       string  `json:"operation"`
	Type            string  `json:"type"`
	Le              float64 `json:"le"`
	CumulativeCount uint64  `json:"cumulative_count"`
}

type EtcdRequestDurationSecondsCount struct {
	Operation   string  `json:"operation"`
	Type        string  `json:"type"`
	SampleCount float64 `json:"sample_count"`
}

type EtcdRequestDurationSecondsSum struct {
	Operation string  `json:"operation"`
	Type      string  `json:"type"`
	SampleSum float64 `json:"sample_sum"`
}
