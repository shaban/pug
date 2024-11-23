
[0  Code   "for i := 0; i < 3; i++ {"]	
  [2  Tag   "li"]	[3  Text   "item"]	
[0  Code   "}"]	

[0  Code   "
  var list = []string{"Uno", "Dos", "Tres",
          "Cuatro", "Cinco", "Seis"}"]	
[0  Each   "item in list"]	
  [2  Tag   "li"]	[3  CodeBuffered   "item"]	


[0  Tag   "p"]	
  [2  CodeBuffered   ""This code is <escaped>!""]	


[0  Tag   "p"]	[1  CodeBuffered   ""This code is" + " <escaped>!""]	


[0  Tag   "p"]	
  [2  CodeUnescaped   ""This code is <strong>not</strong> escaped!""]	


[0  Tag   "p"]	[1  CodeUnescaped   ""This code is" + " <strong>not</strong> escaped!""]	


EOF