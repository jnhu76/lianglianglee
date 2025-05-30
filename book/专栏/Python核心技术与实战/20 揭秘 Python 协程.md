<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=20&#32;揭秘&#32;Python&#32;协程>
        <link rel="icon" href="/static/favicon.png">
        <title>20 揭秘 Python 协程 </title>
        
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
                        <a class="menu-item" id="00 开篇词 从工程的角度深入理解Python.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%bb%8e%e5%b7%a5%e7%a8%8b%e7%9a%84%e8%a7%92%e5%ba%a6%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3Python.md">00 开篇词 从工程的角度深入理解Python.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 如何逐步突破，成为Python高手？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/01%20%e5%a6%82%e4%bd%95%e9%80%90%e6%ad%a5%e7%aa%81%e7%a0%b4%ef%bc%8c%e6%88%90%e4%b8%baPython%e9%ab%98%e6%89%8b%ef%bc%9f.md">01 如何逐步突破，成为Python高手？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 Jupyter Notebook为什么是现代Python的必学技术？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/02%20Jupyter%20Notebook%e4%b8%ba%e4%bb%80%e4%b9%88%e6%98%af%e7%8e%b0%e4%bb%a3Python%e7%9a%84%e5%bf%85%e5%ad%a6%e6%8a%80%e6%9c%af%ef%bc%9f.md">02 Jupyter Notebook为什么是现代Python的必学技术？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 列表和元组，到底用哪一个？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/03%20%e5%88%97%e8%a1%a8%e5%92%8c%e5%85%83%e7%bb%84%ef%bc%8c%e5%88%b0%e5%ba%95%e7%94%a8%e5%93%aa%e4%b8%80%e4%b8%aa%ef%bc%9f.md">03 列表和元组，到底用哪一个？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 字典、集合，你真的了解吗？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/04%20%e5%ad%97%e5%85%b8%e3%80%81%e9%9b%86%e5%90%88%ef%bc%8c%e4%bd%a0%e7%9c%9f%e7%9a%84%e4%ba%86%e8%a7%a3%e5%90%97%ef%bc%9f.md">04 字典、集合，你真的了解吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 深入浅出字符串.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/05%20%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%ad%97%e7%ac%a6%e4%b8%b2.md">05 深入浅出字符串.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 Python “黑箱”：输入与输出.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/06%20Python%20%e2%80%9c%e9%bb%91%e7%ae%b1%e2%80%9d%ef%bc%9a%e8%be%93%e5%85%a5%e4%b8%8e%e8%be%93%e5%87%ba.md">06 Python “黑箱”：输入与输出.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 修炼基本功：条件与循环.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/07%20%e4%bf%ae%e7%82%bc%e5%9f%ba%e6%9c%ac%e5%8a%9f%ef%bc%9a%e6%9d%a1%e4%bb%b6%e4%b8%8e%e5%be%aa%e7%8e%af.md">07 修炼基本功：条件与循环.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 异常处理：如何提高程序的稳定性？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/08%20%e5%bc%82%e5%b8%b8%e5%a4%84%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8f%90%e9%ab%98%e7%a8%8b%e5%ba%8f%e7%9a%84%e7%a8%b3%e5%ae%9a%e6%80%a7%ef%bc%9f.md">08 异常处理：如何提高程序的稳定性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 不可或缺的自定义函数.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/09%20%e4%b8%8d%e5%8f%af%e6%88%96%e7%bc%ba%e7%9a%84%e8%87%aa%e5%ae%9a%e4%b9%89%e5%87%bd%e6%95%b0.md">09 不可或缺的自定义函数.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 简约不简单的匿名函数.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/10%20%e7%ae%80%e7%ba%a6%e4%b8%8d%e7%ae%80%e5%8d%95%e7%9a%84%e5%8c%bf%e5%90%8d%e5%87%bd%e6%95%b0.md">10 简约不简单的匿名函数.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 面向对象（上）：从生活中的类比说起.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/11%20%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e4%bb%8e%e7%94%9f%e6%b4%bb%e4%b8%ad%e7%9a%84%e7%b1%bb%e6%af%94%e8%af%b4%e8%b5%b7.md">11 面向对象（上）：从生活中的类比说起.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 面向对象（下）：如何实现一个搜索引擎？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/12%20%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e6%90%9c%e7%b4%a2%e5%bc%95%e6%93%8e%ef%bc%9f.md">12 面向对象（下）：如何实现一个搜索引擎？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 搭建积木：Python 模块化.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/13%20%e6%90%ad%e5%bb%ba%e7%a7%af%e6%9c%a8%ef%bc%9aPython%20%e6%a8%a1%e5%9d%97%e5%8c%96.md">13 搭建积木：Python 模块化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 答疑（一）：列表和元组的内部实现是怎样的？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/14%20%e7%ad%94%e7%96%91%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%88%97%e8%a1%a8%e5%92%8c%e5%85%83%e7%bb%84%e7%9a%84%e5%86%85%e9%83%a8%e5%ae%9e%e7%8e%b0%e6%98%af%e6%80%8e%e6%a0%b7%e7%9a%84%ef%bc%9f.md">14 答疑（一）：列表和元组的内部实现是怎样的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 Python对象的比较、拷贝.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/15%20Python%e5%af%b9%e8%b1%a1%e7%9a%84%e6%af%94%e8%be%83%e3%80%81%e6%8b%b7%e8%b4%9d.md">15 Python对象的比较、拷贝.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 值传递，引用传递or其他，Python里参数是如何传递的？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/16%20%e5%80%bc%e4%bc%a0%e9%80%92%ef%bc%8c%e5%bc%95%e7%94%a8%e4%bc%a0%e9%80%92or%e5%85%b6%e4%bb%96%ef%bc%8cPython%e9%87%8c%e5%8f%82%e6%95%b0%e6%98%af%e5%a6%82%e4%bd%95%e4%bc%a0%e9%80%92%e7%9a%84%ef%bc%9f.md">16 值传递，引用传递or其他，Python里参数是如何传递的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 强大的装饰器.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/17%20%e5%bc%ba%e5%a4%a7%e7%9a%84%e8%a3%85%e9%a5%b0%e5%99%a8.md">17 强大的装饰器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 metaclass，是潘多拉魔盒还是阿拉丁神灯？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/18%20metaclass%ef%bc%8c%e6%98%af%e6%bd%98%e5%a4%9a%e6%8b%89%e9%ad%94%e7%9b%92%e8%bf%98%e6%98%af%e9%98%bf%e6%8b%89%e4%b8%81%e7%a5%9e%e7%81%af%ef%bc%9f.md">18 metaclass，是潘多拉魔盒还是阿拉丁神灯？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 深入理解迭代器和生成器.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/19%20%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%e8%bf%ad%e4%bb%a3%e5%99%a8%e5%92%8c%e7%94%9f%e6%88%90%e5%99%a8.md">19 深入理解迭代器和生成器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 揭秘 Python 协程.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/20%20%e6%8f%ad%e7%a7%98%20Python%20%e5%8d%8f%e7%a8%8b.md">20 揭秘 Python 协程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 Python并发编程之Futures.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/21%20Python%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e4%b9%8bFutures.md">21 Python并发编程之Futures.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 并发编程之Asyncio.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/22%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e4%b9%8bAsyncio.md">22 并发编程之Asyncio.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 你真的懂Python GIL（全局解释器锁）吗？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/23%20%e4%bd%a0%e7%9c%9f%e7%9a%84%e6%87%82Python%20GIL%ef%bc%88%e5%85%a8%e5%b1%80%e8%a7%a3%e9%87%8a%e5%99%a8%e9%94%81%ef%bc%89%e5%90%97%ef%bc%9f.md">23 你真的懂Python GIL（全局解释器锁）吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 带你解析 Python 垃圾回收机制.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/24%20%e5%b8%a6%e4%bd%a0%e8%a7%a3%e6%9e%90%20Python%20%e5%9e%83%e5%9c%be%e5%9b%9e%e6%94%b6%e6%9c%ba%e5%88%b6.md">24 带你解析 Python 垃圾回收机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 答疑（二）：GIL与多线程是什么关系呢？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/25%20%e7%ad%94%e7%96%91%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9aGIL%e4%b8%8e%e5%a4%9a%e7%ba%bf%e7%a8%8b%e6%98%af%e4%bb%80%e4%b9%88%e5%85%b3%e7%b3%bb%e5%91%a2%ef%bc%9f.md">25 答疑（二）：GIL与多线程是什么关系呢？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 活都来不及干了，还有空注意代码风格？！.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/26%20%e6%b4%bb%e9%83%bd%e6%9d%a5%e4%b8%8d%e5%8f%8a%e5%b9%b2%e4%ba%86%ef%bc%8c%e8%bf%98%e6%9c%89%e7%a9%ba%e6%b3%a8%e6%84%8f%e4%bb%a3%e7%a0%81%e9%a3%8e%e6%a0%bc%ef%bc%9f%ef%bc%81.md">26 活都来不及干了，还有空注意代码风格？！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 学会合理分解代码，提高代码可读性.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/27%20%e5%ad%a6%e4%bc%9a%e5%90%88%e7%90%86%e5%88%86%e8%a7%a3%e4%bb%a3%e7%a0%81%ef%bc%8c%e6%8f%90%e9%ab%98%e4%bb%a3%e7%a0%81%e5%8f%af%e8%af%bb%e6%80%a7.md">27 学会合理分解代码，提高代码可读性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 如何合理利用assert？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/28%20%e5%a6%82%e4%bd%95%e5%90%88%e7%90%86%e5%88%a9%e7%94%a8assert%ef%bc%9f.md">28 如何合理利用assert？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 巧用上下文管理器和With语句精简代码.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/29%20%e5%b7%a7%e7%94%a8%e4%b8%8a%e4%b8%8b%e6%96%87%e7%ae%a1%e7%90%86%e5%99%a8%e5%92%8cWith%e8%af%ad%e5%8f%a5%e7%b2%be%e7%ae%80%e4%bb%a3%e7%a0%81.md">29 巧用上下文管理器和With语句精简代码.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 真的有必要写单元测试吗？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/30%20%e7%9c%9f%e7%9a%84%e6%9c%89%e5%bf%85%e8%a6%81%e5%86%99%e5%8d%95%e5%85%83%e6%b5%8b%e8%af%95%e5%90%97%ef%bc%9f.md">30 真的有必要写单元测试吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 pdb &amp; cProfile：调试和性能分析的法宝.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/31%20pdb%20&amp;%20cProfile%ef%bc%9a%e8%b0%83%e8%af%95%e5%92%8c%e6%80%a7%e8%83%bd%e5%88%86%e6%9e%90%e7%9a%84%e6%b3%95%e5%ae%9d.md">31 pdb &amp; cProfile：调试和性能分析的法宝.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 答疑（三）：如何选择合适的异常处理方式？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/32%20%e7%ad%94%e7%96%91%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%89%e6%8b%a9%e5%90%88%e9%80%82%e7%9a%84%e5%bc%82%e5%b8%b8%e5%a4%84%e7%90%86%e6%96%b9%e5%bc%8f%ef%bc%9f.md">32 答疑（三）：如何选择合适的异常处理方式？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 带你初探量化世界.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/33%20%e5%b8%a6%e4%bd%a0%e5%88%9d%e6%8e%a2%e9%87%8f%e5%8c%96%e4%b8%96%e7%95%8c.md">33 带你初探量化世界.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 RESTful &amp; Socket：搭建交易执行层核心.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/34%20RESTful%20&amp;%20Socket%ef%bc%9a%e6%90%ad%e5%bb%ba%e4%ba%a4%e6%98%93%e6%89%a7%e8%a1%8c%e5%b1%82%e6%a0%b8%e5%bf%83.md">34 RESTful &amp; Socket：搭建交易执行层核心.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 RESTful &amp; Socket：行情数据对接和抓取.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/35%20RESTful%20&amp;%20Socket%ef%bc%9a%e8%a1%8c%e6%83%85%e6%95%b0%e6%8d%ae%e5%af%b9%e6%8e%a5%e5%92%8c%e6%8a%93%e5%8f%96.md">35 RESTful &amp; Socket：行情数据对接和抓取.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 Pandas &amp; Numpy：策略与回测系统.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/36%20Pandas%20&amp;%20Numpy%ef%bc%9a%e7%ad%96%e7%95%a5%e4%b8%8e%e5%9b%9e%e6%b5%8b%e7%b3%bb%e7%bb%9f.md">36 Pandas &amp; Numpy：策略与回测系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 Kafka &amp; ZMQ：自动化交易流水线.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/37%20Kafka%20&amp;%20ZMQ%ef%bc%9a%e8%87%aa%e5%8a%a8%e5%8c%96%e4%ba%a4%e6%98%93%e6%b5%81%e6%b0%b4%e7%ba%bf.md">37 Kafka &amp; ZMQ：自动化交易流水线.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 MySQL：日志和数据存储系统.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/38%20MySQL%ef%bc%9a%e6%97%a5%e5%bf%97%e5%92%8c%e6%95%b0%e6%8d%ae%e5%ad%98%e5%82%a8%e7%b3%bb%e7%bb%9f.md">38 MySQL：日志和数据存储系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 Django：搭建监控平台.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/39%20Django%ef%bc%9a%e6%90%ad%e5%bb%ba%e7%9b%91%e6%8e%a7%e5%b9%b3%e5%8f%b0.md">39 Django：搭建监控平台.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 总结：Python中的数据结构与算法全景.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/40%20%e6%80%bb%e7%bb%93%ef%bc%9aPython%e4%b8%ad%e7%9a%84%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84%e4%b8%8e%e7%ae%97%e6%b3%95%e5%85%a8%e6%99%af.md">40 总结：Python中的数据结构与算法全景.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 硅谷一线互联网公司的工作体验.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/41%20%e7%a1%85%e8%b0%b7%e4%b8%80%e7%ba%bf%e4%ba%92%e8%81%94%e7%bd%91%e5%85%ac%e5%8f%b8%e7%9a%84%e5%b7%a5%e4%bd%9c%e4%bd%93%e9%aa%8c.md">41 硅谷一线互联网公司的工作体验.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 细数技术研发的注意事项.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/42%20%e7%bb%86%e6%95%b0%e6%8a%80%e6%9c%af%e7%a0%94%e5%8f%91%e7%9a%84%e6%b3%a8%e6%84%8f%e4%ba%8b%e9%a1%b9.md">42 细数技术研发的注意事项.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 Q&amp;A：聊一聊职业发展和选择.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/43%20Q&amp;A%ef%bc%9a%e8%81%8a%e4%b8%80%e8%81%8a%e8%81%8c%e4%b8%9a%e5%8f%91%e5%b1%95%e5%92%8c%e9%80%89%e6%8b%a9.md">43 Q&amp;A：聊一聊职业发展和选择.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 带你上手SWIG：一份清晰好用的SWIG编程实践指南.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e5%8a%a0%e9%a4%90%20%e5%b8%a6%e4%bd%a0%e4%b8%8a%e6%89%8bSWIG%ef%bc%9a%e4%b8%80%e4%bb%bd%e6%b8%85%e6%99%b0%e5%a5%bd%e7%94%a8%e7%9a%84SWIG%e7%bc%96%e7%a8%8b%e5%ae%9e%e8%b7%b5%e6%8c%87%e5%8d%97.md">加餐 带你上手SWIG：一份清晰好用的SWIG编程实践指南.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 技术之外的几点成长建议.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%8a%80%e6%9c%af%e4%b9%8b%e5%a4%96%e7%9a%84%e5%87%a0%e7%82%b9%e6%88%90%e9%95%bf%e5%bb%ba%e8%ae%ae.md">结束语 技术之外的几点成长建议.md</a>
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
                            <h1 id="title" data-id="20 揭秘 Python 协程" class="title">20 揭秘 Python 协程</h1>
                            <div><p>你好，我是景霄。</p>

