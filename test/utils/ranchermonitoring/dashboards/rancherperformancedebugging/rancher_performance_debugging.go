package rancherperformancedebugging

const (
	UID                                              = "rancher-performance-debugging"
	UIDFormatted1                                    = "rancher-performance-debugging-1"
	UIDFormatted2                                    = "rancher-performance-debugging-2"
	UIDFormatted3                                    = "rancher-performance-debugging-3"
	HandlerAverageExecutionTimesExpr                 = "topk(20,sum by (handler_name) (rate(lasso_controller_reconcile_time_seconds_sum[$__rate_interval])) / sum by (handler_name) (rate(lasso_controller_reconcile_time_seconds_count[$__rate_interval])))"
	RancherAPIAverageRequestTimesExpr                = "topk(20,sum by (resource, method, code) (rate(steve_api_request_time_sum{resource!=\"subscribe\"}[$__rate_interval])) / sum by (resource, method, code) (rate(steve_api_request_time_count{resource!=\"subscribe\"}[$__rate_interval])))"
	SubscribeAverageRequestTimesExpr                 = "rate(steve_api_request_time_sum{resource=\"subscribe\"}[$__rate_interval]) / rate(steve_api_request_time_count{resource=\"subscribe\"}[$__rate_interval])"
	LassoControllerWorkQueueDepthExpr                = "topk(20,workqueue_depth)"
	NumberOfRancherRequestsExpr                      = "topk(20,sum by (id, resource, method, code) (steve_api_total_requests))"
	NumberOfFailedRancherRequestsExpr                = "topk(20,sum by (id, resource, method) (steve_api_total_requests{code!=\"200\",code!=\"201\"}))"
	K8sProxyStoreAverageRequestTimesExpr             = "topk(20,sum by (resource, method, code) (rate(k8s_proxy_store_request_time_sum[5m])) / sum by (resource, method, code) (rate(k8s_proxy_store_request_time_count[5m])))"
	K8sProxyClientAverageRequestTimesExpr            = "topk(20,sum by (resource, method, code) (rate(k8s_proxy_client_request_time_sum[5m])) / sum by (resource, method, code) (rate(k8s_proxy_client_request_time_count[5m])))"
	CachedObjectsByGroupVersionKindExpr              = "topk(20,lasso_controller_total_cached_object)"
	LassoHandlerExecutionsExpr                       = "topk(20,sum by (handler_name) (lasso_controller_total_handler_execution))"
	LassoHandlerExecutionsByClusterNameExpr          = "topk(20, sum by (handler_name,controller_name) (increase(lasso_controller_total_handler_execution[2m])))"
	LassoHandlerExecutionsWithErrorExpr              = "topk(20,sum by (handler_name) (lasso_controller_total_handler_execution{has_error=\"true\"}))"
	LassoHandlerExecutionsWithErrorByClusterNameExpr = "topk(20,sum by (handler_name,controller_name) (increase(lasso_controller_total_handler_execution{has_error=\"true\"}[2m])))"
	DataTransmittedByRemoteDialerSessionsExpr        = "topk(20,session_server_total_transmit_bytes)"
	ErrorsByRemoteDialerSessionsExpr                 = "session_server_total_transmit_error_bytes"
	RemoteDialerConnectionsRemovedExpr               = "session_server_total_remove_connections"
	RemoteDialerConnectionsAddedByClientExpr         = "session_server_total_add_connections"
)
