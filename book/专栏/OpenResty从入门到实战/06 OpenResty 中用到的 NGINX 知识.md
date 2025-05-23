<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=06&#32;OpenResty&#32;中用到的&#32;NGINX&#32;知识>
        <link rel="icon" href="/static/favicon.png">
        <title>06 OpenResty 中用到的 NGINX 知识 </title>
        
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
                        <a class="menu-item" id="00 开篇词 OpenResty，为你打开高性能开发的大门.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/00%20%e5%bc%80%e7%af%87%e8%af%8d%20OpenResty%ef%bc%8c%e4%b8%ba%e4%bd%a0%e6%89%93%e5%bc%80%e9%ab%98%e6%80%a7%e8%83%bd%e5%bc%80%e5%8f%91%e7%9a%84%e5%a4%a7%e9%97%a8.md">00 开篇词 OpenResty，为你打开高性能开发的大门.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 初探OpenResty的三大特性.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/01%20%e5%88%9d%e6%8e%a2OpenResty%e7%9a%84%e4%b8%89%e5%a4%a7%e7%89%b9%e6%80%a7.md">01 初探OpenResty的三大特性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 如何写出你的“hello world”？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/02%20%e5%a6%82%e4%bd%95%e5%86%99%e5%87%ba%e4%bd%a0%e7%9a%84%e2%80%9chello%20world%e2%80%9d%ef%bc%9f.md">02 如何写出你的“hello world”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 揪出隐藏在背后的那些子项目.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/03%20%e6%8f%aa%e5%87%ba%e9%9a%90%e8%97%8f%e5%9c%a8%e8%83%8c%e5%90%8e%e7%9a%84%e9%82%a3%e4%ba%9b%e5%ad%90%e9%a1%b9%e7%9b%ae.md">03 揪出隐藏在背后的那些子项目.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 如何管理第三方包？从包管理工具luarocks和opm说起.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/04%20%e5%a6%82%e4%bd%95%e7%ae%a1%e7%90%86%e7%ac%ac%e4%b8%89%e6%96%b9%e5%8c%85%ef%bc%9f%e4%bb%8e%e5%8c%85%e7%ae%a1%e7%90%86%e5%b7%a5%e5%85%b7luarocks%e5%92%8copm%e8%af%b4%e8%b5%b7.md">04 如何管理第三方包？从包管理工具luarocks和opm说起.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 [视频]opm项目导读.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/05%20[%e8%a7%86%e9%a2%91]opm%e9%a1%b9%e7%9b%ae%e5%af%bc%e8%af%bb.md">05 [视频]opm项目导读.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 OpenResty 中用到的 NGINX 知识.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/06%20OpenResty%20%e4%b8%ad%e7%94%a8%e5%88%b0%e7%9a%84%20NGINX%20%e7%9f%a5%e8%af%86.md">06 OpenResty 中用到的 NGINX 知识.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 带你快速上手 Lua.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/07%20%e5%b8%a6%e4%bd%a0%e5%bf%ab%e9%80%9f%e4%b8%8a%e6%89%8b%20Lua.md">07 带你快速上手 Lua.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 LuaJIT分支和标准Lua有什么不同？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/08%20LuaJIT%e5%88%86%e6%94%af%e5%92%8c%e6%a0%87%e5%87%86Lua%e6%9c%89%e4%bb%80%e4%b9%88%e4%b8%8d%e5%90%8c%ef%bc%9f.md">08 LuaJIT分支和标准Lua有什么不同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 为什么 lua-resty-core 性能更高一些？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/09%20%e4%b8%ba%e4%bb%80%e4%b9%88%20lua-resty-core%20%e6%80%a7%e8%83%bd%e6%9b%b4%e9%ab%98%e4%b8%80%e4%ba%9b%ef%bc%9f.md">09 为什么 lua-resty-core 性能更高一些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 JIT编译器的死穴：为什么要避免使用 NYI ？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/10%20JIT%e7%bc%96%e8%af%91%e5%99%a8%e7%9a%84%e6%ad%bb%e7%a9%b4%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e9%81%bf%e5%85%8d%e4%bd%bf%e7%94%a8%20NYI%20%ef%bc%9f.md">10 JIT编译器的死穴：为什么要避免使用 NYI ？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 剖析Lua唯一的数据结构table和metatable特性.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/11%20%e5%89%96%e6%9e%90Lua%e5%94%af%e4%b8%80%e7%9a%84%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84table%e5%92%8cmetatable%e7%89%b9%e6%80%a7.md">11 剖析Lua唯一的数据结构table和metatable特性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 高手秘诀：识别Lua的独有概念和坑.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/12%20%e9%ab%98%e6%89%8b%e7%a7%98%e8%af%80%ef%bc%9a%e8%af%86%e5%88%abLua%e7%9a%84%e7%8b%ac%e6%9c%89%e6%a6%82%e5%bf%b5%e5%92%8c%e5%9d%91.md">12 高手秘诀：识别Lua的独有概念和坑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 [视频]实战：基于FFI实现的lua-resty-lrucache.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/13%20[%e8%a7%86%e9%a2%91]%e5%ae%9e%e6%88%98%ef%bc%9a%e5%9f%ba%e4%ba%8eFFI%e5%ae%9e%e7%8e%b0%e7%9a%84lua-resty-lrucache.md">13 [视频]实战：基于FFI实现的lua-resty-lrucache.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 答疑（一）：Lua 规则和 NGINX 配置文件产生冲突怎么办？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/14%20%e7%ad%94%e7%96%91%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9aLua%20%e8%a7%84%e5%88%99%e5%92%8c%20NGINX%20%e9%85%8d%e7%bd%ae%e6%96%87%e4%bb%b6%e4%ba%a7%e7%94%9f%e5%86%b2%e7%aa%81%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">14 答疑（一）：Lua 规则和 NGINX 配置文件产生冲突怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 OpenResty 和别的开发平台有什么不同？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/15%20OpenResty%20%e5%92%8c%e5%88%ab%e7%9a%84%e5%bc%80%e5%8f%91%e5%b9%b3%e5%8f%b0%e6%9c%89%e4%bb%80%e4%b9%88%e4%b8%8d%e5%90%8c%ef%bc%9f.md">15 OpenResty 和别的开发平台有什么不同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 秒杀大多数开发问题的两个利器：文档和测试案例.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/16%20%e7%a7%92%e6%9d%80%e5%a4%a7%e5%a4%9a%e6%95%b0%e5%bc%80%e5%8f%91%e9%97%ae%e9%a2%98%e7%9a%84%e4%b8%a4%e4%b8%aa%e5%88%a9%e5%99%a8%ef%bc%9a%e6%96%87%e6%a1%a3%e5%92%8c%e6%b5%8b%e8%af%95%e6%a1%88%e4%be%8b.md">16 秒杀大多数开发问题的两个利器：文档和测试案例.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 为什么能成为更好的Web服务器？动态处理请求和响应是关键.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/17%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%83%bd%e6%88%90%e4%b8%ba%e6%9b%b4%e5%a5%bd%e7%9a%84Web%e6%9c%8d%e5%8a%a1%e5%99%a8%ef%bc%9f%e5%8a%a8%e6%80%81%e5%a4%84%e7%90%86%e8%af%b7%e6%b1%82%e5%92%8c%e5%93%8d%e5%ba%94%e6%98%af%e5%85%b3%e9%94%ae.md">17 为什么能成为更好的Web服务器？动态处理请求和响应是关键.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 worker间的通信法宝：最重要的数据结构之shared dict.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/18%20worker%e9%97%b4%e7%9a%84%e9%80%9a%e4%bf%a1%e6%b3%95%e5%ae%9d%ef%bc%9a%e6%9c%80%e9%87%8d%e8%a6%81%e7%9a%84%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84%e4%b9%8bshared%20dict.md">18 worker间的通信法宝：最重要的数据结构之shared dict.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 OpenResty 的核心和精髓：cosocket.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/19%20OpenResty%20%e7%9a%84%e6%a0%b8%e5%bf%83%e5%92%8c%e7%b2%be%e9%ab%93%ef%bc%9acosocket.md">19 OpenResty 的核心和精髓：cosocket.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 超越 Web 服务器：特权进程和定时任务.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/20%20%e8%b6%85%e8%b6%8a%20Web%20%e6%9c%8d%e5%8a%a1%e5%99%a8%ef%bc%9a%e7%89%b9%e6%9d%83%e8%bf%9b%e7%a8%8b%e5%92%8c%e5%ae%9a%e6%97%b6%e4%bb%bb%e5%8a%a1.md">20 超越 Web 服务器：特权进程和定时任务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 带你玩转时间、正则表达式等常用API.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/21%20%e5%b8%a6%e4%bd%a0%e7%8e%a9%e8%bd%ac%e6%97%b6%e9%97%b4%e3%80%81%e6%ad%a3%e5%88%99%e8%a1%a8%e8%be%be%e5%bc%8f%e7%ad%89%e5%b8%b8%e7%94%a8API.md">21 带你玩转时间、正则表达式等常用API.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 [视频]从一个安全漏洞说起，探寻API性能和安全的平衡.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/22%20[%e8%a7%86%e9%a2%91]%e4%bb%8e%e4%b8%80%e4%b8%aa%e5%ae%89%e5%85%a8%e6%bc%8f%e6%b4%9e%e8%af%b4%e8%b5%b7%ef%bc%8c%e6%8e%a2%e5%af%bbAPI%e6%80%a7%e8%83%bd%e5%92%8c%e5%ae%89%e5%85%a8%e7%9a%84%e5%b9%b3%e8%a1%a1.md">22 [视频]从一个安全漏洞说起，探寻API性能和安全的平衡.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 [视频]导读lua-resty-requests：优秀的lua-resty-_是如何编写的？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/23%20[%e8%a7%86%e9%a2%91]%e5%af%bc%e8%af%bblua-resty-requests%ef%bc%9a%e4%bc%98%e7%a7%80%e7%9a%84lua-resty-_%e6%98%af%e5%a6%82%e4%bd%95%e7%bc%96%e5%86%99%e7%9a%84%ef%bc%9f.md">23 [视频]导读lua-resty-requests：优秀的lua-resty-_是如何编写的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 实战：处理四层流量，实现Memcached Server.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/24%20%e5%ae%9e%e6%88%98%ef%bc%9a%e5%a4%84%e7%90%86%e5%9b%9b%e5%b1%82%e6%b5%81%e9%87%8f%ef%bc%8c%e5%ae%9e%e7%8e%b0Memcached%20Server.md">24 实战：处理四层流量，实现Memcached Server.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 答疑（二）：特权进程的权限到底是什么？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/25%20%e7%ad%94%e7%96%91%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e7%89%b9%e6%9d%83%e8%bf%9b%e7%a8%8b%e7%9a%84%e6%9d%83%e9%99%90%e5%88%b0%e5%ba%95%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">25 答疑（二）：特权进程的权限到底是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 代码贡献者的拦路虎：test__nginx 简介.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/26%20%e4%bb%a3%e7%a0%81%e8%b4%a1%e7%8c%ae%e8%80%85%e7%9a%84%e6%8b%a6%e8%b7%af%e8%99%8e%ef%bc%9atest__nginx%20%e7%ae%80%e4%bb%8b.md">26 代码贡献者的拦路虎：test__nginx 简介.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 test__nginx 包罗万象的测试方法.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/27%20test__nginx%20%e5%8c%85%e7%bd%97%e4%b8%87%e8%b1%a1%e7%9a%84%e6%b5%8b%e8%af%95%e6%96%b9%e6%b3%95.md">27 test__nginx 包罗万象的测试方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 test__nginx 还可以这样用？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/28%20test__nginx%20%e8%bf%98%e5%8f%af%e4%bb%a5%e8%bf%99%e6%a0%b7%e7%94%a8%ef%bc%9f.md">28 test__nginx 还可以这样用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 最容易失准的性能测试？你需要压测工具界的“悍马”wrk.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/29%20%e6%9c%80%e5%ae%b9%e6%98%93%e5%a4%b1%e5%87%86%e7%9a%84%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%ef%bc%9f%e4%bd%a0%e9%9c%80%e8%a6%81%e5%8e%8b%e6%b5%8b%e5%b7%a5%e5%85%b7%e7%95%8c%e7%9a%84%e2%80%9c%e6%82%8d%e9%a9%ac%e2%80%9dwrk.md">29 最容易失准的性能测试？你需要压测工具界的“悍马”wrk.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 答疑（三）如何搭建测试的网络结构？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/30%20%e7%ad%94%e7%96%91%ef%bc%88%e4%b8%89%ef%bc%89%e5%a6%82%e4%bd%95%e6%90%ad%e5%bb%ba%e6%b5%8b%e8%af%95%e7%9a%84%e7%bd%91%e7%bb%9c%e7%bb%93%e6%9e%84%ef%bc%9f.md">30 答疑（三）如何搭建测试的网络结构？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 性能下降10倍的真凶：阻塞函数.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/31%20%e6%80%a7%e8%83%bd%e4%b8%8b%e9%99%8d10%e5%80%8d%e7%9a%84%e7%9c%9f%e5%87%b6%ef%bc%9a%e9%98%bb%e5%a1%9e%e5%87%bd%e6%95%b0.md">31 性能下降10倍的真凶：阻塞函数.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 让人又恨又爱的字符串操作.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/32%20%e8%ae%a9%e4%ba%ba%e5%8f%88%e6%81%a8%e5%8f%88%e7%88%b1%e7%9a%84%e5%ad%97%e7%ac%a6%e4%b8%b2%e6%93%8d%e4%bd%9c.md">32 让人又恨又爱的字符串操作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 性能提升10倍的秘诀：必须用好 table.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/33%20%e6%80%a7%e8%83%bd%e6%8f%90%e5%8d%8710%e5%80%8d%e7%9a%84%e7%a7%98%e8%af%80%ef%bc%9a%e5%bf%85%e9%a1%bb%e7%94%a8%e5%a5%bd%20table.md">33 性能提升10倍的秘诀：必须用好 table.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 特别放送：OpenResty编码指南.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/34%20%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%ef%bc%9aOpenResty%e7%bc%96%e7%a0%81%e6%8c%87%e5%8d%97.md">34 特别放送：OpenResty编码指南.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 [视频]实际项目中的性能优化：ingress-nginx中的几个PR解读.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/35%20[%e8%a7%86%e9%a2%91]%e5%ae%9e%e9%99%85%e9%a1%b9%e7%9b%ae%e4%b8%ad%e7%9a%84%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%ef%bc%9aingress-nginx%e4%b8%ad%e7%9a%84%e5%87%a0%e4%b8%aaPR%e8%a7%a3%e8%af%bb.md">35 [视频]实际项目中的性能优化：ingress-nginx中的几个PR解读.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 盘点OpenResty的各种调试手段.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/36%20%e7%9b%98%e7%82%b9OpenResty%e7%9a%84%e5%90%84%e7%a7%8d%e8%b0%83%e8%af%95%e6%89%8b%e6%ae%b5.md">36 盘点OpenResty的各种调试手段.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 systemtap-toolkit和stapxx：如何用数据搞定“疑难杂症”？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/37%20systemtap-toolkit%e5%92%8cstapxx%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a8%e6%95%b0%e6%8d%ae%e6%90%9e%e5%ae%9a%e2%80%9c%e7%96%91%e9%9a%be%e6%9d%82%e7%97%87%e2%80%9d%ef%bc%9f.md">37 systemtap-toolkit和stapxx：如何用数据搞定“疑难杂症”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 [视频]巧用wrk和火焰图，科学定位性能瓶颈.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/38%20[%e8%a7%86%e9%a2%91]%e5%b7%a7%e7%94%a8wrk%e5%92%8c%e7%81%ab%e7%84%b0%e5%9b%be%ef%bc%8c%e7%a7%91%e5%ad%a6%e5%ae%9a%e4%bd%8d%e6%80%a7%e8%83%bd%e7%93%b6%e9%a2%88.md">38 [视频]巧用wrk和火焰图，科学定位性能瓶颈.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 高性能的关键：shared dict 缓存和 lru 缓存.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/39%20%e9%ab%98%e6%80%a7%e8%83%bd%e7%9a%84%e5%85%b3%e9%94%ae%ef%bc%9ashared%20dict%20%e7%bc%93%e5%ad%98%e5%92%8c%20lru%20%e7%bc%93%e5%ad%98.md">39 高性能的关键：shared dict 缓存和 lru 缓存.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 缓存与风暴并存，谁说缓存风暴不可避免？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/40%20%e7%bc%93%e5%ad%98%e4%b8%8e%e9%a3%8e%e6%9a%b4%e5%b9%b6%e5%ad%98%ef%bc%8c%e8%b0%81%e8%af%b4%e7%bc%93%e5%ad%98%e9%a3%8e%e6%9a%b4%e4%b8%8d%e5%8f%af%e9%81%bf%e5%85%8d%ef%bc%9f.md">40 缓存与风暴并存，谁说缓存风暴不可避免？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 lua-resty-_ 封装，让你远离多级缓存之痛.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/41%20lua-resty-_%20%e5%b0%81%e8%a3%85%ef%bc%8c%e8%ae%a9%e4%bd%a0%e8%bf%9c%e7%a6%bb%e5%a4%9a%e7%ba%a7%e7%bc%93%e5%ad%98%e4%b9%8b%e7%97%9b.md">41 lua-resty-_ 封装，让你远离多级缓存之痛.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 如何应对突发流量：漏桶和令牌桶的概念.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/42%20%e5%a6%82%e4%bd%95%e5%ba%94%e5%af%b9%e7%aa%81%e5%8f%91%e6%b5%81%e9%87%8f%ef%bc%9a%e6%bc%8f%e6%a1%b6%e5%92%8c%e4%bb%a4%e7%89%8c%e6%a1%b6%e7%9a%84%e6%a6%82%e5%bf%b5.md">42 如何应对突发流量：漏桶和令牌桶的概念.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 灵活实现动态限流限速，其实没有那么难.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/43%20%e7%81%b5%e6%b4%bb%e5%ae%9e%e7%8e%b0%e5%8a%a8%e6%80%81%e9%99%90%e6%b5%81%e9%99%90%e9%80%9f%ef%bc%8c%e5%85%b6%e5%ae%9e%e6%b2%a1%e6%9c%89%e9%82%a3%e4%b9%88%e9%9a%be.md">43 灵活实现动态限流限速，其实没有那么难.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 OpenResty 的杀手锏：动态.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/44%20OpenResty%20%e7%9a%84%e6%9d%80%e6%89%8b%e9%94%8f%ef%bc%9a%e5%8a%a8%e6%80%81.md">44 OpenResty 的杀手锏：动态.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 不得不提的能力外延：OpenResty常用的第三方库.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/45%20%e4%b8%8d%e5%be%97%e4%b8%8d%e6%8f%90%e7%9a%84%e8%83%bd%e5%8a%9b%e5%a4%96%e5%bb%b6%ef%bc%9aOpenResty%e5%b8%b8%e7%94%a8%e7%9a%84%e7%ac%ac%e4%b8%89%e6%96%b9%e5%ba%93.md">45 不得不提的能力外延：OpenResty常用的第三方库.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46 答疑（四）：共享字典的缓存是必须的吗？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/46%20%e7%ad%94%e7%96%91%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9a%e5%85%b1%e4%ba%ab%e5%ad%97%e5%85%b8%e7%9a%84%e7%bc%93%e5%ad%98%e6%98%af%e5%bf%85%e9%a1%bb%e7%9a%84%e5%90%97%ef%bc%9f.md">46 答疑（四）：共享字典的缓存是必须的吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="47 微服务API网关搭建三步曲（一）.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/47%20%e5%be%ae%e6%9c%8d%e5%8a%a1API%e7%bd%91%e5%85%b3%e6%90%ad%e5%bb%ba%e4%b8%89%e6%ad%a5%e6%9b%b2%ef%bc%88%e4%b8%80%ef%bc%89.md">47 微服务API网关搭建三步曲（一）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="48 微服务API网关搭建三步曲（二）.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/48%20%e5%be%ae%e6%9c%8d%e5%8a%a1API%e7%bd%91%e5%85%b3%e6%90%ad%e5%bb%ba%e4%b8%89%e6%ad%a5%e6%9b%b2%ef%bc%88%e4%ba%8c%ef%bc%89.md">48 微服务API网关搭建三步曲（二）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="49 微服务API网关搭建三步曲（三）.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/49%20%e5%be%ae%e6%9c%8d%e5%8a%a1API%e7%bd%91%e5%85%b3%e6%90%ad%e5%bb%ba%e4%b8%89%e6%ad%a5%e6%9b%b2%ef%bc%88%e4%b8%89%ef%bc%89.md">49 微服务API网关搭建三步曲（三）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="50 答疑（五）：如何在工作中引入 OpenResty？.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/50%20%e7%ad%94%e7%96%91%ef%bc%88%e4%ba%94%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9c%a8%e5%b7%a5%e4%bd%9c%e4%b8%ad%e5%bc%95%e5%85%a5%20OpenResty%ef%bc%9f.md">50 答疑（五）：如何在工作中引入 OpenResty？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 行百里者半九十.md" href="/%e4%b8%93%e6%a0%8f/OpenResty%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e8%a1%8c%e7%99%be%e9%87%8c%e8%80%85%e5%8d%8a%e4%b9%9d%e5%8d%81.md">结束语 行百里者半九十.md</a>
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
                            <h1 id="title" data-id="06 OpenResty 中用到的 NGINX 知识" class="title">06 OpenResty 中用到的 NGINX 知识</h1>
                            <div><p>你好，我是温铭。</p>