<p>上一节课的最后，我们留下一个小小的悬念：生成器在 Python 2 中还扮演了一个重要角色，就是用来实现 Python 协程。</p>

<p>那么首先你要明白，什么是协程？</p>

<p>协程是实现并发编程的一种方式。一说并发，你肯定想到了多线程/多进程模型，没错，多线程/多进程，正是解决并发问题的经典模型之一。最初的互联网世界，多线程/多进程在服务器并发中，起到举足轻重的作用。</p>

<p>随着互联网的快速发展，你逐渐遇到了 C10K 瓶颈，也就是同时连接到服务器的客户达到了一万个。于是很多代码跑崩了，进程上下文切换占用了大量的资源，线程也顶不住如此巨大的压力，这时， NGINX 带着事件循环出来拯救世界了。</p>

<p>如果将多进程/多线程类比为起源于唐朝的藩镇割据，那么事件循环，就是宋朝加强的中央集权制。事件循环启动一个统一的调度器，让调度器来决定一个时刻去运行哪个任务，于是省却了多线程中启动线程、管理线程、同步锁等各种开销。同一时期的 NGINX，在高并发下能保持低资源低消耗高性能，相比 Apache 也支持更多的并发连接。</p>

<p>再到后来，出现了一个很有名的名词，叫做回调地狱（callback hell），手撸过 JavaScript 的朋友肯定知道我在说什么。我们大家惊喜地发现，这种工具完美地继承了事件循环的优越性，同时还能提供 async / await 语法糖，解决了执行性和可读性共存的难题。于是，协程逐渐被更多人发现并看好，也有越来越多的人尝试用 Node.js 做起了后端开发。（讲个笑话，JavaScript 是一门编程语言。）</p>

