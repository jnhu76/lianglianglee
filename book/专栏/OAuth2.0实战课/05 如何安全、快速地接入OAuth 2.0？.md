<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=05&#32;如何安全、快速地接入OAuth&#32;2.0？>
        <link rel="icon" href="/static/favicon.png">
        <title>05 如何安全、快速地接入OAuth 2.0？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 为什么要学OAuth 2.0？.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e5%ad%a6OAuth%202.0%ef%bc%9f.md">00 开篇词 为什么要学OAuth 2.0？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 OAuth 2.0是要通过什么方式解决什么问题？.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/01%20OAuth%202.0%e6%98%af%e8%a6%81%e9%80%9a%e8%bf%87%e4%bb%80%e4%b9%88%e6%96%b9%e5%bc%8f%e8%a7%a3%e5%86%b3%e4%bb%80%e4%b9%88%e9%97%ae%e9%a2%98%ef%bc%9f.md">01 OAuth 2.0是要通过什么方式解决什么问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 授权码许可类型中，为什么一定要有授权码？.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/02%20%e6%8e%88%e6%9d%83%e7%a0%81%e8%ae%b8%e5%8f%af%e7%b1%bb%e5%9e%8b%e4%b8%ad%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e4%b8%80%e5%ae%9a%e8%a6%81%e6%9c%89%e6%8e%88%e6%9d%83%e7%a0%81%ef%bc%9f.md">02 授权码许可类型中，为什么一定要有授权码？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 授权服务：授权码和访问令牌的颁发流程是怎样的？.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/03%20%e6%8e%88%e6%9d%83%e6%9c%8d%e5%8a%a1%ef%bc%9a%e6%8e%88%e6%9d%83%e7%a0%81%e5%92%8c%e8%ae%bf%e9%97%ae%e4%bb%a4%e7%89%8c%e7%9a%84%e9%a2%81%e5%8f%91%e6%b5%81%e7%a8%8b%e6%98%af%e6%80%8e%e6%a0%b7%e7%9a%84%ef%bc%9f.md">03 授权服务：授权码和访问令牌的颁发流程是怎样的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 在OAuth 2.0中，如何使用JWT结构化令牌？.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/04%20%e5%9c%a8OAuth%202.0%e4%b8%ad%ef%bc%8c%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8JWT%e7%bb%93%e6%9e%84%e5%8c%96%e4%bb%a4%e7%89%8c%ef%bc%9f.md">04 在OAuth 2.0中，如何使用JWT结构化令牌？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 如何安全、快速地接入OAuth 2.0？.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/05%20%e5%a6%82%e4%bd%95%e5%ae%89%e5%85%a8%e3%80%81%e5%bf%ab%e9%80%9f%e5%9c%b0%e6%8e%a5%e5%85%a5OAuth%202.0%ef%bc%9f.md">05 如何安全、快速地接入OAuth 2.0？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 除了授权码许可类型，OAuth 2.0还支持什么授权流程？.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/06%20%e9%99%a4%e4%ba%86%e6%8e%88%e6%9d%83%e7%a0%81%e8%ae%b8%e5%8f%af%e7%b1%bb%e5%9e%8b%ef%bc%8cOAuth%202.0%e8%bf%98%e6%94%af%e6%8c%81%e4%bb%80%e4%b9%88%e6%8e%88%e6%9d%83%e6%b5%81%e7%a8%8b%ef%bc%9f.md">06 除了授权码许可类型，OAuth 2.0还支持什么授权流程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 如何在移动App中使用OAuth 2.0？.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/07%20%e5%a6%82%e4%bd%95%e5%9c%a8%e7%a7%bb%e5%8a%a8App%e4%b8%ad%e4%bd%bf%e7%94%a8OAuth%202.0%ef%bc%9f.md">07 如何在移动App中使用OAuth 2.0？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 实践OAuth 2.0时，使用不当可能会导致哪些安全漏洞？.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/08%20%e5%ae%9e%e8%b7%b5OAuth%202.0%e6%97%b6%ef%bc%8c%e4%bd%bf%e7%94%a8%e4%b8%8d%e5%bd%93%e5%8f%af%e8%83%bd%e4%bc%9a%e5%af%bc%e8%87%b4%e5%93%aa%e4%ba%9b%e5%ae%89%e5%85%a8%e6%bc%8f%e6%b4%9e%ef%bc%9f.md">08 实践OAuth 2.0时，使用不当可能会导致哪些安全漏洞？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 实战：利用OAuth 2.0实现一个OpenID Connect用户身份认证协议..md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/09%20%e5%ae%9e%e6%88%98%ef%bc%9a%e5%88%a9%e7%94%a8OAuth%202.0%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aaOpenID%20Connect%e7%94%a8%e6%88%b7%e8%ba%ab%e4%bb%bd%e8%ae%a4%e8%af%81%e5%8d%8f%e8%ae%ae..md">09 实战：利用OAuth 2.0实现一个OpenID Connect用户身份认证协议..md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 串讲：OAuth 2.0的工作流程与安全问题.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/10%20%e4%b8%b2%e8%ae%b2%ef%bc%9aOAuth%202.0%e7%9a%84%e5%b7%a5%e4%bd%9c%e6%b5%81%e7%a8%8b%e4%b8%8e%e5%ae%89%e5%85%a8%e9%97%ae%e9%a2%98.md">10 串讲：OAuth 2.0的工作流程与安全问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 实战案例：使用Spring Security搭建一套基于JWT的OAuth 2.0架构.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/11%20%e5%ae%9e%e6%88%98%e6%a1%88%e4%be%8b%ef%bc%9a%e4%bd%bf%e7%94%a8Spring%20Security%e6%90%ad%e5%bb%ba%e4%b8%80%e5%a5%97%e5%9f%ba%e4%ba%8eJWT%e7%9a%84OAuth%202.0%e6%9e%b6%e6%9e%84.md">11 实战案例：使用Spring Security搭建一套基于JWT的OAuth 2.0架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 架构案例：基于OAuth 2.0_JWT的微服务参考架构.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/12%20%e6%9e%b6%e6%9e%84%e6%a1%88%e4%be%8b%ef%bc%9a%e5%9f%ba%e4%ba%8eOAuth%202.0_JWT%e7%9a%84%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%8f%82%e8%80%83%e6%9e%b6%e6%9e%84.md">12 架构案例：基于OAuth 2.0_JWT的微服务参考架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 各大开放平台是如何使用OAuth 2.0的？.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/13%20%e5%90%84%e5%a4%a7%e5%bc%80%e6%94%be%e5%b9%b3%e5%8f%b0%e6%98%af%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8OAuth%202.0%e7%9a%84%ef%bc%9f.md">13 各大开放平台是如何使用OAuth 2.0的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 查漏补缺：OAuth 2.0 常见问题答疑.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/14%20%e6%9f%a5%e6%bc%8f%e8%a1%a5%e7%bc%ba%ef%bc%9aOAuth%202.0%20%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91.md">14 查漏补缺：OAuth 2.0 常见问题答疑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 把学习当成一种习惯.md" href="/%e4%b8%93%e6%a0%8f/OAuth2.0%e5%ae%9e%e6%88%98%e8%af%be/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%8a%8a%e5%ad%a6%e4%b9%a0%e5%bd%93%e6%88%90%e4%b8%80%e7%a7%8d%e4%b9%a0%e6%83%af.md">结束语 把学习当成一种习惯.md</a>
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
                            <h1 id="title" data-id="05 如何安全、快速地接入OAuth 2.0？" class="title">05 如何安全、快速地接入OAuth 2.0？</h1>
                            <div><p>　　你好，我是王新栋。</p>

