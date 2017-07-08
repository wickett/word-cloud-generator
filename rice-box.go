package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1499371281, 0),
		Content:     string("\n<!DOCTYPE html>\n<html>\n  <head>\n    <title>WordCloud</title>\n    <!-- Latest compiled and minified CSS -->\n    <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css\" integrity=\"sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u\" crossorigin=\"anonymous\">\n\n    <!-- Optional theme -->\n    <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap-theme.min.css\" integrity=\"sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp\" crossorigin=\"anonymous\">\n    <link rel=\"stylesheet\" href=\"jqcloud.css\">\n    <script src=\"https://code.jquery.com/jquery-1.10.2.js\"></script>\n    <script type=\"text/javascript\" src=\"jqcloud-1.0.0.min.js\"></script>\n    <style type=\"text/css\">\n      body {\n        background: #eee;\n        font-family: Helvetica, Arial, sans-serif;\n      }\n      #wordcloud {\n        margin: 30px auto;\n        width: 600px;\n        height: 371px;\n        border: none;\n      }\n      #wordcloud span.w10, #wordcloud span.w9, #wordcloud span.w8, #wordcloud span.w7 {\n        text-shadow: 0px 1px 1px #ccc;\n      }\n      #wordcloud span.w3, #wordcloud span.w2, #wordcloud span.w1 {\n        text-shadow: 0px 1px 1px #fff;\n      }\n    </style>\n  </head>\n  <body>\n<form action=\"/api\" id=\"wordcloudForm\" >\n  <input type=\"text\" name=\"text\" placeholder=\"Enter text here...\">\n  <input type=\"submit\" value=\"Submit\">\n</form>\n<div id=\"wordcloud\"></div>\n<script>\n\n$( \"#wordcloudForm\" ).submit(function( event ) {\n  // Stop form from submitting normally\n  event.preventDefault();\n  var word_list = new Array();\n\n  // Get the values from elements on the page:\n  var $form = $( this ),\n  entered = $form.find( \"input[name='text']\" ).val(),\n  url = $form.attr( \"action\" );\n\n  // Send the data using post\n  var posting = $.post( url, { text: entered } );\n\n  posting.done(function( data ) {\n\n    console.log(\"Received response from server...\");\n    console.log(data);\n    $.each(data, function(index, element) {\n      console.log(index, element);\n      word_list.push({text:index, weight:element});\n    });\n    $(\"#wordcloud\").jQCloud(word_list);\n  })\n\n  .complete(function() { console.log(\"complete\"); });\n\n});\n\n</script>\n  </body>\n</html>\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "jqcloud-1.0.0.min.js",
		FileModTime: time.Unix(1499311518, 0),
		Content:     string("/*!\n * jQCloud Plugin for jQuery\n *\n * Version 1.0.0\n *\n * Copyright 2011, Luca Ongaro\n * Licensed under the MIT license.\n *\n * Date: Tue Apr 17 15:06:02 +0200 2012\n*/\n(function(a){\"use strict\",a.fn.jQCloud=function(b,c){var d=this,e=d.attr(\"id\")||Math.floor(Math.random()*1e6).toString(36),f={width:d.width(),height:d.height(),center:{x:(c&&c.width?c.width:d.width())/2,y:(c&&c.height?c.height:d.height())/2},delayedMode:b.length>50,shape:!1};c=a.extend(f,c||{}),d.addClass(\"jqcloud\").width(c.width).height(c.height),d.css(\"position\")===\"static\"&&d.css(\"position\",\"relative\");var g=function(){var f=function(a,b){var c=function(a,b){return Math.abs(2*a.offsetLeft+a.offsetWidth-2*b.offsetLeft-b.offsetWidth)<a.offsetWidth+b.offsetWidth&&Math.abs(2*a.offsetTop+a.offsetHeight-2*b.offsetTop-b.offsetHeight)<a.offsetHeight+b.offsetHeight?!0:!1},d=0;for(d=0;d<b.length;d++)if(c(a,b[d]))return!0;return!1};for(var g=0;g<b.length;g++)b[g].weight=parseFloat(b[g].weight,10);b.sort(function(a,b){return a.weight<b.weight?1:a.weight>b.weight?-1:0});var h=c.shape===\"rectangular\"?18:2,i=[],j=c.width/c.height,k=function(g,k){var l=e+\"_word_\"+g,m=\"#\"+l,n=6.28*Math.random(),o=0,p=0,q=0,r=5,s=\"\",t=\"\",u=\"\";k.html=a.extend(k.html,{id:l}),k.html&&k.html[\"class\"]&&(s=k.html[\"class\"],delete k.html[\"class\"]),b[0].weight>b[b.length-1].weight&&(r=Math.round((k.weight-b[b.length-1].weight)/(b[0].weight-b[b.length-1].weight)*9)+1),u=a(\"<span>\").attr(k.html).addClass(\"w\"+r+\" \"+s),k.link?(typeof k[\"link\"]==\"string\"&&(k.link={href:k.link}),k.link=a.extend(k.link,{href:encodeURI(k.link.href).replace(/'/g,\"%27\")}),t=a(\"<a>\").attr(k.link).text(k.text)):t=k.text,u.append(t);if(!!k.handlers)for(var v in k.handlers)k.handlers.hasOwnProperty(v)&&typeof k.handlers[v]==\"function\"&&a(u).bind(v,k.handlers[v]);d.append(u);var w=u.width(),x=u.height(),y=c.center.x-w/2,z=c.center.y-x/2,A=u[0].style;A.position=\"absolute\",A.left=y+\"px\",A.top=z+\"px\";while(f(document.getElementById(l),i)){if(c.shape===\"rectangular\"){p++,p*h>(1+Math.floor(q/2))*h*(q%4%2===0?1:j)&&(p=0,q++);switch(q%4){case 1:y+=h*j+Math.random()*2;break;case 2:z-=h+Math.random()*2;break;case 3:y-=h*j+Math.random()*2;break;case 0:z+=h+Math.random()*2}}else o+=h,n+=(g%2===0?1:-1)*h,y=c.center.x-w/2+o*Math.cos(n)*j,z=c.center.y+o*Math.sin(n)-x/2;A.left=y+\"px\",A.top=z+\"px\"}i.push(document.getElementById(l)),a.isFunction(k.afterWordRender)&&k.afterWordRender.call(u)},l=function(e){e=e||0,e<b.length?(k(e,b[e]),setTimeout(function(){l(e+1)},10)):a.isFunction(c.afterCloudRender)&&c.afterCloudRender.call(d)};c.delayedMode?l():(a.each(b,k),a.isFunction(c.afterCloudRender)&&c.afterCloudRender.call(d))};return setTimeout(function(){g()},10),d}})(jQuery);"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "jqcloud.css",
		FileModTime: time.Unix(1499311518, 0),
		Content:     string("div#wordcloud {\n  font-family: \"Helvetica\", \"Arial\", sans-serif;\n  color: #09f;\n  overflow: hidden;\n  position: relative;\n}\ndiv#wordcloud a {\n  color: inherit;\n  text-decoration: none;\n}\ndiv#wordcloud a:hover {\n  color: #0df;\n}\ndiv#wordcloud a:hover {\n  color: #0cf;\n}\ndiv#wordcloud span {\n  padding: 0;\n}\ndiv#wordcloud span.w10 {\n  font-size: 54px;\n  color: #0cf;\n}\ndiv#wordcloud span.w9 {\n  font-size: 50px;\n  color: #0cf;\n}\ndiv#wordcloud span.w8 {\n  font-size: 44px;\n  color: #0cf;\n}\ndiv#wordcloud span.w7 {\n  font-size: 40px;\n  color: #39d;\n}\ndiv#wordcloud span.w6 {\n  font-size: 34px;\n  color: #90c5f0;\n}\ndiv#wordcloud span.w5 {\n  font-size: 30px;\n  color: #90a0dd;\n}\ndiv#wordcloud span.w4 {\n  font-size: 24px;\n  color: #90c5f0;\n}\ndiv#wordcloud span.w3 {\n  font-size: 20px;\n  color: #a0ddff;\n}\ndiv#wordcloud span.w2 {\n  font-size: 14px;\n  color: #99ccee;\n}\ndiv#wordcloud span.w1 {\n  font-size: 10px;\n  color: #aab5f0;\n}"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "version",
		FileModTime: time.Unix(1499548250, 0),
		Content:     string("{ \"version\": \"1.DEVELOPMENT\" }\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1499548250, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "index.html"
			file3, // "jqcloud-1.0.0.min.js"
			file4, // "jqcloud.css"
			file5, // "version"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`static`, &embedded.EmbeddedBox{
		Name: `static`,
		Time: time.Unix(1499548250, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"index.html":           file2,
			"jqcloud-1.0.0.min.js": file3,
			"jqcloud.css":          file4,
			"version":              file5,
		},
	})
}