<p>回到我们的 Python。使用生成器，是 Python 2 开头的时代实现协程的老方法了，Python 3.7 提供了新的基于 asyncio 和 async / await 的方法。我们这节课，同样的，跟随时代，抛弃掉不容易理解、也不容易写的旧的基于生成器的方法，直接来讲新方法。</p>

<p>我们先从一个爬虫实例出发，用清晰的讲解思路，带你结合实战来搞懂这个不算特别容易理解的概念。之后，我们再由浅入深，直击协程的核心。</p>

<h2 id="从一个爬虫说起">从一个爬虫说起</h2>

<p>爬虫，就是互联网的蜘蛛，在搜索引擎诞生之时，与其一同来到世上。爬虫每秒钟都会爬取大量的网页，提取关键信息后存储在数据库中，以便日后分析。爬虫有非常简单的 Python 十行代码实现，也有 Google 那样的全球分布式爬虫的上百万行代码，分布在内部上万台服务器上，对全世界的信息进行嗅探。</p>

<p>话不多说，我们先看一个简单的爬虫例子：</p>

<pre><code>import time

def crawl_page(url):
    print('crawling {}'.format(url))
    sleep_time = int(url.split('_')[-1])
    time.sleep(sleep_time)
    print('OK {}'.format(url))

def main(urls):
    for url in urls:
        crawl_page(url)

