<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=35&#32;&#32;OpenResty：更灵活的Web服务器>
        <link rel="icon" href="/static/favicon.png">
        <title>35  OpenResty：更灵活的Web服务器 </title>
        
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
                        <a class="menu-item" id="00 开篇词｜To Be a HTTP Hero.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/00%20%e5%bc%80%e7%af%87%e8%af%8d%ef%bd%9cTo%20Be%20a%20HTTP%20Hero.md">00 开篇词｜To Be a HTTP Hero.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  时势与英雄：HTTP的前世今生.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/01%20%20%e6%97%b6%e5%8a%bf%e4%b8%8e%e8%8b%b1%e9%9b%84%ef%bc%9aHTTP%e7%9a%84%e5%89%8d%e4%b8%96%e4%bb%8a%e7%94%9f.md">01  时势与英雄：HTTP的前世今生.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  HTTP是什么？HTTP又不是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/02%20%20HTTP%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9fHTTP%e5%8f%88%e4%b8%8d%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">02  HTTP是什么？HTTP又不是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  HTTP世界全览（上）：与HTTP相关的各种概念.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/03%20%20HTTP%e4%b8%96%e7%95%8c%e5%85%a8%e8%a7%88%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e4%b8%8eHTTP%e7%9b%b8%e5%85%b3%e7%9a%84%e5%90%84%e7%a7%8d%e6%a6%82%e5%bf%b5.md">03  HTTP世界全览（上）：与HTTP相关的各种概念.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  HTTP世界全览（下）：与HTTP相关的各种协议.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/04%20%20HTTP%e4%b8%96%e7%95%8c%e5%85%a8%e8%a7%88%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e4%b8%8eHTTP%e7%9b%b8%e5%85%b3%e7%9a%84%e5%90%84%e7%a7%8d%e5%8d%8f%e8%ae%ae.md">04  HTTP世界全览（下）：与HTTP相关的各种协议.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  常说的“四层”和“七层”到底是什么？“五层”“六层”哪去了？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/05%20%20%e5%b8%b8%e8%af%b4%e7%9a%84%e2%80%9c%e5%9b%9b%e5%b1%82%e2%80%9d%e5%92%8c%e2%80%9c%e4%b8%83%e5%b1%82%e2%80%9d%e5%88%b0%e5%ba%95%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f%e2%80%9c%e4%ba%94%e5%b1%82%e2%80%9d%e2%80%9c%e5%85%ad%e5%b1%82%e2%80%9d%e5%93%aa%e5%8e%bb%e4%ba%86%ef%bc%9f.md">05  常说的“四层”和“七层”到底是什么？“五层”“六层”哪去了？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  域名里有哪些门道？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/06%20%20%e5%9f%9f%e5%90%8d%e9%87%8c%e6%9c%89%e5%93%aa%e4%ba%9b%e9%97%a8%e9%81%93%ef%bc%9f.md">06  域名里有哪些门道？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  自己动手，搭建HTTP实验环境.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/07%20%20%e8%87%aa%e5%b7%b1%e5%8a%a8%e6%89%8b%ef%bc%8c%e6%90%ad%e5%bb%baHTTP%e5%ae%9e%e9%aa%8c%e7%8e%af%e5%a2%83.md">07  自己动手，搭建HTTP实验环境.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  键入网址再按下回车，后面究竟发生了什么？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/08%20%20%e9%94%ae%e5%85%a5%e7%bd%91%e5%9d%80%e5%86%8d%e6%8c%89%e4%b8%8b%e5%9b%9e%e8%bd%a6%ef%bc%8c%e5%90%8e%e9%9d%a2%e7%a9%b6%e7%ab%9f%e5%8f%91%e7%94%9f%e4%ba%86%e4%bb%80%e4%b9%88%ef%bc%9f.md">08  键入网址再按下回车，后面究竟发生了什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  HTTP报文是什么样子的？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/09%20%20HTTP%e6%8a%a5%e6%96%87%e6%98%af%e4%bb%80%e4%b9%88%e6%a0%b7%e5%ad%90%e7%9a%84%ef%bc%9f.md">09  HTTP报文是什么样子的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  应该如何理解请求方法？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/10%20%20%e5%ba%94%e8%af%a5%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%e8%af%b7%e6%b1%82%e6%96%b9%e6%b3%95%ef%bc%9f.md">10  应该如何理解请求方法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  你能写出正确的网址吗？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/11%20%20%e4%bd%a0%e8%83%bd%e5%86%99%e5%87%ba%e6%ad%a3%e7%a1%ae%e7%9a%84%e7%bd%91%e5%9d%80%e5%90%97%ef%bc%9f.md">11  你能写出正确的网址吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  响应状态码该怎么用？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/12%20%20%e5%93%8d%e5%ba%94%e7%8a%b6%e6%80%81%e7%a0%81%e8%af%a5%e6%80%8e%e4%b9%88%e7%94%a8%ef%bc%9f.md">12  响应状态码该怎么用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  HTTP有哪些特点？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/13%20%20HTTP%e6%9c%89%e5%93%aa%e4%ba%9b%e7%89%b9%e7%82%b9%ef%bc%9f.md">13  HTTP有哪些特点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  HTTP有哪些优点？又有哪些缺点？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/14%20%20HTTP%e6%9c%89%e5%93%aa%e4%ba%9b%e4%bc%98%e7%82%b9%ef%bc%9f%e5%8f%88%e6%9c%89%e5%93%aa%e4%ba%9b%e7%bc%ba%e7%82%b9%ef%bc%9f.md">14  HTTP有哪些优点？又有哪些缺点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  海纳百川：HTTP的实体数据.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/15%20%20%e6%b5%b7%e7%ba%b3%e7%99%be%e5%b7%9d%ef%bc%9aHTTP%e7%9a%84%e5%ae%9e%e4%bd%93%e6%95%b0%e6%8d%ae.md">15  海纳百川：HTTP的实体数据.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  把大象装进冰箱：HTTP传输大文件的方法.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/16%20%20%e6%8a%8a%e5%a4%a7%e8%b1%a1%e8%a3%85%e8%bf%9b%e5%86%b0%e7%ae%b1%ef%bc%9aHTTP%e4%bc%a0%e8%be%93%e5%a4%a7%e6%96%87%e4%bb%b6%e7%9a%84%e6%96%b9%e6%b3%95.md">16  把大象装进冰箱：HTTP传输大文件的方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  排队也要讲效率：HTTP的连接管理.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/17%20%20%e6%8e%92%e9%98%9f%e4%b9%9f%e8%a6%81%e8%ae%b2%e6%95%88%e7%8e%87%ef%bc%9aHTTP%e7%9a%84%e8%bf%9e%e6%8e%a5%e7%ae%a1%e7%90%86.md">17  排队也要讲效率：HTTP的连接管理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  四通八达：HTTP的重定向和跳转.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/18%20%20%e5%9b%9b%e9%80%9a%e5%85%ab%e8%be%be%ef%bc%9aHTTP%e7%9a%84%e9%87%8d%e5%ae%9a%e5%90%91%e5%92%8c%e8%b7%b3%e8%bd%ac.md">18  四通八达：HTTP的重定向和跳转.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  让我知道你是谁：HTTP的Cookie机制.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/19%20%20%e8%ae%a9%e6%88%91%e7%9f%a5%e9%81%93%e4%bd%a0%e6%98%af%e8%b0%81%ef%bc%9aHTTP%e7%9a%84Cookie%e6%9c%ba%e5%88%b6.md">19  让我知道你是谁：HTTP的Cookie机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  生鲜速递：HTTP的缓存控制.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/20%20%20%e7%94%9f%e9%b2%9c%e9%80%9f%e9%80%92%ef%bc%9aHTTP%e7%9a%84%e7%bc%93%e5%ad%98%e6%8e%a7%e5%88%b6.md">20  生鲜速递：HTTP的缓存控制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  良心中间商：HTTP的代理服务.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/21%20%20%e8%89%af%e5%bf%83%e4%b8%ad%e9%97%b4%e5%95%86%ef%bc%9aHTTP%e7%9a%84%e4%bb%a3%e7%90%86%e6%9c%8d%e5%8a%a1.md">21  良心中间商：HTTP的代理服务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  冷链周转：HTTP的缓存代理.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/22%20%20%e5%86%b7%e9%93%be%e5%91%a8%e8%bd%ac%ef%bc%9aHTTP%e7%9a%84%e7%bc%93%e5%ad%98%e4%bb%a3%e7%90%86.md">22  冷链周转：HTTP的缓存代理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  HTTPS是什么？SSLTLS又是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/23%20%20HTTPS%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9fSSLTLS%e5%8f%88%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">23  HTTPS是什么？SSLTLS又是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  固若金汤的根本（上）：对称加密与非对称加密.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/24%20%20%e5%9b%ba%e8%8b%a5%e9%87%91%e6%b1%a4%e7%9a%84%e6%a0%b9%e6%9c%ac%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%af%b9%e7%a7%b0%e5%8a%a0%e5%af%86%e4%b8%8e%e9%9d%9e%e5%af%b9%e7%a7%b0%e5%8a%a0%e5%af%86.md">24  固若金汤的根本（上）：对称加密与非对称加密.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25  固若金汤的根本（下）：数字签名与证书.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/25%20%20%e5%9b%ba%e8%8b%a5%e9%87%91%e6%b1%a4%e7%9a%84%e6%a0%b9%e6%9c%ac%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%95%b0%e5%ad%97%e7%ad%be%e5%90%8d%e4%b8%8e%e8%af%81%e4%b9%a6.md">25  固若金汤的根本（下）：数字签名与证书.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26  信任始于握手：TLS1.2连接过程解析.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/26%20%20%e4%bf%a1%e4%bb%bb%e5%a7%8b%e4%ba%8e%e6%8f%a1%e6%89%8b%ef%bc%9aTLS1.2%e8%bf%9e%e6%8e%a5%e8%bf%87%e7%a8%8b%e8%a7%a3%e6%9e%90.md">26  信任始于握手：TLS1.2连接过程解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27  更好更快的握手：TLS1.3特性解析.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/27%20%20%e6%9b%b4%e5%a5%bd%e6%9b%b4%e5%bf%ab%e7%9a%84%e6%8f%a1%e6%89%8b%ef%bc%9aTLS1.3%e7%89%b9%e6%80%a7%e8%a7%a3%e6%9e%90.md">27  更好更快的握手：TLS1.3特性解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28  连接太慢该怎么办：HTTPS的优化.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/28%20%20%e8%bf%9e%e6%8e%a5%e5%a4%aa%e6%85%a2%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9aHTTPS%e7%9a%84%e4%bc%98%e5%8c%96.md">28  连接太慢该怎么办：HTTPS的优化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29  我应该迁移到HTTPS吗？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/29%20%20%e6%88%91%e5%ba%94%e8%af%a5%e8%bf%81%e7%a7%bb%e5%88%b0HTTPS%e5%90%97%ef%bc%9f.md">29  我应该迁移到HTTPS吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30  时代之风（上）：HTTP2特性概览.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/30%20%20%e6%97%b6%e4%bb%a3%e4%b9%8b%e9%a3%8e%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9aHTTP2%e7%89%b9%e6%80%a7%e6%a6%82%e8%a7%88.md">30  时代之风（上）：HTTP2特性概览.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31  时代之风（下）：HTTP2内核剖析.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/31%20%20%e6%97%b6%e4%bb%a3%e4%b9%8b%e9%a3%8e%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aHTTP2%e5%86%85%e6%a0%b8%e5%89%96%e6%9e%90.md">31  时代之风（下）：HTTP2内核剖析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32  未来之路：HTTP3展望.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/32%20%20%e6%9c%aa%e6%9d%a5%e4%b9%8b%e8%b7%af%ef%bc%9aHTTP3%e5%b1%95%e6%9c%9b.md">32  未来之路：HTTP3展望.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33  我应该迁移到HTTP2吗？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/33%20%20%e6%88%91%e5%ba%94%e8%af%a5%e8%bf%81%e7%a7%bb%e5%88%b0HTTP2%e5%90%97%ef%bc%9f.md">33  我应该迁移到HTTP2吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34  Nginx：高性能的Web服务器.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/34%20%20Nginx%ef%bc%9a%e9%ab%98%e6%80%a7%e8%83%bd%e7%9a%84Web%e6%9c%8d%e5%8a%a1%e5%99%a8.md">34  Nginx：高性能的Web服务器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35  OpenResty：更灵活的Web服务器.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/35%20%20OpenResty%ef%bc%9a%e6%9b%b4%e7%81%b5%e6%b4%bb%e7%9a%84Web%e6%9c%8d%e5%8a%a1%e5%99%a8.md">35  OpenResty：更灵活的Web服务器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36  WAF：保护我们的网络服务.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/36%20%20WAF%ef%bc%9a%e4%bf%9d%e6%8a%a4%e6%88%91%e4%bb%ac%e7%9a%84%e7%bd%91%e7%bb%9c%e6%9c%8d%e5%8a%a1.md">36  WAF：保护我们的网络服务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37  CDN：加速我们的网络服务.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/37%20%20CDN%ef%bc%9a%e5%8a%a0%e9%80%9f%e6%88%91%e4%bb%ac%e7%9a%84%e7%bd%91%e7%bb%9c%e6%9c%8d%e5%8a%a1.md">37  CDN：加速我们的网络服务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38  WebSocket：沙盒里的TCP.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/38%20%20WebSocket%ef%bc%9a%e6%b2%99%e7%9b%92%e9%87%8c%e7%9a%84TCP.md">38  WebSocket：沙盒里的TCP.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39  HTTP性能优化面面观（上）.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/39%20%20HTTP%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e9%9d%a2%e9%9d%a2%e8%a7%82%ef%bc%88%e4%b8%8a%ef%bc%89.md">39  HTTP性能优化面面观（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40  HTTP性能优化面面观（下）.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/40%20%20HTTP%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e9%9d%a2%e9%9d%a2%e8%a7%82%ef%bc%88%e4%b8%8b%ef%bc%89.md">40  HTTP性能优化面面观（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语  做兴趣使然的Hero.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e5%81%9a%e5%85%b4%e8%b6%a3%e4%bd%bf%e7%84%b6%e7%9a%84Hero.md">结束语  做兴趣使然的Hero.md</a>
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
                            <h1 id="title" data-id="35  OpenResty：更灵活的Web服务器" class="title">35  OpenResty：更灵活的Web服务器</h1>
                            <div><p>在上一讲里，我们看到了高性能的 Web 服务器 Nginx，它资源占用少，处理能力高，是搭建网站的首选。</p>

