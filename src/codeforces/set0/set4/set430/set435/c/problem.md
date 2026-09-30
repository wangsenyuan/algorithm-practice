# C. Cardiogram

[题目链接](https://codeforces.com/problemset/problem/435/C)

**比赛：** [Codeforces Round 249 (Div. 2)](https://codeforces.com/contest/435)

**标签：** implementation, \*1600

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## 题意

用 ASCII 图形画一条心电图折线（cardiogram）。

心电图是一条折线，由序列正整数 `a1, a2, …, an` 完全确定：从某个点出发，按段交替地画斜向上的 `/` 与斜向下的 `\`，第 `i` 段长度为 `ai`（占 `ai` 个字符）。第一段向上（`/`），第二段向下（`\`），以此类推。

给定序列 `ai`，请输出该折线的 ASCII 图像。

输出共 `max |yi − yj|` 行（`yk` 为折线第 `k` 个点的纵坐标），每行恰好 `∑ ai` 个字符。每个字符只能是 `/`、`\` 或空格。图像必须与给定折线一致；请仔细对照样例。

**注意：** 评测会考虑空格，不要输出多余字符。官方题面也注明样例输出因技术原因不便从网页复制，样例答案见：

- http://assets.codeforces.com/rounds/435/1.txt
- http://assets.codeforces.com/rounds/435/2.txt

## 约束

- 单组数据
- `n`（`2 ≤ n ≤ 1000`）
- `a1, a2, …, an`（`1 ≤ ai ≤ 1000`）
- 保证所有 `ai` 之和不超过 `1000`

## 输入

第一行一个整数 `n`。

第二行 `n` 个整数 `a1, a2, …, an`。

## 输出

输出折线的 ASCII 图像（行末空格也计入答案）。

## 样例

### 样例 1

输入

```
5
3 1 2 5 1
```

输出

```
     /\     
  /\/  \    
 /      \   
/        \  
          \/
```

### 样例 2

输入

```
3
1 5 1
```

输出

```
/\     
  \    
   \   
    \  
     \/
```
