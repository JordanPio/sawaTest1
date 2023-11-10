package folders

import (
	"github.com/gofrs/uuid"
)
// Recommendations
// fix issues that prevent the code from running - such as removing unused variables
// have better naming conventions as its quite confusing at the moment
// try to run the code after that
// add proper error handling
// create unit tests for the functions
// the first loop seems to be unnecessary as its just copying the data from a pointer to a value and then back to a pointer again - I believe we can remove this for improving performance
// the second loop also seems unnecessary and it seems to be wrong as it reuses the loop variable address which can leap to all elements in 'fp' pointing to the same instance



// this function retrieves all folders related to an organization ID returning in the response and error
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {



	f := []Folder{} // initializes a slice to hold folder structs - SLICE is similar to array in JS

	// this calls the function FetchAllFoldersByOrgID and its ignoring any errors which is not a good practice
	// errors should be checked and handled properly

	r, _ := FetchAllFoldersByOrgID(req.OrgID)

	// iterate over the results from fetchAllFoldersByOrgID
	for _, v := range r { // THIS LOOP SEEMS TO BE UNNECESSARY

		f = append(f, *v) // dereferences each pointer and append the folder struct to the slice 'f'
	}
	var fp []*Folder // initializes  a slice to hold pointers to folder structs

	// iterates over f containing folder structs
	for _, v1 := range f { // THIS LOOP SEEMS TO BE INCORRECT
		fp = append(fp, &v1) // takes the address of each Folder struct and appends it to the slice 'fp'
	}
	var ffr *FetchFolderResponse // declares a pointer to a FetchFolderResponse.
	ffr = &FetchFolderResponse{Folders: fp} // initializes FetchFolderResponse with the slice 'fp'
	return ffr, nil // returns the response struct and a nil error
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
