<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=34&#32;基于JSON的RESTful接口协议：我不关心过程，请给我结果>
        <link rel="icon" href="/static/favicon.png">
        <title>34 基于JSON的RESTful接口协议：我不关心过程，请给我结果 </title>
        
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
                        <a class="menu-item" id="00 开篇词 想成为技术牛人？先搞定网络协议！.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e6%83%b3%e6%88%90%e4%b8%ba%e6%8a%80%e6%9c%af%e7%89%9b%e4%ba%ba%ef%bc%9f%e5%85%88%e6%90%9e%e5%ae%9a%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae%ef%bc%81.md">00 开篇词 想成为技术牛人？先搞定网络协议！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 为什么要学习网络协议？.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/01%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e5%ad%a6%e4%b9%a0%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae%ef%bc%9f.md">01 为什么要学习网络协议？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 网络分层的真实含义是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/02%20%e7%bd%91%e7%bb%9c%e5%88%86%e5%b1%82%e7%9a%84%e7%9c%9f%e5%ae%9e%e5%90%ab%e4%b9%89%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">02 网络分层的真实含义是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 ifconfig：最熟悉又陌生的命令行.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/03%20ifconfig%ef%bc%9a%e6%9c%80%e7%86%9f%e6%82%89%e5%8f%88%e9%99%8c%e7%94%9f%e7%9a%84%e5%91%bd%e4%bb%a4%e8%a1%8c.md">03 ifconfig：最熟悉又陌生的命令行.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 DHCP与PXE：IP是怎么来的，又是怎么没的？.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/04%20DHCP%e4%b8%8ePXE%ef%bc%9aIP%e6%98%af%e6%80%8e%e4%b9%88%e6%9d%a5%e7%9a%84%ef%bc%8c%e5%8f%88%e6%98%af%e6%80%8e%e4%b9%88%e6%b2%a1%e7%9a%84%ef%bc%9f.md">04 DHCP与PXE：IP是怎么来的，又是怎么没的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 从物理层到MAC层：如何在宿舍里自己组网玩联机游戏？.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/05%20%e4%bb%8e%e7%89%a9%e7%90%86%e5%b1%82%e5%88%b0MAC%e5%b1%82%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9c%a8%e5%ae%bf%e8%88%8d%e9%87%8c%e8%87%aa%e5%b7%b1%e7%bb%84%e7%bd%91%e7%8e%a9%e8%81%94%e6%9c%ba%e6%b8%b8%e6%88%8f%ef%bc%9f.md">05 从物理层到MAC层：如何在宿舍里自己组网玩联机游戏？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 交换机与VLAN：办公室太复杂，我要回学校.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/06%20%e4%ba%a4%e6%8d%a2%e6%9c%ba%e4%b8%8eVLAN%ef%bc%9a%e5%8a%9e%e5%85%ac%e5%ae%a4%e5%a4%aa%e5%a4%8d%e6%9d%82%ef%bc%8c%e6%88%91%e8%a6%81%e5%9b%9e%e5%ad%a6%e6%a0%a1.md">06 交换机与VLAN：办公室太复杂，我要回学校.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 ICMP与ping：投石问路的侦察兵.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/07%20ICMP%e4%b8%8eping%ef%bc%9a%e6%8a%95%e7%9f%b3%e9%97%ae%e8%b7%af%e7%9a%84%e4%be%a6%e5%af%9f%e5%85%b5.md">07 ICMP与ping：投石问路的侦察兵.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 世界这么大，我想出网关：欧洲十国游与玄奘西行.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/08%20%e4%b8%96%e7%95%8c%e8%bf%99%e4%b9%88%e5%a4%a7%ef%bc%8c%e6%88%91%e6%83%b3%e5%87%ba%e7%bd%91%e5%85%b3%ef%bc%9a%e6%ac%a7%e6%b4%b2%e5%8d%81%e5%9b%bd%e6%b8%b8%e4%b8%8e%e7%8e%84%e5%a5%98%e8%a5%bf%e8%a1%8c.md">08 世界这么大，我想出网关：欧洲十国游与玄奘西行.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 路由协议：西出网关无故人，敢问路在何方.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/09%20%e8%b7%af%e7%94%b1%e5%8d%8f%e8%ae%ae%ef%bc%9a%e8%a5%bf%e5%87%ba%e7%bd%91%e5%85%b3%e6%97%a0%e6%95%85%e4%ba%ba%ef%bc%8c%e6%95%a2%e9%97%ae%e8%b7%af%e5%9c%a8%e4%bd%95%e6%96%b9.md">09 路由协议：西出网关无故人，敢问路在何方.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 UDP协议：因性善而简单，难免碰到“城会玩”.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/10%20UDP%e5%8d%8f%e8%ae%ae%ef%bc%9a%e5%9b%a0%e6%80%a7%e5%96%84%e8%80%8c%e7%ae%80%e5%8d%95%ef%bc%8c%e9%9a%be%e5%85%8d%e7%a2%b0%e5%88%b0%e2%80%9c%e5%9f%8e%e4%bc%9a%e7%8e%a9%e2%80%9d.md">10 UDP协议：因性善而简单，难免碰到“城会玩”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 TCP协议（上）：因性恶而复杂，先恶后善反轻松.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/11%20TCP%e5%8d%8f%e8%ae%ae%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%9b%a0%e6%80%a7%e6%81%b6%e8%80%8c%e5%a4%8d%e6%9d%82%ef%bc%8c%e5%85%88%e6%81%b6%e5%90%8e%e5%96%84%e5%8f%8d%e8%bd%bb%e6%9d%be.md">11 TCP协议（上）：因性恶而复杂，先恶后善反轻松.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 TCP协议（下）：西行必定多妖孽，恒心智慧消磨难.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/12%20TCP%e5%8d%8f%e8%ae%ae%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e8%a5%bf%e8%a1%8c%e5%bf%85%e5%ae%9a%e5%a4%9a%e5%a6%96%e5%ad%bd%ef%bc%8c%e6%81%92%e5%bf%83%e6%99%ba%e6%85%a7%e6%b6%88%e7%a3%a8%e9%9a%be.md">12 TCP协议（下）：西行必定多妖孽，恒心智慧消磨难.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 套接字Socket：Talk is cheap, show me the code.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/13%20%e5%a5%97%e6%8e%a5%e5%ad%97Socket%ef%bc%9aTalk%20is%20cheap,%20show%20me%20the%20code.md">13 套接字Socket：Talk is cheap, show me the code.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 HTTP协议：看个新闻原来这么麻烦.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/14%20HTTP%e5%8d%8f%e8%ae%ae%ef%bc%9a%e7%9c%8b%e4%b8%aa%e6%96%b0%e9%97%bb%e5%8e%9f%e6%9d%a5%e8%bf%99%e4%b9%88%e9%ba%bb%e7%83%a6.md">14 HTTP协议：看个新闻原来这么麻烦.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 HTTPS协议：点外卖的过程原来这么复杂.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/15%20HTTPS%e5%8d%8f%e8%ae%ae%ef%bc%9a%e7%82%b9%e5%a4%96%e5%8d%96%e7%9a%84%e8%bf%87%e7%a8%8b%e5%8e%9f%e6%9d%a5%e8%bf%99%e4%b9%88%e5%a4%8d%e6%9d%82.md">15 HTTPS协议：点外卖的过程原来这么复杂.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 流媒体协议：如何在直播里看到美女帅哥？.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/16%20%e6%b5%81%e5%aa%92%e4%bd%93%e5%8d%8f%e8%ae%ae%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9c%a8%e7%9b%b4%e6%92%ad%e9%87%8c%e7%9c%8b%e5%88%b0%e7%be%8e%e5%a5%b3%e5%b8%85%e5%93%a5%ef%bc%9f.md">16 流媒体协议：如何在直播里看到美女帅哥？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 P2P协议：我下小电影，99%急死你.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/17%20P2P%e5%8d%8f%e8%ae%ae%ef%bc%9a%e6%88%91%e4%b8%8b%e5%b0%8f%e7%94%b5%e5%bd%b1%ef%bc%8c99%25%e6%80%a5%e6%ad%bb%e4%bd%a0.md">17 P2P协议：我下小电影，99%急死你.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 DNS协议：网络世界的地址簿.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/18%20DNS%e5%8d%8f%e8%ae%ae%ef%bc%9a%e7%bd%91%e7%bb%9c%e4%b8%96%e7%95%8c%e7%9a%84%e5%9c%b0%e5%9d%80%e7%b0%bf.md">18 DNS协议：网络世界的地址簿.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 HttpDNS：网络世界的地址簿也会指错路.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/19%20HttpDNS%ef%bc%9a%e7%bd%91%e7%bb%9c%e4%b8%96%e7%95%8c%e7%9a%84%e5%9c%b0%e5%9d%80%e7%b0%bf%e4%b9%9f%e4%bc%9a%e6%8c%87%e9%94%99%e8%b7%af.md">19 HttpDNS：网络世界的地址簿也会指错路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 CDN：你去小卖部取过快递么？.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/20%20CDN%ef%bc%9a%e4%bd%a0%e5%8e%bb%e5%b0%8f%e5%8d%96%e9%83%a8%e5%8f%96%e8%bf%87%e5%bf%ab%e9%80%92%e4%b9%88%ef%bc%9f.md">20 CDN：你去小卖部取过快递么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 数据中心：我是开发商，自己拿地盖别墅.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/21%20%e6%95%b0%e6%8d%ae%e4%b8%ad%e5%bf%83%ef%bc%9a%e6%88%91%e6%98%af%e5%bc%80%e5%8f%91%e5%95%86%ef%bc%8c%e8%87%aa%e5%b7%b1%e6%8b%bf%e5%9c%b0%e7%9b%96%e5%88%ab%e5%a2%85.md">21 数据中心：我是开发商，自己拿地盖别墅.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 VPN：朝中有人好做官.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/22%20VPN%ef%bc%9a%e6%9c%9d%e4%b8%ad%e6%9c%89%e4%ba%ba%e5%a5%bd%e5%81%9a%e5%ae%98.md">22 VPN：朝中有人好做官.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 移动网络：去巴塞罗那，手机也上不了脸书.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/23%20%e7%a7%bb%e5%8a%a8%e7%bd%91%e7%bb%9c%ef%bc%9a%e5%8e%bb%e5%b7%b4%e5%a1%9e%e7%bd%97%e9%82%a3%ef%bc%8c%e6%89%8b%e6%9c%ba%e4%b9%9f%e4%b8%8a%e4%b8%8d%e4%ba%86%e8%84%b8%e4%b9%a6.md">23 移动网络：去巴塞罗那，手机也上不了脸书.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 云中网络：自己拿地成本高，购买公寓更灵活.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/24%20%e4%ba%91%e4%b8%ad%e7%bd%91%e7%bb%9c%ef%bc%9a%e8%87%aa%e5%b7%b1%e6%8b%bf%e5%9c%b0%e6%88%90%e6%9c%ac%e9%ab%98%ef%bc%8c%e8%b4%ad%e4%b9%b0%e5%85%ac%e5%af%93%e6%9b%b4%e7%81%b5%e6%b4%bb.md">24 云中网络：自己拿地成本高，购买公寓更灵活.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 软件定义网络：共享基础设施的小区物业管理办法.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/25%20%e8%bd%af%e4%bb%b6%e5%ae%9a%e4%b9%89%e7%bd%91%e7%bb%9c%ef%bc%9a%e5%85%b1%e4%ba%ab%e5%9f%ba%e7%a1%80%e8%ae%be%e6%96%bd%e7%9a%84%e5%b0%8f%e5%8c%ba%e7%89%a9%e4%b8%9a%e7%ae%a1%e7%90%86%e5%8a%9e%e6%b3%95.md">25 软件定义网络：共享基础设施的小区物业管理办法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 云中的网络安全：虽然不是土豪，也需要基本安全和保障.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/26%20%e4%ba%91%e4%b8%ad%e7%9a%84%e7%bd%91%e7%bb%9c%e5%ae%89%e5%85%a8%ef%bc%9a%e8%99%bd%e7%84%b6%e4%b8%8d%e6%98%af%e5%9c%9f%e8%b1%aa%ef%bc%8c%e4%b9%9f%e9%9c%80%e8%a6%81%e5%9f%ba%e6%9c%ac%e5%ae%89%e5%85%a8%e5%92%8c%e4%bf%9d%e9%9a%9c.md">26 云中的网络安全：虽然不是土豪，也需要基本安全和保障.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 云中的网络QoS：邻居疯狂下电影，我该怎么办？.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/27%20%e4%ba%91%e4%b8%ad%e7%9a%84%e7%bd%91%e7%bb%9cQoS%ef%bc%9a%e9%82%bb%e5%b1%85%e7%96%af%e7%8b%82%e4%b8%8b%e7%94%b5%e5%bd%b1%ef%bc%8c%e6%88%91%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">27 云中的网络QoS：邻居疯狂下电影，我该怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 云中网络的隔离GRE、VXLAN：虽然住一个小区，也要保护隐私.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/28%20%e4%ba%91%e4%b8%ad%e7%bd%91%e7%bb%9c%e7%9a%84%e9%9a%94%e7%a6%bbGRE%e3%80%81VXLAN%ef%bc%9a%e8%99%bd%e7%84%b6%e4%bd%8f%e4%b8%80%e4%b8%aa%e5%b0%8f%e5%8c%ba%ef%bc%8c%e4%b9%9f%e8%a6%81%e4%bf%9d%e6%8a%a4%e9%9a%90%e7%a7%81.md">28 云中网络的隔离GRE、VXLAN：虽然住一个小区，也要保护隐私.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 容器网络：来去自由的日子，不买公寓去合租.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/29%20%e5%ae%b9%e5%99%a8%e7%bd%91%e7%bb%9c%ef%bc%9a%e6%9d%a5%e5%8e%bb%e8%87%aa%e7%94%b1%e7%9a%84%e6%97%a5%e5%ad%90%ef%bc%8c%e4%b8%8d%e4%b9%b0%e5%85%ac%e5%af%93%e5%8e%bb%e5%90%88%e7%a7%9f.md">29 容器网络：来去自由的日子，不买公寓去合租.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 容器网络之Flannel：每人一亩三分地.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/30%20%e5%ae%b9%e5%99%a8%e7%bd%91%e7%bb%9c%e4%b9%8bFlannel%ef%bc%9a%e6%af%8f%e4%ba%ba%e4%b8%80%e4%ba%a9%e4%b8%89%e5%88%86%e5%9c%b0.md">30 容器网络之Flannel：每人一亩三分地.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 容器网络之Calico：为高效说出善意的谎言.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/31%20%e5%ae%b9%e5%99%a8%e7%bd%91%e7%bb%9c%e4%b9%8bCalico%ef%bc%9a%e4%b8%ba%e9%ab%98%e6%95%88%e8%af%b4%e5%87%ba%e5%96%84%e6%84%8f%e7%9a%84%e8%b0%8e%e8%a8%80.md">31 容器网络之Calico：为高效说出善意的谎言.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 RPC协议综述：远在天边，近在眼前.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/32%20RPC%e5%8d%8f%e8%ae%ae%e7%bb%bc%e8%bf%b0%ef%bc%9a%e8%bf%9c%e5%9c%a8%e5%a4%a9%e8%be%b9%ef%bc%8c%e8%bf%91%e5%9c%a8%e7%9c%bc%e5%89%8d.md">32 RPC协议综述：远在天边，近在眼前.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 基于XML的SOAP协议：不要说NBA，请说美国职业篮球联赛.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/33%20%e5%9f%ba%e4%ba%8eXML%e7%9a%84SOAP%e5%8d%8f%e8%ae%ae%ef%bc%9a%e4%b8%8d%e8%a6%81%e8%af%b4NBA%ef%bc%8c%e8%af%b7%e8%af%b4%e7%be%8e%e5%9b%bd%e8%81%8c%e4%b8%9a%e7%af%ae%e7%90%83%e8%81%94%e8%b5%9b.md">33 基于XML的SOAP协议：不要说NBA，请说美国职业篮球联赛.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 基于JSON的RESTful接口协议：我不关心过程，请给我结果.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/34%20%e5%9f%ba%e4%ba%8eJSON%e7%9a%84RESTful%e6%8e%a5%e5%8f%a3%e5%8d%8f%e8%ae%ae%ef%bc%9a%e6%88%91%e4%b8%8d%e5%85%b3%e5%bf%83%e8%bf%87%e7%a8%8b%ef%bc%8c%e8%af%b7%e7%bb%99%e6%88%91%e7%bb%93%e6%9e%9c.md">34 基于JSON的RESTful接口协议：我不关心过程，请给我结果.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 二进制类RPC协议：还是叫NBA吧，总说全称多费劲.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/35%20%e4%ba%8c%e8%bf%9b%e5%88%b6%e7%b1%bbRPC%e5%8d%8f%e8%ae%ae%ef%bc%9a%e8%bf%98%e6%98%af%e5%8f%abNBA%e5%90%a7%ef%bc%8c%e6%80%bb%e8%af%b4%e5%85%a8%e7%a7%b0%e5%a4%9a%e8%b4%b9%e5%8a%b2.md">35 二进制类RPC协议：还是叫NBA吧，总说全称多费劲.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 跨语言类RPC协议：交流之前，双方先来个专业术语表.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/36%20%e8%b7%a8%e8%af%ad%e8%a8%80%e7%b1%bbRPC%e5%8d%8f%e8%ae%ae%ef%bc%9a%e4%ba%a4%e6%b5%81%e4%b9%8b%e5%89%8d%ef%bc%8c%e5%8f%8c%e6%96%b9%e5%85%88%e6%9d%a5%e4%b8%aa%e4%b8%93%e4%b8%9a%e6%9c%af%e8%af%ad%e8%a1%a8.md">36 跨语言类RPC协议：交流之前，双方先来个专业术语表.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 知识串：用双十一的故事串起碎片的网络协议（上）.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/37%20%e7%9f%a5%e8%af%86%e4%b8%b2%ef%bc%9a%e7%94%a8%e5%8f%8c%e5%8d%81%e4%b8%80%e7%9a%84%e6%95%85%e4%ba%8b%e4%b8%b2%e8%b5%b7%e7%a2%8e%e7%89%87%e7%9a%84%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae%ef%bc%88%e4%b8%8a%ef%bc%89.md">37 知识串：用双十一的故事串起碎片的网络协议（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 知识串：用双十一的故事串起碎片的网络协议（中）.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/38%20%e7%9f%a5%e8%af%86%e4%b8%b2%ef%bc%9a%e7%94%a8%e5%8f%8c%e5%8d%81%e4%b8%80%e7%9a%84%e6%95%85%e4%ba%8b%e4%b8%b2%e8%b5%b7%e7%a2%8e%e7%89%87%e7%9a%84%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae%ef%bc%88%e4%b8%ad%ef%bc%89.md">38 知识串：用双十一的故事串起碎片的网络协议（中）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 知识串：用双十一的故事串起碎片的网络协议（下）.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/39%20%e7%9f%a5%e8%af%86%e4%b8%b2%ef%bc%9a%e7%94%a8%e5%8f%8c%e5%8d%81%e4%b8%80%e7%9a%84%e6%95%85%e4%ba%8b%e4%b8%b2%e8%b5%b7%e7%a2%8e%e7%89%87%e7%9a%84%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae%ef%bc%88%e4%b8%8b%ef%bc%89.md">39 知识串：用双十一的故事串起碎片的网络协议（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 搭建一个网络实验环境：授人以鱼不如授人以渔.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/40%20%e6%90%ad%e5%bb%ba%e4%b8%80%e4%b8%aa%e7%bd%91%e7%bb%9c%e5%ae%9e%e9%aa%8c%e7%8e%af%e5%a2%83%ef%bc%9a%e6%8e%88%e4%ba%ba%e4%bb%a5%e9%b1%bc%e4%b8%8d%e5%a6%82%e6%8e%88%e4%ba%ba%e4%bb%a5%e6%b8%94.md">40 搭建一个网络实验环境：授人以鱼不如授人以渔.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐1 创作故事：我是如何创作“趣谈网络协议”专栏的？.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/%e5%8a%a0%e9%a4%901%20%e5%88%9b%e4%bd%9c%e6%95%85%e4%ba%8b%ef%bc%9a%e6%88%91%e6%98%af%e5%a6%82%e4%bd%95%e5%88%9b%e4%bd%9c%e2%80%9c%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae%e2%80%9d%e4%b8%93%e6%a0%8f%e7%9a%84%ef%bc%9f.md">加餐1 创作故事：我是如何创作“趣谈网络协议”专栏的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="协议专栏特别福利 答疑解惑1期.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/%e5%8d%8f%e8%ae%ae%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e7%a6%8f%e5%88%a9%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%911%e6%9c%9f.md">协议专栏特别福利 答疑解惑1期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="协议专栏特别福利 答疑解惑2期.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/%e5%8d%8f%e8%ae%ae%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e7%a6%8f%e5%88%a9%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%912%e6%9c%9f.md">协议专栏特别福利 答疑解惑2期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="协议专栏特别福利 答疑解惑3期.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/%e5%8d%8f%e8%ae%ae%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e7%a6%8f%e5%88%a9%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%913%e6%9c%9f.md">协议专栏特别福利 答疑解惑3期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="协议专栏特别福利 答疑解惑4期.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/%e5%8d%8f%e8%ae%ae%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e7%a6%8f%e5%88%a9%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%914%e6%9c%9f.md">协议专栏特别福利 答疑解惑4期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="协议专栏特别福利 答疑解惑5期.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/%e5%8d%8f%e8%ae%ae%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e7%a6%8f%e5%88%a9%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%915%e6%9c%9f.md">协议专栏特别福利 答疑解惑5期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 放弃完美主义，执行力就是限时限量认真完成.md" href="/%e4%b8%93%e6%a0%8f/%e8%b6%a3%e8%b0%88%e7%bd%91%e7%bb%9c%e5%8d%8f%e8%ae%ae/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%94%be%e5%bc%83%e5%ae%8c%e7%be%8e%e4%b8%bb%e4%b9%89%ef%bc%8c%e6%89%a7%e8%a1%8c%e5%8a%9b%e5%b0%b1%e6%98%af%e9%99%90%e6%97%b6%e9%99%90%e9%87%8f%e8%ae%a4%e7%9c%9f%e5%ae%8c%e6%88%90.md">结束语 放弃完美主义，执行力就是限时限量认真完成.md</a>
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
                            <h1 id="title" data-id="34 基于JSON的RESTful接口协议：我不关心过程，请给我结果" class="title">34 基于JSON的RESTful接口协议：我不关心过程，请给我结果</h1>
                            <div><p>上一节我们讲了基于XML的SOAP协议，SOAP的S是啥意思来着？是Simple，但是好像一点儿都不简单啊！</p>

