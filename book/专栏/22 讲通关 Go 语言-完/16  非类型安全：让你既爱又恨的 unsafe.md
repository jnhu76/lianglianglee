<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=16&#32;&#32;非类型安全：让你既爱又恨的&#32;unsafe>
        <link rel="icon" href="/static/favicon.png">
        <title>16  非类型安全：让你既爱又恨的 unsafe </title>
        
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
                        <a class="menu-item" id="00 开篇词  Go 为开发者的需求设计，带你实现高效工作.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20Go%20%e4%b8%ba%e5%bc%80%e5%8f%91%e8%80%85%e7%9a%84%e9%9c%80%e6%b1%82%e8%ae%be%e8%ae%a1%ef%bc%8c%e5%b8%a6%e4%bd%a0%e5%ae%9e%e7%8e%b0%e9%ab%98%e6%95%88%e5%b7%a5%e4%bd%9c.md">00 开篇词  Go 为开发者的需求设计，带你实现高效工作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  基础入门：编写你的第一个 Go 语言程序.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/01%20%20%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8%ef%bc%9a%e7%bc%96%e5%86%99%e4%bd%a0%e7%9a%84%e7%ac%ac%e4%b8%80%e4%b8%aa%20Go%20%e8%af%ad%e8%a8%80%e7%a8%8b%e5%ba%8f.md">01  基础入门：编写你的第一个 Go 语言程序.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  数据类型：你必须掌握的数据类型有哪些？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/02%20%20%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b%ef%bc%9a%e4%bd%a0%e5%bf%85%e9%a1%bb%e6%8e%8c%e6%8f%a1%e7%9a%84%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">02  数据类型：你必须掌握的数据类型有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  控制结构：if、for、switch 逻辑语句的那些事儿.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/03%20%20%e6%8e%a7%e5%88%b6%e7%bb%93%e6%9e%84%ef%bc%9aif%e3%80%81for%e3%80%81switch%20%e9%80%bb%e8%be%91%e8%af%ad%e5%8f%a5%e7%9a%84%e9%82%a3%e4%ba%9b%e4%ba%8b%e5%84%bf.md">03  控制结构：if、for、switch 逻辑语句的那些事儿.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  集合类型：如何正确使用 array、slice 和 map？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/04%20%20%e9%9b%86%e5%90%88%e7%b1%bb%e5%9e%8b%ef%bc%9a%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e4%bd%bf%e7%94%a8%20array%e3%80%81slice%20%e5%92%8c%20map%ef%bc%9f.md">04  集合类型：如何正确使用 array、slice 和 map？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  函数和方法：Go 语言中的函数和方法到底有什么不同？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/05%20%20%e5%87%bd%e6%95%b0%e5%92%8c%e6%96%b9%e6%b3%95%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e4%b8%ad%e7%9a%84%e5%87%bd%e6%95%b0%e5%92%8c%e6%96%b9%e6%b3%95%e5%88%b0%e5%ba%95%e6%9c%89%e4%bb%80%e4%b9%88%e4%b8%8d%e5%90%8c%ef%bc%9f.md">05  函数和方法：Go 语言中的函数和方法到底有什么不同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  struct 和 interface：结构体与接口都实现了哪些功能？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/06%20%20struct%20%e5%92%8c%20interface%ef%bc%9a%e7%bb%93%e6%9e%84%e4%bd%93%e4%b8%8e%e6%8e%a5%e5%8f%a3%e9%83%bd%e5%ae%9e%e7%8e%b0%e4%ba%86%e5%93%aa%e4%ba%9b%e5%8a%9f%e8%83%bd%ef%bc%9f.md">06  struct 和 interface：结构体与接口都实现了哪些功能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  错误处理：如何通过 error、deferred、panic 等处理错误？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/07%20%20%e9%94%99%e8%af%af%e5%a4%84%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%20error%e3%80%81deferred%e3%80%81panic%20%e7%ad%89%e5%a4%84%e7%90%86%e9%94%99%e8%af%af%ef%bc%9f.md">07  错误处理：如何通过 error、deferred、panic 等处理错误？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  并发基础：Goroutines 和 Channels 的声明与使用.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/08%20%20%e5%b9%b6%e5%8f%91%e5%9f%ba%e7%a1%80%ef%bc%9aGoroutines%20%e5%92%8c%20Channels%20%e7%9a%84%e5%a3%b0%e6%98%8e%e4%b8%8e%e4%bd%bf%e7%94%a8.md">08  并发基础：Goroutines 和 Channels 的声明与使用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  同步原语：sync 包让你对并发控制得心应手.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/09%20%20%e5%90%8c%e6%ad%a5%e5%8e%9f%e8%af%ad%ef%bc%9async%20%e5%8c%85%e8%ae%a9%e4%bd%a0%e5%af%b9%e5%b9%b6%e5%8f%91%e6%8e%a7%e5%88%b6%e5%be%97%e5%bf%83%e5%ba%94%e6%89%8b.md">09  同步原语：sync 包让你对并发控制得心应手.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  Context：你必须掌握的多线程并发控制神器.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/10%20%20Context%ef%bc%9a%e4%bd%a0%e5%bf%85%e9%a1%bb%e6%8e%8c%e6%8f%a1%e7%9a%84%e5%a4%9a%e7%ba%bf%e7%a8%8b%e5%b9%b6%e5%8f%91%e6%8e%a7%e5%88%b6%e7%a5%9e%e5%99%a8.md">10  Context：你必须掌握的多线程并发控制神器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  并发模式：Go 语言中即学即用的高效并发模式.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/11%20%20%e5%b9%b6%e5%8f%91%e6%a8%a1%e5%bc%8f%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e4%b8%ad%e5%8d%b3%e5%ad%a6%e5%8d%b3%e7%94%a8%e7%9a%84%e9%ab%98%e6%95%88%e5%b9%b6%e5%8f%91%e6%a8%a1%e5%bc%8f.md">11  并发模式：Go 语言中即学即用的高效并发模式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  指针详解：在什么情况下应该使用指针？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/12%20%20%e6%8c%87%e9%92%88%e8%af%a6%e8%a7%a3%ef%bc%9a%e5%9c%a8%e4%bb%80%e4%b9%88%e6%83%85%e5%86%b5%e4%b8%8b%e5%ba%94%e8%af%a5%e4%bd%bf%e7%94%a8%e6%8c%87%e9%92%88%ef%bc%9f.md">12  指针详解：在什么情况下应该使用指针？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  参数传递：值、引用及指针之间的区别？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/13%20%20%e5%8f%82%e6%95%b0%e4%bc%a0%e9%80%92%ef%bc%9a%e5%80%bc%e3%80%81%e5%bc%95%e7%94%a8%e5%8f%8a%e6%8c%87%e9%92%88%e4%b9%8b%e9%97%b4%e7%9a%84%e5%8c%ba%e5%88%ab%ef%bc%9f.md">13  参数传递：值、引用及指针之间的区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  内存分配：new 还是 make？什么情况下该用谁？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/14%20%20%e5%86%85%e5%ad%98%e5%88%86%e9%85%8d%ef%bc%9anew%20%e8%bf%98%e6%98%af%20make%ef%bc%9f%e4%bb%80%e4%b9%88%e6%83%85%e5%86%b5%e4%b8%8b%e8%af%a5%e7%94%a8%e8%b0%81%ef%bc%9f.md">14  内存分配：new 还是 make？什么情况下该用谁？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  运行时反射：字符串和结构体之间如何转换？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/15%20%20%e8%bf%90%e8%a1%8c%e6%97%b6%e5%8f%8d%e5%b0%84%ef%bc%9a%e5%ad%97%e7%ac%a6%e4%b8%b2%e5%92%8c%e7%bb%93%e6%9e%84%e4%bd%93%e4%b9%8b%e9%97%b4%e5%a6%82%e4%bd%95%e8%bd%ac%e6%8d%a2%ef%bc%9f.md">15  运行时反射：字符串和结构体之间如何转换？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  非类型安全：让你既爱又恨的 unsafe.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/16%20%20%e9%9d%9e%e7%b1%bb%e5%9e%8b%e5%ae%89%e5%85%a8%ef%bc%9a%e8%ae%a9%e4%bd%a0%e6%97%a2%e7%88%b1%e5%8f%88%e6%81%a8%e7%9a%84%20unsafe.md">16  非类型安全：让你既爱又恨的 unsafe.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  SliceHeader：slice 如何高效处理数据？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/17%20%20SliceHeader%ef%bc%9aslice%20%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e5%a4%84%e7%90%86%e6%95%b0%e6%8d%ae%ef%bc%9f.md">17  SliceHeader：slice 如何高效处理数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  质量保证：Go 语言如何通过测试保证质量？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/18%20%20%e8%b4%a8%e9%87%8f%e4%bf%9d%e8%af%81%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e6%b5%8b%e8%af%95%e4%bf%9d%e8%af%81%e8%b4%a8%e9%87%8f%ef%bc%9f.md">18  质量保证：Go 语言如何通过测试保证质量？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  性能优化：Go 语言如何进行代码检查和优化？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/19%20%20%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e4%bb%a3%e7%a0%81%e6%a3%80%e6%9f%a5%e5%92%8c%e4%bc%98%e5%8c%96%ef%bc%9f.md">19  性能优化：Go 语言如何进行代码检查和优化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  协作开发：模块化管理为什么能够提升研发效能？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/20%20%20%e5%8d%8f%e4%bd%9c%e5%bc%80%e5%8f%91%ef%bc%9a%e6%a8%a1%e5%9d%97%e5%8c%96%e7%ae%a1%e7%90%86%e4%b8%ba%e4%bb%80%e4%b9%88%e8%83%bd%e5%a4%9f%e6%8f%90%e5%8d%87%e7%a0%94%e5%8f%91%e6%95%88%e8%83%bd%ef%bc%9f.md">20  协作开发：模块化管理为什么能够提升研发效能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  网络编程：Go 语言如何玩转 RESTful API 服务？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/21%20%20%e7%bd%91%e7%bb%9c%e7%bc%96%e7%a8%8b%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e5%a6%82%e4%bd%95%e7%8e%a9%e8%bd%ac%20RESTful%20API%20%e6%9c%8d%e5%8a%a1%ef%bc%9f.md">21  网络编程：Go 语言如何玩转 RESTful API 服务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  网络编程：Go 语言如何通过 RPC 实现跨平台服务？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/22%20%20%e7%bd%91%e7%bb%9c%e7%bc%96%e7%a8%8b%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%20RPC%20%e5%ae%9e%e7%8e%b0%e8%b7%a8%e5%b9%b3%e5%8f%b0%e6%9c%8d%e5%8a%a1%ef%bc%9f.md">22  网络编程：Go 语言如何通过 RPC 实现跨平台服务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 结束语  你的 Go 语言成长之路.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/23%20%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e4%bd%a0%e7%9a%84%20Go%20%e8%af%ad%e8%a8%80%e6%88%90%e9%95%bf%e4%b9%8b%e8%b7%af.md">23 结束语  你的 Go 语言成长之路.md</a>
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
                            <h1 id="title" data-id="16  非类型安全：让你既爱又恨的 unsafe" class="title">16  非类型安全：让你既爱又恨的 unsafe</h1>
                            <div><p>上节课我留了一个小作业，让你练习一下如何使用反射调用一个方法，下面我来进行讲解。</p>

