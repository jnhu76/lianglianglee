<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=06&#32;WebFlux&#32;中&#32;Thymeleaf&#32;和&#32;MongoDB&#32;实践>
        <link rel="icon" href="/static/favicon.png">
        <title>06 WebFlux 中 Thymeleaf 和 MongoDB 实践 </title>
        
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
                        <a class="menu-item" id="01 导读：课程概要.md" href="/%e4%b8%93%e6%a0%8f/%e6%a1%88%e4%be%8b%e4%b8%8a%e6%89%8b%20Spring%20Boot%20WebFlux%ef%bc%88%e5%ae%8c%ef%bc%89/01%20%e5%af%bc%e8%af%bb%ef%bc%9a%e8%af%be%e7%a8%8b%e6%a6%82%e8%a6%81.md">01 导读：课程概要.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 WebFlux 快速入门实践.md" href="/%e4%b8%93%e6%a0%8f/%e6%a1%88%e4%be%8b%e4%b8%8a%e6%89%8b%20Spring%20Boot%20WebFlux%ef%bc%88%e5%ae%8c%ef%bc%89/02%20WebFlux%20%e5%bf%ab%e9%80%9f%e5%85%a5%e9%97%a8%e5%ae%9e%e8%b7%b5.md">02 WebFlux 快速入门实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 WebFlux Web CRUD 实践.md" href="/%e4%b8%93%e6%a0%8f/%e6%a1%88%e4%be%8b%e4%b8%8a%e6%89%8b%20Spring%20Boot%20WebFlux%ef%bc%88%e5%ae%8c%ef%bc%89/03%20WebFlux%20Web%20CRUD%20%e5%ae%9e%e8%b7%b5.md">03 WebFlux Web CRUD 实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 WebFlux 整合 MongoDB.md" href="/%e4%b8%93%e6%a0%8f/%e6%a1%88%e4%be%8b%e4%b8%8a%e6%89%8b%20Spring%20Boot%20WebFlux%ef%bc%88%e5%ae%8c%ef%bc%89/04%20WebFlux%20%e6%95%b4%e5%90%88%20MongoDB.md">04 WebFlux 整合 MongoDB.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 WebFlux 整合 Thymeleaf.md" href="/%e4%b8%93%e6%a0%8f/%e6%a1%88%e4%be%8b%e4%b8%8a%e6%89%8b%20Spring%20Boot%20WebFlux%ef%bc%88%e5%ae%8c%ef%bc%89/05%20WebFlux%20%e6%95%b4%e5%90%88%20Thymeleaf.md">05 WebFlux 整合 Thymeleaf.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 WebFlux 中 Thymeleaf 和 MongoDB 实践.md" href="/%e4%b8%93%e6%a0%8f/%e6%a1%88%e4%be%8b%e4%b8%8a%e6%89%8b%20Spring%20Boot%20WebFlux%ef%bc%88%e5%ae%8c%ef%bc%89/06%20WebFlux%20%e4%b8%ad%20Thymeleaf%20%e5%92%8c%20MongoDB%20%e5%ae%9e%e8%b7%b5.md">06 WebFlux 中 Thymeleaf 和 MongoDB 实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 WebFlux 整合 Redis.md" href="/%e4%b8%93%e6%a0%8f/%e6%a1%88%e4%be%8b%e4%b8%8a%e6%89%8b%20Spring%20Boot%20WebFlux%ef%bc%88%e5%ae%8c%ef%bc%89/07%20WebFlux%20%e6%95%b4%e5%90%88%20Redis.md">07 WebFlux 整合 Redis.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 WebFlux 中 Redis 实现缓存.md" href="/%e4%b8%93%e6%a0%8f/%e6%a1%88%e4%be%8b%e4%b8%8a%e6%89%8b%20Spring%20Boot%20WebFlux%ef%bc%88%e5%ae%8c%ef%bc%89/08%20WebFlux%20%e4%b8%ad%20Redis%20%e5%ae%9e%e7%8e%b0%e7%bc%93%e5%ad%98.md">08 WebFlux 中 Redis 实现缓存.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 WebFlux 中 WebSocket 实现通信.md" href="/%e4%b8%93%e6%a0%8f/%e6%a1%88%e4%be%8b%e4%b8%8a%e6%89%8b%20Spring%20Boot%20WebFlux%ef%bc%88%e5%ae%8c%ef%bc%89/09%20WebFlux%20%e4%b8%ad%20WebSocket%20%e5%ae%9e%e7%8e%b0%e9%80%9a%e4%bf%a1.md">09 WebFlux 中 WebSocket 实现通信.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 WebFlux 集成测试及部署.md" href="/%e4%b8%93%e6%a0%8f/%e6%a1%88%e4%be%8b%e4%b8%8a%e6%89%8b%20Spring%20Boot%20WebFlux%ef%bc%88%e5%ae%8c%ef%bc%89/10%20WebFlux%20%e9%9b%86%e6%88%90%e6%b5%8b%e8%af%95%e5%8f%8a%e9%83%a8%e7%bd%b2.md">10 WebFlux 集成测试及部署.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 WebFlux 实战图书管理系统.md" href="/%e4%b8%93%e6%a0%8f/%e6%a1%88%e4%be%8b%e4%b8%8a%e6%89%8b%20Spring%20Boot%20WebFlux%ef%bc%88%e5%ae%8c%ef%bc%89/11%20WebFlux%20%e5%ae%9e%e6%88%98%e5%9b%be%e4%b9%a6%e7%ae%a1%e7%90%86%e7%b3%bb%e7%bb%9f.md">11 WebFlux 实战图书管理系统.md</a>
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
                            <h1 id="title" data-id="06 WebFlux 中 Thymeleaf 和 MongoDB 实践" class="title">06 WebFlux 中 Thymeleaf 和 MongoDB 实践</h1>
                            <div><h3 id="前言">前言</h3>

