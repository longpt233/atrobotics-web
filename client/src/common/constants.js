export const PageName = {
  HOME_PAGE: 'HomePage',
  LOGIN_PAGE: 'LoginPage',
  REGISTER_PAGE: 'RegisterPage',
  PRODUCT_LIST_PAGE: 'ProductListPage',
  PRODUCT_DETAIL_PAGE: 'ProductDetailPage',
  CHECKOUT_PAGE: 'CheckoutPage',
  ORDER_PAGE: 'OrderPage',
};

export const SUPPORT_LANGUAGE = {
  VI: 'vi',
};

export const DEFAULT_LANGUAGE = SUPPORT_LANGUAGE.VI;

export const HttpStatus = {
  OK: 200,
  BAD_REQUEST: 400,
  UNAUTHORIZED: 401,
  FORBIDDEN: 403,
  NOT_FOUND: 404,
  CONFLICT: 409,
  UNPROCESSABLE_ENTITY: 422,
  GROUP_HAS_CHILDREN: 410,
  GROUP_MAX_LEVEL: 411,
  GROUP_MAX_QUANTITY: 412,
  AWS_ERROR: 413,
  ITEM_NOT_FOUND: 444,
  ITEM_ALREADY_EXIST: 445,
  ITEM_INVALID: 446,
  INTERNAL_SERVER_ERROR: 500,
  SERVICE_UNAVAILABLE: 503,
};

export const DEFAULT_META = [
  {
    property: 'og:title',
    content: 'HoiDapThuoc',
    vmid: 'og:title',
  },
  {
    property: 'og:url',
    content: 'hoidapthuoc.com',
    vmid: 'og:url',
  },
];

export const DATE_TIME_FORMAT = {
  YYYY_MM_DD_HYPHEN: 'YYYY_MM_DD',
  DD_MM_YYYY_DASH: 'DD/MM/YYYY',
  hh_mm_L_COLON: 'h:mm L',
  hh_mm_vi_DD_MM_YYYY_DOT: 'hh:mm [Ngày] DD.MM.YYYY',
  dddd_vi_DD_MM_YYYY_DASH: 'dddd [Ngày] DD/MM/YYYY',
  DD_MM_YY_DASH: 'DD/MM/YYYY',
  dddd_vi_L_SPACE: 'dddd, [ngày] L',
  DD_vi_MM: 'DD [Th]MM',
  DD_vi_M_YYYY: 'DD [Th]M YYYY',
  hh_mm: 'hh:mm',
};

export const FORM_VALIDATION = {
  textMaxLength: 255,
  textAreaMaxLength: 2000,
  passwordMinLength: 6,
  phoneRegExp:
    /^((\\+[1-9]{1,4}[ \\-]*)|(\\([0-9]{2,3}\\)[ \\-]*)|([0-9]{2,4})[ \\-]*)*?[0-9]{3,4}?[ \\-]*[0-9]{3,4}?$/,
  nameRegExp: /^([^!@`~#$:%^*&()<>?\\/\\+|=]+?)$/,
};

export const DeviceType = {
  MOBILE: 'mobile',
  DESKTOP: 'desktop',
};

export const TIME_DELAY_LOADING = 500;