<p>还是以 person 这个结构体类型为例。我为它增加一个方法 Print，功能是打印一段文本，示例代码如下：</p>

<pre><code>func (p person) Print(prefix string){

   fmt.Printf(&quot;%s:Name is %s,Age is %d\n&quot;,prefix,p.Name,p.Age)

}
</code></pre>

<p>然后就可以通过反射调用 Print 方法了，示例代码如下：</p>

<pre><code>func main() {

   p:=person{Name: &quot;飞雪无情&quot;,Age: 20}

   pv:=reflect.ValueOf(p)

   //反射调用person的Print方法

   mPrint:=pv.MethodByName(&quot;Print&quot;)

   args:=[]reflect.Value{reflect.ValueOf(&quot;登录&quot;)}

   mPrint.Call(args)

}
</code></pre>

<p>从示例中可以看到，要想通过反射调用一个方法，首先要通过 MethodByName 方法找到相应的方法。因为 Print 方法需要参数，所以需要声明参数，它的类型是 []reflect.Value，也就是示例中的 args 变量，最后就可以通过 Call 方法反射调用 Print 方法了。其中记得要把 args 作为参数传递给 Call 方法。</p>

<p>运行以上代码，可以看到如下结果：</p>

<pre><code>登录:Name is 飞雪无情,Age is 20
</code></pre>

