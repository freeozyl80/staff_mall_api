package product

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"staff-mall-center/models/dao"
	"staff-mall-center/models/service/category_service"
	"staff-mall-center/models/service/firm_service"
	"staff-mall-center/models/service/product_service"
	"staff-mall-center/models/service/supplier_service"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"staff-mall-center/pkg/setting"
	"staff-mall-center/pkg/tools"
	"strconv"
	"strings"
)

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ProductList(ctx *context.Context) {

	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))

	total := new(int)
	productlist, err := dao.GetProductList(total, (pageIndex-1)*pageSize, pageSize, "")

	if err != nil {
		code := e.INVALID_PARAMS
		ctx.GenResError(code, "查询商品列表数据失败")
		return
	}

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
		item["supplier_id"] = product_item.SupplierId
		item["supplier_realname"] = product_item.SupplierRealname
		item["supplier_name"] = product_item.SupplierName
		productResList = append(productResList, item)
	}

	values := map[string]interface{}{"page": "page", "pageSize": "pageSize", "succMsg": "查询成功", "list": productResList, "total": *total}

	ctx.GenResSuccess(values)
}

func ProductFirmList(ctx *context.Context) {
	url := fmt.Sprintf("/wx/product/firm/list?fid=%v&page_index=%v&page_size=%v", ctx.Query("fid"), ctx.Query("page_index"), ctx.Query("page_size"))
	ctx.Redirect(http.StatusMovedPermanently, url)
}

