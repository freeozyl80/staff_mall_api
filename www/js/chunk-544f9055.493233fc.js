(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-544f9055"],{"0abf":function(e,t,a){"use strict";a.d(t,"i",function(){return o}),a.d(t,"j",function(){return i}),a.d(t,"h",function(){return u}),a.d(t,"p",function(){return s}),a.d(t,"a",function(){return d}),a.d(t,"s",function(){return c}),a.d(t,"d",function(){return m}),a.d(t,"o",function(){return f}),a.d(t,"m",function(){return p}),a.d(t,"l",function(){return l}),a.d(t,"n",function(){return g}),a.d(t,"k",function(){return h}),a.d(t,"t",function(){return _}),a.d(t,"q",function(){return v}),a.d(t,"b",function(){return b}),a.d(t,"f",function(){return y}),a.d(t,"g",function(){return D}),a.d(t,"r",function(){return q}),a.d(t,"c",function(){return w}),a.d(t,"e",function(){return I});var r=a("2777"),n=a("f121"),o=function(e){var t=e.fid,a=n["apiHost"]+"/manage/firm/detail";return r["a"].request({url:a,method:"get",params:{fid:t}})},i=function(e){var t=e.pageIndex,a=e.pageSize,o=n["apiHost"]+"/manage/firm/list";return r["a"].request({url:o,method:"get",params:{page_index:t,page_size:a}})},u=function(e){var t=n["apiHost"]+"/manage/firm/add",a=new FormData;return a.set("firmname",e.firmName),a.set("firm_realname",e.firmRealName),a.set("balance",e.balance),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},s=function(e){var t=n["apiHost"]+"/manage/firm/update",a=new FormData;return a.set("firmname",e.firmName),a.set("firm_realname",e.firmRealName),a.set("balance",e.balance),e.productGroup&&a.set("product_group",e.productGroup),e.categoryGroup&&a.set("category_group",e.categoryGroup),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a,params:{fid:e.fid}})},d=function(e){var t=e.pageIndex,a=e.pageSize,o=e.auth1,i=n["apiHost"]+"/manage/firm/account";return r["a"].request({url:i,method:"get",params:{fid:o,page_index:t,page_size:a}})},c=function(e){var t=e.pageIndex,a=e.pageSize,o=e.auth1,i=n["apiHost"]+"/manage/firm/staff/list";return r["a"].request({url:i,method:"get",params:{fid:o,page_index:t,page_size:a}})},m=function(e){var t=e.auth1,a=e.uid,o=e.uname,i=e.type,u=n["apiHost"]+"/manage/firm/delegate";return r["a"].request({url:u,method:"get",params:{type:i,fid:t,uname:o,uid:a}})},f=function(e){var t=e.pageIndex,a=e.pageSize,o=e.fid,i=n["apiHost"]+"/manage/product/firm/list";return r["a"].request({url:i,method:"get",params:{fid:o,page_index:t,page_size:a}})},p=function(e){var t=e.fid,a=e.pid,o=n["apiHost"]+"/manage/product/firm/detail";return r["a"].request({url:o,method:"get",params:{fid:t,pid:a}})},l=function(e){var t=n["apiHost"]+"/manage/product/firm/new",a=new FormData;return a.set("fid",e.fid),a.set("product_name",e.productName),a.set("product_realname",e.productRealname),a.set("category_id",e.categoryId),a.set("category_name",e.categoryName),a.set("category_realname",e.categoryRealname),a.set("supplier_id",e.supplierId),a.set("supplier_name",e.supplierName),a.set("supplier_realname",e.supplierRealname),a.set("product_price",e.productPrice),a.set("product_count",e.productCount),a.set("product_img",e.productImg),a.set("product_status",e.productStatus),a.set("product_desc",e.productDesc),a.set("product_banner_list",e.productBannerList.join(",")),a.set("product_detail_list",e.productDetailList.join(",")),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},g=function(e){var t=n["apiHost"]+"/manage/product/firm/update",a=new FormData;return a.set("fid",e.fid),a.set("pid",e.pid),a.set("product_name",e.productName),a.set("product_realname",e.productRealname),a.set("category_id",e.categoryId),a.set("category_name",e.categoryName),a.set("category_realname",e.categoryRealname),a.set("supplier_id",e.supplierId),a.set("supplier_name",e.supplierName),a.set("supplier_realname",e.supplierRealname),a.set("product_price",e.productPrice),a.set("product_count",e.productCount),a.set("product_img",e.productImg),a.set("product_status",e.productStatus),a.set("product_desc",e.productDesc),a.set("product_banner_list",e.productBannerList.join(",")),a.set("product_detail_list",e.productDetailList.join(",")),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a})},h=function(e){var t=e.fid,a=n["apiHost"]+"/manage/product/firm/categroy/list";return r["a"].request({url:a,method:"get",params:{fid:t,page_index:1,page_size:100}})},_=function(){var e=n["apiHost"]+"/manage/supplier/list";return r["a"].request({url:e,method:"get",params:{page_index:1,page_size:100}})},v=function(e){var t=e.uid,a=e.fid,o=n["apiHost"]+"/manage/firm/staff";return r["a"].request({url:o,method:"get",params:{fid:a,uid:t}})},b=function(e){var t=n["apiHost"]+"/manage/firm/staff/add",a=new FormData;return a.set("name",e.username),a.set("realname",e.userRealname),a.set("pwd",e.password),r["a"].request({url:t,method:"post",headers:{"Content-Type":"application/form-data"},data:a,params:{fid:e.fid}})},y=function(e){var t=e.uid,a=e.fid,o=e.coin,i=n["apiHost"]+"/manage/firm/staff/coin",u=new FormData;return u.set("data",o),r["a"].request({url:i,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:a,uid:t}})},D=function(e){var t=e.uids,a=e.fid,o=e.coin,i=n["apiHost"]+"/manage/firm/staff/list/coin",u=new FormData;return u.set("data",o),r["a"].request({url:i,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:a,uids:t.join(",")}})},q=function(e){var t=e.fid,a=e.pageIndex,o=e.pageSize,i=n["apiHost"]+"/manage/order/list";return r["a"].request({url:i,method:"get",params:{fid:t,page_index:a,page_size:o}})},w=function(e){var t=e.fid,a=e.orderId,o=n["apiHost"]+"/manage/order/cancel",i=new FormData;return i.set("orderId",a),r["a"].request({url:o,method:"post",data:i,params:{fid:t}})},I=function(e){var t=e.fid,a=e.orderId,o=n["apiHost"]+"/manage/order/deliver",i=new FormData;return i.set("orderId",a),r["a"].request({url:o,method:"post",data:i,params:{fid:t}})}},"1b3d":function(e,t,a){},"1c09":function(e,t,a){"use strict";a.r(t);var r,n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"mg-staff"},[a("Breadcrumb",[a("BreadcrumbItem",{attrs:{to:"/home"}},[e._v("首页")]),a("BreadcrumbItem",{attrs:{to:"/home/association_list"}},[e._v("合作机构列表")]),a("BreadcrumbItem",{attrs:{to:"/home/association_account?fid="+e.$route.query.fid+"&firmname="+e.$route.query.firmname}},[e._v(e._s(e.$route.query.firmname)+": 员工列表")]),a("BreadcrumbItem",[e._v("机构"+e._s(e.$route.query.firmname)+"：员工操作页面")])],1),a("br"),e.isEdit?a("div",[a("div",{staticClass:"staff-info"},[a("div",[e._v("员工姓名: "),a("span",[e._v(e._s(e.staffInfo.username))])]),a("div",[e._v("福利点数: "),a("span",[e._v(e._s(e.staffInfo.coi))])])]),a("div",{staticStyle:{"max-width":"500px","text-align":"left"}}),a("Form",{ref:"formData",attrs:{model:e.formData,rules:e.rules,"label-width":100,width:"400"}},[a("FormItem",{attrs:{label:"福利点数调整",prop:"coin"}},[a("Input",{attrs:{placeholder:"增加或减少福利点数（减少请输入负数，比如 -100）",number:""},model:{value:e.formData.coin,callback:function(t){e.$set(e.formData,"coin",t)},expression:"formData.coin"}})],1),a("FormItem",[a("Button",{attrs:{type:"primary"},on:{click:function(t){return e.handleSubmit("formData")}}},[e._v("确定")])],1)],1)],1):a("div",[a("Form",{ref:"formAddData",attrs:{model:e.formAddData,rules:e.addRules,"label-width":100,width:"400"}},[a("FormItem",{attrs:{label:"员工账号名称（username）",prop:"username"}},[a("Input",{attrs:{placeholder:"输入员工账号名称（英文拼音）"},model:{value:e.formAddData.username,callback:function(t){e.$set(e.formAddData,"username",t)},expression:"formAddData.username"}})],1),a("FormItem",{attrs:{label:"员工姓名（userRealname）",prop:"userRealname"}},[a("Input",{attrs:{placeholder:"输入员工姓名"},model:{value:e.formAddData.userRealname,callback:function(t){e.$set(e.formAddData,"userRealname",t)},expression:"formAddData.userRealname"}})],1),a("FormItem",{attrs:{label:"员工密码（password）",prop:"password"}},[a("Input",{attrs:{placeholder:"输入员工密码"},model:{value:e.formAddData.password,callback:function(t){e.$set(e.formAddData,"password",t)},expression:"formAddData.password"}})],1),a("FormItem",[a("Button",{attrs:{type:"primary"},on:{click:function(t){return e.handleAddSubmit("formAddData")}}},[e._v("确定")])],1)],1)],1)],1)},o=[],i=a("bd86"),u=(a("7f7f"),a("c5f6"),a("7cdf"),a("0abf")),s=(r={name:"staffDetail",data:function(){var e=Boolean("new"!=this.$route.query.mode),t=function(e,t,a){if(!t)return a(new Error("coin cannot be empty"));Number.isInteger(t)?a():a(new Error("Please enter a numeric value"))};return{isEdit:e,staffInfo:{},formData:{coin:null},formAddData:{username:"",userRealname:"",password:null},rules:{coin:[{validator:t,trigger:"blur"}]},addRules:{username:[{required:!0,message:"账号名称不能为空",trigger:"blur"}],userRealname:[{required:!0,message:"用户姓名不能为空",trigger:"blur"}],password:[{required:!0,message:"密码不能为空",trigger:"blur"}]}}},mounted:function(){console.log("This is staffDetail")},watch:{$route:function(e,t){var a=this;"association_staff"==e&&a.fetch()}}},Object(i["a"])(r,"mounted",function(){var e=this;e.fetch()}),Object(i["a"])(r,"methods",{fetch:function(){var e=this;e.isEdit&&Object(u["q"])({fid:e.$route.query.fid,uid:e.$route.query.uid}).then(function(t){t.ok?e.staffInfo=t.data.info:e.$Message.error({content:t.errorMsg+":"+t.info,closable:!0,duration:0})})},handleAddSubmit:function(e){var t=this;this.$Modal.confirm({render:function(e){return e("div",[e("h2","确认操作"),e("h4","确认创建员工".concat(t.formAddData.username,"?"))])},onOk:function(){t.$refs[e].validate(function(e){e?Object(u["b"])({fid:t.$route.query.fid,username:t.formAddData.username,userRealname:t.formAddData.userRealname,password:t.formAddData.password}).then(function(e){0==e.errorCode?(t.$Message.success("员工创建成功, 5s后跳转回员工列表页:"),setTimeout(function(){t.$router.push({name:"association_account",query:{fid:t.$route.query.fid,firmname:t.$route.query.firmname}})},5e3)):t.$Message.error({content:e.errorMsg+":"+e.info,closable:!0,duration:0})},function(){t.$Message.error({content:"员工创建成功失败",closable:!0,duration:0})}):t.$Message.error({content:"请检查填写",closable:!0,duration:0})})}})},handleSubmit:function(e){var t=this;this.$Modal.confirm({render:function(e){return e("div",[e("h2","确认操作"),e("h4","确认添加员工福利点数".concat(t.formData.coin||0,"?"))])},onOk:function(){t.$refs[e].validate(function(e){e?Object(u["f"])({fid:t.$route.query.fid,uid:t.$route.query.uid,coin:t.formData.coin||0}).then(function(e){0==e.errorCode?(t.$Message.success("员工福利点数操作成功:"),t.fetch()):(console.log(e),t.$Message.error({content:e.errorMsg+":"+e.info,closable:!0,duration:0}))},function(){t.$Message.error({content:"员工福利点数操作失败",closable:!0,duration:0})}):t.$Message.error({content:"请检查填写",closable:!0,duration:0})})}})}}),r),d=s,c=(a("817d"),a("2877")),m=Object(c["a"])(d,n,o,!1,null,"23f5509e",null);t["default"]=m.exports},"7cdf":function(e,t,a){var r=a("5ca1");r(r.S,"Number",{isInteger:a("9c12")})},"817d":function(e,t,a){"use strict";var r=a("1b3d"),n=a.n(r);n.a},"9c12":function(e,t,a){var r=a("d3f4"),n=Math.floor;e.exports=function(e){return!r(e)&&isFinite(e)&&n(e)===e}}}]);