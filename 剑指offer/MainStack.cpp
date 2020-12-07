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
 * @Date: 2020-12-05 09:51:15
 * @LastEditTime: 2020-12-05 09:58:02
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/MainStack.cpp
 * @
 */

#include <stack>
using namespace std;

class MainStack {
public:
    void push(int value) {
        if (min_sta.size() == 0 || min_sta.top() > value){
            min_sta.push(value);
        }
        data_sta.push(value);
    }
    void pop() {
        if (data_sta.top() == min_sta.top()){
            min_sta.pop();
        }
        data_sta.pop();
    }
    int top() {
        return data_sta.top();
    }
    int min() {
        return min_sta.top();
    }
private:
    stack<int> data_sta;
    stack<int> min_sta;
};