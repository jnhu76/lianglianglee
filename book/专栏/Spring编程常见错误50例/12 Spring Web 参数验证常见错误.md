<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=12&#32;Spring&#32;Web&#32;参数验证常见错误>
        <link rel="icon" href="/static/favicon.png">
        <title>12 Spring Web 参数验证常见错误 </title>
        
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
                            <h1 id="title" data-id="12 Spring Web 参数验证常见错误" class="title">12 Spring Web 参数验证常见错误</h1>
                            <div><p>你好，我是傅健，这节课我们来聊聊 Spring Web 开发中的参数检验（Validation）。</p>

<p>参数检验是我们在Web编程时经常使用的技术之一，它帮助我们完成请求的合法性校验，可以有效拦截无效请求，从而达到节省系统资源、保护系统的目的。</p>

<p>相比较其他 Spring 技术，Spring提供的参数检验功能具有独立性强、使用难度不高的特点。但是在实践中，我们仍然会犯一些常见的错误，这些错误虽然不会导致致命的后果，但是会影响我们的使用体验，例如非法操作要在业务处理时才被拒绝且返回的响应码不够清晰友好。而且这些错误不经测试很难发现，接下来我们就具体分析下这些常见错误案例及背后的原理。</p>

<h2 id="案例1-对象参数校验失效">案例1：对象参数校验失效</h2>

<p>在构建Web服务时，我们一般都会对一个HTTP请求的 Body 内容进行校验，例如我们来看这样一个案例及对应代码。</p>

<p>当开发一个学籍管理系统时，我们会提供了一个 API 接口去添加学生的相关信息，其对象定义参考下面的代码：</p>

<pre><code>import lombok.Data;
import javax.validation.constraints.Size;
@Data
public class Student {
    @Size(max = 10)
    private String name;
    private short age;
}
</code></pre>

<p>这里我们使用了@Size(max = 10)给学生的姓名做了约束（最大为 10 字节），以拦截姓名过长、不符合“常情”的学生信息的添加。</p>

<p>定义完对象后，我们再定义一个 Controller 去使用它，使用方法如下：</p>

<pre><code>import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
@RestController
@Slf4j
@Validated
public class StudentController {
    @RequestMapping(path = &quot;students&quot;, method = RequestMethod.POST)
    public void addStudent(@RequestBody Student student){
        log.info(&quot;add new student: {}&quot;, student.toString());
        //省略业务代码
    };
}
</code></pre>

<p>我们提供了一个支持学生信息添加的接口。启动服务后，使用 IDEA 自带的 HTTP Client 工具来发送下面的请求以添加一个学生，当然，这个学生的姓名会远超想象（即this_is_my_name_which_is_too_long）：</p>

<pre><code>POST http://localhost:8080/students
Content-Type: application/json
{
 &quot;name&quot;: &quot;this_is_my_name_which_is_too_long&quot;,
 &quot;age&quot;: 10
}
</code></pre>

<p>很明显，发送这样的请求（name 超长）是期待 Spring Validation 能拦截它的，我们的预期响应如下（省略部分响应字段）：</p>

<pre><code>HTTP/1.1 400 
Content-Type: application/json

{
  &quot;timestamp&quot;: &quot;2021-01-03T00:47:23.994+0000&quot;,
  &quot;status&quot;: 400,
  &quot;error&quot;: &quot;Bad Request&quot;,
  &quot;errors&quot;: [
      &quot;defaultMessage&quot;: &quot;个数必须在 0 和 10 之间&quot;,
      &quot;objectName&quot;: &quot;student&quot;,
      &quot;field&quot;: &quot;name&quot;,
      &quot;rejectedValue&quot;: &quot;this_is_my_name_which_is_too_long&quot;,
      &quot;bindingFailure&quot;: false,
      &quot;code&quot;: &quot;Size&quot;
    }
  ],
  &quot;message&quot;: &quot;Validation failed for object='student'. Error count: 1&quot;,
  &quot;path&quot;: &quot;/students&quot;
}
</code></pre>

<p>但是理想与现实往往有差距。实际测试会发现，使用上述代码构建的Web服务并没有做任何拦截。</p>

<h3 id="案例解析">案例解析</h3>

<p>要找到这个问题的根源，我们就需要对 Spring Validation 有一定的了解。首先，我们来看下 RequestBody 接受对象校验发生的位置和条件。</p>

