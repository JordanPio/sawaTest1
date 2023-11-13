package folders

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started

import (
	"encoding/base64"
	"fmt"
	"encoding/json"
	"github.com/gofrs/uuid"
)

// Pagination parameters added to the request struct
type PaginatedFetchFolderRequest struct {
	OrgID  uuid.UUID
	Cursor string // Cursor is the pagination token
	Limit  int    // Limit is the number of items per page
}

// Pagination token and data added to the response struct
type PaginatedFetchFolderResponse struct {
	Folders []*Folder
	NextCursor string // NextCursor is the pagination token for the next page
}

// Pagination logic in GetAllFolders
func GetPaginatedAllFolders(req *PaginatedFetchFolderRequest) (*PaginatedFetchFolderResponse, error) {

	if req.Cursor == "END_OF_DATA" {
        return &PaginatedFetchFolderResponse{
            Folders:    []*Folder{}, // Empty slice
            NextCursor: "END_OF_DATA",
        }, nil
    }
    // Use GetAllFolders to fetch all folders
    nonPaginatedResponse, err := GetAllFolders(&FetchFolderRequest{OrgID: req.OrgID})
    if err != nil {
        return nil, fmt.Errorf("failed to fetch all folders: %w", err)
    }
    allFolders := nonPaginatedResponse.Folders

    // Parse the cursor to get the starting point for this page
    startingAfter, err := parsePaginationToken(req.Cursor)
    if err != nil {
        return nil, err // appropriate error handling
    }

    // Find the starting index based on the cursor's last ID
    startIndex := 0
    if startingAfter != nil {
        for i, folder := range allFolders {
            if folder.Id == startingAfter.LastID {
                startIndex = i + 1
                break
            }
        }
    }

    // Calculate the endIndex, making sure we don't go beyond the number of available folders
    endIndex := startIndex + req.Limit
    if endIndex > len(allFolders) {
        endIndex = len(allFolders)
    }

    // Create the paginated slice of folders
    foldersPage := allFolders[startIndex:endIndex]

    // Generate the next cursor token if there's more data
	isLastPage := endIndex >= len(allFolders)
	nextCursor := generatePaginationToken(allFolders[endIndex-1], isLastPage)

    // Return the paginated response
    return &PaginatedFetchFolderResponse{
        Folders:    foldersPage,
        NextCursor: nextCursor,
    }, nil
}
// PaginationTokenStruct defines the structure of the pagination token.

type PaginationTokenStruct struct {
	LastID uuid.UUID `json:"lastId"`
}

// Helper function to generate a pagination token from the last folder in a page
func generatePaginationToken(lastFolder *Folder, isLastPage bool) string {
    // If lastFolder is nil, return an empty string (no more data to paginate)
	if lastFolder == nil || isLastPage {
        return "END_OF_DATA"
    }

    // Generate token based on the last folder's ID
    tokenStruct := PaginationTokenStruct{
        LastID: lastFolder.Id,
    }
    tokenBytes, err := json.Marshal(tokenStruct)
    if err != nil {
        fmt.Printf("Error generating pagination token: %v\n", err)
        return ""
    }
    return base64.URLEncoding.EncodeToString(tokenBytes)
}

// Helper function to parse a pagination token back into the last ID
func parsePaginationToken(token string) (*PaginationTokenStruct, error) {
	if token == "" {
		return nil, nil // No cursor provided, start from the beginning
	}

	tokenBytes, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	var tokenStruct PaginationTokenStruct
	err = json.Unmarshal(tokenBytes, &tokenStruct)
	if err != nil {
		return nil, err
	}

	return &tokenStruct, nil
}