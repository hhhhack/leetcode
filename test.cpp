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
 * @Date: 2021-02-01 19:12:31
 * @LastEditTime: 2021-02-01 19:19:31
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/c_code/test.cpp
 * @
 */

#include <iostream>
using namespace std;
struct node
{
    int val;
    node *high;
    node *low;
    node *next;
    int end;
    int start;
    node()
    {
        val = 0;
        high = NULL;
        low = NULL;
        next = NULL;
        end = 0;
        start = 0;
    };
    node(int i) : val(i)
    {
        high = NULL;
        low = NULL;
        end = 1;
    };
    node(int i, node *high, node *low, node *next, int end, int start) : val(i), high(high), low(low), next(next), end(end), start(start){};
};

struct jlist
{
    node *top;
    node *base;
    int level;
};

const int n = 10;

jlist data;

node base[n];

int add(int i, int j)
{
    if (i < 0 || i >= n)
    {
        return -1;
    }
    base[i].val += j;

    node *cur = &base[i];
    while (cur->high != NULL)
    {
        cur = cur->high;
        cur->val += j;
    }
    return 0;
}

int dec(int i, int j)
{
    return add(i, (0 - j));
}

int contright(node *cur, int i)
{
    node *tmp = cur->low;
    if (tmp == NULL)
    {
        return cur->val;
    }
    int result = 0;
    while (tmp->end != cur->end)
    {
        if (tmp->end < i)
        {
            tmp = tmp->next;
        }
        else
        {
            if (tmp->start <= i)
            {
                result += tmp->val;
                tmp = tmp->next;
            }
            else
            {
                result += contright(tmp, i);
                tmp = tmp->next;
            }
        }
    }
    return result + tmp->val;
}

int contleft(node *cur, int j)
{
    node *tmp = cur->low;
    if (tmp == NULL)
    {
        return cur->val;
    }
    int result = 0;
    while (tmp->start <= j)
    {
        if (tmp->end <= j)
        {
            result += tmp->val;
            tmp = tmp->next;
        }
        else
        {
            result += contleft(tmp, j);
        }
    }
    return result;
}

int query(int i, int j)
{
    if (i < 0 || i >= n || j < 0 || j >= n)
    {
        return -1;
    }
    int result = 0;
    node *cur = data.top;
    while (true)
    {
        if (cur->start == i && cur->end == j)
        {
            return cur->val;
        }
        else if (cur->start <= i)
        {
            if (cur->end <= i)
            {
                cur = cur->next;
            }
            else if (cur->end >= j)
            {
                cur = cur->low;
            }
            else
            {
                result += contright(cur, i);
                cur = cur->next;
            }
        }
        else
        {
            if (cur->end < j)
            {
                result += cur->val;
                cur = cur->next;
            }
            else if (cur->end == j)
            {
                return result += cur->val;
            }
            else
            {
                return result + contleft(cur, j);
            }
        }
    }
    return result;
}

void init(int *nums)
{
    node *tmp = NULL, *cur = NULL;
    data.base = base;
    data.top = NULL;
    for (int i = 0; i < 10; i++)
    {
        if (i % 3 == 0)
        {
            cur = tmp;
            tmp = new node(0, NULL, &base[i], NULL, i - 1, i);
            if (cur != NULL)
            {
                cur->next = tmp;
            }
        }
        if (data.top == NULL)
        {
            data.top = tmp;
        }
        base[i].val = nums[i];
        base[i].start = i;
        base[i].end = i;
        base[i].high = tmp;
        if (i < 9)
        {
            base[i].next = &base[i + 1];
        }
        else
        {
            base[i].next = NULL;
        }
        tmp->val += base[i].val;
        tmp->end += 1;
    }
}
void deletmem()
{
    //程序直接退出，不用专门释放
}

int main()
{
    int nums[10] = {1, 4, 9, 3, 6, 8, 3, 6, 5, 2};
    init(nums);
    for (int i = 0; i < 10; i++)
    {
        cout << "base " << i << " value is " << base[i].val << " high is " << base[i].high->val << endl;
    }
    add(4, 3);
    dec(1, 1);
    cout << "after op " << endl;
    for (int i = 0; i < 10; i++)
    {
        cout << "base " << i << " value is " << base[i].val << " high is " << base[i].high->val << endl;
    }
    cout << "query 1, 2 " << query(1, 2) << endl;
    cout << "query 1, 1 " << query(1, 1) << endl;
    cout << "query 0, 2 " << query(0, 2) << endl;
    cout << "query 1, 4 " << query(1, 4) << endl;
    deletmem();
    return 0;
}