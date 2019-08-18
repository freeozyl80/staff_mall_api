package product

import (
	"fmt"
	"path"
	"path/filepath"
	"staff-mall-center/models/dao"
	"staff-mall-center/models/service/category_service"
	"staff-mall-center/models/service/firm_service"
	"staff-mall-center/models/service/product_service"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"
	"staff-mall-center/pkg/setting"
	"staff-mall-center/pkg/tools"
	"strconv"
	"strings"
)

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
	fmt.Printf("%+v\n", firm_item)

	seachValue := []int{}
	for _, product_item := range strings.Split(firm_item.ProductGroup, ",") {
		pid, _ := strconv.Atoi(product_item)
		seachValue = append(seachValue, pid)
	}

	productlist, err := dao.GetProductList((pageIndex-1)*pageSize, pageSize, seachValue)

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
	values := map[string]interface{}{"page": pageIndex, "pageSize": pageSize, "list": productResList, "succMsg": "查询成功"}

	ctx.GenResSuccess(values)
}
func ProductList(ctx *context.Context) {

	pageIndex, _ := strconv.Atoi(ctx.Query("page_index"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))

	productlist, err := dao.GetProductList((pageIndex-1)*pageSize, pageSize, "")

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

		productResList = append(productResList, item)
	}

	values := map[string]interface{}{"page": "page", "pageSize": "pageSize", "succMsg": "查询成功", "list": productResList}

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

			var categoryList = category_service.ArrayCategory{}
			var productList = product_service.ArrayProduct{}
			//  品类导入
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
			_, error := categoryList.BuckRegister()

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
