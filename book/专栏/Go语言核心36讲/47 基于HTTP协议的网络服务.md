<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=47&#32;基于HTTP协议的网络服务>
        <link rel="icon" href="/static/favicon.png">
        <title>47 基于HTTP协议的网络服务 </title>
        
        <link rel="stylesheet" href="/static/index.css">
        <link rel="stylesheet" href="/static/highlight.min.css">
        <script src="/static/highlight.min.js"></script>
        
        <meta name="generator" content="Hexo 4.2.0">
        <script defer src="https://umami.lianglianglee.com/script.js"
         data-website-id="83e5d5db-9d06-40e3-b780-cbae722fdf8c"></script>
    </head>

<body>
    <div class="book-container">
        <div class="book-sidebar">
            <div class="book-brand">
                <a href="/">
                    <img src="/static/favicon.png">
                    <span>技术文章摘抄</span>
                </a>
            </div>
            <div class="book-menu uncollapsible">
                <ul class="uncollapsible">
                    <li><a href="/" class="current-tab">首页</a></li>
                    <li><a href="../">上一级</a></li>
                </ul>
                <ul class="uncollapsible">
                    
                    <li>
                        <a class="menu-item" id="00 导读 写给0基础入门的Go语言学习者.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/00%20%e5%af%bc%e8%af%bb%20%e5%86%99%e7%bb%990%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8%e7%9a%84Go%e8%af%ad%e8%a8%80%e5%ad%a6%e4%b9%a0%e8%80%85.md">00 导读 写给0基础入门的Go语言学习者.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 导读 学习专栏的正确姿势.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/00%20%e5%af%bc%e8%af%bb%20%e5%ad%a6%e4%b9%a0%e4%b8%93%e6%a0%8f%e7%9a%84%e6%ad%a3%e7%a1%ae%e5%a7%bf%e5%8a%bf.md">00 导读 学习专栏的正确姿势.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 开篇词 跟着学，你也能成为Go语言高手.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e8%b7%9f%e7%9d%80%e5%ad%a6%ef%bc%8c%e4%bd%a0%e4%b9%9f%e8%83%bd%e6%88%90%e4%b8%baGo%e8%af%ad%e8%a8%80%e9%ab%98%e6%89%8b.md">00 开篇词 跟着学，你也能成为Go语言高手.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 工作区和GOPATH.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/01%20%e5%b7%a5%e4%bd%9c%e5%8c%ba%e5%92%8cGOPATH.md">01 工作区和GOPATH.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 命令源码文件.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/02%20%e5%91%bd%e4%bb%a4%e6%ba%90%e7%a0%81%e6%96%87%e4%bb%b6.md">02 命令源码文件.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 库源码文件.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/03%20%e5%ba%93%e6%ba%90%e7%a0%81%e6%96%87%e4%bb%b6.md">03 库源码文件.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 程序实体的那些事儿（上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/04%20%e7%a8%8b%e5%ba%8f%e5%ae%9e%e4%bd%93%e7%9a%84%e9%82%a3%e4%ba%9b%e4%ba%8b%e5%84%bf%ef%bc%88%e4%b8%8a%ef%bc%89.md">04 程序实体的那些事儿（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 程序实体的那些事儿（中）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/05%20%e7%a8%8b%e5%ba%8f%e5%ae%9e%e4%bd%93%e7%9a%84%e9%82%a3%e4%ba%9b%e4%ba%8b%e5%84%bf%ef%bc%88%e4%b8%ad%ef%bc%89.md">05 程序实体的那些事儿（中）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 程序实体的那些事儿 （下）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/06%20%e7%a8%8b%e5%ba%8f%e5%ae%9e%e4%bd%93%e7%9a%84%e9%82%a3%e4%ba%9b%e4%ba%8b%e5%84%bf%20%ef%bc%88%e4%b8%8b%ef%bc%89.md">06 程序实体的那些事儿 （下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 数组和切片.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/07%20%e6%95%b0%e7%bb%84%e5%92%8c%e5%88%87%e7%89%87.md">07 数组和切片.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 container包中的那些容器.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/08%20container%e5%8c%85%e4%b8%ad%e7%9a%84%e9%82%a3%e4%ba%9b%e5%ae%b9%e5%99%a8.md">08 container包中的那些容器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 字典的操作和约束.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/09%20%e5%ad%97%e5%85%b8%e7%9a%84%e6%93%8d%e4%bd%9c%e5%92%8c%e7%ba%a6%e6%9d%9f.md">09 字典的操作和约束.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 通道的基本操作.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/10%20%e9%80%9a%e9%81%93%e7%9a%84%e5%9f%ba%e6%9c%ac%e6%93%8d%e4%bd%9c.md">10 通道的基本操作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 通道的高级玩法.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/11%20%e9%80%9a%e9%81%93%e7%9a%84%e9%ab%98%e7%ba%a7%e7%8e%a9%e6%b3%95.md">11 通道的高级玩法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 使用函数的正确姿势.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/12%20%e4%bd%bf%e7%94%a8%e5%87%bd%e6%95%b0%e7%9a%84%e6%ad%a3%e7%a1%ae%e5%a7%bf%e5%8a%bf.md">12 使用函数的正确姿势.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 结构体及其方法的使用法门.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/13%20%e7%bb%93%e6%9e%84%e4%bd%93%e5%8f%8a%e5%85%b6%e6%96%b9%e6%b3%95%e7%9a%84%e4%bd%bf%e7%94%a8%e6%b3%95%e9%97%a8.md">13 结构体及其方法的使用法门.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 接口类型的合理运用.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/14%20%e6%8e%a5%e5%8f%a3%e7%b1%bb%e5%9e%8b%e7%9a%84%e5%90%88%e7%90%86%e8%bf%90%e7%94%a8.md">14 接口类型的合理运用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 关于指针的有限操作.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/15%20%e5%85%b3%e4%ba%8e%e6%8c%87%e9%92%88%e7%9a%84%e6%9c%89%e9%99%90%e6%93%8d%e4%bd%9c.md">15 关于指针的有限操作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 go语句及其执行规则（上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/16%20go%e8%af%ad%e5%8f%a5%e5%8f%8a%e5%85%b6%e6%89%a7%e8%a1%8c%e8%a7%84%e5%88%99%ef%bc%88%e4%b8%8a%ef%bc%89.md">16 go语句及其执行规则（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 go语句及其执行规则（下）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/17%20go%e8%af%ad%e5%8f%a5%e5%8f%8a%e5%85%b6%e6%89%a7%e8%a1%8c%e8%a7%84%e5%88%99%ef%bc%88%e4%b8%8b%ef%bc%89.md">17 go语句及其执行规则（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 if语句、for语句和switch语句.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/18%20if%e8%af%ad%e5%8f%a5%e3%80%81for%e8%af%ad%e5%8f%a5%e5%92%8cswitch%e8%af%ad%e5%8f%a5.md">18 if语句、for语句和switch语句.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 错误处理（上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/19%20%e9%94%99%e8%af%af%e5%a4%84%e7%90%86%ef%bc%88%e4%b8%8a%ef%bc%89.md">19 错误处理（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 错误处理 （下）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/20%20%e9%94%99%e8%af%af%e5%a4%84%e7%90%86%20%ef%bc%88%e4%b8%8b%ef%bc%89.md">20 错误处理 （下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 panic函数、recover函数以及defer语句 （上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/21%20panic%e5%87%bd%e6%95%b0%e3%80%81recover%e5%87%bd%e6%95%b0%e4%bb%a5%e5%8f%8adefer%e8%af%ad%e5%8f%a5%20%ef%bc%88%e4%b8%8a%ef%bc%89.md">21 panic函数、recover函数以及defer语句 （上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 panic函数、recover函数以及defer语句（下）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/22%20panic%e5%87%bd%e6%95%b0%e3%80%81recover%e5%87%bd%e6%95%b0%e4%bb%a5%e5%8f%8adefer%e8%af%ad%e5%8f%a5%ef%bc%88%e4%b8%8b%ef%bc%89.md">22 panic函数、recover函数以及defer语句（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 测试的基本规则和流程 （上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/23%20%e6%b5%8b%e8%af%95%e7%9a%84%e5%9f%ba%e6%9c%ac%e8%a7%84%e5%88%99%e5%92%8c%e6%b5%81%e7%a8%8b%20%ef%bc%88%e4%b8%8a%ef%bc%89.md">23 测试的基本规则和流程 （上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 测试的基本规则和流程（下）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/24%20%e6%b5%8b%e8%af%95%e7%9a%84%e5%9f%ba%e6%9c%ac%e8%a7%84%e5%88%99%e5%92%8c%e6%b5%81%e7%a8%8b%ef%bc%88%e4%b8%8b%ef%bc%89.md">24 测试的基本规则和流程（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 更多的测试手法.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/25%20%e6%9b%b4%e5%a4%9a%e7%9a%84%e6%b5%8b%e8%af%95%e6%89%8b%e6%b3%95.md">25 更多的测试手法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 sync.Mutex与sync.RWMutex.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/26%20sync.Mutex%e4%b8%8esync.RWMutex.md">26 sync.Mutex与sync.RWMutex.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 条件变量sync.Cond （上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/27%20%e6%9d%a1%e4%bb%b6%e5%8f%98%e9%87%8fsync.Cond%20%ef%bc%88%e4%b8%8a%ef%bc%89.md">27 条件变量sync.Cond （上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 条件变量sync.Cond （下）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/28%20%e6%9d%a1%e4%bb%b6%e5%8f%98%e9%87%8fsync.Cond%20%ef%bc%88%e4%b8%8b%ef%bc%89.md">28 条件变量sync.Cond （下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 原子操作（上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/29%20%e5%8e%9f%e5%ad%90%e6%93%8d%e4%bd%9c%ef%bc%88%e4%b8%8a%ef%bc%89.md">29 原子操作（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 原子操作（下）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/30%20%e5%8e%9f%e5%ad%90%e6%93%8d%e4%bd%9c%ef%bc%88%e4%b8%8b%ef%bc%89.md">30 原子操作（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 sync.WaitGroup和sync.Once.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/31%20sync.WaitGroup%e5%92%8csync.Once.md">31 sync.WaitGroup和sync.Once.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 context.Context类型.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/32%20context.Context%e7%b1%bb%e5%9e%8b.md">32 context.Context类型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 临时对象池sync.Pool.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/33%20%e4%b8%b4%e6%97%b6%e5%af%b9%e8%b1%a1%e6%b1%a0sync.Pool.md">33 临时对象池sync.Pool.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 并发安全字典sync.Map （上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/34%20%e5%b9%b6%e5%8f%91%e5%ae%89%e5%85%a8%e5%ad%97%e5%85%b8sync.Map%20%ef%bc%88%e4%b8%8a%ef%bc%89.md">34 并发安全字典sync.Map （上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 并发安全字典sync.Map (下).md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/35%20%e5%b9%b6%e5%8f%91%e5%ae%89%e5%85%a8%e5%ad%97%e5%85%b8sync.Map%20%28%e4%b8%8b%29.md">35 并发安全字典sync.Map (下).md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 unicode与字符编码.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/36%20unicode%e4%b8%8e%e5%ad%97%e7%ac%a6%e7%bc%96%e7%a0%81.md">36 unicode与字符编码.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 strings包与字符串操作.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/37%20strings%e5%8c%85%e4%b8%8e%e5%ad%97%e7%ac%a6%e4%b8%b2%e6%93%8d%e4%bd%9c.md">37 strings包与字符串操作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 bytes包与字节串操作（上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/38%20bytes%e5%8c%85%e4%b8%8e%e5%ad%97%e8%8a%82%e4%b8%b2%e6%93%8d%e4%bd%9c%ef%bc%88%e4%b8%8a%ef%bc%89.md">38 bytes包与字节串操作（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 bytes包与字节串操作（下）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/39%20bytes%e5%8c%85%e4%b8%8e%e5%ad%97%e8%8a%82%e4%b8%b2%e6%93%8d%e4%bd%9c%ef%bc%88%e4%b8%8b%ef%bc%89.md">39 bytes包与字节串操作（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 io包中的接口和工具 （上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/40%20io%e5%8c%85%e4%b8%ad%e7%9a%84%e6%8e%a5%e5%8f%a3%e5%92%8c%e5%b7%a5%e5%85%b7%20%ef%bc%88%e4%b8%8a%ef%bc%89.md">40 io包中的接口和工具 （上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 io包中的接口和工具 （下）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/41%20io%e5%8c%85%e4%b8%ad%e7%9a%84%e6%8e%a5%e5%8f%a3%e5%92%8c%e5%b7%a5%e5%85%b7%20%ef%bc%88%e4%b8%8b%ef%bc%89.md">41 io包中的接口和工具 （下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 bufio包中的数据类型 （上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/42%20bufio%e5%8c%85%e4%b8%ad%e7%9a%84%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b%20%ef%bc%88%e4%b8%8a%ef%bc%89.md">42 bufio包中的数据类型 （上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 bufio包中的数据类型（下）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/43%20bufio%e5%8c%85%e4%b8%ad%e7%9a%84%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b%ef%bc%88%e4%b8%8b%ef%bc%89.md">43 bufio包中的数据类型（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 使用os包中的API （上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/44%20%e4%bd%bf%e7%94%a8os%e5%8c%85%e4%b8%ad%e7%9a%84API%20%ef%bc%88%e4%b8%8a%ef%bc%89.md">44 使用os包中的API （上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 使用os包中的API （下）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/45%20%e4%bd%bf%e7%94%a8os%e5%8c%85%e4%b8%ad%e7%9a%84API%20%ef%bc%88%e4%b8%8b%ef%bc%89.md">45 使用os包中的API （下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46 访问网络服务.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/46%20%e8%ae%bf%e9%97%ae%e7%bd%91%e7%bb%9c%e6%9c%8d%e5%8a%a1.md">46 访问网络服务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="47 基于HTTP协议的网络服务.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/47%20%e5%9f%ba%e4%ba%8eHTTP%e5%8d%8f%e8%ae%ae%e7%9a%84%e7%bd%91%e7%bb%9c%e6%9c%8d%e5%8a%a1.md">47 基于HTTP协议的网络服务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="48 程序性能分析基础（上）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/48%20%e7%a8%8b%e5%ba%8f%e6%80%a7%e8%83%bd%e5%88%86%e6%9e%90%e5%9f%ba%e7%a1%80%ef%bc%88%e4%b8%8a%ef%bc%89.md">48 程序性能分析基础（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="49 程序性能分析基础（下）.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/49%20%e7%a8%8b%e5%ba%8f%e6%80%a7%e8%83%bd%e5%88%86%e6%9e%90%e5%9f%ba%e7%a1%80%ef%bc%88%e4%b8%8b%ef%bc%89.md">49 程序性能分析基础（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="尾声 愿你披荆斩棘，所向无敌.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/%e5%b0%be%e5%a3%b0%20%e6%84%bf%e4%bd%a0%e6%8a%ab%e8%8d%86%e6%96%a9%e6%a3%98%ef%bc%8c%e6%89%80%e5%90%91%e6%97%a0%e6%95%8c.md">尾声 愿你披荆斩棘，所向无敌.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="新年彩蛋 完整版思考题答案.md" href="/%e4%b8%93%e6%a0%8f/Go%e8%af%ad%e8%a8%80%e6%a0%b8%e5%bf%8336%e8%ae%b2/%e6%96%b0%e5%b9%b4%e5%bd%a9%e8%9b%8b%20%e5%ae%8c%e6%95%b4%e7%89%88%e6%80%9d%e8%80%83%e9%a2%98%e7%ad%94%e6%a1%88.md">新年彩蛋 完整版思考题答案.md</a>
                    </li>
                    
                    <li><a href="https://lianglianglee.com/assets/%E6%8D%90%E8%B5%A0.md">捐赠</a></li>
                </ul>

            </div>
        </div>

        <div class="sidebar-toggle" onclick="sidebar_toggle()" onmouseover="add_inner()" onmouseleave="remove_inner()">
            <div class="sidebar-toggle-inner"></div>
        </div>
        <div class="off-canvas-content">
            <div class="columns">
                <div class="column col-12 col-lg-12">
                    <div class="book-navbar">
                        
                        <header class="navbar">
                            <section class="navbar-section">
                                <a onclick="open_sidebar()">
                                    <i class="icon icon-menu"></i>
                                </a>
                            </section>
                        </header>
                    </div>
                    <div class="book-content" style="max-width: 960px; margin: 0 auto;
    overflow-x: auto;
    overflow-y: hidden;">
                        <div class="book-post">
                            
                            
                            
                            <p id="tip" align="center"></p>
                            <h1 id="title" data-id="47 基于HTTP协议的网络服务" class="title">47 基于HTTP协议的网络服务</h1>
                            <div><p>我们在上一篇文章中简单地讨论了网络编程和socket，并由此提及了Go语言标准库中的<code>syscall</code>代码包和<code>net</code>代码包。</p>

