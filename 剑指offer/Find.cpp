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
 * @Date: 2020-12-03 10:18:32
 * @LastEditTime: 2020-12-03 10:30:58
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/Find.cpp
 * @
 */

#include <vector>
#include <algorithm>
using namespace std;

bool Find(vector<vector<int>> martx, int num){
    if (martx.size() == 0){
        return false;
    }
    if (martx.size() == 1){
        return find(martx[0].begin(), martx[0].end(), num) != martx[0].end();
    }
    for (int j = martx[0].size() - 1, i = 0; j >= 0 && i < martx.size();){
        if (martx[i][j] == num){
            return true;
        }else if (martx[i][j] > num){
            j--;
        }else{
            i++;
        }
    }
    return false;
}