<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=09&#32;Spring&#32;Web&#32;URL&#32;解析常见错误>
        <link rel="icon" href="/static/favicon.png">
        <title>09 Spring Web URL 解析常见错误 </title>
        
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
                        <a class="menu-item" id="00 导读 5分钟轻松了解Spring基础知识.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/00%20%e5%af%bc%e8%af%bb%205%e5%88%86%e9%92%9f%e8%bd%bb%e6%9d%be%e4%ba%86%e8%a7%a3Spring%e5%9f%ba%e7%a1%80%e7%9f%a5%e8%af%86.md">00 导读 5分钟轻松了解Spring基础知识.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 开篇词 贴心“保姆”Spring罢工了怎么办？.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e8%b4%b4%e5%bf%83%e2%80%9c%e4%bf%9d%e5%a7%86%e2%80%9dSpring%e7%bd%a2%e5%b7%a5%e4%ba%86%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">00 开篇词 贴心“保姆”Spring罢工了怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 Spring Bean 定义常见错误.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/01%20Spring%20Bean%20%e5%ae%9a%e4%b9%89%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af.md">01 Spring Bean 定义常见错误.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 Spring Bean 依赖注入常见错误（上）.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/02%20Spring%20Bean%20%e4%be%9d%e8%b5%96%e6%b3%a8%e5%85%a5%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%ef%bc%88%e4%b8%8a%ef%bc%89.md">02 Spring Bean 依赖注入常见错误（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 Spring Bean 依赖注入常见错误（下）.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/03%20Spring%20Bean%20%e4%be%9d%e8%b5%96%e6%b3%a8%e5%85%a5%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%ef%bc%88%e4%b8%8b%ef%bc%89.md">03 Spring Bean 依赖注入常见错误（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 Spring Bean 生命周期常见错误.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/04%20Spring%20Bean%20%e7%94%9f%e5%91%bd%e5%91%a8%e6%9c%9f%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af.md">04 Spring Bean 生命周期常见错误.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 Spring AOP 常见错误（上）.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/05%20Spring%20AOP%20%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%ef%bc%88%e4%b8%8a%ef%bc%89.md">05 Spring AOP 常见错误（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 Spring AOP 常见错误（下）.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/06%20Spring%20AOP%20%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%ef%bc%88%e4%b8%8b%ef%bc%89.md">06 Spring AOP 常见错误（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 Spring事件常见错误.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/07%20Spring%e4%ba%8b%e4%bb%b6%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af.md">07 Spring事件常见错误.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 答疑现场：Spring Core 篇思考题合集.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/08%20%e7%ad%94%e7%96%91%e7%8e%b0%e5%9c%ba%ef%bc%9aSpring%20Core%20%e7%af%87%e6%80%9d%e8%80%83%e9%a2%98%e5%90%88%e9%9b%86.md">08 答疑现场：Spring Core 篇思考题合集.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 Spring Web URL 解析常见错误.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/09%20Spring%20Web%20URL%20%e8%a7%a3%e6%9e%90%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af.md">09 Spring Web URL 解析常见错误.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 Spring Web Header 解析常见错误.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/10%20Spring%20Web%20Header%20%e8%a7%a3%e6%9e%90%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af.md">10 Spring Web Header 解析常见错误.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 Spring Web Body 转化常见错误.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/11%20Spring%20Web%20Body%20%e8%bd%ac%e5%8c%96%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af.md">11 Spring Web Body 转化常见错误.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 Spring Web 参数验证常见错误.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/12%20Spring%20Web%20%e5%8f%82%e6%95%b0%e9%aa%8c%e8%af%81%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af.md">12 Spring Web 参数验证常见错误.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 Spring Web 过滤器使用常见错误（上）.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/13%20Spring%20Web%20%e8%bf%87%e6%bb%a4%e5%99%a8%e4%bd%bf%e7%94%a8%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%ef%bc%88%e4%b8%8a%ef%bc%89.md">13 Spring Web 过滤器使用常见错误（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 Spring Web 过滤器使用常见错误（下）.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/14%20Spring%20Web%20%e8%bf%87%e6%bb%a4%e5%99%a8%e4%bd%bf%e7%94%a8%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%ef%bc%88%e4%b8%8b%ef%bc%89.md">14 Spring Web 过滤器使用常见错误（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 Spring Security 常见错误.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/15%20Spring%20Security%20%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af.md">15 Spring Security 常见错误.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 Spring Exception 常见错误.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/16%20Spring%20Exception%20%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af.md">16 Spring Exception 常见错误.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 答疑现场：Spring Web 篇思考题合集.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/17%20%e7%ad%94%e7%96%91%e7%8e%b0%e5%9c%ba%ef%bc%9aSpring%20Web%20%e7%af%87%e6%80%9d%e8%80%83%e9%a2%98%e5%90%88%e9%9b%86.md">17 答疑现场：Spring Web 篇思考题合集.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 Spring Data 常见错误.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/18%20Spring%20Data%20%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af.md">18 Spring Data 常见错误.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 Spring 事务常见错误（上）.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/19%20Spring%20%e4%ba%8b%e5%8a%a1%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%ef%bc%88%e4%b8%8a%ef%bc%89.md">19 Spring 事务常见错误（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 Spring 事务常见错误（下）.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/20%20Spring%20%e4%ba%8b%e5%8a%a1%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%ef%bc%88%e4%b8%8b%ef%bc%89.md">20 Spring 事务常见错误（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 Spring Rest Template 常见错误.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/21%20Spring%20Rest%20Template%20%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af.md">21 Spring Rest Template 常见错误.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 Spring Test 常见错误.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/22%20Spring%20Test%20%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af.md">22 Spring Test 常见错误.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 答疑现场：Spring 补充篇思考题合集.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/23%20%e7%ad%94%e7%96%91%e7%8e%b0%e5%9c%ba%ef%bc%9aSpring%20%e8%a1%a5%e5%85%85%e7%af%87%e6%80%9d%e8%80%83%e9%a2%98%e5%90%88%e9%9b%86.md">23 答疑现场：Spring 补充篇思考题合集.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="导读 5分钟轻松了解一个HTTP请求的处理过程.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/%e5%af%bc%e8%af%bb%205%e5%88%86%e9%92%9f%e8%bd%bb%e6%9d%be%e4%ba%86%e8%a7%a3%e4%b8%80%e4%b8%aaHTTP%e8%af%b7%e6%b1%82%e7%9a%84%e5%a4%84%e7%90%86%e8%bf%87%e7%a8%8b.md">导读 5分钟轻松了解一个HTTP请求的处理过程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="知识回顾 系统梳理Spring编程错误根源.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/%e7%9f%a5%e8%af%86%e5%9b%9e%e9%a1%be%20%e7%b3%bb%e7%bb%9f%e6%a2%b3%e7%90%86Spring%e7%bc%96%e7%a8%8b%e9%94%99%e8%af%af%e6%a0%b9%e6%ba%90.md">知识回顾 系统梳理Spring编程错误根源.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 问题总比解决办法多.md" href="/%e4%b8%93%e6%a0%8f/Spring%e7%bc%96%e7%a8%8b%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af50%e4%be%8b/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e9%97%ae%e9%a2%98%e6%80%bb%e6%af%94%e8%a7%a3%e5%86%b3%e5%8a%9e%e6%b3%95%e5%a4%9a.md">结束语 问题总比解决办法多.md</a>
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
                            <h1 id="title" data-id="09 Spring Web URL 解析常见错误" class="title">09 Spring Web URL 解析常见错误</h1>
                            <div><p>你好，我是傅健。</p>

