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
 * @Date: 2020-12-05 09:20:38
 * @LastEditTime: 2020-12-05 10:46:15
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/HasSubtree.cpp
 * @
 */

#include "ComStruct.h"
#include <iostream>
using namespace std;
bool CmpTree(TreeNode* pRoot1, TreeNode* pRoot2){
    if (pRoot2 == NULL)
    {
        return true;
    }else if (pRoot2 != NULL && pRoot1 == NULL){
        return false;
    }else if (pRoot1->val == pRoot2->val){
        if (CmpTree(pRoot1->left, pRoot2->left) && CmpTree(pRoot1->right, pRoot2->right))
            return true;
    }
    return false;
}

bool HasSubtree(TreeNode* pRoot1, TreeNode* pRoot2)
{
    if (pRoot1 == NULL || pRoot2 == NULL)
    {
        return false;
    }
    if (CmpTree(pRoot1, pRoot2))
    {
        return true;
    }
    return HasSubtree(pRoot1->right, pRoot2) || HasSubtree(pRoot1->right, pRoot2);
}
int main(){
    TreeNode *pRoot1 = new TreeNode(8);
    TreeNode *cur = pRoot1;
    cur->left = new TreeNode(8);
    cur = cur->left;
    cur->left = new TreeNode(9);
    cur = cur->left;
    cur->left = new TreeNode(2);
    cur = cur->left;
    cur->left = new TreeNode(5);
    TreeNode *pRoot2 = new TreeNode(8);
    cur->left = new TreeNode(9);
    cur = cur->left;
    cur->left = new TreeNode(2);
    cout<<HasSubtree(pRoot1, pRoot2);
    return 0;
}
