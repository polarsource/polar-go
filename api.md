# Users

## Benefits

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserBenefitGetResponse">UserBenefitGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserBenefitListResponse">UserBenefitListResponse</a>

Methods:

- <code title="get /v1/users/benefits/{id}">client.Users.Benefits.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserBenefitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserBenefitGetResponse">UserBenefitGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/benefits/">client.Users.Benefits.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserBenefitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserBenefitListParams">UserBenefitListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserBenefitListResponse">UserBenefitListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orders

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserOrderGetResponse">UserOrderGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserOrderListResponse">UserOrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserOrderInvoiceResponse">UserOrderInvoiceResponse</a>

Methods:

- <code title="get /v1/users/orders/{id}">client.Users.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserOrderGetResponse">UserOrderGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/orders/">client.Users.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserOrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserOrderListParams">UserOrderListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserOrderListResponse">UserOrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/orders/{id}/invoice">client.Users.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserOrderService.Invoice">Invoice</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserOrderInvoiceResponse">UserOrderInvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Subscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionNewResponse">UserSubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionGetResponse">UserSubscriptionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionUpdateResponse">UserSubscriptionUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionListResponse">UserSubscriptionListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionDeleteResponse">UserSubscriptionDeleteResponse</a>

Methods:

- <code title="post /v1/users/subscriptions/">client.Users.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionNewParams">UserSubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionNewResponse">UserSubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/subscriptions/{id}">client.Users.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionGetResponse">UserSubscriptionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/users/subscriptions/{id}">client.Users.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionUpdateParams">UserSubscriptionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionUpdateResponse">UserSubscriptionUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/subscriptions/">client.Users.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionListParams">UserSubscriptionListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionListResponse">UserSubscriptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/users/subscriptions/{id}">client.Users.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserSubscriptionDeleteResponse">UserSubscriptionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Advertisements

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementNewResponse">UserAdvertisementNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementGetResponse">UserAdvertisementGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementUpdateResponse">UserAdvertisementUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementListResponse">UserAdvertisementListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementDeleteResponse">UserAdvertisementDeleteResponse</a>

Methods:

- <code title="post /v1/users/advertisements/">client.Users.Advertisements.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementNewParams">UserAdvertisementNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementNewResponse">UserAdvertisementNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/advertisements/{id}">client.Users.Advertisements.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementGetResponse">UserAdvertisementGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/users/advertisements/{id}">client.Users.Advertisements.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementUpdateParams">UserAdvertisementUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementUpdateResponse">UserAdvertisementUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/advertisements/">client.Users.Advertisements.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementListParams">UserAdvertisementListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementListResponse">UserAdvertisementListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/users/advertisements/{id}">client.Users.Advertisements.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementDeleteResponse">UserAdvertisementDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/users/advertisements/{id}/enable">client.Users.Advertisements.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementService.Enable">Enable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserAdvertisementEnableParams">UserAdvertisementEnableParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Downloadables

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserDownloadableGetResponse">UserDownloadableGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserDownloadableListResponse">UserDownloadableListResponse</a>

Methods:

- <code title="get /v1/users/downloadables/{token}">client.Users.Downloadables.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserDownloadableService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, token <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserDownloadableGetResponse">UserDownloadableGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/downloadables/">client.Users.Downloadables.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserDownloadableService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserDownloadableListParams">UserDownloadableListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserDownloadableListResponse">UserDownloadableListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Funding

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FundingLookupResponse">FundingLookupResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FundingSearchResponse">FundingSearchResponse</a>

Methods:

- <code title="get /v1/funding/lookup">client.Funding.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FundingService.Lookup">Lookup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FundingLookupParams">FundingLookupParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FundingLookupResponse">FundingLookupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/funding/search">client.Funding.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FundingService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FundingSearchParams">FundingSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FundingSearchResponse">FundingSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ExternalOrganizations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ExternalOrganizationListResponse">ExternalOrganizationListResponse</a>

Methods:

- <code title="get /v1/external_organizations/">client.ExternalOrganizations.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ExternalOrganizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ExternalOrganizationListParams">ExternalOrganizationListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ExternalOrganizationListResponse">ExternalOrganizationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Repositories

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RepositoryGetResponse">RepositoryGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RepositoryUpdateResponse">RepositoryUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RepositoryListResponse">RepositoryListResponse</a>