<p>通过前面几篇文章的介绍，相信你对 OpenResty 的轮廓已经有了一个大概的认知。下面几节课里，我会带你熟悉下 OpenResty 的两个基石：NGINX 和 LuaJIT。万丈高楼平地起，掌握些这些基础的知识，才能更好地去学习 OpenResty。</p>

<p>今天我先来讲 NGINX。这里我只会介绍下，OpenResty 中可能会用到的一些 NGINX 基础知识，这些仅仅是 NGINX 很小的一个子集。如果你需要系统和深入学习 NGINX，可以参考陶辉老师的《NGINX 核心知识 100 讲》，这也是极客时间上评价非常高的一门课程。</p>

<p>说到配置，其实，在 OpenResty 的开发中，我们需要注意下面几点：</p>

<ul>
<li>要尽可能少地配置 nginx.conf；</li>
<li>避免使用if、set 、rewrite 等多个指令的配合；</li>
<li>能通过 Lua 代码解决的，就别用 NGINX 的配置、变量和模块来解决。</li>
</ul>

<p>这样可以最大限度地提高可读性、可维护性和可扩展性。</p>

<p>下面这段 NGINX 配置，就是一个典型的反例，可以说是把配置项当成了代码来使用：</p>

<pre><code>location ~ ^/mobile/(web/app.htm) {
            set $type $1;
            set $orig_args $args;
            if ( $http_user_Agent ~ &quot;(iPhone|iPad|Android)&quot; ) {
                rewrite  ^/mobile/(.*) http://touch.foo.com/mobile/$1 last;
            }
            proxy_pass http://foo.com/$type?$orig_args;
}
</code></pre>

