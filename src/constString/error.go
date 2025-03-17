package constString

import (
	"encoding/json"
	"errors"
)

var ErrClosed = errors.New("jsonrpc2: connection is closed")

var JsonNull = json.RawMessage("null")
