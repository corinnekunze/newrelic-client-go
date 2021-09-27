// Code generated by tutone: DO NOT EDIT
package entities

import (
	"context"

	"github.com/newrelic/newrelic-client-go/pkg/common"
	"github.com/newrelic/newrelic-client-go/pkg/errors"
)

// Adds the provided tags to your specified entity, without deleting existing ones.
//  The maximum number of tag-values per entity is 100; if the sum of existing and new tag-values if over the limit this mutation will fail.
//
//  For details and mutation examples, visit [our docs](https://docs.newrelic.com/docs/apis/nerdgraph/examples/nerdgraph-tagging-api-tutorial).
func (a *Entities) TaggingAddTagsToEntity(
	gUID common.EntityGUID,
	tags []TaggingTagInput,
) (*TaggingMutationResult, error) {
	return a.TaggingAddTagsToEntityWithContext(context.Background(),
		gUID,
		tags,
	)
}

// Adds the provided tags to your specified entity, without deleting existing ones.
//  The maximum number of tag-values per entity is 100; if the sum of existing and new tag-values if over the limit this mutation will fail.
//
//  For details and mutation examples, visit [our docs](https://docs.newrelic.com/docs/apis/nerdgraph/examples/nerdgraph-tagging-api-tutorial).
func (a *Entities) TaggingAddTagsToEntityWithContext(
	ctx context.Context,
	gUID common.EntityGUID,
	tags []TaggingTagInput,
) (*TaggingMutationResult, error) {

	resp := TaggingAddTagsToEntityQueryResponse{}
	vars := map[string]interface{}{
		"guid": gUID,
		"tags": tags,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, TaggingAddTagsToEntityMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.TaggingMutationResult, nil
}

type TaggingAddTagsToEntityQueryResponse struct {
	TaggingMutationResult TaggingMutationResult `json:"TaggingAddTagsToEntity"`
}

const TaggingAddTagsToEntityMutation = `mutation(
	$guid: EntityGuid!,
	$tags: [TaggingTagInput!]!,
) { taggingAddTagsToEntity(
	guid: $guid,
	tags: $tags,
) {
	errors {
		message
		type
	}
} }`

// Delete specific tag keys from the entity.
//
//  For details and mutation examples, visit [our docs](https://docs.newrelic.com/docs/apis/nerdgraph/examples/nerdgraph-tagging-api-tutorial).
func (a *Entities) TaggingDeleteTagFromEntity(
	gUID common.EntityGUID,
	tagKeys []string,
) (*TaggingMutationResult, error) {
	return a.TaggingDeleteTagFromEntityWithContext(context.Background(),
		gUID,
		tagKeys,
	)
}

// Delete specific tag keys from the entity.
//
//  For details and mutation examples, visit [our docs](https://docs.newrelic.com/docs/apis/nerdgraph/examples/nerdgraph-tagging-api-tutorial).
func (a *Entities) TaggingDeleteTagFromEntityWithContext(
	ctx context.Context,
	gUID common.EntityGUID,
	tagKeys []string,
) (*TaggingMutationResult, error) {

	resp := TaggingDeleteTagFromEntityQueryResponse{}
	vars := map[string]interface{}{
		"guid":    gUID,
		"tagKeys": tagKeys,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, TaggingDeleteTagFromEntityMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.TaggingMutationResult, nil
}

type TaggingDeleteTagFromEntityQueryResponse struct {
	TaggingMutationResult TaggingMutationResult `json:"TaggingDeleteTagFromEntity"`
}

const TaggingDeleteTagFromEntityMutation = `mutation(
	$guid: EntityGuid!,
	$tagKeys: [String!]!,
) { taggingDeleteTagFromEntity(
	guid: $guid,
	tagKeys: $tagKeys,
) {
	errors {
		message
		type
	}
} }`

// Delete specific tag key-values from the entity.
//
//  For details and mutation examples, visit [our docs](https://docs.newrelic.com/docs/apis/nerdgraph/examples/nerdgraph-tagging-api-tutorial).
func (a *Entities) TaggingDeleteTagValuesFromEntity(
	gUID common.EntityGUID,
	tagValues []TaggingTagValueInput,
) (*TaggingMutationResult, error) {
	return a.TaggingDeleteTagValuesFromEntityWithContext(context.Background(),
		gUID,
		tagValues,
	)
}

// Delete specific tag key-values from the entity.
//
//  For details and mutation examples, visit [our docs](https://docs.newrelic.com/docs/apis/nerdgraph/examples/nerdgraph-tagging-api-tutorial).
func (a *Entities) TaggingDeleteTagValuesFromEntityWithContext(
	ctx context.Context,
	gUID common.EntityGUID,
	tagValues []TaggingTagValueInput,
) (*TaggingMutationResult, error) {

	resp := TaggingDeleteTagValuesFromEntityQueryResponse{}
	vars := map[string]interface{}{
		"guid":      gUID,
		"tagValues": tagValues,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, TaggingDeleteTagValuesFromEntityMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.TaggingMutationResult, nil
}

type TaggingDeleteTagValuesFromEntityQueryResponse struct {
	TaggingMutationResult TaggingMutationResult `json:"TaggingDeleteTagValuesFromEntity"`
}

const TaggingDeleteTagValuesFromEntityMutation = `mutation(
	$guid: EntityGuid!,
	$tagValues: [TaggingTagValueInput!]!,
) { taggingDeleteTagValuesFromEntity(
	guid: $guid,
	tagValues: $tagValues,
) {
	errors {
		message
		type
	}
} }`

// Replaces the entity's entire set of tags with the provided tag set.
//  The maximum number of tag-values per entity is 100; if more than 100 tag-values are provided this mutation will fail.
//
//  For details and mutation examples, visit [our docs](https://docs.newrelic.com/docs/apis/nerdgraph/examples/nerdgraph-tagging-api-tutorial).
func (a *Entities) TaggingReplaceTagsOnEntity(
	gUID common.EntityGUID,
	tags []TaggingTagInput,
) (*TaggingMutationResult, error) {
	return a.TaggingReplaceTagsOnEntityWithContext(context.Background(),
		gUID,
		tags,
	)
}

// Replaces the entity's entire set of tags with the provided tag set.
//  The maximum number of tag-values per entity is 100; if more than 100 tag-values are provided this mutation will fail.
//
//  For details and mutation examples, visit [our docs](https://docs.newrelic.com/docs/apis/nerdgraph/examples/nerdgraph-tagging-api-tutorial).
func (a *Entities) TaggingReplaceTagsOnEntityWithContext(
	ctx context.Context,
	gUID common.EntityGUID,
	tags []TaggingTagInput,
) (*TaggingMutationResult, error) {

	resp := TaggingReplaceTagsOnEntityQueryResponse{}
	vars := map[string]interface{}{
		"guid": gUID,
		"tags": tags,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, TaggingReplaceTagsOnEntityMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.TaggingMutationResult, nil
}

type TaggingReplaceTagsOnEntityQueryResponse struct {
	TaggingMutationResult TaggingMutationResult `json:"TaggingReplaceTagsOnEntity"`
}

const TaggingReplaceTagsOnEntityMutation = `mutation(
	$guid: EntityGuid!,
	$tags: [TaggingTagInput!]!,
) { taggingReplaceTagsOnEntity(
	guid: $guid,
	tags: $tags,
) {
	errors {
		message
		type
	}
} }`

// Fetch a list of entities.
//
// You can fetch a max of 25 entities in one query.
//
// For more details on entities, visit our [entity docs](https://docs.newrelic.com/docs/apis/graphql-api/tutorials/use-new-relic-graphql-api-query-entities).
func (a *Entities) GetEntities(
	gUIDs []common.EntityGUID,
) (*[]EntityInterface, error) {
	return a.GetEntitiesWithContext(context.Background(),
		gUIDs,
	)
}

// Fetch a list of entities.
//
// You can fetch a max of 25 entities in one query.
//
// For more details on entities, visit our [entity docs](https://docs.newrelic.com/docs/apis/graphql-api/tutorials/use-new-relic-graphql-api-query-entities).
func (a *Entities) GetEntitiesWithContext(
	ctx context.Context,
	gUIDs []common.EntityGUID,
) (*[]EntityInterface, error) {

	resp := entitiesResponse{}
	vars := map[string]interface{}{
		"guids": gUIDs,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, getEntitiesQuery, vars, &resp); err != nil {
		return nil, err
	}

	if len(resp.Actor.Entities) == 0 {
		return nil, errors.NewNotFound("")
	}

	return &resp.Actor.Entities, nil
}

const getEntitiesQuery = `query(
	$guids: [EntityGuid]!,
) { actor { entities(
	guids: $guids,
) {
	account {
		id
		name
		reportingEventTypes
	}
	accountId
	alertSeverity
	domain
	entityType
	goldenMetrics {
		context {
			account
			guid
		}
		metrics {
			name
			query
			title
		}
	}
	goldenTags {
		context {
			account
			guid
		}
		tags {
			key
		}
	}
	guid
	indexedAt
	name
	permalink
	recentAlertViolations {
		agentUrl
		alertSeverity
		closedAt
		label
		level
		openedAt
		violationId
		violationUrl
	}
	relatedEntities {
		nextCursor
		results {
			__typename
			createdAt
			type
			... on EntityRelationshipDetectedEdge {
				__typename
			}
			... on EntityRelationshipUserDefinedEdge {
				__typename
			}
		}
	}
	relationships {
		source {
			accountId
			entityType
			guid
		}
		target {
			accountId
			entityType
			guid
		}
		type
	}
	reporting
	serviceLevel {
		indicators {
			createdAt
			description
			entityGuid
			id
			name
			updatedAt
		}
	}
	tags {
		key
		values
	}
	tagsWithMetadata {
		key
		values {
			mutable
			value
		}
	}
	type
	... on ApmApplicationEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		apmBrowserSummary {
			ajaxRequestThroughput
			ajaxResponseTimeAverage
			jsErrorRate
			pageLoadThroughput
			pageLoadTimeAverage
		}
		apmSummary {
			apdexScore
			errorRate
			hostCount
			instanceCount
			nonWebResponseTimeAverage
			nonWebThroughput
			responseTimeAverage
			throughput
			webResponseTimeAverage
			webThroughput
		}
		applicationId
		deployments {
			changelog
			description
			permalink
			revision
			timestamp
			user
		}
		language
		metricNormalizationRules {
			action
			applicationGuid
			applicationName
			createdAt
			enabled
			evalOrder
			id
			matchExpression
			notes
			replacement
			terminateChain
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		runningAgentVersions {
			maxVersion
			minVersion
		}
		settings {
			apdexTarget
			serverSideConfig
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on ApmDatabaseInstanceEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		host
		portOrPath
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
		vendor
	}
	... on ApmExternalServiceEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		externalSummary {
			responseTimeAverage
			throughput
		}
		host
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on BrowserApplicationEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		agentInstallType
		applicationId
		browserSummary {
			ajaxRequestThroughput
			ajaxResponseTimeAverage
			jsErrorRate
			pageLoadThroughput
			pageLoadTimeAverage
			pageLoadTimeMedian
			spaResponseTimeAverage
			spaResponseTimeMedian
		}
		metricNormalizationRules {
			action
			applicationGuid
			applicationName
			createdAt
			enabled
			evalOrder
			id
			matchExpression
			notes
			replacement
			terminateChain
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		runningAgentVersions {
			maxVersion
			minVersion
		}
		servingApmApplicationId
		settings {
			apdexTarget
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on DashboardEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		createdAt
		dashboardParentGuid
		description
		owner {
			email
			userId
		}
		pages {
			createdAt
			description
			guid
			name
			updatedAt
		}
		permissions
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
		updatedAt
	}
	... on ExternalEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on GenericEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on GenericInfrastructureEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		integrationTypeCode
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on InfrastructureAwsLambdaFunctionEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		integrationTypeCode
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		runtime
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on InfrastructureHostEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		hostSummary {
			cpuUtilizationPercent
			diskUsedPercent
			memoryUsedPercent
			networkReceiveRate
			networkTransmitRate
			servicesCount
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on MobileApplicationEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		applicationId
		metricNormalizationRules {
			action
			applicationGuid
			applicationName
			createdAt
			enabled
			evalOrder
			id
			matchExpression
			notes
			replacement
			terminateChain
		}
		mobileSummary {
			appLaunchCount
			crashCount
			crashRate
			httpErrorRate
			httpRequestCount
			httpRequestRate
			httpResponseTimeAverage
			mobileSessionCount
			networkFailureRate
			usersAffectedCount
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on SecureCredentialEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		description
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		secureCredentialId
		secureCredentialSummary {
			failingMonitorCount
			monitorCount
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
		updatedAt
	}
	... on SyntheticMonitorEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		monitorId
		monitorSummary {
			locationsFailing
			locationsRunning
			status
			successRate
		}
		monitorType
		monitoredUrl
		period
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on ThirdPartyServiceEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on UnavailableEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on WorkloadEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		createdAt
		createdByUser {
			email
			gravatar
			id
			name
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
		updatedAt
		workloadStatus {
			description
			statusSource
			statusValue
			summary
		}
	}
} } }`

// Fetch a single entity.
//
// For more details on entities, visit our [entity docs](https://docs.newrelic.com/docs/apis/graphql-api/tutorials/use-new-relic-graphql-api-query-entities).
func (a *Entities) GetEntity(
	gUID common.EntityGUID,
) (*EntityInterface, error) {
	return a.GetEntityWithContext(context.Background(),
		gUID,
	)
}

// Fetch a single entity.
//
// For more details on entities, visit our [entity docs](https://docs.newrelic.com/docs/apis/graphql-api/tutorials/use-new-relic-graphql-api-query-entities).
func (a *Entities) GetEntityWithContext(
	ctx context.Context,
	gUID common.EntityGUID,
) (*EntityInterface, error) {

	resp := entityResponse{}
	vars := map[string]interface{}{
		"guid": gUID,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, getEntityQuery, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.Actor.Entity, nil
}

const getEntityQuery = `query(
	$guid: EntityGuid!,
) { actor { entity(
	guid: $guid,
) {
	account {
		id
		name
		reportingEventTypes
	}
	accountId
	alertSeverity
	domain
	entityType
	goldenMetrics {
		context {
			account
			guid
		}
		metrics {
			name
			query
			title
		}
	}
	goldenTags {
		context {
			account
			guid
		}
		tags {
			key
		}
	}
	guid
	indexedAt
	name
	permalink
	recentAlertViolations {
		agentUrl
		alertSeverity
		closedAt
		label
		level
		openedAt
		violationId
		violationUrl
	}
	relatedEntities {
		nextCursor
		results {
			__typename
			createdAt
			type
			... on EntityRelationshipDetectedEdge {
				__typename
			}
			... on EntityRelationshipUserDefinedEdge {
				__typename
			}
		}
	}
	relationships {
		source {
			accountId
			entityType
			guid
		}
		target {
			accountId
			entityType
			guid
		}
		type
	}
	reporting
	serviceLevel {
		indicators {
			createdAt
			description
			entityGuid
			id
			name
			updatedAt
		}
	}
	tags {
		key
		values
	}
	tagsWithMetadata {
		key
		values {
			mutable
			value
		}
	}
	type
	... on ApmApplicationEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		apmBrowserSummary {
			ajaxRequestThroughput
			ajaxResponseTimeAverage
			jsErrorRate
			pageLoadThroughput
			pageLoadTimeAverage
		}
		apmSummary {
			apdexScore
			errorRate
			hostCount
			instanceCount
			nonWebResponseTimeAverage
			nonWebThroughput
			responseTimeAverage
			throughput
			webResponseTimeAverage
			webThroughput
		}
		applicationId
		deployments {
			changelog
			description
			permalink
			revision
			timestamp
			user
		}
		language
		metricNormalizationRules {
			action
			applicationGuid
			applicationName
			createdAt
			enabled
			evalOrder
			id
			matchExpression
			notes
			replacement
			terminateChain
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		runningAgentVersions {
			maxVersion
			minVersion
		}
		settings {
			apdexTarget
			serverSideConfig
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on ApmDatabaseInstanceEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		host
		portOrPath
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
		vendor
	}
	... on ApmExternalServiceEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		externalSummary {
			responseTimeAverage
			throughput
		}
		host
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on BrowserApplicationEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		agentInstallType
		applicationId
		browserSummary {
			ajaxRequestThroughput
			ajaxResponseTimeAverage
			jsErrorRate
			pageLoadThroughput
			pageLoadTimeAverage
			pageLoadTimeMedian
			spaResponseTimeAverage
			spaResponseTimeMedian
		}
		metricNormalizationRules {
			action
			applicationGuid
			applicationName
			createdAt
			enabled
			evalOrder
			id
			matchExpression
			notes
			replacement
			terminateChain
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		runningAgentVersions {
			maxVersion
			minVersion
		}
		servingApmApplicationId
		settings {
			apdexTarget
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on DashboardEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		createdAt
		dashboardParentGuid
		description
		owner {
			email
			userId
		}
		pages {
			createdAt
			description
			guid
			name
			updatedAt
		}
		permissions
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
		updatedAt
	}
	... on ExternalEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on GenericEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on GenericInfrastructureEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		integrationTypeCode
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on InfrastructureAwsLambdaFunctionEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		integrationTypeCode
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		runtime
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on InfrastructureHostEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		hostSummary {
			cpuUtilizationPercent
			diskUsedPercent
			memoryUsedPercent
			networkReceiveRate
			networkTransmitRate
			servicesCount
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on MobileApplicationEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		applicationId
		metricNormalizationRules {
			action
			applicationGuid
			applicationName
			createdAt
			enabled
			evalOrder
			id
			matchExpression
			notes
			replacement
			terminateChain
		}
		mobileSummary {
			appLaunchCount
			crashCount
			crashRate
			httpErrorRate
			httpRequestCount
			httpRequestRate
			httpResponseTimeAverage
			mobileSessionCount
			networkFailureRate
			usersAffectedCount
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on SecureCredentialEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		description
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		secureCredentialId
		secureCredentialSummary {
			failingMonitorCount
			monitorCount
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
		updatedAt
	}
	... on SyntheticMonitorEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		monitorId
		monitorSummary {
			locationsFailing
			locationsRunning
			status
			successRate
		}
		monitorType
		monitoredUrl
		period
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on ThirdPartyServiceEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on UnavailableEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
	}
	... on WorkloadEntity {
		__typename
		account {
			id
			name
			reportingEventTypes
		}
		createdAt
		createdByUser {
			email
			gravatar
			id
			name
		}
		recentAlertViolations {
			agentUrl
			alertSeverity
			closedAt
			label
			level
			openedAt
			violationId
			violationUrl
		}
		relatedEntities {
			nextCursor
		}
		relationships {
			type
		}
		tags {
			key
			values
		}
		tagsWithMetadata {
			key
		}
		updatedAt
		workloadStatus {
			description
			statusSource
			statusValue
			summary
		}
	}
} } }`

// Search for entities using a custom query.
//
// For more details on how to create a custom query
// and what entity data you can request, visit our
// [entity docs](https://docs.newrelic.com/docs/apis/graphql-api/tutorials/use-new-relic-graphql-api-query-entities).
//
// Note: you must supply either a `query` OR a `queryBuilder` argument, not both.
func (a *Entities) GetEntitySearch(
	options EntitySearchOptions,
	query string,
	queryBuilder EntitySearchQueryBuilder,
	sortBy []EntitySearchSortCriteria,
) (*EntitySearch, error) {
	return a.GetEntitySearchWithContext(context.Background(),
		options,
		query,
		queryBuilder,
		sortBy,
	)
}

// Search for entities using a custom query.
//
// For more details on how to create a custom query
// and what entity data you can request, visit our
// [entity docs](https://docs.newrelic.com/docs/apis/graphql-api/tutorials/use-new-relic-graphql-api-query-entities).
//
// Note: you must supply either a `query` OR a `queryBuilder` argument, not both.
func (a *Entities) GetEntitySearchWithContext(
	ctx context.Context,
	options EntitySearchOptions,
	query string,
	queryBuilder EntitySearchQueryBuilder,
	sortBy []EntitySearchSortCriteria,
) (*EntitySearch, error) {

	resp := entitySearchResponse{}
	vars := map[string]interface{}{
		"options":      options,
		"query":        query,
		"queryBuilder": queryBuilder,
		"sortBy":       sortBy,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, getEntitySearchQuery, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.Actor.EntitySearch, nil
}

const getEntitySearchQuery = `query(
	$queryBuilder: EntitySearchQueryBuilder,
) { actor { entitySearch(
	queryBuilder: $queryBuilder,
) {
	count
	query
	results {
		entities {
			__typename
			accountId
			alertSeverity
			domain
			entityType
			guid
			indexedAt
			name
			permalink
			reporting
			type
			... on ApmApplicationEntityOutline {
				__typename
				applicationId
				language
			}
			... on ApmDatabaseInstanceEntityOutline {
				__typename
				host
				portOrPath
				vendor
			}
			... on ApmExternalServiceEntityOutline {
				__typename
				host
			}
			... on BrowserApplicationEntityOutline {
				__typename
				agentInstallType
				applicationId
				servingApmApplicationId
			}
			... on DashboardEntityOutline {
				__typename
				createdAt
				dashboardParentGuid
				permissions
				updatedAt
			}
			... on ExternalEntityOutline {
				__typename
			}
			... on GenericEntityOutline {
				__typename
			}
			... on GenericInfrastructureEntityOutline {
				__typename
				integrationTypeCode
			}
			... on InfrastructureAwsLambdaFunctionEntityOutline {
				__typename
				integrationTypeCode
				runtime
			}
			... on InfrastructureHostEntityOutline {
				__typename
			}
			... on MobileApplicationEntityOutline {
				__typename
				applicationId
			}
			... on SecureCredentialEntityOutline {
				__typename
				description
				secureCredentialId
				updatedAt
			}
			... on SyntheticMonitorEntityOutline {
				__typename
				monitorId
				monitorType
				monitoredUrl
				period
			}
			... on ThirdPartyServiceEntityOutline {
				__typename
			}
			... on UnavailableEntityOutline {
				__typename
			}
			... on WorkloadEntityOutline {
				__typename
				createdAt
				updatedAt
			}
		}
		nextCursor
	}
	types {
		count
		domain
		entityType
		type
	}
} } }`