<p>这是我们在使用 OpenResty 进行开发时需要避免的。</p>

<h2 id="nginx-配置"><strong>NGINX 配置</strong></h2>

<p>我们首先来看下 NGINX 的配置文件。NGINX 通过配置文件来控制自身行为，它的配置可以看作是一个简单的 DSL。NGINX 在进程启动的时候读取配置，并加载到内存中。<strong>如果修改了配置文件，需要你重启或者重载 NGINX，再次读取后才能生效</strong>。只有 NGINX 的商业版本，才会在运行时, 以 API 的形式提供部分动态的能力。</p>

<p>我们先来看下面这段配置，里面的内容非常简单，我相信大部分工程师都能看懂：</p>

<pre><code>worker_processes auto;

pid logs/nginx.pid;
error_log logs/error.log notice;

worker_rlimit_nofile 65535;

events {
    worker_connections 16384;
}

http {
    server {
	listen 80;
	listen 443 ssl;

        location / {
	    proxy_pass https://foo.com;
	    }
    }
}

stream {
    server {
        listen 53 udp;
    }
}
</code></pre>

<p>不过，即使是简单的配置，背后也涉及到了一些很重要的基础概念。</p>

<p>第一，每个指令都有自己适用的上下文（Context），也就是NGINX 配置文件中指令的作用域。</p>

