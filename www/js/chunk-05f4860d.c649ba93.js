(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-05f4860d"],{"0abf":function(t,e,a){"use strict";a.d(e,"i",function(){return i}),a.d(e,"j",function(){return o}),a.d(e,"h",function(){return u}),a.d(e,"p",function(){return s}),a.d(e,"a",function(){return d}),a.d(e,"s",function(){return c}),a.d(e,"d",function(){return p}),a.d(e,"o",function(){return f}),a.d(e,"m",function(){return m}),a.d(e,"l",function(){return l}),a.d(e,"n",function(){return g}),a.d(e,"k",function(){return h}),a.d(e,"t",function(){return _}),a.d(e,"q",function(){return v}),a.d(e,"b",function(){return y}),a.d(e,"f",function(){return q}),a.d(e,"g",function(){return w}),a.d(e,"r",function(){return x}),a.d(e,"c",function(){return b}),a.d(e,"e",function(){return D});var r=a("2777"),n=a("f121"),i=function(t){var e=t.fid,a=n["apiHost"]+"/manage/firm/detail";return r["a"].request({url:a,method:"get",params:{fid:e}})},o=function(t){var e=t.pageIndex,a=t.pageSize,i=n["apiHost"]+"/manage/firm/list";return r["a"].request({url:i,method:"get",params:{page_index:e,page_size:a}})},u=function(t){var e=n["apiHost"]+"/manage/firm/add",a=new FormData;return a.set("firmname",t.firmName),a.set("firm_realname",t.firmRealName),a.set("balance",t.balance),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},s=function(t){var e=n["apiHost"]+"/manage/firm/update",a=new FormData;return a.set("firmname",t.firmName),a.set("firm_realname",t.firmRealName),a.set("balance",t.balance),t.productGroup&&a.set("product_group",t.productGroup),t.categoryGroup&&a.set("category_group",t.categoryGroup),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a,params:{fid:t.fid}})},d=function(t){var e=t.pageIndex,a=t.pageSize,i=t.auth1,o=n["apiHost"]+"/manage/firm/account";return r["a"].request({url:o,method:"get",params:{fid:i,page_index:e,page_size:a}})},c=function(t){var e=t.pageIndex,a=t.pageSize,i=t.auth1,o=n["apiHost"]+"/manage/firm/staff/list";return r["a"].request({url:o,method:"get",params:{fid:i,page_index:e,page_size:a}})},p=function(t){var e=t.auth1,a=t.uid,i=t.uname,o=t.type,u=n["apiHost"]+"/manage/firm/delegate";return r["a"].request({url:u,method:"get",params:{type:o,fid:e,uname:i,uid:a}})},f=function(t){var e=t.pageIndex,a=t.pageSize,i=t.fid,o=n["apiHost"]+"/manage/product/firm/list";return r["a"].request({url:o,method:"get",params:{fid:i,page_index:e,page_size:a}})},m=function(t){var e=t.fid,a=t.pid,i=n["apiHost"]+"/manage/product/firm/detail";return r["a"].request({url:i,method:"get",params:{fid:e,pid:a}})},l=function(t){var e=n["apiHost"]+"/manage/product/firm/new",a=new FormData;return a.set("fid",t.fid),a.set("product_name",t.productName),a.set("product_realname",t.productRealname),a.set("category_id",t.categoryId),a.set("category_name",t.categoryName),a.set("category_realname",t.categoryRealname),a.set("supplier_id",t.supplierId),a.set("supplier_name",t.supplierName),a.set("supplier_realname",t.supplierRealname),a.set("product_price",t.productPrice),a.set("product_count",t.productCount),a.set("product_img",t.productImg),a.set("product_status",t.productStatus),a.set("product_desc",t.productDesc),a.set("product_banner_list",t.productBannerList.join(",")),a.set("product_detail_list",t.productDetailList.join(",")),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},g=function(t){var e=n["apiHost"]+"/manage/product/firm/update",a=new FormData;return a.set("fid",t.fid),a.set("pid",t.pid),a.set("product_name",t.productName),a.set("product_realname",t.productRealname),a.set("category_id",t.categoryId),a.set("category_name",t.categoryName),a.set("category_realname",t.categoryRealname),a.set("supplier_id",t.supplierId),a.set("supplier_name",t.supplierName),a.set("supplier_realname",t.supplierRealname),a.set("product_price",t.productPrice),a.set("product_count",t.productCount),a.set("product_img",t.productImg),a.set("product_status",t.productStatus),a.set("product_desc",t.productDesc),a.set("product_banner_list",t.productBannerList.join(",")),a.set("product_detail_list",t.productDetailList.join(",")),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},h=function(t){var e=t.fid,a=n["apiHost"]+"/manage/product/firm/categroy/list";return r["a"].request({url:a,method:"get",params:{fid:e,page_index:1,page_size:100}})},_=function(){var t=n["apiHost"]+"/manage/supplier/list";return r["a"].request({url:t,method:"get",params:{page_index:1,page_size:100}})},v=function(t){var e=t.uid,a=t.fid,i=n["apiHost"]+"/manage/firm/staff";return r["a"].request({url:i,method:"get",params:{fid:a,uid:e}})},y=function(t){var e=n["apiHost"]+"/manage/firm/staff/add",a=new FormData;return a.set("name",t.username),a.set("realname",t.userRealname),a.set("pwd",t.password),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a,params:{fid:t.fid}})},q=function(t){var e=t.uid,a=t.fid,i=t.coin,o=n["apiHost"]+"/manage/firm/staff/coin",u=new FormData;return u.set("data",i),r["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:a,uid:e}})},w=function(t){var e=t.uids,a=t.fid,i=t.coin,o=n["apiHost"]+"/manage/firm/staff/list/coin",u=new FormData;return u.set("data",i),r["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:a,uids:e.join(",")}})},x=function(t){var e=t.fid,a=t.pageIndex,i=t.pageSize,o=n["apiHost"]+"/manage/order/list";return r["a"].request({url:o,method:"get",params:{fid:e,page_index:a,page_size:i}})},b=function(t){var e=t.fid,a=t.orderId,i=n["apiHost"]+"/manage/order/cancel",o=new FormData;return o.set("orderId",a),r["a"].request({url:i,method:"post",data:o,params:{fid:e}})},D=function(t){var e=t.fid,a=t.orderId,i=n["apiHost"]+"/manage/order/deliver",o=new FormData;return o.set("orderId",a),r["a"].request({url:i,method:"post",data:o,params:{fid:e}})}},"5df3":function(t,e,a){"use strict";var r=a("02f4")(!0);a("01f9")(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,a=this._i;return a>=e.length?{value:void 0,done:!0}:(t=r(e,a),this._i+=t.length,{value:t,done:!1})})},"671f":function(t,e,a){},"9caf":function(t,e,a){"use strict";var r=a("671f"),n=a.n(r);n.a},c1db3:function(t,e,a){"use strict";a.r(e);var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"mg-detail"},[a("Breadcrumb",[a("BreadcrumbItem",{attrs:{to:"/home"}},[t._v("首页")]),a("BreadcrumbItem",{attrs:{to:"/home/association_list"}},[t._v("合作机构列表")]),a("BreadcrumbItem",[t._v(t._s(t.$route.query.firmname)+": 员工列表")])],1),a("Button",{staticStyle:{margin:"30px 0 30px"},attrs:{type:"primary"},on:{click:function(e){return t.importStaff()}}},[t._v("+ 导入员工列表")]),t._v("\n       \n"),a("Button",{staticStyle:{margin:"30px 0 30px"},attrs:{type:"default"},on:{click:function(e){return t.addStaff()}}},[t._v("+ 新增单个员工")]),t._v("\n          \n    "),a("span",[t._v("福利点数调整：")]),t._v(" \n"),a("Input",{staticStyle:{width:"100px"},attrs:{placeholder:"福利点数值"},model:{value:t.coin,callback:function(e){t.coin=e},expression:"coin"}}),t._v("\n     \n"),a("Button",{staticStyle:{margin:"30px 0 30px"},attrs:{type:"warning"},on:{click:function(e){return t.addCoin()}}},[t._v("+ 批量增加福利点数")]),a("Table",{attrs:{columns:t.columns,data:t.listData},scopedSlots:t._u([{key:"action",fn:function(e){var r=e.row;e.index;return["1000"==r.auth2?a("Button",{staticClass:"mg-button",attrs:{type:"default",size:"small"},on:{click:function(e){return t.delegate(r,"off")}}},[t._v("取缔管理员")]):t._e(),"1000"!=r.auth2?a("Button",{staticClass:"mg-button",attrs:{type:"primary",size:"small"},on:{click:function(e){return t.delegate(r,"on")}}},[t._v("设定为管理员")]):t._e(),t._v("\n            \n      "),a("Button",{staticClass:"mg-button",attrs:{type:"default",size:"small"},on:{click:function(e){return t.staff(r,"on")}}},[t._v("staff操作")]),t._v("\n\n            \n      "),a("Button",{staticClass:"mg-button",attrs:{type:"warning",size:"small"},on:{click:function(e){return t.reset(r)}}},[t._v("密码重置")])]}}])}),[a("Page",{attrs:{total:t.total,"show-sizer":"false","page-size":t.pageData.pageSize,current:t.pageData.pageIndex},on:{"on-change":t.jump}})]],2)},n=[],i=a("cebc"),o=(a("ac6a"),a("5df3"),a("7f7f"),a("0abf")),u=a("df02"),s={name:"AssociationAccount",data:function(){return{total:0,pageData:{pageSize:20,pageIndex:1},coin:0,columns:[{title:"用户id",key:"uid"},{title:"账户名",key:"username"},{title:"真实姓名",key:"realname"},{title:"手机号",key:"tel"},{title:"福利币点数",key:"coin"},{title:"操作",key:"action",slot:"action",width:500}],listData:[]}},watch:{$route:function(t,e){var a=this;"association_staff"==t&&a.fetch()}},mounted:function(){var t=this;t.fetch()},methods:{importStaff:function(){this.$router.push({name:"account_import",query:{fid:this.$route.query.fid,firmname:this.$route.query.firmname,from:this.$route.name}})},jump:function(t){var e=this;e.pageData.pageIndex=t,e.fetch(e.pageData.pageIndex,e.pageData.pageSize)},fetch:function(t,e){var a=this;Promise.all([Object(o["a"])({auth1:a.$route.query.fid,pageIndex:t||a.pageData.pageIndex,pageSize:e||a.pageData.pageSize}),Object(o["s"])({auth1:a.$route.query.fid,pageIndex:t||a.pageData.pageIndex,pageSize:e||a.pageData.pageSize})]).then(function(t){var e=!0;if(t.forEach(function(t){t.ok||(e=!1,a.$Message.error(t.errorMsg+":"+t.info))}),e){var r=t[0].data.list||[],n=t[1].data.list||[],o=n.map(function(t){var e=t;return r.forEach(function(a){t.uid===a.uid&&(e=Object(i["a"])({},t,e))}),e});a.listData=o||[],a.total=t[1].data.total}})},delegate:function(t,e){var a=this;this.$Modal.confirm({render:function(t){return t("div",[t("h2","确认操作"),t("h4","off"==e?"请确认取缔管理员？":"请确认设定管理员？")])},onOk:function(){Object(o["d"])({type:e,auth1:a.$route.query.fid,uid:t.uid,uanme:t.username}).then(function(t){t.ok?(a.$Message.success("管理员操作成功"),setTimeout(function(){a.fetch()},2e3)):a.$Message.error(t.errorMsg+":"+t.info)})},onCancel:function(){console.log("取消")}})},staff:function(t){var e=this;e.$router.push({name:"association_staff",query:{mode:"edit",fid:e.$route.query.fid,firmname:e.$route.query.firmname,uid:t.uid}})},addStaff:function(){var t=this;t.$router.push({name:"association_staff",query:{mode:"new",fid:t.$route.query.fid,firmname:t.$route.query.firmname}})},addCoin:function(){var t=this,e=t.listData.map(function(t){return t.uid});this.$Modal.confirm({render:function(e){return e("div",[e("h2","确认操作"),e("h4","确认批量添加员工福利点数".concat(t.coin,"?"))])},onOk:function(){Object(o["g"])({fid:t.$route.query.fid,coin:t.coin,uids:e}).then(function(e){0==e.errorCode?(t.$Message.success("员工福利点数操作成功:"),t.fetch()):t.$Message.error(e.errorMsg+":"+e.info)},function(){t.$Message.error("员工福利点数操作失败")})}})},reset:function(t){var e=this;this.$Modal.confirm({render:function(e){return e("div",[e("h2","确认操作"),e("h4","重置操作后，账号(".concat(t.username,")密码将为6个0"))])},onOk:function(){console.log("密码重置"),Object(u["d"])({uid:t.uid}).then(function(t){0==t.errorCode?e.$Message.success("员工账号密码初始化成功"):(console.log(t),e.$Message.error(t.errorMsg+":"+t.info))},function(){e.$Message.error("员工账号密码初始化失败")})}})}}},d=s,c=(a("9caf"),a("2877")),p=Object(c["a"])(d,r,n,!1,null,"407a305a",null);e["default"]=p.exports},df02:function(t,e,a){"use strict";a.d(e,"a",function(){return i}),a.d(e,"b",function(){return o}),a.d(e,"c",function(){return u}),a.d(e,"d",function(){return s}),a.d(e,"e",function(){return d}),a.d(e,"f",function(){return c});a("7f7f");var r=a("2777"),n=a("f121"),i=function(t){var e=t.name,a=t.pwd,i=t.type,o=n["apiHost"]+"/manage/admin/login",u=new FormData;return u.set("name",e),u.set("pwd",a),u.set("type",i),r["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:u})},o=n["apiHost"]+"/manage/user/import",u=function(t){var e=t.pageIndex,a=t.pageSize,i=t.userType,o=n["apiHost"]+"/manage/user/list";return r["a"].request({url:o,method:"get",params:{page_index:e,page_size:a,user_type:i}})},s=function(t){var e=t.uid,a=n["apiHost"]+"/manage/user/reset",i=new FormData;return i.set("uid",e),r["a"].request({url:a,method:"post",headers:{"Content-Type":"application/form-data"},data:i})},d=function(t){var e=t.name,a=t.realname,i=t.pwd,o=n["apiHost"]+"/manage/service/create",u=new FormData;return u.set("name",e),u.set("realname",a),u.set("pwd",i),r["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:u})},c=function(t){var e=t.uid,a=n["apiHost"]+"/manage/user/off",i=new FormData;return i.set("uid",e),r["a"].request({url:a,method:"post",headers:{"Content-Type":"application/form-data"},data:i})}}}]);