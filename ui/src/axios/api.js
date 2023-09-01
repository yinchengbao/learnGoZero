import service from './request'
export const getPersonInfo = data => {
    return service({
        url: '/login?username=admin&password=admin',
        method: 'get',
        data
    })
};
