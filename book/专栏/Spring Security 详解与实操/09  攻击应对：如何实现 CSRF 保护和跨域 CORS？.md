<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=09&#32;&#32;攻击应对：如何实现&#32;CSRF&#32;保护和跨域&#32;CORS？>
        <link rel="icon" href="/static/favicon.png">
        <title>09  攻击应对：如何实现 CSRF 保护和跨域 CORS？ </title>
        
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
                        <a class="menu-item" id="00 开篇词  Spring Security，为你的应用安全与职业之路保驾护航.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20Spring%20Security%ef%bc%8c%e4%b8%ba%e4%bd%a0%e7%9a%84%e5%ba%94%e7%94%a8%e5%ae%89%e5%85%a8%e4%b8%8e%e8%81%8c%e4%b8%9a%e4%b9%8b%e8%b7%af%e4%bf%9d%e9%a9%be%e6%8a%a4%e8%88%aa.md">00 开篇词  Spring Security，为你的应用安全与职业之路保驾护航.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  顶级框架：Spring Security 是一款什么样的安全性框架？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/01%20%20%e9%a1%b6%e7%ba%a7%e6%a1%86%e6%9e%b6%ef%bc%9aSpring%20Security%20%e6%98%af%e4%b8%80%e6%ac%be%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84%e5%ae%89%e5%85%a8%e6%80%a7%e6%a1%86%e6%9e%b6%ef%bc%9f.md">01  顶级框架：Spring Security 是一款什么样的安全性框架？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  用户认证：如何使用 Spring Security 构建用户认证体系？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/02%20%20%e7%94%a8%e6%88%b7%e8%ae%a4%e8%af%81%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20Spring%20Security%20%e6%9e%84%e5%bb%ba%e7%94%a8%e6%88%b7%e8%ae%a4%e8%af%81%e4%bd%93%e7%b3%bb%ef%bc%9f.md">02  用户认证：如何使用 Spring Security 构建用户认证体系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  认证体系：如何深入理解 Spring Security 用户认证机制？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/03%20%20%e8%ae%a4%e8%af%81%e4%bd%93%e7%b3%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Spring%20Security%20%e7%94%a8%e6%88%b7%e8%ae%a4%e8%af%81%e6%9c%ba%e5%88%b6%ef%bc%9f.md">03  认证体系：如何深入理解 Spring Security 用户认证机制？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  密码安全：Spring Security 中包含哪些加解密技术？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/04%20%20%e5%af%86%e7%a0%81%e5%ae%89%e5%85%a8%ef%bc%9aSpring%20Security%20%e4%b8%ad%e5%8c%85%e5%90%ab%e5%93%aa%e4%ba%9b%e5%8a%a0%e8%a7%a3%e5%af%86%e6%8a%80%e6%9c%af%ef%bc%9f.md">04  密码安全：Spring Security 中包含哪些加解密技术？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  访问授权：如何对请求的安全访问过程进行有效配置？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/05%20%20%e8%ae%bf%e9%97%ae%e6%8e%88%e6%9d%83%ef%bc%9a%e5%a6%82%e4%bd%95%e5%af%b9%e8%af%b7%e6%b1%82%e7%9a%84%e5%ae%89%e5%85%a8%e8%ae%bf%e9%97%ae%e8%bf%87%e7%a8%8b%e8%bf%9b%e8%a1%8c%e6%9c%89%e6%95%88%e9%85%8d%e7%bd%ae%ef%bc%9f.md">05  访问授权：如何对请求的安全访问过程进行有效配置？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  权限管理：如何剖析 Spring Security 的授权原理？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/06%20%20%e6%9d%83%e9%99%90%e7%ae%a1%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e5%89%96%e6%9e%90%20Spring%20Security%20%e7%9a%84%e6%8e%88%e6%9d%83%e5%8e%9f%e7%90%86%ef%bc%9f.md">06  权限管理：如何剖析 Spring Security 的授权原理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  案例实战：使用 Spring Security 基础功能保护 Web 应用.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/07%20%20%e6%a1%88%e4%be%8b%e5%ae%9e%e6%88%98%ef%bc%9a%e4%bd%bf%e7%94%a8%20Spring%20Security%20%e5%9f%ba%e7%a1%80%e5%8a%9f%e8%83%bd%e4%bf%9d%e6%8a%a4%20Web%20%e5%ba%94%e7%94%a8.md">07  案例实战：使用 Spring Security 基础功能保护 Web 应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  管道过滤：如何基于 Spring Security 过滤器扩展安全性？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/08%20%20%e7%ae%a1%e9%81%93%e8%bf%87%e6%bb%a4%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%20Spring%20Security%20%e8%bf%87%e6%bb%a4%e5%99%a8%e6%89%a9%e5%b1%95%e5%ae%89%e5%85%a8%e6%80%a7%ef%bc%9f.md">08  管道过滤：如何基于 Spring Security 过滤器扩展安全性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  攻击应对：如何实现 CSRF 保护和跨域 CORS？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/09%20%20%e6%94%bb%e5%87%bb%e5%ba%94%e5%af%b9%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%20CSRF%20%e4%bf%9d%e6%8a%a4%e5%92%8c%e8%b7%a8%e5%9f%9f%20CORS%ef%bc%9f.md">09  攻击应对：如何实现 CSRF 保护和跨域 CORS？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  全局方法：如何确保方法级别的安全访问？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/10%20%20%e5%85%a8%e5%b1%80%e6%96%b9%e6%b3%95%ef%bc%9a%e5%a6%82%e4%bd%95%e7%a1%ae%e4%bf%9d%e6%96%b9%e6%b3%95%e7%ba%a7%e5%88%ab%e7%9a%84%e5%ae%89%e5%85%a8%e8%ae%bf%e9%97%ae%ef%bc%9f.md">10  全局方法：如何确保方法级别的安全访问？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  案例实战：使用 Spring Security 高级主题保护 Web 应用.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/11%20%20%e6%a1%88%e4%be%8b%e5%ae%9e%e6%88%98%ef%bc%9a%e4%bd%bf%e7%94%a8%20Spring%20Security%20%e9%ab%98%e7%ba%a7%e4%b8%bb%e9%a2%98%e4%bf%9d%e6%8a%a4%20Web%20%e5%ba%94%e7%94%a8.md">11  案例实战：使用 Spring Security 高级主题保护 Web 应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  开放协议：OAuth2 协议解决的是什么问题？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/12%20%20%e5%bc%80%e6%94%be%e5%8d%8f%e8%ae%ae%ef%bc%9aOAuth2%20%e5%8d%8f%e8%ae%ae%e8%a7%a3%e5%86%b3%e7%9a%84%e6%98%af%e4%bb%80%e4%b9%88%e9%97%ae%e9%a2%98%ef%bc%9f.md">12  开放协议：OAuth2 协议解决的是什么问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  授权体系：如何构建 OAuth2 授权服务器？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/13%20%20%e6%8e%88%e6%9d%83%e4%bd%93%e7%b3%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e6%9e%84%e5%bb%ba%20OAuth2%20%e6%8e%88%e6%9d%83%e6%9c%8d%e5%8a%a1%e5%99%a8%ef%bc%9f.md">13  授权体系：如何构建 OAuth2 授权服务器？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  资源保护：如何基于 OAuth2 协议配置授权过程？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/14%20%20%e8%b5%84%e6%ba%90%e4%bf%9d%e6%8a%a4%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%20OAuth2%20%e5%8d%8f%e8%ae%ae%e9%85%8d%e7%bd%ae%e6%8e%88%e6%9d%83%e8%bf%87%e7%a8%8b%ef%bc%9f.md">14  资源保护：如何基于 OAuth2 协议配置授权过程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  令牌扩展：如何使用 JWT 实现定制化 Token？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/15%20%20%e4%bb%a4%e7%89%8c%e6%89%a9%e5%b1%95%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20JWT%20%e5%ae%9e%e7%8e%b0%e5%ae%9a%e5%88%b6%e5%8c%96%20Token%ef%bc%9f.md">15  令牌扩展：如何使用 JWT 实现定制化 Token？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  案例实战：基于 Spring Security 和 Spring Cloud 构建微服务安全架构.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/16%20%20%e6%a1%88%e4%be%8b%e5%ae%9e%e6%88%98%ef%bc%9a%e5%9f%ba%e4%ba%8e%20Spring%20Security%20%e5%92%8c%20Spring%20Cloud%20%e6%9e%84%e5%bb%ba%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%89%e5%85%a8%e6%9e%b6%e6%9e%84.md">16  案例实战：基于 Spring Security 和 Spring Cloud 构建微服务安全架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  案例实战：基于 Spring Security 和 OAuth2 实现单点登录.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/17%20%20%e6%a1%88%e4%be%8b%e5%ae%9e%e6%88%98%ef%bc%9a%e5%9f%ba%e4%ba%8e%20Spring%20Security%20%e5%92%8c%20OAuth2%20%e5%ae%9e%e7%8e%b0%e5%8d%95%e7%82%b9%e7%99%bb%e5%bd%95.md">17  案例实战：基于 Spring Security 和 OAuth2 实现单点登录.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  技术趋势：如何为 Spring Security 添加响应式编程特性？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/18%20%20%e6%8a%80%e6%9c%af%e8%b6%8b%e5%8a%bf%ef%bc%9a%e5%a6%82%e4%bd%95%e4%b8%ba%20Spring%20Security%20%e6%b7%bb%e5%8a%a0%e5%93%8d%e5%ba%94%e5%bc%8f%e7%bc%96%e7%a8%8b%e7%89%b9%e6%80%a7%ef%bc%9f.md">18  技术趋势：如何为 Spring Security 添加响应式编程特性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  测试驱动：如何基于 Spring Security 测试系统安全性？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/19%20%20%e6%b5%8b%e8%af%95%e9%a9%b1%e5%8a%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%20Spring%20Security%20%e6%b5%8b%e8%af%95%e7%b3%bb%e7%bb%9f%e5%ae%89%e5%85%a8%e6%80%a7%ef%bc%9f.md">19  测试驱动：如何基于 Spring Security 测试系统安全性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 结束语  以终为始，Spring Security 的学习总结.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Security%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e5%ae%9e%e6%93%8d/20%20%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e4%bb%a5%e7%bb%88%e4%b8%ba%e5%a7%8b%ef%bc%8cSpring%20Security%20%e7%9a%84%e5%ad%a6%e4%b9%a0%e6%80%bb%e7%bb%93.md">20 结束语  以终为始，Spring Security 的学习总结.md</a>
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
                            <h1 id="title" data-id="09  攻击应对：如何实现 CSRF 保护和跨域 CORS？" class="title">09  攻击应对：如何实现 CSRF 保护和跨域 CORS？</h1>
                            <div><p>现在我们已经掌握了 Spring Security 提供的多项核心功能，但正如在开篇词中提到的，我们面临的系统安全性问题不止如此。今天我们就来讨论在日常开发过程中常见的两个安全性话题，即 CSRF 和 CORS。这两个缩写名称看似陌生，但和应用程序的每次请求都有关联，Spring Security 对它们也提供了良好的开发支持。</p>