%time main(['url_1', 'url_2', 'url_3', 'url_4'])

########## 输出 ##########

crawling url_1
OK url_1
crawling url_2
OK url_2
crawling url_3
OK url_3
crawling url_4
OK url_4
Wall time: 10 s
</code></pre>

<p>（注意：本节的主要目的是协程的基础概念，因此我们简化爬虫的 scrawl_page 函数为休眠数秒，休眠时间取决于 url 最后的那个数字。）</p>

<p>这是一个很简单的爬虫，main() 函数执行时，调取 crawl_page() 函数进行网络通信，经过若干秒等待后收到结果，然后执行下一个。</p>

<p>看起来很简单，但你仔细一算，它也占用了不少时间，五个页面分别用了 1 秒到 4 秒的时间，加起来一共用了 10 秒。这显然效率低下，该怎么优化呢？</p>

<p>于是，一个很简单的思路出现了——我们这种爬取操作，完全可以并发化。我们就来看看使用协程怎么写。</p>

<pre><code>import asyncio

async def crawl_page(url):
    print('crawling {}'.format(url))
    sleep_time = int(url.split('_')[-1])
    await asyncio.sleep(sleep_time)
    print('OK {}'.format(url))

async def main(urls):
    for url in urls:
        await crawl_page(url)

%time asyncio.run(main(['url_1', 'url_2', 'url_3', 'url_4']))

