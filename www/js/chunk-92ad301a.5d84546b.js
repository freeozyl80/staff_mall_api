(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-92ad301a"],{"5d2c":function(t,e,r){"use strict";r.d(e,"a",function(){return n}),r.d(e,"b",function(){return u});var o=r("2777"),a=r("f121"),n=a["apiHost"]+"/manage/product/import",u=function(t){var e=t.pageIndex,r=t.pageSize,n=a["apiHost"]+"/manage/product/list";return o["a"].request({url:n,method:"get",params:{page_index:e,page_size:r}})}},"74a8":function(t,e,r){},"9f0f":function(t,e,r){"use strict";var o=r("74a8"),a=r.n(o);a.a},a195:function(t,e,r){"use strict";r.r(e);var o=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"mg-import"},[r("Divider",{attrs:{orientation:"left"}},[t._v("批量导入 商品列表")]),t._m(0),r("div",{staticStyle:{"margin-top":"20px"}},[r("Upload",{attrs:{type:"drag",data:t.$route.query.fid?{fid:t.$route.query.fid}:{},headers:t.headers,"show-upload-list":!1,action:t.uploadUrl,"on-success":t.uploadCb}},[r("div",{staticStyle:{padding:"20px 0"}},[r("Icon",{staticStyle:{color:"#3399ff"},attrs:{type:"ios-cloud-upload",size:"52"}}),r("p",[t._v("导入商品列表")])],1)])],1)],1)},a=[function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"tips"},[r("p",[t._v("1. 导入文件必须使用xlsx 文件")]),r("p",[t._v("2. 默认sheet1放置登录信息")]),r("p",[t._v("3. 登录信息依次为商品简称，商品全称，商品类别简称，商品类别全称， 库存，价格，商品描述，商品图片")])])}],n=r("f121"),u=r.n(n),i=r("a78e"),s=r.n(i),c=r("5d2c"),d={name:"PooductImport",mounted:function(){console.log("This is PooductImport")},computed:{uploadUrl:function(){return c["a"]},headers:function(){return{hualvmall_authorization:s.a.get(u.a.loginOpts.cookieKey)||""}}},methods:{uploadCb:function(t){var e=this;0!=t.errorCode?e.$Message.error(t.errorMsg+":"+t.info):(e.$Message.success("上传成功"),setTimeout(function(){e.$route.query.from?e.$router.push({name:e.$route.query.from,query:{fid:e.$route.query.fid,firmname:e.$route.query.firmname}}):e.$router.push({name:"product_list"})},1e3))}}},p=d,l=(r("9f0f"),r("2877")),f=Object(l["a"])(p,o,a,!1,null,"64610676",null);e["default"]=f.exports}}]);