<h3 id="使用-spring-security-提供-csrf-保护">使用 Spring Security 提供 CSRF 保护</h3>

<p>我们先来看 CSRF。CSRF 的全称是 Cross-Site Request Forgery，翻译成中文就是<strong>跨站请求伪造</strong>。那么，究竟什么是跨站请求伪造，面对这个问题我们又该如何应对呢？请继续往下看。</p>

<h4 id="什么是-csrf">什么是 CSRF？</h4>

<p>从安全的角度来讲，你可以将 CSRF 理解为一种攻击手段，即攻击者盗用了你的身份，然后以你的名义向第三方网站发送恶意请求。我们可以使用如下所示的流程图来描述 CSRF：</p>

<p><img src="assets/CioPOWDJ3xCAWERSAACPjO6QDvw503.png" alt="Drawing 0.png" /></p>

<p>CSRF 运行流程图</p>

<p>具体流程如下：</p>

<ul>
<li>用户浏览并登录信任的网站 A，通过用户认证后，会在浏览器中生成针对 A 网站的 Cookie；</li>
<li>用户在没有退出网站 A 的情况下访问网站 B，然后网站 B 向网站 A 发起一个请求；</li>
<li>用户浏览器根据网站 B 的请求，携带 Cookie 访问网站 A；</li>
<li>由于浏览器会自动带上用户的 Cookie，所以网站 A 接收到请求之后会根据用户具备的权限进行访问控制，这样相当于用户本身在访问网站 A，从而网站 B 就达到了模拟用户访问网站 A 的操作过程。</li>
</ul>