<p>最上层的是 main，里面是和具体业务无关的一些指令，比如上面出现的 worker_processes、pid 和 error_log，都属于 main 这个上下文。另外，上下文是有层级关系的，比如 location 的上下文是 server，server 的上下文是 http，http 的上下文是 main。</p>

<p>指令不能运行在错误的上下文中，NGINX 在启动时会检测 nginx.conf 是否合法。比如我们把</p>

<p><code>listen 80;</code> 从 server 上下文换到 main 上下文，然后启动 NGINX 服务，会看到类似这样的报错：</p>

<pre><code>&quot;listen&quot; directive is not allowed here ......
</code></pre>

<p>第二，NGINX 不仅可以处理 HTTP 请求 和 HTTPS 流量，还可以处理 UDP 和 TCP 流量。</p>

<p>其中，七层的放在 HTTP 中，四层的放在 stream中。在 OpenResty 里面， lua-nginx-module 和 stream-lua-nginx-module 分别和这俩对应。</p>

<p>这里有一点需要注意，<strong>NGINX 支持的功能，OpenResty 并不一定支持，需要看 OpenResty 的版本号</strong>。OpenResty 的版本号是和 NGINX 保持一致的，所以很容易识别。比如 NGINX 在 2018 年 3 月份发布的 1.13.10 版本中，增加了对 gRPC 的支持，但 OpenResty 在 2019 年 4 月份时的最新版本是 1.13.6.2，由此可以推断 OpenResty 还不支持 gRPC。</p>