<p>你会发现，对于SOAP来讲，无论XML中调用的是什么函数，多是通过HTTP的POST方法发送的。但是咱们原来学HTTP的时候，我们知道HTTP除了POST，还有PUT、DELETE、GET等方法，这些也可以代表一个个动作，而且基本满足增、删、查、改的需求，比如增是POST，删是DELETE，查是GET，改是PUT。</p>

<h2 id="传输协议问题">传输协议问题</h2>

<p>对于SOAP来讲，比如我创建一个订单，用POST，在XML里面写明动作是CreateOrder；删除一个订单，还是用POST，在XML里面写明了动作是DeleteOrder。其实创建订单完全可以使用POST动作，然后在XML里面放一个订单的信息就可以了，而删除用DELETE动作，然后在XML里面放一个订单的ID就可以了。</p>

<p>于是上面的那个SOAP就变成下面这个简单的模样。</p>

<pre><code>POST /purchaseOrder HTTP/1.1
Host: www.geektime.com
Content-Type: application/xml; charset=utf-8
Content-Length: nnn

&lt;?xml version=&quot;1.0&quot;?&gt;
 &lt;order&gt;
     &lt;date&gt;2018-07-01&lt;/date&gt;
      &lt;className&gt;趣谈网络协议&lt;/className&gt;
       &lt;Author&gt;刘超&lt;/Author&gt;
       &lt;price&gt;68&lt;/price&gt;
  &lt;/order&gt;