<p>虽然 Nginx 成为了 Web 服务器领域无可争议的“王者”，但它也并不是没有缺点的，毕竟它已经 15 岁了。</p>

<p>“一个人很难超越时代，而时代却可以轻易超越所有人”，Nginx 当初设计时针对的应用场景已经发生了变化，它的一些缺点也就暴露出来了。</p>

<p>Nginx 的服务管理思路延续了当时的流行做法，使用磁盘上的静态配置文件，所以每次修改后必须重启才能生效。</p>

<p>这在业务频繁变动的时候是非常致命的（例如流行的微服务架构），特别是对于拥有成千上万台服务器的网站来说，仅仅增加或者删除一行配置就要分发、重启所有的机器，对运维是一个非常大的挑战，要耗费很多的时间和精力，成本很高，很不灵活，难以“随需应变”。</p>

<p>那么，有没有这样的一个 Web 服务器，它有 Nginx 的优点却没有 Nginx 的缺点，既轻量级、高性能，又灵活、可动态配置呢？</p>

<p>这就是我今天要说的 OpenResty，它是一个“更好更灵活的 Nginx”。</p>

<h2 id="openresty-是什么">OpenResty 是什么？</h2>

<p>其实你对 OpenResty 并不陌生，这个专栏的实验环境就是用 OpenResty 搭建的，这么多节课程下来，你应该或多或少对它有了一些印象吧。</p>

