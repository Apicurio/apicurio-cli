package remote

import (
	"context"
	"testing"

	"github.com/apicurio/apicurio-cli/internal/mockutil"
)

func TestGetRemoteServiceConstants(t *testing.T) {
	err, constants := GetRemoteServiceConstants(context.Background(), mockutil.NewLoggerMock())

	if err != nil {
		t.Errorf("GetRemoteServiceConstants() failed with error %s", err)
	}

	if constants == nil {
		t.Errorf("GetRemoteServiceConstants() returned nil")
	}

	t.Logf("GetRemoteServiceConstants() returned %+v", constants)
}
