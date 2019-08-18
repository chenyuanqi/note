
### 工具类
Java 的核心库提供了大量的现成的类供我们使用。  

org.apache.commons.io.IOUtils  
> closeQuietly：关闭一个 IO 流、socket、或者 selector 且不抛出异常，通常放在 finally 块  
> toString：转换 IO 流、 Uri、 byte[] 为 String  
> copy：IO 流数据复制，从输入流写到输出流中，最大支持 2GB  
> toByteArray：从输入流、URI 获取 byte[]  
> write：把字节. 字符等写入输出流  
> toInputStream：把字符转换为输入流  
> readLines：从输入流中读取多行数据，返回 List<String>  
> copyLarge：同 copy，支持 2GB 以上数据的复制  
> lineIterator：从输入流返回一个迭代器，根据参数要求读取的数据量，全部读取，如果数据不够，则失败  

org.apache.commons.io.FileUtils   
> deleteDirectory：删除文件夹    
> readFileToString：以字符形式读取文件内容  
> deleteQueitly：删除文件或文件夹且不会抛出异常  
> copyFile：复制文件  
> writeStringToFile：把字符写到目标文件，如果文件不存在，则创建  
> forceMkdir：强制创建文件夹，如果该文件夹父级目录不存在，则创建父级  
> write：把字符写到指定文件中  
> listFiles：列举某个目录下的文件(根据过滤器)  
> copyDirectory：复制文件夹  
> forceDelete：强制删除文件  

org.apache.commons.lang.StringUtils  
> isBlank：字符串是否为空 (trim 后判断)  
> isEmpty：字符串是否为空 (不 trim 并判断)  
> equals：字符串是否相等  
> join：合并数组为单一字符串，可传分隔符  
> split：分割字符串  
> EMPTY：返回空字符串  
> trimToNull：trim 后为空字符串则转换为 null  
> replace：替换字符串  

org.apache.http.util.EntityUtils
> toString：把 Entity 转换为字符串  
> consume：确保 Entity 中的内容全部被消费。可以看到源码里又一次消费了 Entity 的内容，假如用户没有消费，那调用 Entity 时候将会把它消费掉  
> toByteArray：把 Entity 转换为字节流  
> consumeQuietly：和 consume 一样，但不抛异常  
> getContentCharset：获取内容的编码  

org.apache.commons.lang3.StringUtils \ org.apache.commons.lang.StringEscapeUtils   
> isBlank：字符串是否为空 (trim 后判断)  
> isEmpty：字符串是否为空 (不 trim 并判断)  
> equals：字符串是否相等  
> join：合并数组为单一字符串，可传分隔符  
> split：分割字符串  
> EMPTY：返回空字符串  
> replace：替换字符串  
> capitalize：首字符大写  

org.apache.commons.io.FilenameUtils  
> getExtension：返回文件后缀名  
> getBaseName：返回文件名，不包含后缀名  
> getName：返回文件全名  
> concat：按命令行风格组合文件路径(详见方法注释)  
> removeExtension：删除后缀名  
> normalize：使路径正常化  
> wildcardMatch：匹配通配符  
> seperatorToUnix：路径分隔符改成 unix 系统格式的，即 /  
> getFullPath：获取文件路径，不包括文件名  
> isExtension：检查文件后缀名是不是传入参数(List<String>)中的一个  

org.springframework.util.StringUtils   
> hasText：检查字符串中是否包含文本  
> hasLength：检测字符串是否长度大于 0  
> isEmpty：检测字符串是否为空（若传入为对象，则判断对象是否为 null）  
> commaDelimitedStringToArray：逗号分隔的 String 转换为数组  
> collectionToDelimitedString：把集合转为 CSV 格式字符串  
> replace 替换字符串  
> delimitedListToStringArray：相当于 split  
> uncapitalize：首字母小写  
> collectionToDelimitedCommaString：把集合转为 CSV 格式字符串  
> tokenizeToStringArray：和 split 基本一样，但能自动去掉空白的单词  

org.apache.commons.lang.ArrayUtils  
> contains：是否包含某字符串  
> addAll：添加整个数组  
> clone：克隆一个数组  
> isEmpty：是否空数组  
> add：向数组添加元素  
> subarray：截取数组  
> indexOf：查找某个元素的下标   
> isEquals：比较数组是否相等  
> toObject：基础类型数据数组转换为对应的 Object 数组  

org.apache.http.client.utils.URLEncodedUtils  
> format：格式化参数，返回一个 HTTP POST 或者 HTTP PUT 可用 application/x-www-form-urlencoded 字符串  
> parse：把 String 或者 URI 等转换为 List<NameValuePair>  

org.apache.commons.codec.digest.DigestUtils  
> md5Hex：MD5 加密，返回 32 位字符串  
> sha1Hex：SHA-1 加密  
> sha256Hex：SHA-256 加密  
> sha512Hex：SHA-512 加密  
> md5：MD5 加密，返回 16 位字符串  

org.apache.commons.collections.CollectionUtils  
> isEmpty：是否为空  
> select：根据条件筛选集合元素  
> transform：根据指定方法处理集合元素，类似 List 的 map()  
> filter：过滤元素，类似 List 的 filter()  
> find：基本和 select 一样  
> collect：和 transform 差不多一样，但是返回新数组  
> forAllDo：调用每个元素的指定方法  
> isEqualCollection：判断两个集合是否一致  