</code></pre>

<p>而且XML的格式也可以改成另外一种简单的文本化的对象表示格式JSON。</p>

<pre><code>POST /purchaseOrder HTTP/1.1
Host: www.geektime.com
Content-Type: application/json; charset=utf-8
Content-Length: nnn

{
 &quot;order&quot;: {
  &quot;date&quot;: &quot;2018-07-01&quot;,
  &quot;className&quot;: &quot;趣谈网络协议&quot;,
  &quot;Author&quot;: &quot;刘超&quot;,
  &quot;price&quot;: &quot;68&quot;
 }
}
</code></pre>

<p>经常写Web应用的应该已经发现，这就是RESTful格式的API的样子。</p>

<h2 id="协议约定问题">协议约定问题</h2>

<p>然而RESTful可不仅仅是指API，而是一种架构风格，全称Representational State Transfer，表述性状态转移，来自一篇重要的论文《架构风格与基于网络的软件架构设计》（Architectural Styles and the Design of Network-based Software Architectures）。</p>

<p>这篇文章从深层次，更加抽象地论证了一个互联网应用应该有的设计要点，而这些设计要点，成为后来我们能看到的所有高并发应用设计都必须要考虑的问题，再加上REST API比较简单直接，所以后来几乎成为互联网应用的标准接口。</p>

