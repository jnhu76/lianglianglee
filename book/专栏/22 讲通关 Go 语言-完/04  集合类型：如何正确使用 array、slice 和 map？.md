<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=04&#32;&#32;集合类型：如何正确使用&#32;array、slice&#32;和&#32;map？>
        <link rel="icon" href="/static/favicon.png">
        <title>04  集合类型：如何正确使用 array、slice 和 map？ </title>
        
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
                            <h1 id="title" data-id="04  集合类型：如何正确使用 array、slice 和 map？" class="title">04  集合类型：如何正确使用 array、slice 和 map？</h1>
                            <div><p>上节课的思考题是练习使用 for 循环中的 continue，通过上节课的学习，你已经了解 continue 是跳出本次循环的意思，现在我就以计算 100 以内的偶数之和为例，演示 continue 的用法：</p>

<pre><code class="language-go">sum := 0

for i:=1; i&lt;100; i++{

   if i%2!=0 {

      continue

   }

   sum+=i

}

fmt.Println(&quot;the sum is&quot;,sum)
</code></pre>

<p>这个示例的关键在于：如果 i 不是偶数，就会用 continue 跳出本次循环，继续下个循环；如果是偶数，则继续执行 sum+=i，然后继续循环，这样就达到了只计算 100 以内偶数之和的目的。</p>

<p>下面我们开始本节课的学习，我将介绍 Go 语言的集合类型。</p>

<p>在实际需求中，我们会有很多同一类型的元素放在一起的场景，这就是集合，例如 100 个数字，10 个字符串等。在 Go 语言中，数组（array）、切片（slice）、映射（map）这些都是集合类型，用于存放同一类元素。虽然都是集合，但用处又不太一样，这节课我就为你详细地介绍。</p>

<h3 id="array-数组">Array（数组）</h3>

<p>数组存放的是固定长度、相同类型的数据，而且这些存放的元素是连续的。所存放的数据类型没有限制，可以是整型、字符串甚至自定义。</p>

<h4 id="数组声明">数组声明</h4>

<p>要声明一个数组非常简单，语法和第二课时介绍的声明基础类型是一样的。</p>

<p>在下面的代码示例中，我声明了一个字符串数组，长度是 5，所以其类型定义为 [5]string，其中大括号中的元素用于初始化数组。此外，在类型名前加 [] 中括号，并设置好长度，就可以通过它来推测数组的类型。</p>

<blockquote>
<p>注意：[5]string 和 [4]string 不是同一种类型，也就是说长度也是数组类型的一部分。</p>
</blockquote>

<p><strong><em>ch04/main.go</em></strong></p>

<pre><code class="language-go">array:=[5]string{&quot;a&quot;,&quot;b&quot;,&quot;c&quot;,&quot;d&quot;,&quot;e&quot;}
</code></pre>

<p>数组在内存中都是连续存放的，下面通过一幅图片形象地展示数组在内存中如何存放：</p>

<p><img src="assets/Ciqc1F-pBzmAWUQ0AAAttSjgTjQ158.png" alt="Drawing 1.png" /></p>

<p>可以看到，数组的每个元素都是连续存放的，每一个元素都有一个下标（Index）。下标从 0 开始，比如第一个元素 a 对应的下标是 0，第二个元素 b 对应的下标是 1。以此类推，通过 array+[下标] 的方式，我们可以快速地定位元素。</p>

<p>如下面代码所示，运行它，可以看到输出打印的结果是 c，也就是数组 array 的第三个元素：</p>

<p><strong><em>ch04/main.go</em></strong></p>

<pre><code class="language-go">func main() {

    array:=[5]string{&quot;a&quot;,&quot;b&quot;,&quot;c&quot;,&quot;d&quot;,&quot;e&quot;}

    fmt.Println(array[2])

}
</code></pre>

<p>在定义数组的时候，数组的长度可以省略，这个时候 Go 语言会自动根据大括号 {} 中元素的个数推导出长度，所以以上示例也可以像下面这样声明：</p>

<pre><code class="language-go">array:=[...]string{&quot;a&quot;,&quot;b&quot;,&quot;c&quot;,&quot;d&quot;,&quot;e&quot;}
</code></pre>