<p>假设我们构建Web服务使用的是Spring Boot技术，我们可以参考下面的时序图了解它的核心执行步骤：</p>

<p><img src="assets/936435d2eaba4a30a37276cce2ca80d8.jpg" alt="" /></p>

<p>如上图所示，当一个请求来临时，都会进入 DispatcherServlet，执行其 doDispatch()，此方法会根据 Path、Method 等关键信息定位到负责处理的 Controller 层方法（即 addStudent 方法），然后通过反射去执行这个方法，具体反射执行过程参考下面的代码（InvocableHandlerMethod#invokeForRequest）：</p>

<pre><code>public Object invokeForRequest(NativeWebRequest request, @Nullable ModelAndViewContainer mavContainer,
      Object... providedArgs) throws Exception {
   //根据请求内容和方法定义获取方法参数实例
   Object[] args = getMethodArgumentValues(request, mavContainer, providedArgs);
   if (logger.isTraceEnabled()) {
      logger.trace(&quot;Arguments: &quot; + Arrays.toString(args));
   }
   //携带方法参数实例去“反射”调用方法
   return doInvoke(args);
}
</code></pre>

<p>要使用 Java 反射去执行一个方法，需要先获取调用的参数，上述代码正好验证了这一点：getMethodArgumentValues() 负责获取方法执行参数，doInvoke() 负责使用这些获取到的参数去执行。</p>

<p>而具体到getMethodArgumentValues() 如何获取方法调用参数，可以参考 addStudent 的方法定义，我们需要从当前的请求（NativeWebRequest ）中构建出 Student 这个方法参数的实例。</p>

<blockquote>
<p>public void addStudent(@RequestBody Student student)</p>
</blockquote>

<p>那么如何构建出这个方法参数实例？Spring 内置了相当多的 HandlerMethodArgumentResolver，参考下图：</p>

<p><img src="assets/1e1b6c3f990d4554ad99711af7ca1ed4.jpg" alt="" /></p>

<p>当试图构建出一个方法参数时，会遍历所有支持的解析器（Resolver）以找出适合的解析器，查找代码参考HandlerMethodArgumentResolverComposite#getArgumentResolver：</p>

<pre><code>@Nullable
private HandlerMethodArgumentResolver getArgumentResolver(MethodParameter parameter) {
   HandlerMethodArgumentResolver result = this.argumentResolverCache.get(parameter);
   if (result == null) {
      //轮询所有的HandlerMethodArgumentResolver
      for (HandlerMethodArgumentResolver resolver : this.argumentResolvers) {
         //判断是否匹配当前HandlerMethodArgumentResolver 
         if (resolver.supportsParameter(parameter)) {
            result = resolver;            
            this.argumentResolverCache.put(parameter, result);
            break;
         }
      }
   }
   return result;
}
</code></pre>

<p>对于 student 参数而言，它被标记为@RequestBody，当遍历到 RequestResponseBodyMethodProcessor 时就会匹配上。匹配代码参考其 RequestResponseBodyMethodProcessor 的supportsParameter 方法：</p>

<pre><code>@Override
public boolean supportsParameter(MethodParameter parameter) {
   return parameter.hasParameterAnnotation(RequestBody.class);
}
</code></pre>

<p>找到 Resolver 后，就会执行 HandlerMethodArgumentResolver#resolveArgument 方法。它首先会根据当前的请求（NativeWebRequest）组装出 Student 对象并对这个对象进行必要的校验，校验的执行参考AbstractMessageConverterMethodArgumentResolver#validateIfApplicable：</p>

<pre><code>protected void validateIfApplicable(WebDataBinder binder, MethodParameter parameter) {
   Annotation[] annotations = parameter.getParameterAnnotations();
   for (Annotation ann : annotations) {
      Validated validatedAnn = AnnotationUtils.getAnnotation(ann, Validated.class);
      //判断是否需要校验
      if (validatedAnn != null || ann.annotationType().getSimpleName().startsWith(&quot;Valid&quot;)) {
         Object hints = (validatedAnn != null ? validatedAnn.value() : AnnotationUtils.getValue(ann));
         Object[] validationHints = (hints instanceof Object[] ? (Object[]) hints : new Object[] {hints});
         //执行校验
         binder.validate(validationHints);
         break;
      }
   }
}
</code></pre>

