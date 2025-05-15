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

Read the docs at https://docs.polar.sh/api-reference
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
			CreatedAt:              types.MustTimeFromString("2025-11-12T14:26:42.882Z"),
			ModifiedAt:             types.MustNewTimeFromString("2024-05-27T05:08:06.235Z"),
			ID:                     "<value>",
			PaymentProcessor:       components.PaymentProcessorStripe,
			Status:                 components.CheckoutStatusFailed,
			ClientSecret:           "<value>",
			URL:                    "https://heavy-beret.com/",
			ExpiresAt:              types.MustTimeFromString("2023-02-25T02:26:48.460Z"),
			SuccessURL:             "https://sardonic-final.info/",
			EmbedOrigin:            polargo.String("<value>"),
			Amount:                 962818,
			DiscountAmount:         6400,
			NetAmount:              426090,
			TaxAmount:              polargo.Int64(648726),
			TotalAmount:            210702,
			Currency:               "Riel",
			ProductID:              "<value>",
			ProductPriceID:         "<value>",
			DiscountID:             polargo.String("<value>"),
			AllowDiscountCodes:     false,
			RequireBillingAddress:  false,
			IsDiscountApplicable:   false,
			IsFreeProductPrice:     false,
			IsPaymentRequired:      false,
			IsPaymentSetupRequired: false,
			IsPaymentFormRequired:  false,
			CustomerID:             polargo.String("<value>"),
			CustomerName:           polargo.String("<value>"),
			CustomerEmail:          polargo.String("<value>"),
			CustomerIPAddress:      polargo.String("<value>"),
			CustomerBillingAddress: &components.Address{
				Country: "FR",
			},
			CustomerTaxID: polargo.String("<id>"),
			PaymentProcessorMetadata: map[string]string{
				"key": "<value>",
			},
			Metadata: map[string]components.CheckoutMetadata{
				"key": components.CreateCheckoutMetadataInteger(
					366960,
				),
				"key1": components.CreateCheckoutMetadataInteger(
					961375,
				),
			},
			CustomerExternalID: polargo.String("<id>"),
			Products: []components.CheckoutProduct{
				components.CheckoutProduct{
					CreatedAt:         types.MustTimeFromString("2023-01-21T11:17:02.173Z"),
					ModifiedAt:        types.MustNewTimeFromString("2024-08-14T19:11:10.463Z"),
					ID:                "<value>",
					Name:              "<value>",
					Description:       polargo.String("accessorise ack pave"),
					RecurringInterval: components.SubscriptionRecurringIntervalMonth.ToPointer(),
					IsRecurring:       false,
					IsArchived:        false,
					OrganizationID:    "<value>",
					Prices: []components.CheckoutProductPrices{
						components.CreateCheckoutProductPricesProductPrice(
							components.CreateProductPriceProductPriceMeteredUnit(
								components.ProductPriceMeteredUnit{
									CreatedAt:         types.MustTimeFromString("2024-12-13T18:25:33.693Z"),
									ModifiedAt:        types.MustNewTimeFromString("2023-01-09T04:38:53.436Z"),
									ID:                "<value>",
									IsArchived:        false,
									ProductID:         "<value>",
									Type:              components.ProductPriceTypeRecurring,
									RecurringInterval: components.SubscriptionRecurringIntervalMonth.ToPointer(),
									PriceCurrency:     "<value>",
									UnitAmount:        "<value>",
									CapAmount:         polargo.Int64(199664),
									MeterID:           "<value>",
									Meter: components.ProductPriceMeter{
										ID:   "<value>",
										Name: "<value>",
									},
								},
							),
						),
						components.CreateCheckoutProductPricesLegacyRecurringProductPrice(
							components.CreateLegacyRecurringProductPriceFree(
								components.LegacyRecurringProductPriceFree{
									CreatedAt:         types.MustTimeFromString("2023-01-25T05:44:56.791Z"),
									ModifiedAt:        types.MustNewTimeFromString("2025-05-25T15:20:50.694Z"),
									ID:                "<value>",
									IsArchived:        false,
									ProductID:         "<value>",
									RecurringInterval: components.SubscriptionRecurringIntervalYear,
								},
							),
						),
					},
					Benefits: []components.BenefitPublic{},
					Medias: []components.ProductMediaFileRead{
						components.ProductMediaFileRead{
							ID:                   "<value>",
							OrganizationID:       "<value>",
							Name:                 "<value>",
							Path:                 "/media",
							MimeType:             "<value>",
							Size:                 757308,
							StorageVersion:       polargo.String("<value>"),
							ChecksumEtag:         polargo.String("<value>"),
							ChecksumSha256Base64: polargo.String("<value>"),
							ChecksumSha256Hex:    polargo.String("<value>"),
							LastModifiedAt:       types.MustNewTimeFromString("2024-07-16T16:32:06.101Z"),
							Version:              polargo.String("<value>"),
							IsUploaded:           false,
							CreatedAt:            types.MustTimeFromString("2025-12-08T10:05:37.208Z"),
							SizeReadable:         "<value>",
							PublicURL:            "https://separate-pinstripe.com",
						},
					},
				},
			},
			Product: components.CheckoutProduct{
				CreatedAt:         types.MustTimeFromString("2025-08-22T02:17:20.661Z"),
				ModifiedAt:        types.MustNewTimeFromString("2025-09-12T19:41:14.261Z"),
				ID:                "<value>",
				Name:              "<value>",
				Description:       polargo.String("darling lined nephew crank"),
				RecurringInterval: components.SubscriptionRecurringIntervalMonth.ToPointer(),
				IsRecurring:       true,
				IsArchived:        false,
				OrganizationID:    "<value>",
				Prices:            []components.CheckoutProductPrices{},
				Benefits: []components.BenefitPublic{
					components.BenefitPublic{
						ID:             "<value>",
						CreatedAt:      types.MustTimeFromString("2025-11-18T07:12:26.749Z"),
						ModifiedAt:     types.MustNewTimeFromString("2024-08-14T14:29:29.201Z"),
						Type:           components.BenefitTypeGithubRepository,
						Description:    "phew empty flawed worth radiant anti of blissfully",
						Selectable:     false,
						Deletable:      false,
						OrganizationID: "<value>",
					},
					components.BenefitPublic{
						ID:             "<value>",
						CreatedAt:      types.MustTimeFromString("2024-03-12T06:20:21.161Z"),
						ModifiedAt:     types.MustNewTimeFromString("2024-06-27T20:20:27.400Z"),
						Type:           components.BenefitTypeLicenseKeys,
						Description:    "ack neatly jagged as weekly drat nor why",
						Selectable:     true,
						Deletable:      false,
						OrganizationID: "<value>",
					},
					components.BenefitPublic{
						ID:             "<value>",
						CreatedAt:      types.MustTimeFromString("2024-03-31T07:42:35.077Z"),
						ModifiedAt:     types.MustNewTimeFromString("2025-04-30T13:13:01.707Z"),
						Type:           components.BenefitTypeGithubRepository,
						Description:    "ironclad chilly seriously finger unbearably and bourgeoisie deplore unless",
						Selectable:     false,
						Deletable:      true,
						OrganizationID: "<value>",
					},
				},
				Medias: []components.ProductMediaFileRead{
					components.ProductMediaFileRead{
						ID:                   "<value>",
						OrganizationID:       "<value>",
						Name:                 "<value>",
						Path:                 "/usr/share",
						MimeType:             "<value>",
						Size:                 366209,
						StorageVersion:       polargo.String("<value>"),
						ChecksumEtag:         polargo.String("<value>"),
						ChecksumSha256Base64: polargo.String("<value>"),
						ChecksumSha256Hex:    polargo.String("<value>"),
						LastModifiedAt:       types.MustNewTimeFromString("2024-09-08T17:09:22.377Z"),
						Version:              polargo.String("<value>"),
						IsUploaded:           false,
						CreatedAt:            types.MustTimeFromString("2025-01-15T08:45:21.256Z"),
						SizeReadable:         "<value>",
						PublicURL:            "https://cumbersome-seafood.net",
					},
				},
			},
			ProductPrice: components.CreateCheckoutProductPriceProductPrice(
				components.CreateProductPriceProductPriceCustom(
					components.ProductPriceCustom{
						CreatedAt:         types.MustTimeFromString("2024-12-18T06:26:25.293Z"),
						ModifiedAt:        types.MustNewTimeFromString("2025-03-28T00:06:24.086Z"),
						ID:                "<value>",
						IsArchived:        false,
						ProductID:         "<value>",
						Type:              components.ProductPriceTypeOneTime,
						RecurringInterval: components.SubscriptionRecurringIntervalYear.ToPointer(),
						PriceCurrency:     "<value>",
						MinimumAmount:     polargo.Int64(476477),
						MaximumAmount:     polargo.Int64(729611),
						PresetAmount:      polargo.Int64(905974),
					},
				),
			),
			Discount: polargo.Pointer(components.CreateCheckoutDiscountCheckoutDiscountPercentageRepeatDuration(
				components.CheckoutDiscountPercentageRepeatDuration{
					Duration:         components.DiscountDurationForever,
					DurationInMonths: 928294,
					Type:             components.DiscountTypeFixed,
					BasisPoints:      887018,
					ID:               "<value>",
					Name:             "<value>",
					Code:             polargo.String("<value>"),
				},
			)),
			SubscriptionID: polargo.String("<value>"),
			AttachedCustomFields: []components.AttachedCustomField{
				components.AttachedCustomField{
					CustomFieldID: "<value>",
					CustomField: components.CreateCustomFieldNumber(
						components.CustomFieldNumber{
							CreatedAt:  types.MustTimeFromString("2025-10-29T00:32:24.235Z"),
							ModifiedAt: types.MustNewTimeFromString("2025-04-28T15:46:25.145Z"),
							ID:         "<value>",
							Metadata: map[string]components.CustomFieldNumberMetadata{
								"key": components.CreateCustomFieldNumberMetadataNumber(
									7295.4,
								),
							},
							Slug:           "<value>",
							Name:           "<value>",
							OrganizationID: "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
							Properties:     components.CustomFieldNumberProperties{},
						},
					),
					Order:    192934,
					Required: true,
				},
				components.AttachedCustomField{
					CustomFieldID: "<value>",
					CustomField: components.CreateCustomFieldSelect(
						components.CustomFieldSelect{
							CreatedAt:  types.MustTimeFromString("2024-06-12T07:35:26.368Z"),
							ModifiedAt: types.MustNewTimeFromString("2023-01-01T03:32:54.177Z"),
							ID:         "<value>",
							Metadata: map[string]components.CustomFieldSelectMetadata{
								"key": components.CreateCustomFieldSelectMetadataInteger(
									882089,
								),
							},
							Slug:           "<value>",
							Name:           "<value>",
							OrganizationID: "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
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
					Order:    597839,
					Required: false,
				},
				components.AttachedCustomField{
					CustomFieldID: "<value>",
					CustomField: components.CreateCustomFieldDate(
						components.CustomFieldDate{
							CreatedAt:  types.MustTimeFromString("2023-10-09T19:13:38.530Z"),
							ModifiedAt: types.MustNewTimeFromString("2025-03-10T05:19:02.590Z"),
							ID:         "<value>",
							Metadata: map[string]components.CustomFieldDateMetadata{
								"key": components.CreateCustomFieldDateMetadataNumber(
									6244.25,
								),
							},
							Slug:           "<value>",
							Name:           "<value>",
							OrganizationID: "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
							Properties:     components.CustomFieldDateProperties{},
						},
					),
					Order:    660585,
					Required: false,
				},
			},
			CustomerMetadata: map[string]components.CustomerMetadata{
				"key": components.CreateCustomerMetadataBoolean(
					true,
				),
			},
			CustomerBillingAddressFields: components.CheckoutCustomerBillingAddressFields{
				Country:    true,
				State:      false,
				City:       true,
				PostalCode: false,
				Line1:      false,
				Line2:      false,
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
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.Organizations.List(ctx, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceOrganization != nil {
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

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
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

	s := polargo.New()

	res, err := s.CustomerPortal.BenefitGrants.List(ctx, operations.CustomerPortalBenefitGrantsListRequest{
		OrganizationID: polargo.Pointer(operations.CreateCustomerPortalBenefitGrantsListQueryParamOrganizationIDFilterArrayOfStr(
			[]string{
				"1dbfc517-0bbf-4301-9ba8-555ca42b9737",
			},
		)),
	}, operations.CustomerPortalBenefitGrantsListSecurity{
		CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceCustomerBenefitGrant != nil {
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

* [List](docs/sdks/checkouts/README.md#list) - List Checkout Sessions
* [Create](docs/sdks/checkouts/README.md#create) - Create Checkout Session
* [Get](docs/sdks/checkouts/README.md#get) - Get Checkout Session
* [Update](docs/sdks/checkouts/README.md#update) - Update Checkout Session
* [ClientGet](docs/sdks/checkouts/README.md#clientget) - Get Checkout Session from Client
* [ClientUpdate](docs/sdks/checkouts/README.md#clientupdate) - Update Checkout Session from Client
* [ClientConfirm](docs/sdks/checkouts/README.md#clientconfirm) - Confirm Checkout Session from Client

### [CustomerMeters](docs/sdks/customermeters/README.md)

* [List](docs/sdks/customermeters/README.md#list) - List Customer Meters
* [Get](docs/sdks/customermeters/README.md#get) - Get Customer Meter

### [CustomerPortal](docs/sdks/customerportal/README.md)


#### [CustomerPortal.BenefitGrants](docs/sdks/benefitgrants/README.md)

* [List](docs/sdks/benefitgrants/README.md#list) - List Benefit Grants
* [Get](docs/sdks/benefitgrants/README.md#get) - Get Benefit Grant
* [Update](docs/sdks/benefitgrants/README.md#update) - Update Benefit Grant

#### [CustomerPortal.CustomerMeters](docs/sdks/polarcustomermeters/README.md)

* [List](docs/sdks/polarcustomermeters/README.md#list) - List Meters
* [Get](docs/sdks/polarcustomermeters/README.md#get) - Get Customer Meter

#### [CustomerPortal.Customers](docs/sdks/polarcustomers/README.md)

* [Get](docs/sdks/polarcustomers/README.md#get) - Get Customer
* [Update](docs/sdks/polarcustomers/README.md#update) - Update Customer
* [GetPaymentMethods](docs/sdks/polarcustomers/README.md#getpaymentmethods) - Get Customer Payment Methods
* [AddPaymentMethod](docs/sdks/polarcustomers/README.md#addpaymentmethod) - Add Customer Payment Method
* [DeletePaymentMethod](docs/sdks/polarcustomers/README.md#deletepaymentmethod) - Delete Customer Payment Method

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
* [GetExternal](docs/sdks/customers/README.md#getexternal) - Get Customer by External ID
* [UpdateExternal](docs/sdks/customers/README.md#updateexternal) - Update Customer by External ID
* [DeleteExternal](docs/sdks/customers/README.md#deleteexternal) - Delete Customer by External ID
* [GetState](docs/sdks/customers/README.md#getstate) - Get Customer State
* [GetStateExternal](docs/sdks/customers/README.md#getstateexternal) - Get Customer State by External ID

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

### [Events](docs/sdks/events/README.md)

* [List](docs/sdks/events/README.md#list) - List Events
* [ListNames](docs/sdks/events/README.md#listnames) - List Event Names
* [Get](docs/sdks/events/README.md#get) - Get Event
* [Ingest](docs/sdks/events/README.md#ingest) - Ingest Events

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

### [Meters](docs/sdks/meters/README.md)

* [List](docs/sdks/meters/README.md#list) - List Meters
* [Create](docs/sdks/meters/README.md#create) - Create Meter
* [Get](docs/sdks/meters/README.md#get) - Get Meter
* [Update](docs/sdks/meters/README.md#update) - Update Meter
* [Quantities](docs/sdks/meters/README.md#quantities) - Get Meter Quantities

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

### [Payments](docs/sdks/payments/README.md)

* [List](docs/sdks/payments/README.md#list) - List Payments
* [Get](docs/sdks/payments/README.md#get) - Get Payment


### [Products](docs/sdks/products/README.md)

* [List](docs/sdks/products/README.md#list) - List Products
* [Create](docs/sdks/products/README.md#create) - Create Product
* [Get](docs/sdks/products/README.md#get) - Get Product
* [Update](docs/sdks/products/README.md#update) - Update Product
* [UpdateBenefits](docs/sdks/products/README.md#updatebenefits) - Update Product Benefits

### [Refunds](docs/sdks/refunds/README.md)

* [List](docs/sdks/refunds/README.md#list) - List Refunds
* [Create](docs/sdks/refunds/README.md#create) - Create Refund

### [Subscriptions](docs/sdks/subscriptions/README.md)

* [List](docs/sdks/subscriptions/README.md#list) - List Subscriptions
* [Export](docs/sdks/subscriptions/README.md#export) - Export Subscriptions
* [Get](docs/sdks/subscriptions/README.md#get) - Get Subscription
* [Update](docs/sdks/subscriptions/README.md#update) - Update Subscription
* [Revoke](docs/sdks/subscriptions/README.md#revoke) - Revoke Subscription

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
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.Organizations.List(ctx, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceOrganization != nil {
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

	res, err := s.Organizations.List(ctx, nil, nil, nil, nil, operations.WithRetries(
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
	if res.ListResourceOrganization != nil {
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

	res, err := s.Organizations.List(ctx, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceOrganization != nil {
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
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.Organizations.List(ctx, nil, nil, nil, nil)
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

| Name         | Server                         | Description            |
| ------------ | ------------------------------ | ---------------------- |
| `production` | `https://api.polar.sh`         | Production environment |
| `sandbox`    | `https://sandbox-api.polar.sh` | Sandbox environment    |

#### Example

```go
package main

import (
	"context"
	polargo "github.com/polarsource/polar-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithServer("sandbox"),
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.Organizations.List(ctx, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceOrganization != nil {
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
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithServerURL("https://api.polar.sh"),
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.Organizations.List(ctx, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceOrganization != nil {
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
