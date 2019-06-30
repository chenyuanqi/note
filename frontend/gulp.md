
### gulp 是什么
gulp 是基于 node 实现 Web 前端自动化开发的工具，利用它能够极大的提高开发效率。  
在 Web 前端开发工作中有很多 “重复工作”，比如压缩 CSS/JS 文件。而这些工作都是有规律的。找到这些规律，并编写 gulp 配置代码，让 gulp 自动执行这些 “重复工作”。

### gulp 安装
gulp 是基于 node 实现的，那么我们就需要先安装 node。  
```bash
npm install -g gulp 
```

### gulp 压缩 JS
压缩 js 代码可降低 js 文件大小，提高页面打开速度。在不利用 gulp 时我们需要通过各种工具手动完成压缩工作，所有的 gulp 代码编写都可以看做是将规律转化为代码的过程。  
规律：找到 js/ 目录下的所有 js 文件，压缩它们，将压缩后的文件存放在 dist/js/ 目录下。  

gulp 的所有配置代码都写在 gulpfile.js 文件。  
```js
// 获取 gulp
var gulp = require('gulp')

// 获取 uglify 模块（用于压缩 JS）
// 需安装执行命令：npm install gulp-uglify
var uglify = require('gulp-uglify')

// 压缩 js 文件
// 在命令行使用 gulp script 启动此任务
// gulp.task(name, fn) - 定义任务，第一个参数是任务名，第二个参数是任务内容
gulp.task('script', function() {
    	// 1. 找到文件，选择文件，传入参数是文件路径
    gulp.src('js/*.js')
    	// 2. 压缩文件，通过管道将操作加入执行队列
        .pipe(uglify())
    	// 3. 另存压缩后的文件，输出文件
        .pipe(gulp.dest('dist/js'))
})

// 检测代码修改，自动执行任务
// 在命令行使用 gulp auto 启动此任务
gulp.task('auto', function () {
    // 监听文件修改，当文件被修改则执行 script 任务
    gulp.watch('js/*.js', ['script'])
})

// 使用 gulp.task('default') 定义默认任务，即执行 gulp 命令会执行的任务
// 在命令行使用 gulp 启动 script 任务和 auto 任务
gulp.task('default', ['script', 'auto'])
```
打开命令行使用 cd 命令跳转至 gulpfile.js 文件所在目录。执行 gulp script 命令即可。

### gulp 压缩 CSS
压缩 css 代码可降低 css 文件大小，提高页面打开速度。  
规律：找到 css/ 目录下的所有 css 文件，压缩它们，将压缩后的文件存放在 dist/css/ 目录下。  

在对应目录创建 gulpfile.js 文件并写入如下内容：
```js
// 获取 gulp
var gulp = require('gulp')

// 获取 minify-css 模块（用于压缩 CSS）
// 需安装执行命令：npm install gulp-minify-css
var minifyCSS = require('gulp-minify-css')

// 压缩 css 文件
// 在命令行使用 gulp css 启动此任务
gulp.task('css', function () {
    	// 1. 找到文件
    gulp.src('css/*.css')
    	// 2. 压缩文件
        .pipe(minifyCSS())
    	// 3. 另存为压缩文件
        .pipe(gulp.dest('dist/css'))
})

// 在命令行使用 gulp auto 启动此任务
gulp.task('auto', function () {
    // 监听文件修改，当文件被修改则执行 css 任务
    gulp.watch('css/*.css', ['css'])
});

// 使用 gulp.task('default') 定义默认任务
// 在命令行使用 gulp 启动 css 任务和 auto 任务
gulp.task('default', ['css', 'auto'])
```

### gulp 压缩图片
压缩图片文件可降低文件大小，提高图片加载速度。  
规律：找到 images/ 目录下的所有文件，压缩它们，将压缩后的文件存放在 dist/images/ 目录下。  