<p>如上述代码所示，要对 student 实例进行校验（执行binder.validate(validationHints)方法），必须匹配下面两个条件的其中之一：</p>

<ol>
<li>标记了 org.springframework.validation.annotation.Validated 注解；</li>
<li>标记了其他类型的注解，且注解名称以Valid关键字开头。</li>
</ol>

<p>因此，结合案例程序，我们知道：student 方法参数并不符合这两个条件，所以即使它的内部成员添加了校验（即@Size(max = 10)），也不能生效。</p>

<h3 id="问题修正">问题修正</h3>

<p>针对这个案例，有了源码的剖析，我们就可以很快地找到解决方案。即对于 RequestBody 接受的对象参数而言，要启动 Validation，必须将对象参数标记上 @Validated 或者其他以@Valid关键字开头的注解，因此，我们可以采用对应的策略去修正问题。</p>

<ol>
<li>标记 @Validated</li>
</ol>

<p>修正后关键代码行如下：</p>

<blockquote>
<p>public void addStudent(<strong>@Validated</strong> @RequestBody Student student)</p>
</blockquote>

<ol>
<li>标记@Valid关键字开头的注解</li>
</ol>

<p>这里我们可以直接使用熟识的 javax.validation.Valid 注解，它就是一种以@Valid关键字开头的注解，修正后关键代码行如下：</p>

<blockquote>
<p>public void addStudent(<strong>@Valid</strong> @RequestBody Student student)</p>
</blockquote>

<p>另外，我们也可以自定义一个以Valid关键字开头的注解，定义如下：</p>

<pre><code>import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
@Retention(RetentionPolicy.RUNTIME)
public @interface ValidCustomized {
}
</code></pre>

<p>定义完成后，将它标记给 student 参数对象，关键代码行如下：</p>

<blockquote>
<p>public void addStudent(<strong>@</strong>ValidCustomized @RequestBody Student student)</p>
</blockquote>

<p>通过上述2种策略、3种具体修正方法，我们最终让参数校验生效且符合预期，不过需要提醒你的是：当使用第3种修正方法时，一定要注意自定义的注解要显式标记@Retention(RetentionPolicy.RUNTIME)，否则校验仍不生效。这也是另外一个容易疏忽的地方，究其原因，不显式标记RetentionPolicy 时，默认使用的是 RetentionPolicy.CLASS，而这种类型的注解信息虽然会被保留在字节码文件（.class）中，但在加载进 JVM 时就会丢失了。所以在运行时，依据这个注解来判断是否校验，肯定会失效。</p>

<h2 id="案例2-嵌套校验失效">案例2：嵌套校验失效</h2>

<p>前面这个案例虽然比较经典，但是，它只是初学者容易犯的错误。实际上，关于 Validation 最容易忽略的是对嵌套对象的校验，我们沿用上面的案例举这样一个例子。</p>

<p>学生可能还需要一个联系电话信息，所以我们可以定义一个 Phone 对象，然后关联上学生对象，代码如下：</p>

<pre><code>public class Student {
    @Size(max = 10)
    private String name;
    private short age;   
    private Phone phone;
}
@Data
class Phone {
    @Size(max = 10)
    private String number;
}
</code></pre>

<p>这里我们也给 Phone 对象做了合法性要求（@Size(max = 10)），当我们使用下面的请求（请求 body 携带一个联系电话信息超过 10 位），测试校验会发现这个约束并不生效。</p>

<pre><code>POST http://localhost:8080/students
Content-Type: application/json
{
  &quot;name&quot;: &quot;xiaoming&quot;,
  &quot;age&quot;: 10,
  &quot;phone&quot;: {&quot;number&quot;:&quot;12306123061230612306&quot;}
}
</code></pre>

<p>为什么会不生效？</p>

<h3 id="案例解析-1">案例解析</h3>

<p>在解析案例 1 时，我们提及只要给对象参数 student 加上@Valid（或@Validated 等注解）就可以开启这个对象的校验。但实际上，关于 student 本身的 Phone 类型成员是否校验是在校验过程中（即案例1中的代码行binder.validate(validationHints)）决定的。</p>

<p>在校验执行时，首先会根据 Student 的类型定义找出所有的校验点，然后对 Student 对象实例执行校验，这个逻辑过程可以参考代码 ValidatorImpl#validate：</p>

