import axios from 'axios';

async function getProductList() {
  const response = await axios.get(
    'http://atroboticsvn.com/api/v1/user/products?limit=3&offset=1'
  );
  return response.data;
}

export const productList = getProductList();
