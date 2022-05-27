const productList = {
  label: {
    title: 'Danh sách toàn bộ sản phẩm',
    resultCount: 'kết quả phù hợp',
  },
  filter: {
    sort: {
      title: 'Sắp xếp',
      placeholder: 'Mặc định',
      options: [
        {
          value: 'product_updated_at.desc',
          label: 'Mới nhất',
        },
        {
          value: 'product_sold.desc',
          label: 'Bán chạy nhất',
        },
        {
          value: 'product_price.asc',
          label: 'Giá tăng dần',
        },
        {
          value: 'product_price.desc',
          label: 'Giá giảm dần',
        },
      ],
      optionsLength: '4',
    },
    brand: {
      title: 'Hãng',
      placeholder: 'Tất cả',
    },
  },
};
export default productList;
