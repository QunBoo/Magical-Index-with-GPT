# Magical-Index-with-GPT
Magical Index with GPT 魔法禁书目录
## Introduction

GPT Index GPT魔法禁书目录是一个以GPT作为控制核心的接口项目。基于index GPT魔法禁书目录，可以实现通过自然语言调用对应的函数功能，并完成你的某些目的。

Index GPT魔法禁书目录相比于纯GPT的优势在于：

* 安全性：本项目为开源模板，基于本项目，您可以在您的服务器上部署，并通过您的服务器连接到内网访问信息，而不需要上传给GPT
* 成本低：相比于完全通过给GPT输入咒语让其生成，通过程序检索的特性能够大大降低调用GPT所需要的成本，并且更加高效
* 接口丰富：本项目提供和飞书、钉钉等程序的接口，您可以在群内加入“机器人”的方式实现**向机器人小助手交代任务，并让他执行**

## 需求

基于GPT“听得懂人话”的功能特性，让GPT负责解析自然语言，并构建json

```json
{
  "Idea": "GPT是依据什么来推断，并说明自己想用哪些函数的",
  "function": "GPT指示应该调用的函数",
  "param": "GPT从自然语言中提取有效信息，构造成param用来作为函数输入"
}
```

程序基于该json中function字段，调用对应的函数，并将结果返回

## 支持功能

功能

| 功能       | 解释                                                                    | 对应函数        |
| ------------ | ------------------------------------------------------------------------- | ----------------- |
| 解析       | 根据输入的自然语言，分析需要调用哪个功能                                | analysis()      |
| 搜索       | 为了获取信息，需要在google上搜索，以备进一步分析/评价                   | search()        |
| 评价       | 对信息进行评价                                                          | comment\()   |
| 查库SQL    | 编写SQL语句，在用户检查确认后执行查库功能并返回结果                     | makeSQL()       |
| 编写脚本   | 编写python代码，保存为文件，用户检查确认后执行                          | makeScript()    |
| 知识库查找 | 在一个指定的知识库中查找需要的信息，<br />返回查到的信息和对应的知识库url<br /> | findKnowledge() |

### 解析

一切的基础功能，起到**解析**需求的功能，**编排**需要调用的函数，并且说明自己的**思路**

**原理：**

GPT


**函数输入：**


**在什么情况下会使用（trigger）：**


**咒语：**


### 搜索

**原理：**

GPT


**函数输入：**


**在什么情况下会使用：**


**咒语：**


### 查库SQL

**原理：**

GPT


**函数输入：**


**在什么情况下会使用：**


**咒语：**


### 编写脚本

**原理：**

GPT根据提取的需求编写python脚本，并保存在某个位置（服务器上或代码仓库中）


**函数输入：**


**在什么情况下会使用：**


**咒语：**


### 知识库查找

**原理：**

GPT提取需求并生成对应的关键字，在知识库中搜索对应关键字，并返回搜索到的段落，以及**搜索思路（按照xx、xx关键字搜索，得到结果为xxxx）**


**函数输入：**

url（放在配置文件中）、需要搜索的关键字


**在什么情况下会使用：**


**咒语：**



## 接口

### 和飞书的接口支持


### 调用GPT的api接口

参考github轮子

[acheong08/ChatGPT-to-API: Scalable unofficial ChatGPT API for production. (github.com)](https://github.com/acheong08/ChatGPT-to-API)


## 展望，未来

### 通过索引构建工具实现 data connector

GPT通过数据连接器实现对文件的访问，查询(API, PDF, DOC, SQL)

[run-llama/llama_index: LlamaIndex (formerly GPT Index) is a data framework for your LLM applications (github.com)](https://github.com/run-llama/llama_index)

[zilliztech/GPTCache: Semantic cache for LLMs. Fully integrated with LangChain and llama_index. (github.com)](https://github.com/zilliztech/GPTCache)