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
 * @Date: 2020-12-16 09:59:30
 * @LastEditTime: 2020-12-16 17:24:25
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/算法导论/Comalg.h
 * @
 */

#include <vector>
#include <algorithm>
#include <iostream>
#include <iomanip>
#include <ctime>

void inline swap(std::vector<int> &input, int i, int j){
    int tmp = input[i];
    input[i] = input[j];
    input[j] =tmp;
}

void insertSort(std::vector<int> &input);

void mergeSort(std::vector<int>& input, int i, int j);

void QuickSort(std::vector<int> &input, int i, int j);