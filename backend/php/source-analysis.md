
### PHP 源码分析
获取源码：http://php.net/git.php  
解读参考：https://github.com/laruence/php7-internal  

- PHP 源码结构  
> 只关注两个目录：ext 和 Zend，其他的文件和目录对于 PHP 扩展和开发来说很重要。  
> 
> PHP 程序被分为两个主要的部分。  
> 第一部分是 Zend 引擎，控制 PHP 代码运行时候的运行环境。它处理 PHP 提供的所有“语言层”的特性，包括：变量，表达式，语法解析，代码执行和错误处理。可以说，没有这个引擎，就没有 PHP。引擎的源码就放在了 Zend 目录。  
> 第二个部分是包含在 PHP 里面的扩展。这些扩展包括我们可以在 PHP 调用的每一个核心函数（例如 strpos,substr,array_diff,mysql_connect 等等）。也包括核心的类（MySQLi,SplFixedArray,PDO 等等）。  
> 
> 在核心代码中，决定在哪里找到你想查看的功能最简单的方法是，查看 [PHP 的文档首页](https://secure.php.net/manual/zh/index.php)。  
> PHP 的文档也被分为两个主要的部分：语言参考和函数参考。  
> 作为一个庞大的概括，如果想查看的是在语言参考中的定义，很有可能可以在 Zend 文件夹找到。如果是在函数参考中，可以在 ext 文件夹中找到。  
> 

- 一些基本的 C 语言概念
> 变量  
> 在 C 语言里面，变量是静态和强类型的。这意味着变量必须要使用一个类型定义之后才能使用。一旦定义之后，你不能改变它的类型（你可以在之后转换成其他类型，但你需要使用不同的变量来实现）。  
> 因为，在 C 语言里面，变量并不真实地存在，它们只是为了我们使用的方便的内存地址的标签。正因为如此，C 语言没有 PHP 中的引用。取而代之，它有指针。为了我们的目的，把指针想象成指向其他变量的变量。把它当作 PHP 中变量的变量。  
> 变量名前面可以有一个或这多个符号。星号(*)表明变量是指向某个类型的指针（一个引用），两个星号表明变量是指向指针的指针，三个星号表明变量是指向一个指向其他指针的指针。  
> 
> 预处理说明  
> C 在编译之前使用一步叫做 “预处理” 的步骤。这一步包含优化和根据你传递给编译器的选项动态使用部分代码。  
> 下面说说两个主要的预处理器说明：条件语句和宏。  
> 条件语句允许代码在编译输出或者不是基于定义时被引入。这看起来很像下面的例子。这允许不同的代码根据不同的操作系统被使用（因此尽管它们使用不同的 API，也可以在 Windows 和 Linux 中很好的使用）。另外，它允许一部分代码被引入或者不是基于定义的指示。事实上，这是配置步骤中如何编译 PHP 的执行过程。  
> 宏，这是最简单简化代码的迷你函数。它们不是真正的函数，但是在编译预处理是会执行简单的文本替换。因此，宏不会真正地调用函数。你可以为函数定义写一个宏，宏允许在预处理编译时使用更简单的代码。  
> 
> 源文件  
> 在 C 源码使用的类型的文件，主要有两种文件：.c 和.h。  
> .c 文件是包含了源码准备编译的文件。通常来说，.c 文件包含了不能分享到其他文件的私有函数的实现。  
> .h（或者说头文件）定义了在.c 文件中可以被其他文件看到的函数，包括预处理宏。  
> 头文件定义公共 API 的方式，是通过不使用函数体重新声明函数的签名（跟 PHP 中的接口和抽象方法相似）。这样，源码就可以通过头文件 链接在一起了。  

- PHP 内部函数的定义
> PHP 内部函数一般是在 ext/standard 文件夹，以 .h 后缀结尾的文件和以 .c 后缀结尾的文件。以 .h 后缀结尾的文件是一个典型的头文件的样子：单纯的函数列表，函数在其他地方定义；而 .c 文件包含了函数真正的源代码。  
> 所有的 PHP 函数都使用同一个基本结构。在函数顶部定义了各个变量，然后调用 zend_parse_parameters 函数，然后到了主要的逻辑，当中有 RETURN_*** 和 php_error_docref 的调用。
```c
// strpos 函数的定义
zval *needle; // 定义一个指向 zval 的指针 needle（zval 是在 PHP 内部代表任意一个 PHP 变量的定义）
char *haystack; // 定义了指向单个字符的指针 haystack，haystack 变量会指向你所传递的 $haystack 字符串变量的第一个字符。haystack + 1 会指向第二个字符，haystack + 2 指向第三个，以此类推。因此，通过逐个递增指针，可以读取整个字符串。
char *found = NULL;
char  needle_char[2]; // 在 C 语言里，数组代表指向它们第一个元素的指针
long  offset = 0; // 这个变量用来保存函数的第三个参数：开始搜索的偏移量
int   haystack_len; // PHP 需要知道字符串在哪里结束（不然的话，它会一直递增指针而不会停止）为了解决这个问题，PHP 也保存了明确的长度，这就是 haystack_len 变量

// 获取传递到函数的参数，然后把它们存储到上面声明的变量中
// 传递参数的数量，通过 ZEND_NUM_ARGS() 宏提供
// TSRMLS_CC 宏，这是 PHP 的一种特性；它是线程安全资源管理器（TSRM）的一部分，它保证 PHP 不会在多线程之间混乱变量
// ”sz|l” 字符串标记了函数接收的参数，s 第一个参数是字符串，z 第二个参数是一个zval结构体、任意的变量，| 标识接下来的参数是可选的，l 第三个参数是long类型（整型）
// &haystack，&haystack_len，&needle，&offset 指定了需要赋值的参数的变量
if (zend_parse_parameters(ZEND_NUM_ARGS() TSRMLS_CC, "sz|l", &haystack, &haystack_len, &needle, &offset) == FAILURE) {
    return;
}
```

- PHP 变量的实现
> 在 PHP 的核心代码中，变量被称为 ZVAL。  
> ZVAL 这个结构之所以那么重要是有原因的，不仅仅是因为 PHP 使用弱类型而 C 使用强类型
```c
struct _zval_struct {
    /* Variable information */
    zvalue_value value; /* value */
    zend_uint refcount__gc; // 指向 PHP 变量容器的指针的计数器，表示这个 zval 的引用数目
    zend_uchar type; /* active type，比如使用 zval.type = IS_LONG 来定义整型数据 */
    zend_uchar is_ref__gc; // 标识变量是否为引用
};

// union 是单独的类型，它根据怎么被访问而使用不同的方式解释
typedef union _zvalue_value {
    long lval; /* long value */
    double dval; /* double value */
    struct {
        char *val;
        int len;
    } str;
    HashTable *ht; /* hash table value */
    zend_object_value obj;
} zvalue_value;
```

- PHP 数组的实现
> PHP 里面的所有东西都是哈希表。在 C 里面，数组是内存块，你可以通过下标访问这些内存块。即在 C 里面的数组只能使用整数且有序的键值。  
> 哈希表是这样的：它们使用哈希函数转换字符串键值为正常的整型键值。哈希后的结果可以被作为正常的 C 数组的键值（又名为内存块）。现在的问题是，哈希函数会有冲突，那就是说，多个字符串键值可能会生成一样的哈希值。这个问题可以通过存储可能冲突的值到链表中，而不是直接将值存储到生成的下标里。  
> Zend Engine 定义了大量的 API 函数供哈希表使用。低级的哈希表函数预览可以在 zend_hash.h 文件里面找到。另外 Zend Engine 在 zend_API.h 文件定义了稍微高级一些的 API。  
```c
// hash 表结构
typedef struct _hashtable {
    uint nTableSize;
    uint nTableMask;
    uint nNumOfElements;
    ulong nNextFreeElement;
    Bucket *pInternalPointer;
    Bucket *pListHead;
    Bucket *pListTail;
    Bucket **arBuckets;
    dtor_func_t pDestructor;
    zend_bool persistent;
    unsigned char nApplyCount;
    zend_bool bApplyProtection;
     #if ZEND_DEBUG
        int inconsistent;
     #endif
} HashTable;
```
