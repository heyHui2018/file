##strings包

###1、func Compare(a, b string) int{}
```
比较两个，若a==b则返回0，若a<b则返回-1，若a>b则返回1。
```

###2、func Count(s, substr string) int{}
```
统计s中substr出现的次数，返回类型为int。若substr为空，则返回1+len(s)。
```

###3、func Contains(s, substr string) bool{}
```
判断s中是否包含substr，返回类型为布尔值。
```

###4、func ContainsAny(s, chars string) bool{}
```
判断s中是否包含chars中的任何字符，返回类型为布尔值。
```

###5、func ContainsRune(s string, r rune) bool{}
```
判断s中是否包含r，返回类型为布尔值。
```

###6、func LastIndex(s, substr string) int{}
```
返回substr在s中最后一次出现的索引值，若无则返回-1。
```

###7、func IndexByte(s string, c byte) int{}
```
返回c在s中第一次出现的索引值，若无则返回-1。
```

###8、func IndexRune(s string, r rune) int{}
```
返回r在s中第一次出现的索引值，若无则返回-1。
```

###9、func IndexAny(s, chars string) int{}
```
返回chars的任何一个字符在s中第一次出现的索引值，若无则返回-1。
```

###10、func LastIndexAny(s, chars string) int{}
```
返回chars的任何一个字符在s中最后一次出现的索引值，若无则返回-1。
```

###11、func LastIndexByte(s string, c byte) int{}
```
返回s中c的最后一个索引值，若无则返回-1。
```

###12、func SplitN(s, sep string, n int) []string {}
```

```

###13、func SplitAfterN(s, sep string, n int) []string{}
```
按照sep对s进行切分，结果中包含sep本身，参数n表示最多切分出的子串个数，超出部分将不再切分。
若n=0则返回nil，若n<0则全部切分。
```

###14、func Split(s, sep string) []string{}
```
按照sep对s进行切分，返回值为切片类型。
```

###15、func SplitAfter(s, sep string) []string{}
```
按照sep对s进行切分，结果中包含sep本身。
若sep为空，将s切分成Unicode字符列表。
若s不包含sep，则返回长度为1的切片，其唯一元素为s。
若s和sep都为空，则返回一个空切片。
```

###16、func Fields(s string) []string{}
```
按照空格对s进行切分，返回类型为切片类型。
```

###17、func FieldsFunc(s string, f func(rune) bool) []string{}
```
按照自定义方法f定义的rune对s进行切分，返回类型为切片类型。
```

###18、func Join(a []string, sep string) string{}
```
将切片通过sep进行连接，形成一个。
```

###19、func HasPrefix(s, prefix string) bool{}
```
判断s是否是以prefix开头，返回类型为布尔值。
```

###20、func HasSuffix(s, suffix string) bool{}
```
判断s是否是以suffix结尾，返回类型为布尔值。
```

###21、func Map(mapping func(rune) rune, s string) string{}
```
返回s的副本，其中mapping的key里存在的rune都将被替换为value。若映射返回负值，则该key将从s中删除。
```

###22、func Repeat(s string, count int) string{}
```
返回一个由count个s组成的新。
```

###23、func ToUpper(s string) string{}
```
将s转换为大写
```

###24、func ToLower(s string) string{}
```
返回s的副本，其中所有Unicode字母都映射到它们的小写字母。
```

###25、func ToTitle(s string) string {}
```

```

###26、func ToUpperSpecial(c unicode.SpecialCase, s string) string {}
```

```

###27、func ToLowerSpecial(c unicode.SpecialCase, s string) string {}
```
使用c指定的大小写映射返回s的副本，其中所有Unicode字母都映射到它们的小写字母。
```

###28、func ToTitleSpecial(c unicode.SpecialCase, s string) string {}
```

```

###29、func Title(s string) string{}
```
返回s的副本，其中包含所有Unicode字母，这些字母开始映射到其标题大小写的单词。
```

###30、func TrimLeftFunc(s string, f func(rune) bool) string{}
```
返回s的一个片段，其中所有前导的Unicode代码点c都满足f（c）被删除。
```

###31、func TrimRightFunc(s string, f func(rune) bool) string{}
```
返回s的一个片段，其中所有尾随的Unicode代码点c满足f（c）被删除。
```

###32、func TrimFunc(s string, f func(rune) bool) string{}
```
返回s的一个片段，其中所有前导和尾随的Unicode代码点c都满足f（c）被删除。
```

###33、func IndexFunc(s string, f func(rune) bool) int{}
```
返回s中f定义的rune第一个索引值，若无则返回-1。
```

###34、func LastIndexFunc(s string, f func(rune) bool) int{}
```
返回s中f定义的rune的最后一个索引值，若无则返回-1。
```

###35、func Trim(s string, cutset string) string{}
```
返回s的一个片段，其中包含cutset中包含的所有前导和尾随Unicode代码点。
```

###36、func TrimLeft(s string, cutset string) string{}
```
返回s的一个片段，其中包含cutset中包含的所有前导Unicode代码点。
```

###37、func TrimRight(s string, cutset string) string{}
```
返回s的一个片段，删除了cutset中包含的所有尾随Unicode代码点。
```

###38、func TrimSpace(s string) string{}
```
返回s的一部分，删除所有前导和尾随空格，如Unicode所定义。
```

###39、func TrimPrefix(s, prefix string) string{}
```
返回s而没有提供的前导前缀。若s不以prefix开头，则返回s不变。
```

###40、func TrimSuffix(s, suffix string) string{}
```
返回s而没有提供的尾随后缀。若s不以后缀结尾，则s保持不变。
```

###41、func Replace(s, old, new string, n int) string{}
```
将s的old替换为new，n为取代的次数。
```

###42、func ReplaceAll(s, old, new string) string{}
```
将s中的old全部替换为new。
```

###43、func EqualFold（s，t string）bool{}
```
不考虑大小写，判断s、t是否相同，返回类型为布尔值。
```

###44、func Index(s, substr string) int{}
```
返回s中substr的第一个索引值，若无则返回-1。
```