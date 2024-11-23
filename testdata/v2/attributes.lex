
[0  TagInline   "a"]	[0  AttrStart   "("]	[0  Attr   "href"]	[0  AttrEqual   "="]	[0  Attr   "`google.com`+`google.com`"]	[0  AttrEnd   ")"]	[1  Text   "Google"]	


[0  TagInline   "a"]	[0  AttrStart   "("]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  Attr   "'button'"]	[0  AttrSpace   " "]	[0  Attr   "href"]	[0  AttrEqual   "="]	[0  Attr   "'google.com'"]	[0  AttrEnd   ")"]	[1  Text   "Google"]	


[0  TagInline   "a"]	[0  AttrStart   "("]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  Attr   "'button'"]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   "href"]	[0  AttrEqual   "="]	[0  Attr   "'google.com'"]	[0  AttrEnd   ")"]	[1  Text   "Google"]	


[0  Code   "var authenticated = true"]	
[0  Tag   "body"]	[0  AttrStart   "("]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  Attr   "authenticated"]	[0  AttrSpace   " "]	[0  Attr   "?"]	[0  AttrSpace   " "]	[0  Attr   ""authed""]	[0  AttrSpace   " "]	[0  Attr   ":"]	[0  AttrSpace   " "]	[0  Attr   ""anon""]	[0  AttrEnd   ")"]	


[0  TagVoid   "input"]	[0  AttrStart   "("]	[0  AttrComma   "
"]	[0  AttrSpace   "  "]	[0  Attr   "type"]	[0  AttrEqual   "="]	[0  Attr   "'checkbox'"]	[0  AttrComma   "
"]	[0  AttrSpace   "  "]	[0  Attr   "name"]	[0  AttrEqual   "="]	[0  Attr   "'agreement'"]	[0  AttrComma   "
"]	[0  AttrSpace   "  "]	[0  Attr   "checked"]	[0  AttrComma   "
"]	[0  AttrEnd   ")"]	