########## 输出 ##########

crawling url_1
OK url_1
crawling url_2
OK url_2
crawling url_3
OK url_3
crawling url_4
OK url_4
Wall time: 10 s
</code></pre>

<p>看到这段代码，你应该发现了，在 Python 3.7 以上版本中，使用协程写异步程序非常简单。</p>

<p>首先来看 import asyncio，这个库包含了大部分我们实现协程所需的魔法工具。</p>

<p>async 修饰词声明异步函数，于是，这里的 crawl_page 和 main 都变成了异步函数。而调用异步函数，我们便可得到一个协程对象（coroutine object）。</p>

<p>举个例子，如果你 <code>print(crawl_page(''))</code>，便会输出<code>&lt;coroutine object crawl_page at 0x000002BEDF141148&gt;</code>，提示你这是一个 Python 的协程对象，而并不会真正执行这个函数。</p>

<p>再来说说协程的执行。执行协程有多种方法，这里我介绍一下常用的三种。</p>

<p>首先，我们可以通过 await 来调用。await 执行的效果，和 Python 正常执行是一样的，也就是说程序会阻塞在这里，进入被调用的协程函数，执行完毕返回后再继续，而这也是 await 的字面意思。代码中 <code>await asyncio.sleep(sleep_time)</code> 会在这里休息若干秒，<code>await crawl_page(url)</code> 则会执行 crawl_page() 函数。</p>

<p>其次，我们可以通过 asyncio.create_task() 来创建任务，这个我们下节课会详细讲一下，你先简单知道即可。</p>

<p>最后，我们需要 asyncio.run 来触发运行。asyncio.run 这个函数是 Python 3.7 之后才有的特性，可以让 Python 的协程接口变得非常简单，你不用去理会事件循环怎么定义和怎么使用的问题（我们会在下面讲）。一个非常好的编程规范是，asyncio.run(main()) 作为主程序的入口函数，在程序运行周期内，只调用一次 asyncio.run。</p>

<p>这样，你就大概看懂了协程是怎么用的吧。不妨试着跑一下代码，欸，怎么还是 10 秒？</p>

<p>10 秒就对了，还记得上面所说的，await 是同步调用，因此， crawl_page(url) 在当前的调用结束之前，是不会触发下一次调用的。于是，这个代码效果就和上面完全一样了，相当于我们用异步接口写了个同步代码。</p>

<p>现在又该怎么办呢？</p>

<p>其实很简单，也正是我接下来要讲的协程中的一个重要概念，任务（Task）。老规矩，先看代码。</p>

<pre><code>import asyncio

async def crawl_page(url):
    print('crawling {}'.format(url))
    sleep_time = int(url.split('_')[-1])
    await asyncio.sleep(sleep_time)
    print('OK {}'.format(url))

async def main(urls):
    tasks = [asyncio.create_task(crawl_page(url)) for url in urls]
    for task in tasks:
        await task

