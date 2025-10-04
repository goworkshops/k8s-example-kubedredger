# Exercise: Report Status, use conditions

Follow these steps to implement status reporting:

**Note**: existing tests are *broken*, and that's fine.
Don't feel constrained to reimplement the API they expect.
Feel free to rewrite them as you see fit.

## 1. Implement status reporting

Enhance the code in `internal/controller/configuration_controller.go` to fill the `ConfigurationStatus`.
You may want to fill and use the functions in `internal/controller/conversion.go`.

## 2. Implement the conditions

Once you report the status, implement the reporting of the `conf.Status.Conditions`.
Make sure to report at least these condition **types**:

```go
const (
       ConditionAvailable   = "Available"
       ConditionProgressing = "Progressing"
       ConditionDegraded    = "Degraded"
)
```

You may want to use these reasons:

```go
const (
       ConditionReasonAsExpected      = "AsExpected"
       ConditionReasonUpToDate        = "UpToDate"
       ConditionReasonWriteError      = "WriteError"
       ConditionReasonUpdatingContent = "UpdatingContent"
)
```

Here's an abridged definition of condition:

```go
type Condition struct {
	// type of condition in CamelCase or in foo.example.com/CamelCase.
	// ---
	// Many .condition.type values are consistent across resources like Available, but because arbitrary conditions can be
	// useful (see .node.status.conditions), the ability to deconflict is important.
	// The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)
	// +required
	Type string
	// status of the condition, one of True, False, Unknown.
	// +required
	Status ConditionStatus
	// observedGeneration represents the .metadata.generation that the condition was set based upon.
	// For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date
	// with respect to the current state of the instance.
	// +optional
	ObservedGeneration int64
	// lastTransitionTime is the last time the condition transitioned from one status to another.
	// This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.
	// +required
	LastTransitionTime Time
	// reason contains a programmatic identifier indicating the reason for the condition's last transition.
	// Producers of specific condition types may define expected values and meanings for this field,
	// and whether the values are considered a guaranteed API.
	// The value should be a CamelCase string.
	// This field may not be empty.
	// +required
	Reason string
	// message is a human readable message indicating details about the transition.
	// This may be an empty string.
	// +required
	Message string
}
```
