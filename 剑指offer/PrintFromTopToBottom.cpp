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
 * @Date: 2020-12-05 11:01:57
 * @LastEditTime: 2020-12-05 11:12:43
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/PrintFromTopToBottom.cpp
 * @
 */
#include "ComStruct.h"
using namespace std;
vector<int> PrintFromTopToBottom(TreeNode* root){
    queue<TreeNode*> treeque;
    treeque.push(root);
    vector<int>ret;
    while(!treeque.empty()){
        TreeNode *tmp = *treeque.front()->left;
        if (tmp){
            treeque.push(tmp);
        }
        tmp = *treeque.front()->right;
        if (tmp){
            treeque.push(tmp);
        }
        ret.push_back(*treeque.front()->val);
        treeque.pop();
    }
    return ret;
}