<p>因此，和SOAP不一样，REST不是一种严格规定的标准，它其实是一种设计风格。如果按这种风格进行设计，RESTful接口和SOAP接口都能做到，只不过后面的架构是REST倡导的，而SOAP相对比较关注前面的接口。</p>

<p>而且由于能够通过WSDL生成客户端的Stub，因而SOAP常常被用于类似传统的RPC方式，也即调用远端和调用本地是一样的。</p>

<p>然而本地调用和远程跨网络调用毕竟不一样，这里的不一样还不仅仅是因为有网络而导致的客户端和服务端的分离，从而带来的网络性能问题。更重要的问题是，客户端和服务端谁来维护状态。所谓的状态就是对某个数据当前处理到什么程度了。</p>

<p>这里举几个例子，例如，我浏览到哪个目录了，我看到第几页了，我要买个东西，需要扣减一下库存，这些都是状态。本地调用其实没有人纠结这个问题，因为数据都在本地，谁处理都一样，而且一边处理了，另一边马上就能看到。</p>

<p>当有了RPC之后，我们本来期望对上层透明，就像上一节说的“远在天边，尽在眼前”。于是使用RPC的时候，对于状态的问题也没有太多的考虑。</p>

<p>就像NFS一样，客户端会告诉服务端，我要进入哪个目录，服务端必须要为某个客户端维护一个状态，就是当前这个客户端浏览到哪个目录了。例如，客户端输入cd hello，服务端要在某个地方记住，上次浏览到/root/liuchao了，因而客户的这次输入，应该给它显示/root/liuchao/hello下面的文件列表。而如果有另一个客户端，同样输入cd hello，服务端也在某个地方记住，上次浏览到/var/lib，因而要给客户显示的是/var/lib/hello。</p>

