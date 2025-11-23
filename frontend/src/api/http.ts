import axios from 'axios';
import type { AxiosInstance, AxiosRequestConfig, AxiosResponse, AxiosRequestHeaders } from 'axios';
import { useLoadingStore } from '@src/stores/LoadingStore';
/* 
基于 axios 封装一个 请求库
调用方式为 
export const GetIndexManualData = (data) => {
  return MyFetch({
    url: "/api/front/manual/commission/index",
    data,
    method: "POST",
  });
};

export const GetManualProjectList = (data) => {
  return MyFetch({
    url: "/api/front/manual/item/list",
    data,
    method: "GET",
  });
};

其中 method 不区分大小写
*/

export const GetBaseUrl = (): string => {
  let basteUrl = window.location.origin;

  const hostname = window.location.hostname;
  if (hostname.indexOf('test') > -1) {
    basteUrl = '//test-api.example.com';
  }

  if (hostname.indexOf('example.com') > -1) {
    basteUrl = '//api.example.com';
  }

  return basteUrl;
};

let service: AxiosInstance;

function SetAxiosConfig() {
  const LoadingStore = useLoadingStore();

  service = axios.create();

  service.defaults.timeout = 30000; // 超时 30 秒
  // 请求拦截
  service.interceptors.request.use(
    (config) => {
      LoadingStore.Show({
        info: `正在请求:${config.url}`,
      });
      return config;
    },
    (error) => {
      LoadingStore.Hide();
      return Promise.reject(error);
    },
  );

  // 响应拦截
  service.interceptors.response.use(
    (response) => {
      LoadingStore.Hide();
      return response.data;
    },
    (error) => {
      LoadingStore.Hide();
      return error;
    },
  );
}

type MyFetchOpt = {
  Path: string;
  Method: string;
  Data?: unknown;
  Headers?: AxiosRequestHeaders;
};

const MyFetch = (opt: MyFetchOpt): Promise<AxiosResponse> => {
  const config: AxiosRequestConfig = {};

  const baseURL = GetBaseUrl();
  config.url = `${baseURL}${opt.Path}`;

  const method = opt.Method.toUpperCase();
  if (method === 'POST') {
    config.method = 'post';
    config.data = opt.Data || {};
  }
  if (method === 'GET') {
    config.method = 'get';
    config.params = opt.Data || {};
  }

  config.headers = {
    'Content-Type': 'application/json',
    ...opt.Headers,
  };

  return service.request(config);
};

export { SetAxiosConfig, MyFetch };