<p>上一章节我们讲解了各式各样的错误案例，这些案例都是围绕 Spring 的核心功能展开的，例如依赖注入、AOP 等诸多方面。然而，从现实情况来看，在使用上，我们更多地是使用 Spring 来构建一个 Web 服务，所以从这节课开始，我们会重点解析在 Spring Web 开发中经常遇到的一些错误，帮助你规避这些问题。</p>

<p>不言而喻，这里说的 Web 服务就是指使用 HTTP 协议的服务。而对于 HTTP 请求，首先要处理的就是 URL，所以今天我们就先来介绍下，在 URL 的处理上，Spring 都有哪些经典的案例。闲话少叙，下面我们直接开始演示吧。</p>

<h2 id="案例-1-当-pathvariable-遇到">案例 1：当@PathVariable 遇到 /</h2>

<p>在解析一个 URL 时，我们经常会使用 @PathVariable 这个注解。例如我们会经常见到如下风格的代码：</p>

<pre><code>@RestController
@Slf4j
public class HelloWorldController {
    @RequestMapping(path = &quot;/hi1/{name}&quot;, method = RequestMethod.GET)
    public String hello1(@PathVariable(&quot;name&quot;) String name){
        return name;
        
    };  
}
</code></pre>

<p>当我们使用 <a href="http://localhost:8080/hi1/xiaoming" target="_blank">http://localhost:8080/hi1/xiaoming</a> 访问这个服务时，会返回&rdquo;xiaoming&rdquo;，即 Spring 会把 name 设置为 URL 中对应的值。</p>

<p>看起来顺风顺水，但是假设这个 name 中含有特殊字符/时（例如<a href="http://localhost:8080/hi1/xiaoming" target="_blank">http://localhost:8080/hi1/xiao/ming</a> ），会如何？如果我们不假思索，或许答案是&rdquo;xiao/ming&rdquo;？然而稍微敏锐点的程序员都会判定这个访问是会报错的，具体错误参考：</p>

<p><img src="assets/da282ca3e0c942c6a124b3672d9962b1.jpg" alt="" /></p>

<p>如图所示，当 name 中含有/，这个接口不会为 name 获取任何值，而是直接报Not Found错误。当然这里的“找不到”并不是指name找不到，而是指服务于这个特殊请求的接口。</p>

<p>实际上，这里还存在另外一种错误，即当 name 的字符串以/结尾时，/会被自动去掉。例如我们访问 <a href="http://localhost:8080/hi1/xiaoming/" target="_blank">http://localhost:8080/hi1/xiaoming/</a>，Spring 并不会报错，而是返回xiaoming。</p>

<p>针对这两种类型的错误，应该如何理解并修正呢？</p>

<h3 id="案例解析">案例解析</h3>

<p>实际上，这两种错误都是 URL 匹配执行方法的相关问题，所以我们有必要先了解下 URL 匹配执行方法的大致过程。参考 AbstractHandlerMethodMapping#lookupHandlerMethod：</p>