<p>不光NFS，如果浏览翻页，我们经常要实现函数next()，在一个列表中取下一页，但是这就需要服务端记住，客户端A上次浏览到20～30页了，那它调用next()，应该显示30～40页，而客户端B上次浏览到100～110页了，调用next()应该显示110～120页。</p>

<p>上面的例子都是在RPC场景下，由服务端来维护状态，很多SOAP接口设计的时候，也常常按这种模式。这种模式原来没有问题，是因为客户端和服务端之间的比例没有失衡。因为一般不会同时有太多的客户端同时连上来，所以NFS还能把每个客户端的状态都记住。</p>

<p>公司内部使用的ERP系统，如果使用SOAP的方式实现，并且服务端为每个登录的用户维护浏览到报表那一页的状态，由于一个公司内部的人也不会太多，把ERP放在一个强大的物理机上，也能记得过来。</p>

<p>但是互联网场景下，客户端和服务端就彻底失衡了。你可以想象“双十一”，多少人同时来购物，作为服务端，它能记得过来吗？当然不可能，只好多个服务端同时提供服务，大家分担一下。但是这就存在一个问题，服务端怎么把自己记住的客户端状态告诉另一个服务端呢？或者说，你让我给你分担工作，你也要把工作的前因后果给我说清楚啊！</p>

