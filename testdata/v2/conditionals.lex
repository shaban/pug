
[0  Code   " 
  var user = struct{
                  description, name string
                  isAnonymous bool
                    }{ description: "foo bar baz", name: "zxc" }"]	
[0  Code   "var authorised = false"]	
[0  Div   "#"]	[0  ID   "user"]	
  [2  If   "len(user.description) > 0"]	
    [4  Tag   "h2"]	[4  Class   "green"]	[5  Text   "Description"]	
    [4  Tag   "p"]	[4  Class   "description"]	[5  CodeBuffered   "user.description"]	
  [2  Else   "else"]	[2  If   "authorised"]	
    [4  Tag   "h2"]	[4  Class   "blue"]	[5  Text   "Description"]	
    [4  Tag   "p"]	[4  Class   "description"]	[5  Text   "      User has no description,
      why not add one..."]	
  [2  Else   "else"]	
    [4  Tag   "h2"]	[4  Class   "red"]	[5  Text   "Description"]	
    [4  Tag   "p"]	[4  Class   "description"]	[5  Text   "User has no description"]	



[0  Unless   "user.isAnonymous"]	
  [2  Tag   "p"]	[3  Text   "You're logged in as "]	[3  CodeBuffered   "user.name"]	

[0  If   "!user.isAnonymous"]	
  [2  Tag   "p"]	[3  Text   "You're logged in as "]	[3  CodeBuffered   "user.name"]	


EOF