<p>OpenResty 诞生于 2009 年，到现在刚好满 10 周岁。它的创造者是当时就职于某宝的“神级”程序员<strong>章亦春</strong>，网名叫“agentzh”。</p>

<p>OpenResty 并不是一个全新的 Web 服务器，而是基于 Nginx，它利用了 Nginx 模块化、可扩展的特性，开发了一系列的增强模块，并把它们打包整合，形成了一个<strong>“一站式”的 Web 开发平台</strong>。</p>

<p>虽然 OpenResty 的核心是 Nginx，但它又超越了 Nginx，关键就在于其中的 ngx_lua 模块，把小巧灵活的 Lua 语言嵌入了 Nginx，可以用脚本的方式操作 Nginx 内部的进程、多路复用、阶段式处理等各种构件。</p>

<p>脚本语言的好处你一定知道，它不需要编译，随写随执行，这就免去了 C 语言编写模块漫长的开发周期。而且 OpenResty 还把 Lua 自身的协程与 Nginx 的事件机制完美结合在一起，优雅地实现了许多其他语言所没有的“<strong>同步非阻塞</strong>”编程范式，能够轻松开发出高性能的 Web 应用。</p>

<p>目前 OpenResty 有两个分支，分别是开源、免费的“OpenResty”和闭源、商业产品的“OpenResty+”，运作方式有社区支持、OpenResty 基金会、OpenResty.Inc 公司，还有其他的一些外界赞助（例如 Kong、CloudFlare），正在蓬勃发展。</p>