<p>那服务端索性就要想了，既然这么多客户端，那大家就分分工吧。服务端就只记录资源的状态，例如文件的状态，报表的状态，库存的状态，而客户端自己维护自己的状态。比如，你访问到哪个目录了啊，报表的哪一页了啊，等等。</p>

<p>这样对于API也有影响，也就是说，当客户端维护了自己的状态，就不能这样调用服务端了。例如客户端说，我想访问当前目录下的hello路径。服务端说，我怎么知道你的当前路径。所以客户端要先看看自己当前路径是/root/liuchao，然后告诉服务端说，我想访问/root/liuchao/hello路径。</p>

<p>再比如，客户端说我想访问下一页，服务端说，我怎么知道你当前访问到哪一页了。所以客户端要先看看自己访问到了100～110页，然后告诉服务器说，我想访问110～120页。</p>

<p>这就是服务端的无状态化。这样服务端就可以横向扩展了，一百个人一起服务，不用交接，每个人都能处理。</p>

<p>所谓的无状态，其实是服务端维护资源的状态，客户端维护会话的状态。对于服务端来讲，只有资源的状态改变了，客户端才调用POST、PUT、DELETE方法来找我；如果资源的状态没变，只是客户端的状态变了，就不用告诉我了，对于我来说都是统一的GET。</p>

