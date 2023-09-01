import axios from "axios";
const service = axios.create({
    baseURL: "/apis",
    timeout: 5000  //
});

service.interceptors.request.use(config => {
    config.method === 'post'
        ? config.data = qs.stringify({...config.data})
        : config.params = {...config.params};
    config.headers['Content-Type'] = 'application/x-www-form-urlencoded';

    return config;
}, error => {  //请求错误处理
    Promise.reject(error)
});
service.interceptors.response.use(
    response => {  //成功请求到数据
        //这里根据后端提供的数据进行对应的处理
        if (response.data.result === 'TRUE') {
            return response.data;
        } else {

        }
    },
    error => {  //响应错误处理
        console.log('error');
        console.log(error);
        console.log(JSON.stringify(error));

        let text = JSON.parse(JSON.stringify(error)).response.status === 404
            ? '404'
            : '网络异常，请重试';


        return Promise.reject(error)
    }
);
export default service;
