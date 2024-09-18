package common

const (
	//accessLogSelect = "SELECT * FROM access_log {where} order by start_time limit 2"

	AccessLogResource = "access-log"
	AccessLogSelect   = "SELECT region,customer_id,start_time,duration_str,traffic,rate_limit FROM access_log {where} order by start_time desc limit 2"
	AccessLogInsert   = "INSERT INTO access_log (" +
		"customer_id,start_time,duration_ms,duration_str,traffic," +
		"region,zone,sub_zone,service,instance_id,route_name," +
		"request_id,url,protocol,method,host,path,status_code,bytes_sent,status_flags," +
		"timeout,rate_limit,rate_burst,retry,retry_rate_limit,retry_rate_burst,failover) VALUES"

	CustomerIdName = "customer_id"
	StartTimeName  = "start_time"
	DurationName   = "duration_ms"
	TrafficName    = "traffic"
	CreatedTSName  = "created_ts"

	RegionName     = "region"
	ZoneName       = "zone"
	SubZoneName    = "sub_zone"
	HostName       = "host"
	InstanceIdName = "instance_id"

	RequestIdName = "request_id"
	RelatesToName = "relates_to"
	ProtocolName  = "protocol"
	MethodName    = "method"
	FromName      = "from"
	ToName        = "to"
	UriName       = "url"
	PathName      = "path"

	StatusCodeName = "status_code"
	EncodingName   = "encoding"
	BytesName      = "bytes"

	RouteName        = "route"
	RouteToName      = "route_to"
	RoutePercentName = "route_percent"
	RouteCodeName    = "rc"

	TimeoutName        = "timeout"
	RateLimitName      = "rate_limit"
	RateBurstName      = "rate_burst"
	ControllerCodeName = "cc"
)