<p>从打印的结果可以看到，和我们直接调用 Print 方法是一样的结果，这也证明了通过反射调用 Print 方法是可行的。</p>

<p>下面我们继续深入 Go 的世界，这节课会介绍 Go 语言自带的 unsafe 包的高级用法。</p>

<p>顾名思义，unsafe 是不安全的。Go 将其定义为这个包名，也是为了让我们尽可能地不使用它。不过虽然不安全，它也有优势，那就是可以绕过 Go 的内存安全机制，直接对内存进行读写。所以有时候出于性能需要，还是会冒险使用它来对内存进行操作。</p>

<h3 id="指针类型转换">指针类型转换</h3>

<p>Go 的设计者为了编写方便、提高效率且降低复杂度，将其设计成一门强类型的静态语言。强类型意味着一旦定义了，类型就不能改变；静态意味着类型检查在运行前就做了。同时出于安全考虑，Go 语言是不允许两个指针类型进行转换的。</p>

<p>我们一般使用 *T 作为一个指针类型，表示一个指向类型 T 变量的指针。为了安全的考虑，两个不同的指针类型不能相互转换，比如 *int 不能转为 *float64。</p>

<p>我们来看下面的代码：</p>

<pre><code>func main() {

   i:= 10

   ip:=&amp;i

   var fp *float64 = (*float64)(ip)

   fmt.Println(fp)

}
</code></pre>

