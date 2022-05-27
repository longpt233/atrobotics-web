const checkout = {
  cartList: {
    header: {
      product: 'Sản phẩm',
      price: 'Đơn giá',
      quantity: 'Số lượng',
      total: 'Thành tiền',
    },
  },
  cartItem: {
    multiple: 'x',
    incQuantity: '+',
    decQuantity: '-',
  },
  cartTotals: {
    title: 'Tổng đơn hàng',
    content: {
      totals: 'Tổng tiền',
    },
    additionalInfo: 'Mức giá trên đã bao gồm thuế và phí vận chuyển.',
    submit: 'Đặt hàng',
  },
  deliveryAddress: {
    title: 'Địa chỉ giao hàng',
    content: {
      noAddress: {
        title: 'Hiện không có thông tin về địa chỉ giao hàng, mời bổ sung thêm địa chỉ.',
        addNewAddress: 'Thêm địa chỉ',
      },
    },
  },
};

export default checkout;
