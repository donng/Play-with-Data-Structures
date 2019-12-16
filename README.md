# Play-with-Data-Structures

本仓库使用 Go 语言实现常用的数据结构。实现逻辑基于 liuyubobobo 老师在慕课网上的实战课程[《玩儿转数据结构》](https://coding.imooc.com/class/207.html)，以及使用 Java 实现的原仓库地址：https://github.com/liuyubobobo/Play-with-Data-Structures 。

代码当前正基于 go module 和 [effective-go](https://go-zh.org/doc/effective_go.html) 重写中，原版本请查看 [tag v1.0](https://github.com/donng/Play-with-Data-Structures/tree/v1.0)。

重写进度：

- [x] 02-Arrays
- [x] 03-Stacks-and-Queues
- [ ] 04-Linked-List
- [ ] 05-Recursion
- [ ] 06-Binary-Search-Tree
- [ ] 07-Set-and-Map
- [ ] 08-Heap-and-Priority-Queue
- [ ] 09-Segment-Tree
- [ ] 10-Trie
- [ ] 11-Union-Find
- [ ] 12-Avl-Tree
- [ ] 13-Red-Black-Tree
- [ ] 14-Hash-Table 

## Installation

```
cd $GOPATH/src

# 基于Go Module的实现版本
git clone https://github.com/donng/Play-with-Data-Structures.git
# 原基于Go Path的实现版本
git clone -b v1.0 https://github.com/donng/Play-with-Data-Structures.git
```

注意： 本仓库暂未支持 `go module`，需要执行请先确认环境变量 `GO111MODULE=off`，然后到 main.og 对应的目录下执行 `go run main.go` 即可。

## 数据结构的分类

### 1. 线性结构

数组，栈，队列，链表，哈希表...

### 2. 树结构

二叉树，二分搜索树，AVL，红黑树，Treap，Splay，堆，Trie，线段树，K-D树，并查集，哈夫曼树...

### 3. 图结构

邻接矩阵，邻接表

## 此项目涉及的数据结构
不包含图结构，图论领域以算法为主。

- 数组 栈 队列 链表
- 二分搜索树 堆 线段树 Trie
- 并查集 AVL 红黑树 哈希表
