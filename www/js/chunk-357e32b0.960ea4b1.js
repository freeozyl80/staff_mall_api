(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-357e32b0"],{"0abf":function(e,t,a){"use strict";a.d(t,"g",function(){return i}),a.d(t,"h",function(){return o}),a.d(t,"f",function(){return u}),a.d(t,"n",function(){return d}),a.d(t,"a",function(){return s}),a.d(t,"c",function(){return c}),a.d(t,"m",function(){return p}),a.d(t,"k",function(){return m}),a.d(t,"j",function(){return f}),a.d(t,"l",function(){return l}),a.d(t,"i",function(){return _}),a.d(t,"q",function(){return g}),a.d(t,"o",function(){return v}),a.d(t,"d",function(){return h}),a.d(t,"e",function(){return y}),a.d(t,"p",function(){return q}),a.d(t,"b",function(){return b});var r=a("2777"),n=a("f121"),i=function(e){var t=e.fid,a=n["apiHost"]+"/manage/firm/detail";return r["a"].request({url:a,method:"get",params:{fid:t}})},o=function(e){var t=e.pageIndex,a=e.pageSize,i=n["apiHost"]+"/manage/firm/list";return r["a"].request({url:i,method:"get",params:{page_index:t,page_size:a}})},u=function(e){var t=n["apiHost"]+"/manage/firm/add",a=new FormData;return a.set("firmname",e.firmName),a.set("firm_realname",e.firmRealName),a.set("balance",e.balance),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},d=function(e){var t=n["apiHost"]+"/manage/firm/update",a=new FormData;return a.set("firmname",e.firmName),a.set("firm_realname",e.firmRealName),a.set("balance",e.balance),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a,params:{fid:e.fid}})},s=function(e){var t=e.pageIndex,a=e.pageSize,i=e.auth1,o=n["apiHost"]+"/manage/firm/account";return r["a"].request({url:o,method:"get",params:{fid:i,page_index:t,page_size:a}})},c=function(e){var t=e.auth1,a=e.uid,i=e.uname,o=e.type,u=n["apiHost"]+"/manage/firm/delegate";return r["a"].request({url:u,method:"get",params:{type:o,fid:t,uname:i,uid:a}})},p=function(e){var t=e.pageIndex,a=e.pageSize,i=e.fid,o=n["apiHost"]+"/manage/product/firm/list";return r["a"].request({url:o,method:"get",params:{fid:i,page_index:t,page_size:a}})},m=function(e){var t=e.fid,a=e.pid,i=n["apiHost"]+"/manage/product/firm/detail";return r["a"].request({url:i,method:"get",params:{fid:t,pid:a}})},f=function(e){var t=n["apiHost"]+"/manage/product/firm/new",a=new FormData;return a.set("fid",e.fid),a.set("product_name",e.productName),a.set("product_realname",e.productRealname),a.set("category_id",e.categoryId),a.set("category_name",e.categoryName),a.set("category_realname",e.categoryRealname),a.set("supplier_id",e.supplierId),a.set("supplier_name",e.supplierName),a.set("supplier_realname",e.supplierRealname),a.set("product_price",e.productPrice),a.set("product_count",e.productCount),a.set("product_img",e.productImg),a.set("product_status",e.productStatus),a.set("product_desc",e.productDesc),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},l=function(e){var t=n["apiHost"]+"/manage/product/firm/update",a=new FormData;return a.set("fid",e.fid),a.set("pid",e.pid),a.set("product_name",e.productName),a.set("product_realname",e.productRealname),a.set("category_id",e.categoryId),a.set("category_name",e.categoryName),a.set("category_realname",e.categoryRealname),a.set("supplier_id",e.supplierId),a.set("supplier_name",e.supplierName),a.set("supplier_realname",e.supplierRealname),a.set("product_price",e.productPrice),a.set("product_count",e.productCount),a.set("product_img",e.productImg),a.set("product_status",e.productStatus),a.set("product_desc",e.productDesc),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},_=function(e){var t=e.fid,a=n["apiHost"]+"/manage/product/firm/categroy/list";return r["a"].request({url:a,method:"get",params:{fid:t,page_index:1,page_size:100}})},g=function(){var e=n["apiHost"]+"/manage/supplier/list";return r["a"].request({url:e,method:"get",params:{page_index:1,page_size:100}})},v=function(e){var t=e.uid,a=e.fid,i=n["apiHost"]+"/manage/firm/staff";return r["a"].request({url:i,method:"get",params:{fid:a,uid:t}})},h=function(e){var t=e.uid,a=e.fid,i=e.coin,o=n["apiHost"]+"/manage/firm/staff/coin",u=new FormData;return u.set("data",i),r["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:a,uid:t}})},y=function(e){var t=e.uids,a=e.fid,i=e.coin,o=n["apiHost"]+"/manage/firm/staff/list/coin",u=new FormData;return u.set("data",i),r["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:a,uids:t.join(",")}})},q=function(e){var t=e.fid,a=e.pageIndex,i=e.pageSize,o=n["apiHost"]+"/manage/order/list";return r["a"].request({url:o,method:"get",params:{fid:t,page_index:a,page_size:i}})},b=function(e){var t=e.fid,a=e.orderId,i=n["apiHost"]+"/manage/order/cancel",o=new FormData;return o.set("orderId",a),r["a"].request({url:i,method:"get",data:o,params:{fid:t}})}},"51eb":function(e,t,a){"use strict";a.r(t);var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"mg-detail"},[a("Divider",{attrs:{orientation:"left"}},[e._v(e._s(e.$route.query.firmname)+": 订单列表")]),a("Table",{attrs:{columns:e.columns,data:e.listData},scopedSlots:e._u([{key:"status",fn:function(t){var r=t.row;t.index;return[1==r.order_status?a("div",[e._v("下定")]):e._e(),2==r.order_status?a("div",[e._v("支付完成")]):e._e(),3==r.order_status?a("div",[e._v("发货完成")]):e._e(),4==r.order_status?a("div",[e._v("收货完成")]):e._e(),5==r.order_status?a("div",[e._v("取消")]):e._e(),6==r.order_status?a("div",[e._v("退换中")]):e._e(),7==r.order_status?a("div",[e._v("7异常")]):e._e()]}},{key:"action",fn:function(t){var r=t.row;t.index;return[5!=r.order_status?a("Button",{staticClass:"mg-button",attrs:{type:"error",size:"small"},on:{click:function(t){return e.cancelOrder(r)}}},[e._v("立即退款")]):e._e()]}}])})],1)},n=[],i=a("0abf"),o={name:"AssociationAccount",data:function(){return{columns:[{title:"订单id",key:"order_id"},{title:"收货地址",key:"order_receiving_address"},{title:"收货城市",key:"order_receiving_city"},{title:"收货电话",key:"order_receiving_tel"},{title:"收货人",key:"order_receiving_username"},{title:"订单价格",key:"order_total_price"},{title:"订单状态",key:"oder_status",slot:"status"},{title:"操作",key:"action",slot:"action"}],listData:[]}},watch:{$route:"fetch"},mounted:function(){var e=this;e.fetch()},methods:{fetch:function(){var e=this;Object(i["p"])({fid:e.$route.query.fid,pageIndex:1,pageSize:10}).then(function(t){t.ok?e.listData=t.data.list||[]:e.$Message.error(t.errorMsg+":"+t.info)})},cancelOrder:function(e){var t=this;Object(i["b"])({fid:t.$route.query.fid,orderId:e.order_id}).then(function(e){e.ok?(t.listData=e.data.list||[],t.fetch()):t.$Message.error(e.errorMsg+":"+e.info)})}}},u=o,d=(a("d788"),a("2877")),s=Object(d["a"])(u,r,n,!1,null,"715941ff",null);t["default"]=s.exports},"79ed":function(e,t,a){},d788:function(e,t,a){"use strict";var r=a("79ed"),n=a.n(r);n.a}}]);