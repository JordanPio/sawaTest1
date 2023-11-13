## Instructions

- the code doesn't run because of declared variables not being used
- first loop seems unnecessary

  considering some sample data results from FetchAllFoldersByOrgId

r := []\*Folder{
&Folder{"1", "Folder1", "Org1", false},
&Folder{"2", "Folder2", "Org1", false},
&Folder{"3", "Folder3", "Org2", true},
}

after the first loop we get the following

f := []Folder{
Folder{"1", "Folder1", "Org1", false},
Folder{"2", "Folder2", "Org1", false},
Folder{"3", "Folder3", "Org2", true},
}

its better to just use the values from result straight

the second loop looks weird as well because the results look like this

fp := []\*Folder{
&Folder{"3", "Folder3", "Org2", true}, // seems incorrect: all pointing to the last folder
&Folder{"3", "Folder3", "Org2", true}, // because &v1 is the same in each iteration
&Folder{"3", "Folder3", "Org2", true},
}

-- creating tests

note: added a library testify to help with assertion methods

-- I've added more that to the sample.json so its easy for me to track the results - I've just done this because its a sample data and not real data

- by printing out the results of my first test on terminal I could validate that the second loop was creating what I think its a bug by changing the data.
  -- After I was also able to validate that the first loop was unnecessary since I could simplify using the returns from FetchAllFoldersByOrgID straight in my response
  --I've refactor the function and renamed variables to make the code cleaner and more concise
- created additional test to check what would happen if we pass an OrgId that does not exist on it

the tests validate that my function is working correct after refactoring

-- started working on the second task pagination

- I've research between the different types such as offset-based, cursor-based, keyset and seek method and considering the Requirements I decided to go with the cursor-based.
  I believe it is a good fit for this assignment because it is efficient for the server as there is no counting of rows and can handles the potential of new data well.
- I've created some tests for the function
- Found a bug on my function due to my test "Paginate Beyond Data Set" and fixed

Description of my function

## Function: GetPaginatedAllFolders

1. Check for 'No More Data'
   First, this function looks if the cursor says "END_OF_DATA". If yes, it returns no folders and the same "END_OF_DATA". This means there are no more folders to give.

2. Get All Folders
   Then, it takes all folders using a function called GetAllFolders. It uses OrgID from our request for this.

3. What If Error Happens
   If there's a problem getting folders, the function stops and says there is an error.

4. Understand the Token Logic
   the function uses parsePaginationToken to understand the cursor. This helps to know where to start showing folders.

5. Find Start Point
   It then finds where to start showing folders. If the cursor has a LastID, it searches for this ID in the folders.

6. Decide End Point
   It calculates where to stop showing folders. It adds the limit to the start point but makes sure it's not more than the total folders.

7. Select Folders to Show
   The function selects a part of the folder list from the start to the end point.

8. Prepare for Next Time
   It prepares for the next request. It makes a new cursor for where to start next time using generatePaginationToken.

9. Send Back the Folders and Token
   Lastly, it sends back the folders and new token in something called `PaginatedFetchFolderResponse.`

## Helper Functions

1. generatePaginationToken
   This makes a new cursor from the last folder's ID. If it's the last folder or there's none, it gives "END_OF_DATA".

2. parsePaginationToken
   This function changes a cursor from base64 into a structure the program can use, including the LastID.
   So, this is how the function works. It's like a process of getting folders in parts, based on where to start and end, and preparing for the next part.
