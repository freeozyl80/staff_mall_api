(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-1d737178"],{"0abf":function(t,e,a){"use strict";a.d(e,"g",function(){return i}),a.d(e,"h",function(){return o}),a.d(e,"f",function(){return u}),a.d(e,"n",function(){return s}),a.d(e,"a",function(){return d}),a.d(e,"c",function(){return c}),a.d(e,"m",function(){return p}),a.d(e,"k",function(){return f}),a.d(e,"j",function(){return m}),a.d(e,"l",function(){return l}),a.d(e,"i",function(){return g}),a.d(e,"q",function(){return h}),a.d(e,"o",function(){return _}),a.d(e,"d",function(){return y}),a.d(e,"e",function(){return v}),a.d(e,"p",function(){return q}),a.d(e,"b",function(){return b});var r=a("2777"),n=a("f121"),i=function(t){var e=t.fid,a=n["apiHost"]+"/manage/firm/detail";return r["a"].request({url:a,method:"get",params:{fid:e}})},o=function(t){var e=t.pageIndex,a=t.pageSize,i=n["apiHost"]+"/manage/firm/list";return r["a"].request({url:i,method:"get",params:{page_index:e,page_size:a}})},u=function(t){var e=n["apiHost"]+"/manage/firm/add",a=new FormData;return a.set("firmname",t.firmName),a.set("firm_realname",t.firmRealName),a.set("balance",t.balance),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},s=function(t){var e=n["apiHost"]+"/manage/firm/update",a=new FormData;return a.set("firmname",t.firmName),a.set("firm_realname",t.firmRealName),a.set("balance",t.balance),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a,params:{fid:t.fid}})},d=function(t){var e=t.pageIndex,a=t.pageSize,i=t.auth1,o=n["apiHost"]+"/manage/firm/account";return r["a"].request({url:o,method:"get",params:{fid:i,page_index:e,page_size:a}})},c=function(t){var e=t.auth1,a=t.uid,i=t.uname,o=t.type,u=n["apiHost"]+"/manage/firm/delegate";return r["a"].request({url:u,method:"get",params:{type:o,fid:e,uname:i,uid:a}})},p=function(t){var e=t.pageIndex,a=t.pageSize,i=t.fid,o=n["apiHost"]+"/manage/product/firm/list";return r["a"].request({url:o,method:"get",params:{fid:i,page_index:e,page_size:a}})},f=function(t){var e=t.fid,a=t.pid,i=n["apiHost"]+"/manage/product/firm/detail";return r["a"].request({url:i,method:"get",params:{fid:e,pid:a}})},m=function(t){var e=n["apiHost"]+"/manage/product/firm/new",a=new FormData;return a.set("fid",t.fid),a.set("product_name",t.productName),a.set("product_realname",t.productRealname),a.set("category_id",t.categoryId),a.set("category_name",t.categoryName),a.set("category_realname",t.categoryRealname),a.set("supplier_id",t.supplierId),a.set("supplier_name",t.supplierName),a.set("supplier_realname",t.supplierRealname),a.set("product_price",t.productPrice),a.set("product_count",t.productCount),a.set("product_img",t.productImg),a.set("product_status",t.productStatus),a.set("product_desc",t.productDesc),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},l=function(t){var e=n["apiHost"]+"/manage/product/firm/update",a=new FormData;return a.set("fid",t.fid),a.set("pid",t.pid),a.set("product_name",t.productName),a.set("product_realname",t.productRealname),a.set("category_id",t.categoryId),a.set("category_name",t.categoryName),a.set("category_realname",t.categoryRealname),a.set("supplier_id",t.supplierId),a.set("supplier_name",t.supplierName),a.set("supplier_realname",t.supplierRealname),a.set("product_price",t.productPrice),a.set("product_count",t.productCount),a.set("product_img",t.productImg),a.set("product_status",t.productStatus),a.set("product_desc",t.productDesc),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},g=function(t){var e=t.fid,a=n["apiHost"]+"/manage/product/firm/categroy/list";return r["a"].request({url:a,method:"get",params:{fid:e,page_index:1,page_size:100}})},h=function(){var t=n["apiHost"]+"/manage/supplier/list";return r["a"].request({url:t,method:"get",params:{page_index:1,page_size:100}})},_=function(t){var e=t.uid,a=t.fid,i=n["apiHost"]+"/manage/firm/staff";return r["a"].request({url:i,method:"get",params:{fid:a,uid:e}})},y=function(t){var e=t.uid,a=t.fid,i=t.coin,o=n["apiHost"]+"/manage/firm/staff/coin",u=new FormData;return u.set("data",i),r["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:a,uid:e}})},v=function(t){var e=t.uids,a=t.fid,i=t.coin,o=n["apiHost"]+"/manage/firm/staff/list/coin",u=new FormData;return u.set("data",i),r["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:a,uids:e.join(",")}})},q=function(t){var e=t.fid,a=t.pageIndex,i=t.pageSize,o=n["apiHost"]+"/manage/order/list";return r["a"].request({url:o,method:"get",params:{fid:e,page_index:a,page_size:i}})},b=function(t){var e=t.fid,a=t.orderId,i=n["apiHost"]+"/manage/order/cancel",o=new FormData;return o.set("orderId",a),r["a"].request({url:i,method:"get",data:o,params:{fid:e}})}},8462:function(t,e,a){},b1c0:function(t,e,a){"use strict";var r=a("8462"),n=a.n(r);n.a},c1db3:function(t,e,a){"use strict";a.r(e);var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"mg-detail"},[a("Divider",{attrs:{orientation:"left"}},[t._v(t._s(t.$route.query.firmname)+": 员工列表")]),a("Button",{staticStyle:{margin:"30px 0 30px"},attrs:{type:"primary"},on:{click:function(e){return t.importStaff()}}},[t._v("+ 导入员工列表")]),a("br"),a("Input",{staticStyle:{width:"100px"},attrs:{placeholder:"福利点数值"},model:{value:t.coin,callback:function(e){t.coin=e},expression:"coin"}}),t._v("\n  \n"),a("Button",{staticStyle:{margin:"30px 0 30px"},attrs:{type:"warning"},on:{click:function(e){return t.addCoin()}}},[t._v("+ 批量增加福利点数")]),a("Table",{attrs:{columns:t.columns,data:t.listData},scopedSlots:t._u([{key:"action",fn:function(e){var r=e.row;e.index;return["1000"==r.auth2?a("Button",{staticClass:"mg-button",attrs:{type:"default",size:"small"},on:{click:function(e){return t.delegate(r,"off")}}},[t._v("取缔管理员")]):t._e(),"1000"!=r.auth2?a("Button",{staticClass:"mg-button",attrs:{type:"primary",size:"small"},on:{click:function(e){return t.delegate(r,"on")}}},[t._v("设定为管理员")]):t._e(),t._v("\n            \n      "),a("Button",{staticClass:"mg-button",attrs:{type:"default",size:"small"},on:{click:function(e){return t.staff(r,"on")}}},[t._v("staff操作")])]}}])})],1)},n=[],i=(a("7f7f"),a("0abf")),o={name:"AssociationAccount",data:function(){return{coin:0,columns:[{title:"用户id",key:"id"},{title:"账户名",key:"username"},{title:"操作",key:"action",slot:"action",width:500}],listData:[]}},watch:{$route:function(t,e){var a=this;"association_staff"==t&&a.fetch()}},mounted:function(){var t=this;t.fetch()},methods:{importStaff:function(){this.$router.push({name:"account_import",query:{fid:this.$route.query.fid,firmname:this.$route.query.firmname,from:this.$route.name}})},fetch:function(){var t=this;Object(i["a"])({auth1:t.$route.query.fid,pageIndex:1,pageSize:10}).then(function(e){e.ok?t.listData=e.data.list||[]:t.$Message.error(e.errorMsg+":"+e.info)})},delegate:function(t,e){var a=this;Object(i["c"])({type:e,auth1:a.$route.query.fid,uid:t.id,uanme:t.username}).then(function(t){t.ok?(a.$Message.success("管理员操作成功"),setTimeout(function(){a.fetch()},1e3)):a.$Message.error(t.errorMsg+":"+t.info)})},staff:function(t){var e=this;e.$router.push({name:"association_staff",query:{fid:e.$route.query.fid,firmname:e.$route.query.firmname,uid:t.id}})},addCoin:function(){var t=this,e=t.listData.map(function(t){return t.id});Object(i["e"])({fid:t.$route.query.fid,coin:t.coin,uids:e}).then(function(e){0==e.errorCode?(t.$Message.success("员工福利点数操作成功:"),t.fetch()):(console.log(e),t.$Message.error(e.errorMsg+":"+e.info))},function(){t.$Message.error("员工福利点数操作失败")})}}},u=o,s=(a("b1c0"),a("2877")),d=Object(s["a"])(u,r,n,!1,null,"9d058844",null);e["default"]=d.exports}}]);