<p>以上省略数组长度的声明只适用于所有元素都被初始化的数组，如果是只针对特定索引元素初始化的情况，就不适合了，如下示例：</p>

<pre><code class="language-go">array1:=[5]string{1:&quot;b&quot;,3:&quot;d&quot;}
</code></pre>

<p>示例中的「1:&ldquo;b&rdquo;,3:&ldquo;d&rdquo;」的意思表示初始化索引 1 的值为 b，初始化索引 3 的值为 d，整个数组的长度为 5。如果我省略长度 5，那么整个数组的长度只有 4，显然不符合我们定义数组的初衷。</p>

<p>此外，没有初始化的索引，其默认值都是数组类型的零值，也就是 string 类型的零值 &ldquo;&rdquo; 空字符串。</p>

<p>除了使用 [] 操作符根据索引快速定位数组的元素外，还可以通过 for 循环打印所有的数组元素，如下面的代码所示：</p>

<p><strong><em>ch04/main.go</em></strong></p>

<pre><code class="language-go">for i:=0;i&lt;5;i++{

    fmt.Printf(&quot;数组索引:%d,对应值:%s\n&quot;, i, array[i])

}
</code></pre>

<h4 id="数组循环">数组循环</h4>

<p>使用传统的 for 循环遍历数组，输出对应的索引和对应的值，这种方式很烦琐，一般不使用，大部分情况下，我们使用的是 for range 这种 Go 语言的新型循环，如下面的代码所示：</p>

<pre><code class="language-go">for i,v:=range array{

    fmt.Printf(&quot;数组索引:%d,对应值:%s\n&quot;, i, v)

}
</code></pre>

<p>这种方式和传统 for 循环的结果是一样的。对于数组，range 表达式返回两个结果：</p>

<ol>
<li>第一个是数组的索引；</li>
<li>第二个是数组的值。</li>
</ol>

<p>在上面的示例中，把返回的两个结果分别赋值给 i 和 v 这两个变量，就可以使用它们了。</p>

<p>相比传统的 for 循环，for range 要更简洁，如果返回的值用不到，可以使用 _ 下划线丢弃，如下面的代码所示：</p>

<pre><code class="language-go">for _,v:=range array{

    fmt.Printf(&quot;对应值:%s\n&quot;, v)

}
</code></pre>

<p>数组的索引通过 _ 就被丢弃了，只使用数组的值 v 即可。</p>

<h3 id="slice-切片">Slice（切片）</h3>

<p>切片和数组类似，可以把它理解为动态数组。切片是基于数组实现的，它的底层就是一个数组。对数组任意分隔，就可以得到一个切片。现在我们通过一个例子来更好地理解它，同样还是基于上述例子的 array。</p>

<h4 id="基于数组生成切片">基于数组生成切片</h4>

<p>下面代码中的 array[2:5] 就是获取一个切片的操作，它包含从数组 array 的索引 2 开始到索引 5 结束的元素：</p>

<pre><code class="language-go">array:=[5]string{&quot;a&quot;,&quot;b&quot;,&quot;c&quot;,&quot;d&quot;,&quot;e&quot;}

slice:=array[2:5]

fmt.Println(slice)
</code></pre>

<blockquote>
<p>注意：这里是包含索引 2，但是不包含索引 5 的元素，即在 : 右边的数字不会被包含。</p>
</blockquote>

<p><strong><em>ch04/main.go</em></strong></p>

<pre><code class="language-go">//基于数组生成切片，包含索引start，但是不包含索引end

slice:=array[start:end]
</code></pre>

<p>所以 array[2:5] 获取到的是 c、d、e 这三个元素，然后这三个元素作为一个切片赋值给变量 slice。</p>

<p>切片和数组一样，也可以通过索引定位元素。这里以新获取的 slice 切片为例，slice[0] 的值为 c，slice[1] 的值为 d。</p>

<p>有没有发现，在数组 array 中，元素 c 的索引其实是 2，但是对数组切片后，在新生成的切片 slice 中，它的索引是 0，这就是切片。虽然切片底层用的也是 array 数组，<strong>但是经过切片后，切片的索引范围改变了</strong>。</p>

<p>通过下图可以看出，切片是一个具备三个字段的数据结构，分别是指向数组的指针 data，长度 len 和容量 cap：</p>

