
### DOM 事件
DOM 事件数量非常多，即使分类也有十多种，比如键盘事件、鼠标事件、表单事件等，而且不同事件对象属性也有差异。  

*防抖*  
一个场景，有一个搜索输入框，为了提升用户体验，希望在用户输入后可以立即展现搜索结果，而不是每次输入完后还要点击搜索按钮。最基本的实现方式应该很容易想到，那就是绑定 input 元素的键盘事件，然后在监听函数中发送 AJAX 请求。  
```js
const ipt = document.querySelector('input')
ipt.addEventListener('input', e => {
  search(e.target.value).then(resp => {
    // ...
  }, e => {
    // ...
  })
})
```
但其实这样的写法很容易造成性能问题。比如当用户在搜索 “htmlt” 这个词的时候，每一次输入字符都会触发搜索，而实际上，只有最后一次搜索结果是用户想要的，前面进行了 3 次无效查询，浪费了网络带宽和服务器资源。所以对于这类连续触发的事件，需要添加一个“防抖” 功能，为函数的执行设置一个合理的时间间隔，避免事件在时间间隔内频繁触发，同时又保证用户输入后能即时看到搜索结果。  

具体代码如下所示，首先将原函数作为参数传入 debounce () 函数中，同时指定延迟等待时间，返回一个新的函数，这个函数包含 cancel 属性，用来取消原函数执行。flush 属性用来立即调用原函数，同时将原函数的执行结果以 Promise 的形式返回。  
```js
const debounce = (func, wait = 0) => {
  let timeout = null
  let args

  function debounced(...arg) {
    args = arg
    if(timeout) {
      clearTimeout(timeout)
      timeout = null
    }

    // 以Promise的形式返回函数执行结果
    return new Promise((res, rej) => {
      timeout = setTimeout(async () => {
        try {
          const result = await func.apply(this, args)
          res(result)
        } catch(e) {
          rej(e)
        }
      }, wait)
    })
  }

  // 允许取消
  function cancel() {
    clearTimeout(timeout)
    timeout = null
  }

  // 允许立即执行
  function flush() {
    cancel()
    return func.apply(this, args)
  }

  debounced.cancel = cancel
  debounced.flush = flush
  return debounced
}
```

*节流*  
另一个场景，一个左右两列布局的查看文章页面，左侧为文章大纲结构，右侧为文章内容。现在需要添加一个功能，就是当用户滚动阅读右侧文章内容时，左侧大纲相对应部分高亮显示，提示用户当前阅读位置。  
这个功能的实现思路比较简单，滚动前先记录大纲中各个章节的垂直距离，然后监听 scroll 事件的滚动距离，根据距离的比较来判断需要高亮的章节。  
这样功能是实现了，但这并不是最优方法，因为滚动事件的触发频率是很高的，持续调用判断函数很可能会影响渲染性能。实际上也不需要过于频繁地调用，因为当鼠标滚动 1 像素的时候，很有可能当前章节的阅读并没有发生变化。所以我们可以设置在指定一段时间内只调用一次函数，从而降低函数调用频率，这种方式我们称之为 “节流”。  
实现节流函数的过程和防抖函数有些类似，只是对于节流函数而言，有两种执行方式，在调用函数时执行最先一次调用还是最近一次调用，所以需要设置时间戳加以判断。我们可以基于 debounce () 函数加以修改，代码如下所示：  
```js
const throttle = (func, wait = 0, execFirstCall) => {
  let timeout = null
  let args
  let firstCallTimestamp

  function throttled(...arg) {
    if (!firstCallTimestamp) firstCallTimestamp = new Date().getTime()
    if (!execFirstCall || !args) {
      console.log('set args:', arg)
      args = arg
    }

    if (timeout) {
      clearTimeout(timeout)
      timeout = null
    }

    // 以Promise的形式返回函数执行结果
    return new Promise(async(res, rej) => {
      if (new Date().getTime() - firstCallTimestamp >= wait) {
        try {
          const result = await func.apply(this, args)
          res(result)
        } catch (e) {
          rej(e)
        } finally {
          cancel()
        }
      } else {
        timeout = setTimeout(async () => {
          try {
            const result = await func.apply(this, args)
            res(result)
          } catch (e) {
            rej(e)
          } finally {
            cancel()
          }
        }, firstCallTimestamp + wait - new Date().getTime())
      }
    })
  }

  // 允许取消
  function cancel() {
    clearTimeout(timeout)
    args = null
    timeout = null
    firstCallTimestamp = null
  }

  // 允许立即执行

  function flush() {
    cancel()
    return func.apply(this, args)
  }

  throttled.cancel = cancel
  throttled.flush = flush
  return throttled
}
```
节流与防抖都是通过延迟执行，减少调用次数，来优化频繁调用函数时的性能。不同的是，对于一段时间内的频繁调用，防抖是延迟执行后一次调用，节流是延迟定时多次调用。  


*代理*  
下面的 HTML 代码是一个简单的无序列表，现在希望点击每个项目的时候调用 getInfo () 函数，当点击 “编辑” 时，调用一个 edit () 函数，当点击 “删除” 时，调用一个 del () 函数。  
```html
<ul class="list">
  <li class="item" id="item1">项目1<span class="edit">编辑</span><span class="delete">删除</span></li>
  <li class="item" id="item2">项目2<span class="edit">编辑</span><span class="delete" >删除</span></li>
  <li class="item" id="item3">项目3<span class="edit">编辑</span><span class="delete">删除</span></li>
  ...
</ul>
```
要实现这个功能并不难，只需要对列表中每一项，分别监听 3 个元素的 click 事件即可。但如果数据量一旦增大，事件绑定占用的内存以及执行时间将会成线性增加，而其实这些事件监听函数逻辑一致，只是参数不同而已。此时我们可以以事件代理或事件委托来进行优化。  
我们知道，DOM 事件的触发流程主要分为 3 个阶段：  
捕获，事件对象 Window 传播到目标的父对象；  
目标，事件对象到达事件对象的事件目标；  
冒泡，事件对象从目标的父节点开始传播到 Window。  
事件代理的实现原理就是利用上述 DOM 事件的触发流程来对一类事件进行统一处理。比如对于上面的列表，我们在 ul 元素上绑定事件统一处理，通过得到的事件对象来获取参数，调用对应的函数。  
```js
const ul = document.querySelector('.list')
ul.addEventListener('click', e => {
  const t = e.target || e.srcElement
  if (t.classList.contains('item')) {
    getInfo(t.id)
  } else {
    id = t.parentElement.id
    if (t.classList.contains('edit')) {
      edit(id)
    } else if (t.classList.contains('delete')) {
      del(id)
    }
  }
})
```
虽然这里我们选择了默认在冒泡阶段监听事件，但和捕获阶段监听并没有区别。对于其他情况还需要具体情况具体细分析，比如有些列表项目需要在目标阶段进行一些预处理操作，那么可以选择冒泡阶段进行事件代理。  

说说下面 3 种事件监听方式的区别？  
```js
// 方式1
<input type="text" onclick="click()"/>

// 方式2
document.querySelector('input').onClick = function(e) {
  // ...
}

// 方式3
document.querySelector('input').addEventListener('click', function(e) {
  //...
})
```
方式 1 和方式 2 同属于 DOM0 标准，通过这种方式进行事件监会覆盖之前的事件监听函数。  
方式 3 属于 DOM2 标准，推荐使用这种方式。同一元素上的事件监听函数互不影响，而且可以独立取消，调用顺序和监听顺序一致。  
