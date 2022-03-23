package main

const (
	ApiOverview                                     = "/api/overview"
	ApiCluster_name                                 = "/api/cluster_name"
	ApiNodes                                        = "/api/nodes"
	ApiNodesName                                    = "/api/nodes/name"
	ApiExtensions                                   = "/api/extensions"
	ApiDefinitions                                  = "/api/definitions"
	ApiDefinitionsVhost                             = "/api/definitions/vhost"
	ApiConnections                                  = "/api/connections"
	ApiVhostsVhostConnections                       = "/api/vhosts/vhost/connections"
	ApiConnectionsName                              = "/api/connections/name"
	ApiConnectionsNameChannels                      = "/api/connections/name/channels"
	ApiChannels                                     = "/api/channels"
	ApiVhostsVhostChannels                          = "/api/vhosts/vhost/channels"
	ApiChannelsChannel                              = "/api/channels/channel"
	ApiConsumers                                    = "/api/consumers"
	ApiConsumersVhost                               = "/api/consumers/vhost"
	ApiExchanges                                    = "/api/exchanges"
	ApiExchangesVhost                               = "/api/exchanges/vhost"
	ApiExchangesVhostName                           = "/api/exchanges/vhost/name"
	ApiExchangesVhostNameBindingsSource             = "/api/exchanges/vhost/name/bindings/source"
	ApiExchangesVhostNameBindingsDestination        = "/api/exchanges/vhost/name/bindings/destination"
	ApiExchangesVhostNamePublish                    = "/api/exchanges/vhost/name/publish"
	ApiQueues                                       = "/api/queues"
	ApiQueuesVhost                                  = "/api/queues/vhost"
	ApiQueuesVhostName                              = "/api/queues/vhost/name"
	ApiQueuesVhostNameBindings                      = "/api/queues/vhost/name/bindings"
	ApiQueuesVhostNameContents                      = "/api/queues/vhost/name/contents"
	ApiQueuesVhostNameActions                       = "/api/queues/vhost/name/actions"
	ApiQueuesVhostNameGet                           = "/api/queues/vhost/name/get"
	ApiBindings                                     = "/api/bindings"
	ApiBindingsVhost                                = "/api/bindings/vhost"
	ApiBindingsVhostEExchangeQQueue                 = "/api/bindings/vhost/e/exchange/q/queue"
	ApiBindingsVhostEExchangeQQueueProps            = "/api/bindings/vhost/e/exchange/q/queue/props"
	ApiBindingsVhostESourceEDestination             = "/api/bindings/vhost/e/source/e/destination"
	ApiBindingsVhostESourceEDestinationProps        = "/api/bindings/vhost/e/source/e/destination/props"
	ApiVhosts                                       = "/api/vhosts"
	ApiVhostsName                                   = "/api/vhosts/name"
	ApiVhostsNamePermissions                        = "/api/vhosts/name/permissions"
	ApiVhostsNameTopic_permissions                  = "/api/vhosts/name/topic_permissions"
	ApiVhostsNameStartNode                          = "/api/vhosts/name/start/node"
	ApiUsers                                        = "/api/users"
	ApiUsersWithout_permissions                     = "/api/users/without_permissions"
	ApiUsersBulk_delete                             = "/api/users/bulk_delete"
	ApiUsersName                                    = "/api/users/name"
	ApiUsersUserPermissions                         = "/api/users/user/permissions"
	ApiUsersUserTopic_permissions                   = "/api/users/user/topic_permissions"
	ApiUser_limits                                  = "/api/user_limits"
	ApiUser_limitsUser                              = "/api/user_limits/user"
	ApiUser_limitsUserName                          = "/api/user_limits/user/name"
	ApiWhoami                                       = "/api/whoami"
	ApiPermissions                                  = "/api/permissions"
	ApiPermissionsVhostUser                         = "/api/permissions/vhost/user"
	ApiTopic_permissions                            = "/api/topic_permissions"
	ApiTopic_permissionsVhostUser                   = "/api/topic_permissions/vhost/user"
	ApiParameters                                   = "/api/parameters"
	ApiParametersComponent                          = "/api/parameters/component"
	ApiParametersComponentVhost                     = "/api/parameters/component/vhost"
	ApiParametersComponentVhostName                 = "/api/parameters/component/vhost/name"
	ApiGlobal_parameters                            = "/api/global_parameters"
	ApiGlobal_parametersName                        = "/api/global_parameters/name"
	ApiPolicies                                     = "/api/policies"
	ApiPoliciesVhost                                = "/api/policies/vhost"
	ApiPoliciesVhostName                            = "/api/policies/vhost/name"
	ApiOperator_policies                            = "/api/operator_policies"
	ApiOperator_policiesVhost                       = "/api/operator_policies/vhost"
	ApiOperator_policiesVhostName                   = "/api/operator_policies/vhost/name"
	ApiAliveness_testVhost                          = "/api/aliveness_test/vhost"
	ApiHealthChecksAlarms                           = "/api/health/checks/alarms"
	ApiHealthChecksLocal_alarms                     = "/api/health/checks/local_alarms"
	ApiHealthChecksCertificate_expirationWithinUnit = "/api/health/checks/certificate_expiration/within/unit"
	ApiHealthChecksPort_listenerPort                = "/api/health/checks/port_listener/port"
	ApiHealthChecksProtocol_listenerProtocol        = "/api/health/checks/protocol_listener/protocol"
	ApiHealthChecksVirtual_hosts                    = "/api/health/checks/virtual_hosts"
	ApiHealthChecksNode_is_mirror_sync_critical     = "/api/health/checks/node_is_mirror_sync_critical"
	ApiHealthChecksNode_is_quorum_critical          = "/api/health/checks/node_is_quorum_critical"
	ApiVhost_limits                                 = "/api/vhost_limits"
	ApiVhost_limitsVhost                            = "/api/vhost_limits/vhost"
	ApiVhost_limitsVhostName                        = "/api/vhost_limits/vhost/name"
	ApiAuth                                         = "/api/auth"
	ApiRebalanceQueues                              = "/api/rebalance/queues"
	ApiFederation_linksVhost                        = "/api/federation_links/vhost"
	ApiAuthAttemptsNode                             = "/api/auth/attempts/node"
)