%time asyncio.run(main(['url_1', 'url_2', 'url_3', 'url_4']))

########## 输出 ##########

crawling url_1
crawling url_2
crawling url_3
crawling url_4
OK url_1
OK url_2
OK url_3
OK url_4
Wall time: 3.99 s
</code></pre>

<p>你可以看到，我们有了协程对象后，便可以通过 <code>asyncio.create_task</code> 来创建任务。任务创建后很快就会被调度执行，这样，我们的代码也不会阻塞在任务这里。所以，我们要等所有任务都结束才行，用<code>for task in tasks: await task</code> 即可。</p>

<p>这次，你就看到效果了吧，结果显示，运行总时长等于运行时间最长的爬虫。</p>

<p>当然，你也可以想一想，这里用多线程应该怎么写？而如果需要爬取的页面有上万个又该怎么办呢？再对比下协程的写法，谁更清晰自是一目了然。</p>

<p>其实，对于执行 tasks，还有另一种做法：</p>

<pre><code>import asyncio

async def crawl_page(url):
    print('crawling {}'.format(url))
    sleep_time = int(url.split('_')[-1])
    await asyncio.sleep(sleep_time)
    print('OK {}'.format(url))

async def main(urls):
    tasks = [asyncio.create_task(crawl_page(url)) for url in urls]
    await asyncio.gather(*tasks)

%time asyncio.run(main(['url_1', 'url_2', 'url_3', 'url_4']))

########## 输出 ##########

crawling url_1
crawling url_2
crawling url_3
crawling url_4
OK url_1
OK url_2
OK url_3
OK url_4
Wall time: 4.01 s
</code></pre>

<p>这里的代码也很好理解。唯一要注意的是，<code>*tasks</code> 解包列表，将列表变成了函数的参数；与之对应的是， <code>** dict</code> 将字典变成了函数的参数。</p>

<p>另外，<code>asyncio.create_task</code>，<code>asyncio.run</code> 这些函数都是 Python 3.7 以上的版本才提供的，自然，相比于旧接口它们也更容易理解和阅读。</p>

<h2 id="解密协程运行时">解密协程运行时</h2>

<p>说了这么多，现在，我们不妨来深入代码底层看看。有了前面的知识做基础，你应该很容易理解这两段代码。</p>

<pre><code>import asyncio

async def worker_1():
    print('worker_1 start')
    await asyncio.sleep(1)
    print('worker_1 done')

async def worker_2():
    print('worker_2 start')
    await asyncio.sleep(2)
    print('worker_2 done')

async def main():
    print('before await')
    await worker_1()
    print('awaited worker_1')
    await worker_2()
    print('awaited worker_2')

%time asyncio.run(main())

########## 输出 ##########

before await
worker_1 start
worker_1 done
awaited worker_1
worker_2 start
worker_2 done
awaited worker_2
Wall time: 3 s
</code></pre>

<pre><code>import asyncio

async def worker_1():
    print('worker_1 start')
    await asyncio.sleep(1)
    print('worker_1 done')

async def worker_2():
    print('worker_2 start')
    await asyncio.sleep(2)
    print('worker_2 done')

async def main():
    task1 = asyncio.create_task(worker_1())
    task2 = asyncio.create_task(worker_2())
    print('before await')
    await task1
    print('awaited worker_1')
    await task2
    print('awaited worker_2')

%time asyncio.run(main())

########## 输出 ##########

before await
worker_1 start
worker_2 start
worker_1 done
awaited worker_1
worker_2 done
awaited worker_2
Wall time: 2.01 s
</code></pre>

<p>不过，第二个代码，到底发生了什么呢？为了让你更详细了解到协程和线程的具体区别，这里我详细地分析了整个过程。步骤有点多，别着急，我们慢慢来看。</p>

<ol>
<li><code>asyncio.run(main())</code>，程序进入 main() 函数，事件循环开启；</li>
<li>task1 和 task2 任务被创建，并进入事件循环等待运行；运行到 print，输出 <code>'before await'</code>；</li>
<li>await task1 执行，用户选择从当前的主任务中切出，事件调度器开始调度 worker_1；</li>
<li>worker_1 开始运行，运行 print 输出<code>'worker_1 start'</code>，然后运行到 <code>await asyncio.sleep(1)</code>， 从当前任务切出，事件调度器开始调度 worker_2；</li>
<li>worker_2 开始运行，运行 print 输出 <code>'worker_2 start'</code>，然后运行 <code>await asyncio.sleep(2)</code> 从当前任务切出；</li>
<li>以上所有事件的运行时间，都应该在 1ms 到 10ms 之间，甚至可能更短，事件调度器从这个时候开始暂停调度；</li>
<li>一秒钟后，worker_1 的 sleep 完成，事件调度器将控制权重新传给 task_1，输出 <code>'worker_1 done'</code>，task_1 完成任务，从事件循环中退出；</li>
<li>await task1 完成，事件调度器将控制器传给主任务，输出 <code>'awaited worker_1'</code>，·然后在 await task2 处继续等待；</li>
<li>两秒钟后，worker_2 的 sleep 完成，事件调度器将控制权重新传给 task_2，输出 <code>'worker_2 done'</code>，task_2 完成任务，从事件循环中退出；</li>
<li>主任务输出 <code>'awaited worker_2'</code>，协程全任务结束，事件循环结束。</li>
</ol>