<pre><code>@Override
public final &lt;T&gt; Set&lt;ConstraintViolation&lt;T&gt;&gt; validate(T object, Class&lt;?&gt;... groups) {
   //省略部分非关键代码
   Class&lt;T&gt; rootBeanClass = (Class&lt;T&gt;) object.getClass();
   //获取校验对象类型的“信息”（包含“约束”）
   BeanMetaData&lt;T&gt; rootBeanMetaData = beanMetaDataManager.getBeanMetaData( rootBeanClass );

   if ( !rootBeanMetaData.hasConstraints() ) {
      return Collections.emptySet();
   }

   //省略部分非关键代码
   //执行校验
   return validateInContext( validationContext, valueContext, validationOrder );
}
</code></pre>

<p>这里语句&rdquo;beanMetaDataManager.getBeanMetaData( rootBeanClass )&ldquo;根据 Student 类型组装出 BeanMetaData，BeanMetaData 即包含了需要做的校验（即 Constraint）。</p>

<p>在组装 BeanMetaData 过程中，会根据成员字段是否标记了@Valid 来决定（记录）这个字段以后是否做级联校验，参考代码 AnnotationMetaDataProvider#getCascadingMetaData：</p>

<pre><code>private CascadingMetaDataBuilder getCascadingMetaData(Type type, AnnotatedElement annotatedElement,
      Map&lt;TypeVariable&lt;?&gt;, CascadingMetaDataBuilder&gt; containerElementTypesCascadingMetaData) {
   return CascadingMetaDataBuilder.annotatedObject( type, annotatedElement.isAnnotationPresent( Valid.class ), containerElementTypesCascadingMetaData,
               getGroupConversions( annotatedElement ) );
}
</code></pre>

<p>在上述代码中&rdquo;annotatedElement.isAnnotationPresent( Valid.class )&ldquo;决定了 CascadingMetaDataBuilder#cascading 是否为 true。如果是，则在后续做具体校验时，做级联校验，而级联校验的过程与宿主对象（即Student）的校验过程大体相同，即先根据对象类型获取定义再来做校验。</p>

<p>在当前案例代码中，phone字段并没有被@Valid标记，所以关于这个字段信息的 cascading 属性肯定是false，因此在校验Student时并不会级联校验它。</p>

<h3 id="问题修正-1">问题修正</h3>

<p>从源码级别了解了嵌套 Validation 失败的原因后，我们会发现，要让嵌套校验生效，解决的方法只有一种，就是加上@Valid，修正代码如下：</p>

<blockquote>
<p>@Valid-
private Phone phone;</p>
</blockquote>

<p>当修正完问题后，我们会发现校验生效了。而如果此时去调试修正后的案例代码，会看到 phone 字段 MetaData 信息中的 cascading 确实为 true 了，参考下图：</p>

<p><img src="assets/61adff03683345cca71f42e95e6318f6.jpg" alt="" /></p>

<p>另外，假设我们不去解读源码，我们很可能会按照案例 1 所述的其他修正方法去修正这个问题。例如，使用 @Validated 来修正这个问题，但是此时你会发现，不考虑源码是否支持，代码本身也编译不过，这主要在于 @Validated 的定义是不允许修饰一个 Field 的：</p>