func CategroyList(ctx *context.Context) {
	url := fmt.Sprintf("/wx/category/firm/list?fid=%v&page_index=%v&page_size=%v", ctx.Query("fid"), ctx.Query("page_index"), ctx.Query("page_size"))
	ctx.Redirect(http.StatusMovedPermanently, url)
}
func SupplierList(ctx *context.Context) {
	var code int
	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))

	supplierList, err := dao.GetSupplierList((pageIndex-1)*pageSize, pageSize, "")

	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, err.Error())
		return
	}
	var supplierResList []map[string]interface{}

	for _, categrouy_item := range supplierList {
		item := make(map[string]interface{})

		item["supplier_id"] = categrouy_item.ID
		item["supplier_name"] = categrouy_item.SupplierName
		item["supplier_realname"] = categrouy_item.SupplierRealname

		supplierResList = append(supplierResList, item)
	}
	values := map[string]interface{}{"page": pageIndex, "pageSize": pageSize, "list": supplierResList, "succMsg": "查询成功"}

	ctx.GenResSuccess(values)
}
func ProductImport(ctx *context.Context) {
	var code int
	var categoryGroup string
	var productGroup string

	file, info, err := ctx.Request.FormFile("file")
	auth1 := ctx.PostForm("fid")

	if err != nil {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, "请选择文件")
		return
	}

	upload_file_ext := filepath.Ext(info.Filename)

	if upload_file_ext != ".xlsx" {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, "请上传正确格式的文件")
		return
	}

	filename := setting.XlsxSetting.UserXlsx

	filename = path.Join("tmp", filename)

	err = tools.Upload(file, filename)
	// 处理文件
	if err == nil {
		_err, arr := tools.GenerateDataList(filename)

		if _err == nil {

			var supplierList = supplier_service.ArraySupplier{}
			var categoryList = category_service.ArrayCategory{}
			var productList = product_service.ArrayProduct{}

			// 供应商导入
			for num, row := range arr {
				if num != 0 {
					if err != nil {
						break
					}
					supplier_item := supplier_service.Supplier{
						SupplierName:     row[8],
						SupplierRealname: row[9],
						SupplierTel:      0,
					}
					supplierList = append(supplierList, supplier_item)
				}
			}
			_, error := supplierList.BuckRegister()

			if error != nil {
				code = e.INVALID_PARAMS
				ctx.GenResError(code, error.Error())
				return
			}

			// 品类导入
			for num, row := range arr {
				if num != 0 {
					if err != nil {
						break
					}
					category_item := category_service.Category{
						CategoryName:     row[2],
						CategoryRealname: row[3],
						CategoryDesc:     "导入商品 默认品类",
						CategoryStatus:   1,
					}
					categoryList = append(categoryList, category_item)
				}
			}
			_, error = categoryList.BuckRegister()

			if error != nil {
				code = e.INVALID_PARAMS
				ctx.GenResError(code, error.Error())
				return
			}

			// 商品导入

			for num, row := range arr {
				if num != 0 {
					product_category_item := category_service.Category{
						CategoryName: row[2],
					}
					err = product_category_item.FindCategory()

					product_supplier_item := supplier_service.Supplier{
						SupplierName: row[8],
					}
					err = product_supplier_item.FindSupplier()

					if err != nil {
						code = e.INVALID_PARAMS
						ctx.GenResError(code, err.Error())
						return
					}

					product_item_count, _ := strconv.Atoi(row[4])
					product_item_price, _ := strconv.Atoi(row[5])

					product_item := product_service.Product{
						ProductName:      row[0],
						ProductRealname:  row[1],
						CategoryName:     row[2],
						CategoryRealname: row[3],
						ProductCount:     product_item_count,
						ProductPrice:     product_item_price,
						ProductDesc:      row[6],
						ProductImg:       row[7],
						ProductStatus:    1,
						CategoryID:       product_category_item.CID,

						SupplierId:       product_supplier_item.SupplierId,
						SupplierName:     row[8],
						SupplierRealname: row[9],
						SupplierTel:      0,
					}

					productList = append(productList, product_item)

					cid := strconv.Itoa(product_category_item.CID)
					categoryGroup = categoryGroup + cid + ","
				}
			}
			ids, err := productList.BuckRegister()
			for _, id := range ids {
				pid := strconv.Itoa(id)
				productGroup = productGroup + pid + ","
			}
			// 公司导入
			if len(auth1) > 0 {
				var updateFirmValues map[string]interface{}
				fid, _ := strconv.Atoi(auth1)

				var firm_item = firm_service.Firm{
					Fid: fid,
				}
				err = firm_item.FindFirm()

				if err != nil {
					code = e.ERROR
					ctx.GenResError(code, err.Error())
					return
				}

				if len(firm_item.CategoryGroup) >= 1 {

					if len(categoryGroup) >= 1 {
						categoryGroup = categoryGroup[0 : len(categoryGroup)-1]
						categoryGroup = categoryGroup + "," + firm_item.CategoryGroup
					} else {
						categoryGroup = firm_item.CategoryGroup
					}

				}
				if len(firm_item.ProductGroup) >= 1 {
					if len(productGroup) >= 1 {
						productGroup = productGroup[0 : len(productGroup)-1]
						productGroup = productGroup + "," + firm_item.ProductGroup
					} else {
						productGroup = firm_item.ProductGroup
					}
				}

				temCgroup := unique(strings.Split(categoryGroup, ","))
				temPgroup := unique(strings.Split(productGroup, ","))

				categoryGroup = strings.Join(temCgroup, ",")
				categoryGroup = strings.Join(temPgroup, ",")

				updateFirmValues = map[string]interface{}{"category_group": categoryGroup, "product_group": productGroup}

				err := dao.UpdateFirm(fid, updateFirmValues)

				if err != nil {
					code = e.ERROR
					ctx.GenResError(code, err.Error())
					return
				}
			}
			if err != nil {
				code = e.INVALID_PARAMS
				ctx.GenResError(code, err.Error())
				return
			}

			code = e.SUCCESS
			values := map[string]string{"succMsg": "账户批量导入成功, 已存在商品不会被覆盖"}
			ctx.GenResSuccess(values)
			return

		} else {
			code = e.INVALID_PARAMS
			ctx.GenResError(code, "检查xlsx文件格式，默认sheet1放置登录信息")
			return
		}
	}

	if err != nil {
		code = e.ERROR
		ctx.GenResError(code, "")
		return
	}

}

