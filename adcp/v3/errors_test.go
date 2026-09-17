package adcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultRecoveryUsesSchemaValues(t *testing.T) {
	assert.Equal(t, "transient", defaultRecovery("RATE_LIMITED"))
	assert.Equal(t, "transient", defaultRecovery("SERVICE_UNAVAILABLE"))
	assert.Equal(t, "correctable", defaultRecovery("INVALID_REQUEST"))
	assert.Equal(t, "correctable", defaultRecovery("MISSING_FIELD"))
	assert.Equal(t, "correctable", defaultRecovery("BUDGET_TOO_LOW"))
	assert.Equal(t, "terminal", defaultRecovery("INTERNAL_ERROR"))
	assert.Equal(t, "terminal", defaultRecovery("UNKNOWN_CODE"))
}