<p>显然，从应用程序开发的角度来讲，CSRF 就是系统的一个安全漏洞，这种安全漏洞也在 Web 开发中广泛存在。</p>

<p>基于 CSRF 的工作流程，进行 CSRF 保护的基本思想就是<strong>为系统中的每一个连接请求加上一个随机值</strong>，我们称之为 csrf_token。这样，当用户向网站 A 发送请求时，网站 A 在生成的 Cookie 中就会设置一个 csrf_token 值。而在浏览器发送请求时，提交的表单数据中也有一个隐藏的 csrf_token 值，这样网站 A 接收到请求后，一方面从 Cookie 中提取出 csrf_token，另一方面也从表单提交的数据中获取隐藏的 csrf_token，将两者进行比对，如果不一致就代表这就是一个伪造的请求。</p>

<h4 id="使用-csrffilter">使用 CsrfFilter</h4>

<p>在 Spring Security 中，专门提供了一个 CsrfFilter 来实现对 CSRF 的保护。CsrfFilter 拦截请求，并允许使用 GET、HEAD、TRACE 和 OPTIONS 等 HTTP 方法的请求。而针对 PUT、POST、DELETE 等可能会修改数据的其他请求，CsrfFilter 则希望接<strong>收包含 csrf_token 的消息头</strong>。如果这个消息头不存在或包含不正确的 csrf_token 值，应用程序将拒绝该请求并将响应的状态设置为 403。</p>