<p>虽然这只改进了GET，但是已经带来了很大的进步。因为对于互联网应用，大多数是读多写少的。而且只要服务端的资源状态不变，就给了我们缓存的可能。例如可以将状态缓存到接入层，甚至缓存到CDN的边缘节点，这都是资源状态不变的好处。</p>

<p>按照这种思路，对于API的设计，就慢慢变成了以资源为核心，而非以过程为核心。也就是说，客户端只要告诉服务端你想让资源状态最终变成什么样就可以了，而不用告诉我过程，不用告诉我动作。</p>

<p>还是文件目录的例子。客户端应该访问哪个绝对路径，而非一个动作，我就要进入某个路径。再如，库存的调用，应该查看当前的库存数目，然后减去购买的数量，得到结果的库存数。这个时候应该设置为目标库存数（但是当前库存数要匹配），而非告知减去多少库存。</p>

<p>这种API的设计需要实现幂等，因为网络不稳定，就会经常出错，因而需要重试，但是一旦重试，就会存在幂等的问题，也就是同一个调用，多次调用的结果应该一样，不能一次支付调用，因为调用三次变成了支付三次。不能进入cd a，做了三次，就变成了cd a/a/a。也不能扣减库存，调用了三次，就扣减三次库存。</p>

<p>当然按照这种设计模式，无论RESTful API还是SOAP API都可以将架构实现成无状态的，面向资源的、幂等的、横向扩展的、可缓存的。</p>