<p>我还重点讲述了<code>net.Dial</code>函数和<code>syscall.Socket</code>函数的参数含义。前者间接地调用了后者，所以正确理解后者，会对用好前者有很大裨益。</p>

<p>之后，我们把视线转移到了<code>net.DialTimeout</code>函数以及它对操作超时的处理上，这又涉及了<code>net.Dialer</code>类型。实际上，这个类型正是<code>net</code>包中这两个“拨号”函数的底层实现。</p>

<p>我们像上一篇文章的示例代码那样用<code>net.Dial</code>或<code>net.DialTimeout</code>函数来访问基于HTTP协议的网络服务是完全没有问题的。HTTP协议是基于TCP/IP协议栈的，并且它也是一个面向普通文本的协议。</p>

<p>原则上，我们使用任何一个文本编辑器，都可以轻易地写出一个完整的HTTP请求报文。只要你搞清楚了请求报文的头部（header）和主体（body）应该包含的内容，这样做就会很容易。所以，在这种情况下，即便直接使用<code>net.Dial</code>函数，你应该也不会感觉到困难。</p>

<p>不过，不困难并不意味着很方便。如果我们只是访问基于HTTP协议的网络服务的话，那么使用<code>net/http</code>代码包中的程序实体来做，显然会更加便捷。</p>