Methods:

- <code title="get /v1/repositories/{id}">client.Repositories.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RepositoryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RepositoryGetResponse">RepositoryGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/repositories/{id}">client.Repositories.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RepositoryService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RepositoryUpdateParams">RepositoryUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RepositoryUpdateResponse">RepositoryUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/repositories/">client.Repositories.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RepositoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RepositoryListParams">RepositoryListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RepositoryListResponse">RepositoryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Rewards

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RewardSearchResponse">RewardSearchResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RewardSummaryResponse">RewardSummaryResponse</a>

Methods:

- <code title="get /v1/rewards/search">client.Rewards.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RewardService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RewardSearchParams">RewardSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RewardSearchResponse">RewardSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/rewards/summary">client.Rewards.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RewardService.Summary">Summary</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RewardSummaryParams">RewardSummaryParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RewardSummaryResponse">RewardSummaryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PullRequests

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PullRequestSearchResponse">PullRequestSearchResponse</a>

Methods:

- <code title="get /v1/pull_requests/search">client.PullRequests.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PullRequestService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PullRequestSearchParams">PullRequestSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PullRequestSearchResponse">PullRequestSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountNewResponse">AccountNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountGetResponse">AccountGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountDashboardLinkResponse">AccountDashboardLinkResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountOnboardingLinkResponse">AccountOnboardingLinkResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountSearchResponse">AccountSearchResponse</a>

Methods:

- <code title="post /v1/accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountNewParams">AccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountNewResponse">AccountNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts/{id}">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountGetResponse">AccountGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/accounts/{id}/dashboard_link">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountService.DashboardLink">DashboardLink</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountDashboardLinkResponse">AccountDashboardLinkResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/accounts/{id}/onboarding_link">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountService.OnboardingLink">OnboardingLink</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountOnboardingLinkParams">AccountOnboardingLinkParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountOnboardingLinkResponse">AccountOnboardingLinkResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts/search">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountSearchParams">AccountSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AccountSearchResponse">AccountSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Issues

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueGetResponse">IssueGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueUpdateResponse">IssueUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueListResponse">IssueListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueConfirmSolvedResponse">IssueConfirmSolvedResponse</a>

Methods:

- <code title="get /v1/issues/{id}">client.Issues.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueGetResponse">IssueGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/issues/{id}">client.Issues.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueUpdateParams">IssueUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueUpdateResponse">IssueUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/issues/">client.Issues.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueListParams">IssueListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueListResponse">IssueListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/issues/{id}/confirm_solved">client.Issues.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueService.ConfirmSolved">ConfirmSolved</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueConfirmSolvedParams">IssueConfirmSolvedParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueConfirmSolvedResponse">IssueConfirmSolvedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Lookup

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueLookupGetResponse">IssueLookupGetResponse</a>

Methods:

- <code title="get /v1/issues/lookup">client.Issues.Lookup.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueLookupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueLookupGetParams">IssueLookupGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueLookupGetResponse">IssueLookupGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Body

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueBodyGetResponse">IssueBodyGetResponse</a>

Methods:

- <code title="get /v1/issues/{id}/body">client.Issues.Body.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueBodyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IssueBodyGetResponse">IssueBodyGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Pledges

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeGetResponse">PledgeGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeSearchResponse">PledgeSearchResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeSpendingResponse">PledgeSpendingResponse</a>

Methods:

- <code title="get /v1/pledges/{id}">client.Pledges.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeGetResponse">PledgeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/pledges/search">client.Pledges.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeSearchParams">PledgeSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeSearchResponse">PledgeSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/pledges/spending">client.Pledges.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeService.Spending">Spending</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeSpendingParams">PledgeSpendingParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeSpendingResponse">PledgeSpendingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Summary

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeSummaryGetResponse">PledgeSummaryGetResponse</a>

Methods:

- <code title="get /v1/pledges/summary">client.Pledges.Summary.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeSummaryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeSummaryGetParams">PledgeSummaryGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PledgeSummaryGetResponse">PledgeSummaryGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Organizations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationNewResponse">OrganizationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationGetResponse">OrganizationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationUpdateResponse">OrganizationUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationListResponse">OrganizationListResponse</a>

Methods:

