package icalendar

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2

// icalparameter = altrepparam       ; Alternate text representation
// / cnparam           ; Common name
// / cutypeparam       ; Calendar user type
// / delfromparam      ; Delegator
// / deltoparam        ; Delegatee
// / dirparam          ; Directory entry
// / encodingparam     ; Inline encoding
// / fmttypeparam      ; Format type
// / fbtypeparam       ; Free/busy time type
// / languageparam     ; Language for text
// / memberparam       ; Group or list membership
// / partstatparam     ; Participation status
// / rangeparam        ; Recurrence identifier range
// / trigrelparam      ; Alarm trigger relationship
// / reltypeparam      ; Relationship type
// / roleparam         ; Participation role
// / rsvpparam         ; RSVP expectation
// / sentbyparam       ; Sent by
// / tzidparam         ; Reference to time zone object
// / valuetypeparam    ; Property value data type
// / other-param

// other-param   = (iana-param / x-param)

// iana-param  = iana-token "=" param-value *("," param-value)
// ; Some other IANA-registered iCalendar parameter.

// x-param     = x-name "=" param-value *("," param-value)
// ; A non-standard, experimental parameter.

// Applications MUST ignore x-param and iana-param values they don't
// recognize.

type Param struct {
	BEGIN string

	// Value Type:  TEXT
	PRODID string

	// Value Type:  TEXT
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.4
	VERSION string

	// Value Type:  TEXT
	DTSTART string

	// Value Type:  TEXT
	DTEND string

	// Value Type:  TEXT
	LOCATION string

	// Value Type:  TEXT
	CATEGORIES string

	// Value Type:  TEXT
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.5
	DESCRIPTION string

	// Value Type:  TEXT
	SUMMARY string

	// Value Type:  TEXT
	PRIORITY string
}

func Set(parameter Param) {

}
