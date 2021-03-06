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
 * @Date: 2020-12-07 19:20:53
 * @LastEditTime: 2020-12-07 20:50:10
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/MaxInWindows.cpp
 * @
 */

#include "ComStruct.h"
using namespace std;
vector<int> maxInWindows(const vector<int>& num, unsigned int size)
{
    vector<int>ret;
    unsigned int maxIndex = 0;
    for(int i = 0; i < num.size(); i++){
        if (num[i] >= num[maxIndex]){
            maxIndex = i;
        }else if (maxIndex == i - size){
            maxIndex ++;
            for (int j = maxIndex + 1; j <= i; j++){
                if (num[j] >= num[maxIndex]){
                    maxIndex = j;
                }
            }
        }
        if (i >= size - 1){
            ret.push_back(num[maxIndex]);
        }
    }
    return ret;
}