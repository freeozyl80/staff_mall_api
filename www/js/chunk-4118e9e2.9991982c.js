(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-4118e9e2"],{"0abf":function(t,e,r){"use strict";r.d(e,"i",function(){return n}),r.d(e,"j",function(){return i}),r.d(e,"h",function(){return u}),r.d(e,"p",function(){return s}),r.d(e,"a",function(){return c}),r.d(e,"d",function(){return d}),r.d(e,"o",function(){return p}),r.d(e,"m",function(){return m}),r.d(e,"l",function(){return l}),r.d(e,"n",function(){return f}),r.d(e,"k",function(){return g}),r.d(e,"s",function(){return h}),r.d(e,"q",function(){return v}),r.d(e,"b",function(){return D}),r.d(e,"f",function(){return _}),r.d(e,"g",function(){return b}),r.d(e,"r",function(){return I}),r.d(e,"c",function(){return y}),r.d(e,"e",function(){return $});var a=r("2777"),o=r("f121"),n=function(t){var e=t.fid,r=o["apiHost"]+"/manage/firm/detail";return a["a"].request({url:r,method:"get",params:{fid:e}})},i=function(t){var e=t.pageIndex,r=t.pageSize,n=o["apiHost"]+"/manage/firm/list";return a["a"].request({url:n,method:"get",params:{page_index:e,page_size:r}})},u=function(t){var e=o["apiHost"]+"/manage/firm/add",r=new FormData;return r.set("firmname",t.firmName),r.set("firm_realname",t.firmRealName),r.set("balance",t.balance),a["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:r})},s=function(t){var e=o["apiHost"]+"/manage/firm/update",r=new FormData;return r.set("firmname",t.firmName),r.set("firm_realname",t.firmRealName),r.set("balance",t.balance),a["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:r,params:{fid:t.fid}})},c=function(t){var e=t.pageIndex,r=t.pageSize,n=t.auth1,i=o["apiHost"]+"/manage/firm/account";return a["a"].request({url:i,method:"get",params:{fid:n,page_index:e,page_size:r}})},d=function(t){var e=t.auth1,r=t.uid,n=t.uname,i=t.type,u=o["apiHost"]+"/manage/firm/delegate";return a["a"].request({url:u,method:"get",params:{type:i,fid:e,uname:n,uid:r}})},p=function(t){var e=t.pageIndex,r=t.pageSize,n=t.fid,i=o["apiHost"]+"/manage/product/firm/list";return a["a"].request({url:i,method:"get",params:{fid:n,page_index:e,page_size:r}})},m=function(t){var e=t.fid,r=t.pid,n=o["apiHost"]+"/manage/product/firm/detail";return a["a"].request({url:n,method:"get",params:{fid:e,pid:r}})},l=function(t){var e=o["apiHost"]+"/manage/product/firm/new",r=new FormData;return r.set("fid",t.fid),r.set("product_name",t.productName),r.set("product_realname",t.productRealname),r.set("category_id",t.categoryId),r.set("category_name",t.categoryName),r.set("category_realname",t.categoryRealname),r.set("supplier_id",t.supplierId),r.set("supplier_name",t.supplierName),r.set("supplier_realname",t.supplierRealname),r.set("product_price",t.productPrice),r.set("product_count",t.productCount),r.set("product_img",t.productImg),r.set("product_status",t.productStatus),r.set("product_desc",t.productDesc),r.set("product_banner_list",t.productBannerList.join(",")),r.set("product_detail_list",t.productDetailList.join(",")),a["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:r})},f=function(t){var e=o["apiHost"]+"/manage/product/firm/update",r=new FormData;return r.set("fid",t.fid),r.set("pid",t.pid),r.set("product_name",t.productName),r.set("product_realname",t.productRealname),r.set("category_id",t.categoryId),r.set("category_name",t.categoryName),r.set("category_realname",t.categoryRealname),r.set("supplier_id",t.supplierId),r.set("supplier_name",t.supplierName),r.set("supplier_realname",t.supplierRealname),r.set("product_price",t.productPrice),r.set("product_count",t.productCount),r.set("product_img",t.productImg),r.set("product_status",t.productStatus),r.set("product_desc",t.productDesc),r.set("product_banner_list",t.productBannerList.join(",")),r.set("product_detail_list",t.productDetailList.join(",")),a["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:r})},g=function(t){var e=t.fid,r=o["apiHost"]+"/manage/product/firm/categroy/list";return a["a"].request({url:r,method:"get",params:{fid:e,page_index:1,page_size:100}})},h=function(){var t=o["apiHost"]+"/manage/supplier/list";return a["a"].request({url:t,method:"get",params:{page_index:1,page_size:100}})},v=function(t){var e=t.uid,r=t.fid,n=o["apiHost"]+"/manage/firm/staff";return a["a"].request({url:n,method:"get",params:{fid:r,uid:e}})},D=function(t){var e=o["apiHost"]+"/manage/firm/staff/add",r=new FormData;return r.set("name",t.username),r.set("realname",t.userRealname),r.set("pwd",t.password),a["a"].request({url:e,method:"post",headers:{"Content-Type":"application/form-data"},data:r,params:{fid:t.fid}})},_=function(t){var e=t.uid,r=t.fid,n=t.coin,i=o["apiHost"]+"/manage/firm/staff/coin",u=new FormData;return u.set("data",n),a["a"].request({url:i,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:r,uid:e}})},b=function(t){var e=t.uids,r=t.fid,n=t.coin,i=o["apiHost"]+"/manage/firm/staff/list/coin",u=new FormData;return u.set("data",n),a["a"].request({url:i,method:"post",headers:{"Content-Type":"application/form-data"},data:u,params:{fid:r,uids:e.join(",")}})},I=function(t){var e=t.fid,r=t.pageIndex,n=t.pageSize,i=o["apiHost"]+"/manage/order/list";return a["a"].request({url:i,method:"get",params:{fid:e,page_index:r,page_size:n}})},y=function(t){var e=t.fid,r=t.orderId,n=o["apiHost"]+"/manage/order/cancel",i=new FormData;return i.set("orderId",r),a["a"].request({url:n,method:"post",data:i,params:{fid:e}})},$=function(t){var e=t.fid,r=t.orderId,n=o["apiHost"]+"/manage/order/deliver",i=new FormData;return i.set("orderId",r),a["a"].request({url:n,method:"post",data:i,params:{fid:e}})}},9088:function(t,e,r){"use strict";var a=r("fd35"),o=r.n(a);o.a},9512:function(t,e,r){"use strict";r.r(e);var a=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"mg-detail"},[r("Breadcrumb",[r("BreadcrumbItem",{attrs:{to:"/home"}},[t._v("首页")]),r("BreadcrumbItem",{attrs:{to:"/home/association_list"}},[t._v("合作机构列表")]),r("BreadcrumbItem",{attrs:{to:"/home/association_business?fid="+t.$route.query.fid+"&firmname="+t.$route.query.firmname}},[t._v(t._s(t.$route.query.firmname)+": 商品列表")]),t.isEdit?t._e():r("BreadcrumbItem",[t._v("新增 商品")]),t.isEdit?r("BreadcrumbItem",[t._v("编辑 商品入")]):t._e()],1),r("br"),r("div",{staticStyle:{"max-width":"500px","text-align":"left"}},[r("Form",{ref:"formData",attrs:{model:t.formData,rules:t.rules,"label-width":80}},[r("FormItem",{attrs:{label:"商品名（英文）",prop:"productName"}},[r("Input",{attrs:{placeholder:"输入商品英文简写"},model:{value:t.formData.productName,callback:function(e){t.$set(t.formData,"productName",e)},expression:"formData.productName"}})],1),r("FormItem",{attrs:{label:"商品名（全称）",prop:"productRealname"}},[r("Input",{attrs:{placeholder:"输入商品名称"},model:{value:t.formData.productRealname,callback:function(e){t.$set(t.formData,"productRealname",e)},expression:"formData.productRealname"}})],1),r("FormItem",{attrs:{label:"商品类别",prop:"categoryId"}},[r("Select",{attrs:{disabled:t.isEdit},on:{"on-change":t.confimCategory},model:{value:t.formData.categoryId,callback:function(e){t.$set(t.formData,"categoryId",e)},expression:"formData.categoryId"}},t._l(t.categoryOptions,function(e,a){return r("Option",{key:a,attrs:{value:e.cid}},[t._v(t._s(e.category_name)+" ("+t._s(e.category_realname)+")")])}),1)],1),r("FormItem",{attrs:{label:"商品价格",prop:"productPrice"}},[r("Input",{attrs:{placeholder:"请输入商品价格",type:"number"},model:{value:t.formData.productPrice,callback:function(e){t.$set(t.formData,"productPrice",e)},expression:"formData.productPrice"}})],1),r("FormItem",{attrs:{label:"商品库存",prop:"productCount"}},[r("Input",{attrs:{placeholder:"请输入商品库存",type:"number"},model:{value:t.formData.productCount,callback:function(e){t.$set(t.formData,"productCount",e)},expression:"formData.productCount"}})],1),r("FormItem",{attrs:{label:"商品图片",prop:"productImg"}},[r("Input",{attrs:{placeholder:"输入商品图片"},model:{value:t.formData.productImg,callback:function(e){t.$set(t.formData,"productImg",e)},expression:"formData.productImg"}})],1),r("FormItem",{attrs:{label:"商品图片",prop:"productImg"}},[r("Input",{attrs:{placeholder:"输入商品图片"},model:{value:t.formData.productImg,callback:function(e){t.$set(t.formData,"productImg",e)},expression:"formData.productImg"}})],1),r("FormItem",{attrs:{label:"商品状态",prop:"productStatus"}},[r("Select",{model:{value:t.formData.productStatus,callback:function(e){t.$set(t.formData,"productStatus",e)},expression:"formData.productStatus"}},t._l(t.productStatusOptions,function(e,a){return r("Option",{key:a,attrs:{value:e.id}},[t._v(t._s(e.value))])}),1)],1),r("FormItem",{attrs:{label:"商品描述",prop:"productDesc"}},[r("Input",{attrs:{placeholder:"输入商品描述"},model:{value:t.formData.productDesc,callback:function(e){t.$set(t.formData,"productDesc",e)},expression:"formData.productDesc"}})],1),r("FormItem",{attrs:{label:"供应商名称",prop:"supplierId"}},[r("Select",{attrs:{disabled:t.isEdit},on:{"on-change":t.confimSupplier},model:{value:t.formData.supplierId,callback:function(e){t.$set(t.formData,"supplierId",e)},expression:"formData.supplierId"}},t._l(t.supplierOptions,function(e,a){return r("Option",{key:a,attrs:{value:e.supplier_id}},[t._v(t._s(e.supplier_name)+" ("+t._s(e.supplier_realname)+")")])}),1)],1),r("Divider"),r("div",{staticClass:"custom-tag"},t._l(t.formData.productBannerList,function(e,a){return r("Tag",{attrs:{closable:""},on:{"on-close":function(e){return t.handleBannerClose(a)}}},[r("img",{staticStyle:{width:"100px","object-fit":"contain"},attrs:{src:e}})])}),1),r("FormItem",{attrs:{label:"商品轮播图片list",prop:"productBannerItem"}},[r("Input",{attrs:{placeholder:"输入商品图片, 回车确定"},on:{"on-enter":t.enterBanner},model:{value:t.formData.productBannerItem,callback:function(e){t.$set(t.formData,"productBannerItem",e)},expression:"formData.productBannerItem"}})],1),r("div",{staticClass:"custom-tag"},t._l(t.formData.productDetailList,function(e,a){return r("Tag",{attrs:{size:"large",closable:""},on:{"on-close":function(e){return t.handleDetailClose(a)}}},[r("img",{staticStyle:{width:"100px","object-fit":"contain"},attrs:{src:e}})])}),1),r("FormItem",{attrs:{label:"商品详情图片list",prop:"productDetailItem"}},[r("Input",{attrs:{placeholder:"确定，回车确定"},on:{"on-enter":t.enterDetail},model:{value:t.formData.productDetailItem,callback:function(e){t.$set(t.formData,"productDetailItem",e)},expression:"formData.productDetailItem"}})],1),r("FormItem",[t.isEdit?r("Button",{attrs:{type:"primary"},on:{click:function(e){return t.handleSubmit("formData")}}},[t._v("修改商品")]):t._e(),t.isEdit?t._e():r("Button",{attrs:{type:"primary"},on:{click:function(e){return t.handleSubmit("formData")}}},[t._v("新建商品")])],1)],1)],1)],1)},o=[],n=(r("7f7f"),r("c5f6"),r("a481"),r("0abf")),i=function(t){var e={};for(var r in t){var a=t[r],o=r.replace(/\_(\w)/g,function(t,e){return e.toUpperCase()});e[o]=a}return console.log(e),e},u={name:"AssociationBusinessDetail",data:function(){var t=Boolean("new"!=this.$route.query.mode);return{isEdit:t,formData:{productName:"",productRealname:"",categoryId:null,categoryName:"",categoryRealname:"",supplierId:null,supplierName:"",supplierRealname:"",productPrice:0,productCount:0,productImg:"",productStatus:1,productDesc:"",productDetailItem:"",productBannerItem:"",productBannerList:[],productDetailList:[]},rules:{productName:[{required:!0,message:"商品名称不能为空",trigger:"blur"}],productRealname:[{required:!0,message:"商品全称不能为空",trigger:"blur"}],categoryId:[{required:!0,type:"number",message:"请选择商品类别",trigger:"blur"}],supplierId:[{required:!0,type:"number",message:"请选择供应商",trigger:"blur"}],productPrice:[{required:!0,type:"number",message:"商品价格不能为空",trigger:"blur",transform:function(t){return Number(t)}}],productCount:[{required:!0,type:"number",message:"商品数量不能为空",trigger:"blur",transform:function(t){return Number(t)}}],productImg:[{required:!0,message:"商品图片不能为空",trigger:"blur"}],productStatus:[{required:!0,type:"number",message:"请选择商品状态",trigger:"blur"}],productDesc:[{required:!0,message:"商品描述不能为空",trigger:"blur"}]},categoryOptions:[],supplierOptions:[],categoryData:{},SupplierData:{},productStatusOptions:[{id:1,value:"状态正常"},{id:2,value:"状态失效"}]}},watch:{$route:function(t,e){var r=this;"association_business_detail"==t.name&&r.fetch()}},mounted:function(){window.test=this;var t=this;t.fetch()},methods:{handleSubmit:function(t){var e,r=this;e=r.isEdit?n["n"]:n["l"],this.$Modal.confirm({render:function(t){return t("div",[t("h2","确认操作"),t("h4","商品操作确认？")])},onOk:function(){r.$refs[t].validate(function(t){if(t){var a=Object.assign({},r.formData,r.isEdit?{fid:r.$route.query.fid,pid:r.$route.query.pid}:{fid:r.$route.query.fid});e(a).then(function(t){0==t.errorCode&&(r.$Message.success("商品操作成功"),setTimeout(function(){r.$router.push({name:"association_business",query:{fid:r.$route.query.fid,firmname:r.$route.query.firmname}})},2e3))},function(){r.$Message.error("商品操作失败")})}else r.$Message.error("请检查填写")})}})},fetch:function(){var t=this;Object(n["k"])({fid:t.$route.query.fid}).then(function(e){0==e.errorCode?(t.categoryOptions=e.data.list||[],e.data.list.map(function(e){t.categoryData[e.cid]={categoryName:e.category_name,categoryRealname:e.category_realname}})):t.$Message.error(e.errorMsg+":"+e.info)},function(){t.$Message.error("商品列表拉取失败")}),Object(n["s"])().then(function(e){0==e.errorCode?(t.supplierOptions=e.data.list||[],e.data.list.map(function(e){t.SupplierData[e.supplier_id]={supplierName:e.supplier_name,supplierRealname:e.supplier_realname}})):t.$Message.error(e.errorMsg+":"+e.info)},function(){t.$Message.error("供应商列表拉取失败")}),t.isEdit&&Object(n["m"])({fid:t.$route.query.fid,pid:t.$route.query.pid}).then(function(e){0==e.errorCode?(e.data.product_status=parseInt(e.data.productStatus,10),e.data.category_id=parseInt(e.data.productStatus,10),e.data.product_price=parseInt(e.data.productPrice,10),e.data.product_count=parseInt(e.data.productCount,10),t.formData=i(e.data.info)||[]):t.$Message.error(e.errorMsg+":"+e.info)},function(){t.$Message.error("商品详情拉取失败")})},confimCategory:function(t){var e=this;e.formData.categoryName=e.categoryData[t].categoryName,e.formData.categoryRealname=e.categoryData[t].categoryRealname},confimSupplier:function(t){var e=this;e.formData.supplierName=e.SupplierData[t].supplierName,e.formData.supplierRealname=e.SupplierData[t].supplierRealname},enterDetail:function(){var t=this;t.formData.productDetailItem.length>0&&t.formData.productDetailList.push(t.formData.productDetailItem)},enterBanner:function(){var t=this;t.formData.productBannerItem.length>0&&t.formData.productBannerList.push(t.formData.productBannerItem)},handleBannerClose:function(t){var e=this;e.formData.productBannerList.splice(t,1)},handleDetailClose:function(t){var e=this;e.formData.productDetailList.splice(t,1)}}},s=u,c=(r("edd2"),r("9088"),r("2877")),d=Object(c["a"])(s,a,o,!1,null,"e15c9690",null);e["default"]=d.exports},a481:function(t,e,r){"use strict";var a=r("cb7c"),o=r("4bf8"),n=r("9def"),i=r("4588"),u=r("0390"),s=r("5f1b"),c=Math.max,d=Math.min,p=Math.floor,m=/\$([$&`']|\d\d?|<[^>]*>)/g,l=/\$([$&`']|\d\d?)/g,f=function(t){return void 0===t?t:String(t)};r("214f")("replace",2,function(t,e,r,g){return[function(a,o){var n=t(this),i=void 0==a?void 0:a[e];return void 0!==i?i.call(a,n,o):r.call(String(n),a,o)},function(t,e){var o=g(r,t,this,e);if(o.done)return o.value;var p=a(t),m=String(this),l="function"===typeof e;l||(e=String(e));var v=p.global;if(v){var D=p.unicode;p.lastIndex=0}var _=[];while(1){var b=s(p,m);if(null===b)break;if(_.push(b),!v)break;var I=String(b[0]);""===I&&(p.lastIndex=u(m,n(p.lastIndex),D))}for(var y="",$=0,q=0;q<_.length;q++){b=_[q];for(var S=String(b[0]),x=c(d(i(b.index),m.length),0),C=[],w=1;w<b.length;w++)C.push(f(b[w]));var B=b.groups;if(l){var k=[S].concat(C,x,m);void 0!==B&&k.push(B);var N=String(e.apply(void 0,k))}else N=h(S,m,x,C,B,e);x>=$&&(y+=m.slice($,x)+N,$=x+S.length)}return y+m.slice($)}];function h(t,e,a,n,i,u){var s=a+t.length,c=n.length,d=l;return void 0!==i&&(i=o(i),d=m),r.call(u,d,function(r,o){var u;switch(o.charAt(0)){case"$":return"$";case"&":return t;case"`":return e.slice(0,a);case"'":return e.slice(s);case"<":u=i[o.slice(1,-1)];break;default:var d=+o;if(0===d)return r;if(d>c){var m=p(d/10);return 0===m?r:m<=c?void 0===n[m-1]?o.charAt(1):n[m-1]+o.charAt(1):r}u=n[d-1]}return void 0===u?"":u})}})},c9d3:function(t,e,r){},edd2:function(t,e,r){"use strict";var a=r("c9d3"),o=r.n(a);o.a},fd35:function(t,e,r){}}]);