<p>接下来，我们进阶一下。如果我们想给某些协程任务限定运行时间，一旦超时就取消，又该怎么做呢？再进一步，如果某些协程运行时出现错误，又该怎么处理呢？同样的，来看代码。</p>

<pre><code>import asyncio

async def worker_1():
    await asyncio.sleep(1)
    return 1

async def worker_2():
    await asyncio.sleep(2)
    return 2 / 0

async def worker_3():
    await asyncio.sleep(3)
    return 3

async def main():
    task_1 = asyncio.create_task(worker_1())
    task_2 = asyncio.create_task(worker_2())
    task_3 = asyncio.create_task(worker_3())

    await asyncio.sleep(2)
    task_3.cancel()

    res = await asyncio.gather(task_1, task_2, task_3, return_exceptions=True)
    print(res)

%time asyncio.run(main())

########## 输出 ##########

[1, ZeroDivisionError('division by zero'), CancelledError()]
Wall time: 2 s
</code></pre>

<p>你可以看到，worker_1 正常运行，worker_2 运行中出现错误，worker_3 执行时间过长被我们 cancel 掉了，这些信息会全部体现在最终的返回结果 res 中。</p>

<p>不过要注意<code>return_exceptions=True</code>这行代码。如果不设置这个参数，错误就会完整地 throw 到我们这个执行层，从而需要 try except 来捕捉，这也就意味着其他还没被执行的任务会被全部取消掉。为了避免这个局面，我们将 return_exceptions 设置为 True 即可。</p>

<p>到这里，发现了没，线程能实现的，协程都能做到。那就让我们温习一下这些知识点，用协程来实现一个经典的生产者消费者模型吧。</p>

<pre><code>import asyncio
import random

async def consumer(queue, id):
    while True:
        val = await queue.get()
        print('{} get a val: {}'.format(id, val))
        await asyncio.sleep(1)

async def producer(queue, id):
    for i in range(5):
        val = random.randint(1, 10)
        await queue.put(val)
        print('{} put a val: {}'.format(id, val))
        await asyncio.sleep(1)

async def main():
    queue = asyncio.Queue()

    consumer_1 = asyncio.create_task(consumer(queue, 'consumer_1'))
    consumer_2 = asyncio.create_task(consumer(queue, 'consumer_2'))

    producer_1 = asyncio.create_task(producer(queue, 'producer_1'))
    producer_2 = asyncio.create_task(producer(queue, 'producer_2'))

    await asyncio.sleep(10)
    consumer_1.cancel()
    consumer_2.cancel()
    
    await asyncio.gather(consumer_1, consumer_2, producer_1, producer_2, return_exceptions=True)

%time asyncio.run(main())

########## 输出 ##########

producer_1 put a val: 5
producer_2 put a val: 3
consumer_1 get a val: 5
consumer_2 get a val: 3
producer_1 put a val: 1
producer_2 put a val: 3
consumer_2 get a val: 1
consumer_1 get a val: 3
producer_1 put a val: 6
producer_2 put a val: 10
consumer_1 get a val: 6
consumer_2 get a val: 10
producer_1 put a val: 4
producer_2 put a val: 5
consumer_2 get a val: 4
consumer_1 get a val: 5
producer_1 put a val: 2
producer_2 put a val: 8
consumer_1 get a val: 2
consumer_2 get a val: 8
Wall time: 10 s
</code></pre>

<h2 id="实战-豆瓣近日推荐电影爬虫">实战：豆瓣近日推荐电影爬虫</h2>

<p>最后，进入今天的实战环节——实现一个完整的协程爬虫。</p>

<p>任务描述：<a href="https://movie.douban.com/cinema/later/beijing/" target="_blank">https://movie.douban.com/cinema/later/beijing/</a> 这个页面描述了北京最近上映的电影，你能否通过 Python 得到这些电影的名称、上映时间和海报呢？这个页面的海报是缩小版的，我希望你能从具体的电影描述页面中抓取到海报。</p>

<p>听起来难度不是很大吧？我在下面给出了同步版本的代码和协程版本的代码，通过运行时间和代码写法的对比，希望你能对协程有更深的了解。（注意：为了突出重点、简化代码，这里我省略了异常处理。）</p>

