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
 * @Date: 2021-05-24 10:57:08
 * @LastEditTime: 2021-05-24 12:04:46
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /leetcode/剑指offer/ReList.cpp
 * 
 */
#include <vector>
#include "ComStruct.h"
#include <iostream>
using namespace std;

void ListNode::output()
{
    ListNode *tmp = this;
    while (tmp != nullptr)
    {
        cout << tmp->val;
        tmp = tmp->next;
        if (tmp != nullptr)
        {
            cout << " -> ";
        }
    }
    cout << endl;
}

ListNode *reList(ListNode *head, int m, int n)
{
    if (m > n || m < 0)
    {
        cout << "input err m is " << m << " n is " << n << endl;
        return nullptr;
    }
    if (m == n)
    {
        return head;
    }
    ListNode *cur = head, *pre = nullptr, *next = nullptr;
    int i = 1;
    while (cur != nullptr)
    {
        if (i < m)
        {
            pre = cur;
        }
        if (i <= n)
        {
            next = cur->next;
        }
        else
        {
            break;
        }
        cur = cur->next;
        i++;
    }
    cout << "pre is " << (pre == nullptr ? "null" : to_string(pre->val)) << " next is " << (next == nullptr ? "null" : to_string(next->val)) << endl;
    if (cur == nullptr && i < n)
    {
        cout << "out of list";
        return nullptr;
    }
    ListNode *tmp_head = nullptr, *tmp = nullptr;
    if (pre == nullptr)
    {
        tmp_head = head;
    }
    else
    {
        tmp_head = pre->next;
    }
    for (int i = m; i <= n; i++)
    {
        tmp = tmp_head->next;
        tmp_head->next = next;
        next = tmp_head;
        tmp_head = tmp;
    }
    if (pre == nullptr)
    {
        return next;
    }
    else
    {
        pre->next = next;
    }
    return head;
}

int main()
{
    vector<int> data = {1, 5, 6, 8, 0, 43, 75, 76, 3, 2};
    ListNode *head = nullptr, *tmp = nullptr;
    for (auto i : data)
    {
        tmp = head;
        head = new ListNode(i);
        head->next = tmp;
    }
    head->output();
    int start, end;
    cout << "sizeof list is " << data.size() << endl;
    cin >> start >> end;
    tmp = reList(head, start, end);
    tmp->output();
    return 0;
}