<p>　　在[第 3 讲]，我已经讲了授权服务的流程，如果你还记得的话，当时我特意强调了一点，就是<strong>授权服务将 OAuth 2.0 的复杂性都揽在了自己身上</strong>，这也是授权服务为什么是 OAuth 2.0 体系的核心的原因之一。</p>

<p>　　虽然授权服务做了大部分工作，但是呢，在 OAuth 2.0 的体系里面，除了资源拥有者是作为用户参与，还有另外两个系统角色，也就是第三方软件和受保护资源服务。那么今天这一讲，我们就站在这两个角色的角度，看看它们应该做哪些工作，才能接入到 OAuth 2.0 的体系里面呢？</p>

<p>　　现在，就让我们来看看，作为第三方软件的小兔和京东的受保护资源服务，具体需要着重处理哪些工作吧。</p>

<p>　　注：另外说明一点，为了脱敏的需要，在下面的讲述中，我只是把京东商家开放平台作为一个角色使用，以便有场景感，来帮助你理解。</p>

<h2 id="构建第三方软件应用">构建第三方软件应用</h2>

<p>　　我们先来思考一下：如果要基于京东商家开放平台构建一个小兔打单软件的应用，小兔软件的研发人员应该做哪些工作？</p>

<p>　　是不是要到京东商家开放平台申请注册为开发者，在成为开发者以后再创建一个应用，之后我们就开始开发了，对吧？没错，一定是这样的流程。那么，开发第三方软件应用的过程中，我们需要重点关注哪些内容呢？</p>

