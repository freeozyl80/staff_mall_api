(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-39d157b5"],{"35c8":function(e,t,a){"use strict";var n=a("81c4"),r=a.n(n);r.a},"81c4":function(e,t,a){},"8e36":function(e,t,a){"use strict";a.r(t);var n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"mg-list"},[a("Divider",{attrs:{orientation:"left"}},[e._v("客服账号列表")]),a("Button",{staticStyle:{margin:"30px 0 30px"},attrs:{type:"default"},on:{click:function(t){return e.addService()}}},[e._v("+ 新增客服")]),a("br"),a("Table",{attrs:{columns:e.columns1,data:e.listData},scopedSlots:e._u([{key:"usertype",fn:function(t){var n=t.row;return[4==n.usertype?a("p",[e._v("客服")]):e._e()]}},{key:"action",fn:function(t){var n=t.row;t.index;return[a("Button",{staticClass:"mg-button",attrs:{type:"default",size:"small"},on:{click:function(t){return e.detail(n,"on")}}},[e._v("客服操作")])]}}])}),[a("Page",{attrs:{total:e.total,"show-sizer":"false","page-size":e.pageData.pageSize,current:e.pageData.pageIndex},on:{"on-change":e.jump}})]],2)},r=[],i=a("df02"),o={name:"AccountList",mounted:function(){var e=this;e.fetch()},data:function(){return{total:0,pageData:{pageSize:10,pageIndex:1},columns1:[{title:"用户id",key:"id"},{title:"用户类型",key:"usertype",slot:"usertype"},{title:"账户名",key:"username"},{title:"姓名",key:"realname"},{title:"操作",key:"action",slot:"action"}],listData:[]}},watch:{$route:"fetch"},methods:{jump:function(e){var t=this;t.pageData.pageIndex=e,t.fetch(t.pageData.pageIndex,t.pageData.pageSize)},fetch:function(e,t){var a=this;Object(i["c"])({pageIndex:e||a.pageData.pageIndex,pageSize:t||a.pageData.pageSize,userType:4}).then(function(e){e.ok?(a.listData=e.data.list,a.total=e.data.total):a.$Message.error(e.errorMsg+":"+e.info)})},addService:function(){var e=this;e.$router.push({name:"service_detail",query:{mode:"new"}})},detail:function(e){var t=this;t.$router.push({name:"service_detail",query:{mode:"edit",uid:e.id,username:e.username,realname:e.realname}})}}},s=o,u=(a("35c8"),a("2877")),p=Object(u["a"])(s,n,r,!1,null,"1fbae234",null);t["default"]=p.exports},df02:function(e,t,a){"use strict";a.d(t,"a",function(){return i}),a.d(t,"b",function(){return o}),a.d(t,"c",function(){return s}),a.d(t,"d",function(){return u}),a.d(t,"e",function(){return p}),a.d(t,"f",function(){return c});a("7f7f");var n=a("2777"),r=a("f121"),i=function(e){var t=e.name,a=e.pwd,i=e.type,o=r["apiHost"]+"/manage/admin/login",s=new FormData;return s.set("name",t),s.set("pwd",a),s.set("type",i),n["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:s})},o=r["apiHost"]+"/manage/user/import",s=function(e){var t=e.pageIndex,a=e.pageSize,i=e.userType,o=r["apiHost"]+"/manage/user/list";return n["a"].request({url:o,method:"get",params:{page_index:t,page_size:a,user_type:i}})},u=function(e){var t=e.uid,a=r["apiHost"]+"/manage/user/reset",i=new FormData;return i.set("uid",t),n["a"].request({url:a,method:"post",headers:{"Content-Type":"application/form-data"},data:i})},p=function(e){var t=e.name,a=e.realname,i=e.pwd,o=r["apiHost"]+"/manage/service/create",s=new FormData;return s.set("name",t),s.set("realname",a),s.set("pwd",i),n["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:s})},c=function(e){var t=e.uid,a=r["apiHost"]+"/manage/user/off",i=new FormData;return i.set("uid",t),n["a"].request({url:a,method:"post",headers:{"Content-Type":"application/form-data"},data:i})}}}]);