<p>但是SOAP的XML正文中，是可以放任何动作的。例如XML里面可以写&lt; ADD &gt;，&lt; MINUS &gt;等。这就方便使用SOAP的人，将大量的动作放在API里面。</p>

<p>RESTful没这么复杂，也没给客户提供这么多的可能性，正文里的JSON基本描述的就是资源的状态，没办法描述动作，而且能够出发的动作只有CRUD，也即POST、GET、PUT、DELETE，也就是对于状态的改变。</p>

<p>所以，从接口角度，就让你死了这条心。当然也有很多技巧的方法，在使用RESTful API的情况下，依然提供基于动作的有状态请求，这属于反模式了。</p>

<h2 id="服务发现问题">服务发现问题</h2>

<p>对于RESTful API来讲，我们已经解决了传输协议的问题——基于HTTP，协议约定问题——基于JSON，最后要解决的是服务发现问题。</p>

<p>有个著名的基于RESTful API的跨系统调用框架叫Spring Cloud。在Spring Cloud中有一个组件叫 Eureka。传说，阿基米德在洗澡时发现浮力原理，高兴得来不及穿上裤子，跑到街上大喊：“Eureka（我找到了）！”所以Eureka是用来实现注册中心的，负责维护注册的服务列表。</p>

<p>服务分服务提供方，它向Eureka做服务注册、续约和下线等操作，注册的主要数据包括服务名、机器IP、端口号、域名等等。</p>

<p>另外一方是服务消费方，向Eureka获取服务提供方的注册信息。为了实现负载均衡和容错，服务提供方可以注册多个。</p>

<p>当消费方要调用服务的时候，会从注册中心读出多个服务来，那怎么调用呢？当然是RESTful方式了。</p>

<p>Spring Cloud提供一个RestTemplate工具，用于将请求对象转换为JSON，并发起Rest调用，RestTemplate的调用也是分POST、PUT、GET、 DELETE的，当结果返回的时候，根据返回的JSON解析成对象。</p>

<p>通过这样封装，调用起来也很方便。</p>

<h2 id="小结">小结</h2>

<p>好了，这一节就到这里了，我们来总结一下。</p>

<ul>
<li><p>SOAP过于复杂，而且设计是面向动作的，因而往往因为架构问题导致并发量上不去。</p></li>

<li><p>RESTful不仅仅是一个API，而且是一种架构模式，主要面向资源，提供无状态服务，有利于横向扩展应对高并发。</p></li>
</ul>

<p>最后，给你留两个思考题：</p>

<ol>
<li><p>在讨论RESTful模型的时候，举了一个库存的例子，但是这种方法有很大问题，那你知道为什么要这样设计吗？</p></li>

<li><p>基于文本的RPC虽然解决了二进制的问题，但是它本身也有问题，你能举出一些例子吗？</p></li>
</ol>

<p>我们的专栏更新到第34讲，不知你掌握得如何？每节课后我留的思考题，你都有没有认真思考，并在留言区写下答案呢？我会从<strong>已发布的文章中选出一批认真留言的同学</strong>，赠送学习奖励礼券和我整理的独家网络协议知识图谱。</p>

<p>欢迎你留言和我讨论。趣谈网络协议，我们下期见！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#68040404515c5959585f280f05090104460b0705" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f199afaace7cdba',t:'MTczNDEzMTE5NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>