<p>上面 nginx.conf 涉及到的配置指令，都在 NGINX 的核心模块 <a href="http://nginx.org/en/docs/ngx_core_module.html" target="_blank">ngx_core_module</a>、<a href="http://nginx.org/en/docs/http/ngx_http_core_module.html" target="_blank">ngx_http_core<em>module</em></a> 和 <a href="http://nginx.org/en/docs/stream/ngx_stream_core_module.html" target="_blank">ngx_stream_core<em>module</em></a> 中，你可以点击这几个链接去查看具体的文档说明。</p>

<h2 id="master-worker-模式"><strong>MASTER-WORKER 模式</strong></h2>

<p>了解完配置文件，我们再来看下 NGINX 的多进程模式。这里我放了一张图来表示，你可以看到，NGINX 启动后，会有一个 Master 进程和多个 Worker 进程（也可以只有一个 Worker 进程，看你如何配置）。</p>

<p><img src="assets/a7304c2c8af0e1e6c54819c97611b992.jpg" alt="" /></p>

<p>先来说 Master 进程，一如其名，扮演“管理者”的角色，并不负责处理终端的请求。它是用来管理 Worker 进程的，包括接受管理员发送的信号量、监控 Worker 的运行状态。当 Worker 进程异常退出时，Master 进程会重新启动一个新的 Worker 进程。</p>

