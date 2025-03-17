package modelStruct

import (
	"bytes"
	"encoding/json"
	"errors"
)

type AnyMessage struct {
	Response  *ResponseModel
	Request   *RequestModel
	Broadcast *RPCBroadcastModel
}

func (m AnyMessage) MarshalJSON() ([]byte, error) {
	var v interface{}
	switch {
	case m.Request != nil && m.Response == nil:
		v = m.Request
	case m.Request == nil && m.Response != nil:
		v = m.Response
	}
	if v != nil {
		return json.Marshal(v)
	}
	return nil, errors.New("jsonrpc2: message must have exactly one of the request or response fields set")
}

func (m *AnyMessage) UnmarshalJSON(data []byte) error {
	// The presence of these fields distinguishes between the 2
	// message types.
	type msg struct {
		ID     interface{}              `json:"id"`
		Method *string                  `json:"method"`
		Result anyValueWithExplicitNull `json:"result"`
		Error  interface{}              `json:"error"`
	}

	var isRequest, isResponse bool
	checkType := func(m *msg) error {
		mIsRequest := m.Method != nil
		mIsResponse := m.Result.null || m.Result.value != nil || m.Error != nil
		if (!mIsRequest && !mIsResponse) || (mIsRequest && mIsResponse) {
			return errors.New("jsonrpc2: unable to determine message type (request or response)")
		}
		if (mIsRequest && isResponse) || (mIsResponse && isRequest) {
			return errors.New("jsonrpc2: batch message type mismatch (must be all requests or all responses)")
		}
		isRequest = mIsRequest
		isResponse = mIsResponse
		return nil
	}

	if isArray := len(data) > 0 && data[0] == '['; isArray {
		var msgs []msg
		if err := json.Unmarshal(data, &msgs); err != nil {
			return err
		}
		if len(msgs) == 0 {
			return errors.New("jsonrpc2: invalid empty batch")
		}
		for i := range msgs {
			if err := checkType(&msg{
				ID:     msgs[i].ID,
				Method: msgs[i].Method,
				Result: msgs[i].Result,
				Error:  msgs[i].Error,
			}); err != nil {
				return err
			}
		}
	} else {
		var m msg
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}
		if err := checkType(&m); err != nil {
			return err
		}
	}

	var v interface{}
	switch {
	case isRequest && !isResponse:
		v = &m.Broadcast
	case !isRequest && isResponse:
		v = &m.Response
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	/*	if !isRequest && isResponse && m.Response.Error == nil && m.Response.Result == nil {
		m.response.Result = &jsonNull
	}*/
	return nil
}

// anyValueWithExplicitNull is used to distinguish {} from
// {"result":null} by anyMessage's JSON unmarshaler.
type anyValueWithExplicitNull struct {
	null  bool // JSON "null"
	value interface{}
}

func (v anyValueWithExplicitNull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *anyValueWithExplicitNull) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if string(data) == "null" {
		*v = anyValueWithExplicitNull{null: true}
		return nil
	}
	*v = anyValueWithExplicitNull{}
	return json.Unmarshal(data, &v.value)
}