func ProductFirmAdd(ctx *context.Context) {
	var code int
	fid, _ := strconv.Atoi(ctx.PostForm("fid"))

	productName := ctx.PostForm("product_name")
	productRealname := ctx.PostForm("product_realname")
	categoryID, _ := strconv.Atoi(ctx.PostForm("category_id"))
	categoryName := ctx.PostForm("category_name")
	categoryRealname := ctx.PostForm("category_realname")
	productPrice, _ := strconv.Atoi(ctx.PostForm("product_price"))
	productCount, _ := strconv.Atoi(ctx.PostForm("product_count"))
	productImg := ctx.PostForm("product_img")
	productStatus, _ := strconv.Atoi(ctx.PostForm("product_status"))
	productDesc := ctx.PostForm("product_desc")

	supplierId, _ := strconv.Atoi(ctx.PostForm("supplier_id"))
	supplierName := ctx.PostForm("supplier_name")
	supplierRealname := ctx.PostForm("supplier_realname")

	productBannerList := ctx.PostForm("product_banner_list")

	productDetailList := ctx.PostForm("product_detail_list")

	product_item := product_service.Product{
		ProductName:       productName,
		ProductRealname:   productRealname,
		CategoryID:        categoryID,
		CategoryName:      categoryName,
		CategoryRealname:  categoryRealname,
		ProductPrice:      productPrice,
		ProductCount:      productCount,
		ProductImg:        productImg,
		ProductStatus:     productStatus,
		ProductDesc:       productDesc,
		SupplierId:        supplierId,
		SupplierName:      supplierName,
		SupplierRealname:  supplierRealname,
		ProductBannerList: productBannerList,
		ProductDetailList: productDetailList,
	}
	err := product_item.Register()

	firm := firm_service.Firm{
		Fid: fid,
	}
	if err == nil {
		err = firm.FindFirm()
	}

	updateFirmValues := map[string]interface{}{
		"product_group": firm.ProductGroup + strconv.Itoa(product_item.PID) + ",",
	}
	if err == nil {
		err = dao.UpdateFirm(fid, updateFirmValues)
	}

	if err != nil {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, err.Error())
		return
	}

	code = e.SUCCESS
	values := map[string]string{"succMsg": "创建商品成功"}
	ctx.GenResSuccess(values)
	return
}
func ProductFirmUpdate(ctx *context.Context) {

	var code int
	//fid, _ := strconv.Atoi(ctx.PostForm("fid"))
	pid, _ := strconv.Atoi(ctx.PostForm("pid"))

	productName := ctx.PostForm("product_name")
	productRealname := ctx.PostForm("product_realname")
	categoryID, _ := strconv.Atoi(ctx.PostForm("category_id"))
	categoryName := ctx.PostForm("category_name")
	categoryRealname := ctx.PostForm("category_realname")

	supplierId, _ := strconv.Atoi(ctx.PostForm("supplier_id"))
	supplierName := ctx.PostForm("supplier_name")
	supplierRealname := ctx.PostForm("supplier_realname")

	productPrice, _ := strconv.Atoi(ctx.PostForm("product_price"))
	productCount, _ := strconv.Atoi(ctx.PostForm("product_count"))
	productImg := ctx.PostForm("product_img")
	productStatus, _ := strconv.Atoi(ctx.PostForm("product_status"))
	productDesc := ctx.PostForm("product_desc")

	ProductBannerList := ctx.PostForm("product_banner_list")

	ProductDetailList := ctx.PostForm("product_detail_list")

	updateProductValues := map[string]interface{}{
		"product_name":        productName,
		"product_realname":    productRealname,
		"category_id":         categoryID,
		"category_name":       categoryName,
		"category_realname":   categoryRealname,
		"supplier_id":         supplierId,
		"supplier_name":       supplierName,
		"supplier_realname":   supplierRealname,
		"product_price":       productPrice,
		"product_count":       productCount,
		"product_img":         productImg,
		"product_status":      productStatus,
		"product_desc":        productDesc,
		"product_banner_list": ProductBannerList,
		"product_detail_list": ProductDetailList,
	}

	err := dao.UpdateProduct(pid, updateProductValues)

	if err != nil {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, err.Error())
		return
	}

	code = e.SUCCESS
	values := map[string]string{"succMsg": "创建商品成功"}
	ctx.GenResSuccess(values)
	return
}
func ProductFirmDetail(ctx *context.Context) {
	var code int

	PID, _ := strconv.Atoi(ctx.Query("pid"))

	product_item := product_service.Product{
		PID: PID,
	}
	err := product_item.Find()

	if err != nil {
		code = e.INVALID_PARAMS
		ctx.GenResError(code, err.Error())
		return
	}

	code = e.SUCCESS

	var bannerList []string
	var detailList []string
	if len(product_item.ProductDetailList) > 0 {
		detailList = strings.Split(product_item.ProductDetailList, ",")
	} else {
		detailList = []string{}
	}

	if len(product_item.ProductBannerList) > 0 {
		bannerList = strings.Split(product_item.ProductBannerList, ",")
	} else {
		bannerList = []string{}
	}

	productValues := map[string]interface{}{
		"product_name":        product_item.ProductName,
		"product_realname":    product_item.ProductRealname,
		"category_id":         product_item.CategoryID,
		"category_name":       product_item.CategoryName,
		"category_realname":   product_item.CategoryRealname,
		"supplier_id":         product_item.SupplierId,
		"supplier_name":       product_item.SupplierName,
		"supplier_realname":   product_item.SupplierRealname,
		"product_price":       product_item.ProductPrice,
		"product_count":       product_item.ProductCount,
		"product_img":         product_item.ProductImg,
		"product_status":      product_item.ProductStatus,
		"product_desc":        product_item.ProductDesc,
		"product_banner_list": bannerList,
		"product_detail_list": detailList,
	}

	values := map[string]interface{}{"info": productValues}
	ctx.GenResSuccess(values)
	return
}
