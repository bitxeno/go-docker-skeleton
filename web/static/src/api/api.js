import request from "@/utils/request";

export default {
  get: function (params) {
    return request({
      url: "/hello",
      method: "get",
      params,
    });
  },
  post: function (data) {
    return request({
      url: "/hello",
      method: "post",
      data,
    });
  },
};
