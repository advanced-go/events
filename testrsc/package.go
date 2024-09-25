package testrsc

const (
	log1BasePath = "file:///f:/files/log1"

	LOG1EgressEntry      = log1BasePath + "/egress-entry.json"
	LOG1EgressEntryTest  = log1BasePath + "/egress-entry-test.json"
	LOG1IngressEntry     = log1BasePath + "/ingress-entry.json"
	LOG1IngressEntryTest = log1BasePath + "/ingress-entry-test.json"

	LOG1GetReq  = log1BasePath + "/get-req.txt"
	LOG1GetResp = log1BasePath + "/get-resp.txt"

	log2BasePath         = "file:///f:/files/log2"
	LOG2EgressEntryTest  = log2BasePath + "/egress-entry-test.json"
	LOG2IngressEntryTest = log2BasePath + "/ingress-entry-test.json"
)

const (
	ts1BasePath = "file:///f:/files/timeseries1"

	TS1PercentileThresholdTest = ts1BasePath + "/percentile-threshold-test.json"
	TS1StatusCodeThresholdTest = ts1BasePath + "/stat-code-threshold-test.json"

	// test

)

/*
	ts1BasePath = "file:///f:/files/timeseries1"

	TS1EgressEntry      = ts1BasePath + "/egress-entry.json"
	TS1EgressEntryTest  = ts1BasePath + "/egress-percentile-threshold-test.json"
	TS1IngressEntry     = ts1BasePath + "/ingress-entry.json"
	TS1IngressEntryTest = ts1BasePath + "/ingress-percentile-threshold-test.json"

	TS1GetReq  = ts1BasePath + "/get-req.txt"
	TS1GetResp = ts1BasePath + "/get-resp.txt"

	ts2BasePath         = "file:///f:/files/timeseries2"
	TS2IngressEntry     = ts2BasePath + "/ingress-entry.json"
	TS2IngressEntryTest = ts2BasePath + "/ingress-percentile-threshold-test.json"
	TS2EgressEntry      = ts2BasePath + "/egress-entry.json"
	TS2EgressEntryTest  = ts2BasePath + "/percentile-threshold-test.json"


*/