<p>看到这里，你可能会问，这个 csrf_token 到底长什么样子呢？其实它本质上就是一个<strong>字符串</strong>。在 Spring Security 中，专门定义了一个 CsrfToken 接口来约定它的格式：</p>

<pre><code>public interface CsrfToken extends Serializable {

 

    //获取消息头名称

    String getHeaderName();

 

    //获取应该包含 Token 的参数名称

    String getParameterName();

	 

	//获取具体的 Token 值

    String getToken();

}
</code></pre>

<p>而在 CsrfFilter 类中，我们也找到了如下所示的针对 CsrfToken 的处理过程：</p>

<pre><code>@Override

protected void doFilterInternal(HttpServletRequest request,

             HttpServletResponse response, FilterChain filterChain)

                     throws ServletException, IOException {

        request.setAttribute(HttpServletResponse.class.getName(), response);

 

        //从 CsrfTokenRepository 中获取 CsrfToken

        CsrfToken csrfToken = this.tokenRepository.loadToken(request);

        final boolean missingToken = csrfToken == null;

 

        //如果找不到 CsrfToken 就生成一个并保存到 CsrfTokenRepository 中

        if (missingToken) {

             csrfToken = this.tokenRepository.generateToken(request);

             this.tokenRepository.saveToken(csrfToken, request, response);

        }

 

        //在请求中添加 CsrfToken

        request.setAttribute(CsrfToken.class.getName(), csrfToken);

        request.setAttribute(csrfToken.getParameterName(), csrfToken);

 

        if (!this.requireCsrfProtectionMatcher.matches(request)) {

             filterChain.doFilter(request, response);

             return;

        }

 

        //从请求中获取 CsrfToken

        String actualToken = request.getHeader(csrfToken.getHeaderName());

        if (actualToken == null) {

             actualToken = request.getParameter(csrfToken.getParameterName());

        }

 

        //如果请求所携带的 CsrfToken 与从 Repository 中获取的不同，则抛出异常

        if (!csrfToken.getToken().equals(actualToken)) {

             if (this.logger.isDebugEnabled()) {

                 this.logger.debug(&quot;Invalid CSRF token found for &quot;

                         + UrlUtils.buildFullRequestUrl(request));

             }

             if (missingToken) {

                 this.accessDeniedHandler.handle(request, response,

                         new MissingCsrfTokenException(actualToken));

             }

             else {

                 this.accessDeniedHandler.handle(request, response,

                         new InvalidCsrfTokenException(csrfToken, actualToken));

             }

             return;

        }

        

        //正常情况下继续执行过滤器链的后续流程

        filterChain.doFilter(request, response);

}
</code></pre>

