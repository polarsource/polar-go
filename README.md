# github.com/polarsource/polar-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/polarsource/polar-go* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/polarsource/polar-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/polar/polar). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Polar API: Polar HTTP and Webhooks API

Read the docs at https://docs.polar.sh/api
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/polarsource/polar-go](#githubcompolarsourcepolar-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/polarsource/polar-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/components"
	"github.com/polarsource/polar-go/types"
	"log"
)

func main() {
	ctx := context.Background()

	s := polargo.New()

	res, err := s.EndpointcheckoutCreatedPost(ctx, components.WebhookCheckoutCreatedPayload{
		Data: components.Checkout{
			CreatedAt:              types.MustTimeFromString("2024-11-12T14:26:42.882Z"),
			ModifiedAt:             types.MustNewTimeFromString("2023-05-28T05:08:06.235Z"),
			ID:                     "<value>",
			Status:                 components.CheckoutStatusFailed,
			ClientSecret:           "<value>",
			URL:                    "https://heavy-beret.com/",
			ExpiresAt:              types.MustTimeFromString("2022-02-25T02:26:48.460Z"),
			SuccessURL:             "https://sardonic-final.info/",
			EmbedOrigin:            polargo.String("<value>"),
			Amount:                 polargo.Int64(962818),
			TaxAmount:              polargo.Int64(6400),
			Currency:               polargo.String("Yen"),
			SubtotalAmount:         polargo.Int64(648726),
			TotalAmount:            polargo.Int64(210702),
			ProductID:              "<value>",
			ProductPriceID:         "<value>",
			DiscountID:             polargo.String("<value>"),
			AllowDiscountCodes:     true,
			IsDiscountApplicable:   false,
			IsFreeProductPrice:     false,
			IsPaymentRequired:      false,
			IsPaymentSetupRequired: false,
			IsPaymentFormRequired:  false,
			CustomerID:             polargo.String("<value>"),
			CustomerName:           polargo.String("<value>"),
			CustomerEmail:          polargo.String("Ryley_Erdman@hotmail.com"),
			CustomerIPAddress:      polargo.String("<value>"),
			CustomerBillingAddress: &components.Address{
				Country: "South Africa",
			},
			CustomerTaxID:            polargo.String("<id>"),
			PaymentProcessorMetadata: components.PaymentProcessorMetadata{},
			Metadata: map[string]components.CheckoutMetadata{
				"key": components.CreateCheckoutMetadataInteger(
					18677,
				),
				"key1": components.CreateCheckoutMetadataInteger(
					95370,
				),
			},
			Product: components.CheckoutProduct{
				CreatedAt:      types.MustTimeFromString("2022-04-02T00:05:42.586Z"),
				ModifiedAt:     types.MustNewTimeFromString("2023-12-16T03:02:38.803Z"),
				ID:             "<value>",
				Name:           "<value>",
				Description:    polargo.String("for embarrassment untidy long-term near honestly separate yet"),
				IsRecurring:    true,
				IsArchived:     false,
				OrganizationID: "<value>",
				Prices: []components.ProductPrice{
					components.CreateProductPriceProductPriceOneTime(
						components.CreateProductPriceOneTimeProductPriceOneTimeCustom(
							components.ProductPriceOneTimeCustom{
								CreatedAt:     types.MustTimeFromString("2023-02-07T04:30:48.802Z"),
								ModifiedAt:    types.MustNewTimeFromString("2024-06-25T22:47:14.264Z"),
								ID:            "<value>",
								IsArchived:    false,
								ProductID:     "<value>",
								PriceCurrency: "<value>",
								MinimumAmount: polargo.Int64(691423),
								MaximumAmount: polargo.Int64(499526),
								PresetAmount:  polargo.Int64(18677),
							},
						),
					),
				},
				Benefits: []components.BenefitBase{
					components.BenefitBase{
						CreatedAt:      types.MustTimeFromString("2023-08-22T00:47:02.059Z"),
						ModifiedAt:     types.MustNewTimeFromString("2023-06-04T10:32:44.101Z"),
						ID:             "<value>",
						Type:           components.BenefitTypeLicenseKeys,
						Description:    "within jacket unless",
						Selectable:     false,
						Deletable:      false,
						OrganizationID: "<value>",
					},
				},
				Medias: []components.ProductMediaFileRead{
					components.ProductMediaFileRead{
						ID:                   "<value>",
						OrganizationID:       "<value>",
						Name:                 "<value>",
						Path:                 "/private/var",
						MimeType:             "<value>",
						Size:                 245189,
						StorageVersion:       polargo.String("<value>"),
						ChecksumEtag:         polargo.String("<value>"),
						ChecksumSha256Base64: polargo.String("<value>"),
						ChecksumSha256Hex:    polargo.String("<value>"),
						LastModifiedAt:       types.MustNewTimeFromString("2022-11-03T15:00:03.276Z"),
						Version:              polargo.String("<value>"),
						IsUploaded:           false,
						CreatedAt:            types.MustTimeFromString("2024-06-07T13:47:02.365Z"),
						SizeReadable:         "<value>",
						PublicURL:            "https://webbed-experience.name/",
					},
				},
			},
			ProductPrice: components.CreateProductPriceProductPriceOneTime(
				components.CreateProductPriceOneTimeProductPriceOneTimeFixed(
					components.ProductPriceOneTimeFixed{
						CreatedAt:     types.MustTimeFromString("2022-04-02T00:05:42.586Z"),
						ModifiedAt:    types.MustNewTimeFromString("2023-12-16T03:02:38.803Z"),
						ID:            "<value>",
						IsArchived:    false,
						ProductID:     "<value>",
						PriceCurrency: "<value>",
						PriceAmount:   740296,
					},
				),
			),
			Discount: components.CreateCheckoutDiscountCheckoutDiscountPercentageOnceForeverDuration(
				components.CheckoutDiscountPercentageOnceForeverDuration{
					Duration:    components.DiscountDurationRepeating,
					Type:        components.DiscountTypeFixed,
					BasisPoints: 341163,
					ID:          "<value>",
					Name:        "<value>",
					Code:        polargo.String("<value>"),
				},
			),
			SubscriptionID: polargo.String("<value>"),
			AttachedCustomFields: []components.AttachedCustomField{
				components.AttachedCustomField{
					CustomFieldID: "<value>",
					CustomField: components.CreateCustomFieldCustomFieldNumber(
						components.CustomFieldNumber{
							CreatedAt:  types.MustTimeFromString("2024-06-23T16:57:50.081Z"),
							ModifiedAt: types.MustNewTimeFromString("2023-12-14T18:25:33.693Z"),
							ID:         "<value>",
							Metadata: map[string]components.CustomFieldNumberMetadata{
								"key": components.CreateCustomFieldNumberMetadataStr(
									"<value>",
								),
							},
							Slug:           "<value>",
							Name:           "<value>",
							OrganizationID: "<value>",
							Properties:     components.CustomFieldNumberProperties{},
						},
					),
					Order:    996863,
					Required: false,
				},
				components.AttachedCustomField{
					CustomFieldID: "<value>",
					CustomField: components.CreateCustomFieldCustomFieldSelect(
						components.CustomFieldSelect{
							CreatedAt:  types.MustTimeFromString("2022-04-26T22:34:57.487Z"),
							ModifiedAt: types.MustNewTimeFromString("2022-08-07T19:57:51.694Z"),
							ID:         "<value>",
							Metadata: map[string]components.CustomFieldSelectMetadata{
								"key": components.CreateCustomFieldSelectMetadataInteger(
									856200,
								),
							},
							Slug:           "<value>",
							Name:           "<value>",
							OrganizationID: "<value>",
							Properties: components.CustomFieldSelectProperties{
								Options: []components.CustomFieldSelectOption{
									components.CustomFieldSelectOption{
										Value: "<value>",
										Label: "<value>",
									},
								},
							},
						},
					),
					Order:    72589,
					Required: true,
				},
				components.AttachedCustomField{
					CustomFieldID: "<value>",
					CustomField: components.CreateCustomFieldCustomFieldCheckbox(
						components.CustomFieldCheckbox{
							CreatedAt:  types.MustTimeFromString("2024-05-25T15:20:50.694Z"),
							ModifiedAt: types.MustNewTimeFromString("2023-11-28T14:29:40.329Z"),
							ID:         "<value>",
							Metadata: map[string]components.CustomFieldCheckboxMetadata{
								"key": components.CreateCustomFieldCheckboxMetadataBoolean(
									false,
								),
							},
							Slug:           "<value>",
							Name:           "<value>",
							OrganizationID: "<value>",
							Properties:     components.CustomFieldCheckboxProperties{},
						},
					),
					Order:    161325,
					Required: true,
				},
			},
			CustomerMetadata: map[string]components.CustomerMetadata{
				"key": components.CreateCustomerMetadataStr(
					"<value>",
				),
				"key1": components.CreateCustomerMetadataStr(
					"<value>",
				),
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Any != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name          | Type | Scheme      | Environment Variable |
| ------------- | ---- | ----------- | -------------------- |
| `AccessToken` | http | HTTP Bearer | `POLAR_ACCESS_TOKEN` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.ExternalOrganizations.List(ctx, operations.ExternalOrganizationsListRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceExternalOrganization != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Advertisements](docs/sdks/advertisements/README.md)

* [List](docs/sdks/advertisements/README.md#list) - List Campaigns
* [Get](docs/sdks/advertisements/README.md#get) - Get Campaign

### [Benefits](docs/sdks/benefits/README.md)

* [List](docs/sdks/benefits/README.md#list) - List Benefits
* [Create](docs/sdks/benefits/README.md#create) - Create Benefit
* [Get](docs/sdks/benefits/README.md#get) - Get Benefit
* [Update](docs/sdks/benefits/README.md#update) - Update Benefit
* [Delete](docs/sdks/benefits/README.md#delete) - Delete Benefit
* [Grants](docs/sdks/benefits/README.md#grants) - List Benefit Grants

### [CheckoutLinks](docs/sdks/checkoutlinks/README.md)

* [List](docs/sdks/checkoutlinks/README.md#list) - List Checkout Links
* [Create](docs/sdks/checkoutlinks/README.md#create) - Create Checkout Link
* [Get](docs/sdks/checkoutlinks/README.md#get) - Get Checkout Link
* [Update](docs/sdks/checkoutlinks/README.md#update) - Update Checkout Link
* [Delete](docs/sdks/checkoutlinks/README.md#delete) - Delete Checkout Link

### [Checkouts](docs/sdks/checkouts/README.md)

* [~~Create~~](docs/sdks/checkouts/README.md#create) - Create Checkout :warning: **Deprecated** Use [Create](docs/sdks/custom/README.md#create) instead.
* [~~Get~~](docs/sdks/checkouts/README.md#get) - Get Checkout :warning: **Deprecated**

#### [Checkouts.Custom](docs/sdks/custom/README.md)

* [List](docs/sdks/custom/README.md#list) - List Checkout Sessions
* [Create](docs/sdks/custom/README.md#create) - Create Checkout Session
* [Get](docs/sdks/custom/README.md#get) - Get Checkout Session
* [Update](docs/sdks/custom/README.md#update) - Update Checkout Session
* [ClientGet](docs/sdks/custom/README.md#clientget) - Get Checkout Session from Client
* [ClientUpdate](docs/sdks/custom/README.md#clientupdate) - Update Checkout Session from Client
* [ClientConfirm](docs/sdks/custom/README.md#clientconfirm) - Confirm Checkout Session from Client

### [CustomerPortal](docs/sdks/customerportal/README.md)


#### [CustomerPortal.BenefitGrants](docs/sdks/benefitgrants/README.md)

* [List](docs/sdks/benefitgrants/README.md#list) - List Benefit Grants
* [Get](docs/sdks/benefitgrants/README.md#get) - Get Benefit Grant
* [Update](docs/sdks/benefitgrants/README.md#update) - Update Benefit Grant

#### [CustomerPortal.Customers](docs/sdks/polarcustomers/README.md)

* [Get](docs/sdks/polarcustomers/README.md#get) - Get Customer

#### [CustomerPortal.Downloadables](docs/sdks/downloadables/README.md)

* [List](docs/sdks/downloadables/README.md#list) - List Downloadables
* [Get](docs/sdks/downloadables/README.md#get) - Get Downloadable

#### [CustomerPortal.LicenseKeys](docs/sdks/polarlicensekeys/README.md)

* [List](docs/sdks/polarlicensekeys/README.md#list) - List License Keys
* [Get](docs/sdks/polarlicensekeys/README.md#get) - Get License Key
* [Validate](docs/sdks/polarlicensekeys/README.md#validate) - Validate License Key
* [Activate](docs/sdks/polarlicensekeys/README.md#activate) - Activate License Key
* [Deactivate](docs/sdks/polarlicensekeys/README.md#deactivate) - Deactivate License Key

#### [CustomerPortal.Orders](docs/sdks/polarorders/README.md)

* [List](docs/sdks/polarorders/README.md#list) - List Orders
* [Get](docs/sdks/polarorders/README.md#get) - Get Order
* [Invoice](docs/sdks/polarorders/README.md#invoice) - Get Order Invoice

#### [CustomerPortal.Organizations](docs/sdks/polarorganizations/README.md)

* [Get](docs/sdks/polarorganizations/README.md#get) - Get Organization

#### [CustomerPortal.Subscriptions](docs/sdks/polarsubscriptions/README.md)

* [List](docs/sdks/polarsubscriptions/README.md#list) - List Subscriptions
* [Get](docs/sdks/polarsubscriptions/README.md#get) - Get Subscription
* [Update](docs/sdks/polarsubscriptions/README.md#update) - Update Subscription
* [Cancel](docs/sdks/polarsubscriptions/README.md#cancel) - Cancel Subscription

### [Customers](docs/sdks/customers/README.md)

* [List](docs/sdks/customers/README.md#list) - List Customers
* [Create](docs/sdks/customers/README.md#create) - Create Customer
* [Get](docs/sdks/customers/README.md#get) - Get Customer
* [Update](docs/sdks/customers/README.md#update) - Update Customer
* [Delete](docs/sdks/customers/README.md#delete) - Delete Customer

### [CustomerSessions](docs/sdks/customersessions/README.md)

* [Create](docs/sdks/customersessions/README.md#create) - Create Customer Session

### [CustomFields](docs/sdks/customfields/README.md)

* [List](docs/sdks/customfields/README.md#list) - List Custom Fields
* [Create](docs/sdks/customfields/README.md#create) - Create Custom Field
* [Get](docs/sdks/customfields/README.md#get) - Get Custom Field
* [Update](docs/sdks/customfields/README.md#update) - Update Custom Field
* [Delete](docs/sdks/customfields/README.md#delete) - Delete Custom Field

### [Discounts](docs/sdks/discounts/README.md)

* [List](docs/sdks/discounts/README.md#list) - List Discounts
* [Create](docs/sdks/discounts/README.md#create) - Create Discount
* [Get](docs/sdks/discounts/README.md#get) - Get Discount
* [Update](docs/sdks/discounts/README.md#update) - Update Discount
* [Delete](docs/sdks/discounts/README.md#delete) - Delete Discount

### [ExternalOrganizations](docs/sdks/externalorganizations/README.md)

* [List](docs/sdks/externalorganizations/README.md#list) - List External Organizations

### [Files](docs/sdks/files/README.md)

* [List](docs/sdks/files/README.md#list) - List Files
* [Create](docs/sdks/files/README.md#create) - Create File
* [Uploaded](docs/sdks/files/README.md#uploaded) - Complete File Upload
* [Update](docs/sdks/files/README.md#update) - Update File
* [Delete](docs/sdks/files/README.md#delete) - Delete File

### [LicenseKeys](docs/sdks/licensekeys/README.md)

* [List](docs/sdks/licensekeys/README.md#list) - List License Keys
* [Get](docs/sdks/licensekeys/README.md#get) - Get License Key
* [Update](docs/sdks/licensekeys/README.md#update) - Update License Key
* [GetActivation](docs/sdks/licensekeys/README.md#getactivation) - Get Activation

### [Metrics](docs/sdks/metrics/README.md)

* [Get](docs/sdks/metrics/README.md#get) - Get Metrics
* [Limits](docs/sdks/metrics/README.md#limits) - Get Metrics Limits

### [Oauth2](docs/sdks/oauth2/README.md)

* [Authorize](docs/sdks/oauth2/README.md#authorize) - Authorize
* [Token](docs/sdks/oauth2/README.md#token) - Request Token
* [Revoke](docs/sdks/oauth2/README.md#revoke) - Revoke Token
* [Introspect](docs/sdks/oauth2/README.md#introspect) - Introspect Token
* [Userinfo](docs/sdks/oauth2/README.md#userinfo) - Get User Info

#### [Oauth2.Clients](docs/sdks/clients/README.md)

* [List](docs/sdks/clients/README.md#list) - List Clients
* [Create](docs/sdks/clients/README.md#create) - Create Client
* [Get](docs/sdks/clients/README.md#get) - Get Client
* [Update](docs/sdks/clients/README.md#update) - Update Client
* [Delete](docs/sdks/clients/README.md#delete) - Delete Client

### [Orders](docs/sdks/orders/README.md)

* [List](docs/sdks/orders/README.md#list) - List Orders
* [Get](docs/sdks/orders/README.md#get) - Get Order
* [Invoice](docs/sdks/orders/README.md#invoice) - Get Order Invoice

### [Organizations](docs/sdks/organizations/README.md)

* [List](docs/sdks/organizations/README.md#list) - List Organizations
* [Create](docs/sdks/organizations/README.md#create) - Create Organization
* [Get](docs/sdks/organizations/README.md#get) - Get Organization
* [Update](docs/sdks/organizations/README.md#update) - Update Organization


### [Products](docs/sdks/products/README.md)

* [List](docs/sdks/products/README.md#list) - List Products
* [Create](docs/sdks/products/README.md#create) - Create Product
* [Get](docs/sdks/products/README.md#get) - Get Product
* [Update](docs/sdks/products/README.md#update) - Update Product
* [UpdateBenefits](docs/sdks/products/README.md#updatebenefits) - Update Product Benefits

### [Repositories](docs/sdks/repositories/README.md)

* [List](docs/sdks/repositories/README.md#list) - List Repositories
* [Get](docs/sdks/repositories/README.md#get) - Get Repository
* [Update](docs/sdks/repositories/README.md#update) - Update Repository

### [Subscriptions](docs/sdks/subscriptions/README.md)

* [List](docs/sdks/subscriptions/README.md#list) - List Subscriptions
* [Export](docs/sdks/subscriptions/README.md#export) - Export Subscriptions

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.ExternalOrganizations.List(ctx, operations.ExternalOrganizationsListRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceExternalOrganization != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/operations"
	"github.com/polarsource/polar-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.ExternalOrganizations.List(ctx, operations.ExternalOrganizationsListRequest{}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceExternalOrganization != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/operations"
	"github.com/polarsource/polar-go/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.ExternalOrganizations.List(ctx, operations.ExternalOrganizationsListRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceExternalOrganization != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `List` function may return the following errors:

| Error Type                    | Status Code | Content Type     |
| ----------------------------- | ----------- | ---------------- |
| apierrors.HTTPValidationError | 422         | application/json |
| apierrors.APIError            | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/apierrors"
	"github.com/polarsource/polar-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.ExternalOrganizations.List(ctx, operations.ExternalOrganizationsListRequest{})
	if err != nil {

		var e *apierrors.HTTPValidationError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Name

You can override the default server globally using the `WithServer(server string)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the names associated with the available servers:

| Name         | Server                         |
| ------------ | ------------------------------ |
| `production` | `https://api.polar.sh`         |
| `sandbox`    | `https://sandbox-api.polar.sh` |

#### Example

```go
package main

import (
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithServer("sandbox"),
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.ExternalOrganizations.List(ctx, operations.ExternalOrganizationsListRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceExternalOrganization != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithServerURL("https://api.polar.sh"),
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.ExternalOrganizations.List(ctx, operations.ExternalOrganizationsListRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceExternalOrganization != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/polarsource/polar-go&utm_campaign=go)