<p>其中，最便捷的是使用<code>http.Get</code>函数。我们在调用它的时候只需要传给它一个URL就可以了，比如像下面这样：</p>

<pre><code>url1 := &quot;http://google.cn&quot;
fmt.Printf(&quot;Send request to %q with method GET ...\n&quot;, url1)
resp1, err := http.Get(url1)
if err != nil {
	fmt.Printf(&quot;request sending error: %v\n&quot;, err)
}
defer resp1.Body.Close()
line1 := resp1.Proto + &quot; &quot; + resp1.Status
fmt.Printf(&quot;The first line of response:\n%s\n&quot;, line1)
</code></pre>

<p><code>http.Get</code>函数会返回两个结果值。第一个结果值的类型是<code>*http.Response</code>，它是网络服务给我们传回来的响应内容的结构化表示。</p>

<p>第二个结果值是<code>error</code>类型的，它代表了在创建和发送HTTP请求，以及接收和解析HTTP响应的过程中可能发生的错误。</p>

<p><code>http.Get</code>函数会在内部使用缺省的HTTP客户端，并且调用它的<code>Get</code>方法以完成功能。这个缺省的HTTP客户端是由<code>net/http</code>包中的公开变量<code>DefaultClient</code>代表的，其类型是<code>*http.Client</code>。它的基本类型也是可以被拿来使用的，甚至它还是开箱即用的。下面的这两行代码：</p>

