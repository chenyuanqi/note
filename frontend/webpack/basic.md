
### Webpack 概况
Webpack 可以看做一个模块化打包机，分析项目结构，处理模块化依赖，转换成为浏览器可运行的代码，把一系列前端代码自动化去构建处理复杂的流程，从而解放生产力。  

- 代码转换: TypeScript 编译成 JavaScript、SCSS,LESS 编译成 CSS。  
- 文件优化：压缩 JavaScript、CSS、HTML 代码，压缩合并图片。  
- 代码分割：提取多个页面的公共代码、提取首屏不需要执行部分的代码让其异步加载。  
- 模块合并：在采用模块化的项目里会有很多个模块和文件，需要构建功能把模块分类合并成一个文件。  
- 自动刷新：监听本地源代码的变化，自动重新构建、刷新浏览器。  

### Webpack 基本使用
webpack4 抽离出了 webpack-cli, 所以我们需要下载 2 个依赖。  
```bash
npm install webpack webpack-cli -D
```

webpack 需要在项目根目录下创建一个 webpack.config.js 来导出 webpack 的配置，配置多样化，可以自行定制，下面是最基础的配置。
```javascript
module.exports = {
	entry: './src/index.js',
	output: {
		path: path.join(__dirname, './dist'),
		filename: 'main.js',
	}
}
```

- entry 代表入口，webpack 会找到该文件进行解析  
- output 代表输入文件配置  
- path 把最终输出的文件放在哪里  
- filename 输出文件的名字  

有时候我们的项目并不是 spa，需要生成多个 js html，那么我们就需要配置多入口。  
```javascript
module.exports = {
	entry: {
		pageA: './src/pageA.js',
		pageB: './src/pageB.js'
	},
	output: {
		path: path.join(__dirname, './dist'),
		filename: '[name].[hash:8].js',
	},
}
```
entry 配置一个对象，key 值就是 chunk： 代码块，一个 Chunk 由多个模块组合而成，用于代码合并与分割。看看 filename[name]: 这个 name 指的就是 chunk 的名字，我们配置的 key 值 pageA pageB，这样打包出来的文件名是不同的，再来看看 [hash]，这个是给输出文件一个 hash 值，避免缓存，那么:8 是取前 8 位。

项目是多页面的，应该有 pageA.html``pageA.js``pageA.css, 那么我应该生成多个 html，这个只是做了 JS 的入口区分，我不想每一个页面都去复制粘贴一个 html, 并且 html 是大部分重复的，可能不同页面只需要修改 title，这时候需要引入一个 webpack 的 plugin。  
```bash
npm install html-webpack-plugin -D
```
该插件可以给每一个 chunk 生成 html, 指定一个 template, 可以接收参数，在模板里面使用。  
```javascript
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
	entry: {
		pageA: './src/pageA.js',
		pageB: './src/pageB.js'
	},
	output: {
		path: path.join(__dirname, './dist'),
		filename: '[name].[hash:8].js',
	},
	plugins: [
		 new HtmlWebpackPlugin({
            template: './src/templet.html',
            filename: 'pageA.html',
            title: 'pageA',
            chunks: ['pageA'],
            hash: true,
            minify: {
                removeAttributeQuotes: true
            }
        }),
        new HtmlWebpackPlugin({
            template: './src/templet.html',
            filename: 'pageB.html',
            title: 'pageB',
            chunks: ['pageB'],
            hash: true,
            minify: {
                removeAttributeQuotes: true
            }
        }),
	]
}
```
在 webpack 中，插件的引入顺序没有规定。  

- template: html 模板的路径地址  
- filename: 生成的文件名  
- title: 传入的参数  
- chunks: 需要引入的 chunk  
- hash: 在引入 JS 里面加入 hash 值 比如: <script src='index.js?2f373be992fc073e2ef5'></script>  
- removeAttributeQuotes: 去掉引号，减少文件大小 <script src=index.js></script>  

这样在 dist 目录下就生成了 pageA.html 和 pageB.html 并且通过配置 chunks，让 pageA.html 里加上了 script 标签去引入 pageA.js。那么现在还剩下 css 没有导入，css 需要借助 loader 去做，所以现在要下载几个依赖，以 scss 为例，less 同理。  
```bash
# css-loader: 支持 css 中的 import，style-loader: 把 css 写入 style 内嵌标签，sass-loader: scss 转换为 css，node-sass: scss 转换依赖
npm install css-loader style-loader sass-loader node-sass -D
```
如何配置 loader 呢？  
```javascript
module.exports = {
	module: {
        rules: [
        		{
        			test: /\.scss$/,
        			use: ['style-loader', 'css-loader', 'sass-loader'],
        			exclude: /node_modules/
        		}
        ]
    }
}
```

- test: 一个正则表达式，匹配文件名  
- use: 一个数组，里面放需要执行的 loader，倒序执行，从右至左  
- exclude: 取消匹配 node_modules 里面的文件  

如果想把 css 作为一个单独的文件，需要用到一个插件来做。  
```bash
npm i extract-text-webpack-plugin@next -D
```

使用 extract-text-webpack-plugin，需要在 plugins 里加入插件 name: chunk 名字 contenthash:8: 根据内容生成 hash 值取前 8 位；
修改 loader 配置下的 use: fallback: 兼容方案。  
```javascript
const ExtractTextPlugin = require('extract-text-webpack-plugin');
module.exports = {
	entry: './src/index.js',
	output: {
		path: path.join(__dirname, './dist'),
		filename: 'main.js',
    },
    module: {
        rules: [
            {
                test: /\.scss$/,
                use: ExtractTextPlugin.extract({
                    // style-loader 把css直接写入html中style标签
                    fallback: 'style-loader',
                    // css-loader css中import支持
                    // loader执行顺序 从右往左执行
                    use: ['css-loader', 'sass-loader']
                }),
                exclude: /node_modules/
            }
        ]
    },
    plugins: [
        new ExtractTextPlugin('[name].[contenthash:8].css'),
    ]
}
```
