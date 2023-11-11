package folders_test

import (
	"testing"
	"fmt"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/stretchr/testify/assert"
)


func Test_GetAllFolders(t *testing.T) {
	orgID, err := uuid.FromString("3b9a868b-8cd9-4b6b-ba23-fd1e08f3e9fa")
	if err != nil {
		t.Fatalf("Failed to parse orgID: %v", err)
	}

	t.Run("test", func(t *testing.T) {
		// your test/s here
		req := &folders.FetchFolderRequest{OrgID: orgID}
        resp, err := folders.GetAllFolders(req)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
		 // Marshal the response into JSON for readability
        respJSON, err := json.MarshalIndent(resp, "", "    ")
        if err != nil {
            t.Fatalf("Failed to marshal response: %v", err)
        }
        fmt.Println("Response!!!!!:", string(respJSON))
	})
}

