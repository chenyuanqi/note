
### Jquery 常用选择器
```
基本选择器  
$('*')	匹配页面所有元素  
$('#id')	id 选择器  
$('.class')	类选择器  
$('element')	标签选择器  

组合 / 层次选择器  
$('E,F')	多元素选择器，用”, 分隔，同时匹配元素 E 或元素 F  
$('E F')	后代选择器，用空格分隔，匹配 E 元素所有的后代（不只是子元素、子元素向下递归）元素 F  
$(E>F)	子元素选择器，用”>” 分隔，匹配 E 元素的所有直接子元素  
$('E+F')	直接相邻选择器，匹配 E 元素之后的相邻的同级元素 F  
$('E~F')	普通相邻选择器（弟弟选择器），匹配 E 元素之后的同级元素 F（无论直接相邻与否）  
$('.class1.class2')	匹配类名中既包含 class1 又包含 class2 的元素  

基本过滤选择器  
$("E:first")	所有 E 中的第一个  
$("E:last")	所有 E 中的最后一个  
$("E:not(selector)")	按照 selector 过滤 E  
$("E:even")	所有 E 中 index 是偶数  
$("E:odd")	所有 E 中 index 是奇数  
$("E:eq(n)")	所有 E 中 index 为 n 的元素  
$("E:gt(n)")	所有 E 中 index 大于 n 的元素  
$("E:ll(n)")	所有 E 中 index 小于 n 的元素  
$(":header")	选择 h1~h7 元素  
$("div:animated")	选择正在执行动画效果的元素  

内容过滤器  
$('E:contains(value)')	内容中包含 value 值的元素  
$('E:empty')	内容为空的元素  
$('E:has(F)')	子元素中有 F 的元素，$('div:has (a)'): 包含 a 标签的 div  
$('E: parent')	父元素是 E 的元素，$('td: parent'): 父元素是 td 的元素  

可视化选择器  
$('E:hidden')	所有被隐藏的 E  
$('E:visible')	所有可见的 E  

属性过滤选择器  
$('E[attr]')	含有属性 attr 的 E  
$('E[attr=value]')	属性 attr=value 的 E  
$('E[attr !=value]')	属性 attr！=value 的 E  
$('E[attr ^=value]')	属性 attr 以 value 开头的 E  
$('E[attr $=value]')	属性 attr 以 value 结尾的 E  
$('E[attr =value]')	属性 attr 包含 value 的 E  
$('E[attr][attr =value]')	可以连用  

子元素过滤器  
$('E:nth-child(n)')	E 的第 n 个子节点  
$('E:nth-child(3n+1)')	E 的 index 符合 3n+1 表达式的子节点  
$('E:nth-child(even)')	E 的 index 为偶数的子节点  
$('E:nth-child(odd)')	E 的 index 为奇数的子节点  
$('E:first-clild')	所有 E 的第一个子节点  
$('E:last-clild')	所有 E 的最后一个子节点  
$('E:only-clild')	只有唯一子节点的 E 的子节点  

表单元素选择器  
$('E:type')	特定类型的 input  
$(':checked')	被选中的 checkbox 或 radio  
$('option: selected')	被选中的 option  

除了选择器，筛选节点的方法还可以：  
.find (selector) 查找集合每个元素的子节点  
.filter (selector) 过滤当前集合内元素  
```

