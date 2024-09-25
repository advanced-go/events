package testrsc

const (
	log1BasePath = "file:///f:/files/log1"

	LOG1IngressEntry      = log1BasePath + "/ingress-entry.json"
	LOG1IngressGetAllReq  = log1BasePath + "/ingress-get-all-req.txt"
	LOG1IngressGetAllResp = log1BasePath + "/ingress-get-all-resp.txt"

	LOG1EgressEntry      = log1BasePath + "/egress-entry.json"
	LOG1EgressGetAllReq  = log1BasePath + "/egress-get-all-req.txt"
	LOG1EgressGetAllResp = log1BasePath + "/egress-get-all-resp.txt"

	log2BasePath = "file:///f:/files/log2"

	LOG2IngressEntry      = log2BasePath + "/ingress-entry.json"
	LOG2IngressGetAllReq  = log2BasePath + "/ingress-get-all-req.txt"
	LOG2IngressGetAllResp = log2BasePath + "/ingress-get-all-resp.txt"

	LOG2EgressEntry      = log2BasePath + "/egress-entry.json"
	LOG2EgressGetAllReq  = log2BasePath + "/egress-get-all-req.txt"
	LOG2EgressGetAllResp = log2BasePath + "/egress-get-all-resp.txt"
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

	TS1GetReq  = ts1BasePath + "/egress-get-all-req.txt"
	TS1GetResp = ts1BasePath + "/egress-get-all-resp.txt"

	ts2BasePath         = "file:///f:/files/timeseries2"
	TS2IngressEntry     = ts2BasePath + "/ingress-entry.json"
	TS2IngressEntryTest = ts2BasePath + "/ingress-percentile-threshold-test.json"
	TS2EgressEntry      = ts2BasePath + "/egress-entry.json"
	TS2EgressEntryTest  = ts2BasePath + "/percentile-threshold-test.json"


*/