org.apache.commons.lang3.ArrayUtils  
> contains：是否包含某个字符串  
> addAll：添加整个数组  
> clone：克隆一个数组  
> isEmpty：是否空数组  
> add：向数组添加元素  
> subarray：截取数组  
> indexOf：查找某个元素的下标  
> isEquals：比较数组是否相等  
> toObject：基础类型数据数组转换为对应的 Object 数组  

org.apache.commons.beanutils.PropertyUtils  
> getProperty：获取对象属性值  
> setProperty：设置对象属性值  
> getPropertyDiscriptor：获取属性描述器  
> isReadable：检查属性是否可访问  
> copyProperties：复制属性值，从一个对象到另一个对象  
> getPropertyDiscriptors：获取所有属性描述器  
> isWriteable：检查属性是否可写  
> getPropertyType：获取对象属性类型  

org.apache.commons.lang3.StringEscapeUtils 这个现在已经废弃了，建议使用 commons-text 包里面的方法  
> unescapeHtml4：转义 html  
> escapeHtml4：反转义 html  
> escapeXml：转义 xml  
> unescapeXml：反转义 xml  
> escapeJava：转义 unicode 编码  
> escapeEcmaScript：转义 EcmaScript 字符  
> unescapeJava：反转义 unicode 编码  
> escapeJson：转义 json 字符  
> escapeXml10：转义 Xml10  

org.apache.commons.beanutils.BeanUtils  
> copyPeoperties：复制属性值，从一个对象到另一个对象  
> getProperty：获取对象属性值  
> setProperty：设置对象属性值  
> populate：根据 Map 给属性复制  
> copyPeoperty：复制单个值，从一个对象到另一个对象  
> cloneBean：克隆 bean 实例  

#### Math
Math 类就是用来进行数学计算的，它提供了大量的静态方法来便于我们实现数学计算。  
```java
// 数学常量
double pi = Math.PI; // 3.14159...
double e = Math.E; // 2.7182818...

// 求绝对值
Math.abs(-100); // 100
Math.abs(-7.8); // 7.8

// 取最大或最小值
Math.max(100, 99); // 100
Math.min(1.2, 2.3); // 1.2

// 计算 power
Math.pow(2, 10); // 2 的 10 次方 = 1024
// 计算√x
Math.sqrt(2); // 1.414...
// 计算 e 的 x 次方
Math.exp(2); // 7.389...
// 计算以 e 为底的对数
Math.log(4); // 1.386...
// 计算以 10 为底的对数
Math.log10(100); // 2

// 三角函数
Math.sin(3.14); // 0.00159...
Math.sin(Math.PI / 6); // sin(π/6) = 0.5
Math.cos(3.14); // -0.9999...
Math.tan(3.14); // -0.0015...
Math.asin(1.0); // 1.57079...
Math.acos(1.0); // 0.0

// 生成一个随机数 x，x 的范围是 0 <= x < 1
Math.random(); // 0.53907... 每次都不一样
```
StrictMath 保证所有平台计算结果都是完全相同的，而 Math 会尽量针对平台优化计算速度。

#### Random 和 SecureRandom
Random 用来创建伪随机数。  
所谓伪随机数，是指只要给定一个初始的种子，产生的随机数序列是完全一样的。  
```java
// 如果不给定种子，就使用系统当前时间戳作为种子
Random r = new Random();
r.nextInt(); // 2071575453,每次都不一样
r.nextInt(10); // 5,生成一个 [0,10) 之间的 int
r.nextLong(); // 8811649292570369305,每次都不一样
r.nextFloat(); // 0.54335...生成一个[0,1)之间的float
r.nextDouble(); // 0.3716...生成一个[0,1)之间的double

// 指定种子，就会得到完全确定的随机数序列
Random r = new Random(12345);
for (int i = 0; i < 10; i++) {
    System.out.println(r.nextInt(100));
}
```

有伪随机数，就有真随机数。  
实际上真正的真随机数只能通过量子力学原理来获取，而我们想要的是一个不可预测的安全的随机数，SecureRandom 就是用来创建安全的随机数的。  
SecureRandom 无法指定种子，它使用 RNG（random number generator）算法。
```java
SecureRandom sr = new SecureRandom();
System.out.println(sr.nextInt(100));

// JDK 的 SecureRandom 实际上有多种不同的底层实现
// 有的使用安全随机种子加上伪随机数算法来产生安全的随机数，有的使用真正的随机数生成器
// 实际使用的时候，可以优先获取高强度的安全随机数生成器，如果没有提供，再使用普通等级的安全随机数生成器
import java.util.Arrays;
import java.security.SecureRandom;
import java.security.NoSuchAlgorithmException;
SecureRandom sr = null;
try {
    sr = SecureRandom.getInstanceStrong(); // 获取高强度安全随机数生成器
} catch (NoSuchAlgorithmException e) {
    sr = new SecureRandom(); // 获取普通的安全随机数生成器
}
byte[] buffer = new byte[16];
sr.nextBytes(buffer); // 用安全随机数填充 buffer
System.out.println(Arrays.toString(buffer));
```
SecureRandom 的安全性是通过操作系统提供的安全的随机种子来生成随机数。这个种子是通过 CPU 的热噪声、读写磁盘的字节、网络流量等各种随机事件产生的 “熵”。  
在密码学中，安全的随机数非常重要。如果使用不安全的伪随机数，所有加密体系都将被攻破。  
