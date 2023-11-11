package folders

import (
	"github.com/gofrs/uuid"
	"fmt"
)

// this function retrieves all folders related to an organization ID returning in the response and error
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	 // Fetch all folders by organization ID and handle any errors
    folders, err := FetchAllFoldersByOrgID(req.OrgID)
    if err != nil {
        // Return the error to the caller
        return nil, fmt.Errorf("failed to fetch folders by org ID: %w", err)
    }

    // Create the FetchFolderResponse with the slice 'folders'
    response := &FetchFolderResponse{Folders: folders}
    return response, nil
}

// this function retrieves all Folder instances that match an organization ID
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData() // calls function that return a sample set of folder data

	resFolder := []*Folder{} // initializes a slice to hold points to the Folder structs

	// iterates over folders
	for _, folder := range folders {
		if folder.OrgId == orgID { // checks if the folder's organization ID matches the provided 'orgID'.
			resFolder = append(resFolder, folder) // appends the pointer to the matching Folder to 'resFolder'
		}
	}
	return resFolder, nil //returns the filtered slice and a nil error
}
