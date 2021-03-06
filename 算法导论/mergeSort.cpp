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
 * @Date: 2020-12-16 10:51:44
 * @LastEditTime: 2020-12-18 16:26:44
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/算法导论/mergeSort.cpp
 * @
 */

#include "Comalg.h"
using namespace std;

void merge(vector<int> &input, int i, int k, int j){
    vector<int> a(k - i + 1);
    vector<int> b(j - k);
    for(int m = i; m <= k; m++){
        a[m - i] = input[m];
    }
    for (int n = k + 1; n <= j; n++){
        b[n - k - 1] = input[n];
    }
    int cur = i, cur_a = 0, cur_b = 0;
    
    while(cur <= j){
        if (cur_a > k - i){
            input[cur++] = b[cur_b++];
        }else if (cur_b > j - k - 1){
            input[cur++] = a[cur_a++];
        }else if (a[cur_a] > b[cur_b]){
            input[cur++] = b[cur_b++];
        }else {
            input[cur++] = a[cur_a++];
        }

    }
}

void mergeSort(vector<int>& input, int i, int j){
    if (i < j){
        int q = (i + j) / 2;    
        mergeSort(input, i, q);
        mergeSort(input, q + 1, j);
        merge(input, i, q, j);
    }
}