<pre><code>var httpClient1 http.Client
resp2, err := httpClient1.Get(url1)
</code></pre>

<p>与前面的这一行代码</p>

<pre><code>resp1, err := http.Get(url1)
</code></pre>

<p>是等价的。</p>

<p><code>http.Client</code>是一个结构体类型，并且它包含的字段都是公开的。之所以该类型的零值仍然可用，是因为它的这些字段要么存在着相应的缺省值，要么其零值直接就可以使用，且代表着特定的含义。</p>

<p>现在，我问你一个问题，是关于这个类型中的最重要的一个字段的。</p>

<p><strong>今天的问题是：<code>http.Client</code>类型中的<code>Transport</code>字段代表着什么？</strong></p>

<p>这道题的<strong>典型回答</strong>是这样的。</p>

<p><code>http.Client</code>类型中的<code>Transport</code>字段代表着：向网络服务发送HTTP请求，并从网络服务接收HTTP响应的操作过程。也就是说，该字段的方法<code>RoundTrip</code>应该实现单次HTTP事务（或者说基于HTTP协议的单次交互）需要的所有步骤。</p>

<p>这个字段是<code>http.RoundTripper</code>接口类型的，它有一个由<code>http.DefaultTransport</code>变量代表的缺省值（以下简称<code>DefaultTransport</code>）。当我们在初始化一个<code>http.Client</code>类型的值（以下简称<code>Client</code>值）的时候，如果没有显式地为该字段赋值，那么这个<code>Client</code>值就会直接使用<code>DefaultTransport</code>。</p>

