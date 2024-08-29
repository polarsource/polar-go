# Users

## Benefits

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserBenefitGetResponse">UserBenefitGetResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserBenefitListResponse">UserBenefitListResponse</a>

Methods:

- <code title="get /v1/users/benefits/{id}">client.Users.Benefits.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserBenefitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserBenefitGetResponse">UserBenefitGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/benefits/">client.Users.Benefits.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserBenefitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserBenefitListParams">UserBenefitListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserBenefitListResponse">UserBenefitListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orders

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserOrderGetResponse">UserOrderGetResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserOrderListResponse">UserOrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserOrderInvoiceResponse">UserOrderInvoiceResponse</a>

Methods:

- <code title="get /v1/users/orders/{id}">client.Users.Orders.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserOrderGetResponse">UserOrderGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/orders/">client.Users.Orders.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserOrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserOrderListParams">UserOrderListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserOrderListResponse">UserOrderListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/orders/{id}/invoice">client.Users.Orders.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserOrderService.Invoice">Invoice</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserOrderInvoiceResponse">UserOrderInvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Subscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionNewResponse">UserSubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionGetResponse">UserSubscriptionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionUpdateResponse">UserSubscriptionUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionListResponse">UserSubscriptionListResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionDeleteResponse">UserSubscriptionDeleteResponse</a>

Methods:

- <code title="post /v1/users/subscriptions/">client.Users.Subscriptions.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionNewParams">UserSubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionNewResponse">UserSubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/subscriptions/{id}">client.Users.Subscriptions.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionGetResponse">UserSubscriptionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/users/subscriptions/{id}">client.Users.Subscriptions.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionUpdateParams">UserSubscriptionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionUpdateResponse">UserSubscriptionUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/subscriptions/">client.Users.Subscriptions.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionListParams">UserSubscriptionListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionListResponse">UserSubscriptionListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/users/subscriptions/{id}">client.Users.Subscriptions.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserSubscriptionDeleteResponse">UserSubscriptionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Advertisements

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementNewResponse">UserAdvertisementNewResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementGetResponse">UserAdvertisementGetResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementUpdateResponse">UserAdvertisementUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementListResponse">UserAdvertisementListResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementDeleteResponse">UserAdvertisementDeleteResponse</a>

Methods:

- <code title="post /v1/users/advertisements/">client.Users.Advertisements.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementNewParams">UserAdvertisementNewParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementNewResponse">UserAdvertisementNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/advertisements/{id}">client.Users.Advertisements.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementGetResponse">UserAdvertisementGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/users/advertisements/{id}">client.Users.Advertisements.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementUpdateParams">UserAdvertisementUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementUpdateResponse">UserAdvertisementUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/advertisements/">client.Users.Advertisements.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementListParams">UserAdvertisementListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementListResponse">UserAdvertisementListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/users/advertisements/{id}">client.Users.Advertisements.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementDeleteResponse">UserAdvertisementDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/users/advertisements/{id}/enable">client.Users.Advertisements.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementService.Enable">Enable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserAdvertisementEnableParams">UserAdvertisementEnableParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Downloadables

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserDownloadableGetResponse">UserDownloadableGetResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserDownloadableListResponse">UserDownloadableListResponse</a>

Methods:

- <code title="get /v1/users/downloadables/{token}">client.Users.Downloadables.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserDownloadableService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, token <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserDownloadableGetResponse">UserDownloadableGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/users/downloadables/">client.Users.Downloadables.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserDownloadableService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserDownloadableListParams">UserDownloadableListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#UserDownloadableListResponse">UserDownloadableListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ExternalOrganizations

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ExternalOrganizationListResponse">ExternalOrganizationListResponse</a>

Methods:

- <code title="get /v1/external_organizations/">client.ExternalOrganizations.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ExternalOrganizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ExternalOrganizationListParams">ExternalOrganizationListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ExternalOrganizationListResponse">ExternalOrganizationListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Repositories

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#RepositoryGetResponse">RepositoryGetResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#RepositoryUpdateResponse">RepositoryUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#RepositoryListResponse">RepositoryListResponse</a>

Methods:

- <code title="get /v1/repositories/{id}">client.Repositories.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#RepositoryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#RepositoryGetResponse">RepositoryGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/repositories/{id}">client.Repositories.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#RepositoryService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#RepositoryUpdateParams">RepositoryUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#RepositoryUpdateResponse">RepositoryUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/repositories/">client.Repositories.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#RepositoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#RepositoryListParams">RepositoryListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#RepositoryListResponse">RepositoryListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Organizations

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationNewResponse">OrganizationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationGetResponse">OrganizationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationUpdateResponse">OrganizationUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationListResponse">OrganizationListResponse</a>

Methods:

- <code title="post /v1/organizations/">client.Organizations.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationNewParams">OrganizationNewParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationNewResponse">OrganizationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/organizations/{id}">client.Organizations.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationGetResponse">OrganizationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/organizations/{id}">client.Organizations.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationUpdateParams">OrganizationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationUpdateResponse">OrganizationUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/organizations/">client.Organizations.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationListParams">OrganizationListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationListResponse">OrganizationListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Subscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionNewResponse">SubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionListResponse">SubscriptionListResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionExportResponse">SubscriptionExportResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionImportResponse">SubscriptionImportResponse</a>

Methods:

- <code title="post /v1/subscriptions/">client.Subscriptions.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionNewParams">SubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionNewResponse">SubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/subscriptions/">client.Subscriptions.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionListParams">SubscriptionListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionListResponse">SubscriptionListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/subscriptions/export">client.Subscriptions.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionService.Export">Export</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionExportParams">SubscriptionExportParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionExportResponse">SubscriptionExportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/subscriptions/import">client.Subscriptions.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionService.Import">Import</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionImportParams">SubscriptionImportParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#SubscriptionImportResponse">SubscriptionImportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Articles

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Article">Article</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleReceivers">ArticleReceivers</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ListResourceArticle">ListResourceArticle</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticlePreviewResponse">ArticlePreviewResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleSendResponse">ArticleSendResponse</a>