[0  TagVoid   "input"]	[0  AttrStart   "("]	[0  Attr   "data-json"]	[0  AttrEqual   "="]	[0  Attr   "`
  {
    "very-long": "piece of ",
    "data": true
  }
`"]	[0  AttrEnd   ")"]	





[0  Comment   " pug error"]	
[0  Tag   "div"]	[0  AttrStart   "("]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  Attr   "'div-class'"]	[0  AttrSpace   " "]	[0  Attr   "(click)"]	[0  AttrEqual   "="]	[0  Attr   "'play()'"]	[0  AttrEnd   ")"]	


[0  Tag   "div"]	[0  AttrStart   "("]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  Attr   "'div-class'"]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   "(click)"]	[0  AttrEqual   "="]	[0  Attr   "'play()'"]	[0  AttrEnd   ")"]	
[0  Tag   "div"]	[0  AttrStart   "("]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  Attr   "'div-class'"]	[0  AttrSpace   " "]	[0  Attr   "'(click)'"]	[0  AttrEqual   "="]	[0  Attr   "'play()'"]	[0  AttrEnd   ")"]	


[0  TagInline   "a"]	[0  AttrStart   "("]	[0  Attr   "href"]	[0  AttrEqual   "="]	[0  Attr   ""/#{url}""]	[0  AttrEnd   ")"]	[1  Text   "Link"]	

   
   [3  Code   "var url = "pug-test.html""]	
   [3  TagInline   "a"]	[3  AttrStart   "("]	[3  Attr   "href"]	[3  AttrEqual   "="]	[3  Attr   ""/""]	[3  AttrSpace   " "]	[3  Attr   "+"]	[3  AttrSpace   " "]	[3  Attr   "url"]	[3  AttrEnd   ")"]	[4  Text   "Link"]	
   
   
   [3  Code   "url = "https://example.com/""]	
   [3  TagInline   "a"]	[3  AttrStart   "("]	[3  Attr   "href"]	[3  AttrEqual   "="]	[3  Attr   "url"]	[3  AttrEnd   ")"]	[4  Text   "Another link"]	
   
   
   [3  Code   "var btnType = "info""]	
   [3  Code   "var btnSize = "lg""]	
   [3  Tag   "button"]	[3  AttrStart   "("]	[3  Attr   "type"]	[3  AttrEqual   "="]	[3  Attr   "'button'"]	[3  AttrSpace   " "]	[3  Attr   "class"]	[3  AttrEqual   "="]	[3  Attr   ""btn btn-""]	[3  AttrSpace   " "]	[3  Attr   "+"]	[3  AttrSpace   " "]	[3  Attr   "btnType"]	[3  AttrSpace   " "]	[3  Attr   "+"]	[3  AttrSpace   " "]	[3  Attr   "" btn-""]	[3  AttrSpace   " "]	[3  Attr   "+"]	[3  AttrSpace   " "]	[3  Attr   "btnSize"]	[3  AttrEnd   ")"]	
   
   
   [3  Tag   "button"]	[3  AttrStart   "("]	[3  Attr   "type"]	[3  AttrEqual   "="]	[3  Attr   "'button'"]	[3  AttrSpace   " "]	[3  Attr   "class"]	[3  AttrEqual   "="]	[3  Attr   "`btn btn-${btnType} btn-${btnSize}`"]	[3  AttrEnd   ")"]	
   

[0  Tag   "div"]	[0  AttrStart   "("]	[0  Attr   "escaped"]	[0  AttrEqual   "="]	[0  Attr   ""<code>""]	[0  AttrEnd   ")"]	
[0  Tag   "div"]	[0  AttrStart   "("]	[0  Attr   "unescaped"]	[0  AttrEqualUn   "!="]	[0  Attr   ""<code>""]	[0  AttrEnd   ")"]	


[0  TagVoid   "input"]	[0  AttrStart   "("]	[0  Attr   "type"]	[0  AttrEqual   "="]	[0  Attr   "'checkbox'"]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   "checked"]	[0  AttrEnd   ")"]	


[0  TagVoid   "input"]	[0  AttrStart   "("]	[0  Attr   "type"]	[0  AttrEqual   "="]	[0  Attr   "'checkbox'"]	[0  AttrSpace   " "]	[0  Attr   "checked"]	[0  AttrEqual   "="]	[0  Attr   "true"]	[0  AttrEnd   ")"]	


[0  TagVoid   "input"]	[0  AttrStart   "("]	[0  Attr   "type"]	[0  AttrEqual   "="]	[0  Attr   "'checkbox'"]	[0  AttrSpace   " "]	[0  Attr   "checked"]	[0  AttrEqual   "="]	[0  Attr   "false"]	[0  AttrEnd   ")"]	


[0  TagVoid   "input"]	[0  AttrStart   "("]	[0  Attr   "type"]	[0  AttrEqual   "="]	[0  Attr   "'checkbox'"]	[0  AttrSpace   " "]	[0  Attr   "checked"]	[0  AttrEqual   "="]	[0  Attr   ""true""]	[0  AttrEnd   ")"]	


[0  Doctype   "html"]	


[0  TagVoid   "input"]	[0  AttrStart   "("]	[0  Attr   "type"]	[0  AttrEqual   "="]	[0  Attr   "'checkbox'"]	[0  AttrComma   ","]	[0  AttrSpace   " "]	[0  Attr   "checked"]	[0  AttrEnd   ")"]	


[0  TagVoid   "input"]	[0  AttrStart   "("]	[0  Attr   "type"]	[0  AttrEqual   "="]	[0  Attr   "'checkbox'"]	[0  AttrSpace   " "]	[0  Attr   "checked"]	[0  AttrEqual   "="]	[0  Attr   "true"]	[0  AttrEnd   ")"]	


[0  TagVoid   "input"]	[0  AttrStart   "("]	[0  Attr   "type"]	[0  AttrEqual   "="]	[0  Attr   "'checkbox'"]	[0  AttrSpace   " "]	[0  Attr   "checked"]	[0  AttrEqual   "="]	[0  Attr   "false"]	[0  AttrEnd   ")"]	


[0  TagVoid   "input"]	[0  AttrStart   "("]	[0  Attr   "type"]	[0  AttrEqual   "="]	[0  Attr   "'checkbox'"]	[0  AttrSpace   " "]	[0  Attr   "checked"]	[0  AttrEqual   "="]	[0  Attr   "true"]	[0  AttrSpace   " "]	[0  Attr   "&&"]	[0  AttrSpace   " "]	[0  Attr   ""checked""]	[0  AttrSpace   " "]	[0  Attr   "=="]	[0  AttrSpace   " "]	[0  Attr   ""checked""]	[0  AttrEnd   ")"]	


[0  TagInline   "a"]	[0  AttrStart   "("]	[0  Attr   "style"]	[0  AttrEqual   "="]	[0  Attr   "map[string]string{"color":"]	[0  AttrSpace   " "]	[0  Attr   ""red","]	[0  AttrSpace   " "]	[0  Attr   ""background":"]	[0  AttrSpace   " "]	[0  Attr   ""green"}"]	[0  AttrEnd   ")"]	


[0  Code   "var classes = []string{"foo", "bar", "baz"}"]	
[0  TagInline   "a"]	[0  AttrStart   "("]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  Attr   "classes"]	[0  AttrEnd   ")"]	



[0  TagInline   "a"]	[0  Class   "bang"]	[0  AttrStart   "("]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  Attr   ""classes""]	[0  AttrSpace   " "]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  Attr   ""['bing']""]	[0  AttrEnd   ")"]	


[0  Code   "var currentUrl = "/about""]	
[0  TagInline   "a"]	[0  AttrStart   "("]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  AttrSpace   " "]	[0  Attr   "currentUrl"]	[0  AttrSpace   " "]	[0  Attr   "=="]	[0  AttrSpace   " "]	[0  Attr   ""/""]	[0  AttrSpace   " "]	[0  Attr   "?"]	[0  AttrSpace   " "]	[0  Attr   ""active""]	[0  AttrSpace   " "]	[0  Attr   ":"]	[0  AttrSpace   " "]	[0  Attr   """"]	[0  AttrSpace   " "]	[0  Attr   "href"]	[0  AttrEqual   "="]	[0  Attr   "'/'"]	[0  AttrEnd   ")"]	[1  Text   "Home"]	


[0  TagInline   "a"]	[0  AttrStart   "("]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  AttrSpace   " "]	[0  Attr   "currentUrl"]	[0  AttrSpace   " "]	[0  Attr   "=="]	[0  AttrSpace   " "]	[0  Attr   ""/about""]	[0  AttrSpace   " "]	[0  Attr   "?"]	[0  AttrSpace   " "]	[0  Attr   ""active""]	[0  AttrSpace   " "]	[0  Attr   ":"]	[0  AttrSpace   " "]	[0  Attr   """"]	[0  AttrSpace   " "]	[0  Attr   "href"]	[0  AttrEqual   "="]	[0  Attr   "'/about'"]	[0  AttrEnd   ")"]	[1  Text   "About"]	


[0  TagInline   "a"]	[0  Class   "button"]	


[0  Div   "."]	[0  Class   "content"]	


[0  TagInline   "a"]	[0  ID   "main-link"]	


[0  Div   "#"]	[0  ID   "content"]	


[0  Tag   "div"]	[0  ID   "foo"]	[0  AttrStart   "("]	[0  Attr   "data-bar"]	[0  AttrEqual   "="]	[0  Attr   ""foo""]	[0  AttrEnd   ")"]	


[0  Code   "var attributes = struct{class string}{};"]	
[0  Code   "attributes.class = "baz";"]	
[0  Tag   "div"]	[0  ID   "foo"]	[0  AttrStart   "("]	[0  Attr   "data-bar"]	[0  AttrEqual   "="]	[0  Attr   ""foo""]	[0  AttrEnd   ")"]	

[0  Tag   "zxc"]	[0  Class   "asd"]	[0  AttrStart   "("]	[0  Attr   "num"]	[0  AttrEqual   "="]	[0  Attr   "1"]	[0  AttrSpace   " "]	[0  Attr   "class"]	[0  AttrEqual   "="]	[0  Attr   ""qwe""]	[0  AttrEnd   ")"]	[0  Class   "zxc"]	
[0  Tag   "zxc"]	[0  AttrStart   "("]	[0  Attr   "num"]	[0  AttrEqual   "="]	[0  Attr   "1.1"]	[0  AttrEnd   ")"]	
EOF