<p>顺便说一下，<code>http.Client</code>类型的<code>Timeout</code>字段，代表的正是前面所说的单次HTTP事务的超时时间，它是<code>time.Duration</code>类型的。它的零值是可用的，用于表示没有设置超时时间。</p>

<h2 id="问题解析">问题解析</h2>

<p>下面，我们再通过该字段的缺省值<code>DefaultTransport</code>，来深入地了解一下这个<code>Transport</code>字段。</p>

<p><code>DefaultTransport</code>的实际类型是<code>*http.Transport</code>，后者即为<code>http.RoundTripper</code>接口的默认实现。这个类型是可以被复用的，也推荐被复用，同时，它也是并发安全的。正因为如此，<code>http.Client</code>类型也拥有着同样的特质。</p>

<p><code>http.Transport</code>类型，会在内部使用一个<code>net.Dialer</code>类型的值（以下简称<code>Dialer</code>值），并且，它会把该值的<code>Timeout</code>字段的值，设定为<code>30</code>秒。</p>

<p>也就是说，这个<code>Dialer</code>值如果在30秒内还没有建立好网络连接，那么就会被判定为操作超时。在<code>DefaultTransport</code>的值被初始化的时候，这样的<code>Dialer</code>值的<code>DialContext</code>方法会被赋给前者的<code>DialContext</code>字段。</p>