<p>不过，在参考我给出的代码之前，你是不是可以自己先动手写一下、跑一下呢？</p>

<pre><code>import requests
from bs4 import BeautifulSoup

def main():
    url = &quot;https://movie.douban.com/cinema/later/beijing/&quot;
    init_page = requests.get(url).content
    init_soup = BeautifulSoup(init_page, 'lxml')

    all_movies = init_soup.find('div', id=&quot;showing-soon&quot;)
    for each_movie in all_movies.find_all('div', class_=&quot;item&quot;):
        all_a_tag = each_movie.find_all('a')
        all_li_tag = each_movie.find_all('li')

        movie_name = all_a_tag[1].text
        url_to_fetch = all_a_tag[1]['href']
        movie_date = all_li_tag[0].text

        response_item = requests.get(url_to_fetch).content
        soup_item = BeautifulSoup(response_item, 'lxml')
        img_tag = soup_item.find('img')

        print('{} {} {}'.format(movie_name, movie_date, img_tag['src']))

%time main()

########## 输出 ##########

阿拉丁 05月24日 https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2553992741.jpg
龙珠超：布罗利 05月24日 https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2557371503.jpg
五月天人生无限公司 05月24日 https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2554324453.jpg
... ...
直播攻略 06月04日 https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2555957974.jpg
Wall time: 56.6 s
</code></pre>

<pre><code>import asyncio
import aiohttp

from bs4 import BeautifulSoup

async def fetch_content(url):
    async with aiohttp.ClientSession(
        headers=header, connector=aiohttp.TCPConnector(ssl=False)
    ) as session:
        async with session.get(url) as response:
            return await response.text()

async def main():
    url = &quot;https://movie.douban.com/cinema/later/beijing/&quot;
    init_page = await fetch_content(url)
    init_soup = BeautifulSoup(init_page, 'lxml')

    movie_names, urls_to_fetch, movie_dates = [], [], []

    all_movies = init_soup.find('div', id=&quot;showing-soon&quot;)
    for each_movie in all_movies.find_all('div', class_=&quot;item&quot;):
        all_a_tag = each_movie.find_all('a')
        all_li_tag = each_movie.find_all('li')

        movie_names.append(all_a_tag[1].text)
        urls_to_fetch.append(all_a_tag[1]['href'])
        movie_dates.append(all_li_tag[0].text)

    tasks = [fetch_content(url) for url in urls_to_fetch]
    pages = await asyncio.gather(*tasks)

    for movie_name, movie_date, page in zip(movie_names, movie_dates, pages):
        soup_item = BeautifulSoup(page, 'lxml')
        img_tag = soup_item.find('img')

        print('{} {} {}'.format(movie_name, movie_date, img_tag['src']))

%time asyncio.run(main())

########## 输出 ##########

阿拉丁 05月24日 https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2553992741.jpg
龙珠超：布罗利 05月24日 https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2557371503.jpg
五月天人生无限公司 05月24日 https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2554324453.jpg
... ...
直播攻略 06月04日 https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2555957974.jpg
Wall time: 4.98 s
</code></pre>

<h2 id="总结">总结</h2>

<p>到这里，今天的主要内容就讲完了。今天我用了较长的篇幅，从一个简单的爬虫开始，到一个真正的爬虫结束，在中间穿插讲解了 Python 协程最新的基本概念和用法。这里带你简单复习一下。</p>

<ul>
<li>协程和多线程的区别，主要在于两点，一是协程为单线程；二是协程由用户决定，在哪些地方交出控制权，切换到下一个任务。</li>
<li>协程的写法更加简洁清晰，把async / await 语法和 create_task 结合来用，对于中小级别的并发需求已经毫无压力。</li>
<li>写协程程序的时候，你的脑海中要有清晰的事件循环概念，知道程序在什么时候需要暂停、等待 I/O，什么时候需要一并执行到底。</li>
</ul>

<p>最后的最后，请一定不要轻易炫技。多线程模型也一定有其优点，一个真正牛逼的程序员，应该懂得，在什么时候用什么模型能达到工程上的最优，而不是自觉某个技术非常牛逼，所有项目创造条件也要上。技术是工程，而工程则是时间、资源、人力等纷繁复杂的事情的折衷。</p>

<h2 id="思考题">思考题</h2>

<p>最后给你留一个思考题。协程怎么实现回调函数呢？欢迎留言和我讨论，也欢迎你把这篇文章分享给你的同事朋友，我们一起交流，一起进步。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#8ee2e2e2b7babfbfbeb9cee9e3efe7e2a0ede1e3" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f11d3522fb7cd35',t:'MTczNDA0OTYxNS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>