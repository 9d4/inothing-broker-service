package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// RabbitMQ Endpoints. How to read:
//
// ApiOverview		-->	/api/overview
//
// ApiCluster_name	--> /api/cluster-name
//
// Take a look at ApiNodesName. It has "%s" in it.
// It means we have to substitute the %s before using
// it to make request. We can use format(fmt) to do it.
// e.g. fmt.Sprintf(ApiNodesName, "nodename")
//
// More information about API can be found in web management.
// e.g. http://127.0.0.1:15672/api
const (
	ApiOverview = "/api/overview"

	ApiCluster_name = "/api/cluster-name"

	ApiNodes = "/api/nodes"

	// /api/nodes/<name>
	ApiNodesName = "/api/nodes/%s"

	ApiExtensions = "/api/extensions"

	ApiDefinitions = "/api/definitions"

	// /api/definitions/<vhost>
	ApiDefinitionsVhost = "/api/definitions/%s"

	ApiWhoami = "/api/whoami"

	// ------------------------------------------------------
	// AUTHS
	// ------------------------------------------------------

	ApiAuth = "/api/auth"

	// /api/auth/attempts/<node>
	ApiAuthAttemptsNode = "/api/auth/attempts/%s"

	// /api/auth/attempts/<node>/source
	ApiAuthAttemptsNodeSource = "/api/auth/%s/source"

	// ------------------------------------------------------

	// ------------------------------------------------------
	// BINDINGS
	// ------------------------------------------------------

	ApiBindings = "/api/bindings"

	// /api/bindings/<vhost>
	ApiBindingsVhost = "/api/bindings/%s"

	// /api/bindings/<vhost>/e/<exchange>/q/<queue>
	ApiBindingsVhostEExchangeQQueue = "/api/bindings/%s/e/%s/q/%s"

	// /api/bindings/<vhost>/e/<exchange>/q/<queue>/<props>
	ApiBindingsVhostEExchangeQQueueProps = "/api/bindings/%s/e/%s/q/%s/%s"

	// /api/bindings/<vhost>/e/<source>/e/<destination>
	ApiBindingsVhostESourceEDestination = "/api/bindings/%s/e/%s/e/%s"

	// /api/bindings/<vhost>/e/<source>/e/<destination>/<props>
	ApiBindingsVhostESourceEDestinationProps = "/api/bindings/%s/e/%s/e/%s/%s"

	// ------------------------------------------------------

	// ------------------------------------------------------
	// CONNECTIONS
	// ------------------------------------------------------
	ApiConnections = "/api/connections"

	// /api/connections/<name>
	ApiConnectionsName = "/api/connections/%s"

	// /api/connections/<name>/channels
	ApiConnectionsNameChannels = "/api/connections/%s/channels"

	ApiChannels = "/api/channels"

	// /api/channels/<channel>
	ApiChannelsChannel = "/api/channels/%s"

	ApiConsumers = "/api/consumers"

	// /api/consumers/<vhost>
	ApiConsumersVhost = "/api/consumers/%s"

	// ------------------------------------------------------

	// ------------------------------------------------------
	// EXCHANGES
	// ------------------------------------------------------

	ApiExchanges = "/api/exchanges"

	// /api/exchanges/<vhost>
	ApiExchangesVhost = "/api/exchanges/%s"

	// /api/exchanges/<vhost>/<name>
	ApiExchangesVhostName = "/api/exchanges/%s/%s"

	// /api/exchanges/<vhost>/<name>/bindings/source
	ApiExchangesVhostNameBindingsSource = "/api/exchanges/%s/%s/bindings/source"

	// /api/exchanges/<vhost>/<name>/bindings/destination
	ApiExchangesVhostNameBindingsDestination = "/api/exchanges/%s/%s/bindings/destination"

	// /api/exchanges/<vhost>/<name>/publish
	ApiExchangesVhostNamePublish = "/api/exchanges/%s/%s/publish"

	// ------------------------------------------------------

	// ------------------------------------------------------
	// FEDERATIONS
	// ------------------------------------------------------

	ApiFederation_links = "/api/federation-links"

	// /api/federation-links/<vhost>
	ApiFederation_linksVhost = "/api/federation-links/%s"

	// ------------------------------------------------------

	// ------------------------------------------------------
	// HEALTHS
	// ------------------------------------------------------

	// /api/aliveness-test/<vhost>
	ApiAliveness_testVhost = "/api/aliveness-test/%s"

	ApiHealthChecksAlarms = "/api/health/checks/alarms"

	ApiHealthChecksLocal_alarms = "/api/health/checks/local-alarms"

	// /api/health/checks/certificate-expiration/<within>/<unit>
	ApiHealthChecksCertificate_expirationWithinUnit = "/api/health/checks/certificate-expiration/%s/%s"

	// /api/health/checks/port-listener/<port>
	ApiHealthChecksPort_listenerPort = "/api/health/checks/port-listener/%s"

	// /api/health/checks/protocol-listener/<protocol>
	ApiHealthChecksProtocol_listenerProtocol = "/api/health/checks/protocol-listener/%s"

	ApiHealthChecksVirtual_hosts = "/api/health/checks/virtual-hosts"

	ApiHealthChecksNode_is_mirror_sync_critical = "/api/health/checks/node-is-mirror-sync-critical"

	ApiHealthChecksNode_is_quorum_critical = "/api/health/checks/node-is-quorum-critical"

	// ------------------------------------------------------

	// ------------------------------------------------------
	// PARAMETERS
	// ------------------------------------------------------

	ApiParameters = "/api/parameters"

	// /api/parameters/<component>
	ApiParametersComponent = "/api/parameters/%s"

	// /api/parameters/<component>/<vhost>
	ApiParametersComponentVhost = "/api/parameters/%s/%s"

	// /api/parameters/<component>/<vhost>/<name>
	ApiParametersComponentVhostName = "/api/parameters/%s/%s/%s"

	ApiGlobal_parameters = "/api/global-parameters"

	// /api/global-parameters/<name>
	ApiGlobal_parametersName = "/api/global-parameters/%s"

	// ------------------------------------------------------

	// ------------------------------------------------------
	// PERMISSIONS
	// ------------------------------------------------------

	ApiPermissions = "/api/permissions"

	// /api/permissions/<vhost>/<user>
	ApiPermissionsVhostUser = "/api/permissions/%s/%s"

	ApiTopic_permissions = "/api/topic-permissions"

	// /api/topic-permissions/<vhost>/<user>
	ApiTopic_permissionsVhostUser = "/api/topic-permissions/%s/%s"

	// ------------------------------------------------------

	// ------------------------------------------------------
	// POLICIES
	// ------------------------------------------------------

	ApiPolicies = "/api/policies"

	// /api/policies/<vhost>
	ApiPoliciesVhost = "/api/policies/%s"

	// /api/policies/<vhost>/<name>
	ApiPoliciesVhostName = "/api/policies/%s/%s"
	ApiOperator_policies = "/api/operator-policies"

	// /api/operator-policies/<vhost>
	ApiOperator_policiesVhost = "/api/operator-policies/%s"

	// /api/operator-policies/<vhost>/<name>
	ApiOperator_policiesVhostName = "/api/operator-policies/%s/%s"

	// ------------------------------------------------------

	// ------------------------------------------------------
	// QUEUES
	// ------------------------------------------------------

	ApiQueues = "/api/queues"

	// /api/queues/<vhost>/<name>
	ApiQueuesVhost = "/api/queues/%s"

	// /api/queues/<vhost>/<name>
	ApiQueuesVhostName = "/api/queues/%s/%s"

	// /api/queues/<vhost>/<name>/bindings
	ApiQueuesVhostNameBindings = "/api/queues/%s/%s/bindings"

	// /api/queues/<vhost>/<name>/contents
	ApiQueuesVhostNameContents = "/api/queues/%s/%s/contents"

	// /api/queues/<vhost>/<name>/actions
	ApiQueuesVhostNameActions = "/api/queues/%s/%s/actions"

	// /api/queues/<vhost>/<name>/get
	ApiQueuesVhostNameGet = "/api/queues/%s/%s/get"

	ApiRebalanceQueues = "/api/rebalance/queues"

	// ------------------------------------------------------

	// ------------------------------------------------------
	// USERS
	// ------------------------------------------------------

	ApiUsers                    = "/api/users"
	ApiUsersWithout_permissions = "/api/users/without-permissions"
	ApiUsersBulk_delete         = "/api/users/bulk-delete"

	// /api/users/<name>
	ApiUsersName = "/api/users/%s"

	// /api/users/<user>/permissions
	ApiUsersUserPermissions = "/api/users/%s/permissions"

	// ------------------------------------------------------

	// /api/users/<user>/topic-permissions
	ApiUsersUserTopic_permissions = "/api/users/%s/topic-permissions"

	// ------------------------------------------------------
	// USERS LIMITS
	// ------------------------------------------------------

	ApiUser_limits = "/api/user-limits"

	// /api/user-limits/<user>
	ApiUser_limitsUser = "/api/user-limits/%s"

	// /api/user-limits/<user>/<name>
	ApiUser_limitsUserName = "/api/user-limits/%s/%s"

	// ------------------------------------------------------

	// ------------------------------------------------------
	// VHOSTS
	// ------------------------------------------------------

	ApiVhosts = "/api/vhosts"

	// /api/vhosts/<name>
	ApiVhostsName = "/api/vhosts/%s"

	// /api/vhosts/<vhost>/channels
	ApiVhostsVhostChannels = "/api/vhosts/%s/channels"

	// /api/vhosts/<vhost>/connections
	ApiVhostsVhostConnections = "/api/vhosts/%s/connections"

	// /api/vhosts/<name>/permissions
	ApiVhostsNamePermissions = "/api/vhosts/%s/permissions"

	// /api/vhosts/<name>/start/<node>
	ApiVhostsNameStartNode = "/api/vhosts/%s/start/%s"

	// /api/vhosts/<name>/topic-permissions
	ApiVhostsNameTopic_permissions = "/api/vhosts/%s/topic-permissions"

	ApiVhost_limits = "/api/vhost-limits"

	// /api/vhost-limits/<vhost>
	ApiVhost_limitsVhost = "/api/vhost-limits/%s"

	// /api/vhost-limits/<vhost>/<name>
	ApiVhost_limitsVhostName = "/api/vhost-limits/%s/%s"

	// ------------------------------------------------------
)

