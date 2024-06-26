package fetch

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/cdvelop/vanify/pkg/cdproto/cdp"
	"github.com/cdvelop/vanify/pkg/cdproto/network"
)

// EventRequestPaused issued when the domain is enabled and the request URL
// matches the specified filter. The request is paused until the client responds
// with one of continueRequest, failRequest or fulfillRequest. The stage of the
// request can be determined by presence of responseErrorReason and
// responseStatusCode -- the request is at the response stage if either of these
// fields is present and in the request stage otherwise. Redirect responses and
// subsequent requests are reported similarly to regular responses and requests.
// Redirect responses may be distinguished by the value of responseStatusCode
// (which is one of 301, 302, 303, 307, 308) along with presence of the location
// header. Requests resulting from a redirect will have redirectedRequestId
// field set.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#event-requestPaused
type EventRequestPaused struct {
	RequestID           RequestID            `json:"requestId"`                     // Each request the page makes will have a unique id.
	Request             *network.Request     `json:"request"`                       // The details of the request.
	FrameID             cdp.FrameID          `json:"frameId"`                       // The id of the frame that initiated the request.
	ResourceType        network.ResourceType `json:"resourceType"`                  // How the requested resource will be used.
	ResponseErrorReason network.ErrorReason  `json:"responseErrorReason,omitempty"` // Response error if intercepted at response stage.
	ResponseStatusCode  int64                `json:"responseStatusCode,omitempty"`  // Response code if intercepted at response stage.
	ResponseStatusText  string               `json:"responseStatusText,omitempty"`  // Response status text if intercepted at response stage.
	ResponseHeaders     []*HeaderEntry       `json:"responseHeaders,omitempty"`     // Response headers if intercepted at the response stage.
	NetworkID           network.RequestID    `json:"networkId,omitempty"`           // If the intercepted request had a corresponding Network.requestWillBeSent event fired for it, then this networkId will be the same as the requestId present in the requestWillBeSent event.
	RedirectedRequestID RequestID            `json:"redirectedRequestId,omitempty"` // If the request is due to a redirect response from the server, the id of the request that has caused the redirect.
}

// EventAuthRequired issued when the domain is enabled with
// handleAuthRequests set to true. The request is paused until client responds
// with continueWithAuth.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#event-authRequired
type EventAuthRequired struct {
	RequestID     RequestID            `json:"requestId"`     // Each request the page makes will have a unique id.
	Request       *network.Request     `json:"request"`       // The details of the request.
	FrameID       cdp.FrameID          `json:"frameId"`       // The id of the frame that initiated the request.
	ResourceType  network.ResourceType `json:"resourceType"`  // How the requested resource will be used.
	AuthChallenge *AuthChallenge       `json:"authChallenge"` // Details of the Authorization Challenge encountered. If this is set, client should respond with continueRequest that contains AuthChallengeResponse.
}