- <code title="post /v1/organizations/">client.Organizations.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationNewParams">OrganizationNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationNewResponse">OrganizationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/organizations/{id}">client.Organizations.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationGetResponse">OrganizationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/organizations/{id}">client.Organizations.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationUpdateParams">OrganizationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationUpdateResponse">OrganizationUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/organizations/">client.Organizations.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationListParams">OrganizationListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationListResponse">OrganizationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Customers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationCustomerListResponse">OrganizationCustomerListResponse</a>

Methods:

- <code title="get /v1/organizations/{id}/customers">client.Organizations.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationCustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationCustomerListParams">OrganizationCustomerListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationCustomerListResponse">OrganizationCustomerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Subscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionNewResponse">SubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionListResponse">SubscriptionListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionExportResponse">SubscriptionExportResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionImportResponse">SubscriptionImportResponse</a>

Methods:

- <code title="post /v1/subscriptions/">client.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionNewParams">SubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionNewResponse">SubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/subscriptions/">client.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionListParams">SubscriptionListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionListResponse">SubscriptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/subscriptions/export">client.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionService.Export">Export</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionExportParams">SubscriptionExportParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionExportResponse">SubscriptionExportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/subscriptions/import">client.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionService.Import">Import</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionImportParams">SubscriptionImportParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#SubscriptionImportResponse">SubscriptionImportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Articles

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Article">Article</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleReceivers">ArticleReceivers</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceArticle">ListResourceArticle</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticlePreviewResponse">ArticlePreviewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleSendResponse">ArticleSendResponse</a>

Methods:

- <code title="post /v1/articles/">client.Articles.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleNewParams">ArticleNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Article">Article</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/articles/{id}">client.Articles.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Article">Article</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/articles/{id}">client.Articles.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleUpdateParams">ArticleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Article">Article</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/articles/">client.Articles.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleListParams">ArticleListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceArticle">ListResourceArticle</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/articles/{id}">client.Articles.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/articles/{id}/preview">client.Articles.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleService.Preview">Preview</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticlePreviewParams">ArticlePreviewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticlePreviewResponse">ArticlePreviewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/articles/{id}/send">client.Articles.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleSendResponse">ArticleSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Receivers

Methods:

- <code title="get /v1/articles/{id}/receivers">client.Articles.Receivers.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleReceiverService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ArticleReceivers">ArticleReceivers</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceTransaction">ListResourceTransaction</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionDetails">TransactionDetails</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionsSummary">TransactionsSummary</a>

Methods:

- <code title="get /v1/transactions/lookup">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionService.Lookup">Lookup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionLookupParams">TransactionLookupParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionDetails">TransactionDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/transactions/search">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionSearchParams">TransactionSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceTransaction">ListResourceTransaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/transactions/summary">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionService.Summary">Summary</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionSummaryParams">TransactionSummaryParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionsSummary">TransactionsSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Payouts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PayoutEstimate">PayoutEstimate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Transaction">Transaction</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionPayoutCsvResponse">TransactionPayoutCsvResponse</a>

Methods:

- <code title="post /v1/transactions/payouts">client.Transactions.Payouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionPayoutService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionPayoutNewParams">TransactionPayoutNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/transactions/payouts">client.Transactions.Payouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionPayoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionPayoutListParams">TransactionPayoutListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#PayoutEstimate">PayoutEstimate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/transactions/payouts/{id}/csv">client.Transactions.Payouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionPayoutService.Csv">Csv</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TransactionPayoutCsvResponse">TransactionPayoutCsvResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Advertisements

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AdvertisementCampaign">AdvertisementCampaign</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AdvertisementCampaignListResource">AdvertisementCampaignListResource</a>

Methods:

- <code title="get /v1/advertisements/{id}">client.Advertisements.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AdvertisementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AdvertisementCampaign">AdvertisementCampaign</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/advertisements/">client.Advertisements.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AdvertisementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AdvertisementListParams">AdvertisementListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AdvertisementCampaignListResource">AdvertisementCampaignListResource</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Donations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationStatistics">DonationStatistics</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceDonation">ListResourceDonation</a>

Methods:

- <code title="get /v1/donations/search">client.Donations.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationSearchParams">DonationSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceDonation">ListResourceDonation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/donations/statistics">client.Donations.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationService.Statistics">Statistics</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationStatisticsParams">DonationStatisticsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationStatistics">DonationStatistics</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PaymentIntent

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationStripePaymentIntentMutationResponse">DonationStripePaymentIntentMutationResponse</a>