<p>整个过滤器执行流程还是比较清晰的，基本就是<strong>围绕 CsrfToken 的校验工作</strong>。我们注意到这里引入了一个 CsrfTokenRepository，这个 Repository 组件实现了对 CsrfToken 的存储管理，其中就包含前面提到的专门针对 Cookie 的 CookieCsrfTokenRepository。从 CookieCsrfTokenRepository 中，首先我们能看到一组常量定义，包括针对 CSRF 的 Cookie 名称、参数名称以及消息头名称，如下所示：</p>

<pre><code>static final String DEFAULT_CSRF_COOKIE_NAME = &quot;XSRF-TOKEN&quot;;

static final String DEFAULT_CSRF_PARAMETER_NAME = &quot;_csrf&quot;;

static final String DEFAULT_CSRF_HEADER_NAME = &quot;X-XSRF-TOKEN&quot;;
</code></pre>

<p>CookieCsrfTokenRepository 的 saveToken() 方法也比较简单，就是基于 Cookie 对象进行了 CsrfToken 的设置工作，如下所示：</p>

<pre><code>@Override

public void saveToken(CsrfToken token, HttpServletRequest request,

             HttpServletResponse response) {

        String tokenValue = token == null ? &quot;&quot; : token.getToken();

        Cookie cookie = new Cookie(this.cookieName, tokenValue);

        cookie.setSecure(request.isSecure());

        if (this.cookiePath != null &amp;&amp; !this.cookiePath.isEmpty()) {

                 cookie.setPath(this.cookiePath);

        } else {

                 cookie.setPath(this.getRequestContext(request));

        }

        if (token == null) {

             cookie.setMaxAge(0);

        }

        else {

             cookie.setMaxAge(-1);

        }

        cookie.setHttpOnly(cookieHttpOnly);

        if (this.cookieDomain != null &amp;&amp; !this.cookieDomain.isEmpty()) {

             cookie.setDomain(this.cookieDomain);

        }

 

        response.addCookie(cookie);

}
</code></pre>

<p>在 Spring Security 中，CsrfTokenRepository 接口具有一批实现类，除了 CookieCsrfTokenRepository，还有 HttpSessionCsrfTokenRepository 等，这里不再一一展开。</p>

<p>了解了 CsrfFilter 的基本实现流程，下面我们继续讨论如何使用它来实现 CSRF 保护。从 Spring Security 4.0 开始，默认启用 CSRF 保护，以防止 CSRF 攻击应用程序。Spring Security CSRF 会针对 POST、PUT 和 DELETE 方法进行防护。因此，对于开发人员而言，实际上你并不需要做什么额外工作就能使用这个功能了。当然，如果你不想使用这个功能，也可以通过如下配置方法进行关闭：</p>

<pre><code>http.csrf().disable();
</code></pre>

<h4 id="定制化-csrf-保护">定制化 CSRF 保护</h4>

<p>根据前面的讨论，如果你想获取 HTTP 请求中的 CsrfToken，只需要使用如下所示的代码：</p>

<pre><code>CsrfToken token = (CsrfToken)request.getAttribute(&quot;_csrf&quot;);
</code></pre>

<p>如果你不想使用 Spring Security 内置的存储方式，而是想基于自身需求把 CsrfToken 存储起来，要做的事情就是<strong>实现 CsrfTokenRepository 接口</strong>。这里我们尝试把 CsrfToken 保存到关系型数据库中，所以可以通过扩展 Spring Data 中的 JpaRepository 来定义一个 JpaTokenRepository，如下所示：</p>

<pre><code>public interface JpaTokenRepository extends JpaRepository&lt;Token, Integer&gt; {

 

    Optional&lt;Token&gt; findTokenByIdentifier(String identifier);

}
</code></pre>

<p>JpaTokenRepository 很简单，只有一个根据 identifier 获取 Token 的查询方法，而新增接口则是 JpaRepository 默认提供的，我们可以直接使用。</p>

<p>然后，我们基于 JpaTokenRepository 来构建一个 DatabaseCsrfTokenRepository，如下所示：</p>