<p><img src="assets/Ciqc1F-pC5mAbYE_AABTz22B9TY540.png" alt="图片28.png" /></p>

<p>这里有一些小技巧，切片表达式 array[start:end] 中的 start 和 end 索引都是可以省略的，如果省略 start，那么 start 的值默认为 0，如果省略 end，那么 end 的默认值为数组的长度。如下面的示例：</p>

<ol>
<li>array[:4] 等价于 array[0:4]。</li>
<li>array[1:] 等价于 array[1:5]。</li>
<li>array[:] 等价于 array[0:5]。</li>
</ol>

<h4 id="切片修改">切片修改</h4>

<p>切片的值也可以被修改，这里也同时可以证明切片的底层是数组。</p>

<p>对切片相应的索引元素赋值就是修改，在下面的代码中，把切片 slice 索引 1 的值修改为 f，然后打印输出数组 array：</p>

<pre><code class="language-go">slice:=array[2:5]

slice[1] =&quot;f&quot;

fmt.Println(array)
</code></pre>

<p>可以看到如下结果：</p>

<pre><code class="language-go">[a b c f e]
</code></pre>

<p>数组对应的值已经被修改为 f，所以这也证明了基于数组的切片，使用的底层数组还是原来的数组，一旦修改切片的元素值，那么底层数组对应的值也会被修改。</p>

<h4 id="切片声明"><strong>切片声明</strong></h4>

<p>除了可以从一个数组得到切片外，还可以声明切片，比较简单的是使用 make 函数。</p>

<p>下面的代码是声明了一个元素类型为 string 的切片，长度是 4，make 函数还可以传入一个容量参数：</p>

<pre><code class="language-go">slice1:=make([]string,4)
</code></pre>

<p>在下面的例子中，指定了新创建的切片 []string 容量为 8：</p>

<pre><code class="language-go">slice1:=make([]string,4,8)
</code></pre>

<p>这里需要注意的是，切片的容量不能比切片的长度小。</p>

<p>切片的长度你已经知道了，就是切片内元素的个数。那么容量是什么呢？其实就是切片的空间。</p>

<p>上面的示例说明，Go 语言在内存上划分了一块容量为 8 的内容空间（容量为 8），但是只有 4 个内存空间才有元素（长度为 4），其他的内存空间处于空闲状态，当通过 append 函数往切片中追加元素的时候，会追加到空闲的内存上，当切片的长度要超过容量的时候，会进行扩容。</p>

<p>切片不仅可以通过 make 函数声明，也可以通过字面量的方式声明和初始化，如下所示：</p>

<pre><code class="language-go">slice1:=[]string{&quot;a&quot;,&quot;b&quot;,&quot;c&quot;,&quot;d&quot;,&quot;e&quot;}

fmt.Println(len(slice1),cap(slice1))
</code></pre>

<p>可以注意到，切片和数组的字面量初始化方式，差别就是中括号 [] 里的长度。此外，通过字面量初始化的切片，长度和容量相同。</p>

<h4 id="append">Append</h4>

<p>我们可以通过内置的 append 函数对一个切片追加元素，返回新切片，如下面的代码所示：</p>

<pre><code class="language-go">//追加一个元素

slice2:=append(slice1,&quot;f&quot;)

//多加多个元素

slice2:=append(slice1,&quot;f&quot;,&quot;g&quot;)

//追加另一个切片

slice2:=append(slice1,slice...)
</code></pre>

<p>append 函数可以有以上三种操作，你可以根据自己的实际需求进行选择，append 会自动处理切片容量不足需要扩容的问题。</p>

<blockquote>
<p>小技巧：在创建新切片的时候，最好要让新切片的长度和容量一样，这样在追加操作的时候就会生成新的底层数组，从而和原有数组分离，就不会因为共用底层数组导致修改内容的时候影响多个切片。</p>
</blockquote>

<h4 id="切片元素循环">切片元素循环</h4>

<p>切片的循环和数组一模一样，常用的也是 for range 方式，这里就不再进行举例，当作练习题留给你。</p>

<p>在 Go 语言开发中，切片是使用最多的，尤其是作为函数的参数时，相比数组，通常会优先选择切片，因为它高效，内存占用小。</p>