<p>　　我先来和你总结下，这些内容包括 4 部分，分别是：<strong>注册信息、引导授权、使用访问令牌、使用刷新令牌。</strong></p>

<p>　　<img src="assets/ee18ea7aab4fbee26cf23d7613801078-20220724223042-7zumqqr.png" alt="" /></p>

<p>　　图1 开发第三方软件应用，应该关注的内容</p>

<p>　　第一点，注册信息。</p>

<p>　　首先，小兔软件只有先有了身份，才可以参与到 OAuth 2.0 的流程中去。也就是说，小兔软件需要先拥有自己的 app_id 和 app_serect 等信息，同时还要填写自己的回调地址 redirect_uri、申请权限等信息。</p>

<p>　　这种方式的注册呢，我们有时候也称它为<strong>静态注册</strong>，也就是小兔软件的研发人员提前登录到京东商家开放平台进行手动注册，以便后续使用这些注册的相关信息来请求访问令牌。</p>

<p>　　第二点，引导授权。</p>

<p>　　当用户需要使用第三方软件，来操作其在受保护资源上的数据，就需要第三方软件来引导授权。比如，小明要使用小兔打单软件来对店铺里面的订单发货打印，那小明首先访问的一定是小兔软件（原则上是直接访问第三方软件，不过我们在后面讲到服务市场这种场景的时候，会有稍微不同），不会是授权服务，更不会是受保护资源服务。</p>

<p>　　但是呢，小兔软件需要小明的授权，只有授权服务才能允许小明这样做。所以呢，小兔软件需要 “配合” 小明做的第一件事儿，就是将小明引导至授权服务，如下面代码所示。</p>

<p>　　那去做什么呢？其实就是让用户为第三方软件授权，得到了授权之后，第三方软件才可以代表用户去访问数据。也就是说，小兔打单软件获得授权之后，才能够代表小明处理其在京东店铺上的订单数据。</p>

<pre><code class="language-java">　　String oauthUrl = &quot;http://localhost:8081/OauthServlet-ch03?reqType=oauth&quot;;
　　response.sendRedirect(toOauthUrl);
</code></pre>

<p>　　第三点，使用访问令牌。</p>

<p>　　<strong>拿到令牌后去使用令牌，才是第三方软件的最终目的</strong>。然后我们看看如何使用令牌。目前 OAuth 2.0 的令牌只支持一种类型，那就是 bearer 令牌，也就是我之前讲到的可以是任意字符串格式的令牌。</p>

<p>　　官方规范给出的使用访问令牌请求的方式，有三种，分别是：</p>

<pre><code class="language-java">　　Form-Encoded Body Parameter（表单参数）
　　POST /resource HTTP/1.1
　　Host: server.example.com
　　Content-Type: application/x-www-form-urlencoded
　　access_token=b1a64d5c-5e0c-4a70-9711-7af6568a61fb
　　
　　URI Query Parameter（URI 查询参数）
　　GET /resource?access_token=b1a64d5c-5e0c-4a70-9711-7af6568a61fb HTTP/1.1
　　Host: server.example.com
　　Authorization Request Header Field（授权请求头部字段）
　　GET /resource HTTP/1.1
　　Host: server.example.com
　　Authorization: Bearer b1a64d5c-5e0c-4a70-9711-7af6568a61fb
</code></pre>

<p>　　也就是说，这三种方式都可以请求到受保护资源服务。那么，我们采用哪种方式最合适呢？</p>