在对应目录创建 gulpfile.js 文件并写入如下内容：  
```js
// 获取 gulp
var gulp = require('gulp');

// 获取 gulp-imagemin 模块
// 需安装执行命令：npm install gulp-imagemin
var imagemin = require('gulp-imagemin')

// 压缩图片任务
// 在命令行输入 gulp images 启动此任务
gulp.task('images', function () {
    	// 1. 找到图片
    gulp.src('images/*.*')
    	// 2. 压缩图片
        .pipe(imagemin({
            progressive: true
        }))
    	// 3. 另存图片
        .pipe(gulp.dest('dist/images'))
});

// 在命令行使用 gulp auto 启动此任务
gulp.task('auto', function () {
    // 监听文件修改，当文件被修改则执行 images 任务
    gulp.watch('images/*.*)', ['images'])
});

// 使用 gulp.task('default') 定义默认任务
// 在命令行使用 gulp 启动 images 任务和 auto 任务
gulp.task('default', ['images', 'auto'])
```

### gulp 编译 LESS
Less 是一门 CSS 预处理语言，它扩充了 CSS 语言，增加了诸如变量、混合（mixin）、函数等功能，让 CSS 更易维护。

在对应目录创建 gulpfile.js 文件并写入如下内容： 
```js
// 获取 gulp
var gulp = require('gulp')
// 获取 gulp-less 模块
// 需安装执行命令：npm install gulp-less
var less = require('gulp-less')

// 编译less
// 在命令行输入 gulp less 启动此任务
gulp.task('less', function () {
    	// 1. 找到 less 文件
    gulp.src('less/**.less')
    	// 2. 编译为css
        .pipe(less())
    	// 3. 另存文件
        .pipe(gulp.dest('dist/css'))
});

// 在命令行使用 gulp auto 启动此任务
gulp.task('auto', function () {
    // 监听文件修改，当文件被修改则执行 less 任务
    gulp.watch('less/**.less', ['less'])
})

// 使用 gulp.task('default') 定义默认任务
// 在命令行使用 gulp 启动 less 任务和 auto 任务
gulp.task('default', ['less', 'auto'])
```

### gulp 编译 Sass
Sass 是一种 CSS 的开发工具，提供了许多便利的写法，大大节省了开发者的时间，使得 CSS 的开发，变得简单和可维护。  

`注意：无论是 node-sass 还是 ruby-sass 使用 npm 安装都非常的慢，甚至会有可能安装不上，极其不利于团队协作。  
建议使用 less 作为 css 预处理器。 如果因为 less 不支持自定义函数选择用 sass 可以使用 less-plugin-functions 让 less 支持自定义函数。`  
在对应目录创建 gulpfile.js 文件并写入如下内容： 
```js
// 获取 gulp
var gulp = require('gulp')
// 获取 gulp-ruby-sass 模块
// 需安装执行命令：npm install gulp-ruby-sass	
// 如果不方便安装 ruby 或编译速度慢，建议安装执行命令：npm install gulp-sass
var sass = require('gulp-ruby-sass')

// 编译sass
// 在命令行输入 gulp sass 启动此任务
gulp.task('sass', function() {
    return sass('sass/') 
    .on('error', function (err) {
      console.error('Error!', err.message);
   })
    .pipe(gulp.dest('dist/css'))
});


// 在命令行使用 gulp auto 启动此任务
gulp.task('auto', function () {
    // 监听文件修改，当文件被修改则执行 images 任务
    gulp.watch('sass/**/*.scss', ['sass'])
});

// 使用 gulp.task('default') 定义默认任务
// 在命令行使用 gulp 启动 sass 任务和 auto 任务
gulp.task('default', ['sass', 'auto'])
```

### gulp 构建一个项目
首先，创建 package.json。  
```bash
npm init
```

安装依赖  
```bash
npm install gulp --save-dev
npm install gulp-uglify gulp-watch-path stream-combiner2 gulp-sourcemaps gulp-minify-css gulp-autoprefixer gulp-less gulp-ruby-sass gulp-imagemin gulp-util --save-dev
```

设计目录结构  
└── src/
	├── less/    *.less 文件  
	├── sass/    *.scss *.sass 文件  
	├── css/     *.css  文件  
	├── js/      *.js 文件  
	├── fonts/   字体文件  
    └── images/   图片  
└── dist/   

