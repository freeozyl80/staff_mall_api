(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-3f7dd0f5"],{"0abf":function(t,e,a){"use strict";a.d(e,"i",function(){return i}),a.d(e,"j",function(){return o}),a.d(e,"h",function(){return u}),a.d(e,"p",function(){return s}),a.d(e,"a",function(){return d}),a.d(e,"d",function(){return p}),a.d(e,"o",function(){return c}),a.d(e,"m",function(){return m}),a.d(e,"l",function(){return f}),a.d(e,"n",function(){return l}),a.d(e,"k",function(){return g}),a.d(e,"s",function(){return _}),a.d(e,"q",function(){return h}),a.d(e,"b",function(){return y}),a.d(e,"f",function(){return v}),a.d(e,"g",function(){return q}),a.d(e,"r",function(){return w}),a.d(e,"c",function(){return x}),a.d(e,"e",function(){return b});var r=a("2777"),n=a("f121"),i=function(t){var e=t.fid,a=n["apiHost"]+"/manage/firm/detail";return r["a"].request({url:a,method:"get",params:{fid:e}})},o=function(t){var e=t.pageIndex,a=t.pageSize,i=n["apiHost"]+"/manage/firm/list";return r["a"].request({url:i,method:"get",params:{page_index:e,page_size:a}})},u=function(t){var e=n["apiHost"]+"/manage/firm/add",a=new FormData;return a.set("firmname",t.firmName),a.set("firm_realname",t.firmRealName),a.set("balance",t.balance),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},s=function(t){var e=n["apiHost"]+"/manage/firm/update",a=new FormData;return a.set("firmname",t.firmName),a.set("firm_realname",t.firmRealName),a.set("balance",t.balance),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a,params:{fid:t.fid}})},d=function(t){var e=t.pageIndex,a=t.pageSize,i=t.auth1,o=n["apiHost"]+"/manage/firm/account";return r["a"].request({url:o,method:"get",params:{fid:i,page_index:e,page_size:a}})},p=function(t){var e=t.auth1,a=t.uid,i=t.uname,o=t.type,u=n["apiHost"]+"/manage/firm/delegate";return r["a"].request({url:u,method:"get",params:{type:o,fid:e,uname:i,uid:a}})},c=function(t){var e=t.pageIndex,a=t.pageSize,i=t.fid,o=n["apiHost"]+"/manage/product/firm/list";return r["a"].request({url:o,method:"get",params:{fid:i,page_index:e,page_size:a}})},m=function(t){var e=t.fid,a=t.pid,i=n["apiHost"]+"/manage/product/firm/detail";return r["a"].request({url:i,method:"get",params:{fid:e,pid:a}})},f=function(t){var e=n["apiHost"]+"/manage/product/firm/new",a=new FormData;return a.set("fid",t.fid),a.set("product_name",t.productName),a.set("product_realname",t.productRealname),a.set("category_id",t.categoryId),a.set("category_name",t.categoryName),a.set("category_realname",t.categoryRealname),a.set("supplier_id",t.supplierId),a.set("supplier_name",t.supplierName),a.set("supplier_realname",t.supplierRealname),a.set("product_price",t.productPrice),a.set("product_count",t.productCount),a.set("product_img",t.productImg),a.set("product_status",t.productStatus),a.set("product_desc",t.productDesc),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},l=function(t){var e=n["apiHost"]+"/manage/product/firm/update",a=new FormData;return a.set("fid",t.fid),a.set("pid",t.pid),a.set("product_name",t.productName),a.set("product_realname",t.productRealname),a.set("category_id",t.categoryId),a.set("category_name",t.categoryName),a.set("category_realname",t.categoryRealname),a.set("supplier_id",t.supplierId),a.set("supplier_name",t.supplierName),a.set("supplier_realname",t.supplierRealname),a.set("product_price",t.productPrice),a.set("product_count",t.productCount),a.set("product_img",t.productImg),a.set("product_status",t.productStatus),a.set("product_desc",t.productDesc),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},g=function(t){var e=t.fid,a=n["apiHost"]+"/manage/product/firm/categroy/list";return r["a"].request({url:a,method:"get",params:{fid:e,page_index:1,page_size:100}})},_=function(){var t=n["apiHost"]+"/manage/supplier/list";return r["a"].request({url:t,method:"get",params:{page_index:1,page_size:100}})},h=function(t){var e=t.uid,a=t.fid,i=n["apiHost"]+"/manage/firm/staff";return r["a"].request({url:i,method:"get",params:{fid:a,uid:e}})},y=function(t){var e=n["apiHost"]+"/manage/firm/staff/add",a=new FormData;return a.set("name",t.username),a.set("realname",t.userRealname),a.set("pwd",t.password),r["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:a,params:{fid:t.fid}})},v=function(t){var e=t.uid,a=t.fid,i=t.coin,o=n["apiHost"]+"/manage/firm/staff/coin",u=new FormData;return u.set("data",i),r["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:a,uid:e}})},q=function(t){var e=t.uids,a=t.fid,i=t.coin,o=n["apiHost"]+"/manage/firm/staff/list/coin",u=new FormData;return u.set("data",i),r["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:a,uids:e.join(",")}})},w=function(t){var e=t.fid,a=t.pageIndex,i=t.pageSize,o=n["apiHost"]+"/manage/order/list";return r["a"].request({url:o,method:"get",params:{fid:e,page_index:a,page_size:i}})},x=function(t){var e=t.fid,a=t.orderId,i=n["apiHost"]+"/manage/order/cancel",o=new FormData;return o.set("orderId",a),r["a"].request({url:i,method:"post",data:o,params:{fid:e}})},b=function(t){var e=t.fid,a=t.orderId,i=n["apiHost"]+"/manage/order/deliver",o=new FormData;return o.set("orderId",a),r["a"].request({url:i,method:"post",data:o,params:{fid:e}})}},3279:function(t,e,a){"use strict";var r=a("d6f6"),n=a.n(r);n.a},"67f8":function(t,e,a){"use strict";a.r(e);var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"mg-detail"},[a("Breadcrumb",[a("BreadcrumbItem",{attrs:{to:"/home"}},[t._v("首页")]),a("BreadcrumbItem",{attrs:{to:"/home/association_list"}},[t._v("合作机构列表")]),a("BreadcrumbItem",[t._v(t._s(t.$route.query.firmname)+": 商品列表")])],1),a("br"),a("Button",{staticStyle:{margin:"30px 0 30px"},attrs:{type:"primary"},on:{click:function(e){return t.importProduct()}}},[t._v("+ 导入商品列表")]),t._v("  \n"),a("Button",{staticStyle:{margin:"30px 0 30px"},attrs:{type:"default"},on:{click:function(e){return t.createProduct()}}},[t._v("+ 新建单个商品")]),a("Table",{attrs:{columns:t.columns,data:t.listData},scopedSlots:t._u([{key:"product_img",fn:function(t){var e=t.row;t.index;return[a("img",{staticStyle:{width:"100px","object-fit":"contain"},attrs:{src:e.product_img}})]}},{key:"product_status",fn:function(e){var r=e.row;e.index;return[1==r.product_status?a("div",{staticStyle:{color:"#2b85e4"}},[t._v("状态正常")]):a("div",[t._v("已下线")])]}},{key:"action",fn:function(e){var r=e.row;e.index;return[a("Button",{staticClass:"mg-button",attrs:{type:"primary",size:"small"},on:{click:function(e){return t.manage(r)}}},[t._v("管理商品相关")])]}}])}),[a("Page",{attrs:{total:t.total,"show-sizer":"false","page-size":t.pageData.pageSize,current:t.pageData.pageIndex},on:{"on-change":t.jump}})]],2)},n=[],i=(a("7f7f"),a("0abf")),o={name:"AssociationProduct",data:function(){return{total:0,pageData:{pageSize:10,pageIndex:1},columns:[{title:"商品id",key:"pid"},{title:"商品名",key:"product_name"},{title:"商品全称",key:"product_realname"},{title:"商品类别",key:"category_name"},{title:"商品类别全称",key:"category_realname"},{title:"商品价格",key:"product_price"},{title:"商品描述",key:"product_desc"},{title:"商品库存",key:"product_count"},{title:"商品图片",key:"product_img",slot:"product_img",width:100},{title:"商品状态",key:"product_status",slot:"product_status"},{title:"操作",key:"action",slot:"action"}],listData:[]}},watch:{$route:function(t,e){var a=this;"association_business"==t&&a.fetch()}},mounted:function(){var t=this;t.fetch()},methods:{importProduct:function(){this.$router.push({name:"product_import",query:{fid:this.$route.query.fid,firmname:this.$route.query.firmname,from:this.$route.name}})},jump:function(t){var e=this;e.pageData.pageIndex=t,e.fetch(e.pageData.pageIndex,e.pageData.pageSize)},fetch:function(t,e){var a=this;Object(i["o"])({fid:a.$route.query.fid,pageIndex:t||a.pageData.pageIndex,pageSize:e||a.pageData.pageSize}).then(function(t){t.ok?(a.listData=t.data.list||[],a.total=t.data.total):a.$Message.error(t.errorMsg+":"+t.info)})},manage:function(t){var e=this;e.$router.push({name:"association_business_detail",query:{fid:e.$route.query.fid,firmname:e.$route.query.firmname,pid:t.pid,mode:"edit"}})},createProduct:function(){var t=this;t.$router.push({name:"association_business_detail",query:{fid:t.$route.query.fid,firmname:t.$route.query.firmname,mode:"new"}})}}},u=o,s=(a("3279"),a("2877")),d=Object(s["a"])(u,r,n,!1,null,"01446c40",null);e["default"]=d.exports},d6f6:function(t,e,a){}}]);