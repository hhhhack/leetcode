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
 * @Date: 2020-12-16 15:46:36
 * @LastEditTime: 2020-12-16 17:20:03
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/算法导论/heapSort.cpp
 * @
 */

#include "Comalg.h"
using namespace std;
#define PARENT(i) ((i - 1)/2)
#define LEFT(i) (2 * (i) + 1)
#define RIGHT(i) (2 * ((i) + 1)) 

void max_heapity(vector<int> &input, int i){
    int largest = i;
    if (LEFT(i) < input.size() && input[i] < input[LEFT(i)]){
        largest = LEFT(i);
    }
    if (RIGHT(i) < input.size() && input[largest] < input[RIGHT(i)]){
        largest = RIGHT(i);
    }
    if (i != largest){
        swap(input, i, largest);
        max_heapity(input, largest);
    }
}

void BuildMaxHeap(vector<int>input){
    for(int i = input.size() / 2; i >= 0; i--){
        max_heapity(input, i);
    }
}

void heapSort(vector<int> &input){
    vector<int> ret(input.size());
    BuildMaxHeap(input);
    for(int i = 1; i < input.size(); i++){

    }
}