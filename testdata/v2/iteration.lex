
[0  Tag   "ul"]	
  [2  Each   "val in []int{1, 2, 3, 4, 5}"]	
    [4  Tag   "li"]	[5  CodeBuffered   "val"]	


[0  Tag   "ul"]	
  [2  Each   "val, index in []string{"zero", "one", "two"}"]	
    [4  Tag   "li"]	[5  CodeBuffered   "strconv.Itoa(index) + ": " + val"]	


[0  Tag   "ul"]	
  [2  Each   "val, index in map[int]string{1:"one",2:"two",3:"three"}"]	
    [4  Tag   "li"]	[5  CodeBuffered   "strconv.Itoa(index) + ": " + val"]	


[0  Code   " 
  qfs := func (condition bool, iftrue, iffalse []string) []string {
		if condition {
			return iftrue
		} else {
			return iffalse
		}
	}
  var values = []string{}
"]	
[0  Tag   "ul"]	
  [2  Each   "val in qfs(len(values)>0, values, []string{"There are no values"})"]	
    [4  Tag   "li"]	[5  CodeBuffered   "val"]	


[0  Code   "var values1 = []string{}"]	
[0  Tag   "ul"]	
  [2  Each   "val in values1"]	
    [4  Tag   "li"]	[5  CodeBuffered   "val"]	
  [2  Else   "else"]	
    [4  Tag   "li"]	[5  Text   "There are no values1"]	


[0  Code   "var n = 0;"]	
[0  Tag   "ul"]	
  [2  While   "n < 4"]	
    [4  Tag   "li"]	[5  CodeBuffered   "n ; n++"]	


EOF