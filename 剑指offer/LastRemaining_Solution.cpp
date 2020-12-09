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
 * @Date: 2020-12-08 19:14:49
 * @LastEditTime: 2020-12-08 20:00:23
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/LastRemaining_Solution.cpp
 * @
 */
#include "ComStruct.h"
using namespace std;
int LastRemaining_Solution(int n, int m)
{
    if (n == 0){
        return -1;
    }
    ListNode *Head = new ListNode(0);
    Head->next = Head;
    ListNode *tmp = NULL;
    for (int i = 1; i < n; i++){
        tmp = new ListNode(i);
        tmp->next = Head->next;
        Head->next = tmp;
        Head = Head->next;
    }
    while(Head->next != Head){
        int k = m % n - 1;
        if (k == -1){
            k = m - 1;
        }
        for(int i = 0; i < k; i++){
            Head = Head->next;
        }
        cout<<Head->next->val<<endl;
        tmp = Head->next;
        Head->next = tmp->next;
        tmp->next = NULL;
        n--;
    }
    return Head->val;
}
int main(){
    cout << LastRemaining_Solution(5, 3) << endl;
    return 0;
}