<p><code>http.Transport</code>类型还包含了很多其他的字段，其中有一些字段是关于操作超时的。</p>

<ul>
<li><code>IdleConnTimeout</code>：含义是空闲的连接在多久之后就应该被关闭。</li>
<li><code>DefaultTransport</code>会把该字段的值设定为<code>90</code>秒。如果该值为<code>0</code>，那么就表示不关闭空闲的连接。注意，这样很可能会造成资源的泄露。</li>
<li><code>ResponseHeaderTimeout</code>：含义是，从客户端把请求完全递交给操作系统到从操作系统那里接收到响应报文头的最大时长。<code>DefaultTransport</code>并没有设定该字段的值。</li>
<li><code>ExpectContinueTimeout</code>：含义是，在客户端递交了请求报文头之后，等待接收第一个响应报文头的最长时间。在客户端想要使用HTTP的“POST”方法把一个很大的报文体发送给服务端的时候，它可以先通过发送一个包含了“Expect: 100-continue”的请求报文头，来询问服务端是否愿意接收这个大报文体。这个字段就是用于设定在这种情况下的超时时间的。注意，如果该字段的值不大于<code>0</code>，那么无论多大的请求报文体都将会被立即发送出去。这样可能会造成网络资源的浪费。<code>DefaultTransport</code>把该字段的值设定为了<code>1</code>秒。</li>
<li><code>TLSHandshakeTimeout</code>：TLS是Transport Layer Security的缩写，可以被翻译为传输层安全。这个字段代表了基于TLS协议的连接在被建立时的握手阶段的超时时间。若该值为<code>0</code>，则表示对这个时间不设限。<code>DefaultTransport</code>把该字段的值设定为了<code>10</code>秒。</li>
</ul>

<p>此外，还有一些与<code>IdleConnTimeout</code>相关的字段值得我们关注，即：<code>MaxIdleConns</code>、<code>MaxIdleConnsPerHost</code>以及<code>MaxConnsPerHost</code>。</p>

<p>无论当前的<code>http.Transport</code>类型的值（以下简称<code>Transport</code>值）访问了多少个网络服务，<code>MaxIdleConns</code>字段都只会对空闲连接的总数做出限定。而<code>MaxIdleConnsPerHost</code>字段限定的则是，该<code>Transport</code>值访问的每一个网络服务的最大空闲连接数。</p>

<p>每一个网络服务都会有自己的网络地址，可能会使用不同的网络协议，对于一些HTTP请求也可能会用到代理。<code>Transport</code>值正是通过这三个方面的具体情况，来鉴别不同的网络服务的。</p>

<p><code>MaxIdleConnsPerHost</code>字段的缺省值，由<code>http.DefaultMaxIdleConnsPerHost</code>变量代表，值为<code>2</code>。也就是说，在默认情况下，对于某一个<code>Transport</code>值访问的每一个网络服务，它的空闲连接数都最多只能有两个。</p>

<p>与<code>MaxIdleConnsPerHost</code>字段的含义相似的，是<code>MaxConnsPerHost</code>字段。不过，后者限制的是，针对某一个<code>Transport</code>值访问的每一个网络服务的最大连接数，不论这些连接是否是空闲的。并且，该字段没有相应的缺省值，它的零值表示不对此设限。</p>

<p><code>DefaultTransport</code>并没有显式地为<code>MaxIdleConnsPerHost</code>和<code>MaxConnsPerHost</code>这两个字段赋值，但是它却把<code>MaxIdleConns</code>字段的值设定为了<code>100</code>。</p>

<p>换句话说，在默认情况下，空闲连接的总数最大为<code>100</code>，而针对每个网络服务的最大空闲连接数为<code>2</code>。注意，上述两个与空闲连接数有关的字段的值应该是联动的，所以，你有时候需要根据实际情况来定制它们。</p>

<p>当然了，这首先需要我们在初始化<code>Client</code>值的时候，定制它的<code>Transport</code>字段的值。定制这个值的方式，可以参看<code>DefaultTransport</code>变量的声明。</p>

