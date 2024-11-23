
[0  Code   "var friends1 = 10"]	
[0  Case   "friends1"]	
  [2  CaseWhen   "0"]	
    [4  Tag   "p"]	[5  Text   "you have no friends1"]	
  [2  CaseWhen   "1"]	
    [4  Tag   "p"]	[5  Text   "you have a friend"]	
  [2  CaseDefault   "default"]	
    [4  Tag   "p"]	[5  Text   "you have "]	[5  CodeBuffered   "friends1"]	[5  Text   " friends1"]	


[0  Code   "var friends2 = 0"]	
[0  Case   "friends2"]	
  [2  CaseWhen   "0"]	
    [4  Code   "fallthrough"]	
  [2  CaseWhen   "1"]	
    [4  Tag   "p"]	[5  Text   "you have very few friends2"]	
  [2  CaseDefault   "default"]	
    [4  Tag   "p"]	[5  Text   "you have "]	[5  CodeBuffered   "friends2"]	[5  Text   " friends2"]	


[0  Code   "var friends3 = 0"]	
[0  Case   "friends3"]	
  [2  CaseWhen   "0"]	
    [4  Code   "break"]	
  [2  CaseWhen   "1"]	
    [4  Tag   "p"]	[5  Text   "you have very few friends3"]	
  [2  CaseDefault   "default"]	
    [4  Tag   "p"]	[5  Text   "you have "]	[5  CodeBuffered   "friends3"]	[5  Text   " friends3"]	


[0  Code   "var friends = 1"]	
[0  Case   "friends"]	
  [2  CaseWhen   "0"]	[3  Tag   "p"]	[4  Text   "you have no friends"]	
  [2  CaseWhen   "1"]	[3  Tag   "p"]	[4  Text   "you have a friend"]	
  [2  CaseDefault   "default"]	[3  Tag   "p"]	[4  Text   "you have "]	[4  CodeBuffered   "friends"]	[4  Text   " friends"]	


EOF