<pre><code>public class DatabaseCsrfTokenRepository

        implements CsrfTokenRepository {

 

    @Autowired

    private JpaTokenRepository jpaTokenRepository;

 

    @Override

    public CsrfToken generateToken(HttpServletRequest httpServletRequest) {

        String uuid = UUID.randomUUID().toString();

        return new DefaultCsrfToken(&quot;X-CSRF-TOKEN&quot;, &quot;_csrf&quot;, uuid);

    }

 

    @Override

    public void saveToken(CsrfToken csrfToken, HttpServletRequest httpServletRequest, HttpServletResponse httpServletResponse) {

        String identifier = httpServletRequest.getHeader(&quot;X-IDENTIFIER&quot;);

        Optional&lt;Token&gt; existingToken = jpaTokenRepository.findTokenByIdentifier(identifier);

 

        if (existingToken.isPresent()) {

            Token token = existingToken.get();

            token.setToken(csrfToken.getToken());

        } else {

            Token token = new Token();

            token.setToken(csrfToken.getToken());

            token.setIdentifier(identifier);

            jpaTokenRepository.save(token);

        }

    }

 

    @Override

    public CsrfToken loadToken(HttpServletRequest httpServletRequest) {

        String identifier = httpServletRequest.getHeader(&quot;X-IDENTIFIER&quot;);

        Optional&lt;Token&gt; existingToken = jpaTokenRepository.findTokenByIdentifier(identifier);

 

        if (existingToken.isPresent()) {

            Token token = existingToken.get();

            return new DefaultCsrfToken(&quot;X-CSRF-TOKEN&quot;, &quot;_csrf&quot;, token.getToken());

        }

 

        return null;

    }

}
</code></pre>

<p>DatabaseCsrfTokenRepository 类的代码基本都是自解释的，这里借助了 HTTP 请求中的“X-IDENTIFIER”请求头来确定请求的唯一标识，从而将这一唯一标识与特定的 CsrfToken 关联起来。然后我们使用 JpaTokenRepository 完成了针对关系型数据库的持久化工作。</p>

<p>最后，想要上述代码生效，我们需要通过配置方法完成对 CSRF 的设置，如下所示，这里直接通过 csrfTokenRepository 方法集成了自定义的 DatabaseCsrfTokenRepository：</p>

<hr>

<pre><code>@Override

protected void configure(HttpSecurity http) throws Exception {

        http.csrf(c -&gt; {

            c.csrfTokenRepository(databaseCsrfTokenRepository());

        });

        …

}
</code></pre>

<hr>

<p>作为总结，我们可以用如下所示的示意图来梳理整个定制化 CSRF 所包含的各个组件以及它们之间的关联关系：</p>

<p><img src="assets/Cgp9HWDJ3yyAFoiLAABaQZ0qSQE415.png" alt="Drawing 1.png" /></p>

<p>定制化 CSRF 的相关组件示意图</p>

<h3 id="使用-spring-security-实现-cors">使用 Spring Security 实现 CORS</h3>

<p>介绍完 CSRF，我们继续来看 Web 应用程序开发过程中另一个常见的需求——CORS，即跨域资源共享（Cross-Origin Resource Sharing）。那么问题来了，什么叫跨域？</p>

<h4 id="什么是-cors">什么是 CORS？</h4>

<p>当下的 Web 应用程序开发基本都采用了前后端分离的开发模式，数据的获取并非同源，所以跨域问题在我们日常开发中特别常见。例如，当我们从“test.com”这个域名发起请求时，浏览器为了一定的安全因素考虑，并不会允许请求去访问“api.test.com”这个域名，因为请求已经跨越了两个域名。</p>

<p>请注意，<strong>跨域是浏览器的一种同源安全策略，是浏览器单方面限制的，所以仅在客户端运行在浏览器中才需要考虑这个问题</strong>。从原理上讲，实际就是浏览器在 HTTP 请求的消息头部分新增一些字段，如下所示：</p>

<pre><code>//浏览器自己设置的请求域名

Origin     

//浏览器告诉服务器请求需要用到哪些 HTTP 方法

Access-Control-Request-Method

//浏览器告诉服务器请求需要用到哪些 HTTP 消息头

Access-Control-Request-Headers
</code></pre>

<p>当浏览器进行跨域请求时会和服务器端进行一次的握手协议，从响应结果中可以获取如下信息：</p>

<pre><code>//指定哪些客户端的域名允许访问这个资源

Access-Control-Allow-Origin 

//服务器支持的 HTTP 方法

Access-Control-Allow-Methods 

//需要在正式请求中加入的 HTTP 消息头

