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
 * @Date: 2020-12-03 10:31:55
 * @LastEditTime: 2020-12-03 11:04:21
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/ReplaceSpace.cpp
 * @
 */
#include <cstring>
#include <stdlib.h>

void replaceSpace(char *str,int length){
    if (str == NULL){
        return ;
    }
    int j = 0, count = 0;
    while (str[j] != '\0'){
        if (str[j++] == ' ')
            count ++;
    }
    int ret_len = j + 2 * count;
    if (ret_len < length){
        return;
    }
    while(j >= 0 && ret_len >= 0){
        if (str[j] == ' ' && ret_len >= 2){
            str[ret_len--] = '0';
            str[ret_len--] = '2';
            str[ret_len--] = '%';
        }else{
            str[ret_len--] = str[j];
        }
        j--;
    }
}