<p>　　根据 OAuth 2.0 的官方建议，系统在接入 OAuth 2.0 之前信息传递的请求载体是 JSON 格式的，那么如果继续采用表单参数提交的方式，令牌就无法加入进去了，因为格式不符。如果这时采用参数传递的方式呢，整个 URI 会被整体复制，安全性是最差的。而请求头部字段的方式就没有上述的这些“烦恼”，因此官方的建议是采用 Authorization 的方式来传递令牌。</p>

<p>　　但是，<strong>我建议你采用表单提交，也就是 POST 的方式来提交令牌，</strong>类似如下代码所示。原因是这样的，从官方的建议中也可以看出，它指的是在接入 OAuth 2.0 之前，如果你已经采用了 JSON 数据格式请求体的情况下，不建议使用表单提交。但是，刚开始的时候，只要三方软件和平台之间约束好了，大家一致采用表单提交，就没有任何问题了。<strong>因为表单提交的方式在保证安全传输的同时，还不需要去额外处理 Authorization 头部信息</strong>。</p>

<pre><code class="language-java">　　String protectedURl=&quot;http://localhost:8082/ProtectedServlet-ch03&quot;;
　　Map&lt;String, String&gt; paramsMap = new HashMap&lt;String, String&gt;();
　　paramsMap.put(&quot;app_id&quot;,&quot;APPID_RABBIT&quot;);
　　paramsMap.put(&quot;app_secret&quot;,&quot;APPSECRET_RABBIT&quot;);
　　paramsMap.put(&quot;token&quot;,accessToken);
　　String result = HttpURLClient.doPost(protectedURl,HttpURLClient.mapToStr(paramsMap));
</code></pre>

<p>　　第四点，使用刷新令牌。</p>

<p>　　我在讲授权服务的时候提到过，如果访问令牌过期了，小兔软件总不能立马提示并让小明重新授权一次，否则小明的体验将会非常不好。为了解决这个问题呢，就用到了刷新令牌。</p>

<p>　　使用刷新令牌的方式跟使用访问令牌是一样的，具体可以参照上面我们讲的访问令牌的方式。关于刷新令牌的使用，你最需要关心的是，什么时候你会来决定使用刷新令牌。</p>

<p>　　在小兔打单软件收到访问令牌的同时，也会收到访问令牌的过期时间 expires_in。一个设计良好的第三方应用，<strong>应该将 expires_in 值保存下来并定时检测</strong>；如果发现 expires_in 即将过期，则需要利用 refresh_token 去重新请求授权服务，以便获取新的、有效的访问令牌。</p>

<p>　　这种定时检测的方法可以提前发现访问令牌是否即将过期。此外，还有一种方法是“现场”发现。也就是说，比如小兔软件访问小明店铺订单的时候，突然收到一个访问令牌失效的响应，此时小兔软件立即使用 refresh_token 来请求一个访问令牌，以便继续代表小明使用他的数据。</p>

<p>　　综合来看的话，定时检测的方式，需要我们额外开发一个定时任务；而“现场”发现，就没有这种额外的工作量啦。具体采用哪一种方式，你可以结合自己的实际情况。不过呢，我还是建议你采用定时检测这种方式，因为它可以带来“提前量”，以便让我们有更好的主动性，而现场发现就有点被动了。</p>

<p>　　说到这里，我要再次提醒你注意的是，<strong>刷新令牌是一次性的，使用之后就会失效</strong>，但是它的有效期会比访问令牌要长。这个时候我们可能会想到，如果刷新令牌也过期了怎么办？在这种情况下，我们就需要将刷新令牌和访问令牌都放弃，相当于回到了系统的初始状态，只能让用户小明重新授权了。</p>

<p>　　到这里，我们来总结下，在构建第三方应用时，你需要重点关注的就是注册、授权、访问令牌、刷新令牌。只要你掌握了这四部分内容，在类似京东这样的开放平台上开发小兔软件，就不再是什么困难的事情了。</p>

<h2 id="服务市场中的第三方应用软件">服务市场中的第三方应用软件</h2>

<p>　　在构建第三方应用的引导授权时，我们说用户第一次“触摸”到的一定是第三方软件，但这并不是绝对的。这个不绝对，就发生在服务市场这样的场景里。</p>

