(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-08c3796e"],{"0abf":function(e,t,a){"use strict";a.d(t,"g",function(){return i}),a.d(t,"h",function(){return o}),a.d(t,"f",function(){return u}),a.d(t,"n",function(){return s}),a.d(t,"a",function(){return d}),a.d(t,"c",function(){return c}),a.d(t,"m",function(){return f}),a.d(t,"k",function(){return p}),a.d(t,"j",function(){return m}),a.d(t,"l",function(){return l}),a.d(t,"i",function(){return g}),a.d(t,"q",function(){return _}),a.d(t,"o",function(){return h}),a.d(t,"d",function(){return v}),a.d(t,"e",function(){return y}),a.d(t,"p",function(){return b}),a.d(t,"b",function(){return q});var r=a("2777"),n=a("f121"),i=function(e){var t=e.fid,a=n["apiHost"]+"/manage/firm/detail";return r["a"].request({url:a,method:"get",params:{fid:t}})},o=function(e){var t=e.pageIndex,a=e.pageSize,i=n["apiHost"]+"/manage/firm/list";return r["a"].request({url:i,method:"get",params:{page_index:t,page_size:a}})},u=function(e){var t=n["apiHost"]+"/manage/firm/add",a=new FormData;return a.set("firmname",e.firmName),a.set("firm_realname",e.firmRealName),a.set("balance",e.balance),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},s=function(e){var t=n["apiHost"]+"/manage/firm/update",a=new FormData;return a.set("firmname",e.firmName),a.set("firm_realname",e.firmRealName),a.set("balance",e.balance),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a,params:{fid:e.fid}})},d=function(e){var t=e.pageIndex,a=e.pageSize,i=e.auth1,o=n["apiHost"]+"/manage/firm/account";return r["a"].request({url:o,method:"get",params:{fid:i,page_index:t,page_size:a}})},c=function(e){var t=e.auth1,a=e.uid,i=e.uname,o=e.type,u=n["apiHost"]+"/manage/firm/delegate";return r["a"].request({url:u,method:"get",params:{type:o,fid:t,uname:i,uid:a}})},f=function(e){var t=e.pageIndex,a=e.pageSize,i=e.fid,o=n["apiHost"]+"/manage/product/firm/list";return r["a"].request({url:o,method:"get",params:{fid:i,page_index:t,page_size:a}})},p=function(e){var t=e.fid,a=e.pid,i=n["apiHost"]+"/manage/product/firm/detail";return r["a"].request({url:i,method:"get",params:{fid:t,pid:a}})},m=function(e){var t=n["apiHost"]+"/manage/product/firm/new",a=new FormData;return a.set("fid",e.fid),a.set("product_name",e.productName),a.set("product_realname",e.productRealname),a.set("category_id",e.categoryId),a.set("category_name",e.categoryName),a.set("category_realname",e.categoryRealname),a.set("supplier_id",e.supplierId),a.set("supplier_name",e.supplierName),a.set("supplier_realname",e.supplierRealname),a.set("product_price",e.productPrice),a.set("product_count",e.productCount),a.set("product_img",e.productImg),a.set("product_status",e.productStatus),a.set("product_desc",e.productDesc),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},l=function(e){var t=n["apiHost"]+"/manage/product/firm/update",a=new FormData;return a.set("fid",e.fid),a.set("pid",e.pid),a.set("product_name",e.productName),a.set("product_realname",e.productRealname),a.set("category_id",e.categoryId),a.set("category_name",e.categoryName),a.set("category_realname",e.categoryRealname),a.set("supplier_id",e.supplierId),a.set("supplier_name",e.supplierName),a.set("supplier_realname",e.supplierRealname),a.set("product_price",e.productPrice),a.set("product_count",e.productCount),a.set("product_img",e.productImg),a.set("product_status",e.productStatus),a.set("product_desc",e.productDesc),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},g=function(e){var t=e.fid,a=n["apiHost"]+"/manage/product/firm/categroy/list";return r["a"].request({url:a,method:"get",params:{fid:t,page_index:1,page_size:100}})},_=function(){var e=n["apiHost"]+"/manage/supplier/list";return r["a"].request({url:e,method:"get",params:{page_index:1,page_size:100}})},h=function(e){var t=e.uid,a=e.fid,i=n["apiHost"]+"/manage/firm/staff";return r["a"].request({url:i,method:"get",params:{fid:a,uid:t}})},v=function(e){var t=e.uid,a=e.fid,i=e.coin,o=n["apiHost"]+"/manage/firm/staff/coin",u=new FormData;return u.set("data",i),r["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:a,uid:t}})},y=function(e){var t=e.uids,a=e.fid,i=e.coin,o=n["apiHost"]+"/manage/firm/staff/list/coin",u=new FormData;return u.set("data",i),r["a"].request({url:o,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:a,uids:t.join(",")}})},b=function(e){var t=e.fid,a=e.pageIndex,i=e.pageSize,o=n["apiHost"]+"/manage/order/list";return r["a"].request({url:o,method:"get",params:{fid:t,page_index:a,page_size:i}})},q=function(e){var t=e.fid,a=e.orderId,i=n["apiHost"]+"/manage/order/cancel",o=new FormData;return o.set("orderId",a),r["a"].request({url:i,method:"get",data:o,params:{fid:t}})}},"1c09":function(e,t,a){"use strict";a.r(t);var r,n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"mg-staff"},[a("Divider",{attrs:{orientation:"left"}},[e._v("机构"+e._s(e.$route.query.firmname)+"：员工操作页面")]),a("div",{staticClass:"staff-info"},[a("div",[e._v("staff 姓名: "),a("span",[e._v(e._s(e.staffInfo.username))])]),a("div",[e._v("剩余福利点数: "),a("span",[e._v(e._s(e.staffInfo.coi))])])]),a("div",{staticStyle:{"max-width":"500px","text-align":"left"}}),a("Form",{ref:"formData",attrs:{model:e.formData,rules:e.rules,"label-width":100,width:"400"}},[a("FormItem",{attrs:{label:"增加福利点数",prop:"coin"}},[a("Input",{attrs:{placeholder:"输入coin",number:""},model:{value:e.formData.coin,callback:function(t){e.$set(e.formData,"coin",t)},expression:"formData.coin"}})],1),a("FormItem",[a("Button",{attrs:{type:"primary"},on:{click:function(t){return e.handleSubmit("formData")}}},[e._v("确定")])],1)],1)],1)},i=[],o=a("bd86"),u=(a("7f7f"),a("c5f6"),a("7cdf"),a("0abf")),s=(r={name:"staffDetail",data:function(){var e=function(e,t,a){if(!t)return a(new Error("coin cannot be empty"));Number.isInteger(t)?t<0?a(new Error("Must be over of o coin")):a():a(new Error("Please enter a numeric value"))};return{staffInfo:{},formData:{coin:100},rules:{coin:[{validator:e,trigger:"blur"}]}}},mounted:function(){console.log("This is staffDetail")},watch:{$route:function(e,t){var a=this;"association_staff"==e&&a.fetch()}}},Object(o["a"])(r,"mounted",function(){var e=this;e.fetch()}),Object(o["a"])(r,"methods",{fetch:function(){var e=this;Object(u["o"])({fid:e.$route.query.fid,uid:e.$route.query.uid}).then(function(t){t.ok?e.staffInfo=t.data.info:e.$Message.error(t.errorMsg+":"+t.info)})},handleSubmit:function(e){var t=this;this.$refs[e].validate(function(e){e?Object(u["d"])({fid:t.$route.query.fid,uid:t.$route.query.uid,coin:t.formData.coin}).then(function(e){0==e.errorCode?(t.$Message.success("员工福利点数操作成功:"),t.fetch()):(console.log(e),t.$Message.error(e.errorMsg+":"+e.info))},function(){t.$Message.error("员工福利点数操作失败")}):t.$Message.error("请检查填写")})}}),r),d=s,c=(a("68ce"),a("2877")),f=Object(c["a"])(d,n,i,!1,null,"742a2360",null);t["default"]=f.exports},"68ce":function(e,t,a){"use strict";var r=a("cf3c"),n=a.n(r);n.a},"7cdf":function(e,t,a){var r=a("5ca1");r(r.S,"Number",{isInteger:a("9c12")})},"9c12":function(e,t,a){var r=a("d3f4"),n=Math.floor;e.exports=function(e){return!r(e)&&isFinite(e)&&n(e)===e}},cf3c:function(e,t,a){}}]);