<p>本节内容主要还是总结上面两篇内容的操作，并实现一个复杂查询的小案例，那么没安装 MongoDB 的可以进行下面的安装流程。</p>

<p>Docker 安装 MognoDB 并启动如下。</p>

<p>（1）创建挂载目录：</p>

<pre><code>docker volume create mongo_data_db
docker volume create mongo_data_configdb

</code></pre>

<p>（2）启动 MognoDB：</p>

<pre><code>docker run -d \
    --name mongo \
    -v mongo_data_configdb:/data/configdb \
    -v mongo_data_db:/data/db \
    -p 27017:27017 \
    mongo \
    --auth

</code></pre>

<p>（3）初始化管理员账号：</p>

<pre><code>docker exec -it mongo     mongo              admin
                        // 容器名   // mongo命令 数据库名

# 创建最高权限用户
db.createUser({ user: 'admin', pwd: 'admin', roles: [ { role: &quot;root&quot;, db: &quot;admin&quot; } ] });

</code></pre>

<p>（4）测试连通性：</p>

<pre><code>docker run -it --rm --link mongo:mongo mongo mongo -u admin -p admin --authenticationDatabase admin mongo/admin

</code></pre>

<h3 id="mognodb-基本操作">MognoDB 基本操作</h3>

<p>类似 MySQL 命令，显示库列表：</p>

<pre><code>show dbs

</code></pre>

<p>使用某数据库：</p>

<pre><code>use admin

</code></pre>

<p>显示表列表：</p>

<pre><code>show collections

</code></pre>

<p>如果存在 city 表，格式化显示 city 表内容：</p>

<pre><code>db.city.find().pretty()

</code></pre>

<p>如果已经安装后，只要重启即可。</p>

<p>查看已有的镜像：</p>

<pre><code> docker images

</code></pre>

<p><img src="assets/bc15410321f52ee046966e3ecad9f8281524711.png" alt="file" /></p>

<p>然后 docker start mogno 即可，Mongo 是镜像唯一名词。</p>

<h3 id="结构">结构</h3>

<p>类似上面讲的工程搭建，新建一个工程编写此案例，工程如图：</p>

<p><img src="assets/cbf51d78c7955bca4fcc32863ac482881524710.png" alt="file" /></p>

<p>核心目录如下：</p>

<ul>
<li>pom.xml Maven 依赖配置</li>
<li>application.properties 配置文件，配置 mongo 连接属性配置</li>
<li>dao 数据访问层</li>
<li>controller 展示层实现</li>
</ul>