Methods:

- <code title="post /v1/articles/">client.Articles.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleNewParams">ArticleNewParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Article">Article</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/articles/{id}">client.Articles.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Article">Article</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/articles/{id}">client.Articles.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleUpdateParams">ArticleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Article">Article</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/articles/">client.Articles.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleListParams">ArticleListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Article">Article</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/articles/{id}">client.Articles.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/articles/{id}/preview">client.Articles.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleService.Preview">Preview</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticlePreviewParams">ArticlePreviewParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticlePreviewResponse">ArticlePreviewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/articles/{id}/send">client.Articles.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleSendResponse">ArticleSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Receivers

Methods:

- <code title="get /v1/articles/{id}/receivers">client.Articles.Receivers.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleReceiverService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ArticleReceivers">ArticleReceivers</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Advertisements

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#AdvertisementCampaign">AdvertisementCampaign</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#AdvertisementCampaignListResource">AdvertisementCampaignListResource</a>

Methods:

- <code title="get /v1/advertisements/{id}">client.Advertisements.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#AdvertisementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#AdvertisementCampaign">AdvertisementCampaign</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/advertisements/">client.Advertisements.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#AdvertisementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#AdvertisementListParams">AdvertisementListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#AdvertisementCampaign">AdvertisementCampaign</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Benefits

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadables">ListResourceUnionBenefitArticlesBenefitAdsBenefitCustomBenefitDiscordBenefitGitHubRepositoryBenefitDownloadables</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitNewResponse">BenefitNewResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitGetResponse">BenefitGetResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitUpdateResponse">BenefitUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitListResponse">BenefitListResponse</a>

Methods:

- <code title="post /v1/benefits/">client.Benefits.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitNewParams">BenefitNewParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitNewResponse">BenefitNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/benefits/{id}">client.Benefits.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitGetResponse">BenefitGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/benefits/{id}">client.Benefits.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitUpdateParams">BenefitUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitUpdateResponse">BenefitUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/benefits/">client.Benefits.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitListParams">BenefitListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitListResponse">BenefitListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/benefits/{id}">client.Benefits.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Grants

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ListResourceBenefitGrant">ListResourceBenefitGrant</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitGrantListResponse">BenefitGrantListResponse</a>

Methods:

- <code title="get /v1/benefits/{id}/grants">client.Benefits.Grants.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitGrantService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitGrantListParams">BenefitGrantListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#BenefitGrantListResponse">BenefitGrantListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Products

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ListResourceProduct">ListResourceProduct</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Product">Product</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ProductPrice">ProductPrice</a>

Methods:

- <code title="post /v1/products/">client.Products.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ProductService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ProductNewParams">ProductNewParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ProductService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ProductUpdateParams">ProductUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/products/">client.Products.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ProductListParams">ProductListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Product">Product</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Benefits

Methods:

- <code title="post /v1/products/{id}/benefits">client.Products.Benefits.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ProductBenefitService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ProductBenefitUpdateParams">ProductBenefitUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Orders

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ListResourceOrder">ListResourceOrder</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrderOutput">OrderOutput</a>

Methods:

- <code title="get /v1/orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrderOutput">OrderOutput</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/orders/">client.Orders.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrderListParams">OrderListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrderOutput">OrderOutput</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Invoice

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrderInvoice">OrderInvoice</a>

Methods:

- <code title="get /v1/orders/{id}/invoice">client.Orders.Invoice.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrderInvoiceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrderInvoice">OrderInvoice</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Checkouts

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Checkout">Checkout</a>

Methods:

- <code title="post /v1/checkouts/">client.Checkouts.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#CheckoutService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#CheckoutNewParams">CheckoutNewParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Checkout">Checkout</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/checkouts/{id}">client.Checkouts.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#CheckoutService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#Checkout">Checkout</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#DownloadableFileRead">DownloadableFileRead</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileUpload">FileUpload</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ListResourceAnnotatedUnion">ListResourceAnnotatedUnion</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#OrganizationAvatarFileRead">OrganizationAvatarFileRead</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#ProductMediaFileReadOutput">ProductMediaFileReadOutput</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileUpdateResponse">FileUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileListResponse">FileListResponse</a>
- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileUploadedResponse">FileUploadedResponse</a>

Methods:

- <code title="post /v1/files/">client.Files.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileNewParams">FileNewParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileUpload">FileUpload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/files/{id}">client.Files.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileUpdateParams">FileUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileUpdateResponse">FileUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/files/">client.Files.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileListParams">FileListParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go/internal/pagination#PolarPagination">PolarPagination</a>[<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileListResponse">FileListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/files/{id}">client.Files.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/files/{id}/uploaded">client.Files.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileService.Uploaded">Uploaded</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileUploadedParams">FileUploadedParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#FileUploadedResponse">FileUploadedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Metrics

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#MetricsResponse">MetricsResponse</a>

Methods:

- <code title="get /v1/metrics/">client.Metrics.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#MetricService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#MetricGetParams">MetricGetParams</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#MetricsResponse">MetricsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Limits

Response Types:

- <a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#MetricsLimits">MetricsLimits</a>

Methods:

- <code title="get /v1/metrics/limits">client.Metrics.Limits.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#MetricLimitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/polarsource/polar-go">polar</a>.<a href="https://pkg.go.dev/github.com/polarsource/polar-go#MetricsLimits">MetricsLimits</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
