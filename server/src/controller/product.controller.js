const Product = require("../model/product.model")
const httpStatus = require("../constants/httpStatus")
const productController = {}

productController.getProductInfo = async (req, res, next) => {
    let productId = req.query.id
    await Product.findById(productId, (err, data) => {
        if(err){
            if(err.kind === "not_found"){
                res.status(httpStatus.NOT_FOUND).send({
                    code: httpStatus.NOT_FOUND,
                    message: "product id not found"
                })
            }else{
                res.status(httpStatus.INTERNAL_SERVER_ERROR).send({
                    code: httpStatus.INTERNAL_SERVER_ERROR, 
                    message: "Error when retrieving Product with this id"
                })
            }
        }
        res.status(httpStatus.OK).send({
            code: httpStatus.OK,
            message: "Get Product information successfully!",
            data: data
        })
    })

}

module.exports = productController