const AMQP_TOPIC_EXCHANGE = "amq.topic"

// Request Create new user to RabbitMQ API
func ReqCreateNewUser(username string, password string) (*http.Response, error) {
	bodyStr := fmt.Sprintf(`{"password_hash":"%s","tags":"management"}`, hashPasswordBase64(password))
	body := bytes.NewBufferString(bodyStr)

	res, err := Req("PUT", fmt.Sprintf(ApiUsersName, username), body)
	return res, err
}

// Change vhost permission
func ReqChmodVhost(vhost string, user string, body io.Reader) (*http.Response, error) {
	path := fmt.Sprintf(ApiPermissionsVhostUser, normalizeNames(vhost), normalizeNames(user))

	return Req("PUT", path, body)
}

// Change Topic Permission
func ReqChmodTopic(vhost string, user string, body io.Reader) (*http.Response, error) {
	path := fmt.Sprintf(ApiTopic_permissionsVhostUser, normalizeNames(vhost), normalizeNames(user))

	return Req("PUT", path, body)
}

// Create queue in vhost
func ReqCreateNewQueueInVhost(vhost string, queue string) (*http.Response, error) {
	path := fmt.Sprintf(ApiQueuesVhostName, normalizeNames(vhost), normalizeNames(queue))
	bodyStr := `{"auto_delete":false,"durable":true}`
	body := bytes.NewBufferString(bodyStr)

	return Req("PUT", path, body)
}

// Create queue in default vhost
func ReqCreateNewQueue(queue string) (*http.Response, error) {
	return ReqCreateNewQueueInVhost(RABBITMQ_VHOST, queue)
}

// Bind Exchange, Routing Key to Queue in default vhost
func ReqBindExchangeRoutingToQueue(exchange string, toQueue string, routingKey string) (*http.Response, error) {
	// ApiBindingsVhostEExchangeQQueue
	path := fmt.Sprintf(ApiBindingsVhostEExchangeQQueue, normalizeNames(RABBITMQ_VHOST), normalizeNames(exchange), normalizeNames(toQueue))
	bodyStr := fmt.Sprintf(`{"routing_key":"%s"}`, routingKey)
	body := bytes.NewBufferString(bodyStr)

	return Req("POST", path, body)
}

// Bind amq.topic Exchange from routing key to Queue
// Same as #ReqBindExchangeRoutingToQueue, but automatically select
// amq.topic exchange
func ReqBindAmqRoutingQueue(toQueue string, routingKey string) (*http.Response, error) {
	return ReqBindExchangeRoutingToQueue(AMQP_TOPIC_EXCHANGE, toQueue, routingKey)
}
