package configset_json

import (
	"context"
	"encoding/json"

	"github.com/aperturerobotics/controllerbus/bus"
	"github.com/aperturerobotics/controllerbus/config"
	"github.com/aperturerobotics/controllerbus/controller/resolver"
	"github.com/pkg/errors"
	// "github.com/aperturerobotics/controllerbus/controller/configset"
)

// Config implements the JSON unmarshaling and marshaling logic for a configset
// Config.
type Config struct {
	pendingParseData []byte
	underlying       config.Config
}

// Resolve constructs the underlying config from the pending parse data.
func (c *Config) Resolve(ctx context.Context, configID string, b bus.Bus) error {
	configCtorDir := resolver.NewLoadConfigConstructorByID(configID)
	configCtorVal, configCtorRef, err := bus.ExecOneOff(ctx, b, configCtorDir, nil)
	if err != nil {
		return errors.WithMessage(err, "resolve config object")
	}
	defer configCtorRef.Release()

	ctor, ctorOk := configCtorVal.(config.Constructor)
	if !ctorOk {
		return errors.New("load config constructor directive returned invalid object")
	}
	c.underlying = ctor.ConstructConfig()
	if err := json.Unmarshal(c.pendingParseData, c.underlying); err != nil {
		return err
	}
	c.pendingParseData = nil
	return nil
}

// UnmarshalJSON unmarshals a controller config JSON blob pushing the data into
// the pending parse buffer.
func (c *Config) UnmarshalJSON(data []byte) error {
	// assert that the object is a map
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	c.pendingParseData = data
	return nil
}

// MarshalJSON marshals a controller config JSON blob.
func (c *Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.underlying)
}

// GetConfig returns the underlying config after Resolve.
func (c *Config) GetConfig() config.Config {
	return c.underlying
}

// _ is a type assertion
var _ json.Unmarshaler = ((*Config)(nil))

// _ is a type assertion
var _ json.Marshaler = ((*Config)(nil))
