# data structure
go data structure

## 线性表 (linear list)
### 顺序表
### 链表
* 单向链表
* 静态链表
* 循环链表
    * 将单链表中终端结点的指针端由空指针改为指向头结点
* 双向链表
### 栈与队列
1. 栈

 * 分类
   * 栈(顺序存储)
   * 两栈共享空间
   * 链栈(链式存储)
 * 应用
   * 递归
   * 四则运算表达式求职
2. 队列
> 只允许在一端进行插入操作，而在另一端进行删除操作的线性表
* 循环队列
> 头尾相接的顺序存储结构
  * 循环队列(顺序存储)
  * 链队列(链式存储)
### 串
> 串是由零个或多个字符组成的有限序列，又名叫字符串
* 朴素的模式匹配算法
  * 时间复杂度 O((n-m+1)*m)
  * 低效
* KMP模式匹配算法
### 树
> 树是n(n>=0)个结点的有限集
* n=0时称为空树
* 在任意一棵非空树中
  * 有且仅有一个特定的称为根（root）的结点
  * 当n>1时，其余结点可分为m(m>0)个互不相交的有限集T<sub>1</sub>,T<sub>2</sub>...T<sub>m</sub>,其中每一个集合本身又是一棵树,并且称为根的子树
* 结点的度
  * 结点为0的度为叶子结点
  * 度不为0的结点为分支结点或终端结点
  * 树的度是数内各结点的度的最大值
* 二叉树
> 二叉树是n(n>=0)个结点的有限集合，该集合或者为空集（称为空二叉树），或者由一个根结点和两棵互不相交的、分别称为根结点的左子树和右子树的二叉树组成
* 二叉树特点
  * 每个结点最多有两棵子树，所以二叉树中不存在度大于2的结点
  * 左子树和右子树是有顺序的，次序不能任意颠倒
  * 即使树中某结点只有一棵子树，也要区分它是左子树还是右子树
* 二叉树的五种状态
  * 空二叉树
  * 只有一个根结点
  * 根结点只有左子树
  * 根结点只有右子树
  * 根结点既有左子树又有右子树
* 二叉树的分类
  * 斜树
    * 左斜树-所有的结点都只有左子树的二叉树
    * 右斜树-所有的结点都只有右子树的二叉树
  * 满二叉树
    * 在一棵二叉树中，如果所有分支结点都存在左子树和右子树，并且所有叶子都在同一层上
  * 完全二叉树
    * 对一棵具有n个结点的二叉树按层序编号，如果编号为i(1<=i<=n)的结点与同样深度的满二叉树中编号为i的结点在二叉树中位置完全相同
* 二叉树的性质
  * 在二叉树的第i层上至多有2<sup>i-1</sup>个结点（i>=1）
