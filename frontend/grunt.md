
### 什么是 Grunt
基于 Node.js 创建，Grunt 是一个基于任务的命令行工具，它能够同构减少预先需要准备的资源来加速工作流。它将工作包裹进入任务之中，任务会随着你的工作进程自动编译。基本来说，你可以在任何你觉得可以使用 grunt 的项目以及那些需要你手动配置并运行的项目中使用 Grunt。  

Grunt 完成什么样的任务？
Grunt 几乎无所不包，从压缩 JavaScript 到连接 JavaScript；它也可以被用于那些和 JavaScript 不相关的任务之上，例如从 LESS 以及 Sass 编译 CSS。它和 blink 一起使用，可以提醒我们创建过程失败了。

### 为什么使用 Grunt
使用 Grunt 的最大好处在于它带给团队的一致性。  
如果你曾经多人合作完成工作，你就会知道代码中的不一致性是多么令人抓狂。Grunt 使得团队能够使用一组标准化的命令来工作，因此确保了团队中的每个成员都以相同的标准编写代码。毕竟，没有什么比由于代码不一致导致运行失败更加让人痛苦的事情了。  

Grunt 同时也拥有人数日益增多的开发者社区，以及越来越多的新插件。  
学习使用 Grunt 的门槛非常低，因为很多的工具以及自动化任务都已经可以被使用了。

### 设置 Grunt
使用 Grunt 的第一件事是设置 Node.js。  
```bash
# Node.js 安装完成之后，安装 grunt
npm install -g grunt-cli
# 查看 grunt 版本
grunt --version
```

在项目的根目录下创建一个 package.json 以及一个 gruntfile.js 文件。
```
# packge.json
{
    "name": "SampleGrunt",
    "version": "0.1.0",
    "author": "Brandon Random",
    "private": true,
    "devDependencies": {
        "grunt": "~0.4.0",
        "grunt-contrib-cssmin": "*",
        "grunt-contrib-sass": "*",
        "grunt-contrib-uglify": "*",
        "grunt-contrib-watch": "*",
        "grunt-cssc": "*",
        "grunt-htmlhint": "*",
        "matchdep": "*"
    }
}
```

```bash
# 根据 package.json 安装项目依赖，放入 node_modules 文件夹中
npm install
```

```js
// gruntfile.js
// 如果没有 matchdep，我们必须对每一个依赖包写上一段 grunt.loadNpmTasks("grunt-task-name")
require("matched").filterDev("grunt-*").forEach(grunt.loadNpmTasks);

module.exports = function(grunt){

    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json')
    });

    grunt.registerTask('default', []);
};   
```

### 使用 HTMLHint 进行验证
配置 grunt.initConfig，增加如下配置
```
htmlhint: {
    build: {
        options: {
            'tag-pair': true,
            'tagname-lowercase': true,
            'attr-lowercase': true,
            'attr-value-double-quotes': true,
            'doctype-first': true,
            'spec-char-escape': true,
            'id-unique': true,
            'head-script-disabled': true,
            'style-disabled': true
        },
        src: ['index.html']
    }
},
// 监听，当一个文件被保存时自动运行任务	
watch: {
    html: {
        files: ['index.html'],
        tasks: ['htmlhint']
    }
} 
```

```bash
grunt watch
```

### 保持 JavaScript 尽可能精简
创建一个验证用户名的 JavaScript 文件，放到 assets/js/base.js
```js
function Validator()
{
    "use strict";
}

Validator.prototype.checkName = function(name)
{
    "use strict";
    return (/[^a-z]/i.test(name) === false);
};

window.addEventListener('load', function(){
    "use strict";
    document.getElementById('firstname').addEventListener('blur', function(){
        var _this = this;
        var validator = new Validator();
        var validation = document.getElementById('namevalidation');
        if (validator.checkName(_this.value) === true) {
            validation.innerHTML = 'Looks good! :)';
            validation.className = "validation yep";
            _this.className = "yep";
        }
        else {
            validation.innerHTML = 'Looks bad! :(';
            validation.className = "validation nope";
            _this.className = "nope";
        }

    });
}); 
```

使用 UglifyJS 简化源文件，grunt.initConfig 增加如下配置
```
uglify: {
    build: {
        files: {
            'build/js/base.min.js': ['assets/js/base.js']
        }
    }
},
watch: {    
    js: {
        files: ['assets/js/base.js'],
        tasks: ['uglify']
    }
}
```

### 从 Sass 源文件创建 CSS
在创建 CSS 文件时使用 Sass 非常有效，尤其是在团队工作中。源文件中可以包含更少的代码因为 Sass 可以用函数和变量生成大的 CSS 代码块。
```scss
// master.scss
@mixin prefix($property, $value, $prefixes: webkit moz ms o spec) {
    @each $p in $prefixes {
        @if $p == spec {
            #{$property}: $value;
        }
        @else {
            -#{$p}-#{$property}: $value;
        }
    }
}
$input_field:            #999;
$input_focus:           #559ab9;
$validation_passed:     #8aba56;
$validation_failed:     #ba5656;
$bg_colour:             #f4f4f4;
$box_colour:            #fff;
$border_style:          1px solid;
$border_radius:         4px;

html {
    background:         $bg_colour;
}

body {
    width:              720px;
    padding:            40px;
    margin:             80px auto;
    background:         $box_colour;
    box-shadow:         0 1px 3px rgba(0, 0, 0, .1);
    border-radius:      $border_radius;
    font-family:        sans-serif;
}

input[type="text"] {
    @include            prefix(appearance, none, webkit moz);
    @include            prefix(transition, border .3s ease);
    border-radius:      $border_radius;
    border:             $border_style $input_field;
    width:              220px;
}

input[type="text"]:focus {
    border-color:       $input_focus;
    outline:            0;
}

label,
input[type="text"],
.validation {
    line-height:        1;
    font-size:          1em;
    padding:            10px;
    display:            inline;
    margin-right:       20px;
}

input.yep {
    border-color:       $validation_passed;
}

input.nope {
    border-color:       $validation_failed;
}

p.yep {
    color:              $validation_passed;            
}

p.nope {
    color:              $validation_failed;
}
```

添加任务到 gruntfile.js
```js
cssc: {
    build: {
        options: {
            consolidateViaDeclarations: true,
            consolidateViaSelectors:    true,
            consolidateMediaQueries:    true
        },
        files: {
            'build/css/master.css': 'build/css/master.css'
        }
    }
},

cssmin: {
    build: {
        src: 'build/css/master.css',
        dest: 'build/css/master.css'
    }
},

sass: {
    build: {
        files: {
            'build/css/master.css': 'assets/sass/master.scss'
        }
    }
}  
```