Methods:

- <code title="post /v1/donations/payment_intent">client.Donations.PaymentIntent.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationPaymentIntentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationPaymentIntentNewParams">DonationPaymentIntentNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationStripePaymentIntentMutationResponse">DonationStripePaymentIntentMutationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/donations/payment_intent/{id}">client.Donations.PaymentIntent.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationPaymentIntentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationPaymentIntentUpdateParams">DonationPaymentIntentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationStripePaymentIntentMutationResponse">DonationStripePaymentIntentMutationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Public

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourcePublicDonation">ListResourcePublicDonation</a>

Methods:

- <code title="get /v1/donations/public/search">client.Donations.Public.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationPublicService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DonationPublicSearchParams">DonationPublicSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourcePublicDonation">ListResourcePublicDonation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Oauth2

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AuthorizeResponseOrganization">AuthorizeResponseOrganization</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#AuthorizeResponseUser">AuthorizeResponseUser</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IntrospectTokenResponse">IntrospectTokenResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RevokeTokenResponse">RevokeTokenResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TokenResponse">TokenResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2ListResponse">Oauth2ListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2AuthorizeResponse">Oauth2AuthorizeResponse</a>

Methods:

- <code title="get /v1/oauth2/">client.Oauth2.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2ListParams">Oauth2ListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2ListResponse">Oauth2ListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/oauth2/authorize">client.Oauth2.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2Service.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2AuthorizeResponse">Oauth2AuthorizeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/oauth2/introspect">client.Oauth2.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2Service.Introspect">Introspect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2IntrospectParams">Oauth2IntrospectParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#IntrospectTokenResponse">IntrospectTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/oauth2/revoke">client.Oauth2.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2Service.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RevokeParams">Oauth2RevokeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#RevokeTokenResponse">RevokeTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/oauth2/token">client.Oauth2.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2Service.Token">Token</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2TokenParams">Oauth2TokenParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#TokenResponse">TokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Register

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterNewResponse">Oauth2RegisterNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterGetResponse">Oauth2RegisterGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterUpdateResponse">Oauth2RegisterUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterDeleteResponse">Oauth2RegisterDeleteResponse</a>

Methods:

- <code title="post /v1/oauth2/register">client.Oauth2.Register.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterNewParams">Oauth2RegisterNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterNewResponse">Oauth2RegisterNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/oauth2/register/{client_id}">client.Oauth2.Register.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterGetResponse">Oauth2RegisterGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/oauth2/register/{client_id}">client.Oauth2.Register.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterUpdateParams">Oauth2RegisterUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterUpdateResponse">Oauth2RegisterUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/oauth2/register/{client_id}">client.Oauth2.Register.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2RegisterDeleteResponse">Oauth2RegisterDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Userinfo

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserInfoOrganization">UserInfoOrganization</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#UserInfoUser">UserInfoUser</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2UserinfoNewResponse">Oauth2UserinfoNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2UserinfoGetResponse">Oauth2UserinfoGetResponse</a>

Methods:

- <code title="post /v1/oauth2/userinfo">client.Oauth2.Userinfo.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2UserinfoService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2UserinfoNewResponse">Oauth2UserinfoNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/oauth2/userinfo">client.Oauth2.Userinfo.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2UserinfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Oauth2UserinfoGetResponse">Oauth2UserinfoGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Benefits

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadables">ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadables</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitNewResponse">BenefitNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitGetResponse">BenefitGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitUpdateResponse">BenefitUpdateResponse</a>

Methods:

- <code title="post /v1/benefits/">client.Benefits.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitNewParams">BenefitNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitNewResponse">BenefitNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/benefits/{id}">client.Benefits.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitGetResponse">BenefitGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/benefits/{id}">client.Benefits.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitUpdateParams">BenefitUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitUpdateResponse">BenefitUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/benefits/">client.Benefits.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitListParams">BenefitListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadables">ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadables</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/benefits/{id}">client.Benefits.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Grants

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceBenefitGrant">ListResourceBenefitGrant</a>

Methods:

- <code title="get /v1/benefits/{id}/grants">client.Benefits.Grants.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitGrantService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#BenefitGrantListParams">BenefitGrantListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceBenefitGrant">ListResourceBenefitGrant</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