<p>最后，我简单说一下为什么会出现空闲的连接。我们都知道，HTTP协议有一个请求报文头叫做“Connection”。在HTTP协议的1.1版本中，这个报文头的值默认是“keep-alive”。</p>

<p>在这种情况下的网络连接都是持久连接，它们会在当前的HTTP事务完成后仍然保持着连通性，因此是可以被复用的。</p>

<p>既然连接可以被复用，那么就会有两种可能。一种可能是，针对于同一个网络服务，有新的HTTP请求被递交，该连接被再次使用。另一种可能是，不再有对该网络服务的HTTP请求，该连接被闲置。</p>

<p>显然，后一种可能就产生了空闲的连接。另外，如果分配给某一个网络服务的连接过多的话，也可能会导致空闲连接的产生，因为每一个新递交的HTTP请求，都只会征用一个空闲的连接。所以，为空闲连接设定限制，在大多数情况下都是很有必要的，也是需要斟酌的。</p>

<p>如果我们想彻底地杜绝空闲连接的产生，那么可以在初始化<code>Transport</code>值的时候把它的<code>DisableKeepAlives</code>字段的值设定为<code>true</code>。这时，HTTP请求的“Connection”报文头的值就会被设置为“close”。这会告诉网络服务，这个网络连接不必保持，当前的HTTP事务完成后就可以断开它了。</p>

<p>如此一来，每当一个HTTP请求被递交时，就都会产生一个新的网络连接。这样做会明显地加重网络服务以及客户端的负载，并会让每个HTTP事务都耗费更多的时间。所以，在一般情况下，我们都不要去设置这个<code>DisableKeepAlives</code>字段。</p>

<p>顺便说一句，在<code>net.Dialer</code>类型中，也有一个看起来很相似的字段<code>KeepAlive</code>。不过，它与前面所说的HTTP持久连接并不是一个概念，<code>KeepAlive</code>是直接作用在底层的socket上的。</p>

<p>它的背后是一种针对网络连接（更确切地说，是TCP连接）的存活探测机制。它的值用于表示每间隔多长时间发送一次探测包。当该值不大于<code>0</code>时，则表示不开启这种机制。<code>DefaultTransport</code>会把这个字段的值设定为<code>30</code>秒。</p>

<p>好了，以上这些内容阐述的就是，<code>http.Client</code>类型中的<code>Transport</code>字段的含义，以及它的值的定制方式。这涉及了<code>http.RoundTripper</code>接口、<code>http.DefaultTransport</code>变量、<code>http.Transport</code>类型，以及<code>net.Dialer</code>类型。</p>

<h2 id="知识扩展">知识扩展</h2>

<h3 id="问题-http-server-类型的-listenandserve-方法都做了哪些事情">问题：<code>http.Server</code>类型的<code>ListenAndServe</code>方法都做了哪些事情？</h3>

<p><code>http.Server</code>类型与<code>http.Client</code>是相对应的。<code>http.Server</code>代表的是基于HTTP协议的服务端，或者说网络服务。</p>

<p><code>http.Server</code>类型的<code>ListenAndServe</code>方法的功能是：监听一个基于TCP协议的网络地址，并对接收到的HTTP请求进行处理。这个方法会默认开启针对网络连接的存活探测机制，以保证连接是持久的。同时，该方法会一直执行，直到有严重的错误发生或者被外界关掉。当被外界关掉时，它会返回一个由<code>http.ErrServerClosed</code>变量代表的错误值。</p>

<p>对于本问题，典型回答可以像下面这样。</p>

<p>这个<code>ListenAndServe</code>方法主要会做下面这几件事情。</p>

<ol>
<li>检查当前的<code>http.Server</code>类型的值（以下简称当前值）的<code>Addr</code>字段。该字段的值代表了当前的网络服务需要使用的网络地址，即：IP地址和端口号. 如果这个字段的值为空字符串，那么就用<code>&quot;:http&quot;</code>代替。也就是说，使用任何可以代表本机的域名和IP地址，并且端口号为<code>80</code>。</li>
<li>通过调用<code>net.Listen</code>函数在已确定的网络地址上启动基于TCP协议的监听。</li>
<li>检查<code>net.Listen</code>函数返回的错误值。如果该错误值不为<code>nil</code>，那么就直接返回该值。否则，通过调用当前值的<code>Serve</code>方法准备接受和处理将要到来的HTTP请求。</li>
</ol>

<p>可以从当前问题直接衍生出的问题一般有两个，一个是“<code>net.Listen</code>函数都做了哪些事情”，另一个是“<code>http.Server</code>类型的<code>Serve</code>方法是怎样接受和处理HTTP请求的”。</p>

