import capitalize from 'lodash/capitalize';

export const yupVi = {
  mixed: {
    required: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName}`) + ` là trường bắt buộc`;
    },
    default: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} không hợp lệ`);
    },
    selectRequired: `Đây là trường bắt buộc`,
    oneOf: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName}`) + ` là trường bắt buộc`;
    },
    notOneOf: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải là giá trị ngoại trừ: ${ctx.named('values')}`);
    },
    defined: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải được định nghĩa`);
    },
  },

  string: {
    length: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải đúng ${ctx.named('length')} ký tự`);
    },
    min: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải có ít nhất  ${ctx.named('length')} ký tự`);
    },
    max: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} không được vượt quá ${ctx.named('length')} ký tự`);
    },
    matches: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} không hợp lệ`);
    },
    email: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} chưa đúng định dạng`);
    },
    url: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName}`) + ` phải đúng định dạng URL`;
    },
    uuid: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName}`) + ` đúng địng dạng UUID`;
    },
    trim: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải được cắt chuỗi`);
    },
    lowercase: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải là chữ thường`);
    },
    uppercase: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải là chữ hoa`);
    },
  },

  number: {
    min: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải có giá trị thấp nhất là ${ctx.named('min')}`);
    },
    max: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} không được vượt quá ${ctx.named('max')}`);
    },
    lessThan: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải ít hơn ${ctx.named('less')}`);
    },
    moreThan: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải lớn hơn ${ctx.named('more')}`);
    },
    positive: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải là số dương`);
    },
    negative: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải là số âm`);
    },
    integer: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải là số nguyên`);
    },
  },

  date: {
    max: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải nhỏ hơn ngày ${ctx.named('max')}`);
    },
    min: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải lớn hơn ngày ${ctx.named('min')}`);
    },
  },

  boolean: {
    isValue: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải có giá trị là ${ctx.named('value')}`);
    },
  },

  array: {
    min: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải có ít nhất ${ctx.named('min')} phần tử`);
    },
    max: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải ít hơn hoặc bằng ${ctx.named('max')} phần tử`);
    },
    length: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải có ${ctx.named('length')} phần tử`);
    },
  },

  object: {
    noUnknown: (ctx) => {
      const fieldName = ctx.linked(`yupFields.${ctx.named('path')}`);
      return capitalize(`${fieldName} phải có khóa: ${ctx.named('unknown')}`);
    },
  },
};
