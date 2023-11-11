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

	t.Run("Successful Folder Retrieval for Valid OrgID", func(t *testing.T) {
		// your test/s here
		req := &folders.FetchFolderRequest{OrgID: orgID}
        resp, err := folders.GetAllFolders(req)
        assert.NoError(t, err)
        assert.NotNil(t, resp)

// 		assert.Len(t, resp.Folders, 2)
//
// 		assert.Equal(t, "71702b42-aee8-4c03-a05c-1a0cc5102a86", resp.Folders[0].Id.String())
// 		assert.Equal(t, "sawa-test-2", resp.Folders[0].Name)
// 		assert.Equal(t, orgID, resp.Folders[0].OrgId)
// 		assert.True(t, resp.Folders[0].Deleted)

		 // Marshal the response into JSON for readability
        respJSON, err := json.MarshalIndent(resp, "", "    ")
        if err != nil {
            t.Fatalf("Failed to marshal response: %v", err)
        }
        fmt.Println("Response!!!!!:", string(respJSON))
	})

// 	t.Run("No Result for Non-Existing Organization ID", func(t *testing.T) {
// 		emptyOrgID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
// 		req := &folders.FetchFolderRequest{OrgID: emptyOrgID}
// 		resp, err := folders.GetAllFolders(req)
//
// 		assert.NoError(t, err)
// 		assert.NotNil(t, resp)
// 		assert.Empty(t, resp.Folders)
// 	})

// 	t.Run("ErrorHandling", func(t *testing.T) {
// 		randomOrgID, _ := uuid.NewV4()
// 		req := &folders.FetchFolderRequest{OrgID: randomOrgID}
// 		resp, err := folders.GetAllFolders(req)
//
// 		assert.Error(t, err)
// 		assert.Nil(t, resp)
// 	})

}