<p><img src="assets/9f7b79c43c476890f03c2c716a20f301.png" alt="unpreview" /></p>

<p>顺便说一下 OpenResty 的官方 logo，是一只展翅飞翔的海鸥，选择海鸥是因为“鸥”与 OpenResty 的发音相同。另外，这个 logo 的形状也像是左手比出的一个“OK”姿势，正好也是一个“O”。</p>

<h2 id="动态的-lua">动态的 Lua</h2>

<p>刚才说了，OpenResty 里的一个关键模块是 ngx_lua，它为 Nginx 引入了脚本语言 Lua。</p>

<p>Lua 是一个比较“小众”的语言，虽然历史比较悠久，但名气却没有 PHP、Python、JavaScript 大，这主要与它的自身定位有关。</p>

<p><img src="assets/4f24aa3f53969b71baaf7d9c7cf68fd5.png" alt="unpreview" /></p>

<p>Lua 的设计目标是嵌入到其他应用程序里运行，为其他编程语言带来“脚本化”能力，所以它的“个头”比较小，功能集有限，不追求“大而全”，而是“小而美”，大多数时间都“隐匿”在其他应用程序的后面，是“无名英雄”。</p>

<p>你或许玩过或者听说过《魔兽世界》《愤怒的小鸟》吧，它们就在内部嵌入了 Lua，使用 Lua 来调用底层接口，充当“胶水语言”（glue language），编写游戏逻辑脚本，提高开发效率。</p>

