
```python
# -*- coding:utf-8 -*-
# 代码风格遵循 pep8
# 建议使用工具 pylint、pep8 (autopep8)、flake8
# 每行不超过 80 个字符

# 这是单行注释
foo = 'xxx'

""" 这是多行注释的总结陈词

描述 1
描述 2
描述 3
...
"""

THIS_IS_A_CONSTANT = 1
this_is_a_variable = 1

# 如果一个文本字符串在一行放不下，可以使用圆括号来实现隐式行连接
x = ('This will build a very long long '
     'long long long long long long string')
# 在注释中，如果必要，将长的 URL 放在一行上
# See details at
# http://www.example.com/us/developer/documentation/api/content/v2.0/csv_file_name_extension_full_specification.html

def func_name(arg1, arg2):
    """在这里写函数的一句话总结(如: 计算平均值)

    这里是具体描述

    Args:
        arg1: int arg1 的具体描述
        arg2: int arg2 的具体描述

	Returns:
	    int 返回值的具体描述

	[Raises]:
	    异常类型: 异常具体描述
    """
    pass


def _private_func():
    pass

class ThisIsAClass(object):
	"""在这里写类的一句话总结

    这里是类的具体描述

    Attributes:
        attr1: attr1 的具体描述
        attr2: attr2 的具体描述
    """

    def __init__(self, likes_spam=False):
        """Inits SampleClass with blah."""
        self.likes_spam = likes_spam
        self.eggs = 0

    def public_method(self):
        """Performs operation blah."""

# TODO(xx@gmail.com): Use a "*" here for string repetition.

def main():
      pass

if __name__ == '__main__':
    main()
```