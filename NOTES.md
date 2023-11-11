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

its better to just use R straight

the second loop looks weird as well because the results look like this

fp := []\*Folder{
&Folder{"3", "Folder3", "Org2", true}, // Incorrect: all pointing to the last folder
&Folder{"3", "Folder3", "Org2", true}, // because &v1 is the same in each iteration
&Folder{"3", "Folder3", "Org2", true},
}

-- creating tests

note: added a library testify to help with assertion methods

- by printing out the results of my first test on terminal I could validate that the second loop was creating a bug by changing the data.
  -- After I was also able to validate that the first loop was unnecessary since I could simplify using the returns from FetchAllFoldersByOrgID straight in my response
  --I've refactor the function and rename variables to make the code cleaner and more concise
- created additional test to check what would happen if we pass an OrgId that does not exist and its all --

the tests validate that my function is working correct after refactoring