<p>Worker 进程则是“一线员工”，用来处理终端用户的请求。它是从 Master 进程 fork 出来的，彼此之间相互独立，互不影响。多进程的模式比 Apache 多线程的模式要先进很多，没有线程间加锁，也方便调试。即使某个进程崩溃退出了，也不会影响其他 Worker 进程正常工作。</p>

<p>而 OpenResty 在 NGINX Master-Worker 模式的前提下，又增加了独有的特权进程（privileged agent）。这个进程并不监听任何端口，和 NGINX 的 Master 进程拥有同样的权限，所以可以做一些需要高权限才能完成的任务，比如对本地磁盘文件的一些写操作等。</p>

<p>如果特权进程与 NGINX 二进制热升级的机制互相配合，OpenResty 就可以实现自我二进制热升级的整个流程，而不依赖任何外部的程序。</p>

<p>减少对外部程序的依赖，尽量在 OpenResty 进程内解决问题，不仅方便部署、降低运维成本，也可以降低程序出错的概率。可以说，OpenResty 中的特权进程、ngx.pipe 等功能，都是出于这个目的。</p>

<h2 id="执行阶段"><strong>执行阶段</strong></h2>

<p>执行阶段也是 NGINX 重要的特性，与 OpenResty 的具体实现密切相关。NGINX 有 11 个执行阶段，我们可以从 ngx_http_core_module.h 的源码中看到：</p>

