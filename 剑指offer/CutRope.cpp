/*
 * 
 * 　　┏┓　　　┏┓+ +
 * 　┏┛┻━━━┛┻┓ + +
 * 　┃　　　　　　　┃ 　
 * 　┃　　　━　　　┃ ++ + + +
 *  ████━████ ┃+
 * 　┃　　　　　　　┃ +
 * 　┃　　　┻　　　┃
 * 　┃　　　　　　　┃ + +
 * 　┗━┓　　　┏━┛
 * 　　　┃　　　┃　　　　　　　　　　　
 * 　　　┃　　　┃ + + + +
 * 　　　┃　　　┃
 * 　　　┃　　　┃ +  神兽保佑
 * 　　　┃　　　┃    代码无bug　　
 * 　　　┃　　　┃　　+　　　　　　　　　
 * 　　　┃　 　　┗━━━┓ + +
 * 　　　┃ 　　　　　　　┣┓
 * 　　　┃ 　　　　　　　┏┛
 * 　　　┗┓┓┏━┳┓┏┛ + + + +
 * 　　　　┃┫┫　┃┫┫
 * 　　　　┗┻┛　┗┻┛+ + + +
 * 
 * 
 * @Author: hhhhack
 * @Date: 2020-12-07 10:34:40
 * @LastEditTime: 2021-05-24 11:27:31
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /leetcode/剑指offer/CutRope.cpp
 * @
 */

#include "math.h"
int cutRope(int number)
{
    if (number == 2)
    {
        return 1;
    }
    else if (number == 3)
    {
        return 2;
    }
    else
    {
        if (number % 3 == 0)
        {
            return (int)pow(3, number / 3);
        }
        else if (number % 3 == 1)
        {
            return (int)pow(3, (number / 3) - 1) * 4;
        }
        else
        {
            return (int)pow(3, number / 3) * 2;
        }
    }
    return 0;
}