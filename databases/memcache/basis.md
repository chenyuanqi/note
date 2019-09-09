
### Memcache 疑难杂症
- memcache 与 memcached 的区别
> memcache 是基于 php 开发的，memcached 是基于 c 语言通过 libmemcached 与 memcached 服务器通信，因此性能更好（由于需要事先安装 libmemcached，因此 Windows 下不支持），并且支持的功能特性也更多，推荐使用后者。