<pre><code>typedef enum {
    NGX_HTTP_POST_READ_PHASE = 0,

    NGX_HTTP_SERVER_REWRITE_PHASE,

    NGX_HTTP_FIND_CONFIG_PHASE,
    NGX_HTTP_REWRITE_PHASE,
    NGX_HTTP_POST_REWRITE_PHASE,

    NGX_HTTP_PREACCESS_PHASE,

    NGX_HTTP_ACCESS_PHASE,
    NGX_HTTP_POST_ACCESS_PHASE,

    NGX_HTTP_PRECONTENT_PHASE,

    NGX_HTTP_CONTENT_PHASE,

    NGX_HTTP_LOG_PHASE
} ngx_http_phases;
</code></pre>

<p>如果你想详细了解这 11 个阶段的作用，可以学习陶辉老师的视频课程，或者 NGINX 文档，这里我就不再赘述。</p>

<p>不过，巧合的是，OpenResty 也有 11 个 <code>*_by_lua</code>指令，它们和 NGINX 阶段的关系如下图所示（图片来自 lua-nginx-module 文档）：</p>

<p><img src="assets/2a05cb2a679bd1c81b44508666e70273.png" alt="" /></p>

<p>其中， <code>init_by_lua</code> 只会在 Master 进程被创建时执行，<code>init_worker_by_lua</code> 只会在每个 Worker 进程被创建时执行。其他的 <code>*_by_lua</code> 指令则是由终端请求触发，会被反复执行。</p>