<p>　　那什么是服务市场呢？说白了，就是你开发的软件，比如小兔打单软件、店铺装修软件等，都发布到这样一个“市场”里面售卖。这样，当用户购买了这些软件之后，就可以在服务市场里面看到有个“立即使用”的按钮。点击这个按钮，用户就可以直接访问自己购买的第三方软件了。</p>

<p>　　比如，京东的京麦服务市场里有个“我的服务”目录，里面就存放了我购买的打单软件。小明就可以直接点击“立即使用”，继而进入小兔打单软件，如下图所示。</p>

<p>　　<img src="assets/140a4efb622e21b21fcc4ff57653a915-20220724223042-tnhegu1.png" alt="" /></p>

<p>　　图2 京麦服务市场“我的服务”</p>

<p>　　那么，这里需要注意的是，作为第三方开发者来构建第三方软件的时候，在授权码环节除了要接收授权码 code 值之外，还要接收用户的订购相关信息，比如服务的版本号、服务代码标识等信息。</p>

<p>　　好了，以上就是关于构建第三方软件时需要注意的一些细节问题了。接下来，我们再谈谈构建受保护资源服务的时候，又需要重点处理哪些工作呢。</p>

<h2 id="构建受保护资源服务">构建受保护资源服务</h2>

<p>　　你先想一想，实际上在整个开放授权的环境中，受保护资源最终指的还是 Web API，比如说，访问头像的 API、访问昵称的 API。对应到我们的打单软件中，受保护资源就是订单查询 API、批量查询 API 等。</p>

<p>　　在互联网上的系统之间的通信，基本都是以 Web API 为载体的形式进行。因此呢，当我们说受保护资源被授权服务保护着时，实际上说的是授权服务最终保护的是这些 Web API。我们在构建受保护资源服务的时候，除了基本的要检查令牌的合法性，还需要做些什么呢？我认为<strong>最重要的就是权限范围了。</strong></p>

<p>　　在我们处理受保护资源服务中的逻辑的时候，校验权限的处理会占据很大的比重。你想啊，访问令牌递过来，你肯定要多看看令牌到底能操作哪些功能、又能访问哪些数据吧。现在，我们把这些权限的类别总结归纳下来，最常见的大概有以下几类。</p>

<p>　　<img src="assets/e7b134686b9f2e824ffa8410e20f59f6-20220724223042-akcl1dc.jpg" alt="" /></p>

<p>　　图3 3类权限类别</p>

<p>　　接下来，我和你具体说说这些权限是如何使用的。</p>

<p>　　这里的操作，其实对应的是 Web API，比如目前京东商家开放平台提供有查询商品 API、新增商品 API、删除商品 API 这三种。如果小兔软件请求过来的一个访问令牌 access_token 的 scope 权限范围只对应了查询商品 API、新增商品 API，那么包含这个 access_token 值的请求，就不能执行删除商品 API 的操作。</p>

<pre><code class="language-java">　　String[] scope = OauthServlet.tokenScopeMap.get(accessToken);
　　StringBuffer sbuf = new StringBuffer();
　　for(int i=0;i&lt;scope.length;i++){
　　    sbuf.append(scope[i]).append(&quot;|&quot;);
　　}
　　if(sbuf.toString().indexOf(&quot;query&quot;)&gt;0){
　　    queryGoods(&quot;&quot;);
　　}
　　if(sbuf.toString().indexOf(&quot;add&quot;)&gt;0){
　　    addGoods(&quot;&quot;);
　　}
　　if(sbuf.toString().indexOf(&quot;del&quot;)&gt;0){
　　    delGoods(&quot;&quot;);
　　}
</code></pre>

<p>　　这里的数据，就是指某一个 API 里包含的属性字段信息。比如，有一个查询小明信息的 API，返回的信息包括 Contact（email、phone、qq）、Like（Basketball、Swimming）、Personal Data（sex、age、nickname）。如果小兔软件请求过来的一个访问令牌 access_token 的 scope 权限范围只对应了 Personal Data，那么包含该 access_token 值的请求就不能获取到 Contact 和 Like 的信息，关于这部分的代码，实际跟不同权限对应不同操作的代码类似。</p>

<p>　　看到这里，你就明白了，这种权限范围的粒度要比“不同的权限对应不同的操作”的粒度要小。这正是遵循了最小权限范围原则。</p>

