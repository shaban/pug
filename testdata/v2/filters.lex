  [2  Tag   "p"]	
    [4  Filter   "markdown-it"]	[4  FilterArgs   "inline"]	[4  FilterText   " **BOLD TEXT**
"]	
  [2  Tag   "p"]	[3  Text   "    In the midst of a large amount of plain
    text, suddenly a wild "]	[4  Filter   "markdown-it"]	[4  FilterArgs   "inline"]	[4  FilterText   " *Markdown*]"]	
    [4  Tag   "appeared"]	[5  Text   "
"]	
  [2  Tag   "script"]	
    [4  Filter   "cdata-js"]	[4  FilterSubf   "babel"]	[4  FilterArgs   "presets=['es2015']"]	[4  FilterText   "
      const myFunc = () => `This is ES2015 in a CD${'ATA'}`;
"]	
[0  Filter   "go"]	[0  FilterSubf   "import"]	[0  FilterText   "
  "github.com/shaban/pug""]	

[0  Tag   "p"]	
  [2  Filter   "my-own-filter"]	[2  FilterArgs   "addStart addEnd"]	[2  FilterText   "
    Filter
    Body"]	
  
  [2  Tag   "div"]	
    [4  Div   "#"]	[4  ID   "function"]	
      [6  Filter   "go"]	[6  FilterSubf   "func"]	[6  FilterText   " Output(user string, buffer *bytes.Buffer)"]	

EOF