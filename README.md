# 思路实现
copy from https://github.com/tlkh/asitop

# 为什么要做这个
1、练手项目
2、没有较好的mac展示软件(free)
3、asiop使用python做的 依赖python环境 本人mac经常换py环境

# 大体思路
1、powermetrics 刷数据到临时文件
2、根据临时文件匹配数据
3、展示


#######################################
####  MacBookAir10,1 ...... 23D56  ####
####  GPU        0.00 Ghz   0.02 % ####
####  E-Cluster  1.00 Ghz   0.19 % ####
####  P-Cluster  0.00 Ghz   0.03 % ####
#####  2024-05-31 17:48:50 +0000 UTC ##
#####  2024-06-01 01:48:56 ############
#######################################

# 问题 
1、powermetrics 到 temp 太慢了 可能4s-5s才能刷一次
2、powermetrics 转为数据流也不快
3、半成品
