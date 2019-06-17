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
返回c在s中最后一次出现的索引值，若无则返回-1。
```

###12、func SplitN(s, sep string, n int) []string {}
```
以sep为分隔符将s分隔成n个子串,超出的部分不再切分（n=0返回nil,n<0则不限制）,结果不包含sep。
```

###13、func SplitAfterN(s, sep string, n int) []string{}
```
以sep为分隔符将s分隔成n个子串,超出的部分不再切分（n=0返回nil,n<0则不限制）,结果包含sep。
```

###14、func Split(s, sep string) []string{}
```
以sep为分隔符对s进行切分，结果不包含sep。
```

###15、func SplitAfter(s, sep string) []string{}
```
以sep为分隔符对s进行切分，结果包含sep。
```

###16、func Fields(s string) []string{}
```
以空白字符为分隔符对s进行切分，结果不包含空白字符。
```

###17、func FieldsFunc(s string, f func(rune) bool) []string{}
```
按以一个或多个满足 f(rune) 的字符为分隔符对s进行切分，结果不包含分隔符本身。
```

###18、func Join(a []string, sep string) string{}
```
将切片通过sep进行连接形成一个字符串。
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
将s中满足mapping(rune)的字符替换为mapping(rune)的返回值。若mapping(rune)返回负数，则相应的字符将被删除。
```

###22、func Repeat(s string, count int) string{}
```
将count个字符串s连接成一个新的字符串。
```

###23、func ToUpper(s string) string{}
```
将s中的字符改为其大写格式。
```

###24、func ToLower(s string) string{}
```
将s中的字符改为其小写格式。
```

###25、func ToTitle(s string) string {}
```
将s中的字符改为其Title格式。（大部分字符的Title格式就是其Upper格式，只有少数字符的Title格式是特殊字符）
```

###26、func ToUpperSpecial(c unicode.SpecialCase, s string) string {}
```
将s中的字符改为其大写格式。（优先使用_case中的规则进行转换）
```

###27、func ToLowerSpecial(c unicode.SpecialCase, s string) string {}
```
将s中的字符改为其小写格式。（优先使用_case中的规则进行转换）
```

###28、func ToTitleSpecial(c unicode.SpecialCase, s string) string {}
```
将s中的字符改为其Title格式。（优先使用_case中的规则进行转换）
```

###29、func Title(s string) string{}
```
获取非ASCII字符的Title格式列表。
```

###30、func TrimLeftFunc(s string, f func(rune) bool) string{}
```
将删除s头部连续的满足f(rune)的字符。
```

###31、func TrimRightFunc(s string, f func(rune) bool) string{}
```
将删除s尾部连续的满足f(rune)的字符。
```

###32、func TrimFunc(s string, f func(rune) bool) string{}
```
将删除s首尾连续的满足f(rune)的字符。
```

###33、func IndexFunc(s string, f func(rune) bool) int{}
```
返回s中第一个满足f(rune)的字符的字节位置，若无则返回-1。
```

###34、func LastIndexFunc(s string, f func(rune) bool) int{}
```
返回s中最后一个满足f(rune)的字符的字节位置，若无则返回-1。
```

###35、func Trim(s string, cutset string) string{}
```
将删除s首尾连续的包含在cutset中的字符。
```

###36、func TrimLeft(s string, cutset string) string{}
```
将删除s头部连续的包含在cutset中的字符。
```

###37、func TrimRight(s string, cutset string) string{}
```
将删除s尾部连续的包含在cutset中的字符。
```

###38、func TrimSpace(s string) string{}
```
将删除s首尾连续的的空白字符。
```

###39、func TrimPrefix(s, prefix string) string{}
```
删除s头部的prefix字符串。
```

###40、func TrimSuffix(s, suffix string) string{}
```
删除s尾部的suffix字符串。
```

###41、func Replace(s, old, new string, n int) string{}
```
将s中的的old替换为new，n为取代的次数，若n为-1，则全部替换，若old为空，则在每个字符之间都插入一个new。
```

###43、func EqualFold(s，t string)bool{}
```
不考虑大小写，判断s、t是否相同，返回类型为布尔值。
```

###44、func Index(s, substr string) int{}
```
返回substr在s中第一次出现的索引值，若无则返回-1。
```