<p><strong>对于第一个直接的衍生问题，如果概括地说，回答可以是：</strong></p>

<ol>
<li>解析参数值中包含的网络地址隐含的IP地址和端口号；</li>
<li>根据给定的网络协议，确定监听的方法，并开始进行监听。</li>
</ol>

<p>从这里的第二个步骤出发，我们还可以继续提出一些间接的衍生问题。这往往会涉及<code>net.socket</code>函数以及相关的socket知识。</p>

<p><strong>对于第二个直接的衍生问题，我们可以这样回答：</strong></p>

<p>在一个<code>for</code>循环中，网络监听器的<code>Accept</code>方法会被不断地调用，该方法会返回两个结果值；第一个结果值是<code>net.Conn</code>类型的，它会代表包含了新到来的HTTP请求的网络连接；第二个结果值是代表了可能发生的错误的<code>error</code>类型值。</p>

<p>如果这个错误值不为<code>nil</code>，除非它代表了一个暂时性的错误，否则循环都会被终止。如果是暂时性的错误，那么循环的下一次迭代将会在一段时间之后开始执行。</p>

<p>如果这里的<code>Accept</code>方法没有返回非<code>nil</code>的错误值，那么这里的程序将会先把它的第一个结果值包装成一个<code>*http.conn</code>类型的值（以下简称<code>conn</code>值），然后通过在新的goroutine中调用这个<code>conn</code>值的<code>serve</code>方法，来对当前的HTTP请求进行处理。</p>

<p>这个处理的细节还是很多的，所以我们依然可以找出不少的间接的衍生问题。比如，这个<code>conn</code>值的状态有几种，分别代表着处理的哪个阶段？又比如，处理过程中会用到哪些读取器和写入器，它们的作用分别是什么？再比如，这里的程序是怎样调用我们自定义的处理函数的，等等。</p>

<p>诸如此类的问题很多，我就不在这里一一列举和说明了。你只需要记住一句话：“源码之前了无秘密”。上面这些问题的答案都可以在Go语言标准库的源码中找到。如果你想对本问题进行深入的探索，那么一定要去看<code>net/http</code>代码包的源码。</p>

<h2 id="总结">总结</h2>

<p>今天，我们主要讲的是基于HTTP协议的网络服务，侧重点仍然在客户端。</p>

<p>我们在讨论了<code>http.Get</code>函数和<code>http.Client</code>类型的简单使用方式之后，把目光聚焦在了后者的<code>Transport</code>字段。</p>

<p>这个字段代表着单次HTTP事务的操作过程。它是<code>http.RoundTripper</code>接口类型的。它的缺省值由<code>http.DefaultTransport</code>变量代表，其实际类型是<code>*http.Transport</code>。</p>

<p><code>http.Transport</code>包含的字段非常多。我们先讲了<code>DefaultTransport</code>中的<code>DialContext</code>字段会被赋予什么样的值，又详细说明了一些关于操作超时的字段。</p>

<p>比如<code>IdleConnTimeout</code>和<code>ExpectContinueTimeout</code>，以及相关的<code>MaxIdleConns</code>和<code>MaxIdleConnsPerHost</code>等等。之后，我又简单地解释了出现空闲连接的原因，以及相关的定制方式。</p>

<p>最后，作为扩展，我还为你简要地梳理了<code>http.Server</code>类型的<code>ListenAndServe</code>方法，执行的主要流程。不过，由于篇幅原因，我没有做深入讲述。但是，这并不意味着没有必要深入下去。相反，这个方法很重要，值得我们认真地去探索一番。</p>

<p>在你需要或者有兴趣的时候，我希望你能去好好地看一看<code>net/http</code>包中的相关源码。一切秘密都在其中。</p>

<h2 id="思考题">思考题</h2>

<p>我今天留给你的思考题比较简单，即：怎样优雅地停止基于HTTP协议的网络服务程序？</p>

<p><a href="https://github.com/hyper0x/Golang_Puzzlers" target="_blank">戳此查看Go语言专栏文章配套详细代码。</a></p>
</div>
                        </div>
                        <div>
                            <div id="prePage" style="float: left">

                            </div>
                            <div id="nextPage" style="float: right">

                            </div>
                        </div>

                    </div>
                </div>
            </div>
            <div class="copyright">
                <hr />
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#cea2a2a2f7fafffffef98ea9a3afa7a2e0ada1a3" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0e04effbc6653b',t:'MTczNDAwOTcwNC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>