<h3 id="map-映射">Map（映射）</h3>

<p>在 Go 语言中，map 是一个无序的 K-V 键值对集合，结构为 map[K]V。其中 K 对应 Key，V 对应 Value。map 中所有的 Key 必须具有相同的类型，Value 也同样，但 Key 和 Value 的类型可以不同。此外，Key 的类型必须支持 == 比较运算符，这样才可以判断它是否存在，并保证 Key 的唯一。</p>

<h4 id="map-声明初始化">Map 声明初始化</h4>

<p>创建一个 map 可以通过内置的 make 函数，如下面的代码所示：</p>

<pre><code class="language-go">nameAgeMap:=make(map[string]int)
</code></pre>

<p>它的 Key 类型为 string，Value 类型为 int。有了创建好的 map 变量，就可以对它进行操作了。</p>

<p>在下面的示例中，我添加了一个键值对，Key 为飞雪无情，Value 为 20，如果 Key 已经存在，则更新 Key 对应的 Value：</p>

<pre><code class="language-go">nameAgeMap[&quot;飞雪无情&quot;] = 20
</code></pre>

<p>除了可以通过 make 函数创建 map 外，还可以通过字面量的方式。同样是上面的示例，我们用字面量的方式做如下操作：</p>

<pre><code class="language-go">nameAgeMap:=map[string]int{&quot;飞雪无情&quot;:20}
</code></pre>

<p>在创建 map 的同时添加键值对，如果不想添加键值对，使用空大括号 {} 即可，要注意的是，大括号一定不能省略。</p>

<h4 id="map-获取和删除">Map 获取和删除</h4>

<p>map 的操作和切片、数组差不多，都是通过 [] 操作符，只不过数组切片的 [] 中是索引，而 map 的 [] 中是 Key，如下面的代码所示：</p>

<pre><code class="language-go">//添加键值对或者更新对应 Key 的 Value

nameAgeMap[&quot;飞雪无情&quot;] = 20

//获取指定 Key 对应的 Value

age:=nameAgeMap[&quot;飞雪无情&quot;]
</code></pre>

<p>Go 语言的 map 可以获取不存在的 K-V 键值对，如果 Key 不存在，返回的 Value 是该类型的零值，比如 int 的零值就是 0。所以很多时候，我们需要先判断 map 中的 Key 是否存在。</p>

<p>map 的 [] 操作符可以返回两个值：</p>

<ol>
<li>第一个值是对应的 Value；</li>
<li>第二个值标记该 Key 是否存在，如果存在，它的值为 true。</li>
</ol>

<p>我们通过下面的代码进行演示：</p>

<p><strong><em>ch04/main.go</em></strong></p>

<pre><code class="language-go">nameAgeMap:=make(map[string]int)

nameAgeMap[&quot;飞雪无情&quot;] = 20

age,ok:=nameAgeMap[&quot;飞雪无情1&quot;]

if ok {

    fmt.Println(age)

}
</code></pre>

<p>在示例中，age 是返回的 Value，ok 用来标记该 Key 是否存在，如果存在则打印 age。</p>

<p>如果要删除 map 中的键值对，使用内置的 delete 函数即可，比如要删除 nameAgeMap 中 Key 为飞雪无情的键值对。我们用下面的代码进行演示：</p>

<pre><code class="language-go">delete(nameAgeMap,&quot;飞雪无情&quot;)
</code></pre>

<p>delete 有两个参数：第一个参数是 map，第二个参数是要删除键值对的 Key。</p>

<h4 id="遍历-map">遍历 Map</h4>

<p>map 是一个键值对集合，它同样可以被遍历，在 Go 语言中，map 的遍历使用 for range 循环。</p>

<p>对于 map，for range 返回两个值：</p>

<ol>
<li>第一个是 map 的 Key；</li>
<li>第二个是 map 的 Value。</li>
</ol>

<p>我们用下面的代码进行演示：</p>

<p><strong><em>ch04/main.go</em></strong></p>

<pre><code class="language-go">//测试 for range

nameAgeMap[&quot;飞雪无情&quot;] = 20

nameAgeMap[&quot;飞雪无情1&quot;] = 21

nameAgeMap[&quot;飞雪无情2&quot;] = 22