<h3 id="新增-pom-依赖与配置">新增 POM 依赖与配置</h3>

<p>在 pom.xml 配置新的依赖：</p>

<pre><code>    &lt;!-- Spring Boot 响应式 MongoDB 依赖 --&gt;
    &lt;dependency&gt;
      &lt;groupId&gt;org.springframework.boot&lt;/groupId&gt;
      &lt;artifactId&gt;spring-boot-starter-data-mongodb-reactive&lt;/artifactId&gt;
    &lt;/dependency&gt;

    &lt;!-- 模板引擎 Thymeleaf 依赖 --&gt;
    &lt;dependency&gt;
      &lt;groupId&gt;org.springframework.boot&lt;/groupId&gt;
      &lt;artifactId&gt;spring-boot-starter-thymeleaf&lt;/artifactId&gt;
    &lt;/dependency&gt;

</code></pre>

<p>类似配了 MySQL 和 JDBC 驱动，肯定得去配置数据库。在 application.properties 配置中启动 MongoDB 配置。</p>

<p>数据库名为 admin，账号密码也为 admin。</p>

<pre><code>spring.data.mongodb.host=localhost
spring.data.mongodb.database=admin
spring.data.mongodb.port=27017
spring.data.mongodb.username=admin
spring.data.mongodb.password=admin

</code></pre>

<h3 id="mongodb-数据访问层-cityrepository">MongoDB 数据访问层 CityRepository</h3>

<p>修改 CityRepository 类，代码如下：</p>

<pre><code>import org.spring.springboot.domain.City;
import org.springframework.data.mongodb.repository.ReactiveMongoRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface CityRepository extends ReactiveMongoRepository&lt;City, Long&gt; {

    Mono&lt;City&gt; findByCityName(String cityName);

}

</code></pre>

<p>CityRepository 接口只要继承 ReactiveMongoRepository 类即可。</p>

<p>这里实现了通过城市名找出唯一的城市对象方法：</p>

<pre><code>Mono&lt;City&gt; findByCityName(String cityName);

</code></pre>

<p>复杂查询语句实现也很简单，只要依照接口实现规范，即可实现对应 MySQL 的 where 查询语句。这里 findByxxx 的 xxx 可以映射任何字段，包括主键等。</p>

<p>接口的命名是遵循规范的，常用命名规则如下：</p>

<table>
<thead>
<tr>
<th align="left">关键字</th>
<th align="left">方法命名</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">And</td>
<td align="left">findByNameAndPwd</td>
</tr>

<tr>
<td align="left">Or</td>
<td align="left">findByNameOrSex</td>
</tr>

<tr>
<td align="left">Is</td>
<td align="left">findById</td>
</tr>

<tr>
<td align="left">Between</td>
<td align="left">findByIdBetween</td>
</tr>

<tr>
<td align="left">Like</td>
<td align="left">findByNameLike</td>
</tr>

<tr>
<td align="left">NotLike</td>
<td align="left">findByNameNotLike</td>
</tr>

<tr>
<td align="left">OrderBy</td>
<td align="left">findByIdOrderByXDesc</td>
</tr>

<tr>
<td align="left">Not</td>
<td align="left">findByNameNot</td>
</tr>
</tbody>
</table>

<h3 id="处理器类-handler-和控制器类-controller">处理器类 Handler 和控制器类 Controller</h3>

<p>修改下 Handler，代码如下：</p>

<pre><code>@Component
public class CityHandler {

    private final CityRepository cityRepository;

    @Autowired
    public CityHandler(CityRepository cityRepository) {
        this.cityRepository = cityRepository;
    }

    public Mono&lt;City&gt; save(City city) {
        return cityRepository.save(city);
    }

    public Mono&lt;City&gt; findCityById(Long id) {

        return cityRepository.findById(id);
    }

    public Flux&lt;City&gt; findAllCity() {

        return cityRepository.findAll();
    }

    public Mono&lt;City&gt; modifyCity(City city) {

        return cityRepository.save(city);
    }

    public Mono&lt;Long&gt; deleteCity(Long id) {
        cityRepository.deleteById(id);
        return Mono.create(cityMonoSink -&gt; cityMonoSink.success(id));
    }

