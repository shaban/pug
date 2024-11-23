
[0  Tag   "p"]	[1  Text   "This is plain old <em>text</em> content."]	


[0  HTMLTag   "<html>"]	

[0  Tag   "body"]	
  [2  Tag   "p"]	[3  Text   "Indenting the body tag here would make no difference."]	
  [2  Tag   "p"]	[3  Text   "HTML itself isn't whitespace-sensitive."]	

[0  HTMLTag   "</html>"]	


[0  Tag   "p"]	
  [2  Text   "The pipe always goes at the beginning of its own line,"]	
  [2  Text   "not counting indentation."]	


[0  Tag   "script"]	[1  Text   "  if (usingPug)
    console.log('you are awesome')
  else
    console.log('use pug')

"]	
[0  Tag   "div"]	
  [2  Tag   "p"]	[3  Text   "This text belongs to the paragraph tag."]	
  [2  TagVoidInline   "br"]	
  [2  Text   "
    This text belongs to the div tag.

"]	
[0  Text   "You put the em"]	
[0  TagInline   "em"]	[1  Text   "pha"]	
[0  Text   "sis on the wrong syl"]	
[0  TagInline   "em"]	[1  Text   "la"]	
[0  Text   "ble."]	


[0  TagInline   "a"]	[1  Text   "...sentence ending with a link"]	
[0  Text   "."]	


[0  Text   "Don't"]	

[0  Tag   "button"]	[0  ID   "self-destruct"]	[1  Text   "touch"]	

[0  Text   "me!"]	


[0  Tag   "p"]	[1  Text   "  Using regular tags can help keep your lines short,
  but interpolated tags may be easier to "]	[2  TagInline   "em"]	[3  Text   "visualize"]	[2  Text   "
  whether the tags and text are whitespace-separated.

"]	
[0  Text   "Hey, check out "]	
[0  TagInline   "a"]	[0  AttrStart   "("]	[0  Attr   "href"]	[0  AttrEqual   "="]	[0  Attr   ""http://example.biz/kitteh.png""]	[0  AttrEnd   ")"]	[1  Text   "this picture"]	
[0  Text   " of my cat!"]	

[0  Tag   "script"]	[1  Text   "	const newWS = (url) => {
		let socket = new WebSocket(url)
			socket.onopen    =    _    => console.log("Open")
			socket.onmessage = (event) => console.log(`Message: ${event.data}`)
			socket.onerror   = (error) => console.error(`Error: ${error.message}`)
			socket.onclose   = (event) => {
				if (event.wasClean) console.log(`Clean; code: ${event.code} ${event.reason}`)
							else console.log(`Err; code: ${event.code} ${event.reason}`)
			}
		return socket
	}"]	
EOF