for k,v:=range nameAgeMap{

    fmt.Println(&quot;Key is&quot;,k,&quot;,Value is&quot;,v)

}
</code></pre>

<p>需要注意的是 map 的遍历是无序的，也就是说你每次遍历，键值对的顺序可能会不一样。如果想按顺序遍历，可以先获取所有的 Key，并对 Key 排序，然后根据排序好的 Key 获取对应的 Value。这里我不再进行演示，你可以当作练习题。</p>

<blockquote>
<p>小技巧：for range map 的时候，也可以使用一个值返回。使用一个返回值的时候，这个返回值默认是 map 的 Key。</p>
</blockquote>

<h4 id="map-的大小">Map 的大小</h4>

<p>和数组切片不一样，map 是没有容量的，它只有长度，也就是 map 的大小（键值对的个数）。要获取 map 的大小，使用内置的 len 函数即可，如下代码所示：</p>

<pre><code class="language-go">fmt.Println(len(nameAgeMap))
</code></pre>

<h3 id="string-和-byte">String 和 []byte</h3>

<p>字符串 string 也是一个不可变的字节序列，所以可以直接转为字节切片 []byte，如下面的代码所示：</p>

<p><strong><em>ch04/main.go</em></strong></p>

<pre><code class="language-go">s:=&quot;Hello飞雪无情&quot;

bs:=[]byte(s)
</code></pre>

<p>string 不止可以直接转为 []byte，还可以使用 [] 操作符获取指定索引的字节值，如以下示例：</p>

<p><strong><em>ch04/main.go</em></strong></p>

<pre><code class="language-go">s:=&quot;Hello飞雪无情&quot;

bs:=[]byte(s)

fmt.Println(bs)

fmt.Println(s[0],s[1],s[15])
</code></pre>

<p>你可能会有疑惑，在这个示例中，字符串 s 里的字母和中文加起来不是 9 个字符吗？怎么可以使用 s[15] 超过 9 的索引呢？其实恰恰就是因为字符串是字节序列，每一个索引对应的是一个字节，而在 UTF8 编码下，一个汉字对应三个字节，所以字符串 s 的长度其实是 17。</p>

<p>运行下面的代码，就可以看到打印的结果是 17。</p>

<pre><code class="language-go">fmt.Println(len(s))
</code></pre>

<p>如果你想把一个汉字当成一个长度计算，可以使用 utf8.RuneCountInString 函数。运行下面的代码，可以看到打印结果是 9，也就是 9 个 unicode（utf8）字符，和我们看到的字符的个数一致。</p>

<pre><code class="language-go">fmt.Println(utf8.RuneCountInString(s))
</code></pre>

<p>而使用 for range 对字符串进行循环时，也恰好是按照 unicode 字符进行循环的，所以对于字符串 s 来说，循环了 9 次。</p>

<p>在下面示例的代码中，i 是索引，r 是 unicode 字符对应的 unicode 码点，这也说明了 for range 循环在处理字符串的时候，会自动地隐式解码 unicode 字符串。</p>

<p><strong><em>ch04/main.go</em></strong></p>

<pre><code class="language-go">for i,r:=range s{

    fmt.Println(i,r)

}
</code></pre>

<h3 id="总结">总结</h3>

<p>这节课到这里就要结束了，在这节课里我讲解了数组、切片和映射的声明和使用，有了这些集合类型，你就可以把你需要的某一类数据放到集合类型中了，比如获取用户列表、商品列表等。</p>

<p>数组、切片还可以分为二维和多维，比如二维字节切片就是 [][]byte，三维就是 [][][]byte，因为不常用，所以本节课中没有详细介绍，你可以结合我讲的一维 []byte 切片自己尝试练习，这也是本节课要给你留的<strong>思考题</strong>，创建一个二维数组并使用它。</p>

<p>此外，如果 map 的 Key 的类型是整型，并且集合中的元素比较少，应该尽量选择切片，因为效率更高。在实际的项目开发中，数组并不常用，尤其是在函数间作为参数传递的时候，用得最多的是切片，它更灵活，并且内存占用少。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#f39f9f9fcac7c2c2c3c4b3949e929a9fdd909c9e" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0ce6914f8794de',t:'MTczMzk5Nzk3NS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>