<p>OpenResty 选择 Lua 作为“工作语言”也是基于同样的考虑。因为 Nginx C 开发实在是太麻烦了，限制了 Nginx 的真正实力。而 Lua 作为“最快的脚本语言”恰好可以成为 Nginx 的完美搭档，既可以简化开发，性能上又不会有太多的损耗。</p>

<p>作为脚本语言，Lua 还有一个重要的“<strong>代码热加载</strong>”特性，不需要重启进程，就能够从磁盘、Redis 或者任何其他地方加载数据，随时替换内存里的代码片段。这就带来了“<strong>动态配置</strong>”，让 OpenResty 能够永不停机，在微秒、毫秒级别实现配置和业务逻辑的实时更新，比起 Nginx 秒级的重启是一个极大的进步。</p>

<p>你可以看一下实验环境的“www/lua”目录，里面存放了我写的一些测试 HTTP 特性的 Lua 脚本，代码都非常简单易懂，就像是普通的英语“阅读理解”，这也是 Lua 的另一个优势：易学习、易上手。</p>

<h2 id="高效率的-lua">高效率的 Lua</h2>

<p>OpenResty 能够高效运行的一大“秘技”是它的“<strong>同步非阻塞</strong>”编程范式，如果你要开发 OpenResty 应用就必须时刻铭记于心。</p>