<pre><code>@Target({ElementType.TYPE, ElementType.METHOD, ElementType.PARAMETER})
@Retention(RetentionPolicy.RUNTIME)
@Documented
public @interface Validated {
</code></pre>

<p>通过上述方法修正问题，最终我们让嵌套验证生效了。但是你可能还是会觉得这个错误看起来不容易犯，那么可以试想一下，我们的案例仅仅是嵌套一层，而产品代码往往都是嵌套 n 层，此时我们是否能保证每一级都不会疏忽漏加@Valid呢？所以这仍然是一个典型的错误，需要你格外注意。</p>

<h2 id="案例3-误解校验执行">案例3：误解校验执行</h2>

<p>通过前面两个案例的填坑，我们一般都能让参数校验生效起来，但是校验本身有时候是一个无止境的完善过程，校验本身已经生效，但是否完美匹配我们所有苛刻的要求是另外一个容易疏忽的地方。例如，我们可能在实践中误解一些校验的使用。这里我们可以继续沿用前面的案例，变形一下。</p>

<p>之前我们定义的学生对象的姓名要求是小于 10 字节的（即@Size(max = 10)）。此时我们可能想完善校验，例如，我们希望姓名不能是空，此时你可能很容易想到去修改关键行代码如下：</p>

<pre><code>@Size(min = 1, max = 10)
private String name;
</code></pre>

<p>然后，我们以下面的 JSON Body 做测试：</p>

<pre><code>{
  &quot;name&quot;: &quot;&quot;,
  &quot;age&quot;: 10,
  &quot;phone&quot;: {&quot;number&quot;:&quot;12306&quot;}
}
</code></pre>

<p>测试结果符合我们的预期，但是假设更进一步，用下面的 JSON Body（去除 name 字段）做测试呢？</p>

<pre><code>{
  &quot;age&quot;: 10,
  &quot;phone&quot;: {&quot;number&quot;:&quot;12306&quot;}
}

</code></pre>

<p>我们会发现校验失败了。这结果难免让我们有一些惊讶，也倍感困惑：@Size(min = 1, max = 10) 都已经要求最小字节为 1 了，难道还只能约束空字符串（即“”），不能约束 null?</p>

<h3 id="案例解析-2">案例解析</h3>

<p>如果我们稍微留心点的话，就会发现其实 @Size 的 Javadoc 已经明确了这种情况，参考下图：</p>

<p><img src="assets/7c04e256f690499b9dd2a30150988dac.jpg" alt="" /></p>

<p>如图所示，&rdquo;null elements are considered valid&rdquo; 很好地解释了约束不住null的原因。当然纸上得来终觉浅，我们还需要从源码级别解读下@Size 的校验过程。</p>

<p>这里我们找到了完成@Size 约束的执行方法，参考 SizeValidatorForCharSequence#isValid 方法：</p>

<pre><code>   public boolean isValid(CharSequence charSequence, ConstraintValidatorContext constraintValidatorContext) {
      if ( charSequence == null ) {
         return true;
      }
      int length = charSequence.length();
      return length &gt;= min &amp;&amp; length &lt;= max;
   }
</code></pre>

<p>如代码所示，当字符串为 null 时，直接通过了校验，而不会做任何进一步的约束检查。</p>

<h3 id="问题修正-2">问题修正</h3>

<p>关于这个问题的修正，其实很简单，我们可以使用其他的注解（@NotNull 或@NotEmpty）来加强约束，修正代码如下：</p>

<pre><code>@NotEmpty
@Size(min = 1, max = 10)
private String name;
</code></pre>

<p>完成代码修改后，重新测试，你就会发现约束已经完全满足我们的需求了。</p>

<h2 id="重点回顾">重点回顾</h2>

<p>看完上面的一些案例，我们会发现，这些错误的直接结果都是校验完全失败或者部分失败，并不会造成严重的后果，但是就像本讲开头所讲的那样，这些错误会影响我们的使用体验，所以我们还是需要去规避这些错误，把校验做强最好！</p>

<p>另外，关于@Valid 和@Validation 是我们经常犯迷糊的地方，不知道到底有什么区别。同时我们也经常产生一些困惑，例如能用其中一种时，能不能用另外一种呢？</p>

<p>通过解析，我们会发现，在很多场景下，我们不一定要寄希望于搜索引擎去区别，只需要稍微研读下代码，反而更容易理解。例如，对于案例 1，研读完代码后，我们发现它们不仅可以互换，而且完全可以自定义一个以@Valid开头的注解来使用；而对于案例 2，只能用@Valid 去开启级联校验。</p>

<h2 id="思考题">思考题</h2>

<p>在上面的学籍管理系统中，我们还存在一个接口，负责根据学生的学号删除他的信息，代码如下：</p>

<pre><code>@RequestMapping(path = &quot;students/{id}&quot;, method = RequestMethod.DELETE)
public void deleteStudent(@PathVariable(&quot;id&quot;) @Range(min = 1,max = 10000) String id){
    log.info(&quot;delete student: {}&quot;,id);
    //省略业务代码
};
</code></pre>

<p>这个学生的编号是从请求的Path中获取的，而且它做了范围约束，必须在1到10000之间。那么你能找出负责解出 ID 的解析器（HandlerMethodArgumentResolver）是哪一种吗？校验又是如何触发的？</p>

<p>期待你的思考，我们留言区见！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#e4888888ddd0d5d5d4d3a48389858d88ca878b89" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f128ad53ab3ede4',t:'MTczNDA1NzEzMi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>