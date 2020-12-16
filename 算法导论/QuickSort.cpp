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
 * @Date: 2020-12-16 17:00:23
 * @LastEditTime: 2020-12-16 17:26:46
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/算法导论/QuickSort.cpp
 * @
 */

#include "Comalg.h"
using namespace std;

int partition(vector<int> &input, int i, int j){
    int tmp = input[j];
    for(int k = i; k <= j - 1; k++){
        if (input[k] > tmp){
            swap(input, i, k);
            i++;
        }
    }
    swap(input, i, j);
    return i;
}

void QuickSort(vector<int> &input, int i, int j){
    if (i < j){
        int k = partition(input, i, j);
        QuickSort(input, i, k - 1);
        QuickSort(input, k + 1, j);
    }
}