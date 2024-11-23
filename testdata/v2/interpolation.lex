
[0  Code   "var title = "On Dogs: Man's Best Friend";"]	
[0  Code   "var author = "enlore";"]	
[0  Code   "var theGreat = "<span>escape!</span>";"]	

[0  Tag   "h1"]	[1  CodeBuffered   "title"]	
[0  Tag   "p"]	[1  Text   "Written with love by "]	[1  CodeBuffered   "author"]	
[0  Tag   "p"]	[1  Text   "This will be safe: "]	[1  CodeBuffered   "theGreat"]	


[0  Code   "var msg = "not my inside voice";"]	
[0  Tag   "p"]	[1  Text   "This is "]	[1  CodeBuffered   "strings.ToUpper(msg)"]	


[0  Tag   "p"]	[1  Text   "No escaping for "]	[1  CodeBuffered   "`}`"]	[1  Text   "!"]	


[0  Tag   "p"]	[1  Text   "Escaping works with \#{interpolation}"]	
[0  Tag   "p"]	[1  Text   "Interpolation works with "]	[1  CodeBuffered   "'#{interpolation}'"]	[1  Text   " too!"]	


[0  Code   "var riskyBusiness = "<em>Some of the girls are wearing my mother's clothing.</em>";"]	
[0  Div   "."]	[0  Class   "quote"]	
  [2  Tag   "p"]	[3  Text   "Joel: "]	[3  CodeUnescaped   "riskyBusiness"]	


[0  Tag   "p"]	[1  Text   "  This is a very long and boring paragraph that spans multiple lines.
  Suddenly there is a "]	[2  TagInline   "strong"]	[3  Text   "strongly worded phrase"]	[2  Text   " that cannot be
  "]	[3  TagInline   "em"]	[4  Text   "ignored"]	[3  Text   "."]	
[0  Tag   "p"]	[1  Text   "  And here's an example of an interpolated tag with an attribute:
  "]	[2  Tag   "q"]	[2  AttrStart   "("]	[2  Attr   "lang"]	[2  AttrEqual   "="]	[2  Attr   ""es""]	[2  AttrEnd   ")"]	[3  Text   "¡Hola Mundo!"]	[2  Text   "

"]	
[0  Tag   "p"]	
  [2  Text   "If I don't write the paragraph with tag interpolation, tags like"]	
  [2  TagInline   "strong"]	[3  Text   "strong"]	
  [2  Text   "and"]	
  [2  TagInline   "em"]	[3  Text   "em"]	
  [2  Text   "might produce unexpected results."]	
[0  Tag   "p"]	[1  Text   "  If I do, whitespace is "]	[2  TagInline   "strong"]	[3  Text   "respected"]	[2  Text   " and "]	[3  TagInline   "em"]	[4  Text   "everybody"]	[3  Text   " is happy.
"]	

EOF