# CustomerPortalSubscriptionsUpdateResponse403CustomerPortalSubscriptionsUpdate

Customer subscription is already canceled or will be at the end of the period, the user lacks billing permissions, or pausing/resuming is not enabled for the organization.


## Supported Types

### AlreadyCanceledSubscription

```go
customerPortalSubscriptionsUpdateResponse403CustomerPortalSubscriptionsUpdate := apierrors.CreateCustomerPortalSubscriptionsUpdateResponse403CustomerPortalSubscriptionsUpdateAlreadyCanceledSubscription(components.AlreadyCanceledSubscription{/* values here */})
```

### PauseResumeNotAllowed

```go
customerPortalSubscriptionsUpdateResponse403CustomerPortalSubscriptionsUpdate := apierrors.CreateCustomerPortalSubscriptionsUpdateResponse403CustomerPortalSubscriptionsUpdatePauseResumeNotAllowed(components.PauseResumeNotAllowed{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPortalSubscriptionsUpdateResponse403CustomerPortalSubscriptionsUpdate.Type {
	case apierrors.CustomerPortalSubscriptionsUpdateResponse403CustomerPortalSubscriptionsUpdateTypeAlreadyCanceledSubscription:
		// customerPortalSubscriptionsUpdateResponse403CustomerPortalSubscriptionsUpdate.AlreadyCanceledSubscription is populated
	case apierrors.CustomerPortalSubscriptionsUpdateResponse403CustomerPortalSubscriptionsUpdateTypePauseResumeNotAllowed:
		// customerPortalSubscriptionsUpdateResponse403CustomerPortalSubscriptionsUpdate.PauseResumeNotAllowed is populated
}
```