<p>所以在 init_by_lua 阶段，我们可以预先加载 Lua 模块和公共的只读数据，这样可以利用操作系统的 COW（copy on write）特性，来节省一些内存。</p>

<p>对于业务代码来说，其实大部分的操作都可以在 content_by_lua 里面完成，但我更推荐的做法，是根据不同的功能来进行拆分，比如下面这样：</p>

<ul>
<li>set_by_lua：设置变量；</li>
<li>rewrite_by_lua：转发、重定向等；</li>
<li>access_by_lua：准入、权限等；</li>
<li>content_by_lua：生成返回内容；</li>
<li>header_filter_by_lua：应答头过滤处理；</li>
<li>body_filter_by_lua：应答体过滤处理；</li>
<li>log_by_lua：日志记录。</li>
</ul>

<p>我举一个例子来说明这样拆分的好处。我们假设，你对外提供了很多明文 API，现在需要增加自定义的加密和解密逻辑。那么请问，你需要修改所有 API 的代码吗？</p>

<pre><code># 明文协议版本
location /mixed {
    content_by_lua '...';       # 处理请求
}
</code></pre>

<p>当然不用。事实上，利用阶段的特性，我们只需要简单地在 access 阶段解密，在 body filter 阶段加密就可以了，原来 content 阶段的代码是不用做任何修改的：</p>

<pre><code># 加密协议版本
location /mixed {
    access_by_lua '...';        # 请求体解密
    content_by_lua '...';       # 处理请求，不需要关心通信协议
    body_filter_by_lua '...';   # 应答体加密
}
</code></pre>

<h2 id="二进制热升级"><strong>二进制热升级</strong></h2>

<p>最后，我来简单说一下 NGINX 的二进制热升级。我们知道，在你修改完 NGINX 的配置文件后，还需要重启才能生效。但在 NGINX 升级自身版本的时候，却可以做到热升级。这看上去有点儿本末倒置，不过，考虑到 NGINX 是从传统静态的负载均衡、反向代理、文件缓存起家的，这倒也可以理解。</p>

<p>热升级通过向旧的 Master 进程发送 USR2 和 WINCH 信号量来完成。对于这两步，前者的作用，是启动新的 Master 进程；后者的作用，是逐步关闭 Worker 进程。</p>

<p>执行完这两步后，新的 Master 和新的 Worker 就已经启动了。不过此时，旧的 Master 并没有退出。不退出的原因也很简单，如果你需要回退，依旧可以给旧的 Master 发送 HUP 信号量。当然，如果你已经确定不需要回退，就可以给旧 Master 发送 KILL 信号量来退出。</p>

<p>至此，大功告成，二进制的热升级就完成了。</p>

<p>关于二进制升级，我主要就讲这些。如果你想了解这方面更详细的资料，可以查阅<a href="http://nginx.org/en/docs/control.html#upgrade" target="_blank">官方文档</a>继续学习。</p>

<h2 id="课外延伸"><strong>课外延伸</strong></h2>

<p>OpenResty 的作者多年前写过一个 <a href="https://openresty.org/download/agentzh-nginx-tutorials-zhcn.html" target="_blank">NGINX 教程</a>，如果你对此感兴趣，可以自己学习下。这里面的内容比较多，即使看不懂也没有关系，并不会影响你学习 OpenResty。</p>

<h2 id="写在最后">写在最后</h2>

<p>总的来说，在 OpenResty 中用到的都是 Nginx 的基础知识，主要涉及到配置、主从进程、执行阶段等。而<strong>其他能用 Lua 代码解决的，尽量用代码来解决，而非使用Nginx 的模块和配置</strong>，这是在学习 OpenResty 中的一个思路转变。</p>

<p>最后，我给你留了一道开放的思考题。Nginx 官方支持 NJS，也就是可以用 JS 写控制部分 Nginx 的逻辑，和 OpenResty 的思路很类似。对此，你是怎么看待的呢？</p>

<p>欢迎留言和我分享，也欢迎你把这篇文章转发给你的同事、朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#056969693c3134343532456268646c692b666a68" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f118e694fe294f1',t:'MTczNDA0Njc5My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>