<p>这个代码在编译的时候，会提示 *cannot convert ip (type * int) to type * float64*，也就是不能进行强制转型。那如果还是需要转换呢？这就需要使用 unsafe 包里的 Pointer 了。下面我先为你介绍 unsafe.Pointer 是什么，然后再介绍如何转换。</p>

<h3 id="unsafe-pointer">unsafe.Pointer</h3>

<p>unsafe.Pointer 是一种特殊意义的指针，可以表示任意类型的地址，类似 C 语言里的 void* 指针，是全能型的。</p>

<p>正常情况下，*int 无法转换为 *float64，但是通过 unsafe.Pointer 做中转就可以了。在下面的示例中，我通过 unsafe.Pointer 把 *int 转换为 *float64，并且对新的 *float64 进行 3 倍的乘法操作，你会发现原来变量 i 的值也被改变了，变为 30。</p>

<p><strong><em>ch16/main.go</em></strong></p>

<pre><code>func main() {

   i:= 10

   ip:=&amp;i

   var fp *float64 = (*float64)(unsafe.Pointer(ip))

   *fp = *fp * 3

   fmt.Println(i)

}
</code></pre>

<p>这个例子没有任何实际意义，但是说明了通过 unsafe.Pointer 这个万能的指针，我们可以在 *T 之间做任何转换。那么 unsafe.Pointer 到底是什么？为什么其他类型的指针可以转换为 unsafe.Pointer 呢？这就要看 unsafe.Pointer 的源代码定义了，如下所示：</p>

<pre><code>// ArbitraryType is here for the purposes of documentation

// only and is not actually part of the unsafe package. 

// It represents the type of an arbitrary Go expression.

type ArbitraryType int

type Pointer *ArbitraryType
</code></pre>

