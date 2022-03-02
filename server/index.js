const express = require("express")
const route = require("./src/routes")
const app = express()
require("dotenv").config()


app.use(express.json())
app.use(express.urlencoded({ extended: true }))
console.log(process.env.MYSQL_USER)
route(app)
// set port, listen for requests
const PORT = process.env.PORT || 8080
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}.`)
});