## Endpoints

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceWebhookEndpoint">ListResourceWebhookEndpoint</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEndpoint">WebhookEndpoint</a>

Methods:

- <code title="post /v1/webhooks/endpoints">client.Webhooks.Endpoints.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEndpointService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEndpointNewParams">WebhookEndpointNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEndpoint">WebhookEndpoint</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/webhooks/endpoints/{id}">client.Webhooks.Endpoints.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEndpointService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEndpoint">WebhookEndpoint</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/webhooks/endpoints/{id}">client.Webhooks.Endpoints.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEndpointService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEndpointUpdateParams">WebhookEndpointUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEndpoint">WebhookEndpoint</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/webhooks/endpoints">client.Webhooks.Endpoints.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEndpointService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEndpointListParams">WebhookEndpointListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceWebhookEndpoint">ListResourceWebhookEndpoint</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/webhooks/endpoints/{id}">client.Webhooks.Endpoints.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEndpointService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Deliveries

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceWebhookDelivery">ListResourceWebhookDelivery</a>

Methods:

- <code title="get /v1/webhooks/deliveries">client.Webhooks.Deliveries.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookDeliveryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookDeliveryListParams">WebhookDeliveryListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceWebhookDelivery">ListResourceWebhookDelivery</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Events

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEventRedeliverResponse">WebhookEventRedeliverResponse</a>

Methods:

- <code title="post /v1/webhooks/events/{id}/redeliver">client.Webhooks.Events.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEventService.Redeliver">Redeliver</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#WebhookEventRedeliverResponse">WebhookEventRedeliverResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Products

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceProduct">ListResourceProduct</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductOutput">ProductOutput</a>

Methods:

- <code title="post /v1/products/">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductNewParams">ProductNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductOutput">ProductOutput</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductOutput">ProductOutput</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductUpdateParams">ProductUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductOutput">ProductOutput</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/products/">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductListParams">ProductListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceProduct">ListResourceProduct</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Benefits

Methods:

- <code title="post /v1/products/{id}/benefits">client.Products.Benefits.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductBenefitService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductBenefitUpdateParams">ProductBenefitUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductOutput">ProductOutput</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Orders

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceOrder">ListResourceOrder</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrderOutput">OrderOutput</a>

Methods:

- <code title="get /v1/orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrderOutput">OrderOutput</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/orders/">client.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrderListParams">OrderListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceOrder">ListResourceOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Invoice

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrderInvoice">OrderInvoice</a>

Methods:

- <code title="get /v1/orders/{id}/invoice">client.Orders.Invoice.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrderInvoiceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrderInvoice">OrderInvoice</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Checkouts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Checkout">Checkout</a>

Methods:

- <code title="post /v1/checkouts/">client.Checkouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#CheckoutService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#CheckoutNewParams">CheckoutNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Checkout">Checkout</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/checkouts/{id}">client.Checkouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#CheckoutService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#Checkout">Checkout</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#DownloadableFileRead">DownloadableFileRead</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileUpload">FileUpload</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceAnnotatedUnion">ListResourceAnnotatedUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#OrganizationAvatarFileRead">OrganizationAvatarFileRead</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ProductMediaFileReadOutput">ProductMediaFileReadOutput</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileUpdateResponse">FileUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileUploadedResponse">FileUploadedResponse</a>

Methods:

- <code title="post /v1/files/">client.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileNewParams">FileNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileUpload">FileUpload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/files/{id}">client.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileUpdateParams">FileUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileUpdateResponse">FileUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/files/">client.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileListParams">FileListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#ListResourceAnnotatedUnion">ListResourceAnnotatedUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/files/{id}">client.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/files/{id}/uploaded">client.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileService.Uploaded">Uploaded</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileUploadedParams">FileUploadedParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#FileUploadedResponse">FileUploadedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Metrics

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#MetricsResponse">MetricsResponse</a>

Methods:

- <code title="get /v1/metrics/">client.Metrics.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#MetricService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#MetricListParams">MetricListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#MetricsResponse">MetricsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Limits

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#MetricsLimits">MetricsLimits</a>

Methods:

- <code title="get /v1/metrics/limits">client.Metrics.Limits.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#MetricLimitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/polar-go#MetricsLimits">MetricsLimits</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