Access-Control-Allow-Headers 
</code></pre>

<p>因此，实现 CORS 的关键是<strong>服务器</strong>。只要服务器<strong>合理设置这些响应结果中的消息头</strong>，就相当于实现了对 CORS 的支持，从而支持跨源通信。</p>

<h4 id="使用-corsfilter">使用 CorsFilter</h4>

<p>和 CsrfFilter 注解一样，在 Spring 中也存在一个 CorsFilter 过滤器，不过这个过滤器并不是 Spring Security 提供的，而是来自<strong>Spring Web MVC</strong>。在 CorsFilter 这个过滤器中，首先应该判断来自客户端的请求是不是一个跨域请求，然后根据 CORS 配置来判断该请求是否合法，如下所示：</p>

<pre><code>@Override

protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response,

             FilterChain filterChain) throws ServletException, IOException {

 

        if (CorsUtils.isCorsRequest(request)) {

             CorsConfiguration corsConfiguration = this.configSource.getCorsConfiguration(request);

             if (corsConfiguration != null) {

                 boolean isValid = this.processor.processRequest(corsConfiguration, request, response);

                 if (!isValid || CorsUtils.isPreFlightRequest(request)) {

                     return;

                 }

             }

        }

 

        filterChain.doFilter(request, response);

}
</code></pre>

<p>上述操作的内容是创建合适的配置类 CorsConfiguration。根据 CorsFilter，Spring Security 也在 HttpSecurity 工具类通过提供了 cors() 方法来创建 CorsConfiguration，使用方式如下所示：</p>

<pre><code>@Override

protected void configure(HttpSecurity http) throws Exception {

        http.cors(c -&gt; {

            CorsConfigurationSource source = request -&gt; {

                CorsConfiguration config = new CorsConfiguration();

                config.setAllowedOrigins(Arrays.asList(&quot;*&quot;));

                config.setAllowedMethods(Arrays.asList(&quot;*&quot;));

                return config;

            };

            c.configurationSource(source);

        });

        …

}
</code></pre>

<p>我们可以通过 setAllowedOrigins() 和 setAllowedMethods() 方法实现对 HTTP 响应消息头的设置。<strong>这里将它们都设置成“*”，意味着所有请求都可以进行跨域访问</strong>。你也可以根据需要设置特定的域名和 HTTP 方法。</p>

<h4 id="使用-crossorigin-注解">使用 @CrossOrigin 注解</h4>

<p>通过 CorsFilter，我们实现了全局级别的跨域设置。但有时候，我们可能只需要针对某些请求实现这一功能，通过 Spring Security 也是可以做到这一点的，我们可以在特定的 HTTP 端点上使用如下所示的 @CrossOrigin 注解：</p>

<pre><code>@Controller

public class TestController {

        

    @PostMapping(&quot;/hello&quot;)

	@CrossOrigin(&quot;http://api.test.com:8080&quot;)

    public String hello() {

        return &quot;hello&quot;;

    }

}
</code></pre>

<p>默认情况下，@CrossOrigin 注解允许使用所有的域和消息头，同时会将 Controller 中的方法映射到所有的 HTTP 方法。</p>

<h3 id="小结与预告">小结与预告</h3>

<p>这一讲关注的是对 Web 请求安全性的讨论，我们讨论了日常开发过程中常见的两个概念，即 CSRF 和 CORS。这两个概念有时候容易混淆，但应对的是完全不同的两种场景。</p>

<p>CSRF 是一种攻击行为，所以我们需要对系统进行保护，而 CORS 更多的是一种前后端开发模式上的约定。在 Spring Security 中，针对这两个场景都提供了对应的过滤器，我们只需要通过简单的配置方法就能在系统中自动集成想要的功能。</p>

<p>本讲主要内容如下：</p>

<p><img src="assets/CioPOWDJ3zuAS8aXAADIA10IfrQ469.png" alt="Drawing 2.png" /></p>

<p>最后我想给你留一道思考题：在 Spring Security 中，如何定制化一套对 CsrfToken 的处理机制？欢迎你在留言区和我分享观点。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#87ebebebbeb3b6b6b7b0c7e0eae6eeeba9e4e8ea" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f127ad21b0aede4',t:'MTczNDA1NjQ3Ni4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>