<p>按 Go 语言官方的注释，ArbitraryType 可以表示任何类型（这里的 ArbitraryType 仅仅是文档需要，不用太关注它本身，只要记住可以表示任何类型即可）。 而 unsafe.Pointer 又是 *ArbitraryType，也就是说 unsafe.Pointer 是任何类型的指针，也就是一个通用型的指针，足以表示任何内存地址。</p>

<h3 id="uintptr-指针类型">uintptr 指针类型</h3>

<p>uintptr 也是一种指针类型，它足够大，可以表示任何指针。它的类型定义如下所示：</p>

<pre><code>// uintptr is an integer type that is large enough 

// to hold the bit pattern of any pointer.

type uintptr uintptr
</code></pre>

<p>既然已经有了 unsafe.Pointer，为什么还要设计 uintptr 类型呢？这是因为 unsafe.Pointer 不能进行运算，比如不支持 +（加号）运算符操作，但是 uintptr 可以。通过它，可以对指针偏移进行计算，这样就可以访问特定的内存，达到对特定内存读写的目的，这是真正内存级别的操作。</p>

<p>在下面的代码中，我以通过指针偏移修改 struct 结构体内的字段为例，演示 uintptr 的用法。</p>

<pre><code>func main() {

   p :=new(person)

   //Name是person的第一个字段不用偏移，即可通过指针修改

   pName:=(*string)(unsafe.Pointer(p))

   *pName=&quot;飞雪无情&quot;

   //Age并不是person的第一个字段，所以需要进行偏移，这样才能正确定位到Age字段这块内存，才可以正确的修改

   pAge:=(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p))+unsafe.Offsetof(p.Age)))

   *pAge = 20

   fmt.Println(*p)

}

type person struct {

   Name string

   Age int

}
</code></pre>

<p>这个示例不是通过直接访问相应字段的方式对 person 结构体字段赋值，而是通过指针偏移找到相应的内存，然后对内存操作进行赋值。</p>

<p>下面我详细介绍操作步骤。</p>

<ol>
<li>先使用 new 函数声明一个 *person 类型的指针变量 p。</li>
<li>然后把 *person 类型的指针变量 p 通过 unsafe.Pointer，转换为 *string 类型的指针变量 pName。</li>
<li>因为 person 这个结构体的第一个字段就是 string 类型的 Name，所以 pName 这个指针就指向 Name 字段（偏移为 0），对 pName 进行修改其实就是修改字段 Name 的值。</li>
<li>因为 Age 字段不是 person 的第一个字段，要修改它必须要进行指针偏移运算。所以需要先把指针变量 p 通过 unsafe.Pointer 转换为 uintptr，这样才能进行地址运算。既然要进行指针偏移，那么要偏移多少呢？这个偏移量可以通过函数 unsafe.Offsetof 计算出来，该函数返回的是一个 uintptr 类型的偏移量，有了这个偏移量就可以通过 + 号运算符获得正确的 Age 字段的内存地址了，也就是通过 unsafe.Pointer 转换后的 *int 类型的指针变量 pAge。</li>
<li>然后需要注意的是，如果要进行指针运算，要先通过 unsafe.Pointer 转换为 uintptr 类型的指针。指针运算完毕后，还要通过 unsafe.Pointer 转换为真实的指针类型（比如示例中的 *int 类型），这样可以对这块内存进行赋值或取值操作。</li>
<li>有了指向字段 Age 的指针变量 pAge，就可以对其进行赋值操作，修改字段 Age 的值了。</li>
</ol>

<p>运行以上示例，你可以看到如下结果：</p>

<pre><code>{飞雪无情 20}
</code></pre>

<p>这个示例主要是为了讲解 uintptr 指针运算，所以一个结构体字段的赋值才会写得这么复杂，如果按照正常的编码，以上示例代码会和下面的代码结果一样。</p>

<pre><code>func main() {

   p :=new(person)

   p.Name = &quot;飞雪无情&quot;

   p.Age = 20

   fmt.Println(*p)

}
</code></pre>

<p>指针运算的核心在于它操作的是一个个内存地址，通过内存地址的增减，就可以指向一块块不同的内存并对其进行操作，而且不必知道这块内存被起了什么名字（变量名）。</p>

