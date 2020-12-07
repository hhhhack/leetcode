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
 * @Date: 2020-12-07 11:11:33
 * @LastEditTime: 2020-12-07 14:48:33
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/MoveCount.cpp
 * @
 */

#include "ComStruct.h"
#include <iostream>
using namespace std;
int move(vector<vector<int>> &ret, int threshold, int rows, int cols){
    if (rows > ret.size() - 1 || cols > ret[0].size() - 1 || rows < 0 || cols < 0){
        return 0;
    }
    cout << "run here"<<endl;
    int cont = 0;
    int trows = rows, tcols = cols;
    if (ret[rows][cols] != 0){
        return 0;
    }
    while(trows > 0){
        cont += trows % 10;
        trows = trows / 10;
    }
    while(tcols > 0){
        cont += tcols % 10;
        tcols = tcols / 10;
    }
    if (cont > threshold){
        ret[rows][cols] = -1;
        return 0;
    }
    //cout<< "cont is " << cont << " rows is "<< rows << " cols is "<< cols << endl;
    ret[rows][cols] = 1;
    return 1 + move(ret, threshold, rows - 1, cols) + move(ret, threshold, rows + 1, cols) + move(ret, threshold, rows, cols - 1) + move(ret, threshold, rows, cols + 1);
}

int movingCount(int threshold, int rows, int cols)
{
    vector<vector<int>> ret(rows, vector<int>(cols, 0));
    return move(ret, threshold, 0, 0);
}

int main(){
    cout<<movingCount(5, 10, 10);
    return 0;
}