<p>　　这种权限是什么意思呢？其实，这种权限实际上只是换了一种维度，将其定位到了用户上面。</p>

<p>　　一些基础类信息，比如获取地理位置、获取天气预报等，不会带有用户归属属性，也就是说这些信息并不归属于某个用户，是一类公有信息。对于这样的信息，平台提供出去的 API 接口都是“中性”的，没有用户属性。</p>

<p>　　但是，更多的场景却是基于用户属性的。还是以小兔打单软件为例，商家每次打印物流面单的时候，小兔打单软件都要知道是哪个商家的订单。这种情况下，商家为小兔软件授权，小兔软件获取的 access_token 实际上就包含了商家这个用户属性。</p>

<p>　　京东商家开放平台的受保护资源服务每次接收到小兔软件的请求时，都会根据该请求中 access_token 的值找到对应的商家 ID，继而根据商家 ID 查询到商家的订单信息，也就是不同的商家对应不同的订单数据。</p>

<pre><code class="language-java">　　String user = OauthServlet.tokenMap.get(accessToken);
　　queryOrders(user);
</code></pre>

<p>　　在上面讲三种权限的时候，我举的例子实际上都属于一个系统提供了查询、添加、删除这样的所有服务。此时你可能会想到，现在的系统不已经是分布式系统环境了么，如果有很多个受保护资源服务，比如提供用户信息查询的用户资源服务、提供商品查询的商品资源服务、提供订单查询的订单资源服务，那么每个受保护资源服务岂不是都要把上述的权限范围校验执行一遍吗，这样不就会有大量的重复工作产生么？</p>

<p>　　在这里，我特别高兴你能想到这一点。为了应对这种情况，我们应该有一个统一的网关层来处理这样的校验，所有的请求都会经过 API GATEWAY 跳转到不同的受保护资源服务。这样呢，我们就不需要在每一个受保护资源服务上都做一遍权限校验的工作了，而只需要在 API GATEWAY 这一层做权限校验就可以了。系统结构如下图所示。</p>

<p>　　<img src="assets/a5175438e76411c808dd5e72d3d3dbb0-20220724223042-jii2953.png" alt="" /></p>

<p>　　图4 由统一的网关层处理权限校验</p>

<h2 id="总结">总结</h2>

<p>　　截止到这一讲呢，我们已经把 OAuth 2.0 中授权码相关的流程所涉及到的内容都讲完了。通过 02 到 05 这 4 讲，你可以很清晰地理解授权码流程的核心原理了，也可以弄清楚如何使用以及如何接入这一授权流程了。</p>

<p>　　我在本讲开始的时候，提到 OAuth 2.0 的复杂性实际上都给了授权服务来承担，接着我从第三方软件和受保护资源的角度，分别介绍了这两部分系统在接入 OAuth 2.0 的时候应该注意哪些方面。总结下来，我其实希望你能够记住以下两点。</p>

<p>　　对于第三方软件，比如小兔打单软件来讲，<strong>它的主要目的就是获取访问令牌，使用访问令牌</strong>，这当然也是整个 OAuth 2.0 的目的，就是让第三方软件来做这两件事。在这个过程中需要强调的是，第三方软件在使用访问令牌的时候有三种方式，我们建议在平台和第三方软件约定好的前提下，<strong>优先采用 Post 表单提交的方式</strong>。</p>

<p>　　受保护资源系统，比如小兔软件要访问开放平台的订单数据服务，它需要注意的是权限的问题，这个权限范围主要包括，<strong>不同的权限会有不同的操作，不同的权限也会对应不同的数据，不同的用户也会对应不同的数据</strong>。</p>

<h2 id="思考题">思考题</h2>

<p>　　如果使用刷新令牌 refresh_token 请求回来一个新的访问令牌 access_token，按照一般规则授权服务上旧的访问令牌应该要立即失效，但是如果在这之前已经有使用旧的访问令牌发出去的请求，不就受到影响了吗，这种情况下应该如何处理呢？</p>

<p>　　欢迎你在留言区分享你的观点，也欢迎你把今天的内容分享给其他朋友，我们一起交流。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#721e1e1e4b464343424532151f131b1e5c111d1f" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1187934f7f94f1',t:'MTczNDA0NjUxMy4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>