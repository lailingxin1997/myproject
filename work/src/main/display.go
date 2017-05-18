package main

import (
	"net/http"
	"fmt"
	"os"
	"io"



)

/*初始化网址*/
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(tpl))
}


const tpl = `<html>
<!-- 四川大学腾讯俱乐部-技术部 -->
<head>
    <meta http-equiv="Content-type" content="text/html; charset=utf-8">
    <title>SCUTEC5.0</title>
</head>

<body>
 <h2>大文件上传有点慢！请等久一点</h2>
<form enctype="multipart/form-data" action="/api/" method="post">
    选择文件：
    <input name="pic" type="file" /><br>
    <input type="submit" value="提交">

</form>
</body>

</html>`
/*接收前端字符串，未实现，字符串好像没有传过来，应该是url写错了*/
func recieve(w http.ResponseWriter, r *http.Request)  {
	dd:=r.Form["data"]
	fmt.Println(dd)
}
/*文件上传后的跳转页面*/
func index2(w http.ResponseWriter, r *http.Request)  {
	//makefile()
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("pic")

	if err != nil {
	fmt.Println(err)
	return
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
	fmt.Println(err)
	return
	}
	defer f.Close()
	io.Copy(f, file)
	mirro(handler.Filename)
	change(handler.Filename)
	small(handler.Filename)
	quality(handler.Filename)
	sharpen(handler.Filename)
	convolve(handler.Filename)
	invert(handler.Filename)
	fliph(handler.Filename)
	gray(handler.Filename)
	w.Write([]byte(tps))
	recieve(w,r)
	//fmt.Fprintln(w, "upload ok!")
/*想把图片转成base64格式再显示，但是没用好像*/
	/*file5,err := os.Open("1.png")
	if err != nil {
		panic(err)
	}
	src,e:=imaging.Decode(file5)
	if e != nil {
		log.Fatalf("Save failed: %v", err)
	}
	emptyBuff:=bytes.NewBuffer(nil)
	jpeg.Encode(emptyBuff,src,nil)
	dist:=make([]byte,50000)
	base64.StdEncoding.Encode(dist,emptyBuff.Bytes())
	string:=string(dist)
	tp3:=constructtp(string)
	w.Write([]byte(tp3))*/

}

	const tp2=
		`<!DOCTYPE html>
		<html>
		<head>
		<meta charset="utf-8" />
		<title>图片</title>
		</head>

		<body>
		<p>图片</p>
		<p>
		<img src="坐标系.png" />
		</p>

		</body>
		</html>`

const tps  =`
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>Insert title here</title>
<script src="http://apps.bdimg.com/libs/jquery/2.1.4/jquery.js"></script>
<script type="text/javascript">
$(document).ready(function(){
        $("button").click(function(){
            $("p").hide();
        });
    });
     $("#a").click(function () {
        var str = "quality";
        alert(str);
        $.ajax({
            type: "POST",
            url: "/recieve",
            data: str,
            datatype:Text,
            success: function (msg) {
                alert("Data Saved: " + msg);
            }
        });
    });
         $("#b").click(function () {
          var str = "small";
              $.ajax(
                {
                   type: "POST",
            url: "/recieve",
            data: str,
            datatype:Text,
            success: function (msg) {
                alert("Data Saved: " + msg);
            }

                });
            });


         $("#c").click(function () {
          var str = "cut";
              $.ajax(
                {
                     type: "POST",
            url: "/recieve",
            data: str,
            datatype:Text,
            success: function (msg) {
                alert("Data Saved: " + msg);
            }
                });
            });

           $("#d").click(function () {
            var str = "invert";
              $.ajax(
                {
                   type: "POST",
            url: "/recieve",
            data: str,
            datatype:Text,
            success: function (msg) {
                alert("Data Saved: " + msg);
            }
                });
            });

      $("#e").click(function () {
       var str = "fourto1";
              $.ajax(
                {
                  type: "POST",
            url: "/recieve",
            data: str,
            datatype:Text,
            success: function (msg) {
                alert("Data Saved: " + msg);
            }
                });
            });
    $("#f").click(function () {
     var str = "mirro";
              $.ajax(
                {
                     type: "POST",
            url: "/recieve",
            data: str,
            datatype:Text,
            success: function (msg) {
                alert("Data Saved: " + msg);
            }
                });
            });

          $("#g").click(function () {
           var str = "convert";
              $.ajax(
                {
                    type: "POST",
            url: "/recieve",
            data: str,
            datatype:Text,
            success: function (msg) {
                alert("Data Saved: " + msg);
            }
                });
            });
</script>
</head>
<body>

   <input type="button" value=" 四 合一 " id="a" />
   <input type="button" value=" 对称 " id="b" />
   <input type="button" value=" 裁 剪 " id="c" />
   <input type="button" value=" 模 糊 " id="d" />
   <input type="button" value=" 锐 化 " id="e" />
   <input type="button" value=" 卷 积 " id="f" />
    <input type="button" value="  灰 调  " id="g" />
   <input type="button" value="  镜 像 " id="h" /><br />
   <h2>很绝望这个页面的点击事件没有实现，图片在本工程下的文件夹查看</h2>
<p>我先走了</p>
  <button>绝望！请点我</button>

</body>
</html> `