让命令行输出的文字带颜色  
gulp 自带的输出都带时间和颜色，这样很容易识别。我们也可以利用 gulp-util 实现同样的效果
```js
var gulp = require('gulp')
var gutil = require('gulp-util')

gulp.task('default', function () {
    gutil.log('message')
    gutil.log(gutil.colors.red('error'))
    gutil.log(gutil.colors.green('message:') + "some")
})
```	 	

配置任务  
在对应目录创建 gulpfile.js 文件并写入如下内容： 
```js
var gulp = require('gulp')
var gutil = require('gulp-util')
var uglify = require('gulp-uglify')
var watchPath = require('gulp-watch-path')
var combiner = require('stream-combiner2')
var sourcemaps = require('gulp-sourcemaps')
var minifycss = require('gulp-minify-css')
var autoprefixer = require('gulp-autoprefixer')
var less = require('gulp-less')
var sass = require('gulp-ruby-sass')
var imagemin = require('gulp-imagemin')

var handlebars = require('gulp-handlebars');
var wrap = require('gulp-wrap');
var declare = require('gulp-declare');

var handleError = function (err) {
    var colors = gutil.colors;
    console.log('\n')
    gutil.log(colors.red('Error!'))
    gutil.log('fileName: ' + colors.red(err.fileName))
    gutil.log('lineNumber: ' + colors.red(err.lineNumber))
    gutil.log('message: ' + err.message)
    gutil.log('plugin: ' + colors.yellow(err.plugin))
}

gulp.task('watchjs', function () {
    gulp.watch('src/js/**/*.js', function (event) {
        var paths = watchPath(event, 'src/', 'dist/')
        /*
        paths
            { srcPath: 'src/js/log.js',
              srcDir: 'src/js/',
              distPath: 'dist/js/log.js',
              distDir: 'dist/js/',
              srcFilename: 'log.js',
              distFilename: 'log.js' }
        */
        gutil.log(gutil.colors.green(event.type) + ' ' + paths.srcPath)
        gutil.log('Dist ' + paths.distPath)

        var combined = combiner.obj([
            gulp.src(paths.srcPath),
            sourcemaps.init(),
            uglify(),
            sourcemaps.write('./'),
            gulp.dest(paths.distDir)
        ])

        combined.on('error', handleError)
    })
})

gulp.task('uglifyjs', function () {
    var combined = combiner.obj([
        gulp.src('src/js/**/*.js'),
        sourcemaps.init(),
        uglify(),
        sourcemaps.write('./'),
        gulp.dest('dist/js/')
    ])
    combined.on('error', handleError)
})


gulp.task('watchcss', function () {
    gulp.watch('src/css/**/*.css', function (event) {
        var paths = watchPath(event, 'src/', 'dist/')

        gutil.log(gutil.colors.green(event.type) + ' ' + paths.srcPath)
        gutil.log('Dist ' + paths.distPath)

        gulp.src(paths.srcPath)
            .pipe(sourcemaps.init())
            .pipe(autoprefixer({
              browsers: 'last 2 versions'
            }))
            .pipe(minifycss())
            .pipe(sourcemaps.write('./'))
            .pipe(gulp.dest(paths.distDir))
    })
})

gulp.task('minifycss', function () {
    gulp.src('src/css/**/*.css')
        .pipe(sourcemaps.init())
        .pipe(autoprefixer({
          browsers: 'last 2 versions'
        }))
        .pipe(minifycss())
        .pipe(sourcemaps.write('./'))
        .pipe(gulp.dest('dist/css/'))
})

gulp.task('watchless', function () {
    gulp.watch('src/less/**/*.less', function (event) {
        var paths = watchPath(event, 'src/less/', 'dist/css/')

        gutil.log(gutil.colors.green(event.type) + ' ' + paths.srcPath)
        gutil.log('Dist ' + paths.distPath)
        var combined = combiner.obj([
            gulp.src(paths.srcPath),
            sourcemaps.init(),
            autoprefixer({
              browsers: 'last 2 versions'
            }),
            less(),
            minifycss(),
            sourcemaps.write('./'),
            gulp.dest(paths.distDir)
        ])
        combined.on('error', handleError)
    })
})

gulp.task('lesscss', function () {
    var combined = combiner.obj([
            gulp.src('src/less/**/*.less'),
            sourcemaps.init(),
            autoprefixer({
              browsers: 'last 2 versions'
            }),
            less(),
            minifycss(),
            sourcemaps.write('./'),
            gulp.dest('dist/css/')
        ])
    combined.on('error', handleError)
})


gulp.task('watchsass',function () {
    gulp.watch('src/sass/**/*', function (event) {
        var paths = watchPath(event, 'src/sass/', 'dist/css/')

        gutil.log(gutil.colors.green(event.type) + ' ' + paths.srcPath)
        gutil.log('Dist ' + paths.distPath)
        sass(paths.srcPath)
            .on('error', function (err) {
                console.error('Error!', err.message);
            })
            .pipe(sourcemaps.init())
            .pipe(minifycss())
            .pipe(autoprefixer({
              browsers: 'last 2 versions'
            }))
            .pipe(sourcemaps.write('./'))
            .pipe(gulp.dest(paths.distDir))
    })
})

gulp.task('sasscss', function () {
        sass('src/sass/*')
        .on('error', function (err) {
            console.error('Error!', err.message);
        })
        .pipe(sourcemaps.init())
        .pipe(minifycss())
        .pipe(autoprefixer({
          browsers: 'last 2 versions'
        }))
        .pipe(sourcemaps.write('./'))
        .pipe(gulp.dest('dist/css'))
})

gulp.task('watchimage', function () {
    gulp.watch('src/images/**/*', function (event) {
        var paths = watchPath(event,'src/','dist/')

        gutil.log(gutil.colors.green(event.type) + ' ' + paths.srcPath)
        gutil.log('Dist ' + paths.distPath)

        gulp.src(paths.srcPath)
            .pipe(imagemin({
                progressive: true
            }))
            .pipe(gulp.dest(paths.distDir))
    })
})

gulp.task('image', function () {
    gulp.src('src/images/**/*')
        .pipe(imagemin({
            progressive: true
        }))
        .pipe(gulp.dest('dist/images'))
})

gulp.task('watchcopy', function () {
    gulp.watch('src/fonts/**/*', function (event) {
        var paths = watchPath(event,'src/', 'dist/')

        gutil.log(gutil.colors.green(event.type) + ' ' + paths.srcPath)
        gutil.log('Dist ' + paths.distPath)

        gulp.src(paths.srcPath)
            .pipe(gulp.dest(paths.distDir))
    })
})

gulp.task('copy', function () {
    gulp.src('src/fonts/**/*')
        .pipe(gulp.dest('dist/fonts/'))
})

gulp.task('watchtemplates', function () {
    gulp.watch('src/templates/**/*', function (event) {
        var paths = watchPath(event, 'src/', 'dist/')

        gutil.log(gutil.colors.green(event.type) + ' ' + paths.srcPath)
        gutil.log('Dist ' + paths.distPath)

        var combined = combiner.obj([
            gulp.src(paths.srcPath),
            handlebars({
              // 3.0.1
              handlebars: require('handlebars')
            }),
            wrap('Handlebars.template(<%= contents %>)'),
            declare({
              namespace: 'S.templates',
              noRedeclare: true
            }),
            gulp.dest(paths.distDir)
        ])
        combined.on('error', handleError)
    })
})

gulp.task('templates', function () {
        gulp.src('src/templates/**/*')
        .pipe(handlebars({
          // 3.0.1
          handlebars: require('handlebars')
        }))
        .pipe(wrap('Handlebars.template(<%= contents %>)'))
        .pipe(declare({
          namespace: 'S.templates',
          noRedeclare: true
        }))
        .pipe(gulp.dest('dist/templates'))
})


gulp.task('default', [
    // build
    'uglifyjs', 'minifycss', 'lesscss', 'sasscss', 'image', 'copy', 'templates',
    // watch
    'watchjs', 'watchcss', 'watchless', 'watchsass', 'watchimage', 'watchcopy', 'watchtemplates'
    ]
)
```