<p>“同步非阻塞”本质上还是一种“<strong>多路复用</strong>”，我拿上一讲的 Nginx epoll 来对比解释一下。</p>

<p>epoll 是操作系统级别的“多路复用”，运行在内核空间。而 OpenResty 的“同步非阻塞”则是基于 Lua 内建的“<strong>协程</strong>”，是应用程序级别的“多路复用”，运行在用户空间，所以它的资源消耗要更少。</p>

<p>OpenResty 里每一段 Lua 程序都由协程来调度运行。和 Linux 的 epoll 一样，每当可能发生阻塞的时候“协程”就会立刻切换出去，执行其他的程序。这样单个处理流程是“阻塞”的，但整个 OpenResty 却是“非阻塞的”，多个程序都“复用”在一个 Lua 虚拟机里运行。</p>

<p><img src="assets/9fc3df52df7d6c11aa02b8013f8e9bc6.png" alt="img" /></p>

<p>下面的代码是一个简单的例子，读取 POST 发送的 body 数据，然后再发回客户端：</p>

<pre><code>ngx.req.read_body()                  -- 同步非阻塞 (1)
 
local data = ngx.req.get_body_data()
if data then
    ngx.print(&quot;body: &quot;, data)        -- 同步非阻塞 (2)
end
</code></pre>

<p>代码中的“ngx.req.read_body”和“ngx.print”分别是数据的收发动作，只有收到数据才能发送数据，所以是“同步”的。</p>

<p>但即使因为网络原因没收到或者发不出去，OpenResty 也不会在这里阻塞“干等着”，而是做个“记号”，把等待的这段 CPU 时间用来处理其他的请求，等网络可读或者可写时再“回来”接着运行。</p>

<p>假设收发数据的等待时间是 10 毫秒，而真正 CPU 处理的时间是 0.1 毫秒，那么 OpenResty 就可以在这 10 毫秒内同时处理 100 个请求，而不是把这 100 个请求阻塞排队，用 1000 毫秒来处理。</p>

<p>除了“同步非阻塞”，OpenResty 还选用了<strong>LuaJIT</strong>作为 Lua 语言的“运行时（Runtime）”，进一步“挖潜增效”。</p>

<p>LuaJIT 是一个高效的 Lua 虚拟机，支持 JIT（Just In Time）技术，可以把 Lua 代码即时编译成“本地机器码”，这样就消除了脚本语言解释运行的劣势，让 Lua 脚本跑得和原生 C 代码一样快。</p>

<p>另外，LuaJIT 还为 Lua 语言添加了一些特别的增强，比如二进制位运算库 bit，内存优化库 table，还有 FFI（Foreign Function Interface），让 Lua 直接调用底层 C 函数，比原生的压栈调用快很多。</p>

<h2 id="阶段式处理">阶段式处理</h2>

<p>和 Nginx 一样，OpenResty 也使用“流水线”来处理 HTTP 请求，底层的运行基础是 Nginx 的“阶段式处理”，但它又有自己的特色。</p>

