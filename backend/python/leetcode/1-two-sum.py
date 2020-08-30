# -*- coding:utf-8 -*-


"""求解两个数之和

    Args:
        numbers: int 数列
        target: int 目标和

	Returns:
	    两数的位置
"""
def two_sum(numbers, target):
    numbers_len = len(numbers)
    if numbers_len != 0:
        for i, item in enumerate(numbers):
            for j in range(i + 1, numbers_len):
                if numbers[i] + numbers[j] == target:
                    return i, j

    return -1, -1


def main():
    # 给定一个数组和一个目标和，从数组中找两个数字相加等于目标和，输出这两个数字的下标
    # etc. numbers = [2, 3, 4], target = 5; return [0, 1]
    numbers = [2, 3, 4]
    target = 5
    result = two_sum(numbers, target)
    print(result)


if __name__ == '__main__':
    main()
