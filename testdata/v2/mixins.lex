

[0  Mixin   "list"]	
  [2  Tag   "ul"]	
    [4  Tag   "li"]	[5  Text   "foo"]	
    [4  Tag   "li"]	[5  Text   "bar"]	
    [4  Tag   "li"]	[5  Text   "baz"]	

[0  MixinCall   "list"]	
[0  MixinCall   "list"]	



[0  Mixin   "pet"]	[0  AttrStart   "("]	[0  Attr   "name"]	[0  AttrEnd   ")"]	
  [2  Tag   "li"]	[2  Class   "pet"]	[3  CodeBuffered   "name"]	

[0  Tag   "ul"]	
  [2  MixinCall   "pet"]	[2  AttrStart   "("]	[2  Attr   "'cat'"]	[2  AttrEnd   ")"]	
  [2  MixinCall   "pet"]	[2  AttrStart   "("]	[2  Attr   "'dog'"]	[2  AttrEnd   ")"]	
  [2  MixinCall   "pet"]	[2  AttrStart   "("]	[2  Attr   "'pig'"]	[2  AttrEnd   ")"]	



[0  Mixin   "article"]	[0  AttrStart   "("]	[0  Attr   "title"]	[0  AttrEnd   ")"]	
  [2  Div   "."]	[2  Class   "article"]	
    [4  Div   "."]	[4  Class   "article-wrapper"]	
      [6  Tag   "h1"]	[7  CodeBuffered   "title"]	
      [6  If   "len(block) > 0"]	
        [8  MixinBlock   ""]	
      [6  Else   "else"]	
        [8  Tag   "p"]	[9  Text   "No content provided"]	

[0  MixinCall   "article"]	[0  AttrStart   "("]	[0  Attr   ""Hello world""]	[0  AttrEnd   ")"]	
[0  MixinCall   "article"]	[0  AttrStart   "("]	[0  Attr   ""Hello world""]	[0  AttrEnd   ")"]	
  [2  Tag   "p"]	[3  Text   "This is my"]	
  [2  Tag   "p"]	[3  Text   "Amazing article"]	



[0  Mixin   "link"]	[0  AttrStart   "("]	[0  Attr   "href"]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   "name"]	[0  AttrEnd   ")"]	
  [2  Code   "attributes := struct{class string}{class: "btn"}"]	
  [2  TagInline   "a"]	[2  AttrStart   "("]	[2  Attr   "class"]	[2  AttrEqualUn   "!="]	[2  Attr   "attributes.class"]	[2  AttrSpace   " "]	[2  Attr   "href"]	[2  AttrEqual   "="]	[2  Attr   "href"]	[2  AttrEnd   ")"]	[3  CodeBuffered   "name"]	

[0  MixinCall   "link"]	[0  AttrStart   "("]	[0  Attr   ""/foo""]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   ""foo""]	[0  AttrEnd   ")"]	[0  AttrStart   "("]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  Attr   ""btn""]	[0  AttrEnd   ")"]	
[0  MixinCall   "link"]	[0  AttrStart   "("]	[0  Attr   "fn("/foo", "bar", "baz")"]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   ""foo""]	[0  AttrEnd   ")"]	



[0  Mixin   "link"]	[0  AttrStart   "("]	[0  Attr   "href"]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   "name"]	[0  AttrEnd   ")"]	
  [2  TagInline   "a"]	[2  AttrStart   "("]	[2  Attr   "href"]	[2  AttrEqual   "="]	[2  Attr   "href"]	[2  AttrEnd   ")"]	[3  CodeBuffered   "name"]	

[0  MixinCall   "link"]	[0  AttrStart   "("]	[0  Attr   ""/foo""]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   ""foo""]	[0  AttrEnd   ")"]	[0  AttrStart   "("]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  Attr   ""btn""]	[0  AttrEnd   ")"]	



[0  Mixin   "article"]	[0  AttrStart   "("]	[0  Attr   "title"]	[0  AttrEqual   "="]	[0  Attr   ""Default Title""]	[0  AttrEnd   ")"]	
  [2  Div   "."]	[2  Class   "article"]	
    [4  Div   "."]	[4  Class   "article-wrapper"]	
      [6  Tag   "h1"]	[7  CodeBuffered   "title"]	

[0  MixinCall   "article"]	[0  AttrStart   "("]	[0  AttrEnd   ")"]	
[0  MixinCall   "article"]	[0  AttrStart   "("]	[0  Attr   ""Hello world""]	[0  AttrEnd   ")"]	



[0  Mixin   "list"]	[0  AttrStart   "("]	[0  Attr   "id"]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   "...items"]	[0  AttrEnd   ")"]	
  [2  Tag   "ul"]	[2  AttrStart   "("]	[2  Attr   "id"]	[2  AttrEqual   "="]	[2  Attr   "id"]	[2  AttrEnd   ")"]	
    [4  Each   "item in items"]	
      [6  Tag   "li"]	[7  CodeBuffered   "item"]	

[0  Comment   " TODO for string"]	
[0  MixinCall   "list"]	[0  AttrStart   "("]	[0  Attr   "fn("my-list")"]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   ""string""]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   "2"]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   "3.5"]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   "4"]	[0  AttrEnd   ")"]	

EOF