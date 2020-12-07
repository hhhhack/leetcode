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
 * @Date: 2020-12-05 10:52:27
 * @LastEditTime: 2020-12-05 11:01:13
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/IsPopOrder.cpp
 * @
 */

#include <vector>
#include <stack>
using namespace std;

bool IsPopOrder(vector<int> pushV,vector<int> popV) {
    stack<int>sta;
    for (int i = 0, j = 0; i < pushV.size(); i++){
        sta.push(pushV[i]);
        while (sta.size() != 0)
        {
            if (sta.top() == popV[j]){
                sta.pop();
                j++;
            }else{
                break;
            }
        }
    }
    if (sta.size() == 0)
        return true;
    return false;
}