    public Mono&lt;City&gt; getByCityName(String cityName) {
        return cityRepository.findByCityName(cityName);
    }
}

</code></pre>

<p>新增对应的方法，直接返回 Mono 对象，不需要对 Mono 进行转换，因为 Mono 本身是个对象，可以被 View 层渲染。继续修改控制器类 Controller，代码如下：</p>

<pre><code>    @Autowired
    private CityHandler cityHandler;

    @GetMapping(value = &quot;/{id}&quot;)
    @ResponseBody
    public Mono&lt;City&gt; findCityById(@PathVariable(&quot;id&quot;) Long id) {
        return cityHandler.findCityById(id);
    }

    @GetMapping()
    @ResponseBody
    public Flux&lt;City&gt; findAllCity() {
        return cityHandler.findAllCity();
    }

    @PostMapping()
    @ResponseBody
    public Mono&lt;City&gt; saveCity(@RequestBody City city) {
        return cityHandler.save(city);
    }

    @PutMapping()
    @ResponseBody
    public Mono&lt;City&gt; modifyCity(@RequestBody City city) {
        return cityHandler.modifyCity(city);
    }

    @DeleteMapping(value = &quot;/{id}&quot;)
    @ResponseBody
    public Mono&lt;Long&gt; deleteCity(@PathVariable(&quot;id&quot;) Long id) {
        return cityHandler.deleteCity(id);
    }

    private static final String CITY_LIST_PATH_NAME = &quot;cityList&quot;;
    private static final String CITY_PATH_NAME = &quot;city&quot;;

    @GetMapping(&quot;/page/list&quot;)
    public String listPage(final Model model) {
        final Flux&lt;City&gt; cityFluxList = cityHandler.findAllCity();
        model.addAttribute(&quot;cityList&quot;, cityFluxList);
        return CITY_LIST_PATH_NAME;
    }

    @GetMapping(&quot;/getByName&quot;)
    public String getByCityName(final Model model,
                                @RequestParam(&quot;cityName&quot;) String cityName) {
        final Mono&lt;City&gt; city = cityHandler.getByCityName(cityName);
        model.addAttribute(&quot;city&quot;, city);
        return CITY_PATH_NAME;
    }

</code></pre>

<p>新增 getByName 路径，指向了新的页面 city。使用 @RequestParam 接收 GET 请求入参，接收的参数为 cityName，城市名称。视图返回值 Mono 或者 String 都行。</p>

<h3 id="tymeleaf-视图">Tymeleaf 视图</h3>

<p>然后编写两个视图 city 和 cityList，代码分别如下。</p>

<p>city.html：</p>

<pre><code>&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;zh-CN&quot;&gt;
&lt;head&gt;
    &lt;meta charset=&quot;UTF-8&quot;/&gt;
    &lt;title&gt;城市&lt;/title&gt;
&lt;/head&gt;

&lt;body&gt;

&lt;div&gt;

    &lt;table&gt;
        &lt;legend&gt;
            &lt;strong&gt;城市单个查询&lt;/strong&gt;
        &lt;/legend&gt;
        &lt;tbody&gt;
            &lt;td th:text=&quot;${city.id}&quot;&gt;&lt;/td&gt;
            &lt;td th:text=&quot;${city.provinceId}&quot;&gt;&lt;/td&gt;
            &lt;td th:text=&quot;${city.cityName}&quot;&gt;&lt;/td&gt;
            &lt;td th:text=&quot;${city.description}&quot;&gt;&lt;/td&gt;
        &lt;/tbody&gt;
    &lt;/table&gt;

&lt;/div&gt;

&lt;/body&gt;
&lt;/html&gt;

</code></pre>

<p>cityList.html：</p>

<pre><code>&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;zh-CN&quot;&gt;
&lt;head&gt;
    &lt;meta charset=&quot;UTF-8&quot;/&gt;
    &lt;title&gt;城市列表&lt;/title&gt;
&lt;/head&gt;

&lt;body&gt;