<pre><code>@Nullable
protected HandlerMethod lookupHandlerMethod(String lookupPath, HttpServletRequest request) throws Exception {
   List&lt;Match&gt; matches = new ArrayList&lt;&gt;();
   //尝试按照 URL 进行精准匹配
   List&lt;T&gt; directPathMatches = this.mappingRegistry.getMappingsByUrl(lookupPath);
   if (directPathMatches != null) {
      //精确匹配上，存储匹配结果
      addMatchingMappings(directPathMatches, matches, request);
   }
   if (matches.isEmpty()) {
      //没有精确匹配上，尝试根据请求来匹配
      addMatchingMappings(this.mappingRegistry.getMappings().keySet(), matches, request);
   }

   if (!matches.isEmpty()) {
      Comparator&lt;Match&gt; comparator = new MatchComparator(getMappingComparator(request));
      matches.sort(comparator);
      Match bestMatch = matches.get(0);
      if (matches.size() &gt; 1) {
         //处理多个匹配的情况
      }
      //省略其他非关键代码
      return bestMatch.handlerMethod;
   }
   else {
      //匹配不上，直接报错
      return handleNoMatch(this.mappingRegistry.getMappings().keySet(), lookupPath, request);
   }
</code></pre>

<p>大体分为这样几个基本步骤。</p>

<p><strong>1. 根据 Path 进行精确匹配</strong></p>

<p>这个步骤执行的代码语句是&rdquo;this.mappingRegistry.getMappingsByUrl(lookupPath)&ldquo;，实际上，它是查询 MappingRegistry#urlLookup，它的值可以用调试视图查看，如下图所示：</p>

<p><img src="assets/e5228329c945474fa388b57f65e7c22f.jpg" alt="" /></p>

<p>查询 urlLookup 是一个精确匹配 Path 的过程。很明显，<a href="http://localhost:8080/hi1/xiaoming" target="_blank">http://localhost:8080/hi1/xiao/ming</a> 的 lookupPath 是&rdquo;/hi1/xiao/ming&rdquo;，并不能得到任何精确匹配。这里需要补充的是，&rdquo;/hi1/{name}&ldquo;这种定义本身也没有出现在 urlLookup 中。</p>

<p><strong>2. 假设 Path 没有精确匹配上，则执行模糊匹配</strong></p>

<p>在步骤 1 匹配失败时，会根据请求来尝试模糊匹配，待匹配的匹配方法可参考下图：</p>

<p><img src="assets/658f673383934de990aa676c755d9072.jpg" alt="" /></p>

<p>显然，&rdquo;/hi1/{name}&ldquo;这个匹配方法已经出现在待匹配候选中了。具体匹配过程可以参考方法 RequestMappingInfo#getMatchingCondition：</p>

<pre><code>public RequestMappingInfo getMatchingCondition(HttpServletRequest request) {
   RequestMethodsRequestCondition methods = this.methodsCondition.getMatchingCondition(request);
   if (methods == null) {
      return null;
   }
   ParamsRequestCondition params = this.paramsCondition.getMatchingCondition(request);
   if (params == null) {
      return null;
   }
   //省略其他匹配条件
   PatternsRequestCondition patterns = this.patternsCondition.getMatchingCondition(request);
   if (patterns == null) {
      return null;
   }
   //省略其他匹配条件
   return new RequestMappingInfo(this.name, patterns,
         methods, params, headers, consumes, produces, custom.getCondition());
}
</code></pre>

<p>现在我们知道<strong>匹配会查询所有的信息</strong>，例如 Header、Body 类型以及URL 等。如果有一项不符合条件，则不匹配。</p>

<p>在我们的案例中，当使用 <a href="http://localhost:8080/hi1/xiaoming" target="_blank">http://localhost:8080/hi1/xiaoming</a> 访问时，其中 patternsCondition 是可以匹配上的。实际的匹配方法执行是通过 AntPathMatcher#match 来执行，判断的相关参数可参考以下调试视图：</p>

<p><img src="assets/351194df4a6c441b86408010aea66437.jpg" alt="" /></p>

<p>但是当我们使用 <a href="http://localhost:8080/hi1/xiaoming" target="_blank">http://localhost:8080/hi1/xiao/ming</a> 来访问时，AntPathMatcher 执行的结果是&rdquo;/hi1/xiao/ming&rdquo;匹配不上&rdquo;/hi1/{name}&ldquo;。</p>

<p><strong>3. 根据匹配情况返回结果</strong></p>

<p>如果找到匹配的方法，则返回方法；如果没有，则返回 null。</p>

<p>在本案例中，<a href="http://localhost:8080/hi1/xiaoming" target="_blank">http://localhost:8080/hi1/xiao/ming</a> 因为找不到匹配方法最终报 404 错误。追根溯源就是 AntPathMatcher 匹配不了&rdquo;/hi1/xiao/ming&rdquo;和&rdquo;/hi1/{name}&ldquo;。</p>

<p>另外，我们再回头思考 <a href="http://localhost:8080/hi1/xiaoming/" target="_blank">http://localhost:8080/hi1/xiaoming/</a> 为什么没有报错而是直接去掉了/。这里我直接贴出了负责执行 AntPathMatcher 匹配的 PatternsRequestCondition#getMatchingPattern 方法的部分关键代码：</p>

<pre><code>private String getMatchingPattern(String pattern, String lookupPath) {
   //省略其他非关键代码
   if (this.pathMatcher.match(pattern, lookupPath)) {
      return pattern;
   }
   //尝试加一个/来匹配
   if (this.useTrailingSlashMatch) {
      if (!pattern.endsWith(&quot;/&quot;) &amp;&amp; this.pathMatcher.match(pattern + &quot;/&quot;, lookupPath)) {
         return pattern + &quot;/&quot;;
      }
   }
   return null;
}
</code></pre>

<p>在这段代码中，AntPathMatcher 匹配不了&rdquo;/hi1/xiaoming/&ldquo;和&rdquo;/hi1/{name}&ldquo;，所以不会直接返回。进而，在 useTrailingSlashMatch 这个参数启用时（默认启用），会把 Pattern 结尾加上/再尝试匹配一次。如果能匹配上，在最终返回 Pattern 时就隐式自动加/。</p>

<p>很明显，我们的案例符合这种情况，等于说我们最终是用了&rdquo;/hi1/{name}/&ldquo;这个 Pattern，而不再是&rdquo;/hi1/{name}&ldquo;。所以自然 URL 解析 name 结果是去掉/的。</p>

<h3 id="问题修正">问题修正</h3>

<p>针对这个案例，有了源码的剖析，我们可能会想到可以先用&rdquo;**&ldquo;匹配上路径，等进入方法后再尝试去解析，这样就可以万无一失吧。具体修改代码如下：</p>

<pre><code>@RequestMapping(path = &quot;/hi1/**&quot;, method = RequestMethod.GET)
public String hi1(HttpServletRequest request){
    String requestURI = request.getRequestURI();
    return requestURI.split(&quot;/hi1/&quot;)[1];
};
</code></pre>

<p>但是这种修改方法还是存在漏洞，假设我们路径的 name 中刚好又含有&rdquo;/hi1/&ldquo;，则 split 后返回的值就并不是我们想要的。实际上，更合适的修订代码示例如下：</p>

<pre><code>private AntPathMatcher antPathMatcher = new AntPathMatcher();

@RequestMapping(path = &quot;/hi1/**&quot;, method = RequestMethod.GET)
public String hi1(HttpServletRequest request){
    String path = (String) request.getAttribute(HandlerMapping.PATH_WITHIN_HANDLER_MAPPING_ATTRIBUTE);
    //matchPattern 即为&quot;/hi1/**&quot;
    String matchPattern = (String) request.getAttribute(HandlerMapping.BEST_MATCHING_PATTERN_ATTRIBUTE); 
    return antPathMatcher.extractPathWithinPattern(matchPattern, path); 
};
</code></pre>

<p>经过修改，两个错误都得以解决了。当然也存在一些其他的方案，例如对传递的参数进行 URL 编码以避免出现/，或者干脆直接把这个变量作为请求参数、Header 等，而不是作为 URL 的一部分。你完全可以根据具体情况来选择合适的方案。</p>

<h2 id="案例-2-错误使用-requestparam-pathvarible-等注解">案例 2：错误使用@RequestParam、@PathVarible 等注解</h2>

<p>我们常常使用@RequestParam 和@PathVarible 来获取请求参数（request parameters）以及 path 中的部分。但是在频繁使用这些参数时，不知道你有没有觉得它们的使用方式并不友好，例如我们去获取一个请求参数 name，我们会定义如下：</p>

<blockquote>
<p>@RequestParam(&ldquo;name&rdquo;) String name</p>
</blockquote>

<p>此时，我们会发现变量名称大概率会被定义成 RequestParam值。所以我们是不是可以用下面这种方式来定义：</p>

<blockquote>
<p>@RequestParam String name</p>
</blockquote>

<p>这种方式确实是可以的，本地测试也能通过。这里我还给出了完整的代码，你可以感受下这两者的区别。</p>

<pre><code>@RequestMapping(path = &quot;/hi1&quot;, method = RequestMethod.GET)
public String hi1(@RequestParam(&quot;name&quot;) String name){
    return name;
};

@RequestMapping(path = &quot;/hi2&quot;, method = RequestMethod.GET)
public String hi2(@RequestParam String name){
    return name;
};
</code></pre>

<p>很明显，对于喜欢追究极致简洁的同学来说，这个酷炫的功能是一个福音。但当我们换一个项目时，有可能上线后就失效了，然后报错 500，提示匹配不上。</p>

<p><img src="assets/76e5ba7f6d314bcba1bc7552bddaa35e.jpg" alt="" /></p>

<h3 id="案例解析-1">案例解析</h3>

<p>要理解这个问题出现的原因，首先我们需要把这个问题复现出来。例如我们可以修改下 pom.xml 来关掉两个选项：</p>

<pre><code>&lt;plugin&gt;
    &lt;groupId&gt;org.apache.maven.plugins&lt;/groupId&gt;
    &lt;artifactId&gt;maven-compiler-plugin&lt;/artifactId&gt;
   &lt;configuration&gt;
        &lt;debug&gt;false&lt;/debug&gt;
        &lt;parameters&gt;false&lt;/parameters&gt;
    &lt;/configuration&gt;
&lt;/plugin&gt;
</code></pre>

<p>上述配置显示关闭了 parameters 和 debug，这 2 个参数的作用你可以参考下面的表格：</p>

<p><img src="assets/873d203454bb488fbf5ea1f768dd0d69.jpg" alt="" /></p>

<p>通过上述描述，我们可以看出这 2 个参数控制了一些 debug 信息是否加进 class 文件中。我们可以开启这两个参数来编译，然后使用下面的命令来查看信息：</p>

<blockquote>
<p>javap -verbose HelloWorldController.class</p>
</blockquote>

<p>执行完命令后，我们会看到以下 class 信息：</p>

<p><img src="assets/455c0aa051114134bfaa848a2ed20978.jpg" alt="" /></p>

<p>debug 参数开启的部分信息就是 LocalVaribleTable，而 paramters 参数开启的信息就是 MethodParameters。观察它们的信息，你会发现它们都含有参数名name。</p>

<p>如果你关闭这两个参数，则 name 这个名称自然就没有了。而这个方法本身在 @RequestParam 中又没有指定名称，那么 Spring 此时还能找到解析的方法么？</p>

<p>答案是否定的，这里我们可以顺带说下 Spring 解析请求参数名称的过程，参考代码 AbstractNamedValueMethodArgumentResolver#updateNamedValueInfo：</p>

<pre><code>private NamedValueInfo updateNamedValueInfo(MethodParameter parameter, NamedValueInfo info) {
   String name = info.name;
   if (info.name.isEmpty()) {
      name = parameter.getParameterName();
      if (name == null) {
         throw new IllegalArgumentException(
               &quot;Name for argument type [&quot; + parameter.getNestedParameterType().getName() +
               &quot;] not available, and parameter name information not found in class file either.&quot;);
      }
   }
   String defaultValue = (ValueConstants.DEFAULT_NONE.equals(info.defaultValue) ? null : info.defaultValue);
   return new NamedValueInfo(name, info.required, defaultValue);
}
</code></pre>

<p>其中 NamedValueInfo 的 name 为 @RequestParam 指定的值。很明显，在本案例中，为 null。</p>

<p>所以这里我们就会尝试调用 parameter.getParameterName() 来获取参数名作为解析请求参数的名称。但是，很明显，关掉上面两个开关后，就不可能在 class 文件中找到参数名了，这点可以从下面的调试试图中得到验证：</p>

<p><img src="assets/d9288d7982c648e2816dc181cf985069.jpg" alt="" /></p>

<p>当参数名不存在，@RequestParam 也没有指明，自然就无法决定到底要用什么名称去获取请求参数，所以就会报本案例的错误。</p>

<h3 id="问题修正-1">问题修正</h3>

<p>模拟出了问题是如何发生的，我们自然可以通过开启这两个参数让其工作起来。但是思考这两个参数的作用，很明显，它可以让我们的程序体积更小，所以很多项目都会青睐去关闭这两个参数。</p>

<p>为了以不变应万变，正确的修正方式是<strong>必须显式在@RequestParam 中指定请求参数名</strong>。具体修改如下：</p>

<blockquote>
<p>@RequestParam(&ldquo;name&rdquo;) String name</p>
</blockquote>

<p>通过这个案例，我们可以看出：很多功能貌似可以永远工作，但是实际上，只是在特定的条件下而已。另外，这里再拓展下，IDE 都喜欢开启相关 debug 参数，所以 IDE 里运行的程序不见得对产线适应，例如针对 parameters 这个参数，IDEA 默认就开启了。</p>

<p>另外，本案例围绕的都是 @RequestParam，其实 @PathVarible 也有一样的问题。这里你要注意。</p>

<p>那么说到这里，我顺带提一个可能出现的小困惑：我们这里讨论的参数，和 @QueryParam、@PathParam 有什么区别？实际上，后者都是 JAX-RS 自身的注解，不需要额外导包。而 @RequestParam 和 @PathVariable 是 Spring 框架中的注解，需要额外导入依赖包。另外不同注解的参数也不完全一致。</p>

<h2 id="案例-3-未考虑参数是否可选">案例 3：未考虑参数是否可选</h2>

<p>在上面的案例中，我们提到了 @RequestParam 的使用。而对于它的使用，我们常常会遇到另外一个问题。当需要特别多的请求参数时，我们往往会忽略其中一些参数是否可选。例如存在类似这样的代码：</p>

<pre><code>@RequestMapping(path = &quot;/hi4&quot;, method = RequestMethod.GET)
public String hi4(@RequestParam(&quot;name&quot;) String name, @RequestParam(&quot;address&quot;) String address){
    return name + &quot;:&quot; + address;
};
</code></pre>

<p>在访问 <a href="http://localhost:8080/hi2?name=xiaoming&amp;address=beijing" target="_blank">http://localhost:8080/hi4?name=xiaoming&amp;address=beijing</a> 时并不会出问题，但是一旦用户仅仅使用 name 做请求（即 <a href="http://localhost:8080/hi4?name=xiaoming" target="_blank">http://localhost:8080/hi4?name=xiaoming</a> ）时，则会直接报错如下：</p>

<p><img src="assets/87b6a2663d5148bc8c316a93f38a009a.jpg" alt="" /></p>

<p>此时，返回错误码 400，提示请求格式错误：此处缺少 address 参数。</p>

<p>实际上，部分初学者即使面对这个错误，也会觉得惊讶，既然不存在 address，address 应该设置为 null，而不应该是直接报错不是么？接下来我们就分析下。</p>

<h3 id="案例解析-2">案例解析</h3>

<p>要了解这个错误出现的根本原因，你就需要了解请求参数的发生位置。</p>

<p>实际上，这里我们也能按注解名（@RequestParam）来确定解析发生的位置是在 RequestParamMethodArgumentResolver 中。为什么是它？</p>

<p>追根溯源，针对当前案例，当根据 URL 匹配上要执行的方法是 hi4 后，要反射调用它，必须解析出方法参数 name 和 address 才可以。而它们被 @RequestParam 注解修饰，所以解析器借助 RequestParamMethodArgumentResolver 就成了很自然的事情。</p>

<p>接下来我们看下 RequestParamMethodArgumentResolver 对参数解析的一些关键操作，参考其父类方法 AbstractNamedValueMethodArgumentResolver#resolveArgument：</p>

<pre><code>public final Object resolveArgument(MethodParameter parameter, @Nullable ModelAndViewContainer mavContainer,
      NativeWebRequest webRequest, @Nullable WebDataBinderFactory binderFactory) throws Exception {
   NamedValueInfo namedValueInfo = getNamedValueInfo(parameter);
   MethodParameter nestedParameter = parameter.nestedIfOptional();
   //省略其他非关键代码
   //获取请求参数
   Object arg = resolveName(resolvedName.toString(), nestedParameter, webRequest);
   if (arg == null) {
      if (namedValueInfo.defaultValue != null) {
         arg = resolveStringValue(namedValueInfo.defaultValue);
      }
      else if (namedValueInfo.required &amp;&amp; !nestedParameter.isOptional()) {
         handleMissingValue(namedValueInfo.name, nestedParameter, webRequest);
      }
      arg = handleNullValue(namedValueInfo.name, arg, nestedParameter.getNestedParameterType());
   }
   //省略后续代码：类型转化等工作
   return arg;
}
</code></pre>

<p>如代码所示，当缺少请求参数的时候，通常我们会按照以下几个步骤进行处理。</p>

<p><strong>1. 查看 namedValueInfo 的默认值，如果存在则使用它</strong></p>

<p>这个变量实际是通过下面的方法来获取的，参考 RequestParamMethodArgumentResolver#createNamedValueInfo：</p>

<pre><code>@Override
protected NamedValueInfo createNamedValueInfo(MethodParameter parameter) {
   RequestParam ann = parameter.getParameterAnnotation(RequestParam.class);
   return (ann != null ? new RequestParamNamedValueInfo(ann) : new RequestParamNamedValueInfo());
}
</code></pre>

<p>实际上就是 @RequestParam 的相关信息，我们调试下，就可以验证这个结论，具体如下图所示：</p>

<p><img src="assets/e5c02e8275a6492b8bcefed577e07845.jpg" alt="" /></p>

<p><strong>2. 在 @RequestParam 没有指明默认值时，会查看这个参数是否必须，如果必须，则按错误处理</strong></p>

<p>判断参数是否必须的代码即为下述关键代码行：</p>

<blockquote>
<p>namedValueInfo.required &amp;&amp; !nestedParameter.isOptional()</p>
</blockquote>

<p>很明显，若要判定一个参数是否是必须的，需要同时满足两个条件：条件 1 是@RequestParam 指明了必须（即属性 required 为 true，实际上它也是默认值），条件 2 是要求 @RequestParam 标记的参数本身不是可选的。</p>

<p>我们可以通过 MethodParameter#isOptional 方法看下可选的具体含义：</p>

<pre><code>public boolean isOptional() {
   return (getParameterType() == Optional.class || hasNullableAnnotation() ||
         (KotlinDetector.isKotlinReflectPresent() &amp;&amp;
               KotlinDetector.isKotlinType(getContainingClass()) &amp;&amp;
               KotlinDelegate.isOptional(this)));
}
</code></pre>

<p>在不使用 Kotlin 的情况下，所谓可选，就是参数的类型为 Optional，或者任何标记了注解名为 Nullable 且 RetentionPolicy 为 RUNTIM 的注解。</p>

<p><strong>3. 如果不是必须，则按 null 去做具体处理</strong></p>

<p>如果接受类型是 boolean，返回 false，如果是基本类型则直接报错，这里不做展开。</p>

<p>结合我们的案例，我们的参数符合步骤 2 中判定为必选的条件，所以最终会执行方法 AbstractNamedValueMethodArgumentResolver#handleMissingValue：</p>

<pre><code>protected void handleMissingValue(String name, MethodParameter parameter) throws ServletException {
   throw new ServletRequestBindingException(&quot;Missing argument '&quot; + name +
         &quot;' for method parameter of type &quot; + parameter.getNestedParameterType().getSimpleName());
}
</code></pre>

<h3 id="问题修正-2">问题修正</h3>

<p>通过案例解析，我们很容易就能修正这个问题，就是让参数有默认值或为非可选即可，具体方法包含以下几种。</p>

<p><strong>1. 设置 @RequestParam 的默认值</strong></p>

<p>修改代码如下：</p>

<blockquote>
<p>@RequestParam(value = &ldquo;address&rdquo;, defaultValue = &ldquo;no address&rdquo;) String address</p>
</blockquote>

<p><strong>2. 设置 @RequestParam 的 required 值</strong></p>

<p>修改代码如下：</p>

<blockquote>
<p>@RequestParam(value = &ldquo;address&rdquo;, required = false) String address)</p>
</blockquote>

<p><strong>3. 标记任何名为 Nullable 且 RetentionPolicy 为 RUNTIME 的注解</strong></p>

<p>修改代码如下：</p>

<blockquote>
<p><a href="//org.springframework.lang.Nullable" target="_blank">//org.springframework.lang.Nullable</a> 可以-
<a href="//edu.umd.cs.findbugs.annotations.Nullable" target="_blank">//edu.umd.cs.findbugs.annotations.Nullable</a> 可以-
@RequestParam(value = &ldquo;address&rdquo;) @Nullable String address</p>
</blockquote>

<p><strong>4. 修改参数类型为 Optional</strong></p>

<p>修改代码如下：</p>

<blockquote>
<p>@RequestParam(value = &ldquo;address&rdquo;) Optional address</p>
</blockquote>

<p>从这些修正方法不难看出：假设你不学习源码，解决方法就可能只局限于一两种，但是深入源码后，解决方法就变得格外多了。这里要特别强调的是：<strong>在Spring Web 中，默认情况下，请求参数是必选项。</strong></p>

<h2 id="案例-4-请求参数格式错误">案例 4：请求参数格式错误</h2>

<p>当我们使用 Spring URL 相关的注解，会发现 Spring 是能够完成自动转化的。例如在下面的代码中，age 可以被直接定义为 int 这种基本类型（Integer 也可以），而不是必须是 String 类型。</p>

<pre><code>@RequestMapping(path = &quot;/hi5&quot;, method = RequestMethod.GET)
public String hi5(@RequestParam(&quot;name&quot;) String name, @RequestParam(&quot;age&quot;) int age){
    return name + &quot; is &quot; + age + &quot; years old&quot;;
};
</code></pre>

<p>鉴于 Spring 的强大转化功能，我们断定 Spring 也支持日期类型的转化（也确实如此），于是我们可能会写出类似下面这样的代码：</p>

<pre><code>@RequestMapping(path = &quot;/hi6&quot;, method = RequestMethod.GET)
public String hi6(@RequestParam(&quot;Date&quot;) Date date){
    return &quot;date is &quot; + date ;
};
</code></pre>

<p>然后，我们使用一些看似明显符合日期格式的 URL 来访问，例如 <a href="http://localhost:8080/hi6?date=2021-5-1%2020:26:53" target="_blank">http://localhost:8080/hi6?date=2021-5-1 20:26:53</a>，我们会发现 Spring 并不能完成转化，而是报错如下：</p>

<p><img src="assets/0b3885184edf4b5a986f8613169ceabd.jpg" alt="" /></p>

<p>此时，返回错误码 400，错误信息为&rdquo;Failed to convert value of type &lsquo;java.lang.String&rsquo; to required type &lsquo;java.util.Date&rdquo;。</p>

<p>如何理解这个案例？如果实现自动转化，我们又需要做什么？</p>

<h3 id="案例解析-3">案例解析</h3>

<p>不管是使用 @PathVarible 还是 @RequetParam，我们一般解析出的结果都是一个 String 或 String 数组。例如，使用 @RequetParam 解析的关键代码参考 RequestParamMethodArgumentResolver#resolveName 方法：</p>

<pre><code>@Nullable
protected Object resolveName(String name, MethodParameter parameter, NativeWebRequest request) throws Exception {
   //省略其他非关键代码
   if (arg == null) {
      String[] paramValues = request.getParameterValues(name);
      if (paramValues != null) {
         arg = (paramValues.length == 1 ? paramValues[0] : paramValues);
      }
   }
   return arg;
}
</code></pre>

<p>这里我们调用的&rdquo;request.getParameterValues(name)&ldquo;，返回的是一个 String 数组，最终给上层调用者返回的是单个 String（如果只有一个元素时）或者 String 数组。</p>

<p>所以很明显，在这个测试程序中，我们给上层返回的是一个 String，这个 String 的值最终是需要做转化才能赋值给其他类型。例如对于案例中的&rdquo;int age&rdquo;定义，是需要转化为 int 基本类型的。这个基本流程可以通过 AbstractNamedValueMethodArgumentResolver#resolveArgument 的关键代码来验证：</p>

<pre><code>public final Object resolveArgument(MethodParameter parameter, @Nullable ModelAndViewContainer mavContainer,
      NativeWebRequest webRequest, @Nullable WebDataBinderFactory binderFactory) throws Exception {
   //省略其他非关键代码
   Object arg = resolveName(resolvedName.toString(), nestedParameter, webRequest);
   //以此为界，前面代码为解析请求参数,后续代码为转化解析出的参数
   if (binderFactory != null) {
      WebDataBinder binder = binderFactory.createBinder(webRequest, null, namedValueInfo.name);
      try {
         arg = binder.convertIfNecessary(arg, parameter.getParameterType(), parameter);
      }
      //省略其他非关键代码
   }
   //省略其他非关键代码
   return arg;
}
</code></pre>

<p>实际上在前面我们曾经提到过这个转化的基本逻辑，所以这里不再详述它具体是如何发生的。</p>

<p>在这里你只需要回忆出它是需要<strong>根据源类型和目标类型寻找转化器来执行转化的</strong>。在这里，对于 age 而言，最终找出的转化器是 StringToNumberConverterFactory。而对于 Date 型的 Date 变量，在本案例中，最终找到的是 ObjectToObjectConverter。它的转化过程参考下面的代码：</p>

<pre><code>public Object convert(@Nullable Object source, TypeDescriptor sourceType, TypeDescriptor targetType) {
   if (source == null) {
      return null;
   }
   Class&lt;?&gt; sourceClass = sourceType.getType();
   Class&lt;?&gt; targetClass = targetType.getType();
   //根据源类型去获取构建出目标类型的方法：可以是工厂方法（例如 valueOf、from 方法）也可以是构造器
   Member member = getValidatedMember(targetClass, sourceClass);
   try {
      if (member instanceof Method) {
         //如果是工厂方法，通过反射创建目标实例
      }
      else if (member instanceof Constructor) {
         //如果是构造器，通过反射创建实例
         Constructor&lt;?&gt; ctor = (Constructor&lt;?&gt;) member;
         ReflectionUtils.makeAccessible(ctor);
         return ctor.newInstance(source);
      }
   }
   catch (InvocationTargetException ex) {
      throw new ConversionFailedException(sourceType, targetType, source, ex.getTargetException());
   }
   catch (Throwable ex) {
      throw new ConversionFailedException(sourceType, targetType, source, ex);
   }
</code></pre>

<p>当使用 ObjectToObjectConverter 进行转化时，是根据反射机制带着源目标类型来查找可能的构造目标实例方法，例如构造器或者工厂方法，然后再次通过反射机制来创建一个目标对象。所以对于 Date 而言，最终调用的是下面的 Date 构造器：</p>

<pre><code>public Date(String s) {
    this(parse(s));
}
</code></pre>

<p>然而，我们传入的 <a href="http://localhost:8080/hi6?date=2021-5-1%2020:26:53" target="_blank">2021-5-1 20:26:53</a> 虽然确实是一种日期格式，但用来作为 Date 构造器参数是不支持的，最终报错，并被上层捕获，转化为 ConversionFailedException 异常。这就是这个案例背后的故事了。</p>

<h3 id="问题修正-3">问题修正</h3>

<p>那么怎么解决呢？提供两种方法。</p>

<p><strong>1. 使用 Date 支持的格式</strong></p>

<p>例如下面的测试 URL 就可以工作起来：</p>

<blockquote>
<p><a href="http://localhost:8080/hi6?date=Sat" target="_blank">http://localhost:8080/hi6?date=Sat</a>, 12 Aug 1995 13:30:00 GMT</p>
</blockquote>

<p><strong>2. 使用好内置格式转化器</strong></p>

<p>实际上，在Spring中，要完成 String 对于 Date 的转化，ObjectToObjectConverter 并不是最好的转化器。我们可以使用更强大的AnnotationParserConverter。<strong>在Spring 初始化时，会构建一些针对日期型的转化器，即相应的一些 AnnotationParserConverter 的实例。</strong>但是为什么有时候用不上呢？</p>

<p>这是因为 AnnotationParserConverter 有目标类型的要求，这点我们可以通过调试角度来看下，参考 FormattingConversionService#addFormatterForFieldAnnotation 方法的调试试图：</p>

<p><img src="assets/63dcc9f1056c4dc38d3a06d9511ba6bf.jpg" alt="" /></p>

<p>这是适应于 String 到 Date 类型的转化器 AnnotationParserConverter 实例的构造过程，其需要的 annototationType 参数为 DateTimeFormat。</p>

<p>annototationType 的作用正是为了帮助判断是否能用这个转化器，这一点可以参考代码 AnnotationParserConverter#matches：</p>

<pre><code>@Override
public boolean matches(TypeDescriptor sourceType, TypeDescriptor targetType) {
   return targetType.hasAnnotation(this.annotationType);
}
</code></pre>

<p>最终构建出来的转化器相关信息可以参考下图：</p>

<p><img src="assets/df679f0f62984769b2457455626cf616.jpg" alt="" /></p>

<p>图中构造出的转化器是可以用来转化 String 到 Date，但是它要求我们标记 @DateTimeFormat。很明显，我们的参数 Date 并没有标记这个注解，所以这里为了使用这个转化器，我们可以使用上它并提供合适的格式。这样就可以让原来不工作的 URL 工作起来，具体修改代码如下：</p>

<pre><code>@DateTimeFormat(pattern=&quot;yyyy-MM-dd HH:mm:ss&quot;) Date date
</code></pre>

<p>以上即为本案例的解决方案。除此之外，我们完全可以制定一个转化器来帮助我们完成转化，这里不再赘述。另外，通过这个案例，我们可以看出：尽管 Spring 给我们提供了很多内置的转化功能，但是我们一定要注意，格式是否符合对应的要求，否则代码就可能会失效。</p>

<h2 id="重点回顾">重点回顾</h2>

<p>通过这一讲的学习，我们了解到了在Spring解析URL中的一些常见错误及其背后的深层原因。这里再次回顾下重点：</p>

<ol>
<li>当我们使用@PathVariable时，一定要注意传递的值是不是含有 / ;</li>
<li>当我们使用@RequestParam、@PathVarible等注解时，一定要意识到一个问题，虽然下面这两种方式（以@RequestParam使用示例）都可以，但是后者在一些项目中并不能正常工作，因为很多产线的编译配置会去掉不是必须的调试信息。</li>
</ol>

<pre><code>@RequestMapping(path = &quot;/hi1&quot;, method = RequestMethod.GET)
public String hi1(@RequestParam(&quot;name&quot;) String name){
    return name;
};
//方式2：没有显式指定RequestParam的“name”，这种方式有时候会不行
@RequestMapping(path = &quot;/hi2&quot;, method = RequestMethod.GET)
public String hi2(@RequestParam String name){
    return name;
};
</code></pre>

<ol>
<li>任何一个参数，我们都需要考虑它是可选的还是必须的。同时，你一定要想到参数类型的定义到底能不能从请求中自动转化而来。Spring本身给我们内置了很多转化器，但是我们要以合适的方式使用上它。另外，Spring对很多类型的转化设计都很贴心，例如使用下面的注解就能解决自定义日期格式参数转化问题。</li>
</ol>

<pre><code>@DateTimeFormat(pattern=&quot;yyyy-MM-dd HH:mm:ss&quot;) Date date
</code></pre>

<p>希望这些核心知识点，能帮助你高效解析URL。</p>

<h2 id="思考题">思考题</h2>

<p>关于 URL 解析，其实还有许多让我们惊讶的地方，例如案例 2 的部分代码：</p>

<pre><code>@RequestMapping(path = &quot;/hi2&quot;, method = RequestMethod.GET)
public String hi2(@RequestParam(&quot;name&quot;) String name){
    return name;
};
</code></pre>

<p>在上述代码的应用中，我们可以使用 <a href="http://localhost:8080/hi2?name=xiaoming&amp;name=hanmeimei" target="_blank">http://localhost:8080/hi2?name=xiaoming&amp;name=hanmeimei</a> 来测试下，结果会返回什么呢？你猜会是<a href="http://localhost:8080/hi2?name=xiaoming&amp;name=hanmeimei" target="_blank">xiaoming&amp;name=hanmeimei</a> 么？</p>

<p>我们留言区见！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#e5898989dcd1d4d4d5d2a58288848c89cb868a88" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f128997fe73ede4',t:'MTczNDA1NzA4MS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>