<h3 id="指针转换规则">指针转换规则</h3>

<p>你已经知道 Go 语言中存在三种类型的指针，它们分别是：常用的 *T、unsafe.Pointer 及 uintptr。通过以上示例讲解，可以总结出这三者的转换规则：</p>

<ol>
<li>任何类型的 *T 都可以转换为 unsafe.Pointer；</li>
<li>unsafe.Pointer 也可以转换为任何类型的 *T；</li>
<li>unsafe.Pointer 可以转换为 uintptr；</li>
<li>uintptr 也可以转换为 unsafe.Pointer。</li>
</ol>

<p><img src="assets/CgpVE1_ge3eAaiW8AABP2TB3gXA825.png" alt="图片1.png" /></p>

<p>(指针转换示意图)</p>

<p>可以发现，unsafe.Pointer 主要用于指针类型的转换，而且是各个指针类型转换的桥梁。uintptr 主要用于指针运算，尤其是通过偏移量定位不同的内存。</p>

<h3 id="unsafe-sizeof">unsafe.Sizeof</h3>

<p>Sizeof 函数可以返回一个类型所占用的内存大小，这个大小只与类型有关，和类型对应的变量存储的内容大小无关，比如 bool 型占用一个字节、int8 也占用一个字节。</p>

<p>通过 Sizeof 函数你可以查看任何类型（比如字符串、切片、整型）占用的内存大小，示例代码如下：</p>

<pre><code>fmt.Println(unsafe.Sizeof(true))

fmt.Println(unsafe.Sizeof(int8(0)))

fmt.Println(unsafe.Sizeof(int16(10)))

fmt.Println(unsafe.Sizeof(int32(10000000)))

fmt.Println(unsafe.Sizeof(int64(10000000000000)))

fmt.Println(unsafe.Sizeof(int(10000000000000000)))

fmt.Println(unsafe.Sizeof(string(&quot;飞雪无情&quot;)))

fmt.Println(unsafe.Sizeof([]string{&quot;飞雪u无情&quot;,&quot;张三&quot;}))
</code></pre>

<p>对于整型来说，占用的字节数意味着这个类型存储数字范围的大小，比如 int8 占用一个字节，也就是 8bit，所以它可以存储的大小范围是 -128~~127，也就是 −2^(n-1) 到 2^(n-1)−1。其中 n 表示 bit，int8 表示 8bit，int16 表示 16bit，以此类推。</p>

<p>对于和平台有关的 int 类型，要看平台是 32 位还是 64 位，会取最大的。比如我自己测试以上输出，会发现 int 和 int64 的大小是一样的，因为我用的是 64 位平台的电脑。</p>

<blockquote>
<p>小提示：一个 struct 结构体的内存占用大小，等于它包含的字段类型内存占用大小之和。</p>
</blockquote>

<h3 id="总结">总结</h3>

<p>unsafe 包里最常用的就是 Pointer 指针，通过它可以让你在 *T、uintptr 及 Pointer 三者间转换，从而实现自己的需求，比如零内存拷贝或通过 uintptr 进行指针运算，这些都可以提高程序效率。</p>

<p>unsafe 包里的功能虽然不安全，但的确很香，比如指针运算、类型转换等，都可以帮助我们提高性能。不过我还是建议尽可能地不使用，因为它可以绕开 Go 语言编译器的检查，可能会因为你的操作失误而出现问题。当然如果是需要提高性能的必要操作，还是可以使用，比如 []byte 转 string，就可以通过 unsafe.Pointer 实现零内存拷贝，下节课我会详细讲解。
<img src="assets/Ciqc1F_ge7KAW3QcAAVodV4QU6c331.png" alt="16金句.png" />
unsafe 包还有一个函数我这节课没有讲，它是 Alignof，功能就是函数名字字面的意思，比较简单，你可以自己练习使用一下，这也是这节课的思考题。记得来听下节课哦！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#ea868686d3dedbdbdaddaa8d878b8386c4898587" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0ce8c0695094de',t:'MTczMzk5ODA2NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>