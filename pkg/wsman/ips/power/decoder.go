/*********************************************************************
 * Copyright (c) Intel Corporation 2025
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package power

const (
	IPSPowerManagementService       string = "IPS_PowerManagementService"
	RequestOSPowerSavingStateChange string = "RequestOSPowerSavingStateChange"
	ValueNotFound                   string = "Value not found in map"
)

// TODO: This list of contants needs to be scrubbed.
const (
	// Unknown.
	Unknown OSPowerSavingState = 0 // Unknown OS Power Saving State

	// Unsupported.
	Unsupported OSPowerSavingState = 1 // Unsupported

	// Full Power.
	FullPower OSPowerSavingState = 2 // Full Power

	// OS power saving.
	OSPowerSaving OSPowerSavingState = 3 // OS power saving
)

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
)

// enabledStateMap is a map of the EnabledState enumeration.
var enabledStateMap = map[EnabledState]string{
	EnabledStateUnknown:           "Unknown",
	EnabledStateOther:             "Other",
	EnabledStateEnabled:           "Enabled",
	EnabledStateDisabled:          "Disabled",
	EnabledStateShuttingDown:      "ShuttingDown",
	EnabledStateNotApplicable:     "NotApplicable",
	EnabledStateEnabledButOffline: "EnabledButOffline",
	EnabledStateInTest:            "InTest",
	EnabledStateDeferred:          "Deferred",
	EnabledStateQuiesce:           "Quiesce",
	EnabledStateStarting:          "Starting",
}

// String returns a human-readable string representation of the EnabledState enumeration.
func (e EnabledState) String() string {
	if s, ok := enabledStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	RequestedStateUnknown       RequestedState = 0
	RequestedStateEnabled       RequestedState = 2
	RequestedStateDisabled      RequestedState = 3
	RequestedStateShutDown      RequestedState = 4
	RequestedStateNoChange      RequestedState = 5
	RequestedStateOffline       RequestedState = 6
	RequestedStateTest          RequestedState = 7
	RequestedStateDeferred      RequestedState = 8
	RequestedStateQuiesce       RequestedState = 9
	RequestedStateReboot        RequestedState = 10
	RequestedStateReset         RequestedState = 11
	RequestedStateNotApplicable RequestedState = 12
)

// requestedStateMap is a map of the RequestedState enumeration.
var requestedStateMap = map[RequestedState]string{
	RequestedStateUnknown:       "Unknown",
	RequestedStateEnabled:       "Enabled",
	RequestedStateDisabled:      "Disabled",
	RequestedStateShutDown:      "ShutDown",
	RequestedStateNoChange:      "NoChange",
	RequestedStateOffline:       "Offline",
	RequestedStateTest:          "Test",
	RequestedStateDeferred:      "Deferred",
	RequestedStateQuiesce:       "Quiesce",
	RequestedStateReboot:        "Reboot",
	RequestedStateReset:         "Reset",
	RequestedStateNotApplicable: "NotApplicable",
}

// String returns a human-readable string representation of the RequestedState enumeration.
func (e RequestedState) String() string {
	if s, ok := requestedStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	ReturnValueCompletedWithNoError              ReturnValue = 0
	ReturnValueMethodNotSupported                ReturnValue = 1
	ReturnValueUnknownError                      ReturnValue = 2
	ReturnValueCannotCompleteWithinTimeoutPeriod ReturnValue = 3
	ReturnValueFailed                            ReturnValue = 4
	ReturnValueInvalidParameter                  ReturnValue = 5
	ReturnValueInUse                             ReturnValue = 6
	ReturnValueMethodParametersCheckedJobStarted ReturnValue = 4096
	ReturnValueInvalidStateTransition            ReturnValue = 4097
	ReturnValueUseOfTimeoutParameterNotSupported ReturnValue = 4098
	ReturnValueBusy                              ReturnValue = 4099
)

// returnValueMap is a map of the ReturnValue enumeration.
var returnValueMap = map[ReturnValue]string{
	ReturnValueCompletedWithNoError:              "CompletedWithNoError",
	ReturnValueMethodNotSupported:                "MethodNotSupported",
	ReturnValueUnknownError:                      "UnknownError",
	ReturnValueCannotCompleteWithinTimeoutPeriod: "CannotCompleteWithinTimeoutPeriod",
	ReturnValueFailed:                            "Failed",
	ReturnValueInvalidParameter:                  "InvalidParameter",
	ReturnValueInUse:                             "InUse",
	ReturnValueMethodParametersCheckedJobStarted: "MethodParametersCheckedJobStarted",
	ReturnValueInvalidStateTransition:            "InvalidStateTransition",
	ReturnValueUseOfTimeoutParameterNotSupported: "UseOfTimeoutParameterNotSupported",
	ReturnValueBusy:                              "Busy",
}

// String returns a human-readable string representation of the ReturnValue enumeration.
func (e ReturnValue) String() string {
	if s, ok := returnValueMap[e]; ok {
		return s
	}

	return ValueNotFound
}
