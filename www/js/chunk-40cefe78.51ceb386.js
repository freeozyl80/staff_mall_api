(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-40cefe78"],{"0abf":function(e,t,a){"use strict";a.d(t,"i",function(){return u}),a.d(t,"j",function(){return i}),a.d(t,"h",function(){return s}),a.d(t,"p",function(){return d}),a.d(t,"a",function(){return c}),a.d(t,"s",function(){return p}),a.d(t,"d",function(){return f}),a.d(t,"o",function(){return m}),a.d(t,"m",function(){return l}),a.d(t,"l",function(){return g}),a.d(t,"n",function(){return h}),a.d(t,"k",function(){return _}),a.d(t,"t",function(){return y}),a.d(t,"q",function(){return v}),a.d(t,"b",function(){return q}),a.d(t,"f",function(){return w}),a.d(t,"g",function(){return D}),a.d(t,"r",function(){return x}),a.d(t,"c",function(){return b}),a.d(t,"e",function(){return I});var r=a("2777"),n=a("f121"),o=n["apiHost"],u=function(e){var t=e.fid,a=o+"/manage/firm/detail";return r["a"].request({url:a,method:"get",params:{fid:t}})},i=function(e){var t=e.pageIndex,a=e.pageSize,n=o+"/manage/firm/list";return r["a"].request({url:n,method:"get",params:{page_index:t,page_size:a}})},s=function(e){var t=o+"/manage/firm/add",a=new FormData;return a.set("firmname",e.firmName),a.set("firm_realname",e.firmRealName),a.set("balance",e.balance),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},d=function(e){var t=o+"/manage/firm/update",a=new FormData;return a.set("firmname",e.firmName),a.set("firm_realname",e.firmRealName),a.set("balance",e.balance),e.productGroup&&a.set("product_group",e.productGroup),e.categoryGroup&&a.set("category_group",e.categoryGroup),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a,params:{fid:e.fid}})},c=function(e){var t=e.pageIndex,a=e.pageSize,n=e.auth1,u=o+"/manage/firm/account";return r["a"].request({url:u,method:"get",params:{fid:n,page_index:t,page_size:a}})},p=function(e){var t=e.pageIndex,a=e.pageSize,n=e.auth1,u=o+"/manage/firm/staff/list";return r["a"].request({url:u,method:"get",params:{fid:n,page_index:t,page_size:a}})},f=function(e){var t=e.auth1,a=e.uid,n=e.uname,u=e.type,i=o+"/manage/firm/delegate";return r["a"].request({url:i,method:"get",params:{type:u,fid:t,uname:n,uid:a}})},m=function(e){var t=e.pageIndex,a=e.pageSize,n=e.fid,u=o+"/manage/product/firm/list";return r["a"].request({url:u,method:"get",params:{fid:n,product_status:1,page_index:t,page_size:a}})},l=function(e){var t=e.fid,a=e.pid,n=o+"/manage/product/firm/detail";return r["a"].request({url:n,method:"get",params:{fid:t,pid:a}})},g=function(e){var t=o+"/manage/product/firm/new",a=new FormData;return a.set("fid",e.fid),a.set("product_name",e.productName),a.set("product_realname",e.productRealname),a.set("category_id",e.categoryId),a.set("category_name",e.categoryName),a.set("category_realname",e.categoryRealname),a.set("supplier_id",e.supplierId),a.set("supplier_name",e.supplierName),a.set("supplier_realname",e.supplierRealname),a.set("product_price",e.productPrice),a.set("product_count",e.productCount),a.set("product_img",e.productImg),a.set("product_status",e.productStatus),a.set("product_desc",e.productDesc),a.set("product_banner_list",e.productBannerList.join(",")),a.set("product_detail_list",e.productDetailList.join(",")),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},h=function(e){var t=o+"/manage/product/firm/update",a=new FormData;return a.set("fid",e.fid),a.set("pid",e.pid),a.set("product_name",e.productName),a.set("product_realname",e.productRealname),a.set("category_id",e.categoryId),a.set("category_name",e.categoryName),a.set("category_realname",e.categoryRealname),a.set("supplier_id",e.supplierId),a.set("supplier_name",e.supplierName),a.set("supplier_realname",e.supplierRealname),a.set("product_price",e.productPrice),a.set("product_count",e.productCount),a.set("product_img",e.productImg),a.set("product_status",e.productStatus),a.set("product_desc",e.productDesc),a.set("product_banner_list",e.productBannerList.join(",")),a.set("product_detail_list",e.productDetailList.join(",")),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},_=function(e){var t=e.fid,a=o+"/manage/product/firm/categroy/list";return r["a"].request({url:a,method:"get",params:{fid:t,page_index:1,page_size:100}})},y=function(){var e=o+"/manage/supplier/list";return r["a"].request({url:e,method:"get",params:{page_index:1,page_size:100}})},v=function(e){var t=e.uid,a=e.fid,n=o+"/manage/firm/staff";return r["a"].request({url:n,method:"get",params:{fid:a,uid:t}})},q=function(e){var t=o+"/manage/firm/staff/add",a=new FormData;return a.set("name",e.username),a.set("realname",e.userRealname),a.set("pwd",e.password),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a,params:{fid:e.fid}})},w=function(e){var t=e.uid,a=e.fid,n=e.coin,u=o+"/manage/firm/staff/coin",i=new FormData;return i.set("data",n),r["a"].request({url:u,method:"post",headers:{"Content-Type":"application/form-data"},data:i,params:{fid:a,uid:t}})},D=function(e){var t=e.uids,a=e.fid,n=e.coin,u=o+"/manage/firm/staff/list/coin",i=new FormData;return i.set("data",n),r["a"].request({url:u,method:"post",headers:{"Content-Type":"application/form-data"},data:i,params:{fid:a,uids:t.join(",")}})},x=function(e){var t=e.fid,a=e.pageIndex,n=e.pageSize,u=o+"/manage/order/list";return r["a"].request({url:u,method:"get",params:{fid:t,page_index:a,page_size:n}})},b=function(e){var t=e.fid,a=e.orderId,n=o+"/manage/order/cancel",u=new FormData;return u.set("orderId",a),r["a"].request({url:n,method:"post",data:u,params:{fid:t}})},I=function(e){var t=e.fid,a=e.orderId,n=o+"/manage/order/deliver",u=new FormData;return u.set("orderId",a),r["a"].request({url:n,method:"post",data:u,params:{fid:t}})}},"25fa":function(e,t,a){"use strict";a.r(t);var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"mg-detail"},[a("Breadcrumb",[a("BreadcrumbItem",{attrs:{to:"/home"}},[e._v("首页")]),a("BreadcrumbItem",{attrs:{to:"/home/association_list"}},[e._v("合作机构列表")]),a("BreadcrumbItem",[e._v(e._s(e.$route.query.firmname)+": 员工列表 - 权限管理")])],1),a("Button",{staticStyle:{margin:"30px 0 30px"},attrs:{type:"primary"},on:{click:function(t){return e.importStaff()}}},[e._v("+ 导入员工列表")]),e._v("\n       \n"),a("Button",{staticStyle:{margin:"30px 0 30px"},attrs:{type:"default"},on:{click:function(t){return e.addStaff()}}},[e._v("+ 新增单个员工")]),a("Table",{attrs:{columns:e.columns,data:e.listData},scopedSlots:e._u([{key:"action",fn:function(t){var r=t.row;t.index;return["1000"==r.auth2?a("Button",{staticClass:"mg-button",attrs:{type:"default",size:"small"},on:{click:function(t){return e.delegate(r,"off")}}},[e._v("取缔管理员")]):e._e(),"1000"!=r.auth2?a("Button",{staticClass:"mg-button",attrs:{type:"primary",size:"small"},on:{click:function(t){return e.delegate(r,"on")}}},[e._v("设定为管理员")]):e._e(),e._v("\n            \n      "),a("Button",{staticClass:"mg-button",attrs:{type:"warning",size:"small"},on:{click:function(t){return e.reset(r)}}},[e._v("密码重置")])]}}])}),[a("Page",{attrs:{total:e.total,"show-sizer":"false","page-size":e.pageData.pageSize,current:e.pageData.pageIndex},on:{"on-change":e.jump}})]],2)},n=[],o=(a("7f7f"),a("0abf")),u=a("df02"),i={name:"AssociationAccount",data:function(){return{total:0,pageData:{pageSize:20,pageIndex:1},coin:0,columns:[{title:"用户id",key:"uid"},{title:"账户名",key:"username"},{title:"操作",key:"action",slot:"action",width:500}],listData:[]}},watch:{$route:function(e,t){var a=this;"association_staff"==e&&a.fetch()}},mounted:function(){var e=this;e.fetch()},methods:{importStaff:function(){this.$router.push({name:"account_import",query:{fid:this.$route.query.fid,firmname:this.$route.query.firmname,from:this.$route.name}})},jump:function(e){var t=this;t.pageData.pageIndex=e,t.fetch(t.pageData.pageIndex,t.pageData.pageSize)},fetch:function(e,t){var a=this;Object(o["a"])({auth1:a.$route.query.fid,pageIndex:e||a.pageData.pageIndex,pageSize:t||a.pageData.pageSize}).then(function(e){e.ok?(a.listData=e.data.list||[],a.total=e.data.total):a.$Message.error(e.errorMsg+":"+e.info)})},delegate:function(e,t){var a=this;this.$Modal.confirm({render:function(e){return e("div",[e("h2","确认操作"),e("h4","off"==t?"请确认取缔管理员？":"请确认设定管理员？")])},onOk:function(){Object(o["d"])({type:t,auth1:a.$route.query.fid,uid:e.uid,uanme:e.username}).then(function(e){e.ok?(a.$Message.success("管理员操作成功"),setTimeout(function(){a.fetch()},2e3)):a.$Message.error(e.errorMsg+":"+e.info)})},onCancel:function(){console.log("取消")}})},addStaff:function(){var e=this;e.$router.push({name:"association_staff",query:{mode:"new",fid:e.$route.query.fid,firmname:e.$route.query.firmname}})},addCoin:function(){var e=this,t=e.listData.map(function(e){return e.uid});this.$Modal.confirm({render:function(t){return t("div",[t("h2","确认操作"),t("h4","确认批量添加员工福利点数".concat(e.coin,"?"))])},onOk:function(){Object(o["g"])({fid:e.$route.query.fid,coin:e.coin,uids:t}).then(function(t){0==t.errorCode?(e.$Message.success("员工福利点数操作成功:"),e.fetch()):e.$Message.error(t.errorMsg+":"+t.info)},function(){e.$Message.error("员工福利点数操作失败")})}})},reset:function(e){var t=this;this.$Modal.confirm({render:function(t){return t("div",[t("h2","确认操作"),t("h4","重置操作后，账号(".concat(e.username,")密码将为6个0"))])},onOk:function(){console.log("密码重置"),Object(u["d"])({uid:e.uid}).then(function(e){0==e.errorCode?t.$Message.success("员工账号密码初始化成功"):(console.log(e),t.$Message.error(e.errorMsg+":"+e.info))},function(){t.$Message.error("员工账号密码初始化失败")})}})}}},s=i,d=(a("7de1"),a("2877")),c=Object(d["a"])(s,r,n,!1,null,"2a63e32e",null);t["default"]=c.exports},"28e6":function(e,t,a){},"7de1":function(e,t,a){"use strict";var r=a("28e6"),n=a.n(r);n.a},df02:function(e,t,a){"use strict";a.d(t,"a",function(){return u}),a.d(t,"b",function(){return i}),a.d(t,"c",function(){return s}),a.d(t,"d",function(){return d}),a.d(t,"e",function(){return c}),a.d(t,"f",function(){return p});a("7f7f");var r=a("2777"),n=a("f121"),o=n["apiHost"],u=function(e){var t=e.name,a=e.pwd,n=e.type,u=o+"/manage/admin/login",i=new FormData;return i.set("name",t),i.set("pwd",a),i.set("type",n),r["a"].request({url:u,method:"post",headers:{"Content-Type":"application/form-data"},data:i})},i=o+"/manage/user/import",s=function(e){var t=e.pageIndex,a=e.pageSize,n=e.userType,u=o+"/manage/user/list";return r["a"].request({url:u,method:"get",params:{page_index:t,page_size:a,user_type:n}})},d=function(e){var t=e.uid,a=o+"/manage/user/reset",n=new FormData;return n.set("uid",t),r["a"].request({url:a,method:"post",headers:{"Content-Type":"application/form-data"},data:n})},c=function(e){var t=e.name,a=e.realname,n=e.pwd,u=o+"/manage/service/create",i=new FormData;return i.set("name",t),i.set("realname",a),i.set("pwd",n),r["a"].request({url:u,method:"post",headers:{"Content-Type":"application/form-data"},data:i})},p=function(e){var t=e.uid,a=o+"/manage/user/off",n=new FormData;return n.set("uid",t),r["a"].request({url:a,method:"post",headers:{"Content-Type":"application/form-data"},data:n})}}}]);