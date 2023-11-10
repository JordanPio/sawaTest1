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
