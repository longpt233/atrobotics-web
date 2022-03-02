const sql = require("./db")

const Product = function(product){
    this.id = product.id
    this.name = product.name
    this.age = product.age
}

Product.findById = (id, result) => {
    sql.query(`SELECT * FROM product as pd WHERE pd.id =  ${id}`, (err, res) => {
        if (err) {
            console.log("error: ", err);
            result(err, null);
            return;
        }
      
        if (res.length) {
        console.log("found product: ", res[0]);
        result(null, res[0]);
        return;
        }
      
        // not found Tutorial with the id
        result({ kind: "not_found" }, null);
    })
}


module.exports = Product