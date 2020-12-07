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
 * @Date: 2020-12-07 21:14:47
 * @LastEditTime: 2020-12-07 21:30:19
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/Multiply.cpp
 * @
 */
#include "ComStruct.h"
using namespace std;

vector<int> multiply(const vector<int>& A) {
    vector<int> a (A.size() + 1, 1);
    vector<int> b (A.size() + 1, 1);
    vector<int> ret(A.size());
    for (int i = 0; i < A.size(); i++){
        a[i + 1] = a[i] * A[i];
        b[A.size() - i - 1] = A[A.size() - i - 1] * b[A.size() - i];
    }
    for (int i = 0; i < A.size(); i++){
        ret[i] = a[i] * b[i + 1];
    }
    return ret;
}