&lt;div&gt;

    &lt;table&gt;
        &lt;legend&gt;
            &lt;strong&gt;城市列表&lt;/strong&gt;
        &lt;/legend&gt;
        &lt;thead&gt;
        &lt;tr&gt;
            &lt;th&gt;城市编号&lt;/th&gt;
            &lt;th&gt;省份编号&lt;/th&gt;
            &lt;th&gt;名称&lt;/th&gt;
            &lt;th&gt;描述&lt;/th&gt;
        &lt;/tr&gt;
        &lt;/thead&gt;
        &lt;tbody&gt;
        &lt;tr th:each=&quot;city : ${cityList}&quot;&gt;
            &lt;td th:text=&quot;${city.id}&quot;&gt;&lt;/td&gt;
            &lt;td th:text=&quot;${city.provinceId}&quot;&gt;&lt;/td&gt;
            &lt;td th:text=&quot;${city.cityName}&quot;&gt;&lt;/td&gt;
            &lt;td th:text=&quot;${city.description}&quot;&gt;&lt;/td&gt;
        &lt;/tr&gt;
        &lt;/tbody&gt;
    &lt;/table&gt;

&lt;/div&gt;

&lt;/body&gt;
&lt;/html&gt;

</code></pre>

<h3 id="运行工程">运行工程</h3>

<p>一个 CRUD 的 Spring Boot Webflux 工程就开发完毕了，下面运行工程验证一下。使用 IDEA 右侧工具栏，单击 Maven Project Tab 按钮，然后单击使用下 Maven 插件的 install 命令；或者使用命令行的形式，在工程根目录下，执行 Maven 清理和安装工程的指令：</p>

<pre><code>cd springboot-webflux-5-thymeleaf-mongodb
mvn clean install

</code></pre>

<p>在控制台中看到成功的输出：</p>

<pre><code>... 省略
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 01:30 min
[INFO] Finished at: 2017-10-15T10:00:54+08:00
[INFO] Final Memory: 31M/174M
[INFO] ------------------------------------------------------------------------

</code></pre>

<p>在 IDEA 中执行 Application 类启动，任意正常模式或者 Debug 模式。可以在控制台看到成功运行的输出：</p>

<pre><code>... 省略
2018-04-10 08:43:39.932  INFO 2052 --- [ctor-http-nio-1] r.ipc.netty.tcp.BlockingNettyContext     : Started HttpServer on /0:0:0:0:0:0:0:0:8080
2018-04-10 08:43:39.935  INFO 2052 --- [           main] o.s.b.web.embedded.netty.NettyWebServer  : Netty started on port(s): 8080
2018-04-10 08:43:39.960  INFO 2052 --- [           main] org.spring.springboot.Application        : Started Application in 6.547 seconds (JVM running for 9.851)

</code></pre>

<p>打开 POST MAN 工具，开发必备，进行下面操作。</p>

<p>新增城市信息 POST <a href="http://127.0.0.1:8080/city：" target="_blank">http://127.0.0.1:8080/city：</a></p>

<p><img src="assets/f69fa5b09730de686ef4837824da48e71523883-1608971233683.png" alt="file" /></p>

<p>打开浏览器，访问 <a href="http://localhost:8080/city/getByName?cityName=杭州，可以看到如图的响应：" target="_blank">http://localhost:8080/city/getByName?cityName=杭州，可以看到如图的响应：</a></p>

<p><img src="assets/adfc7fd931e3c7a20f168a6eddc9c1481524710.png" alt="file" /></p>

<p>继续访问 <a href="http://localhost:8080/city/page/list，发现没有值，那么按照上一篇的内容插入几条数据即可有值，如图：" target="_blank">http://localhost:8080/city/page/list，发现没有值，那么按照上一篇的内容插入几条数据即可有值，如图：</a></p>

<p><img src="assets/66e6502b86c02a83bf56daa93eb9bc1a1524186-1608971233702.png" alt="file" /></p>

<h3 id="总结">总结</h3>

<p>这里初步实现了一个简单的整合，具体复杂的案例我们在后面的综合案例中实现，会很酷炫。下面整合 Redis，基于 Redis 可以实现常用的缓存、锁，下一篇我们将学习如何整合 Reids。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#bbd7d7d7828f8a8a8b8cfbdcd6dad2d795d8d4d6" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1579da5c927771',t:'MTczNDA4Nzg5NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>