<p>Nginx 的“流水线”是由一个个 C 模块组成的，只能在静态文件里配置，开发困难，配置麻烦（相对而言）。而 OpenResty 的“流水线”则是由一个个的 Lua 脚本组成的，不仅可以从磁盘上加载，也可以从 Redis、MySQL 里加载，而且编写、调试的过程非常方便快捷。</p>

<p>下面我画了一张图，列出了 OpenResty 的阶段，比起 Nginx，OpenResty 的阶段更注重对 HTTP 请求响应报文的加工和处理。</p>

<p><img src="assets/3689312a970bae0e949b017ad45438df.png" alt="img" /></p>

<p>OpenResty 里有几个阶段与 Nginx 是相同的，比如 rewrite、access、content、filter，这些都是标准的 HTTP 处理。</p>

<p>在这几个阶段里可以用“xxx_by_lua”指令嵌入 Lua 代码，执行重定向跳转、访问控制、产生响应、负载均衡、过滤报文等功能。因为 Lua 的脚本语言特性，不用考虑内存分配、资源回收释放等底层的细节问题，可以专注于编写非常复杂的业务逻辑，比 C 模块的开发效率高很多，即易于扩展又易于维护。</p>

<p>OpenResty 里还有两个不同于 Nginx 的特殊阶段。</p>

<p>一个是“<strong>init 阶段</strong>”，它又分成“master init”和“worker init”，在 master 进程和 worker 进程启动的时候运行。这个阶段还没有开始提供服务，所以慢一点也没关系，可以调用一些阻塞的接口初始化服务器，比如读取磁盘、MySQL，加载黑白名单或者数据模型，然后放进共享内存里供运行时使用。</p>

<p>另一个是“<strong>ssl 阶段</strong>”，这算得上是 OpenResty 的一大创举，可以在 TLS 握手时动态加载证书，或者发送“OCSP Stapling”。</p>

<p>还记得[第 29 讲]里说的“SNI 扩展”吗？Nginx 可以依据“服务器名称指示”来选择证书实现 HTTPS 虚拟主机，但静态配置很不灵活，要编写很多雷同的配置块。虽然后来 Nginx 增加了变量支持，但它每次握手都要读磁盘，效率很低。</p>

<p>而在 OpenResty 里就可以使用指令“ssl_certificate_by_lua”，编写 Lua 脚本，读取 SNI 名字后，直接从共享内存或者 Redis 里获取证书。不仅没有读盘阻塞，而且证书也是完全动态可配置的，无需修改配置文件就能够轻松支持大量的 HTTPS 虚拟主机。</p>

<h2 id="小结">小结</h2>

<ol>
<li>Nginx 依赖于磁盘上的静态配置文件，修改后必须重启才能生效，缺乏灵活性；</li>
<li>OpenResty 基于 Nginx，打包了很多有用的模块和库，是一个高性能的 Web 开发平台；</li>
<li>OpenResty 的工作语言是 Lua，它小巧灵活，执行效率高，支持“代码热加载”；</li>
<li>OpenResty 的核心编程范式是“同步非阻塞”，使用协程，不需要异步回调函数；</li>
<li>OpenResty 也使用“阶段式处理”的工作模式，但因为在阶段里执行的都是 Lua 代码，所以非常灵活，配合 Redis 等外部数据库能够实现各种动态配置。</li>
</ol>

<h2 id="课下作业">课下作业</h2>

<ol>
<li>谈一下这些天你对实验环境里 OpenResty 的感想和认识。</li>
<li>你觉得 Nginx 和 OpenResty 的“阶段式处理”有什么好处？对你的实际工作有没有启发？</li>
</ol>

<p>欢迎你把自己的学习体会写在留言区，与我和其他同学一起讨论。如果你觉得有所收获，也欢迎把文章分享给你的朋友。</p>

<p><img src="assets/c5b7ac40c585c800af0fe3ab98f3449f.png" alt="unpreview" /></p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#660a0a0a5f525757565126010b070f0a4805090b" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f19ff8e5b09888b',t:'MTczNDEzNTMxMy4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>