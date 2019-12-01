package product

import (
	"staff-mall-center/models/dao"
	"staff-mall-center/models/service/firm_service"
	"staff-mall-center/models/service/product_service"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"strconv"
	"strings"
)

func CategoryFirmList(ctx *context.Context) {
	var code int
	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	fid, _ := strconv.Atoi(ctx.Query("fid"))

	var categoryResList []map[string]interface{}

	if fid == 0 {
		categorylist, err := dao.GetCategoryList((pageIndex-1)*pageSize, pageSize, "")

		if err != nil {
			code = e.ERROR
			ctx.GenResError(code, err.Error())
			return
		}
		for _, categrouy_item := range categorylist {
			item := make(map[string]interface{})

			item["cid"] = categrouy_item.ID
			item["category_name"] = categrouy_item.CategoryName
			item["category_realname"] = categrouy_item.CategoryRealname
			item["category_desc"] = categrouy_item.CategoryDesc

			categoryResList = append(categoryResList, item)
		}

	} else {
		var firm_item = firm_service.Firm{
			Fid: fid,
		}
		err := firm_item.FindFirm()

		if err != nil {
			code = e.ERROR
			ctx.GenResError(code, err.Error())
			return
		}

		seachValue := []int{}
		for _, category_item := range strings.Split(firm_item.CategoryGroup, ",") {
			cid, _ := strconv.Atoi(category_item)
			seachValue = append(seachValue, cid)
		}
		categorylist, err := dao.GetCategoryList((pageIndex-1)*pageSize, pageSize, seachValue)

		for _, categrouy_item := range categorylist {
			item := make(map[string]interface{})

			item["cid"] = categrouy_item.ID
			item["category_name"] = categrouy_item.CategoryName
			item["category_realname"] = categrouy_item.CategoryRealname
			item["category_desc"] = categrouy_item.CategoryDesc

			categoryResList = append(categoryResList, item)
		}
	}
	values := map[string]interface{}{"page": pageIndex, "pageSize": pageSize, "list": categoryResList, "succMsg": "查询成功"}

	ctx.GenResSuccess(values)
}

func ProductFirmList(ctx *context.Context) {
	var code int
	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	fid, _ := strconv.Atoi(ctx.Query("fid"))

	var firm_item = firm_service.Firm{
		Fid: fid,
	}
	err := firm_item.FindFirm()

	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, err.Error())
		return
	}

	seachValue := []int{}
	for _, product_item := range strings.Split(firm_item.ProductGroup, ",") {
		pid, _ := strconv.Atoi(product_item)
		seachValue = append(seachValue, pid)
	}

	total := new(int)
	productlist, err := dao.GetProductList(total, (pageIndex-1)*pageSize, pageSize, seachValue)

	var productResList []map[string]interface{}

	for _, product_item := range productlist {
		item := make(map[string]interface{})

		item["pid"] = product_item.ID
		item["product_name"] = product_item.ProductName
		item["product_realname"] = product_item.ProductRealname
		item["category_name"] = product_item.CategoryName
		item["category_realname"] = product_item.CategoryRealname
		item["product_desc"] = product_item.ProductDesc
		item["product_count"] = product_item.ProductCount
		item["product_status"] = product_item.ProductStatus
		item["product_price"] = product_item.ProductPrice
		item["product_img"] = product_item.ProductImg

		productResList = append(productResList, item)
	}
	values := map[string]interface{}{"page": pageIndex, "pageSize": pageSize, "list": productResList, "total": *total, "succMsg": "查询成功"}

	ctx.GenResSuccess(values)
}

func CategroyProductList(ctx *context.Context) {
	var code int

	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	fid, _ := strconv.Atoi(ctx.Query("fid"))
	cid, _ := strconv.Atoi(ctx.Query("cid"))

	var firm_item = firm_service.Firm{
		Fid: fid,
	}
	err := firm_item.FindFirm()

	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, err.Error())
		return
	}
	seachValue := []int{}
	for _, product_item := range strings.Split(firm_item.ProductGroup, ",") {
		pid, _ := strconv.Atoi(product_item)
		seachValue = append(seachValue, pid)
	}

	total := new(int)
	productlist, err := dao.GetProductList(total, (pageIndex-1)*pageSize, pageSize, seachValue)

	var productResList []map[string]interface{}

	for _, product_item := range productlist {
		item := make(map[string]interface{})

		item["pid"] = product_item.ID
		item["product_name"] = product_item.ProductName
		item["product_realname"] = product_item.ProductRealname
		item["category_name"] = product_item.CategoryName
		item["category_realname"] = product_item.CategoryRealname
		item["product_desc"] = product_item.ProductDesc
		item["product_count"] = product_item.ProductCount
		item["product_status"] = product_item.ProductStatus
		item["product_price"] = product_item.ProductPrice
		item["product_img"] = product_item.ProductImg

		if product_item.CategoryID == cid {
			productResList = append(productResList, item)
		}
	}
	values := map[string]interface{}{"page": pageIndex, "pageSize": pageSize, "list": productResList, "total": *total, "succMsg": "查询成功"}

	ctx.GenResSuccess(values)
}

func ProductInfo(ctx *context.Context) {
	var code int
	productId := ctx.Param("id")
	ProductId, _ := strconv.Atoi(productId)

	product_item := product_service.Product{
		PID: ProductId,
	}

	err := product_item.Find()
	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, err.Error())
		return
	}

	item := make(map[string]interface{})
	item["pid"] = product_item.PID
	item["product_name"] = product_item.ProductName
	item["product_realname"] = product_item.ProductRealname
	item["category_name"] = product_item.CategoryName
	item["category_realname"] = product_item.CategoryRealname
	item["product_desc"] = product_item.ProductDesc
	item["product_count"] = product_item.ProductCount
	item["product_status"] = product_item.ProductStatus
	item["product_price"] = product_item.ProductPrice
	item["product_img"] = product_item.ProductImg
	item["product_detail_banner_list"] = product_item.ProductBannerList
	item["product_detail_list"] = product_item.ProductDetailList

	values := map[string]interface{}{"data": item, "succMsg": "查询成功"}

	ctx.GenResSuccess(values)
}
